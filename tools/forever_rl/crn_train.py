#!/usr/bin/env python3
"""Approximate policy iteration with common random numbers: the learner the simulator makes possible.

Each iteration plays fights with the current network, picks one decision per fight, replays the fight exactly (same
seed, labelled random streams, same earlier actions) once per candidate action at that decision, lets the greedy
network finish each copy, and trains the network towards the action whose copy dealt the most damage. Talent picks
are decisions like any other when --learn-talents is on. The network starts from behaviour cloning of the published
rotation and talents.

    python3 tools/forever_rl/crn_train.py --build fire --race Human --minutes 5 --out ~/.cache/forever-sim-audit/rl/runs/crn
"""
import argparse
import collections
import json
import time
from pathlib import Path

import numpy as np
import torch
import torch.nn.functional as F

from forever_env import VecEnv
from train import Net


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--build", default="fire")
    ap.add_argument("--race", default="Human")
    ap.add_argument("--learn-talents", action="store_true")
    ap.add_argument("--threads", type=int, default=8)
    ap.add_argument("--points", type=int, default=512, help="decisions labelled per iteration (one per played fight)")
    ap.add_argument("--k", type=int, default=4, help="candidate actions tried at each labelled decision")
    ap.add_argument("--build-share", type=float, default=0.3, help="share of labelled decisions taken from talent picks")
    ap.add_argument("--minutes", type=float, default=5)
    ap.add_argument("--bc-iters", type=int, default=40)
    ap.add_argument("--bc-envs", type=int, default=1024)
    ap.add_argument("--hidden", type=int, default=256)
    ap.add_argument("--lr", type=float, default=1e-4)
    ap.add_argument("--epochs", type=int, default=2, help="passes over the label buffer per iteration")
    ap.add_argument("--buffer", type=int, default=65536, help="labelled decisions kept")
    ap.add_argument("--temp", type=float, default=3.0, help="softmax temperature over candidate DPS (DPS units)")
    ap.add_argument("--anchor", type=float, default=0.2, help="weight of the cross-entropy to the network's own earlier greedy action on buffer states (keeps unlabelled behaviour)")
    ap.add_argument("--explore", type=float, default=1.0, help="sampling temperature for the played fights")
    ap.add_argument("--eval-fights", type=int, default=256)
    ap.add_argument("--eval-every", type=int, default=5)
    ap.add_argument("--seed", type=int, default=1)
    ap.add_argument("--init", type=Path, help="start from a model.pt instead of behaviour cloning")
    ap.add_argument("--out", type=Path, required=True)
    a = ap.parse_args()
    a.out.mkdir(parents=True, exist_ok=True)
    torch.manual_seed(a.seed)
    rng = np.random.default_rng(a.seed)
    dev = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    base_cfg = dict(build=a.build, race=a.race, learnTalents=a.learn_talents, threads=a.threads, labeledRands=True,
                    rewardScale=1e-5, seed=1_000_000 * a.seed + 3)
    t_start = time.time()
    log = (a.out / "log.jsonl").open("a")

    play = VecEnv(base_cfg, a.points)
    branch = VecEnv(base_cfg, a.points * a.k)
    evalenv = VecEnv({**base_cfg, "seed": 987_654_321}, a.eval_fights)
    spec = play.spec
    D, A, T = spec["obsDim"], spec["actDim"], spec["talents"]
    (a.out / "spec.json").write_text(json.dumps(spec))
    net = Net(D, A, a.hidden).to(dev)
    opt = torch.optim.Adam(net.parameters(), lr=a.lr, eps=1e-5)

    def policy(obs, mask, greedy, temp=1.0):
        with torch.no_grad():
            logits, _ = net(torch.from_numpy(obs).to(dev), torch.from_numpy(mask).to(dev).bool())
            if greedy:
                return logits.argmax(-1).cpu().numpy().astype(np.int32), logits
            return torch.distributions.Categorical(logits=logits / temp).sample().cpu().numpy().astype(np.int32), logits

    # Behaviour cloning of the published rotation (and talents) on a separate vector.
    if a.init:
        net.load_state_dict(torch.load(a.init, map_location=dev)["net"])
    else:
        bc = VecEnv({**base_cfg, "seed": 55_555}, a.bc_envs)
        obs, mask, teacher = bc.reset()
        for it in range(a.bc_iters):
            O, M, Y = [], [], []
            for _ in range(32):
                O.append(obs.copy()); M.append(mask.copy()); Y.append(teacher.copy())
                obs, mask, teacher, *_ = bc.step(teacher)
            o = torch.from_numpy(np.concatenate(O)).to(dev)
            m = torch.from_numpy(np.concatenate(M)).to(dev).bool()
            y = torch.from_numpy(np.concatenate(Y).astype(np.int64)).to(dev)
            net.norm.update(o)
            for _ in range(2):
                perm = torch.randperm(len(o), device=dev)
                for s in range(0, len(o), 8192):
                    idx = perm[s:s + 8192]
                    logits, _ = net(o[idx], m[idx])
                    loss = F.cross_entropy(logits, y[idx])
                    opt.zero_grad(); loss.backward(); opt.step()
        acc = (net(o, m)[0].argmax(-1) == y).float().mean().item()
        bc.close()
        print(json.dumps({"phase": "bc", "elapsed": round(time.time() - t_start, 1), "bc_acc": round(acc, 4)}), flush=True)

    def evaluate():
        eobs, emask, _ = evalenv.restart(np.arange(1, evalenv.n + 1, dtype=np.int64) * 7_777)
        dps, builds = np.full(evalenv.n, np.nan), collections.Counter()
        for _ in range(20000):
            act, _ = policy(eobs, emask, True)
            eobs, emask, eteach, _, edone, einfo = evalenv.step(act)
            for i in np.nonzero(edone)[0]:
                if np.isnan(dps[i]):
                    dps[i] = einfo[i, 0]
                    builds[evalenv.talents(int(i))] += 1
            if not np.isnan(dps).any():
                break
        return float(np.nanmean(dps)), builds.most_common(2)

    # Label buffer: states, masks, candidate actions, target distributions over candidates.
    buf = collections.deque(maxlen=a.buffer)
    it = 0
    seed_counter = 10_000_000 * a.seed
    ev0 = evaluate()
    print(json.dumps({"phase": "start", "eval_dps": round(ev0[0], 2), "eval_builds": ev0[1]}), flush=True)
    log.write(json.dumps({"iter": 0, "elapsed": round(time.time() - t_start, 2), "eval_dps": ev0[0], "eval_builds": ev0[1]}) + "\n")
    labelled = fights = 0
    while time.time() - t_start < a.minutes * 60:
        t0 = time.time()
        it += 1
        P, K = a.points, a.k
        seeds = seed_counter + np.arange(1, P + 1, dtype=np.int64)
        seed_counter += P + 1
        # 1. Play one fight per environment with the sampled policy, recording every action.
        obs, mask, teacher = play.restart(seeds)
        cols, sums = [], []
        build_steps = np.zeros(P, np.int64)
        done_at = np.full(P, -1)
        step = 0
        while (done_at < 0).any() and step < 20000:
            act, _ = policy(obs, mask, False, a.explore)
            cols.append(act.copy())
            sums.append(obs.astype(np.float64).sum(1))
            build_steps += (obs[:, 0] > 0.5) & (done_at < 0)
            obs, mask, teacher, _, done, info = play.step(act)
            step += 1
            newly = (done == 1) & (done_at < 0)
            done_at[newly] = step
        amat = np.stack(cols, 1)  # P x steps; row i is valid up to done_at[i]
        smat = np.stack(sums, 1)  # observation checksums, to confirm the replays reach the same states
        lengths = done_at
        fights += P
        # 2. One decision per fight: a talent pick with probability build-share (when learning talents), else a fight decision.
        t_b = np.zeros(P, np.int64)
        for i in range(P):
            n, nb = int(lengths[i]), int(build_steps[i])
            if nb > 0 and rng.random() < a.build_share:
                t_b[i] = rng.integers(0, nb)
            else:
                t_b[i] = rng.integers(nb, n)
        # 3. Replay K copies of each fight to its decision, try one candidate each, finish greedily.
        bseeds = np.repeat(seeds, K)
        bobs, bmask, bteach = branch.restart(bseeds)
        cand = np.full((P, K), -1, np.int64)
        state_obs = np.zeros((P, D), np.float32)
        state_mask = np.zeros((P, A), np.uint8)
        result = np.full(P * K, np.nan)
        bstep = 0
        mismatched = 0
        tb_rep = np.repeat(t_b, K)
        while np.isnan(result).any() and bstep < 20000:
            greedy, logits = policy(bobs, bmask, True)
            act = greedy.copy()
            if bstep < amat.shape[1]:
                prefix = tb_rep > bstep
                act[prefix] = np.repeat(amat[:, bstep], K)[prefix]
            lg = None
            for i in np.nonzero(t_b == bstep)[0]:
                rows = slice(i * K, (i + 1) * K)
                if lg is None:
                    lg = logits.cpu().numpy()
                if True:
                    o = bobs[i * K]
                    if not np.allclose(bobs[rows], o) or abs(float(o.astype(np.float64).sum()) - smat[i, bstep]) > 1e-3:
                        mismatched += 1
                    state_obs[i], state_mask[i] = o, bmask[i * K]
                    legal = np.nonzero(bmask[i * K])[0]
                    order = legal[np.argsort(-lg[i * K, legal])]
                    picks = [int(order[0]), int(bteach[i * K])]
                    picks += [int(x) for x in order[1:]]
                    seen = []
                    for x in picks:
                        if x not in seen:
                            seen.append(x)
                    cand[i, :min(K, len(seen))] = seen[:K]
                    for k in range(K):
                        act[i * K + k] = cand[i, k] if cand[i, k] >= 0 else cand[i, 0]
            reseed = np.zeros(P * K, np.int64)
            hit = np.repeat(t_b == bstep, K)
            reseed[hit] = np.repeat(seeds * 31 + 17, K)[hit]
            bobs, bmask, bteach, _, bdone, binfo = branch.step_reseed(act, reseed)
            bstep += 1
            newly = (bdone == 1) & np.isnan(result)
            result[newly] = binfo[newly, 0]
        Q = result.reshape(P, K)
        for i in range(P):
            valid = cand[i] >= 0
            if valid.sum() < 2:
                continue
            q = np.where(valid, Q[i], -np.inf)
            w = np.exp((q - q[valid].max()) / a.temp)
            w = w / w.sum()
            buf.append((state_obs[i].copy(), state_mask[i].copy(), cand[i].copy(), w.astype(np.float32),
                        float(np.nanmax(q[valid]) - np.nanmin(q[valid]))))
            labelled += 1
        t1 = time.time()
        # 4. Train on the buffer: cross-entropy to the candidate targets, plus an anchor to the current greedy action
        #    so behaviour at unlabelled states does not drift.
        items = list(buf)
        o = torch.from_numpy(np.stack([x[0] for x in items])).to(dev)
        m = torch.from_numpy(np.stack([x[1] for x in items])).to(dev).bool()
        c = torch.from_numpy(np.stack([x[2] for x in items])).to(dev)
        w = torch.from_numpy(np.stack([x[3] for x in items])).to(dev)
        net.norm.update(o[-P:])
        with torch.no_grad():
            anchor = net(o, m)[0].argmax(-1)
        losses = []
        for _ in range(a.epochs):
            perm = torch.randperm(len(o), device=dev)
            for s in range(0, len(o), 4096):
                idx = perm[s:s + 4096]
                logits, _ = net(o[idx], m[idx])
                logp = F.log_softmax(logits, -1)
                ci = c[idx].clamp(min=0)
                lp = logp.gather(1, ci) * (c[idx] >= 0)
                loss = -(w[idx] * lp).sum(-1).mean() + a.anchor * F.cross_entropy(logits, anchor[idx])
                opt.zero_grad(); loss.backward(); opt.step()
                losses.append(loss.item())
        t2 = time.time()
        gaps = [x[4] for x in items[-P:]]
        row = {"iter": it, "elapsed": round(time.time() - t_start, 2), "fights_played": fights, "labelled": labelled,
               "branch_fights": it * P * K, "t_play_branch": round(t1 - t0, 2), "t_train": round(t2 - t1, 2),
               "loss": round(float(np.mean(losses)), 4), "median_gap_dps": round(float(np.median(gaps)), 2),
               "replay_mismatches": mismatched, "best_is_greedy": round(float(np.mean([np.argmax(np.where(cand[i] >= 0, Q[i], -np.inf)) == 0 for i in range(P)])), 3)}
        if a.eval_every and it % a.eval_every == 0:
            row["eval_dps"], row["eval_builds"] = evaluate()
        log.write(json.dumps(row) + "\n"); log.flush()
        print(json.dumps({k: v for k, v in row.items() if k != "eval_builds"}), flush=True)
    final = evaluate()
    torch.save({"net": net.state_dict(), "spec": spec, "args": vars(a) | {"out": str(a.out), "init": str(a.init)}}, a.out / "model.pt")
    summary = {"final_eval_dps": final[0], "final_builds": final[1], "iterations": it, "labelled": labelled,
               "fights_played": fights, "branch_fights": it * a.points * a.k, "minutes": round((time.time() - t_start) / 60, 2)}
    (a.out / "summary.json").write_text(json.dumps(summary, indent=1))
    print(json.dumps(summary), flush=True)
    for e in (play, branch, evalenv):
        e.close()


if __name__ == "__main__":
    main()
