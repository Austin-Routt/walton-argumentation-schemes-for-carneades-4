#!/usr/bin/env python3
"""Tests of tools/validate_schemes.py (PLAN.md 3.10).

Ported from the third review's run_python_audit.py (20 schema probes) and
full_workflow_mutations.py (two semantic mutations), with the probes turned
into assertions: each malformed fixture must be refused with an ERROR, the
valid control must pass, and a mutated copy of walton.yml must fail.

Usage:  python3 tools/test_validate_schemes.py [repo_root]
"""
import copy
import os
import shutil
import subprocess
import sys
import tempfile
import unittest

import yaml

ROOT = os.path.abspath(sys.argv[1]) if len(sys.argv) > 1 and not sys.argv[1].startswith("-") else \
    os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
sys.argv[1:] = [a for a in sys.argv[1:] if a.startswith("-")]
sys.path.insert(0, os.path.join(ROOT, "tools"))
import validate_schemes as v  # noqa: E402

BASELINE = {
    "language": {"p/1": "%s is P.", "q/1": "%s is Q."},
    "argument_schemes": [{
        "id": "probe", "meta": {"title": "Probe", "source": "Synthetic fixture"},
        "variables": ["X"], "premises": ["p(X)"], "conclusions": ["q(X)"]}],
}


def check(content):
    with tempfile.TemporaryDirectory(prefix="walton-probe-") as td:
        p = os.path.join(td, "probe.yml")
        with open(p, "w", encoding="utf-8") as f:
            f.write(content)
        rep = v.Report()
        v.check_file(p, "probe.yml", rep)
        return rep


def fixture(mutate=None):
    d = copy.deepcopy(BASELINE)
    if mutate:
        mutate(d)
    return yaml.safe_dump(d, sort_keys=False, allow_unicode=True)


def scheme(key, value):
    return lambda d: d["argument_schemes"][0].__setitem__(key, value)


class SchemaProbes(unittest.TestCase):
    def refused(self, content, name):
        rep = check(content)
        self.assertTrue(rep.errors(), f"{name}: no ERROR for a malformed fixture: {rep.items}")

    def test_valid_control(self):
        self.assertEqual(check(fixture()).errors(), [])

    def test_probes(self):
        standard = fixture()
        probes = [
            ("unknown_scheme_key", fixture(scheme("assumption", ["q(X)"]))),
            ("undeclared_premise_variable", fixture(scheme("premises", ["p(Y)"]))),
            ("duplicate_variables", fixture(scheme("variables", ["X", "X"]))),
            ("invalid_declared_variable", fixture(scheme("variables", ["X", "notAVariable"]))),
            ("unhashable_variable", fixture(scheme("variables", ["X", {"bad": "entry"}]))),
            ("malformed_deletions_scalar", fixture(scheme("deletions", "p(X)"))),
            ("malformed_guards_scalar", fixture(scheme("guards", "X >>> 1"))),
            ("malformed_guard_term", fixture(scheme("guards", ["X >>> 1"]))),
            ("undeclared_deletion_variable", fixture(scheme("deletions", ["p(Y)"]))),
            ("undeclared_guard_variable", fixture(scheme("guards", ["Y > 0"]))),
            ("bare_variable_assumption", fixture(scheme("assumptions", ["X"]))),
            ("bare_variable_exception", fixture(scheme("exceptions", ["X"]))),
            ("undefined_weighing_function", fixture(scheme("weight", "definitely_missing_weight"))),
            ("roles_type", fixture(scheme("roles", 17))),
            ("reserved_boolean_predicate", fixture(lambda d: (d["language"].__setitem__("true/1", "%s is true."),
                                                              d["argument_schemes"][0].__setitem__("conclusions", ["true(X)"])))),
            ("repeated_template_argument", fixture(lambda d: d["language"].__setitem__("p/1", "%[1]s agrees with %[1]s."))),
            ("literal_percent", fixture(lambda d: d["language"].__setitem__("p/1", "%s agrees 100%%."))),
            ("duplicate_conclusions_key", standard + "  conclusions: [p(X)]\n"),
            ("inline_duplicate_language_key",
             'language: {"p/1": "%s is old.", "p/1": "%s is P.", "q/1": "%s is Q."}\n'
             + yaml.safe_dump({"argument_schemes": BASELINE["argument_schemes"]}, sort_keys=False)),
            # the profile rules of PLAN.md 3.10
            ("range_restriction", fixture(scheme("conclusions", ["q(Y)"])) .replace("- X\n", "- X\n  - Y\n", 1)),
            ("assumption_repeats_premise", fixture(scheme("assumptions", ["p(X)"]))),
            ("is_as_predicate", fixture(lambda d: (d["language"].__setitem__("is/2", "%s is %s."),
                                                   d["argument_schemes"][0].__setitem__("conclusions", ["is(X,X)"])))),
        ]
        for name, content in probes:
            with self.subTest(name):
                self.refused(content, name)


