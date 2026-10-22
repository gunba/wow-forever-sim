#!/usr/bin/env python3
"""Co-design of talents and rotation: one talent-conditioned policy trained by PPO on many builds at once, and a talent
search fed by the same experience.

Every fight is played by the shared network on a build chosen by the search, so each finished fight is both rotation
experience and one (build, DPS) sample. The search keeps a population of builds scored from those fights, fits an
ensemble surrogate (talent vector -> DPS) on the stream, and assigns the next fights from elites, mutations of elites,
surrogate proposals (predicted DPS plus uncertainty), random builds and the published build. Fights on promising builds
in turn make the network play them better. Shared-seed evaluations rank the leaders.

    python3 tools/forever_rl/codesign.py --build fury --race Orc --minutes 60 --out ~/.cache/forever-sim-audit/rl/runs/fury_codesign
"""
import argparse
import collections
import json
import math
import time
from pathlib import Path

import numpy as np
import torch
import torch.nn as nn
import torch.nn.functional as F

from forever_env import VecEnv
from train import Net


class Talents:
    """Legal talent builds: tier gates (5 points per row in the tree), prerequisites and 51 points."""

    def __init__(self, spec):
        self.max = np.array(spec["talentMax"])
        self.row = np.array(spec["talentRow"])
        self.tree = np.array(spec["talentTree"])
        self.prereq = np.array(spec["talentPrereq"])
        self.sizes = spec["treeSizes"]
        self.n = len(self.max)
        self.names = spec["talentNames"]

    def decode(self, s):
        parts = (s.split("-") + ["", "", ""])[:len(self.sizes)]
        pts = np.zeros(self.n, np.int64)
        off = 0
        for size, part in zip(self.sizes, parts):
            for i, ch in enumerate(part):
                pts[off + i] = int(ch)
            off += size
        return pts

    def encode(self, pts):
        out, off = [], 0
        for size in self.sizes:
            out.append("".join(str(int(x)) for x in pts[off:off + size]).rstrip("0"))
            off += size
        return "-".join(out).rstrip("-")

    def tree_spent(self, pts, t):
        return int(pts[self.tree == t].sum())

    def can_add(self, pts, i):
        if pts[i] >= self.max[i] or pts.sum() >= 51:
            return False
        if self.tree_spent(pts, self.tree[i]) < 5 * self.row[i]:
            return False
        p = self.prereq[i]
        return p < 0 or pts[p] >= self.max[p]

    def valid(self, pts):
        """Every talent with points has 5 points per row above it in its tree and a full prerequisite."""
        if (pts < 0).any() or (pts > self.max).any() or pts.sum() > 51:
            return False
        for i in np.nonzero(pts)[0]:
            same = (self.tree == self.tree[i]) & (self.row < self.row[i])
            if pts[same].sum() < 5 * self.row[i]:
                return False
            p = self.prereq[i]
            if p >= 0 and pts[p] < self.max[p]:
                return False
        return True

    def fill(self, pts, rng, tree_weight=None):
        pts = pts.copy()
        while pts.sum() < 51:
            legal = [i for i in range(self.n) if self.can_add(pts, i)]
            if not legal:
                break
            w = np.ones(len(legal)) if tree_weight is None else np.array([tree_weight[self.tree[i]] for i in legal], float)
            pts[legal[rng.choice(len(legal), p=w / w.sum())]] += 1
        return pts

    def random_build(self, rng):
        w = rng.dirichlet(np.ones(len(self.sizes)) * 0.5)
        return self.fill(np.zeros(self.n, np.int64), rng, w * 10 + 0.01)

    def mutate(self, pts, rng, k):
        pts = pts.copy()
        removed = 0
        for _ in range(k * 6):
            if removed >= k:
                break
            cand = np.nonzero(pts)[0]
            i = cand[rng.integers(len(cand))]
            pts[i] -= 1
            if self.valid(pts):
                removed += 1
            else:
                pts[i] += 1
        w = rng.dirichlet(np.ones(len(self.sizes))) * 3 + 1
        return self.fill(pts, rng, w)


