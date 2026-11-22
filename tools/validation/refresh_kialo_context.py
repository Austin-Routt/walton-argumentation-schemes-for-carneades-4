#!/usr/bin/env python3
"""Refresh only the supplied Kialo machine-draft library context.

Inputs are explicit and immutable Git objects are used, never unstaged repository
files. Historical annotations and graph bytes are never changed. This program
does not run an engine, create semantic gold labels, authenticate reviewers, or
recheck source books. Requires Python 3 and PyYAML.
"""
from __future__ import annotations

import argparse
import copy
import hashlib
import json
from pathlib import Path
import re
import subprocess
import sys

import yaml

LIBRARY = "walton.yml"
LEDGER = "tools/vv/source_ledger.yml"
REGISTRY = "docs/reviews/independent-approvals.json"
VERSION = "third_party/carneades-4/src/common/globals.go"
REVIEW_KEYS = ("by", "date", "decision", "independent_of_author", "independence_basis")
SOURCE_FILES = ("pilot_annotations.jsonl", "library_candidate_catalog.json",
                "pilot_manifest.json", "independent_review_blank.jsonl")
SIX_IDS = {"kialo2629-u01", "kialo2629-u04", "kialo2629-u07",
           "kialo2629-u09", "kialo2629-u11", "kialo2629-u12"}


def require(ok, message):
    if not ok:
        raise ValueError(message)


def sha(data):
    return hashlib.sha256(data).hexdigest()


def json_bytes(data):
    return (json.dumps(data, ensure_ascii=False, sort_keys=True, indent=2,
                       allow_nan=False) + "\n").encode("utf-8")


def lines_bytes(data):
    return b"".join((json.dumps(r, ensure_ascii=False, separators=(",", ":"),
                               allow_nan=False) + "\n").encode("utf-8") for r in data)


def unique_pairs(pairs):
    result = {}
    for key, value in pairs:
        require(key not in result, f"duplicate JSON key: {key}")
        result[key] = value
    return result


def jload(data):
    return json.loads(data, object_pairs_hook=unique_pairs)


class UniqueLoader(yaml.SafeLoader):
    pass


def unique_mapping(loader, node, deep=False):
    result = {}
    for key_node, value_node in node.value:
        key = loader.construct_object(key_node, deep=deep)
        require(key not in result, f"duplicate YAML key: {key}")
        result[key] = loader.construct_object(value_node, deep=deep)
    return result


UniqueLoader.add_constructor(yaml.resolver.BaseResolver.DEFAULT_MAPPING_TAG, unique_mapping)


def index(rows, key):
    require(isinstance(rows, list), f"expected list for {key} inventory")
    result = {}
    for row in rows:
        require(isinstance(row, dict) and isinstance(row.get(key), str), f"invalid {key} row")
        require(row[key] not in result, f"duplicate {key}: {row[key]}")
        result[row[key]] = row
    return result


def git(repo, *args):
    return subprocess.check_output(["git", "-C", str(repo), *args], stderr=subprocess.PIPE,
                                   timeout=20)


def snapshot(repo, ref, expected_commit=None):
    commit = git(repo, "rev-parse", "--verify", "--end-of-options", ref + "^{commit}").decode().strip()
    require(re.fullmatch(r"[0-9a-f]{40}", commit), "expected full Git commit")
    if expected_commit is not None:
        require(commit == expected_commit, "repository commit differs from --expect-commit")
    files = {p: git(repo, "show", f"{commit}:{p}") for p in (LIBRARY, LEDGER, REGISTRY, VERSION)}
    library = yaml.load(files[LIBRARY], Loader=UniqueLoader)
    ledger = yaml.load(files[LEDGER], Loader=UniqueLoader)
    registry = jload(files[REGISTRY])
    versions = re.findall(rb'^const Version = "([^"\n]+)"\s*$', files[VERSION], re.M)
    require(len(versions) == 1, "engine version declaration missing or ambiguous")
    result = {
        "binding": {"repository_commit": commit,
                    "repository_tree": git(repo, "rev-parse", f"{commit}^{{tree}}").decode().strip(),
                    "inputs_read_from_immutable_git_objects": True,
                    "library_version": library.get("meta", {}).get("version"),
                    "engine_version": versions[0].decode(),
                    "file_sha256": {p: sha(data) for p, data in files.items()}},
        "library": library,
        "ledger": ledger,
        "registry": registry,
    }
    validate_snapshot(result)
    return result


