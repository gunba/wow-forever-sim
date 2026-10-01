import unittest

from seed_intervals import require_independent_seed_ranges
from research_summary import summarize


class SeedIntervalTest(unittest.TestCase):
    def test_overlapping_starts_are_not_independent(self):
        for ranges in ([(20296070, 5000), (20296071, 5000)],
                       [(1, 100000), (25001, 5000)],
                       [(1, 5000), (5000, 5000)]):
            with self.subTest(ranges=ranges), self.assertRaises(ValueError):
                require_independent_seed_ranges(ranges)

    def test_adjacent_and_unsorted_ranges_are_disjoint(self):
        require_independent_seed_ranges([(5001, 5000), (1, 5000)])

    def test_historical_summary_does_not_assume_independence(self):
        data = {"EngineBaseRevision": "fixture", "Correction": "fixture",
                "Seeds": [1, 2], "IterationsPerArmPerSeed": 5000,
                "Choices": [{"Key": "fixture", "Candidate": "candidate"}],
                "Comparisons": [], "Runs": {}}
        for arm, dps in (("baseline", 100), ("candidate", 110)):
            for seed in data["Seeds"]:
                path = f"{arm}/{seed}"
                data["Comparisons"].append({"build": "fixture", "race": "fixture",
                                            "candidate": arm, "seed": seed, "path": path})
                data["Runs"][path] = {"Warnings": [], "Hit": {"Balance": 0},
                                      "DPS": dps, "StandardError": 2, "OOMSeconds": 0}
        result = summarize(data)
        self.assertFalse(result["SeedIntervalsIndependent"])
        self.assertEqual(result["ErrorAggregation"], "conservative-marginal-maximum")
        self.assertAlmostEqual(result["Choices"][0]["Races"][0]["Conservative95BoundDPS"], 7.84)

    def test_invalid_iteration_ranges(self):
        for ranges in ([(0, 5000)], [(1, 0)], [(1, -1)],
                       [(2**63 - 1, 2)]):
            with self.subTest(ranges=ranges), self.assertRaises(ValueError):
                require_independent_seed_ranges(ranges)