class Builds:
    """A build is its talent string; the search and the evaluation tools work on these."""

    def __init__(self, spec):
        self.tal = Talents(spec)
        self.published = self.canon(spec["baseTalents"])
        self.dim = self.tal.n

    def canon(self, b):
        return self.tal.encode(self.tal.decode(b))

    def points(self, b):
        return self.tal.decode(b)

    def vec(self, b):
        """Surrogate input: talent ranks over their maxima."""
        return (self.tal.decode(b) / self.tal.max).astype(np.float32)

    def mutate(self, b, rng, k):
        """k talent points moved."""
        return self.tal.encode(self.tal.mutate(self.tal.decode(b), rng, k))

    def random(self, rng):
        return self.tal.encode(self.tal.random_build(rng))

    def describe(self, b):
        return b


class Surrogate(nn.Module):
    """An ensemble of small networks: talent ranks -> DPS. The spread across members is the uncertainty."""

    def __init__(self, n, members=5, hidden=128):
        super().__init__()
        self.nets = nn.ModuleList(nn.Sequential(nn.Linear(n, hidden), nn.ReLU(), nn.Linear(hidden, hidden), nn.ReLU(),
                                                nn.Linear(hidden, 1)) for _ in range(members))
        self.mu, self.sd = 0.0, 1.0

    def forward(self, x):
        return torch.stack([m(x).squeeze(-1) for m in self.nets], 0) * self.sd + self.mu

    def fit(self, X, y, steps=300, lr=1e-3):
        self.mu, self.sd = float(y.mean()), float(y.std() + 1e-6)
        opt = torch.optim.Adam(self.parameters(), lr=lr, weight_decay=1e-5)
        yn = (y - self.mu) / self.sd
        for _ in range(steps):
            idx = torch.randint(0, len(X), (len(self.nets), min(4096, len(X))), device=X.device)
            loss = sum(F.mse_loss(m(X[i]).squeeze(-1), yn[i]) for m, i in zip(self.nets, idx))
            opt.zero_grad(); loss.backward(); opt.step()
        return float(loss.item()) / len(self.nets)


def gumbel_sample(logits):
    """One categorical draw per row (masked logits are -1e9) without building a distribution object."""
    u = torch.rand_like(logits).clamp_(1e-12, 1.0)
    return (logits - torch.log(-torch.log(u))).argmax(-1)


