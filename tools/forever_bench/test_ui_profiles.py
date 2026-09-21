import json
import hashlib
from pathlib import Path
import unittest

from export_ui_profiles import ui_settings


class UIProfilesTest(unittest.TestCase):
    def test_bundled_defaults_match_published_requests(self):
        root = Path(__file__).resolve().parents[2]
        source = (root / "artifacts/forever_dps_5min.json").read_bytes()
        results = json.loads(source)["Results"]
        bundle = json.loads((root / "ui/core/forever_ranked_profiles.json").read_text())
        self.assertEqual(bundle["sourceSHA256"], hashlib.sha256(source).hexdigest())
        self.assertEqual(len(bundle["profiles"]), 147)
        profiles = {(p["key"], p["race"]): p for p in bundle["profiles"]}
        self.assertEqual(len(profiles), 147)
        for row in results:
            profile = profiles[row["Key"], row["Race"]]
            self.assertEqual(profile["settings"], ui_settings(row))
            self.assertEqual(profile["dps"], row["DPS"])

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
