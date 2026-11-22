#!/usr/bin/env python3
"""Check recorded source approvals, or print candidate semantic fingerprints.

Default: every promoted row needs a matching approval; every recorded approval
must match its current record and have evidence and existing example IDs.
--require-stable additionally requires every record in walton.yml to be at
example-passed. --fingerprints only inventories hashes; it grants no approval.

The hash covers the raw scheme, statically used language templates, global
issue schemes, and selected source-ledger contract fields. Ledger stages and
notes are excluded. Dynamic templates supplied through statement variables,
engine code, test results and reviewer identity are not authenticated here.
An explicit non-author attestation is required; an integrity-valid self-review
is not independent acceptance. The attestation is evidence, not authentication.
"""
import argparse
from datetime import date
import hashlib
import json
from pathlib import Path
import sys
import yaml

sys.path.insert(0, str(Path(__file__).resolve().parent))
from validate_schemes import LIST_FIELDS, load_unique, parse_statement

SCHEMA = "walton-source-approvals/v2"
STAGES = ("mapped", "transcription-checked", "interpretation-approved", "questions-sorted", "example-passed")
CONTRACT_FIELDS = ("origin", "upstream", "entry", "cites", "claim", "source_questions")
REVIEW_FIELDS = ("by", "date", "decision", "independent_of_author", "independence_basis")


def read(path):
    data, duplicates = load_unique(path.read_text(encoding="utf-8"))
    if duplicates or not isinstance(data, dict):
        raise ValueError(f"{path}: expected a mapping without duplicate keys")
    return data


def index(items, key, label):
    if not isinstance(items, list):
        raise ValueError(f"{label}: expected a list")
    result = {}
    for item in items:
        name = item.get(key) if isinstance(item, dict) else None
        if not isinstance(name, str) or not name.strip() or name in result:
            raise ValueError(f"{label}: missing or duplicate {key} {name!r}")
        result[name] = item
    return result


def inventory(repo):
    master = read(repo / "walton.yml")
    schemes = index(master.get("argument_schemes"), "id", "schemes")
    rows = index(read(repo / "tools/vv/source_ledger.yml").get("records"), "record", "ledger")
    if not schemes or set(rows) != set(schemes):
        raise ValueError(f"ledger IDs differ from nonempty library: missing {sorted(set(schemes)-set(rows))}; unknown {sorted(set(rows)-set(schemes))}")
    return master, schemes, rows


def semantic_fingerprint(master, scheme, row):
    used, language = set(), master.get("language", {})
    if not isinstance(language, dict):
        raise ValueError("language must be a mapping")
    for field in LIST_FIELDS + ("deletions", "guards"):
        for statement in scheme.get(field, []) or []:
            _, name, arity, _ = parse_statement(statement)
            if name is not None:
                used.add(f"{name}/{arity}")
    if any(key not in language or not isinstance(language[key], str) for key in used):
        raise ValueError(f"{scheme['id']}: a referenced language template is missing or not text")
    contract = {key: row.get(key) for key in CONTRACT_FIELDS}
    if row.get("comparison") is not None:
        # the comparison contract an alternative question names (added after
        # the first approvals: only rows that name one carry the key, so the
        # fingerprints of the others are unchanged)
        contract["comparison"] = row["comparison"]
    contract["questions"] = [{k: v for k, v in q.items() if k != "note"} for q in row.get("questions", [])]
    payload = {"format": "walton-source-semantics/v1", "scheme": scheme,
               "language": {key: language[key] for key in sorted(used)},
               "issue_schemes": master.get("issue_schemes", {}), "source_contract": contract}
    encoded = json.dumps(payload, ensure_ascii=False, sort_keys=True, separators=(",", ":"), allow_nan=False).encode("utf-8")
    return "sha256:" + hashlib.sha256(encoded).hexdigest()