class Registry:
    """Build strings <-> ids, with decayed fight statistics per id."""

    def __init__(self):
        self.ids, self.names = {}, []
        self.n = np.zeros(1024); self.s = np.zeros(1024); self.ss = np.zeros(1024); self.seen = np.zeros(1024)

    def id(self, b):
        k = self.ids.get(b)
        if k is None:
            k = self.ids[b] = len(self.names)
            self.names.append(b)
            if k >= len(self.n):
                for name in ("n", "s", "ss", "seen"):
                    a = getattr(self, name)
                    setattr(self, name, np.concatenate([a, np.zeros_like(a)]))
        return k

    def add(self, ids, dps):
        np.add.at(self.n, ids, 1.0); np.add.at(self.s, ids, dps); np.add.at(self.ss, ids, dps * dps)
        np.add.at(self.seen, ids, 1.0)

    def decay(self, f):
        m = len(self.names)
        self.n[:m] *= f; self.s[:m] *= f; self.ss[:m] *= f

    def table(self, min_n=3.0, k=64):
        """(mean, standard error, fights, build) of the k best builds with at least min_n (decayed) fights; the
        per-fight spread is floored at 25 DPS so a few lucky, similar fights are not trusted."""
        m = len(self.names)
        n = self.n[:m]
        ok = np.nonzero(n >= min_n)[0]
        if len(ok) == 0:
            return []
        mean = self.s[ok] / n[ok]
        sd = np.sqrt(np.maximum(self.ss[ok] / n[ok] - mean * mean, 625.0))
        se = sd / np.sqrt(n[ok])
        order = np.argsort(-mean)[:k]
        return [(float(mean[t]), float(se[t]), float(n[ok][t]), self.names[ok[t]]) for t in order]


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--build", default="fury")
    ap.add_argument("--race", default="Orc")
    ap.add_argument("--envs", type=int, default=8192)
    ap.add_argument("--threads", type=int, default=24)
    ap.add_argument("--rollout", type=int, default=128, help="decisions per environment per iteration (~14 s of a melee fight)")
    ap.add_argument("--max-decisions", type=float, default=500e6)
    ap.add_argument("--min-decisions", type=float, default=150e6, help="no early stop before this")
    ap.add_argument("--patience", type=float, default=120e6, help="stop once the best evaluation is this many decisions old")
    ap.add_argument("--minutes", type=float, default=30, help="wall-clock cap")
    ap.add_argument("--bc-iters", type=int, default=5)
    ap.add_argument("--hidden", type=int, default=384)
    ap.add_argument("--lr", type=float, default=1e-4)
    ap.add_argument("--lam", type=float, default=0.97, help="GAE lambda (0.97: ~3.6 s of a melee fight at full fidelity)")
    ap.add_argument("--clip", type=float, default=0.2)
    ap.add_argument("--epochs", type=int, default=4)
    ap.add_argument("--minibatch", type=int, default=16384)
    ap.add_argument("--ent", type=float, default=0.003)
    ap.add_argument("--bc-coef", type=float, default=0.5, help="teacher cross-entropy kept during PPO, fading to 0")
    ap.add_argument("--bc-decay", type=float, default=100e6, help="decisions over which --bc-coef falls to 0")
    ap.add_argument("--anneal", type=float, default=0.9, help="learning rate and entropy bonus fall linearly to (1 - this) of their start over --max-decisions")
    ap.add_argument("--fights-per-build", type=int, default=0, help="fights an environment plays on a build before the search assigns another (0: 4, or 16 below 1,000 decisions a fight)")
    ap.add_argument("--target-kl", type=float, default=0.02)
    ap.add_argument("--poll", type=float, default=0.0, help="0: the agent may act at every engine poll, as the APL does")
    ap.add_argument("--elite", type=int, default=16)
    ap.add_argument("--pool", type=int, default=64)
    ap.add_argument("--mix", default="0.10,0.40,0.30,0.15,0.05", help="published, elite, mutation, surrogate, random shares")
    ap.add_argument("--half-life", type=float, default=100e6, help="decisions over which build statistics halve")
    ap.add_argument("--beta", type=float, default=1.0, help="surrogate uncertainty bonus")
    ap.add_argument("--eval-every", type=float, default=30e6, help="decisions between evaluations")
    ap.add_argument("--eval-fights", type=int, default=512, help="paired fights on the published build per evaluation")
    ap.add_argument("--leader-fights", type=int, default=128)
    ap.add_argument("--leaders", type=int, default=4)
    ap.add_argument("--final-fights", type=int, default=4096)
    ap.add_argument("--finalist-fights", type=int, default=512)
    ap.add_argument("--report-fights", type=int, default=1024)
    ap.add_argument("--seed", type=int, default=1)
    ap.add_argument("--rotation", type=Path, help="teacher and reference rotation (APLRotation JSON) instead of the published APL")
    ap.add_argument("--init", type=Path, help="start from this model and skip behaviour cloning")
    ap.add_argument("--out", type=Path, required=True)
    a = ap.parse_args()
    a.out.mkdir(parents=True, exist_ok=True)
    torch.manual_seed(a.seed)
    torch.distributions.Distribution.set_default_validate_args(False)
    rng = np.random.default_rng(a.seed)
    dev = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    # Two cached simulations per environment (the running build and the next): with thousands of builds in turn, a
    # larger cache only costs memory (pet classes ran out at four).
    cfg = dict(build=a.build, race=a.race, fullChoices=True, pollSeconds=a.poll, rewardScale=1e-5, threads=a.threads,
               seed=1_000_000 * a.seed + 11, rotation=a.rotation.read_text() if a.rotation else "", cachedSims=2)
    env = VecEnv(cfg, a.envs, pinned=dev.type == "cuda")
    spec = env.spec
    B = Builds(spec)
    tal = B.tal
    D, A = spec["obsDim"], spec["actDim"]
    published = B.published
    (a.out / "spec.json").write_text(json.dumps(spec))
    log = (a.out / "log.jsonl").open("a")
    if a.init:
        ck = torch.load(a.init, map_location="cpu", weights_only=False)
        if ck["spec"]["obsDim"] != D or ck["spec"]["actDim"] != A:
            raise SystemExit("--init model spaces differ from this environment")
        a.hidden = ck["net"]["trunk.0.weight"].shape[0]
        a.bc_iters = 0
    net = Net(D, A, a.hidden).to(dev)
    if a.init:
        net.load_state_dict(ck["net"])
    opt = torch.optim.Adam(net.parameters(), lr=a.lr, eps=1e-5)
    sur = Surrogate(B.dim).to(dev)
    t0 = time.time()
    N, R = a.envs, a.rollout

    # ------------------------------------------------------------------ build search
    reg = Registry()
    pub_id = reg.id(published)
    shares = np.cumsum([float(x) for x in a.mix.split(",")])
    vec_cache = {}
    samples_id = np.zeros(400_000, np.int64); samples_dps = np.zeros(400_000); samples_dec = np.zeros(400_000)
    n_samples = 0
    proposals = []
    board = {}  # build -> {"vs_pub": mean network gain over the network on the published build, "se": ...}
    state = {"table": [], "pool": []}

    def vec(k):
        v = vec_cache.get(k)
        if v is None:
            v = vec_cache[k] = B.vec(reg.names[k])
        return v

    def refresh_tables():
        state["table"] = reg.table(3.0, max(a.elite, a.pool))
        el = state["table"][:a.elite]
        names = {e[3] for e in el}
        for b, r in sorted(board.items(), key=lambda kv: -kv[1]["vs_pub"])[:4]:
            k = reg.ids.get(b)
            if b not in names and k is not None and reg.n[k] > 0:
                el.append((reg.s[k] / reg.n[k], r["se"], reg.n[k], b))
        state["elites"] = el
        state["pool"] = state["table"][:a.pool]

    def next_build():
        r = rng.random()
        el = state.get("elites") or []
        if r < shares[0] or not el:
            return published
        if r < shares[1] and state["pool"]:
            pool = state["pool"]
            draws = np.array([m for m, se, n, b in pool]) + rng.standard_normal(len(pool)) * np.array([se for m, se, n, b in pool])
            return pool[int(np.argmax(draws))][3]
        if r < shares[2]:
            b = el[rng.integers(len(el))][3]
            return B.mutate(b, rng, min(1 + int(rng.geometric(0.45)), 8))
        if r < shares[3] and proposals:
            return proposals.pop()
        return B.random(rng)

    def propose(k=256):
        base = [e[3] for e in (state.get("elites") or [])] or [published]
        cands = set()
        for _ in range(600):
            cands.add(B.mutate(base[rng.integers(len(base))], rng, int(rng.integers(1, 7))))
        for _ in range(100):
            cands.add(B.random(rng))
        cands = list(cands)
        X = torch.tensor(np.stack([B.vec(c) for c in cands]), dtype=torch.float32, device=dev)
        with torch.no_grad():
            p = sur(X)
        ucb = (p.mean(0) + a.beta * p.std(0)).cpu().numpy()
        return [cands[i] for i in np.argsort(-ucb)[:k]]

    # ------------------------------------------------------------------ paired evaluation
    ev_n = a.eval_fights + a.leaders * a.leader_fights
    ev = VecEnv({**cfg, "seed": 424_242, "labeledRands": True, "threads": a.threads}, ev_n, pinned=dev.type == "cuda")
    ev_seeds = np.concatenate([np.arange(1, a.eval_fights + 1, dtype=np.int64) * 9_973,
                               np.tile(np.arange(1, a.leader_fights + 1, dtype=np.int64) * 9_973, a.leaders)])

    def play(vec_env, builds, seeds, use_apl, decisions=None):
        """Greedy fights: environment i plays builds[i] on seeds[i]; use_apl[i] plays the published APL instead.
        decisions, if given, receives each fight's decision count."""
        for i, b in enumerate(builds):
            vec_env.set_talents(i, b)
        obs, mask, tch = vec_env.restart(seeds)
        got = np.full(vec_env.n, np.nan)
        while np.isnan(got).any():
            with torch.no_grad():
                if dev.type == "cuda":
                    o = vec_env.obs_t.to(dev, non_blocking=True); m = vec_env.mask_t.to(dev, non_blocking=True).bool()
                else:
                    o = torch.from_numpy(obs); m = torch.from_numpy(mask).bool()
                logits, _ = net(o, m)
            act = logits.argmax(-1).to(torch.int32).cpu().numpy()
            act[use_apl] = tch[use_apl]
            obs, mask, tch, _, done, info = vec_env.step(act)
            new = (done == 1) & np.isnan(got)
            got[new] = info[new, 0]
            if decisions is not None:
                decisions[new] = info[new, 1]
        return got

    # The published APL on the published build, once: the reference every evaluation is paired with.
    ref_dec = np.zeros(ev_n)
    apl_ref = play(ev, [published] * ev_n, ev_seeds, np.ones(ev_n, bool), ref_dec)[:a.eval_fights]
    per_fight_decisions = float(ref_dec.mean())

    def evaluate(leaders):
        builds = [published] * a.eval_fights
        for j in range(a.leaders):
            builds += [leaders[j] if j < len(leaders) else published] * a.leader_fights
        got = play(ev, builds, ev_seeds, np.zeros(ev_n, bool))
        net_pub = got[:a.eval_fights]
        d = net_pub - apl_ref
        out = {"vs_apl": float(d.mean()), "vs_apl_se": float(d.std() / math.sqrt(len(d))), "network": float(net_pub.mean()),
               "apl": float(apl_ref.mean()), "leaders": []}
        for j, b in enumerate(leaders[:a.leaders]):
            g = got[a.eval_fights + j * a.leader_fights: a.eval_fights + (j + 1) * a.leader_fights]
            dd = g - net_pub[:a.leader_fights]
            board[b] = {"vs_pub": float(dd.mean()), "se": float(dd.std() / math.sqrt(len(dd)))}
            out["leaders"].append({"talents": b, **board[b]})
        return out

    # ------------------------------------------------------------------ rollout buffers
    fpb = a.fights_per_build or (4 if per_fight_decisions >= 1000 else 16)
    cur = np.zeros(N, np.int64)
    pending = np.zeros(N, np.int64)
    for i in range(N):
        r = rng.random()
        b = published if r < 0.4 else (B.mutate(published, rng, int(rng.integers(1, 8))) if r < 0.8 else B.random(rng))
        k = reg.id(b)
        env.set_talents(i, b)
        cur[i] = pending[i] = k
    obs, mask, teacher = env.reset()
    # Environments start together; spread them over the fight with the published rotation, so every batch holds states
    # from every part of the fight. Each environment's first (partly APL-played) fight is left out of the statistics.
    obs, mask, teacher = env.warmup(rng.integers(0, max(1, int(per_fight_decisions)), N))
    first = np.ones(N, bool)
    b_obs = torch.zeros(R, N, D, device=dev)
    b_mask = torch.zeros(R, N, A, dtype=torch.bool, device=dev)
    b_act = torch.zeros(R, N, dtype=torch.long, device=dev)
    b_logp = torch.zeros(R, N, device=dev)
    b_val = torch.zeros(R, N, device=dev)
    b_rew = torch.zeros(R, N, device=dev)
    b_done = torch.zeros(R, N, device=dev)
    b_teach = torch.zeros(R, N, dtype=torch.long, device=dev)
    env_fights = np.zeros(N, np.int64)
    totals = {"decisions": 0, "fights": 0}
    recent = collections.deque(maxlen=20000)
    per_fight = []

    def to_dev():
        if dev.type == "cuda":
            return env.obs_t.to(dev, non_blocking=True), env.mask_t.to(dev, non_blocking=True).bool()
        return torch.from_numpy(obs), torch.from_numpy(mask).bool()

    def rollout(use_teacher):
        nonlocal obs, mask, teacher, n_samples, fpb
        for t in range(R):
            o, m = to_dev()
            tch = torch.from_numpy(teacher).to(dev, non_blocking=True).long()
            with torch.no_grad():
                net.norm.update(o)
                logits, v = net(o, m)
                act = tch if use_teacher else gumbel_sample(logits)
                logp = F.log_softmax(logits, -1).gather(1, act[:, None]).squeeze(1)
            b_obs[t], b_mask[t], b_act[t], b_logp[t], b_val[t], b_teach[t] = o, m, act, logp, v, tch
            obs, mask, teacher, rew, done, info = env.step(act.to(torch.int32).cpu().numpy())
            b_rew[t] = torch.from_numpy(rew).to(dev, non_blocking=True)
            b_done[t] = torch.from_numpy(done).to(dev, non_blocking=True).float()
            idx = np.nonzero(done)[0]
            if len(idx):
                counted = ~first[idx]
                first[idx] = False
                cur_idx = idx
                idx = idx[counted]
                cur[cur_idx[~counted]] = pending[cur_idx[~counted]]
            if len(idx):
                dps = info[idx, 0].astype(np.float64)
                fin = cur[idx]
                totals["fights"] += len(idx)
                recent.extend(dps.tolist())
                per_fight.append(info[idx, 1].mean())
                if not use_teacher:
                    reg.add(fin, dps)
                    k = len(idx)
                    pos = (np.arange(k) + n_samples) % len(samples_id)
                    samples_id[pos], samples_dps[pos], samples_dec[pos] = fin, dps, totals["decisions"]
                    n_samples += k
                cur[idx] = pending[idx]
                env_fights[idx] += 1
                for i in idx[env_fights[idx] % fpb == 0]:
                    if not use_teacher:
                        b = next_build()
                    elif rng.random() < 0.4:
                        b = published
                    else:
                        b = B.mutate(published, rng, int(rng.integers(1, 8)))
                    pending[i] = reg.id(b)
                    env.set_talents(int(i), b)
        totals["decisions"] += R * N
        with torch.no_grad():
            o, m = to_dev()
            _, last_v = net(o, m)
        adv = torch.zeros_like(b_rew)
        last = torch.zeros(N, device=dev)
        for t in reversed(range(R)):
            nv = last_v if t == R - 1 else b_val[t + 1]
            nonterm = 1.0 - b_done[t]
            delta = b_rew[t] + nv * nonterm - b_val[t]
            last = delta + a.lam * nonterm * last
            adv[t] = last
        return adv, adv + b_val

    coef = {"ent": a.ent}

    def update(adv, ret, bc, bc_weight=0.0):
        n = R * N
        f = lambda x: x.reshape(n, *x.shape[2:])
        o, m, act, old_logp, ret_f, adv_f, tch = f(b_obs), f(b_mask), f(b_act), f(b_logp), f(ret), f(adv), f(b_teach)
        adv_f = (adv_f - adv_f.mean()) / (adv_f.std() + 1e-8)
        st = collections.defaultdict(float)
        k = 0
        for _ in range(a.epochs):
            ekl, ek = 0.0, 0
            perm = torch.randperm(n, device=dev)
            for s0 in range(0, n, a.minibatch):
                idx = perm[s0:s0 + a.minibatch]
                logits, v = net(o[idx], m[idx])
                lp_all = F.log_softmax(logits, -1)
                lp = lp_all.gather(1, act[idx][:, None]).squeeze(1)
                ent = -(lp_all.exp() * lp_all.masked_fill(~m[idx], 0)).sum(-1)
                vloss = F.mse_loss(v, ret_f[idx])
                if bc:
                    loss = F.cross_entropy(logits, tch[idx]) + 0.5 * vloss
                    st["bc_acc"] += (logits.argmax(-1) == tch[idx]).float().mean().item()
                else:
                    ratio = (lp - old_logp[idx]).exp()
                    pl = -torch.min(ratio * adv_f[idx], ratio.clamp(1 - a.clip, 1 + a.clip) * adv_f[idx]).mean()
                    loss = pl + 0.5 * vloss - coef["ent"] * ent.mean()
                    if bc_weight > 0:
                        loss = loss + bc_weight * F.cross_entropy(logits, tch[idx])
                    kl = (old_logp[idx] - lp).mean().item()
                    st["kl"] += kl; ekl += kl; ek += 1
                opt.zero_grad(); loss.backward(); nn.utils.clip_grad_norm_(net.parameters(), 0.5); opt.step()
                st["vloss"] += vloss.item(); st["ent"] += ent.mean().item()
                k += 1
            if not bc and a.target_kl and ek and ekl / ek > 1.5 * a.target_kl:
                break
        return {key: round(v / k, 4) for key, v in st.items()}

    # ------------------------------------------------------------------ training
    best = {"score": -1e9, "state": None, "decisions": 0, "eval": None}
    it, last_eval, last_sur, ppo_start = 0, 0, time.time(), None
    decay_f = 0.5 ** (R * N / a.half_life)
    stop = "budget"
    while True:
        it += 1
        bc = it <= a.bc_iters
        t1 = time.time()
        frac = min(1.0, totals["decisions"] / a.max_decisions)
        for g in opt.param_groups:
            g["lr"] = a.lr * (1 - a.anneal * frac)
        coef["ent"] = a.ent * (1 - a.anneal * frac)
        adv, ret = rollout(bc)
        if not bc and ppo_start is None:
            ppo_start = totals["decisions"]
        bc_weight = 0.0 if bc else a.bc_coef * max(0.0, 1 - (totals["decisions"] - ppo_start) / a.bc_decay)
        st = update(adv, ret, bc, bc_weight)
        row = {"iter": it, "phase": "bc" if bc else "ppo", "elapsed": round(time.time() - t0, 1), "decisions": totals["decisions"],
               "fights": totals["fights"], "dps": round(float(np.mean(recent)), 1) if recent else None, "builds_seen": len(reg.names),
               "secs": round(time.time() - t1, 2), "decisions_per_s": round(R * N / (time.time() - t1)), **st}
        if not bc:
            reg.decay(decay_f)
            refresh_tables()
            if time.time() - last_sur > 60 and n_samples > 5000:
                m_ = min(n_samples, len(samples_id))
                keep = samples_dec[:m_] >= totals["decisions"] - 3 * a.half_life
                ids = samples_id[:m_][keep]
                X = torch.tensor(np.stack([vec(int(k)) for k in ids]), dtype=torch.float32, device=dev)
                y = torch.tensor(samples_dps[:m_][keep], dtype=torch.float32, device=dev)
                row["sur_loss"] = round(sur.fit(X, y, steps=200), 4)
                proposals[:] = propose()
                last_sur = time.time()
            if totals["decisions"] - last_eval >= a.eval_every:
                last_eval = totals["decisions"]
                leaders = []
                for b, rr in sorted(board.items(), key=lambda kv: -kv[1]["vs_pub"])[:2]:
                    leaders.append(b)
                for e in state.get("elites") or []:
                    if e[3] != published and e[3] not in leaders:
                        leaders.append(e[3])
                for p_ in proposals[:2]:
                    if p_ not in leaders:
                        leaders.append(p_)
                res = evaluate(leaders[:a.leaders])
                row["eval"] = res
                if res["vs_apl"] > best["score"]:
                    best.update(score=res["vs_apl"], decisions=totals["decisions"], eval=res,
                                state={k: v.detach().clone() for k, v in net.state_dict().items()})
                    torch.save({"net": best["state"], "spec": spec, "decisions": totals["decisions"], "eval": res}, a.out / "best.pt")
        log.write(json.dumps(row) + "\n"); log.flush()
        if it % 10 == 0 or "eval" in row:
            print(json.dumps({k: v for k, v in row.items() if k != "eval"}), flush=True)
            if "eval" in row:
                e = row["eval"]
                print(f"   eval: network {e['network']:.1f} vs APL {e['apl']:.1f}: {e['vs_apl']:+.2f} ± {e['vs_apl_se']:.2f}"
                      f" (best {best['score']:+.2f})", flush=True)
        if totals["decisions"] >= a.max_decisions:
            break
        if (totals["decisions"] >= a.min_decisions and best["state"] is not None
                and totals["decisions"] - best["decisions"] >= a.patience):
            stop = "converged"
            break
        if time.time() - t0 > a.minutes * 60:
            stop = "time"
            break
    train_minutes = (time.time() - t0) / 60
    torch.save({"net": net.state_dict(), "spec": spec}, a.out / "last.pt")
    if best["state"] is not None:
        net.load_state_dict(best["state"])
    env.close()
    ev.close()

    # ------------------------------------------------------------------ final: talents, then a fresh paired confirmation
    finalists = [b for b, r in sorted(board.items(), key=lambda kv: -kv[1]["vs_pub"])[:8]]
    for e in (state.get("elites") or [])[:6]:
        if e[3] not in finalists:
            finalists.append(e[3])
    finalists = [b for b in finalists if b != published][:12]
    F_ = a.finalist_fights
    fin_env = VecEnv({**cfg, "seed": 515_151, "labeledRands": True}, F_ * (1 + len(finalists)), pinned=dev.type == "cuda")
    seeds = np.tile(np.arange(1, F_ + 1, dtype=np.int64) * 7_919, 1 + len(finalists))
    got = play(fin_env, [published] * F_ + [b for b in finalists for _ in range(F_)], seeds, np.zeros(fin_env.n, bool))
    fin_env.close()
    base = got[:F_]
    fin_rows = []
    for j, b in enumerate(finalists):
        d = got[F_ * (j + 1):F_ * (j + 2)] - base
        fin_rows.append({"talents": b, "vs_published_talents": float(d.mean()), "se": float(d.std() / math.sqrt(F_))})
    fin_rows.sort(key=lambda r: -r["vs_published_talents"])
    best_talents = published
    if fin_rows and fin_rows[0]["vs_published_talents"] > 2.5 * fin_rows[0]["se"]:
        best_talents = fin_rows[0]["talents"]
    G = a.final_fights
    rows = [(published, True), (published, False)] + ([(best_talents, False)] if best_talents != published else [])
    conf_env = VecEnv({**cfg, "seed": 616_161, "labeledRands": True}, G * len(rows), pinned=dev.type == "cuda")
    seeds = np.tile(31_337 + np.arange(G, dtype=np.int64), len(rows))
    got = play(conf_env, [b for b, _ in rows for _ in range(G)], seeds, np.repeat([u for _, u in rows], G))
    conf_env.close()
    apl = got[:G]
    confirm = {"fights": G, "apl": float(apl.mean()), "apl_se": float(apl.std() / math.sqrt(G)), "rows": []}
    for j, (b, use_apl) in enumerate(rows[1:], start=1):
        g = got[G * j:G * (j + 1)]
        d = g - apl
        confirm["rows"].append({"talents": b, "player": "network", "dps": float(g.mean()), "vs_apl": float(d.mean()),
                                "vs_apl_se": float(d.std() / math.sqrt(G))})
    torch.save({"net": net.state_dict(), "spec": spec, "best_talents": best_talents}, a.out / "model.pt")
    summary = {"build": a.build, "race": a.race, "stop": stop, "train_minutes": round(train_minutes, 1),
               "minutes": round((time.time() - t0) / 60, 1), "iterations": it, "decisions": totals["decisions"],
               "fights": totals["fights"], "builds_seen": len(reg.names), "best_decisions": best["decisions"],
               "decisions_per_fight": per_fight_decisions,
               "best_eval": best["eval"], "published_talents": published, "best_talents": best_talents,
               "finalists": fin_rows, "confirm": confirm}
    (a.out / "summary.json").write_text(json.dumps(summary, indent=1))
    print(json.dumps({k: v for k, v in summary.items() if k not in ("finalists",)}, indent=1), flush=True)


if __name__ == "__main__":
    main()
