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
import apply_source_review_metadata as metadata


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
                    "stage": "example-passed", "questions": [{"n": 1, "kind": "premise", "as": "p(X)"}],
                    "review": {"by": "Synthetic reviewer", "date": "2026-11-21", "decision": "approved",
                               "independent_of_author": True, "independence_basis": "Synthetic non-author fixture."}}
        self.ledger = {"records": [self.row]}
        self.registry = {"schema": gate.SCHEMA, "review_date": "2026-11-21", "approvals": [{
            "record": "probe", "fingerprint": gate.semantic_fingerprint(self.master, self.master["argument_schemes"][0], self.row),
            "evidence": ["Synthetic unit-test evidence, not a real review approval."], "examples": ["probe/book"],
            "review": dict(self.row["review"])}]}
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
                     lambda: self.row.update({"source_questions": 2}),
                     lambda: self.row.update({"comparison": "value-based-choice"})]
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
        self.row["review"] = {"by": "Synthetic reviewer", "date": "2026-11-21", "decision": "hold", "note": "Needs source adjudication."}
        self.assertEqual("", self.errors())
        self.registry_path.unlink()
        self.assertEqual([], gate.audit(self.root)[1])
        self.assertIn("requires example-passed", "\n".join(gate.audit(self.root, True)[1]))

    def test_cli_repo_forms_inventory_and_stable_mode(self):
        self.registry["approvals"] = []
        self.row["stage"] = "mapped"
        self.row.pop("review")
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

    def test_review_attribution_mutations_fail_without_semantic_hash_change(self):
        original = dict(self.row["review"])
        for field, value in (("by", "Someone else"), ("date", "2026-11-22"), ("decision", "hold"),
                             ("independent_of_author", False), ("independence_basis", "Another review")):
            with self.subTest(field=field):
                self.row["review"] = dict(original, **{field: value})
                self.assertEqual(self.registry["approvals"][0]["fingerprint"],
                                 gate.semantic_fingerprint(self.master, self.master["argument_schemes"][0], self.row))
                self.assertIn("review metadata differs", self.errors())
        self.row.pop("review")
        self.assertIn("requires review decision approved", self.errors())

    def test_matching_invalid_review_snapshot_does_not_authorize_approval(self):
        for field, value, message in (("by", " ", "named reviewer"),
                                      ("date", "2026-02-30", "date YYYY-MM-DD"),
                                      ("date", "20261121", "date YYYY-MM-DD"),
                                      ("decision", "hold", "decision approved"),
                                      ("independent_of_author", False, "explicit non-author acceptance"),
                                      ("independent_of_author", "true", "explicit non-author acceptance"),
                                      ("independence_basis", " ", "independence evidence basis")):
            with self.subTest(field=field, value=value):
                stamp = {"by": "Synthetic reviewer", "date": "2026-11-21", "decision": "approved",
                         "independent_of_author": True, "independence_basis": "Synthetic non-author fixture.", field: value}
                self.row["review"] = dict(stamp)
                self.registry["approvals"][0]["review"] = dict(stamp)
                self.assertIn(message, self.errors())

    def test_review_note_is_editorial_but_withdrawn_stage_is_not_approval(self):
        self.row["review"]["note"] = "Additional evidence locator."
        self.assertEqual("", self.errors())
        self.row["stage"] = "transcription-checked"
        self.assertIn("approving review is inconsistent", self.errors())

    def test_registry_independence_is_boolean_not_equal_integer(self):
        self.registry["approvals"][0]["review"]["independent_of_author"] = 1
        self.assertIn("registry recorded approval requires explicit non-author acceptance", self.errors())

    def disposition(self):
        self.write("docs/reviews/record-dispositions.json", {"records": [{
            "record": "probe", "current_stage": "example-passed", "disposition": "approved within declared scope",
            "next_concrete_action": "Retain the reviewed synthetic interpretation."}]})

    def test_metadata_helper_is_idempotent_and_preserves_accepted_semantics(self):
        self.disposition()
        path = self.root / "tools/vv/source_ledger.yml"
        original_hash = self.registry["approvals"][0]["fingerprint"]
        for flow in (True, False):
            for existing_review in (True, False):
                with self.subTest(flow=flow, existing_review=existing_review):
                    original = copy.deepcopy(self.ledger)
                    if not existing_review:
                        original["records"][0].pop("review")
                    path.write_text("# retained comment\n" + yaml.safe_dump(original, default_flow_style=flow, sort_keys=False, width=100000), encoding="utf-8")
                    self.assertTrue(metadata.reconcile(self.root, write=True))
                    self.assertFalse(metadata.reconcile(self.root, write=False))
                    after = gate.read(path)
                    for data in (original, after):
                        data["records"][0].pop("review", None)
                    self.assertEqual(original, after)
                    self.assertTrue(path.read_text().startswith("# retained comment\n"))
                    self.assertEqual(original_hash, gate.read(self.registry_path)["approvals"][0]["fingerprint"])
                    self.assertEqual([], gate.audit(self.root)[1])
                    self.assertEqual("Synthetic reviewer", gate.read(path)["records"][0]["review"]["by"])

    def test_metadata_helper_cannot_turn_self_review_into_independent_acceptance(self):
        self.disposition()
        self.row["review"]["independent_of_author"] = False
        self.registry["approvals"][0]["review"] = dict(self.row["review"])
        self.save()
        path = self.root / "tools/vv/source_ledger.yml"
        before = (path.read_bytes(), self.registry_path.read_bytes())
        with self.assertRaisesRegex(ValueError, "explicit non-author acceptance"):
            metadata.reconcile(self.root, write=True)
        self.assertEqual(before, (path.read_bytes(), self.registry_path.read_bytes()))

    def test_metadata_helper_preserves_explicit_withdrawal_and_reviewer_conflicts(self):
        self.disposition()
        original = dict(self.row["review"])
        for field, value in (("independent_of_author", False), ("decision", "hold"),
                             ("by", "Different reviewer"), ("date", "2026-11-20"),
                             ("independence_basis", "Withdrawn basis")):
            with self.subTest(field=field):
                self.row["review"] = dict(original, **{field: value})
                self.save()
                path = self.root / "tools/vv/source_ledger.yml"
                before = (path.read_bytes(), self.registry_path.read_bytes())
                with self.assertRaisesRegex(ValueError, "explicit ledger review conflicts"):
                    metadata.reconcile(self.root, write=True)
                self.assertEqual(before, (path.read_bytes(), self.registry_path.read_bytes()))

    def test_metadata_helper_rejects_invalid_registry_reviewer_before_writing(self):
        self.disposition()
        original = dict(self.registry["approvals"][0]["review"])
        self.row.pop("review")
        for field, value in (("by", ""), ("date", "bogus"), ("independent_of_author", "true")):
            with self.subTest(field=field):
                self.registry["approvals"][0]["review"] = dict(original, **{field: value})
                self.save()
                path = self.root / "tools/vv/source_ledger.yml"
                before = (path.read_bytes(), self.registry_path.read_bytes())
                with self.assertRaisesRegex(ValueError, "invalid accepted review"):
                    metadata.reconcile(self.root, write=True)
                self.assertEqual(before, (path.read_bytes(), self.registry_path.read_bytes()))

    def test_metadata_helper_refuses_stale_semantics_without_writing(self):
        self.disposition()
        self.master["argument_schemes"][0]["conclusions"] = ["¬q(X)"]
        self.save()
        ledger_path = self.root / "tools/vv/source_ledger.yml"
        before = (ledger_path.read_bytes(), self.registry_path.read_bytes())
        with self.assertRaisesRegex(ValueError, "stale semantic approval"):
            metadata.reconcile(self.root, write=True)
        self.assertEqual(before, (ledger_path.read_bytes(), self.registry_path.read_bytes()))


if __name__ == "__main__":
    unittest.main(verbosity=1)
