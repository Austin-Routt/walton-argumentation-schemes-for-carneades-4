# Integration of the independent baseline review

2026-11-21. Integrator: Claude, the implementer of PR20 and PR21, whose work
the review examined. This report records what was checked before the
reviewer's four commits were merged, the one gap found in them, and its fix.

## What was received

The owner supplied `Walton_Framework_Finalization_Review_2026-11-21.zip`
(115,440,579 bytes, SHA-256
`825164f8cf16c659ffbc390107f258a35ea6a56c75ad59d746a40e8160ce6ec9`). Its
2,618 manifest entries match the extracted files, with no extra file. Its
bundle's head, `da1dac5`, descends from PR21 `98a2335` through four commits:

| Commit | Subject |
|---|---|
| `598aa7e` | Repair import contracts, legacy conversion boundaries and Dung controls |
| `8a201e4` | Repair finite AATS projection and validate independent argument oracles |
| `3e2593d` | Finalize scoped source acceptance and framework baseline plan |
| `da1dac5` | Integrate independent approval attestations with strict Go ledger checks |

Its patches applied to `98a2335` reproduce the bundle's tree `41623f4`. The
four commits are merged unchanged.

## Findings reproduced before merging

Each finding about PR21 that the implementer could test was reproduced on
`98a2335` itself, not taken from the review:

* **Figure 5.** With the single penalty on P that Figure 5 prints, the
  hurried farmer chooses `row_dog`. PR21's statement that no figure/text
  difference changes a reported choice was wrong.
* **State-pair effects.** Two transitions between the same two states with
  different effect labels made the engine build an argument
  (`practical_reasoning_as2(a,a2,j2,q0,q1,p,u)`) that the adapter did not
  predict. The adapter now rejects inconsistent effects for one state pair.
* **The "second derivation".** The random-model check compared the engine
  with argument ids the adapter itself predicted, so it was not independent.
  The candidate derives the expected arguments from the facts alone and adds
  authored engine-level controls (`aats_adapter_test.go`).
* **Question diagnostics.** A legal joint action not named in the model was
  reported under CQ13; it belongs to CQ2. The injectivity requirement on π
  was not in the paper. Audience values outside the value universe were
  ignored silently; they are now errors.
* **Sources.** Entry 35 of the compendium is a conditional threat (with a
  capability premise), so the `ad_baculum` note's "which the compendium does
  not list as a scheme" was too strong; the Smart Car case's source
  description gave V1 the wrong support (E, not B); a migration row cited
  p. 94 for p. 93; and the D20 stamp of `sunk_costs` said the implementer
  wrote neither the record "nor its last change", though the implementer had
  made its reclassification from entry 25 to 24.
* **AS1 comparison.** An AS1 graph does not by itself decide between
  alternatives; that needs a model.

## The candidate rerun

At `da1dac5`, in a clean worktree: 10,428 of 10,428 registered checks pass
with `VV_STRICT=1`; go vet, the suite's unit tests and the engine and GoCHR
tests pass; the validator reports 0 errors and 10 warnings; `test_aats.py`
(71 tests), `aats.py check` (2 models, 17 situations, 32 decisions), the
approval checker and its 23 tests, the dialogue tests and the result
comparison tests pass.

Every example graph in this repository and in the vendored engine gives the
same outcome under the engines built from `98a2335` and from `da1dac5`, with
one intended exception: the old `abduction.yml`, whose `labels:` key the old
importer ignored, is refused by the strict importer; the candidate rewrites
it with `tests:`. The strict importer refuses a misspelt `assumption:` key
and a duplicated top-level key, as documented. This is a visible change for
users: a graph with an unknown or duplicated key that used to load, with the
key silently dropped, is now refused (`docs/import-profiles.md`).

## The gap: the web service

The candidate makes `check` and `eval` refuse CAF, AGXML and LKIF unless
`-experimental-import` is given (D26). `carneades server` did not: its upload
handler (`src/web/web.go`) called the four legacy importers directly, the
server had no such flag, and the upload form offered the formats. The plan
excludes web-service *hosting* from the claims, but the server is part of
the shipped binary, and this entry point contradicted the stated boundary.

Fix (engine **4.3.1+walton.7**): `src/web/profile.go` applies the same profile
before the upload is read; `carneades server -experimental-import` opts in;
the form labels the three formats experimental. `TestWebImportProfile`
covers both settings. Probed against a running server: a CAF upload is
refused with outcome `unsupported` by default, reaches the CAF converter with
the flag, and a YAML upload is evaluated as before.

## Two qualifications to the review's records

* **The eleven D20 approvals.** Their recorded basis is that the implementer
  "did not write the inherited executable records". Strictly, the
  implementer had changed them three times in Phases 5.5–5.9 (the
  assumption `credible(W)` of `expert_opinion`, the argument order of
  `in_domain` in `witness_testimony`, and the compendium citation in each
  source line). All three were in the tree, `49ad425`, that the non-author
  compendium reviews compared (`first-half-adjudication.json`, entries 2, 3
  and 26; `2026-11-21-second-half-source-review.md`, the other eight), so the
  non-author check exists, in those reviews. `d20-approvals.md` now says so.
  No approval changes.
* **Who the non-authors are.** The 149 recorded acceptances are by model
  reviewers (ChatGPT and Claude instances, each not the author of the change
  it accepts) and by the earlier closure reviews. The checker verifies that
  an attestation is present, explicit and matches the fingerprint; as D25
  says, it cannot authenticate the reviewer or prove the attestation true.
  Review by people remains part of the validation stage.

The review corrected the `ad_baculum` note in place; the superseded wording
is in `98a2335`.

## Review of this integration

The web-service fix, the version bump and the documentation changes are the
implementer's and have not been reviewed by a non-author (D14). They change
no library record, ledger row or fingerprint: the approval checker passes
unchanged. They belong in the second check of the engine before 1.0.0.

At the integrated commit, locally: 10,428 of 10,428 checks, all Go, engine
and Python tests, the adapter, approval, dialogue and comparison tests pass.
The repository's CI on Go 1.24.13 and Go 1.27.4, and its cross-version
comparison, are the evidence for the commit that is merged.
