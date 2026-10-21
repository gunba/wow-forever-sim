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


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--build", default="fury")
    ap.add_argument("--race", default="Orc")
    ap.add_argument("--envs", type=int, default=2048)
    ap.add_argument("--threads", type=int, default=12)
    ap.add_argument("--rollout", type=int, default=64)
    ap.add_argument("--minutes", type=float, default=60)
    ap.add_argument("--bc-iters", type=int, default=40)
    ap.add_argument("--hidden", type=int, default=384)
    ap.add_argument("--lr", type=float, default=1e-4)
    ap.add_argument("--lam", type=float, default=0.95)
    ap.add_argument("--clip", type=float, default=0.2)
    ap.add_argument("--epochs", type=int, default=4)
    ap.add_argument("--minibatch", type=int, default=16384)
    ap.add_argument("--ent", type=float, default=0.003)
    ap.add_argument("--bc-coef", type=float, default=0.5, help="teacher cross-entropy kept during PPO (the published APL plays any build)")
    ap.add_argument("--bc-decay", type=int, default=800, help="PPO iterations over which --bc-coef falls to 0")
    ap.add_argument("--fights-per-build", type=int, default=4, help="fights an environment plays on a build before the search assigns another")
    ap.add_argument("--eval-cands", type=int, default=12, help="builds in each periodic shared-seed evaluation")
    ap.add_argument("--target-kl", type=float, default=0.02)
    ap.add_argument("--poll", type=float, default=0.25)
    ap.add_argument("--elite", type=int, default=16, help="builds the elite share samples from")
    ap.add_argument("--mix", default="0.10,0.40,0.30,0.15,0.05", help="published, elite, mutation, surrogate, random shares")
    ap.add_argument("--decay", type=float, default=0.9995, help="per-iteration decay of build statistics (the policy improves)")
    ap.add_argument("--pool", type=int, default=64, help="builds the elite share draws from (Thompson sampling)")
    ap.add_argument("--beta", type=float, default=1.0, help="surrogate uncertainty bonus")
    ap.add_argument("--eval-every", type=int, default=50)
    ap.add_argument("--eval-fights", type=int, default=96)
    ap.add_argument("--final-fights", type=int, default=1024)
    ap.add_argument("--seed", type=int, default=1)
    ap.add_argument("--rotation", type=Path, help="teacher and reference rotation (APLRotation JSON) instead of the published APL")
    ap.add_argument("--init", type=Path, help="start from this model and skip behaviour cloning")
    ap.add_argument("--anneal", type=float, default=0.9, help="learning rate and entropy bonus fall linearly to (1 - this) of their start")
    ap.add_argument("--out", type=Path, required=True)
    a = ap.parse_args()
    a.out.mkdir(parents=True, exist_ok=True)
    torch.manual_seed(a.seed)
    rng = np.random.default_rng(a.seed)
    dev = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    cfg = dict(build=a.build, race=a.race, fullChoices=True, pollSeconds=a.poll, rewardScale=1e-5, threads=a.threads,
               seed=1_000_000 * a.seed + 11, rotation=a.rotation.read_text() if a.rotation else "")
    env = VecEnv(cfg, a.envs)
    spec = env.spec
    B = Builds(spec)
    tal = B.tal
    D, A, T = spec["obsDim"], spec["actDim"], spec["talents"]
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
    best = {"score": -1e9, "state": None, "iter": 0}
    sur = Surrogate(B.dim).to(dev)
    t0 = time.time()

    # ---------------------------------------------------------------- build search state
    stats = {}  # build -> [n, sum, sumsq] (decayed)
    seen = collections.Counter()
    samples = collections.deque(maxlen=200_000)  # (talent ranks / max, DPS, iteration)
    proposals = []
    shares = np.cumsum([float(x) for x in a.mix.split(",")])

    def add_stat(b, dps):
        s = stats.setdefault(b, [0.0, 0.0, 0.0])
        s[0] += 1; s[1] += dps; s[2] += dps * dps
        seen[b] += 1

    def decay():
        for b in list(stats):
            s = stats[b]
            s[0] *= a.decay; s[1] *= a.decay; s[2] *= a.decay
            if s[0] < 0.5:
                del stats[b]
        cache.clear()

    cache = {}

    def score_table(min_n=4.0):
        """(mean, standard error, fights, build), best first; per-fight spread floored at 25 DPS so a build with a
        few lucky, similar fights is not trusted. Recomputed once per iteration."""
        key = ("table", min_n)
        if key not in cache:
            out = []
            for b, (n, sm, sq) in stats.items():
                if n >= min_n:
                    mean = sm / n
                    sd = math.sqrt(max(sq / n - mean * mean, 625.0))
                    out.append((mean, sd / math.sqrt(n), n, b))
            out.sort(reverse=True)
            cache[key] = out
        return cache[key]

    def elites(k=None):
        """Training-fight leaders by mean (at least ~3 fights), plus the builds that led a shared-seed evaluation."""
        el = score_table(min_n=3.0)[:k or a.elite]
        names = {e[3] for e in el}
        for b, r in sorted(board.items(), key=lambda kv: -kv[1]["vs_ref"])[:4]:
            if b not in names and b in stats and stats[b][0] > 0:
                n, sm, sq = stats[b]
                el.append((sm / n, r["vs_ref_se"], n, b))
        return el

    def propose(k=256):
        """Mutations of the elites (1 to 6 points) and random builds, ranked by surrogate mean + beta * spread."""
        base = [e[3] for e in elites()] or [published]
        cands = set()
        for _ in range(600):
            b = base[rng.integers(len(base))]
            cands.add(B.mutate(b, rng, int(rng.integers(1, 7))))
        for _ in range(100):
            cands.add(B.random(rng))
        cands = list(cands)
        X = torch.tensor(np.stack([B.vec(c) for c in cands]), dtype=torch.float32, device=dev)
        with torch.no_grad():
            p = sur(X)
        ucb = (p.mean(0) + a.beta * p.std(0)).cpu().numpy()
        order = np.argsort(-ucb)
        return [cands[i] for i in order[:k]]

    def next_build():
        r = rng.random()
        el = elites()
        if r < shares[0] or not el:
            return published
        if r < shares[1]:
            # Thompson sampling: builds with few fights and a high mean get more fights until they hold or drop.
            pool = elites(a.pool)
            draws = [m + rng.standard_normal() * se for m, se, n, b in pool]
            return pool[int(np.argmax(draws))][3]
        if r < shares[2]:
            b = el[rng.integers(len(el))][3]
            k = min(1 + rng.geometric(0.45), 8)
            return B.mutate(b, rng, k)
        if r < shares[3] and proposals:
            return proposals.pop()
        return B.random(rng)

    # ---------------------------------------------------------------- shared-seed evaluation of builds
    def evaluate(builds, fights, chunk=2048, teacher_first=False, ref_idx=0, apl=None, seed0=9_973):
        """Greedy fights on shared seeds with one random stream per roll label: build j's fight f always uses seed f,
        so differences are paired. With teacher_first, builds[0] is played by the published APL (the reference)
        instead of the network; apl[j] = True plays build j with the published APL."""
        apl = list(apl) if apl is not None else [teacher_first and j == 0 for j in range(len(builds))]
        per = max(1, chunk // fights)
        ev = VecEnv({**cfg, "seed": 424_242, "labeledRands": True}, per * fights)
        seeds = np.tile(np.arange(1, fights + 1, dtype=np.int64) * seed0, per)
        dps = np.full(len(builds) * fights, np.nan)
        for c in range(0, len(builds), per):
            group = builds[c:c + per] + [builds[c]] * (per - len(builds[c:c + per]))
            for j, b in enumerate(group):
                for f in range(fights):
                    ev.set_talents(j * fights + f, b)
            obs, mask, tch = ev.restart(seeds)
            got = np.full(per * fights, np.nan)
            use_t = np.zeros(per * fights, bool)
            for j in range(len(group)):
                if c + j < len(builds) and apl[c + j]:
                    use_t[j * fights:(j + 1) * fights] = True
            for _ in range(100_000):
                with torch.no_grad():
                    logits, _ = net(torch.from_numpy(obs).to(dev), torch.from_numpy(mask).to(dev).bool())
                act = logits.argmax(-1).cpu().numpy().astype(np.int32)
                act[use_t] = tch[use_t]
                obs, mask, tch, _, done, info = ev.step(act)
                new = (done == 1) & np.isnan(got)
                got[new] = info[new, 0]
                if not np.isnan(got).any():
                    break
            k = min(per, len(builds) - c)
            dps[c * fights:(c + k) * fights] = got[:k * fights]
        ev.close()
        res = []
        base = dps[ref_idx * fights:(ref_idx + 1) * fights]
        for j, b in enumerate(builds):
            d = dps[j * fights:(j + 1) * fights]
            row = {"talents": b, "dps": float(d.mean()), "se": float(d.std() / math.sqrt(fights))}
            if j != ref_idx:
                diff = d - base
                row["vs_ref"] = float(diff.mean())
                row["vs_ref_se"] = float(diff.std() / math.sqrt(fights))
            res.append(row)
        return res

    # ---------------------------------------------------------------- behaviour cloning on varied builds
    for i in range(a.envs):
        r = rng.random()
        b = published if r < 0.4 else (B.mutate(published, rng, int(rng.integers(1, 8))) if r < 0.8 else B.random(rng))
        env.set_talents(i, b)
    obs, mask, teacher = env.reset()
    N, R = a.envs, a.rollout
    b_obs = torch.zeros(R, N, D, device=dev)
    b_mask = torch.zeros(R, N, A, dtype=torch.bool, device=dev)
    b_act = torch.zeros(R, N, dtype=torch.long, device=dev)
    b_logp = torch.zeros(R, N, device=dev)
    b_val = torch.zeros(R, N, device=dev)
    b_rew = torch.zeros(R, N, device=dev)
    b_done = torch.zeros(R, N, device=dev)
    b_teach = torch.zeros(R, N, dtype=torch.long, device=dev)
    totals = collections.Counter()
    recent = collections.deque(maxlen=4000)
    env_fights = np.zeros(a.envs, np.int64)
    board = {}  # build -> latest shared-seed result against the published build

    def rollout(use_teacher, it):
        nonlocal obs, mask, teacher
        for t in range(R):
            o = torch.from_numpy(obs).to(dev)
            m = torch.from_numpy(mask).to(dev).bool()
            tch = torch.from_numpy(teacher.astype(np.int64)).to(dev)
            with torch.no_grad():
                net.norm.update(o)
                logits, v = net(o, m)
                dist = torch.distributions.Categorical(logits=logits)
                act = tch if use_teacher else dist.sample()
                logp = dist.log_prob(act)
            b_obs[t], b_mask[t], b_act[t], b_logp[t], b_val[t], b_teach[t] = o, m, act, logp, v, tch
            obs, mask, teacher, rew, done, info = env.step(act.cpu().numpy().astype(np.int32))
            b_rew[t] = torch.from_numpy(rew).to(dev)
            b_done[t] = torch.from_numpy(done.astype(np.float32)).to(dev)
            for i in np.nonzero(done)[0]:
                b = B.canon(env.talents(int(i)))
                dps = float(info[i, 0])
                totals["fights"] += 1
                recent.append(dps)
                if not use_teacher:
                    add_stat(b, dps)
                    samples.append((B.vec(b), dps, it))
                env_fights[i] += 1
                if env_fights[i] % a.fights_per_build:
                    continue
                if not use_teacher:
                    nxt = next_build()
                elif rng.random() < 0.4:
                    nxt = published
                else:
                    nxt = B.mutate(published, rng, int(rng.integers(1, 8)))
                env.set_talents(int(i), nxt)
        totals["decisions"] += R * N
        with torch.no_grad():
            _, last_v = net(torch.from_numpy(obs).to(dev), torch.from_numpy(mask).to(dev).bool())
        adv = torch.zeros_like(b_rew)
        last = torch.zeros(N, device=dev)
        for t in reversed(range(R)):
            nv = last_v if t == R - 1 else b_val[t + 1]
            nonterm = 1.0 - b_done[t]
            delta = b_rew[t] + nv * nonterm - b_val[t]
            last = delta + a.lam * nonterm * last
            adv[t] = last
        return adv, adv + b_val

    ent_coef = a.ent

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
            for s in range(0, n, a.minibatch):
                idx = perm[s:s + a.minibatch]
                logits, v = net(o[idx], m[idx])
                lp_all = F.log_softmax(logits, -1)
                lp = lp_all.gather(1, act[idx][:, None]).squeeze(1)
                ent = -(lp_all.exp() * lp_all.masked_fill(~m[idx], 0)).sum(-1)
                vloss = F.mse_loss(v, ret_f[idx])
                if bc:
                    pl = F.cross_entropy(logits, tch[idx])
                    loss = pl + 0.5 * vloss
                    st["bc_acc"] += (logits.argmax(-1) == tch[idx]).float().mean().item()
                else:
                    ratio = (lp - old_logp[idx]).exp()
                    pl = -torch.min(ratio * adv_f[idx], ratio.clamp(1 - a.clip, 1 + a.clip) * adv_f[idx]).mean()
                    loss = pl + 0.5 * vloss - ent_coef * ent.mean()
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

    it = 0
    history = []
    while time.time() - t0 < a.minutes * 60:
        it += 1
        bc = it <= a.bc_iters
        t1 = time.time()
        frac = min(1.0, (time.time() - t0) / (a.minutes * 60))
        for g in opt.param_groups:
            g["lr"] = a.lr * (1 - a.anneal * frac)
        ent_coef = a.ent * (1 - a.anneal * frac)
        adv, ret = rollout(bc, it)
        bc_weight = 0.0 if bc else a.bc_coef * max(0.0, 1 - (it - a.bc_iters) / max(1, a.bc_decay))
        st = update(adv, ret, bc, bc_weight)
        if not bc:
            decay()
            if it % 10 == 0 and len(samples) > 2000:
                recent_s = [s for s in samples if s[2] >= it - 1500]
                X = torch.tensor(np.stack([s[0] for s in recent_s]), dtype=torch.float32, device=dev)
                y = torch.tensor([s[1] for s in recent_s], dtype=torch.float32, device=dev)
                st["sur_loss"] = round(sur.fit(X, y), 4)
                proposals[:] = propose()
        row = {"iter": it, "phase": "bc" if bc else "ppo", "elapsed": round(time.time() - t0, 1), "decisions": totals["decisions"],
               "fights": totals["fights"], "dps": round(float(np.mean(recent)), 1) if recent else None, "builds_seen": len(seen),
               "secs": round(time.time() - t1, 2), **st}
        el = elites()
        if el:
            row["top"] = [(round(m, 1), round(se, 1), round(n, 1), b) for m, se, n, b in el[:3]]
        if not bc and it % a.eval_every == 0:
            cand = [published, published]
            for e in el:
                if e[3] not in cand and len(cand) < a.eval_cands - 3:
                    cand.append(e[3])
            for prop in proposals[:6]:
                if prop not in cand and len(cand) < a.eval_cands:
                    cand.append(prop)
            # Reference (index 1): the network on the published build; index 0 is the published APL on it.
            res = evaluate(cand, a.eval_fights, teacher_first=True, ref_idx=1)
            res[0]["talents"] = "APL on " + published
            for r in res[2:]:
                board[r["talents"]] = r
            row["eval"] = res
            history.append(res)
            torch.save({"net": net.state_dict(), "spec": spec}, a.out / "model.pt")
            # The network on the published build, against the APL on the same fights.
            score = -res[0]["vs_ref"]
            if score > best["score"]:
                best.update(score=score, iter=it, state={k: v.detach().clone() for k, v in net.state_dict().items()})
                torch.save({"net": best["state"], "spec": spec, "iter": it, "score": score}, a.out / "best.pt")
        log.write(json.dumps(row) + "\n"); log.flush()
        print(json.dumps({k: v for k, v in row.items() if k not in ("top", "eval")}), flush=True)
        if "eval" in row:
            for r in row["eval"]:
                print("   eval", json.dumps(r), flush=True)

    # Final: the best checkpoint (by the periodic evaluations) plays the leaders, against the published build.
    torch.save({"net": net.state_dict(), "spec": spec}, a.out / "last.pt")
    if best["state"] is not None:
        net.load_state_dict(best["state"])
    el = elites()
    lead = [b for b, r in sorted(board.items(), key=lambda kv: -kv[1]["vs_ref"])[:8]]
    lead += [e[3] for e in el[:6] if e[3] not in lead]
    cand = [published] + [b for b in lead if b != published]
    # Every finalist played by the network and by the published APL, on the same fresh seeds; reference: the
    # published APL on the published build.
    res = evaluate(cand + cand, a.final_fights, apl=[False] * len(cand) + [True] * len(cand), ref_idx=len(cand), seed0=7_919)
    final = []
    for j, b in enumerate(cand):
        net_r, apl_r = res[j], res[len(cand) + j]
        final.append({"talents": b, "network": round(net_r["dps"], 1), "network_vs_apl": round(net_r["vs_ref"], 2),
                      "network_se": round(net_r["vs_ref_se"], 2), "apl": round(apl_r["dps"], 1),
                      "apl_vs_apl_on_published": round(apl_r.get("vs_ref", 0.0), 2), "apl_se": round(apl_r.get("vs_ref_se", 0.0), 2)})
    final.sort(key=lambda r: -max(r["network"], r["apl"]))
    # The surrogate's view: predicted DPS change per talent point at the best build.
    x = torch.tensor(B.vec(final[0]["talents"]), dtype=torch.float32, device=dev, requires_grad=True)
    sur(x[None]).mean().backward()
    grads = x.grad.cpu().numpy() / tal.max
    weights = {tal.names[i]: round(float(grads[i]), 2) for i in np.argsort(-np.abs(grads))[:20]}
    board_top = sorted(board.values(), key=lambda r: -r["vs_ref"])[:20]
    table = [{"talents": b, "mean": round(m, 1), "se": round(se, 1), "n": round(n, 1)} for m, se, n, b in score_table()[:30]]
    torch.save({"net": net.state_dict(), "spec": spec}, a.out / "model.pt")
    summary = {"build": a.build, "race": a.race, "minutes": round((time.time() - t0) / 60, 1), "iterations": it,
               "best_iter": best["iter"], "best_eval_vs_apl": round(best["score"], 2),
               "decisions": totals["decisions"], "fights": totals["fights"], "builds_seen": len(seen), "final": final,
               "surrogate_points": weights, "board": board_top, "training_table": table}
    (a.out / "summary.json").write_text(json.dumps(summary, indent=1))
    print(json.dumps(summary, indent=1), flush=True)
    env.close()


if __name__ == "__main__":
    main()
