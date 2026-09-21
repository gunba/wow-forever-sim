import copy
import unittest

from sensitivity import gain, summarize, compare


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
