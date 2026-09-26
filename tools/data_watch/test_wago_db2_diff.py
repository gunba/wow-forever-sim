import unittest
from unittest.mock import patch

from wago_db2_diff import TABLES, report


class ClientDiffTest(unittest.TestCase):
    def test_failed_fetches_do_not_report_no_changes(self):
        def unavailable(_table, _build):
            raise OSError("offline")
        text = report("old", "new", unavailable)
        self.assertNotIn("no changes", text)
        self.assertIn("incomplete", text.lower())

    def test_effect_labels_use_spell_id_not_effect_row_id(self):
        data = {
            ("SpellName", "new"): "ID,Name_lang\n10,Wrong ability\n99,Correct ability\n",
            ("SpellEffect", "old"): "ID,SpellID,Value\n10,99,1\n",
            ("SpellEffect", "new"): "ID,SpellID,Value\n10,99,2\n",
        }
        with patch("wago_db2_diff.TABLES", ["SpellEffect"]):
            text = report("old", "new", lambda t, b: data[(t, b)])
        self.assertIn("Correct ability", text)
        self.assertNotIn("Wrong ability", text)

    def test_spell_costs_trait_curves_and_item_set_thresholds_are_in_scope(self):
        self.assertTrue({"SpellPower", "SpellCategories", "SpellTargetRestrictions",
                         "TraitDefinitionEffectPoints", "CurvePoint", "ItemSetSpell",
                         "ItemEffect", "ItemXItemEffect"} <= set(TABLES))

    def test_html_response_is_not_an_empty_valid_table(self):
        with patch("wago_db2_diff.TABLES", ["SpellEffect"]):
            text = report("old", "new", lambda _t, _b: "<html>Access denied</html>")
        self.assertNotIn("no changes", text)
        self.assertIn("incomplete", text.lower())


if __name__ == "__main__":
    unittest.main()
