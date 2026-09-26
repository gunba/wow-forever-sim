import json
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

from build_display import expected_roster


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
        self.assertIn("earlier mechanics revision", text)
        races = {row["Race"] for row in json.loads(results.read_text())["Results"]}
        rows = [line for line in text.splitlines() if any(line.startswith(f"| {race} |") for race in races)]
        self.assertEqual(len(rows), len(expected_roster()))
