#!/usr/bin/env python3
"""Mutation and control tests for the cross-toolchain result gate."""
import contextlib
import io
import json
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

import compare_results as comparator


class CompareResultsTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.base = {"file": "walton.yml", "scheme": "s", "check": "c", "pass": True}

    def write(self, name, records):
        path = self.root / name
        path.write_text("\n".join(json.dumps(r) for r in records) + "\n", encoding="utf-8")
        return path

    def compare(self, first, second):
        out, err = io.StringIO(), io.StringIO()
        with contextlib.redirect_stdout(out), contextlib.redirect_stderr(err):
            code = comparator.main(["compare_results.py", str(first), str(second)])
        return code, out.getvalue(), err.getvalue()

    def test_valid_reordered_reports_preserve_false_and_distinct_keys(self):
        rows = [self.base, dict(self.base, check="other", **{"pass": False}),
                dict(self.base, file="", scheme="(engine)"), dict(self.base, scheme="t")]
        a, b = self.write("a", rows), self.write("b", list(reversed(rows)))
        loaded = comparator.load(a)
        self.assertEqual(len(loaded), 4)
        self.assertIs(loaded[("walton.yml", "s", "other")], False)
        code, out, err = self.compare(a, b)
        self.assertEqual((code, err), (0, ""))
        self.assertIn("4 checks, 3 pass", out)

    def test_duplicate_same_or_different_outcome_is_bad_input_on_either_side(self):
        clean = self.write("clean", [self.base])
        for outcome in (True, False):
            duplicate = self.write("duplicate", [self.base, dict(self.base, **{"pass": outcome})])
            for a, b in ((duplicate, clean), (clean, duplicate)):
                with self.subTest(outcome=outcome, side=a.name):
                    code, out, err = self.compare(a, b)
                    self.assertEqual(code, 2)
                    self.assertIn("duplicate check key", err)
                    self.assertIn("first at line 1", err)
                    self.assertNotIn("agree", out)

    def test_duplicate_line_numbers_include_blank_lines(self):
        path = self.root / "lines"
        row = json.dumps(self.base)
        path.write_text("\n" + row + "\n\n" + row + "\n", encoding="utf-8")
        with self.assertRaisesRegex(ValueError, r":4: duplicate .*first at line 2"):
            comparator.load(path)

    def test_pass_requires_actual_json_boolean(self):
        for value in ("true", "false", "", 1, 0, None, [], {}):
            with self.subTest(value=value):
                path = self.write("bad-pass", [dict(self.base, **{"pass": value})])
                with self.assertRaisesRegex(ValueError, "pass must be a JSON boolean"):
                    comparator.load(path)

    def test_missing_pass_is_rejected(self):
        row = dict(self.base)
        del row["pass"]
        with self.assertRaisesRegex(ValueError, "pass must be a JSON boolean"):
            comparator.load(self.write("missing-pass", [row]))

    def test_identity_fields_are_required_strings(self):
        for field in ("file", "scheme", "check"):
            for value in (None, 1, [], {}):
                with self.subTest(field=field, value=value):
                    row = dict(self.base, **{field: value})
                    with self.assertRaisesRegex(ValueError, field + " must be a JSON string"):
                        comparator.load(self.write("bad-key", [row]))
            row = dict(self.base)
            del row[field]
            with self.assertRaisesRegex(ValueError, field + " must be a JSON string"):
                comparator.load(self.write("missing-key", [row]))

    def test_nonobject_records_are_bad_input(self):
        for row in ([], None, 1, True, "record"):
            with self.subTest(row=row):
                with self.assertRaisesRegex(ValueError, "result must be a JSON object"):
                    comparator.load(self.write("not-object", [row]))

    def test_key_difference_fails_agreement(self):
        a = self.write("a", [self.base])
        b = self.write("b", [dict(self.base, check="different")])
        code, out, err = self.compare(a, b)
        self.assertEqual((code, err), (1, ""))
        self.assertIn("only in the first run: 1", out)
        self.assertIn("only in the second run: 1", out)

    def test_outcome_difference_fails_agreement(self):
        a = self.write("a", [self.base])
        b = self.write("b", [dict(self.base, **{"pass": False})])
        code, out, err = self.compare(a, b)
        self.assertEqual((code, err), (1, ""))
        self.assertIn("different outcome: 1", out)

    def test_unreadable_file_is_bad_input(self):
        code, out, err = self.compare(self.root / "absent", self.write("b", [self.base]))
        self.assertEqual(code, 2)
        self.assertIn("absent", err)
        self.assertNotIn("agree", out)

    def test_cli_malformed_json_exits_two(self):
        bad = self.root / "bad-json"
        bad.write_text("{\n", encoding="utf-8")
        good = self.write("good", [self.base])
        result = subprocess.run([sys.executable, str(Path(comparator.__file__)), str(bad), str(good)],
                                capture_output=True, text=True, check=False)
        self.assertEqual(result.returncode, 2)
        self.assertIn(str(bad) + ":1:", result.stderr)
        self.assertNotIn("Traceback", result.stderr)


if __name__ == "__main__":
    unittest.main()