def review_errors(stamp):
    """Validate an attestation without treating its assertions as authenticated."""
    errors = []
    if not isinstance(stamp, dict):
        stamp = {}
    if stamp.get("decision") != "approved":
        errors.append("recorded approval requires review decision approved")
    if not isinstance(stamp.get("by"), str) or not stamp["by"].strip():
        errors.append("recorded approval requires a named reviewer")
    if stamp.get("independent_of_author") is not True:
        errors.append("recorded approval requires explicit non-author acceptance")
    if not isinstance(stamp.get("independence_basis"), str) or not stamp["independence_basis"].strip():
        errors.append("recorded approval requires an independence evidence basis")
    try:
        if date.fromisoformat(stamp.get("date", "")).isoformat() != stamp["date"]:
            raise ValueError("noncanonical date")
    except (TypeError, ValueError):
        errors.append("recorded approval requires review date YYYY-MM-DD")
    return errors


def audit(repo, require_stable=False):
    master, schemes, rows = inventory(repo)
    fingerprints = {sid: semantic_fingerprint(master, s, rows[sid]) for sid, s in schemes.items()}
    path = repo / "docs/reviews/independent-approvals.json"
    registry = read(path) if path.exists() else {}
    if registry and registry.get("schema") != SCHEMA:
        raise ValueError(f"{path}: schema must be {SCHEMA}")
    approvals = index(registry.get("approvals", []), "record", "approvals")
    examples = {c["name"] for p in (repo / "tools/vv/cases").glob("*.yml")
                for c in read(p).get("cases", []) if isinstance(c, dict) and isinstance(c.get("name"), str)}
    errors = []
    for sid, approval in approvals.items():
        if sid not in schemes:
            errors.append(f"{sid}: approval names an unknown record")
            continue
        if approval.get("fingerprint") != fingerprints[sid]:
            errors.append(f"{sid}: approval fingerprint is stale or missing")
        review = rows[sid].get("review")
        stamp = {key: review.get(key) for key in REVIEW_FIELDS} if isinstance(review, dict) else {}
        errors.extend(f"{sid}: {error}" for error in review_errors(stamp))
        errors.extend(f"{sid}: registry {error}" for error in review_errors(approval.get("review")))
        if approval.get("review") != stamp:
            errors.append(f"{sid}: approval review metadata differs from the ledger")
        for field in ("evidence", "examples"):
            values = approval.get(field)
            if not isinstance(values, list) or not values or not all(isinstance(v, str) and v.strip() for v in values):
                errors.append(f"{sid}: {field} must be a nonempty list of references")
            elif field == "examples" and set(values) - examples:
                errors.append(f"{sid}: unknown examples {sorted(set(values)-examples)}")
    for sid, row in rows.items():
        stage = row.get("stage")
        if stage not in STAGES:
            errors.append(f"{sid}: unknown stage {stage!r}")
        if stage in STAGES[2:] and sid not in approvals:
            errors.append(f"{sid}: {stage} has no independent approval")
        if isinstance(row.get("review"), dict) and row["review"].get("decision") == "approved" and stage not in STAGES[2:]:
            errors.append(f"{sid}: approving review is inconsistent with stage {stage!r}")
        if require_stable and stage != "example-passed":
            errors.append(f"{sid}: stable release requires example-passed (found {stage!r})")
    return fingerprints, errors


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("root", nargs="?", type=Path)
    parser.add_argument("--repo", type=Path)
    mode = parser.add_mutually_exclusive_group()
    mode.add_argument("--require-stable", action="store_true")
    mode.add_argument("--fingerprints", action="store_true")
    args = parser.parse_args(argv)
    if args.root and args.repo:
        parser.error("use a positional root or --repo, not both")
    repo = (args.repo or args.root or Path(__file__).resolve().parent.parent).resolve()
    try:
        if args.fingerprints:
            master, schemes, rows = inventory(repo)
            print(json.dumps({sid: semantic_fingerprint(master, s, rows[sid]) for sid, s in schemes.items()}, indent=2, sort_keys=True))
            return 0
        fingerprints, errors = audit(repo, args.require_stable)
    except (OSError, ValueError, TypeError, KeyError, AttributeError, yaml.YAMLError) as exc:
        print(f"source approvals: invalid input: {exc}", file=sys.stderr)
        return 2
    for error in errors:
        print(f"source approvals: {error}", file=sys.stderr)
    if not errors:
        print(f"Source approval integrity passed for {len(fingerprints)} records" + ("; stable source-stage gate passed." if args.require_stable else "."))
    return 1 if errors else 0


if __name__ == "__main__":
    sys.exit(main())
