import unittest
from unittest.mock import patch

from spell_client import Client


class ClassSpellDiscoveryTest(unittest.TestCase):
    def test_zero_class_mask_does_not_hide_new_death_ranks(self):
        client = Client.__new__(Client)
        client.build = "fixture"
        rows = [
            {"Spell": "1309636", "ClassMask": "0", "SkillLine": "78"},
            {"Spell": "1309635", "ClassMask": "0", "SkillLine": "78"},
            {"Spell": "10947", "ClassMask": "16", "SkillLine": "78"},
            {"Spell": "1752", "ClassMask": "8", "SkillLine": "38"},
        ]
        with patch("spell_client.table", return_value=rows):
            self.assertEqual(client.learned("priest"), [10947, 1309635, 1309636])


if __name__ == "__main__":
    unittest.main()
