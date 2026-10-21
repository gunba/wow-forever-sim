#!/usr/bin/env python3
"""Train one agent to choose talents and play the benchmark fight (masked PPO), optionally warm-started by
behaviour cloning of the published rotation and talents.

    python3 tools/forever_rl/train.py --build fire --race Human --learn-talents --minutes 5 --out ~/.cache/forever-sim-audit/rl/runs/fire

Every iteration appends one JSON line to <out>/log.jsonl (decisions, fights, DPS, entropy, timings); checkpoints go
to <out>/model.pt. The environment is the real engine (tools/forever_rl/env), built as libforeverrl.so.
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


class RunningNorm(nn.Module):
    def __init__(self, dim):
        super().__init__()
        self.register_buffer("mean", torch.zeros(dim))
        self.register_buffer("var", torch.ones(dim))
        self.register_buffer("count", torch.tensor(1e-4))

    @torch.no_grad()
    def update(self, x):
        b_mean, b_var, b_n = x.mean(0), x.var(0, unbiased=False), x.shape[0]
        delta, tot = b_mean - self.mean, self.count + b_n
        self.mean += delta * b_n / tot
        self.var = (self.var * self.count + b_var * b_n + delta ** 2 * self.count * b_n / tot) / tot
        self.count = tot

    def forward(self, x):
        return ((x - self.mean) / torch.sqrt(self.var + 1e-6)).clamp(-10, 10)


class Net(nn.Module):
    def __init__(self, obs_dim, act_dim, hidden):
        super().__init__()
        self.norm = RunningNorm(obs_dim)
        self.trunk = nn.Sequential(nn.Linear(obs_dim, hidden), nn.ReLU(), nn.Linear(hidden, hidden), nn.ReLU(),
                                   nn.Linear(hidden, hidden), nn.ReLU())
        self.pi = nn.Linear(hidden, act_dim)
        self.v = nn.Linear(hidden, 1)
        nn.init.orthogonal_(self.pi.weight, 0.01)
        nn.init.zeros_(self.pi.bias)

    def forward(self, obs, mask):
        h = self.trunk(self.norm(obs))
        logits = self.pi(h).masked_fill(~mask, -1e9)
        return logits, self.v(h).squeeze(-1)


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--build", default="fire")
    ap.add_argument("--race", default="Human")
    ap.add_argument("--learn-talents", action="store_true")
    ap.add_argument("--talents", default="", help="fixed talents (default: the published ones)")
    ap.add_argument("--full", action="store_true", help="full choices: cooldowns, unqueue, longer waits")
    ap.add_argument("--poll", type=float, default=0.0, help="event-driven polling after a wait (seconds; 0: every 50 ms)")
    ap.add_argument("--init", type=Path, help="start from this model (codesign or train), e.g. a talent-conditioned policy")
    ap.add_argument("--fights-per-build", type=int, default=1)
    ap.add_argument("--envs", type=int, default=512)
    ap.add_argument("--threads", type=int, default=8)
    ap.add_argument("--rollout", type=int, default=128, help="decisions per environment per iteration")
    ap.add_argument("--minutes", type=float, default=5, help="wall-clock budget (BC included)")
    ap.add_argument("--max-decisions", type=float, default=0, help="stop after this many environment decisions (0: no limit)")
    ap.add_argument("--bc-iters", type=int, default=20, help="behaviour-cloning iterations before PPO (0: none)")
    ap.add_argument("--hidden", type=int, default=256)
    ap.add_argument("--lr", type=float, default=3e-4)
    ap.add_argument("--gamma", type=float, default=1.0, help="per fight decision (1: total damage, the benchmark objective; a per-decision discount favours long casts); talent decisions are undiscounted")
    ap.add_argument("--lam", type=float, default=0.95)
    ap.add_argument("--clip", type=float, default=0.2)
    ap.add_argument("--epochs", type=int, default=4)
    ap.add_argument("--minibatch", type=int, default=8192)
    ap.add_argument("--ent", type=float, default=0.01)
    ap.add_argument("--talent-ent", type=float, default=0.02, help="entropy bonus on talent decisions")
    ap.add_argument("--vf", type=float, default=0.5)
    ap.add_argument("--bc-coef", type=float, default=0.0, help="weight of the teacher cross-entropy kept during PPO")
    ap.add_argument("--bc-decay", type=int, default=0, help="PPO iterations over which --bc-coef falls linearly to 0 (0: constant)")
    ap.add_argument("--target-kl", type=float, default=0.02, help="stop an update's epochs past 1.5x this KL (0: off)")
    ap.add_argument("--reward-scale", type=float, default=1e-5)
    ap.add_argument("--seed", type=int, default=1)
    ap.add_argument("--eval-fights", type=int, default=256)
    ap.add_argument("--eval-every", type=int, default=10, help="iterations between greedy evaluations (0: end only)")
    ap.add_argument("--out", type=Path, required=True)
    a = ap.parse_args()
    a.out.mkdir(parents=True, exist_ok=True)
    torch.manual_seed(a.seed)
    np.random.seed(a.seed)
    dev = torch.device("cuda" if torch.cuda.is_available() else "cpu")

    cfg = dict(build=a.build, race=a.race, learnTalents=a.learn_talents, fightsPerBuild=a.fights_per_build,
               seed=1_000_000 * a.seed + 17, rewardScale=a.reward_scale, threads=a.threads, fullChoices=a.full,
               pollSeconds=a.poll, talents=a.talents)
    env = VecEnv(cfg, a.envs)
    spec = env.spec
    D, A, T = spec["obsDim"], spec["actDim"], spec["talents"]
    init = None
    if a.init:
        init = torch.load(a.init, map_location="cpu", weights_only=False)
        if init["spec"]["obsDim"] != D or init["spec"]["actDim"] != A:
            raise SystemExit("--init model spaces differ from this environment")
        a.hidden = init["net"]["trunk.0.weight"].shape[0]
    net = Net(D, A, a.hidden).to(dev)
    if init:
        net.load_state_dict(init["net"])
    opt = torch.optim.Adam(net.parameters(), lr=a.lr, eps=1e-5)
    # Evaluation fights use one random stream per roll label, so the network and the published APL meet the same luck.
    eval_env = VecEnv({**cfg, "seed": 987_654_321, "labeledRands": True}, a.eval_fights)
    eval_seeds = np.arange(1, a.eval_fights + 1, dtype=np.int64) * 7_919 + 3
    (a.out / "spec.json").write_text(json.dumps(spec))
    log = (a.out / "log.jsonl").open("a")
    t_start = time.time()
    totals = collections.Counter()
    ep_dps, ep_builds = collections.deque(maxlen=2000), collections.Counter()

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

    def rollout(use_teacher):
        nonlocal obs, mask, teacher
        t_env = t_inf = 0.0
        for t in range(R):
            t0 = time.time()
            o = torch.from_numpy(obs).to(dev, non_blocking=True)
            m = torch.from_numpy(mask).to(dev, non_blocking=True).bool()
            tch = torch.from_numpy(teacher.astype(np.int64)).to(dev)
            with torch.no_grad():
                net.norm.update(o)
                logits, v = net(o, m)
                dist = torch.distributions.Categorical(logits=logits)
                act = tch if use_teacher else dist.sample()
                logp = dist.log_prob(act)
            b_obs[t], b_mask[t], b_act[t], b_logp[t], b_val[t], b_teach[t] = o, m, act, logp, v, tch
            acts = act.cpu().numpy().astype(np.int32)
            t1 = time.time()
            obs, mask, teacher, rew, done, info = env.step(acts)
            t2 = time.time()
            t_inf += t1 - t0
            t_env += t2 - t1
            b_rew[t] = torch.from_numpy(rew).to(dev)
            b_done[t] = torch.from_numpy(done.astype(np.float32)).to(dev)
            for i in np.nonzero(done)[0]:
                ep_dps.append(float(info[i, 0]))
                totals["fights"] += 1
                if a.learn_talents and info[i, 2] > 0:
                    ep_builds[env.talents(int(i))] += 1
        totals["decisions"] += R * N
        with torch.no_grad():
            o = torch.from_numpy(obs).to(dev)
            m = torch.from_numpy(mask).to(dev).bool()
            _, last_v = net(o, m)
        # Talent decisions are undiscounted so they see the whole fight's damage.
        gam = torch.where(b_obs[:, :, 0] > 0.5, torch.ones_like(b_val), torch.full_like(b_val, a.gamma))
        adv = torch.zeros_like(b_rew)
        last = torch.zeros(N, device=dev)
        for t in reversed(range(R)):
            nv = last_v if t == R - 1 else b_val[t + 1]
            nonterm = 1.0 - b_done[t]
            delta = b_rew[t] + gam[t] * nv * nonterm - b_val[t]
            last = delta + gam[t] * a.lam * nonterm * last
            adv[t] = last
        return adv, adv + b_val, t_env, t_inf

    teacher_dps = None

    def play(use_teacher):
        eobs, emask, etch = eval_env.restart(eval_seeds)
        n = eval_env.n
        dps, builds = np.full(n, np.nan), collections.Counter()
        steps = 0
        while np.isnan(dps).any() and steps < 100000:
            if use_teacher:
                act = etch.astype(np.int32)
            else:
                with torch.no_grad():
                    logits, _ = net(torch.from_numpy(eobs).to(dev), torch.from_numpy(emask).to(dev).bool())
                    act = logits.argmax(-1).cpu().numpy().astype(np.int32)
            eobs, emask, etch, _, edone, einfo = eval_env.step(act)
            steps += 1
            for i in np.nonzero(edone)[0]:
                if np.isnan(dps[i]):
                    dps[i] = float(einfo[i, 0])
                    builds[eval_env.talents(int(i))] += 1
        return dps, builds

    def evaluate():
        """Greedy DPS on the evaluation fights, and its paired difference to the published APL there."""
        nonlocal teacher_dps
        if teacher_dps is None and not a.learn_talents:
            teacher_dps, _ = play(True)
        dps, builds = play(False)
        out = {"eval_dps": round(float(dps.mean()), 2)}
        if teacher_dps is not None:
            d = dps - teacher_dps
            out |= {"teacher_dps": round(float(teacher_dps.mean()), 2), "vs_teacher": round(float(d.mean()), 2),
                    "vs_teacher_se": round(float(d.std() / np.sqrt(len(d))), 2)}
        return out, builds.most_common(3)

    def update(adv, ret, bc, bc_weight=0.0):
        n = R * N
        f = lambda x: x.reshape(n, *x.shape[2:])
        o, m, act, old_logp, ret_f, adv_f, tch = f(b_obs), f(b_mask), f(b_act), f(b_logp), f(ret), f(adv), f(b_teach)
        talent_step = o[:, 0] > 0.5
        adv_f = (adv_f - adv_f.mean()) / (adv_f.std() + 1e-8)
        stats = collections.defaultdict(float)
        k = 0
        for _ in range(a.epochs):
            epoch_kl, epoch_k = 0.0, 0
            perm = torch.randperm(n, device=dev)
            for s in range(0, n, a.minibatch):
                idx = perm[s:s + a.minibatch]
                logits, v = net(o[idx], m[idx])
                logp_all = F.log_softmax(logits, -1)
                logp = logp_all.gather(1, act[idx][:, None]).squeeze(1)
                p = logp_all.exp()
                ent = -(p * logp_all.masked_fill(~m[idx], 0)).sum(-1)
                ts = talent_step[idx].float()
                ent_term = (a.ent * (1 - ts) + a.talent_ent * ts) * ent
                vloss = F.mse_loss(v, ret_f[idx])
                if bc:
                    pl = F.cross_entropy(logits, tch[idx])
                    loss = pl + a.vf * vloss
                    stats["bc_acc"] += (logits.argmax(-1) == tch[idx]).float().mean().item()
                else:
                    ratio = (logp - old_logp[idx]).exp()
                    s1, s2 = ratio * adv_f[idx], ratio.clamp(1 - a.clip, 1 + a.clip) * adv_f[idx]
                    pl = -torch.min(s1, s2).mean()
                    loss = pl + a.vf * vloss - ent_term.mean()
                    if bc_weight > 0:
                        bce = F.cross_entropy(logits, tch[idx])
                        loss = loss + bc_weight * bce
                        stats["bc_ce"] += bce.item()
                    kl = (old_logp[idx] - logp).mean().item()
                    stats["kl"] += kl
                    epoch_kl += kl
                    epoch_k += 1
                    stats["clipfrac"] += ((ratio - 1).abs() > a.clip).float().mean().item()
                opt.zero_grad()
                loss.backward()
                nn.utils.clip_grad_norm_(net.parameters(), 0.5)
                opt.step()
                stats["pl"] += pl.item()
                stats["vloss"] += vloss.item()
                fight_mask = ts < 0.5
                if fight_mask.any():
                    stats["ent_fight"] += ent[fight_mask].mean().item()
                if (~fight_mask).any():
                    stats["ent_talent"] += ent[~fight_mask].mean().item()
                k += 1
            if not bc and a.target_kl and epoch_k and epoch_kl / epoch_k > 1.5 * a.target_kl:
                break
        return {key: val / k for key, val in stats.items()}

    it = 0
    last_eval = None
    best_eval = -1.0
    while True:
        elapsed = time.time() - t_start
        if elapsed > a.minutes * 60 or (a.max_decisions and totals["decisions"] >= a.max_decisions):
            break
        bc = it < a.bc_iters
        t0 = time.time()
        adv, ret, t_env, t_inf = rollout(use_teacher=bc)
        t1 = time.time()
        ppo_it = it - a.bc_iters
        bc_weight = a.bc_coef if not a.bc_decay else a.bc_coef * max(0.0, 1 - ppo_it / a.bc_decay)
        st = update(adv, ret, bc, 0.0 if bc else bc_weight)
        t2 = time.time()
        it += 1
        row = {"iter": it, "phase": "bc" if bc else "ppo", "elapsed": round(time.time() - t_start, 2),
               "decisions": totals["decisions"], "fights": totals["fights"],
               "dps": round(float(np.mean(list(ep_dps)[-max(1, N):])), 2) if ep_dps else None,
               "t_env": round(t_env, 3), "t_inf": round(t_inf, 3), "t_update": round(t2 - t1, 3),
               "decisions_per_s": round(R * N / (t2 - t0)), **{k: round(v, 4) for k, v in st.items()}}
        if a.learn_talents and ep_builds:
            row["top_builds"] = ep_builds.most_common(3)
            ep_builds.clear()
        if a.eval_every and it % a.eval_every == 0:
            ev_row, row["eval_builds"] = evaluate()
            row |= ev_row
            last_eval = row["eval_dps"]
            if ev_row["eval_dps"] > best_eval:
                best_eval = ev_row["eval_dps"]
                torch.save({"net": net.state_dict(), "spec": spec, "args": vars(a) | {"out": str(a.out)}, "eval": ev_row},
                           a.out / "best.pt")
        log.write(json.dumps(row) + "\n")
        log.flush()
        print(json.dumps({k: row[k] for k in row if k not in ("top_builds", "eval_builds")}), flush=True)
    final_eval, final_builds = evaluate()
    torch.save({"net": net.state_dict(), "spec": spec, "args": vars(a) | {"out": str(a.out)}}, a.out / "model.pt")
    summary = {"final_eval": final_eval, "best_eval_dps": best_eval, "final_builds": final_builds, "iterations": it,
               "decisions": totals["decisions"],
               "fights": totals["fights"], "minutes": round((time.time() - t_start) / 60, 2)}
    (a.out / "summary.json").write_text(json.dumps(summary, indent=1))
    print(json.dumps(summary), flush=True)
    env.close()
    eval_env.close()


if __name__ == "__main__":
    main()
