import hashlib
import json
from pathlib import Path
import unittest

from import_forever_ratings import read_coefficients
from import_forever_vendor import build_stats


ROOT = Path(__file__).resolve().parents[2]


class ForeverRatingsTest(unittest.TestCase):
    def test_client_data_and_provenance(self):
        data = (ROOT / "assets/db_inputs/forever_combat_ratings_1.60.1.69913.txt").read_bytes()
        imported = json.loads((ROOT / "assets/db_inputs/forever_combat_ratings.json").read_text())
        rates = read_coefficients(data, 60)
        self.assertEqual(hashlib.sha256(data).hexdigest(), imported["source"]["sha256"])
        self.assertEqual(rates, imported["coefficients"])
        for school in ("Melee", "Ranged", "Spell"):
            self.assertEqual(rates[f"Hit - {school}"], 10)
            self.assertEqual(rates[f"Crit - {school}"], 14)
            self.assertEqual(rates[f"Haste - {school}"], 10)
        self.assertEqual(rates["Dodge"], 12)
        self.assertEqual(rates["Parry"], 15)
        # These coefficients do not follow TBC's spell-hit or level scaling.
        self.assertEqual(read_coefficients(data, 20), rates)

    def test_missing_level(self):
        with self.assertRaisesRegex(ValueError, "exactly one row"):
            read_coefficients(b"Level\tHit - Melee\n1\t10\n", 60)

    def test_generic_item_ratings_are_not_doubled(self):
        stats = build_stats({"stats": {
            "hitrtng": 10, "critstrkrtng": 14, "hastertng": 20,
            "dodgertng": 12, "parryrtng": 15, "blockrtng": 5,
        }})
        self.assertEqual(stats[18], 1)  # Shared equipment hit.
        self.assertEqual(stats[19], 1)  # Shared equipment crit.
        self.assertEqual(stats[13], 0)  # Do not add a second copy for spells.
        self.assertEqual(stats[14], 0)
        self.assertEqual(stats[20], 2)
        self.assertEqual(stats[15], 2)
        self.assertEqual(stats[31], 1)
        self.assertEqual(stats[32], 1)
        self.assertEqual(stats[29], 1)

    def test_invalid_conversion(self):
        data = (ROOT / "assets/db_inputs/forever_combat_ratings_1.60.1.69913.txt").read_bytes()
        header, *rows = data.decode().splitlines()
        columns = header.split("\t")
        for index, row in enumerate(rows):
            values = row.split("\t")
            if values[0] == "60":
                values[columns.index("Hit - Spell")] = "0"
                rows[index] = "\t".join(values)
        with self.assertRaisesRegex(ValueError, "Hit - Spell"):
            read_coefficients(("\n".join([header, *rows]) + "\n").encode(), 60)


if __name__ == "__main__":
    unittest.main()
