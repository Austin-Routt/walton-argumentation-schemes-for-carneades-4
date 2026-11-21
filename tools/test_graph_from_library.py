#!/usr/bin/env python3
"""Tests of tools/graph_from_library.py (PLAN.md Phase 6: how the library is used).

    python3 tools/test_graph_from_library.py

For every scheme of walton.yml, the graph the tool makes passes the checks of
tools/validate_schemes.py with no error and carries no template it does not
use; the scheme block is walton.yml's, verbatim.
"""
import os
import subprocess
import sys
import tempfile
import unittest

import yaml

HERE = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, HERE)
from build_category_files import read_master  # noqa: E402
from graph_from_library import graph  # noqa: E402
from validate_schemes import Report, check_file  # noqa: E402

LIBRARY = os.path.join(os.path.dirname(HERE), "walton.yml")


class Graphs(unittest.TestCase):
    def issue_library(self, directory):
        path = os.path.join(directory, "issue-library.yml")
        with open(path, "w", encoding="utf-8") as stream:
            stream.write('''language:
  p/1: "%s is P."
  q/1: "%s is Q."
  act_on_presumption_false/1: "Act as if %s is false."
  act_on_presumption_true/1: "Act as if %s is true."
issue_schemes:
  practical_presumption: ["act_on_presumption_false(A)", "act_on_presumption_true(A)"]
argument_schemes:
  - id: sample
    meta:
      title: Generator control
      source: Handwritten generator test
    variables: [X]
    premises: [p(X)]
    conclusions: [q(X)]
''')
        return path

    def test_every_scheme(self):
        _, blocks, _ = read_master(LIBRARY)
        with tempfile.TemporaryDirectory() as td:
            for sid, block in blocks.items():
                with self.subTest(scheme=sid):
                    text = graph([sid], LIBRARY)
                    self.assertIn(block + "\n", text)
                    p = os.path.join(td, sid + ".yml")
                    with open(p, "w", encoding="utf-8") as f:
                        f.write(text)
                    rep = Report()
                    d = check_file(p, sid + ".yml", rep)
                    self.assertIsNotNone(d)
                    self.assertEqual([], [(i["code"], i["message"]) for i in rep.items
                                          if i["severity"] == "ERROR" or i["code"] == "TEMPLATE_UNUSED"])
                    self.assertEqual([sid], [s["id"] for s in d["argument_schemes"]])

    def test_several_schemes_share_templates(self):
        d = yaml.safe_load(graph(["practical_inference", "negative_consequences_reasoning", "practical_inference"], LIBRARY))
        self.assertEqual(["practical_inference", "negative_consequences_reasoning"], [s["id"] for s in d["argument_schemes"]])
        self.assertIn("should_do/2", d["language"])
        self.assertIn("negation", d["issue_schemes"])

    def test_issue_dependencies_survive_argument_selection(self):
        # Neither presumption predicate occurs in sample, but the
        # copied issue still needs both declarations for engine validation.
        with tempfile.TemporaryDirectory() as directory:
            d = yaml.safe_load(graph(["sample"], self.issue_library(directory)))
        self.assertIn("act_on_presumption_false/1", d["language"])
        self.assertIn("act_on_presumption_true/1", d["language"])
        self.assertEqual(["act_on_presumption_false(A)", "act_on_presumption_true(A)"],
                         d["issue_schemes"]["practical_presumption"])
        self.assertEqual(["sample"], [s["id"] for s in d["argument_schemes"]])

    def test_missing_issue_language_is_rejected(self):
        with tempfile.TemporaryDirectory() as td:
            d = yaml.safe_load(graph(["sample"], self.issue_library(td)))
            del d["language"]["act_on_presumption_true/1"]
            path = os.path.join(td, "missing-issue-language.yml")
            with open(path, "w", encoding="utf-8") as stream:
                yaml.safe_dump(d, stream, sort_keys=False, allow_unicode=True)
            report = Report()
            check_file(path, "missing-issue-language.yml", report)
            self.assertTrue(any(item["code"] == "ISSUE_TEMPLATE_MISSING" and
                                "act_on_presumption_true/1" in item["message"]
                                for item in report.items))

    def test_unknown_scheme(self):
        r = subprocess.run([sys.executable, "-I", os.path.join(HERE, "graph_from_library.py"), "no_such_scheme"],
                           capture_output=True, text=True)
        self.assertNotEqual(0, r.returncode)
        self.assertIn("no_such_scheme", r.stderr)


if __name__ == "__main__":
    unittest.main(verbosity=1)
