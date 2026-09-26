import struct
import tempfile
from pathlib import Path
import unittest

from damage_meter import records, scrub


def string(value):
    raw = value.encode("utf-8")
    return struct.pack("<H", len(raw) + 1) + raw + b"\0"


class DamageMeterScrubTest(unittest.TestCase):
    def test_short_and_utf8_names_preserve_offsets_and_damage(self):
        for name in ("Al", "Björn", "李雷"):
            with self.subTest(name=name), tempfile.TemporaryDirectory() as directory:
                blob = string(name) + string("ROGUE") + struct.pack("<7I", 2, 1752, 1752, 36, 0, 0, 0)
                blob += string(name)
                output = Path(directory) / "scrubbed.bin"
                names = scrub(blob, records(blob), output)
                cleaned = output.read_bytes()
                self.assertEqual(len(names), 1)
                self.assertNotIn(string(name), cleaned)
                self.assertEqual(len(cleaned), len(blob))
                self.assertEqual(records(cleaned)[0][6], 36)


if __name__ == "__main__":
    unittest.main()
