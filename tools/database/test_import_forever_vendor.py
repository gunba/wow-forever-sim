import unittest
from pathlib import Path

import import_forever_vendor as vendor


ROOT = Path(__file__).resolve().parents[2]


class VendorImportTest(unittest.TestCase):
    def test_intrinsic_shield_block_is_not_in_get_item_stats(self):
        for item_id in (272591, 278469):
            source = self.exports[item_id]
            self.assertEqual(vendor.shield_block_value(source), 44)
            self.assertEqual(vendor.build_stats(self.planner[str(item_id)], source)[30], 44)

    @classmethod
    def setUpClass(cls):
        cls.planner = vendor.parse_planner(ROOT / "assets/db_inputs/wowhead_forever_gearplanner.txt")
        cls.exports, cls.set_names = vendor.read_vendor_exports(ROOT / "assets/db_inputs/forever_vendor")

    def item(self, item_id):
        return vendor.item_to_proto(self.planner[str(item_id)], self.exports.get(item_id),
                                    self.set_names.get(item_id), {})

    def test_exported_weapon_stats(self):
        for item_id in (272592, 272600):
            item = self.item(item_id)
            self.assertEqual(item["stats"][14] + item["stats"][19], 1)
            self.assertEqual(item["stats"][17], 24)
            self.assertEqual(item["stats"][27], 24)
            self.assertNotIn("weaponSkills", item)
            self.assertEqual((item["weaponDamageMin"], item["weaponDamageMax"],
                              item["weaponSpeed"]), (109, 165, 2.9))

    def test_caster_weapons(self):
        for item_id, damage in ((272683, (44, 82)), (272603, (109, 165))):
            item = self.item(item_id)
            self.assertEqual(item["stats"][5], 94)
            self.assertEqual((item["weaponDamageMin"], item["weaponDamageMax"]), damage)

    def test_ranged_weapon(self):
        item = self.item(272595)
        self.assertEqual((item["weaponDamageMin"], item["weaponDamageMax"]), (85, 129))
        self.assertEqual(item["stats"][27], 32)

    def test_exported_armor_matches_planner_conversions(self):
        # The exported class armor is our evidence for retaining the other sets.
        for item_id, source in self.exports.items():
            if source["item"].get("setID"):
                with self.subTest(item=item_id):
                    planner = self.planner[str(item_id)]
                    self.assertEqual(vendor.build_stats(planner, source),
                                     vendor.build_stats(planner))
                    self.assertEqual(vendor.class_allowlist(planner, source),
                                     vendor.class_allowlist(planner, None))

    def test_provisional_armor_and_selection(self):
        ids = vendor.selected_ids(self.planner, self.exports)
        self.assertEqual(len(ids), 116)
        self.assertEqual(len(ids - self.exports.keys()), 64)
        for item_id in ids:
            with self.subTest(item=item_id):
                self.assertEqual(len(self.item(item_id)["stats"]), vendor.STAT_COUNT)
        item = self.item(272647)
        self.assertEqual(item["stats"][14] + item["stats"][19], .5)
        self.assertEqual(item["classAllowlist"], [7])


if __name__ == "__main__":
    unittest.main()
