#!/usr/bin/env python3
"""Count structural vocabulary sharing with an explicit, reproducible convention.

This is an inventory, not a proof of semantic interoperability. The root
predicate/arity of premises, conclusions, defaults and exceptions is counted;
polarity is ignored and all bare statement variables share one separate class.
Nested terms and unused language templates do not add root signatures.
"""
from __future__ import annotations

import argparse
from collections import defaultdict
import hashlib
import json
from pathlib import Path
import subprocess

import yaml
from validate_schemes import LIST_FIELDS, parse_statement


def inventory(raw: bytes) -> dict:
    data = yaml.safe_load(raw)
    owners = defaultdict(set)
    records = data["argument_schemes"]
    for scheme in records:
        for field in LIST_FIELDS:
            for atom in scheme.get(field, []) or []:
                _, name, arity, _ = parse_statement(atom)
                key = "@statement-variable" if name is None else f"{name}/{arity}"
                owners[key].add(scheme["id"])
    shared = {key: sorted(ids) for key, ids in owners.items() if len(ids) > 1}
    connected = set().union(*(set(ids) for ids in shared.values())) if shared else set()
    return {
        "walton_sha256": hashlib.sha256(raw).hexdigest(),
        "records": len(records),
        "predicate_signatures_without_statement_variable": len(owners) - ("@statement-variable" in owners),
        "records_sharing_predicate_or_statement_variable": len(connected),
        "bare_variable_records": sorted(owners.get("@statement-variable", [])),
        "shared_signatures": dict(sorted(shared.items())),
        "caveat": "Structural sharing only; no semantic, role, burden, or source-fidelity claim.",
    }


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--repo", type=Path, default=Path(__file__).resolve().parents[1])
    parser.add_argument("--ref", help="Optional Git revision; otherwise read the working walton.yml")
    args = parser.parse_args()
    raw = subprocess.check_output(["git", "show", f"{args.ref}:walton.yml"], cwd=args.repo) if args.ref else (args.repo / "walton.yml").read_bytes()
    result = inventory(raw)
    result["revision"] = args.ref or "working-file"
    print(json.dumps(result, indent=2, sort_keys=True))


if __name__ == "__main__":
    main()
