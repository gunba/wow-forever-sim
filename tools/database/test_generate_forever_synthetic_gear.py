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
            row for row in generator.proposals()
            if row["Archetype"] == "Intellect / MP5 / crit"
            and row["SlotID"] == 1 and row["Material"] == "Cloth"
        )
        item = generator.record(proposal, 6)
        self.assertEqual(item["class"], 4)
        self.assertEqual(item["subclass"], 1)
        self.assertNotIn("weapon", item)
        self.assertNotEqual(item["modelIconReferenceItemID"], proposal["SourceID"])

    def test_trinket_sensitivity_cannot_enter_benchmark(self):
        proposals = generator.proposals()
        trinkets = [
            generator.record(row, i+1)
            for i, row in enumerate(proposals)
            if row["SlotID"] == 12
        ]
        self.assertEqual(len(trinkets), 16)
        self.assertEqual(sum(item["benchmarkEligible"] for item in trinkets), 8)
        self.assertTrue(all(item["setId"] == 0 and item["effects"] == [] for item in trinkets))

    def test_model_ids_do_not_overlap_real_catalog(self):
        real = set(generator.BY_ID)
        items = [
            generator.record(row, i+1)
            for i, row in enumerate(generator.proposals())
        ]
        self.assertTrue(all(item["id"] not in real for item in items))
        self.assertEqual(len({item["id"] for item in items}), len(items))


if __name__ == "__main__":
    unittest.main()
