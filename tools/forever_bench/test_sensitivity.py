import copy
import unittest

from sensitivity import gain, summarize, compare, scaling_amplification, summarize_comparisons, format_metric


class SensitivityTest(unittest.TestCase):
    def test_gain_denominator_and_bound(self):
        pct, bound = gain({"DPS": 120, "StandardError": 1}, {"DPS": 100, "StandardError": 1})
        self.assertAlmostEqual(pct, 20)
        self.assertAlmostEqual(bound, 1.96 * 2.2)

    def test_equal_race_weights(self):
        rows = [{"Faction": "Alliance", "GainPercent": 10, "MonteCarlo95Bound": 1},
                {"Faction": "Horde", "GainPercent": 30, "MonteCarlo95Bound": 3}]
        self.assertEqual(summarize(rows), {"Races": 2, "GainPercent": 20, "MonteCarlo95Bound": 2})
        self.assertEqual(summarize(rows, "horde")["GainPercent"], 30)

    def test_scaling_amplification(self):
        near = {"Races": 2, "GainPercent": 10, "MonteCarlo95Bound": 1}
        far = {"Races": 2, "GainPercent": 50, "MonteCarlo95Bound": 2}
        linear = scaling_amplification(near, far)
        self.assertEqual(linear["Amplification"], 1)
        self.assertAlmostEqual(linear["MonteCarlo95Bound"], .14)
        self.assertEqual(format_metric("amplification", linear), "1.00×")
        self.assertEqual(format_metric("gear10", near), "+10.0%")
        far["GainPercent"] = 100
        self.assertEqual(scaling_amplification(near, far)["Amplification"], 2)
        far["GainPercent"] = 25
        self.assertEqual(scaling_amplification(near, far)["Amplification"], .5)
        # DPS proportional to the square of scaled gear: accelerating, not exponential.
        near["GainPercent"], far["GainPercent"] = 21, 125
        self.assertAlmostEqual(scaling_amplification(near, far)["Amplification"], 125 / 105)

    def test_unstable_amplification_is_unavailable(self):
        far = {"Races": 1, "GainPercent": 50, "MonteCarlo95Bound": 2}
        for value, bound in [(-1, .1), (0, 0), (.05, 0), (.1, 0), (1, 1), (1, 2)]:
            with self.subTest(value=value, bound=bound):
                result = scaling_amplification(
                    {"Races": 1, "GainPercent": value, "MonteCarlo95Bound": bound}, far)
                self.assertIsNone(result["Amplification"])
                self.assertIsNone(result["MonteCarlo95Bound"])
                self.assertIn("UnavailableReason", result)
                self.assertEqual(format_metric("amplification", result), "—")

    def test_amplification_uses_ratio_of_equal_race_means(self):
        def rows(values):
            return [{"Faction": faction, "GainPercent": value, "MonteCarlo95Bound": .1}
                    for faction, value in zip(("Alliance", "Horde"), values)]
        pairs = {"tier1": rows((10, 20)), "gear10": rows((10, 20)), "gear50": rows((50, 150))}
        result = summarize_comparisons(pairs)
        self.assertEqual(result["amplification"]["Gear10GainPercent"], 15)
        self.assertEqual(result["amplification"]["Gear50GainPercent"], 100)
        self.assertAlmostEqual(result["amplification"]["Amplification"], 100 / 75)
        self.assertEqual(summarize_comparisons(pairs, "alliance")["amplification"]["Amplification"], 1)
        self.assertEqual(summarize_comparisons(pairs, "horde")["amplification"]["Amplification"], 1.5)

    def test_reject_changed_inputs(self):
        base_row = {
            "Key": "test", "Race": "Human", "Faction": "Alliance", "DPS": 120,
            "StandardError": 1, "Iterations": 5000, "OOMSeconds": 0, "Warnings": [],
            "BaselinePlayer": {"foreverTier1Bonuses": True},
            "Request": {"raid": {"parties": [{"players": [{}]}]}, "simOptions": {"randomSeed": 1}},
            "Hit": {"Balance": 0},
        }
        base = {"Duration": 300, "StartingArmor": 3731, "Mechanics": {}, "Results": [base_row]}
        off = copy.deepcopy(base)
        off.update(Tier1Bonuses=False, EquipmentScale=1)
        off["Results"][0]["BaselinePlayer"] = {}
        off["Results"][0]["DPS"] = 100
        self.assertAlmostEqual(compare(base, off, 1, False)[0]["GainPercent"], 20)
        off["Results"][0]["Request"]["simOptions"]["randomSeed"] = 2
        with self.assertRaisesRegex(ValueError, "seed"):
            compare(base, off, 1, False)


if __name__ == "__main__":
    unittest.main()
