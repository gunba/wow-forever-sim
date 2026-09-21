import json
from pathlib import Path
import unittest

from rotation_candidates import spell_id, variants


PROFILES = json.loads((Path(__file__).resolve().parents[2] / "artifacts/forever_optimization_baselines.json").read_text())["profiles"]


class RotationCandidatesTest(unittest.TestCase):
    def test_candidates_preserve_fixed_inputs(self):
        for build in PROFILES:
            original = json.dumps(PROFILES[build], sort_keys=True)
            choices = variants(build, PROFILES[build])
            signatures = set()
            for label, player in choices:
                self.assertEqual(
                    {k: v for k, v in player.items() if k != "rotation"},
                    {k: v for k, v in PROFILES[build].items() if k != "rotation"},
                    (build, label),
                )
                signature = json.dumps(player, sort_keys=True)
                self.assertNotIn(signature, signatures)
                signatures.add(signature)
            self.assertEqual(json.dumps(PROFILES[build], sort_keys=True), original)
            self.assertEqual(choices[0][0], "baseline")

    def test_feral_can_reintroduce_bite_after_omission(self):
        p = json.loads(json.dumps(PROFILES["feral"]))
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != 31018]
        choices = dict(variants("feral", p))
        self.assertTrue(any(spell_id(a) == 31018 for a in choices["rip4-bite40-builder9830-rake0"]["rotation"]["priorityList"]))
        finisher = next(a for a in choices["finish-bite-4s-3cp"]["rotation"]["priorityList"] if spell_id(a) == 31018)
        self.assertEqual(finisher["action"]["condition"]["and"]["vals"][0]["cmp"]["rhs"]["const"]["val"], "4s")

    def test_flame_shock_rank_updates_the_dot_condition(self):
        candidate = dict(variants("elemental", PROFILES["elemental"]))["flame-shock-1"]
        shock = next(a for a in candidate["rotation"]["priorityList"] if spell_id(a) == 8050)
        identifier = shock["action"]["condition"]["not"]["val"]["dotIsActive"]["spellId"]
        self.assertEqual(identifier, {"spellId": 8050})

    def test_rogue_resource_actions_do_not_wait_for_classic_ticks(self):
        candidate = dict(variants("combat", PROFILES["combat"]))["energy-cooldowns"]
        self.assertNotIn("timeToEnergyTick", json.dumps(candidate["rotation"]))
        actions = candidate["rotation"]["priorityList"]
        tea = next(a for a in actions if a["action"].get("castSpell", {}).get("spellId", {}).get("itemId") == 7676)
        rush = next(a for a in actions if spell_id(a) == 13750)
        self.assertEqual(tea["action"]["condition"]["cmp"]["rhs"]["const"]["val"], "10")
        self.assertEqual(rush["action"]["condition"]["cmp"]["rhs"]["const"]["val"], "40")


if __name__ == "__main__":
    unittest.main()
