import json
import hashlib
from pathlib import Path
import unittest

from export_ui_profiles import ui_settings
from build_display import BUILD_CAVEATS, expected_roster


class UIProfilesTest(unittest.TestCase):
    def test_bundled_defaults_match_published_requests(self):
        root = Path(__file__).resolve().parents[2]
        source = (root / "artifacts/modelled_gear/forever_dps_5min.json").read_bytes()
        results = json.loads(source)["Results"]
        bundle = json.loads((root / "ui/core/forever_ranked_profiles.json").read_text())
        self.assertEqual(bundle["gearScenario"], "modeled-65-v2")
        self.assertEqual(bundle["sourceSHA256"], hashlib.sha256(source).hexdigest())
        self.assertEqual(len(bundle["profiles"]), len(expected_roster()))
        profiles = {(p["key"], p["race"]): p for p in bundle["profiles"]}
        self.assertEqual(set(profiles), expected_roster())
        modeled_ids = {item["id"] for item in json.loads((
            root / "assets/db_inputs/forever_synthetic_gear_v2.json").read_text())["items"]}
        for row in results:
            profile = profiles[row["Key"], row["Race"]]
            self.assertEqual(profile["settings"], ui_settings(row))
            self.assertEqual(profile["dps"], row["DPS"])
            self.assertEqual(profile["caveats"], BUILD_CAVEATS.get(row["Key"], []))
            ids = [item.get("id", 0) for item in profile["settings"]["player"]["equipment"]["items"]]
            self.assertTrue(all(item_id in modeled_ids | {0} for item_id in ids))
            self.assertGreater(sum(bool(item_id) for item_id in ids), 14)
            self.assertGreaterEqual(row["Hit"]["RawHitDelta"], 0)
            self.assertEqual(row["Hit"]["Model"], "equipped-only-v2")
            bonus = profile["settings"]["player"].get("bonusStats", {}).get("stats", [])
            self.assertFalse(any(bonus[index] if index < len(bonus) else 0
                                 for index in (13, 18)),
                             "ranked web profile must not contain artificial melee/spell hit")

    def test_replay_preserves_normalized_player_and_scenario(self):
        data = json.loads((Path(__file__).resolve().parents[2] /
                           "artifacts/forever_mechanics_baseline_5min.json").read_text())
        for row in data["Results"]:
            settings = ui_settings(row)
            request = row["Request"]
            self.assertEqual(settings["player"], request["raid"]["parties"][0]["players"][0])
            self.assertEqual(settings["encounter"], request["encounter"])
            self.assertEqual(settings["raidBuffs"], request["raid"]["buffs"])
            self.assertEqual(settings["partyBuffs"], request["raid"]["parties"][0]["buffs"])
            self.assertEqual(settings["debuffs"], request["raid"]["debuffs"])
            self.assertTrue(settings["player"]["foreverTier1Bonuses"])
            self.assertEqual(settings["settings"]["fixedRngSeed"], request["simOptions"]["randomSeed"])
            self.assertIsNot(settings["player"], request["raid"]["parties"][0]["players"][0])


if __name__ == "__main__":
    unittest.main()
