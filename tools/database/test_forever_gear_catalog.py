import unittest

from forever_gear_catalog import TABLES, acquisition, build_catalog, item_stats, planner_record, validate_planner


def sparse(slot, quality, ilvl, mods):
    row = {"InventoryType": str(slot), "OverallQualityID": str(quality), "ItemLevel": str(ilvl)}
    for n, (mod, allocation) in enumerate(mods):
        row[f"StatModifier_bonusStat_{n}"] = str(mod)
        row[f"StatPercentEditor_{n}"] = str(allocation)
    return row


class GearCatalogTest(unittest.TestCase):
    def test_shield_block_uses_vendor_tooltip_before_web_capture(self):
        from forever_gear_catalog import add_shield_block_values
        items = [{"id": 1, "inventoryType": 14}, {"id": 2, "inventoryType": 14}]
        evidence = [{"id": id, "baseBlockValue": 40, "source": "captured tooltip"} for id in (1, 2)]
        vendor = {"item": {"equipLocation": "INVTYPE_SHIELD"},
                  "tooltip": {"lines": [{"left": "44 Block"}]}}
        add_shield_block_values(items, evidence, {1: vendor})
        self.assertEqual(items[0]["baseBlockValue"], 44)
        self.assertEqual(items[1]["baseBlockValue"], 40)

    def test_unused_effect_reference_does_not_block_catalog(self):
        tables = {name: [] for name in TABLES}
        tables["ItemXItemEffect"] = [{"ItemID": "240068", "ItemEffectID": "112960"}]
        self.assertEqual(build_catalog({}, tables)["items"], [])

    def test_crafted_source(self):
        sources, reason = acquisition({"sourcemore": [{"t": 6, "s": 165, "ti": 1306527}]})
        self.assertIsNone(reason)
        self.assertEqual(sources, [{"kind": "crafted", "skill": 165, "spellId": 1306527}])

    def test_dungeon_source(self):
        sources, reason = acquisition({"sourcemore": [{"t": 1, "ti": 1853, "z": 2057}]})
        self.assertIsNone(reason)
        self.assertEqual(sources[0]["kind"], "dungeon")

    def test_exported_honor_purchase_is_available(self):
        export = {"costs": [{"currencyName": "Honor Points", "amount": 8000}]}
        sources, reason = acquisition({}, export)
        self.assertIsNone(reason)
        self.assertEqual(sources, [{"kind": "pvp", "exported": True}])
        # A vendor listing alone must not admit an unavailable token reward.
        self.assertFalse(acquisition({}, {"costs": [{"itemID": 22726, "amount": 1}]})[0])
        self.assertEqual(
            acquisition({"sourcemore": [{"t": 5, "c": 3456}]}, export),
            ([], "raid-derived"),
        )

    def test_nonraid_vendor_and_faction(self):
        item = {"source": [5], "sourcemore": [{"t": 1, "ti": 15127, "z": 45, "n": "Samuel Hawke"}]}
        sources, reason = acquisition(item)
        self.assertIsNone(reason)
        self.assertEqual(sources[0]["faction"], "Alliance")
        self.assertEqual(sources[0]["kind"], "pvp")
        item["sourcemore"] = [{"t": 1, "ti": 227853, "z": 33, "n": "Pix Xizzix"}]
        self.assertEqual(acquisition(item)[0][0]["kind"], "vendor")
        item["sourcemore"] = [{"t": 1, "ti": 15192, "z": 440, "n": "Anachronos"}]
        self.assertFalse(acquisition(item)[0])

    def test_dungeon_quest_and_zone_only_source(self):
        self.assertEqual(acquisition({"sourcemore": [
            {"t": 5, "c": 2017, "ti": 5243, "n": "Houses of the Holy"},
        ]})[0], [{"kind": "dungeon-quest", "zoneId": 2017, "questId": 5243, "name": "Houses of the Holy"}])
        self.assertEqual(acquisition({"sourcemore": [{"z": 1583}]})[0][0]["kind"], "dungeon")
        # Outdoor quest chains can require raids; do not infer availability.
        self.assertFalse(acquisition({"sourcemore": [{"t": 5, "c": 16, "ti": 7486}]})[0])

    def test_published_dungeon_stats_do_not_require_static_client_record(self):
        item = {
            "id": 16707, "name": "Shadowcraft Cap", "class": 4, "subclass": 2,
            "itemLevel": 62, "quality": 3, "inventoryType": 1, "requiredLevel": 57,
            "stats": {"str": 6, "agi": 20, "sta": 20, "spi": 13, "armor": 141, "itemset": 184},
        }
        record = planner_record(item, [{"kind": "dungeon", "zoneId": 2057}])
        self.assertFalse(record["clientStatRecord"])
        self.assertEqual(record["stats"], {"Strength": 6, "Agility": 20, "Stamina": 20, "Spirit": 13})

    def test_raid_vendor_and_token_sources(self):
        for details in ([{"z": 139}], [{"t": 1, "z": 3456}],
                        [{"t": 5, "c": 3456}, {"t": 6, "s": 197, "ti": 28207}]):
            self.assertFalse(acquisition({"sourcemore": details})[0])

    def test_outlaw_collar(self):
        row = sparse(1, 4, 65, [(7, 6666), (3, 4444), (32, 4444)])
        stats, allocations, scale = item_stats(row, {"65": {"EpicF_0": "47"}})
        self.assertEqual(stats, {"Stamina": 31, "Agility": 21, "CritRating": 21})
        self.assertEqual(scale, 47)
        self.assertFalse(validate_planner({"sta": 31, "agi": 21, "critstrkrtng": 21}, allocations))

    def test_hidden_school_power(self):
        # Robes of Fiery Devastation: Wowhead omits item mod 85 entirely.
        row = sparse(20, 4, 65, [(7, 5777), (5, 5066), (85, 5545)])
        stats, allocations, _ = item_stats(row, {"65": {"EpicF_0": "47"}})
        self.assertEqual(stats, {"Stamina": 27, "Intellect": 24, "FirePower": 26})
        self.assertFalse(validate_planner({"sta": 27, "int": 24}, allocations))
        self.assertTrue(validate_planner({"sta": 28}, allocations))

    def test_shield_uses_offhand_points(self):
        row = sparse(14, 4, 65, [(7, 5777), (32, 5066)])
        stats, _, scale = item_stats(row, {"65": {"EpicF_2": "27", "EpicF_3": "20"}})
        self.assertEqual(scale, 27)
        self.assertEqual(stats, {"Stamina": 16, "CritRating": 14})

    def test_unknown_stat_is_retained(self):
        row = sparse(1, 4, 65, [(199, 10000)])
        stats, _, _ = item_stats(row, {"65": {"EpicF_0": "47"}})
        self.assertEqual(stats, {"UnmappedItemMod199": 47})


if __name__ == "__main__":
    unittest.main()
