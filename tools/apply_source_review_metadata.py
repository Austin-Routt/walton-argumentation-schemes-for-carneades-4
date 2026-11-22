#!/usr/bin/env python3
"""Copy accepted disposition metadata into the ledger; never approve semantics.

Default checks for the derived metadata without writing. --write updates only
ledger review fields and registry review snapshots. Existing approval hashes
must already match: semantic changes require independent adjudication first.
"""
import argparse
from datetime import date
import json
from pathlib import Path
import sys

import yaml

import check_source_approvals as gate

REVIEWER = "ChatGPT independent source-review team (integrated adjudication)"
OPEN_DISPOSITIONS = {"book form approved; own-original-source hold", "mapping-only; source interpretation unreviewed"}


def reconcile(repo, write=False):
    master, schemes, rows = gate.inventory(repo)
    ledger_path = repo / "tools/vv/source_ledger.yml"
    registry_path = repo / "docs/reviews/independent-approvals.json"
    registry = gate.read(registry_path)
    approvals = gate.index(registry.get("approvals"), "record", "approvals")
    if registry.get("schema") != gate.SCHEMA:
        raise ValueError("unexpected approval registry schema")
    dispositions = gate.index(gate.read(repo / "docs/reviews/record-dispositions.json").get("records"), "record", "dispositions")
    if set(dispositions) != set(rows):
        raise ValueError("dispositions must cover exactly the current ledger")
    reviewed = registry.get("review_date", "")
    if date.fromisoformat(reviewed).isoformat() != reviewed:
        raise ValueError("registry review_date must be YYYY-MM-DD")
    promoted = {sid for sid, row in rows.items() if row.get("stage") in gate.STAGES[2:]}
    accepted = {sid for sid, row in dispositions.items() if row.get("disposition") == "approved within declared scope"}
    if set(approvals) != promoted or accepted != promoted:
        raise ValueError("registry, promoted stages and accepted dispositions differ; adjudicate before applying metadata")
    for sid, approval in approvals.items():
        if approval.get("fingerprint") != gate.semantic_fingerprint(master, schemes[sid], rows[sid]):
            raise ValueError(f"{sid}: stale semantic approval; independently review the change before applying metadata")
    reviews = {}
    for sid, row in rows.items():
        disposition = dispositions[sid]
        if disposition.get("current_stage") != row.get("stage"):
            raise ValueError(f"{sid}: disposition stage differs from ledger")
        reference = f"docs/reviews/record-dispositions.json: {sid}"
        if sid in approvals:
            decision = "approved"
            evidence = approvals[sid].get("evidence")
            if not isinstance(evidence, list) or not evidence or not all(isinstance(x, str) and x.strip() for x in evidence):
                raise ValueError(f"{sid}: approval needs existing evidence references")
            note = "Approved only within the declared executable interpretation. Evidence: " + "; ".join(evidence + [reference])
            # Reconciliation is not a new review. Preserve the accepted reviewer
            # and non-author attestation instead of manufacturing an attribution.
            accepted_review = approvals[sid].get("review", {})
            invalid = gate.review_errors(accepted_review)
            if invalid:
                raise ValueError(f"{sid}: invalid accepted review: {'; '.join(invalid)}")
            current_review = row.get("review")
            if current_review is not None:
                if not isinstance(current_review, dict):
                    raise ValueError(f"{sid}: explicit ledger review is malformed; adjudicate before reconciliation")
                for key in gate.REVIEW_FIELDS:
                    if key in current_review and current_review[key] != accepted_review.get(key):
                        raise ValueError(f"{sid}: explicit ledger review conflicts on {key}; adjudicate before reconciliation")
        else:
            decision = "open" if disposition.get("disposition") in OPEN_DISPOSITIONS else "hold"
            action = disposition.get("next_concrete_action")
            if not isinstance(action, str) or not action.strip():
                raise ValueError(f"{sid}: held disposition needs a concrete reason/action")
            note = action + " Evidence: " + reference
        reviews[sid] = {"by": REVIEWER, "date": reviewed, "decision": decision, "note": note}
        if sid in approvals:
            reviews[sid].update({key: accepted_review.get(key) for key in gate.REVIEW_FIELDS})
    original = ledger_path.read_text(encoding="utf-8")
    document = yaml.compose(original)
    record_nodes = next(value.value for key, value in document.value if key.value == "records")
    replacements = []
    for node in record_nodes:
        fields = {key.value: value for key, value in node.value}
        sid = fields["record"].value
        text = yaml.safe_dump(reviews[sid], default_flow_style=True, sort_keys=False, allow_unicode=True, width=100000).strip()
        if "review" in fields:
            value = fields["review"]
            start, end = value.start_mark.index, value.end_mark.index
            if isinstance(value, yaml.MappingNode) and not value.flow_style:
                text += "\n" + " " * value.end_mark.column
        elif node.flow_style:
            start = end = node.end_mark.index - 1
            text = ", review: " + text
        else:
            start = end = original.index("\n", node.start_mark.index) + 1
            text = " " * node.start_mark.column + "review: " + text + "\n"
        replacements.append((start, end, text))
    updated = original
    for start, end, text in sorted(replacements, reverse=True):
        updated = updated[:start] + text + updated[end:]
    # Prove this operation changes only review metadata, not source contracts.
    before, after = gate.read(ledger_path), yaml.safe_load(updated)
    for data in (before, after):
        for row in data["records"]:
            row.pop("review", None)
    if before != after:
        raise ValueError("metadata update unexpectedly changed ledger semantics")
    registry_text = json.dumps(registry, ensure_ascii=False, indent=2) + "\n"
    changed = updated != original or registry_text != registry_path.read_text(encoding="utf-8")
    if write:
        ledger_path.write_text(updated, encoding="utf-8")
        registry_path.write_text(registry_text, encoding="utf-8")
    return changed


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--repo", type=Path, default=Path(__file__).resolve().parent.parent)
    parser.add_argument("--write", action="store_true")
    args = parser.parse_args()
    try:
        changed = reconcile(args.repo.resolve(), args.write)
    except (OSError, ValueError, TypeError, KeyError, AttributeError, yaml.YAMLError) as exc:
        print(f"review metadata: {exc}", file=sys.stderr)
        return 2
    print("Review metadata " + ("updated." if args.write and changed else "needs reconciliation." if changed else "is current."))
    return int(changed and not args.write)


if __name__ == "__main__":
    sys.exit(main())