def mutated_repo(sid, old, new):
    """A copy of the repository with one statement of one scheme of walton.yml changed."""
    td = tempfile.mkdtemp(prefix="walton-mutation-")
    target = os.path.join(td, "repo")
    shutil.copytree(ROOT, target, ignore=shutil.ignore_patterns(".git", "*.zip", "__pycache__", "third_party", "tools"))
    shutil.copytree(os.path.join(ROOT, "tools"), os.path.join(target, "tools"),
                    ignore=shutil.ignore_patterns("__pycache__", "vv"))
    p = os.path.join(target, "walton.yml")
    text = open(p, encoding="utf-8").read()
    start = text.index(f"  - id: {sid}\n")
    end = text.find("  - id:", start + 1)
    end = len(text) if end == -1 else end
    block = text[start:end]
    assert block.count(old) == 1, (sid, old)
    open(p, "w", encoding="utf-8").write(text[:start] + block.replace(old, new, 1) + text[end:])
    return td, target


class Mutations(unittest.TestCase):
    """A deliberately wrong scheme must not survive the documented workflow."""
    def test_mutations(self):
        for name, sid, old, new in [
            ("danger_polarity", "danger", "      - ¬ought_to_bring_about(X)", "      - ought_to_bring_about(X)"),
            ("division_roles", "division_generic", "      - all_parts_property(A,P)", "      - all_parts_property(P,A)"),
            ("precedent_negation_dropped", "from_precedent", "      - ¬has_property(A,G)", "      - has_property(A,G)"),
            ("sacrifice_arguments_swapped", "sacrifice", "      - sacrifice_value(A,V)", "      - sacrifice_value(V,A)"),
            ("ad_hominem_premise_dropped", "ad_hominem", "      - advocates_conclusion(A,S,C)\n", ""),
        ]:
            with self.subTest(name):
                td, target = mutated_repo(sid, old, new)
                try:
                    subprocess.run([sys.executable, "tools/build_category_files.py", target], cwd=target, capture_output=True)
                    r = subprocess.run([sys.executable, "tools/validate_schemes.py", target], cwd=target, capture_output=True, text=True)
                    self.assertNotEqual(r.returncode, 0, f"{name} survived the validator:\n{r.stdout[-800:]}")
                finally:
                    shutil.rmtree(td)


class Robustness(unittest.TestCase):
    def test_missing_scheme_id_does_not_crash(self):
        master = yaml.safe_load(open(os.path.join(ROOT, "walton.yml"), encoding="utf-8"))
        master["argument_schemes"][0].pop("id")
        with tempfile.TemporaryDirectory(prefix="walton-missing-id-") as td:
            with open(os.path.join(td, "walton.yml"), "w", encoding="utf-8") as f:
                yaml.safe_dump(master, f, sort_keys=False, allow_unicode=True)
            r = subprocess.run([sys.executable, os.path.join(ROOT, "tools/validate_schemes.py"), td, "--json"],
                               capture_output=True, text=True)
            self.assertNotEqual(r.returncode, 0)
            self.assertNotIn("Traceback", r.stderr)

    def test_library_is_clean(self):
        r = subprocess.run([sys.executable, os.path.join(ROOT, "tools/validate_schemes.py"), ROOT], capture_output=True, text=True)
        self.assertEqual(r.returncode, 0, r.stdout[-2000:])


if __name__ == "__main__":
    unittest.main(verbosity=1)