def validate_snapshot(snap):
    schemes = index(snap["library"].get("argument_schemes"), "id")
    rows = index(snap["ledger"].get("records"), "record")
    require(snap["registry"].get("schema") == "walton-source-approvals/v2", "unexpected approval registry schema")
    approvals = index(snap["registry"].get("approvals"), "record")
    require(len(schemes) == 149, "this bounded refresh expects 149 library records; review protocol if inventory changes")
    require(set(schemes) == set(rows) == set(approvals), "library, live ledger and approval registry inventories differ")
    for sid in schemes:
        row, approval = rows[sid], approvals[sid]
        require(row.get("stage") == "example-passed", f"{sid}: current source stage is not example-passed")
        stamp = {k: row.get("review", {}).get(k) for k in REVIEW_KEYS}
        require(stamp["decision"] == "approved" and stamp["independent_of_author"] is True,
                f"{sid}: missing recorded approval/independence attestation")
        require(stamp == approval.get("review"), f"{sid}: ledger/registry review stamp mismatch")
        require(re.fullmatch(r"sha256:[0-9a-f]{64}", approval.get("fingerprint", "")),
                f"{sid}: missing approval fingerprint")
    require(bool(snap["binding"]["library_version"]), "library version absent")


def historical(source, graph_path):
    files = {name: (source / name).read_bytes() for name in SOURCE_FILES}
    graph_bytes = graph_path.read_bytes()
    records = [jload(line) for line in files["pilot_annotations.jsonl"].splitlines() if line]
    catalog = jload(files["library_candidate_catalog.json"])
    manifest = jload(files["pilot_manifest.json"])
    blanks = [jload(line) for line in files["independent_review_blank.jsonl"].splitlines() if line]
    require(manifest["reproduction_input"]["kind"] == "normalized_graph", "only the supplied normalized graph profile is supported")
    require(sha(graph_bytes) == manifest["reproduction_input"]["sha256"], "historical graph hash mismatch")
    require(manifest["human_adjudicated_units"] == 0 and manifest["held_out_evaluation_allowed"] is False,
            "historical manifest is no longer an unadjudicated development draft")
    require(len(records) == len(blanks) == 12, "expected the supplied twelve draft/source-only units")
    require(len({r["dossier_id"] for r in records}) == 12, "duplicate historical dossier")
    require(len({r["source"]["claim_id"] for r in records}) == 11, "expected eleven historical child claims")
    require(len({r["source"]["location_id"] for r in records}) == 12, "expected twelve historical occurrences")
    cat_rows = index([dict(row, id=row["scheme"]["id"]) for row in catalog["records"]], "id")
    referenced = {sid for r in records for a in r["arguments"]
                  for sid in a["candidate_library_scheme_ids"] + a["considered_but_not_established_scheme_ids"]}
    require(set(cat_rows) == referenced and len(referenced) == 14, "historical candidate inventory mismatch")
    graph = jload(graph_bytes)
    for r in records:
        contract = r["annotation_contract"]
        require(contract["library_commit"] == catalog["library_commit"], "historical annotation/catalog commit mismatch")
        require(contract["library_sha256"] == catalog["walton_sha256"], "historical annotation/catalog library hash mismatch")
        require(contract["library_register_sha256"] == catalog["register_sha256"], "historical annotation/catalog register hash mismatch")
        validate_source_record(r, graph)
    validate_blanks(blanks, records)
    return {"records": records, "catalog": catalog, "manifest": manifest, "blanks": blanks,
            "graph": graph, "input_sha256": {**{k: sha(v) for k, v in files.items()}, "graph.json": sha(graph_bytes)},
            "blank_bytes": files["independent_review_blank.jsonl"]}


