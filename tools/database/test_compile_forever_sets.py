import unittest

from compile_forever_sets import compile_effects


class EquipmentSetEffectsTest(unittest.TestCase):
    def test_passive_power_regeneration_has_no_periodic_tick_field(self):
        for value in (3, 8, 12):
            effect = {"Effect": "6", "EffectAura": "85", "EffectMiscValue_0": "0",
                      "EffectBasePointsF": str(value), "EffectAuraPeriod": "0"}
            self.assertEqual(compile_effects([effect]), ({"MP5": value}, 0))

    def test_periodic_energize_is_not_passive_mp5(self):
        effect = {"Effect": "6", "EffectAura": "24", "EffectMiscValue_0": "0",
                  "EffectBasePointsF": "8", "EffectAuraPeriod": "1000"}
        self.assertIsNone(compile_effects([effect]))


if __name__ == "__main__":
    unittest.main()
