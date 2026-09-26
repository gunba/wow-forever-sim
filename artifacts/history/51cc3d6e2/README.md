# Previous simulator release

These files preserve the 184-profile DPS benchmark and paired tank
comparisons from
[`51cc3d6e2`](https://github.com/gunba/wow-forever-sim/commit/51cc3d6e2).
The baseline and sensitivity archives contain all 736 complete DPS
requests/results; `tanks/` retains the 24 complete tank comparisons.

They predate the [weekly correctness review](../../../docs/weekly_review.md),
including Shield Block refresh, intrinsic shield Block, utility spell GCDs,
Blood Pact and Unbridled Wrath corrections. Replay them with the pinned
engine when reproducing historical numbers. The updated benchmark retains
the same requests; its changes measure engine corrections, not new gear or
rotation selections.