def validate_source_record(r, graph):
    require(r["status"] == "machine_proposed_unadjudicated", "machine draft misrepresented as gold")
    source = r["source"]
    require(source["split"] == "calibration_only_exposed_to_model", "exposed unit assigned outside calibration/development")
    view = graph["views"]["current_thesis_directed"]
    require(source["location_id"] in view["relation_location_ids"], "occurrence is not in the supplied active thesis view")
    loc = graph["locations"][source["location_id"]]
    require((source["claim_id"], source["parent_claim_id"], source["platform_relation"]) ==
            (loc["child_claim_id"], loc["parent_claim_id"], loc["relation"]), "source occurrence/target/relation mismatch")
    require(source["export_navigation"] == {"from_claim_id": loc["parent_claim_id"], "to_claim_id": loc["child_claim_id"]},
            "export navigation mismatch")
    require(source["candidate_reasoning_direction"] == {"reason_claim_id": loc["child_claim_id"], "target_claim_id": loc["parent_claim_id"]},
            "candidate direction mismatch")
    for c in r["source_claims"]:
        actual = graph["claims"][c["claim_id"]]
        require(c["text_exact"] == actual["text"] and c["claim_version"] == actual.get("version"), "source text/version changed")
        require(c["text_sha256_utf8"] == sha(actual["text"].encode()), "source text hash mismatch")
        require(c["span"] == {"unit": "unicode_codepoint", "start": 0, "end": len(actual["text"])}, "source span mismatch")
    ref = r["reference_process"]
    require(not ref["independent_annotation_ids"] and ref["adjudicator_id"] is None and ref["reference_version"] is None,
            "draft reference process claims an annotation or adjudication")
    require(ref["frozen_at"] is None and not r["evaluation_scenarios"], "draft has been promoted to a frozen/executable reference")
    require(all(a["status"] == "candidate_analysis_not_executable_argument" for a in r["arguments"]), "candidate argument promoted")


def validate_blanks(blanks, records):
    require(len(blanks) == len(records), "source-only packet length mismatch")
    for b, r in zip(blanks, records):
        require(b["dossier_id"] == r["dossier_id"], "source-only dossier mismatch")
        for field in ("source", "source_claims", "context_edges"):
            require(b[field] == r[field], f"source-only {field} mismatch")
        require(b["review_status"] == "unannotated" and not any(b["annotations"].values()), "source-only packet contains labels")
        require(not any(k in b for k in ("title", "arguments", "annotation_lesson", "candidate_library_contracts")),
                "source-only packet leaks machine proposals")
        require(b["annotator_local_id"] is None and b["annotation_timestamp"] is None, "source-only packet claims human annotation")


LESSON_REVISIONS = {
    "kialo2629-u03": ("Preserve the library's original-source hold.",
        "The library record is accepted within its current declared source scope; this corpus interpretation remains unadjudicated."),
    "kialo2629-u09": ("This library record remains held on original-source comparison.",
        "The library record is accepted within its current declared source scope; this corpus interpretation remains unadjudicated."),
}


def used_signatures(schemes):
    """Find statically named terms, including nested terms, in this small profile.

    The historical catalog matched a predicate name without its arity and thus
    included an unused relevant_differences/2 gloss. Distinguish arities here.
    This is a catalog selector, not a general Carneades term parser.
    """
    used = set()
    for scheme in schemes:
        for role in ("premises", "assumptions", "exceptions", "conclusions", "guards", "deletions"):
            for expression in scheme.get(role, []):
                for match in re.finditer(r"([a-z][a-z_0-9]*)\(", expression):
                    depth, commas, quote, escaped = 0, 0, None, False
                    end = None
                    for i in range(match.end(), len(expression)):
                        char = expression[i]
                        if quote:
                            if escaped:
                                escaped = False
                            elif char == "\\":
                                escaped = True
                            elif char == quote:
                                quote = None
                        elif char in "\"'":
                            quote = char
                        elif char in "([{":
                            depth += 1
                        elif char in ")]}":
                            if depth == 0:
                                end = i
                                break
                            depth -= 1
                        elif char == "," and depth == 0:
                            commas += 1
                    require(end is not None, "unterminated candidate term during catalog selection")
                    arity = commas + 1 if expression[match.end():end].strip() else 0
                    used.add(f"{match.group(1)}/{arity}")
    return used


