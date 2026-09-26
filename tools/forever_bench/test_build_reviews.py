import json
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

class BuildReviewTests(unittest.TestCase):
    def test_review_portal_discloses_current_warrior_rage_model(self):
        root = Path(__file__).resolve().parents[2]
        with tempfile.TemporaryDirectory() as directory:
            output = Path(directory) / "review"
            subprocess.run([
                sys.executable, str(root / "tools/forever_bench/build_review_site.py"),
                "--output", str(output),
            ], cwd=root, check=True, capture_output=True)
            html = (output / "index.html").read_text()
        self.assertIn('id="WAR-001"', html)
        self.assertIn("Damage-independent rage:", html)
        self.assertLess(html.index('id="questions"'), html.index('id="matrix"'))
        self.assertNotIn("Warrior rage still uses an inherited damage-based model", html)
        self.assertNotIn("Hunter pets still inherit no owner stats", html)
        self.assertNotIn("General haste does not shorten the default spell GCD", html)
        for route in ("tank_warrior", "protection_paladin", "feral_tank_druid"):
            self.assertIn(f'href="../{route}/?profile={route}__', html)

    def test_current_results_do_not_claim_historical_paired_gains(self):
        root = Path(__file__).resolve().parents[2]
        results = root / "artifacts/forever_dps_5min.json"
        with tempfile.TemporaryDirectory() as directory:
            output = Path(directory) / "reviews.md"
            subprocess.run([
                sys.executable, str(root / "tools/forever_bench/build_reviews.py"),
                "--results", str(results), "--output", str(output),
            ], cwd=root, check=True, capture_output=True)
            text = output.read_text()
        self.assertNotIn("Baseline DPS", text)
        self.assertNotIn("under the same engine and seed", text)
        self.assertIn("Standard error", text)
        self.assertIn("hypothetical modeled-only gear results", text)
        self.assertIn("Tank rows include frontal incoming attacks", text)
        races = {row["Race"] for row in json.loads(results.read_text())["Results"]}
        result_tables = [section.split("\n### ", 1)[0] for section in text.split("### Results\n")[1:]]
        rows = [line for table in result_tables for line in table.splitlines()
                if any(line.startswith(f"| {race} |") for race in races)]
        self.assertEqual(len(rows), len(json.loads(results.read_text())["Results"]))
