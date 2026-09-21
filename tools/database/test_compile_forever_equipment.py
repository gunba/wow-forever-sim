import copy
import json
import unittest
from pathlib import Path

from compile_forever_equipment import compile_item, compile_stats
from import_forever_ratings import load_level_60_coefficients

ROOT = Path(__file__).resolve().parents[2]


class EquipmentTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.items = {x["id"]: x for x in json.loads(
            (ROOT / "assets/db_inputs/forever_gear_catalog.json").read_text())["items"]}
        cls.rates = load_level_60_coefficients()

    def test_outlaw_crit_is_one_shared_pool(self):
        item = compile_item(self.items[279253], self.rates)
        self.assertEqual(item["stats"][19], 1.5)
        self.assertEqual(item["stats"][14], 0)
        self.assertEqual(item["stats"][1:3], [21, 31])
        self.assertEqual(item["armorType"], 2)
        self.assertEqual(item["sources"][0]["crafted"]["profession"], 8)

    def test_weapon_and_school_power(self):
        axe = compile_item(self.items[279260], self.rates)
        self.assertEqual((axe["weaponDamageMin"], axe["weaponDamageMax"], axe["weaponSpeed"]), (162, 244, 3.3))
        self.assertEqual((axe["handType"], axe["weaponType"]), (4, 1))
        cord = compile_item(self.items[279267], self.rates)
        self.assertEqual(cord["stats"][6], 20)
        self.assertEqual(cord["stats"][5], 0)

    def test_generic_hit_ap_and_uniqueness(self):
        ring = compile_item(self.items[17713], self.rates)
        self.assertEqual(ring["stats"][18], 1)
        self.assertEqual(ring["stats"][13], 0)
        self.assertEqual(ring["stats"][17], 20)
        self.assertEqual(ring["stats"][27], 20)
        self.assertTrue(ring["unique"])

    def test_no_inferred_damage_from_healing(self):
        fixture = {"stats": {"HealingPower": 42}, "armor": 0}
        stats, _, _ = compile_stats(fixture, self.rates)
        self.assertEqual(stats[41], 42)
        self.assertEqual(stats[42], 0)

    def test_armor_and_unknown_stats(self):
        stats, _, _ = compile_stats({"armor": 44, "stats": {"BonusArmor": 170}}, self.rates)
        self.assertEqual(stats[26], 44)
        self.assertEqual(stats[40], 170)
        with self.assertRaisesRegex(ValueError, "unsupported stat"):
            compile_stats({"armor": 0, "stats": {"UnmappedItemMod124": 15}}, self.rates)

    def test_unique_equipped_groups_are_retained(self):
        for id in (276538, 276539, 276540, 276541):
            item = self.items[id]
            self.assertEqual(item["limitCategory"], 708)
            self.assertEqual(item["limitCategoryQuantity"], 1)
        self.assertEqual(self.items[279261]["limitCategoryQuantity"], 1)
        self.assertTrue(compile_item(self.items[279261], self.rates)["unique"])
        for id in (249469, 249470, 249473):
            self.assertTrue(self.items[id]["uniqueEquipped"])
            self.assertTrue(compile_item(self.items[id], self.rates)["unique"])

    def test_catalog_is_not_mutated(self):
        item = copy.deepcopy(self.items[279253])
        original = copy.deepcopy(item)
        compile_item(item, self.rates)
        self.assertEqual(item, original)


if __name__ == "__main__":
    unittest.main()