def render(hist, snap):
    validate_snapshot(snap)
    schemes = index(snap["library"]["argument_schemes"], "id")
    rows = index(snap["ledger"]["records"], "record")
    approvals = index(snap["registry"]["approvals"], "record")
    old = {r["scheme"]["id"]: r for r in hist["catalog"]["records"]}
    for sid, entry in old.items():
        old_formula = {k: v for k, v in entry["scheme"].items() if k != "meta"}
        new_formula = {k: v for k, v in schemes[sid].items() if k != "meta"}
        require(old_formula == new_formula, f"{sid}: candidate executable contract changed; manual interpretation review required")
    used = used_signatures([schemes[sid] for sid in old])
    old_glosses = hist["catalog"]["language_glosses"]
    relevant_glosses = {k: v for k, v in old_glosses.items() if k in used}
    for key, value in relevant_glosses.items():
        require(snap["library"]["language"].get(key) == value,
                f"{key}: candidate gloss changed; manual interpretation review required")
    require(all(key in snap["library"]["language"] for key in used), "current candidate term lacks a language gloss")
    binding = snap["binding"]
    changed = [{"scheme_id": sid, "old_stage": old[sid]["stage"], "current_stage": rows[sid]["stage"],
                "old_decision": old[sid]["decision"], "current_decision": rows[sid]["review"]["decision"]}
               for sid in sorted(old)
               if (old[sid]["stage"], old[sid]["decision"]) != (rows[sid]["stage"], rows[sid]["review"]["decision"])]
    catalog = {"schema": "current-kialo-candidate-context/v1", "binding": binding,
               "notice": "Current source approval is not approval of any Kialo annotation. Fingerprints are copied attestations, not newly rechecked source semantics or authenticated identities.",
               "language_glosses": {k: snap["library"]["language"][k] for k in sorted(used)},
               "issue_schemes": snap["library"].get("issue_schemes", {}),
               "records": [{"scheme": schemes[sid], "stage": rows[sid]["stage"],
                            "decision": rows[sid]["review"]["decision"],
                            "source_contract": {k: v for k, v in rows[sid].items() if k != "review"},
                            "review_attestation": approvals[sid],
                            "question_mappings": rows[sid].get("questions", [])} for sid in sorted(old)]}
    records = copy.deepcopy(hist["records"])
    for r in records:
        r["historical_annotation_contract"] = copy.deepcopy(r["annotation_contract"])
        r["annotation_contract"].update({
            "library_commit": binding["repository_commit"], "library_version": binding["library_version"],
            "library_sha256": binding["file_sha256"][LIBRARY],
            "library_register_sha256": binding["file_sha256"][REGISTRY],
            "library_register_path": REGISTRY, "source_ledger_sha256": binding["file_sha256"][LEDGER],
            "engine_semantics_version": binding["engine_version"], "engine_executed_by_annotation": False,
            "context_revision": "metadata-only-current-library/v1", "frozen_at": None})
        for c in r["candidate_library_contracts"]:
            sid = c["scheme_id"]
            for out_role, lib_role in (("premises", "premises"), ("defaults", "assumptions"),
                                      ("exceptions", "exceptions"), ("conclusions", "conclusions")):
                c[out_role] = schemes[sid].get(lib_role, [])
            c["library_stage"] = rows[sid]["stage"]
            c["library_review_decision"] = rows[sid]["review"]["decision"]
            c["library_source_scope"] = "Accepted within the current ledger's declared source scope; corpus application remains unadjudicated."
        for q in r["critical_questions"]:
            if q["question_kind"] == "library_slot":
                q["source_reference"] = f"{binding['repository_commit']}:{LIBRARY} and {LEDGER}; retained draft paraphrase of a relevant slot, not a new numbered book CQ"
        if r["dossier_id"] in LESSON_REVISIONS:
            before, after = LESSON_REVISIONS[r["dossier_id"]]
            require(before in r["annotation_lesson"], "historical lesson differs from the reviewed metadata-only correction")
            r["annotation_lesson"] = r["annotation_lesson"].replace(before, after)
        r["context_refresh"] = {
            "kind": "library_metadata_only", "human_labels_created": 0,
            "semantic_proposals_reassessed": False, "historical_draft_file_sha256": hist["input_sha256"]["pilot_annotations.jsonl"],
            "notice": "Existing machine proposals and source text retained; current source acceptance does not approve these annotations."}
    outputs = {"pilot_annotations.jsonl": lines_bytes(records),
               "library_candidate_catalog.json": json_bytes(catalog),
               "independent_review_blank.jsonl": hist["blank_bytes"],
               "six_unit_source_only_training.jsonl": lines_bytes([b for b in hist["blanks"] if b["dossier_id"] in SIX_IDS])}
    manifest = {"schema": "kialo-context-refresh-manifest/v1", "binding": binding,
                "tool_sha256": sha(Path(__file__).read_bytes()),
                "unit_count": 12, "distinct_reason_claims": 11, "discussion_families": 1,
                "candidate_record_count": len(old), "repository_record_count": len(schemes),
                "human_adjudicated_units": 0, "scheme_assignment_is_gold": False,
                "held_out_evaluation_allowed": False, "annotation_semantics_validated": False,
                "engine_executed": False, "source_approval_fingerprints_recomputed": False,
                "source_input_sha256": hist["input_sha256"], "historical_library_commit": hist["catalog"]["library_commit"],
                "status_changes": changed,
                "formula_comparison": "All fourteen candidate record fields except meta unchanged; recorded historical glosses used at the same arity are unchanged.",
                "unused_historical_gloss_keys_excluded": sorted(set(old_glosses) - used),
                "training_packet": {"units": 6, "distinct_child_claims": 5, "discussion_families": 1,
                                    "ids": sorted(SIX_IDS), "split": "exposed_development_training_only"},
                "output_sha256": {name: sha(data) for name, data in outputs.items()},
                "limitations": ["This is provenance and context preparation, not semantic, engine or human validation.",
                                "Approval fingerprints and reviewer declarations are copied; use the repository's source-approval gate independently.",
                                "The normalized derivative is checked; the original canonical export is not reread here.",
                                "All supplied occurrences and duplicate families remain development-exposed."]}
    outputs["refresh_manifest.json"] = json_bytes(manifest)
    validate_bundle(outputs, hist, snap)
    return outputs


