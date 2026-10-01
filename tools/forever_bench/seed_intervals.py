"""Iteration seeds in sim/core/sim.go are start + iteration index."""


def require_independent_seed_ranges(ranges):
    intervals = []
    for start, iterations in ranges:
        start, iterations = int(start), int(iterations)
        stop = start + iterations
        if start == 0 or iterations <= 0 or start < -(2**63) or stop > 2**63:
            raise ValueError("Invalid deterministic iteration-seed range")
        intervals.append((start, stop))
    intervals.sort()
    for previous, current in zip(intervals, intervals[1:]):
        if current[0] < previous[1]:
            raise ValueError(f"Independent iteration-seed ranges overlap: {previous} and {current}")
