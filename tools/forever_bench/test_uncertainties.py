import copy
import json
from pathlib import Path
import re
import tempfile
import unittest

from uncertainties import REGISTER, load_register, markdown, render_html


class UncertaintyRegisterTests(unittest.TestCase):
    def test_duplicate_current_id_is_rejected(self):
        # The former hand-maintained checklist reused T50, T51 and T52.
        data = copy.deepcopy(load_register())
        data["items"].append(data["items"][0])
        with tempfile.TemporaryDirectory() as directory:
            path = Path(directory) / "register.json"
            path.write_text(json.dumps(data))
            with self.assertRaisesRegex(ValueError, "Duplicate uncertainty ID"):
                load_register(path)

    def test_colliding_old_topics_keep_distinct_ids(self):
        data = load_register()
        for legacy in ("T50", "T51", "T52"):
            items = [item for item in data["items"]
                     if any(legacy in alias for alias in item.get("legacy", []))]
            self.assertGreaterEqual(len(items), 2, legacy)
            self.assertEqual(len({item["id"] for item in items}), len(items))
            for item in items:
                aliases = [alias for alias in item["legacy"] if legacy in alias]
                self.assertTrue(all("(" in alias for alias in aliases), aliases)

    def test_generated_views_keep_every_record(self):
        data = load_register()
        html = render_html(data)
        text = markdown(data) + "\n"
        self.assertEqual(REGISTER.with_suffix(".md").read_text(), text)
        for item in data["items"]:
            self.assertEqual(html.count(f'id="{item["id"]}"'), 1)
            self.assertEqual(text.count(f'### {item["id"]} —'), 1)
        # Corrected effects must not survive as active tasks from older audits.
        lookup = {item["id"]: item for item in data["items"]}
        self.assertEqual(lookup["DONE-003"]["status"], "resolved")
        self.assertEqual(lookup["DONE-007"]["status"], "resolved")
        self.assertEqual(lookup["WAR-011"]["status"], "resolved")
        self.assertIn("provisionally", lookup["HUN-001"]["current_model"])

    def test_legacy_checklist_has_a_disposition(self):
        aliases = " ".join(alias for item in load_register()["items"]
                           for alias in item.get("legacy", []))
        self.assertTrue({f"{number:02}" for number in range(1, 61)}
                        <= set(re.findall(r"T(\d{2})(?!\d)", aliases)))


if __name__ == "__main__":
    unittest.main()