def validate_bundle(bundle, hist, snap):
    manifest = jload(bundle["refresh_manifest.json"])
    require(manifest["binding"] == snap["binding"], "stale commit/version/input binding in refreshed manifest")
    require(manifest["source_input_sha256"] == hist["input_sha256"], "historical source input binding mismatch")
    require(manifest["human_adjudicated_units"] == 0 and manifest["held_out_evaluation_allowed"] is False,
            "refreshed manifest falsely promotes human gold or test eligibility")
    records = [jload(line) for line in bundle["pilot_annotations.jsonl"].splitlines() if line]
    require(len(records) == 12, "refreshed unit count mismatch")
    for r, old in zip(records, hist["records"]):
        validate_source_record(r, hist["graph"])
        require(r["dossier_id"] == old["dossier_id"], "refreshed dossier order/identity changed")
        for field in ("source", "source_claims", "context_edges", "propositions", "arguments", "reference_process"):
            require(r[field] == old[field], f"refresh changed source or semantic/human field: {field}")
        c = r["annotation_contract"]
        require(c["library_commit"] == snap["binding"]["repository_commit"] and
                c["engine_semantics_version"] == snap["binding"]["engine_version"] and
                c["library_sha256"] == snap["binding"]["file_sha256"][LIBRARY], "stale annotation commit/version/library binding")
    blanks = [jload(line) for line in bundle["independent_review_blank.jsonl"].splitlines() if line]
    validate_blanks(blanks, records)
    require(bundle["independent_review_blank.jsonl"] == hist["blank_bytes"], "source-only packet bytes changed")
    for name, digest in manifest["output_sha256"].items():
        require(name in bundle and sha(bundle[name]) == digest, f"refreshed artifact hash mismatch: {name}")


