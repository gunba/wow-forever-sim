"""Regression coverage for the modeled catalog's source-type and budget guards."""

import importlib.util
from pathlib import Path
import sys
import unittest


HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE))
spec = importlib.util.spec_from_file_location(
    "generate_forever_synthetic_gear", HERE / "generate_forever_synthetic_gear.py"
)
generator = importlib.util.module_from_spec(spec)
spec.loader.exec_module(generator)


class SyntheticGearTest(unittest.TestCase):
    def test_projected_armor_never_inherits_weapon_class(self):
        # The caster MP5 armor is derived from a one-hand mace.
        proposal = next(
            row for row in generator.proposals_v2()
            if row["Archetype"] == "Intellect / MP5 / crit"
            and row["SlotID"] == 1 and row["Material"] == "Cloth"
        )
        item = generator.record(proposal, 6)
        self.assertEqual(item["class"], 4)
        self.assertEqual(item["subclass"], 1)
        self.assertNotIn("weapon", item)
        self.assertNotEqual(item["modelIconReferenceItemID"], proposal["SourceID"])

    def test_trinket_sensitivity_cannot_enter_benchmark(self):
        proposals = generator.proposals_v2()
        trinkets = [
            generator.record(row, i+1)
            for i, row in enumerate(proposals)
            if row["SlotID"] == 12
        ]
        self.assertEqual(len(trinkets), 19)
        self.assertEqual(sum(item["benchmarkEligible"] for item in trinkets), 11)
        self.assertTrue(all(item["setId"] == 0 and item["effects"] == [] for item in trinkets))

    def test_model_ids_do_not_overlap_real_catalog(self):
        real = set(generator.BY_ID)
        items = [
            generator.record(row, i+1)
            for i, row in enumerate(generator.proposals_v2())
        ]
        self.assertTrue(all(item["id"] not in real for item in items))
        self.assertTrue(all(item["id"] >= 920_000_001 for item in items))
        self.assertEqual(len({item["id"] for item in items}), len(items))

    def test_caster_mirror_has_same_modeled_budget_and_numeric_name(self):
        proposals = generator.proposals_v2()
        physical = next(p for p in proposals if p["SlotID"] == 7 and p["Material"] == "Plate"
                        and p["Archetype"] == "Strength / agility / crit / hit")
        caster = next(p for p in proposals if p["SlotID"] == 7 and p["Material"] == "Cloth"
                      and "(matched cost)" in p["Archetype"])
        self.assertAlmostEqual(physical["ProposedBudget"], caster["ProposedBudget"], places=4)
        item = generator.record(caster, 400)
        self.assertIn(f"+{caster['Stats']['SpellPower']} SP", item["name"])
        self.assertIn(f"+{caster['Stats']['Intellect']} Int", item["name"])

    def test_fixed_stat_relics_and_damage_wand_fill_previous_gaps(self):
        proposals = generator.proposals_v2()
        relics = [generator.record(row, i+1) for i, row in enumerate(proposals)
                  if row["SlotID"] == 28]
        self.assertEqual({item["subclass"] for item in relics}, {7, 8, 9})
        self.assertTrue(all(item["class"] == 4 and not item["effects"] for item in relics))
        wands = [generator.record(row, i+1) for i, row in enumerate(proposals)
                 if row["SlotID"] == 26 and "Spell-power wand" in row["Archetype"]]
        self.assertEqual(len(wands), 3)
        self.assertTrue(all(item["stats"].get("SpellPower", 0) > 0 for item in wands))

    def test_high_hit_caster_mirrors_retain_source_hit(self):
        proposals = generator.proposals_v2()
        for label, material in (("Spell power / crit / hit (matched cost)", "Cloth"),
                                ("Strength / spell power / hit (matched cost)", "Plate")):
            head = next(row for row in proposals if row["Archetype"] == label
                        and row["SlotID"] == 1 and row["Material"] == material)
            self.assertEqual(head["SourceID"], 279253)
            self.assertGreaterEqual(head["Stats"]["HitRating"], 20)
            self.assertGreater(head["Stats"]["SpellPower"], 0)
            self.assertNotIn("HitRating", generator.BY_ID[head["SourceID"]]["stats"])

    def test_projected_armor_respects_material_stat_patterns(self):
        proposals = generator.proposals_v2()
        body = [row for row in proposals if row["SlotID"] in generator.BODY_SLOTS]
        self.assertFalse(any(row["Material"] == "Leather" and row["Stats"].get("Strength")
                             for row in body))
        self.assertFalse(any(row["Material"] == "Cloth" and row["Stats"].get("Strength")
                             for row in body))
        mail_strength = [row for row in body if row["Material"] == "Mail"
                         and row["Stats"].get("Strength")]
        self.assertTrue(mail_strength)
        self.assertEqual({row["SourceID"] for row in mail_strength}, {20203, 22676})
        projected_hit = [row for row in mail_strength
                         if row["Archetype"] == "Strength / intellect / crit / hit (mail)"]
        self.assertEqual(len(projected_hit), len(generator.BODY_SLOTS))
        self.assertTrue(all(row["Stats"].get("HitRating", 0) > 0 and
                            row["SourceID"] == 22676 and
                            abs(row["ProposedBudget"] - 1) < 1e-8
                            for row in projected_hit))
        self.assertNotIn("HitRating", generator.BY_ID[22676]["stats"])


if __name__ == "__main__":
    unittest.main()
