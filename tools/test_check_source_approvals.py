#!/usr/bin/env python3
"""Synthetic approval-integrity controls and mutations; creates no real approvals."""
import copy
import json
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

import yaml

sys.path.insert(0, str(Path(__file__).resolve().parent))
import check_source_approvals as gate


class ApprovalIntegrity(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.addCleanup(self.tmp.cleanup)
        self.root = Path(self.tmp.name)
        self.master = {"language": {"p/1": "%s is P.", "q/1": "%s is Q.", "unused/1": "%s is unused."},
                       "issue_schemes": {"negation": ["S", "¬S"]},
                       "argument_schemes": [{"id": "probe", "meta": {"title": "Synthetic fixture"},
                                             "variables": ["X"], "premises": ["p(X)"], "conclusions": ["q(X)"]}]}
        self.row = {"record": "probe", "origin": "book", "entry": "1", "claim": "interpretation",
                    "stage": "example-passed", "questions": [{"n": 1, "kind": "premise", "as": "p(X)"}]}
        self.ledger = {"records": [self.row]}
        self.registry = {"schema": gate.SCHEMA, "approvals": [{
            "record": "probe", "fingerprint": gate.semantic_fingerprint(self.master, self.master["argument_schemes"][0], self.row),
            "evidence": ["Synthetic unit-test evidence, not a real review approval."], "examples": ["probe/book"]}]}
        self.registry_path = self.root / "docs/reviews/independent-approvals.json"
        self.save()
        self.write("tools/vv/cases/book.yml", {"cases": [{"name": "probe/book", "source": "Synthetic fixture"}]})

    def write(self, path, data):
        p = self.root / path
        p.parent.mkdir(parents=True, exist_ok=True)
        p.write_text(yaml.safe_dump(data, sort_keys=False, allow_unicode=True), encoding="utf-8")

    def save(self):
        self.write("walton.yml", self.master)
        self.write("tools/vv/source_ledger.yml", self.ledger)
        self.registry_path.parent.mkdir(parents=True, exist_ok=True)
        self.registry_path.write_text(json.dumps(self.registry), encoding="utf-8")

    def errors(self, stable=False):
        self.save()
        return "\n".join(gate.audit(self.root, stable)[1])

    def test_unchanged_approved_record_passes(self):
        self.assertEqual("", self.errors(stable=True))

    def test_changed_approved_scheme_fails(self):
        self.master["argument_schemes"][0]["conclusions"] = ["¬q(X)"]
        self.assertIn("fingerprint is stale", self.errors())

    def test_all_raw_scheme_fields_are_bound(self):
        self.master["argument_schemes"][0]["meta"]["note"] = "Changed raw scheme metadata"
        self.assertIn("fingerprint is stale", self.errors())

    def test_missing_approval_fails_at_every_promoted_stage(self):
        self.registry["approvals"] = []
        for stage in gate.STAGES[2:]:
            self.row["stage"] = stage
            with self.subTest(stage=stage):
                self.assertIn("has no independent approval", self.errors())

    def test_changed_question_scope_and_source_contract_fail(self):
        mutations = [lambda: self.row["questions"][0].update({"as": "p(Y)"}),
                     lambda: self.row["questions"][0].update({"kind": "assumption"}),
                     lambda: self.row.update({"origin": "upstream"}),
                     lambda: self.row.update({"claim": "policy"}),
                     lambda: self.row.update({"source_questions": 2})]
        original = copy.deepcopy(self.row)
        for mutate in mutations:
            self.row.clear()
            self.row.update(copy.deepcopy(original))
            mutate()
            self.assertIn("fingerprint is stale", self.errors())

    def test_relevant_templates_and_global_issues_are_bound(self):
        self.master["language"]["q/1"] = "%s is not Q."
        self.assertIn("fingerprint is stale", self.errors())
        self.master["language"]["q/1"] = "%s is Q."
        self.master["issue_schemes"] = {}
        self.assertIn("fingerprint is stale", self.errors())

    def test_irrelevant_template_and_ledger_notes_do_not_stale_approval(self):
        self.master["language"]["unused/1"] = "Reworded unused %s."
        self.row.update({"note": "Editorial clarification", "stage": "interpretation-approved"})
        self.row["questions"][0]["note"] = "Editorial CQ note"
        self.assertEqual("", self.errors())

    def test_mapping_key_order_does_not_change_fingerprint(self):
        self.master["argument_schemes"][0] = dict(reversed(list(self.master["argument_schemes"][0].items())))
        self.assertEqual("", self.errors())

    def test_unknown_and_duplicate_approval_ids_fail(self):
        self.registry["approvals"][0]["record"] = "unknown"
        self.assertIn("unknown record", self.errors())
        self.registry["approvals"][0]["record"] = "probe"
        self.registry["approvals"].append(copy.deepcopy(self.registry["approvals"][0]))
        with self.assertRaisesRegex(ValueError, "duplicate"):
            self.errors()

    def test_unknown_and_duplicate_ledger_and_scheme_ids_fail(self):
        self.row["record"] = "unknown"
        with self.assertRaisesRegex(ValueError, "unknown"):
            self.errors()
        self.row["record"] = "probe"
        self.ledger["records"].append(copy.deepcopy(self.row))
        with self.assertRaisesRegex(ValueError, "duplicate"):
            self.errors()
        self.ledger["records"].pop()
        self.master["argument_schemes"].append(copy.deepcopy(self.master["argument_schemes"][0]))
        with self.assertRaisesRegex(ValueError, "duplicate"):
            self.errors()

    def test_evidence_and_existing_examples_are_required(self):
        for field in ("evidence", "examples"):
            old = self.registry["approvals"][0][field]
            self.registry["approvals"][0][field] = []
            self.assertIn(f"{field} must be a nonempty", self.errors())
            self.registry["approvals"][0][field] = old
        self.registry["approvals"][0]["examples"] = ["nonexistent/book"]
        self.assertIn("unknown examples", self.errors())

    def test_stale_recorded_approval_fails_even_without_promotion(self):
        self.row["stage"] = "mapped"
        self.master["argument_schemes"][0]["premises"] = ["q(X)"]
        self.assertIn("fingerprint is stale", self.errors())

    def test_empty_or_missing_registry_with_unpromoted_rows_passes(self):
        self.registry["approvals"] = []
        self.row["stage"] = "transcription-checked"
        self.assertEqual("", self.errors())
        self.registry_path.unlink()
        self.assertEqual([], gate.audit(self.root)[1])
        self.assertIn("requires example-passed", "\n".join(gate.audit(self.root, True)[1]))

    def test_cli_repo_forms_inventory_and_stable_mode(self):
        self.registry["approvals"] = []
        self.row["stage"] = "mapped"
        self.save()
        self.registry_path.unlink()
        command = [sys.executable, str(Path(gate.__file__).resolve())]
        for args in ([str(self.root)], ["--repo", str(self.root)]):
            result = subprocess.run(command + args, capture_output=True, text=True)
            self.assertEqual(0, result.returncode, result.stderr)
        result = subprocess.run(command + ["--repo", str(self.root), "--require-stable"], capture_output=True, text=True)
        self.assertEqual(1, result.returncode)
        self.assertIn("requires example-passed", result.stderr)
        result = subprocess.run(command + ["--repo", str(self.root), "--fingerprints"], capture_output=True, text=True)
        self.assertEqual(0, result.returncode, result.stderr)
        self.assertTrue(json.loads(result.stdout)["probe"].startswith("sha256:"))
        self.assertFalse(self.registry_path.exists(), "inventory mode created an approval")


if __name__ == "__main__":
    unittest.main(verbosity=1)