def self_test(hist, snap, outputs):
    results = []

    def rejected(name, action):
        try:
            action()
        except ValueError as exc:
            results.append({"case": name, "rejected": True, "reason": str(exc)})
        else:
            raise AssertionError(f"negative control escaped: {name}")

    changed = copy.deepcopy(snap)
    changed["ledger"]["records"][0]["review"]["by"] += " altered"
    rejected("ledger_registry_review_stamp_disagreement", lambda: validate_snapshot(changed))
    changed = copy.deepcopy(snap)
    changed["registry"]["approvals"].pop()
    rejected("missing_current_approval_record", lambda: validate_snapshot(changed))
    changed = copy.deepcopy(snap)
    sid = hist["catalog"]["records"][0]["scheme"]["id"]
    row = next(r for r in changed["library"]["argument_schemes"] if r["id"] == sid)
    row["premises"] = row.get("premises", []) + ["invented_new_requirement(X)"]
    rejected("changed_candidate_formula_needs_interpretation_review", lambda: render(hist, changed))
    changed = copy.deepcopy(snap)
    used = used_signatures([r["scheme"] for r in hist["catalog"]["records"]])
    key = next(k for k in hist["catalog"]["language_glosses"] if k in used)
    changed["library"]["language"][key] += " changed meaning"
    rejected("changed_candidate_gloss_needs_interpretation_review", lambda: render(hist, changed))

    def output_mutation(file, mutate, jsonl=False):
        altered = dict(outputs)
        if jsonl:
            data = [jload(line) for line in altered[file].splitlines() if line]
            mutate(data)
            altered[file] = lines_bytes(data)
        else:
            data = jload(altered[file]); mutate(data); altered[file] = json_bytes(data)
        validate_bundle(altered, hist, snap)

    rejected("stale_commit_binding", lambda: output_mutation("refresh_manifest.json", lambda d: d["binding"].update(repository_commit=hist["catalog"]["library_commit"])))
    rejected("stale_engine_version", lambda: output_mutation("pilot_annotations.jsonl", lambda d: d[0]["annotation_contract"].update(engine_semantics_version="4.3.1+walton.5"), True))
    rejected("machine_draft_promoted_to_human_gold", lambda: output_mutation("pilot_annotations.jsonl", lambda d: d[0].update(status="human_gold"), True))
    rejected("exposed_occurrence_promoted_to_test", lambda: output_mutation("pilot_annotations.jsonl", lambda d: d[0]["source"].update(split="sealed_test"), True))
    rejected("changed_source_claim_text", lambda: output_mutation("pilot_annotations.jsonl", lambda d: d[0]["source_claims"][0].update(text_exact="changed"), True))
    rejected("source_only_packet_leaks_machine_title", lambda: output_mutation("independent_review_blank.jsonl", lambda d: d[0].update(title="machine answer"), True))
    rejected("catalog_contract_corruption", lambda: output_mutation("library_candidate_catalog.json", lambda d: d["records"][0]["scheme"].update(premises=["invented_new_requirement(X)"])))
    return {"schema": "kialo-context-refresh-controls/v1", "binding": snap["binding"],
            "positive_refresh_verified": True, "negative_controls": results,
            "human_gold": 0, "engine_executed": False, "independent_semantic_validation": False}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--repo", type=Path, required=True)
    parser.add_argument("--ref", default="HEAD", help="Commit-ish read as immutable Git objects; default HEAD")
    parser.add_argument("--expect-commit", help="Optional exact 40-character commit lock")
    parser.add_argument("--source-annotations", type=Path, required=True, help="Unchanged historical twelve-draft folder")
    parser.add_argument("--graph", type=Path, required=True, help="Supplied normalized graph.json")
    mode = parser.add_mutually_exclusive_group(required=True)
    mode.add_argument("--output", type=Path, help="New output directory; existing nonempty directories are refused")
    mode.add_argument("--check", type=Path, help="Check an existing refresh against the exact specified repository pin")
    parser.add_argument("--self-test", action="store_true", help="Run deliberate mismatch controls; no engine or corpus benchmark")
    args = parser.parse_args()
    snap = snapshot(args.repo, args.ref, args.expect_commit)
    hist = historical(args.source_annotations, args.graph)
    outputs = render(hist, snap)
    if args.check:
        actual = {name: (args.check / name).read_bytes() for name in outputs}
        validate_bundle(actual, hist, snap)
        require(actual == outputs, "refreshed context differs from deterministic reconstruction")
    else:
        require(not args.output.exists() or not any(args.output.iterdir()), "refusing to overwrite a nonempty output directory")
        require(args.output.resolve() != args.source_annotations.resolve(), "historical inputs cannot be the output directory")
        args.output.mkdir(parents=True, exist_ok=True)
        for name, data in outputs.items():
            (args.output / name).write_bytes(data)
    if args.self_test:
        controls = self_test(hist, snap, outputs)
        if args.output:
            (args.output / "mismatch_controls.json").write_bytes(json_bytes(controls))
    print(json.dumps({"status": "pass", "repository_commit": snap["binding"]["repository_commit"],
                      "library_version": snap["binding"]["library_version"], "engine_version": snap["binding"]["engine_version"],
                      "draft_units": 12, "candidate_records": 14, "current_records": 149,
                      "human_gold": 0, "engine_executed": False,
                      "negative_controls_rejected": len(controls["negative_controls"]) if args.self_test else 0}))


if __name__ == "__main__":
    try:
        main()
    except (ValueError, KeyError, OSError, subprocess.SubprocessError, yaml.YAMLError) as exc:
        print(f"REFUSED: {exc}", file=sys.stderr)
        sys.exit(2)
