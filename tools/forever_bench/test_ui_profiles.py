import json
from pathlib import Path
import unittest

from export_ui_profiles import ui_settings


class UIProfilesTest(unittest.TestCase):
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
