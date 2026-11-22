# Cross-review of the root-authored engine and input documentation

**Reviewer:** the `final_engine` agent, independently reviewing the root reviewer's documentation. The reviewer authored the engine/import repairs and therefore does not use this document as independent approval of their own code. Root's separate inspection and final clean acceptance are required for that purpose.

**Date:** 2026-11-22 UTC. **Baseline:** `98a2335d2c9909cf597ae85a9f359f20d092ae6c` plus the local closure candidate.

## Decision

**Accept the engine/input-profile claims in the four documents at the hashes below.** They accurately distinguish native YAML, the project-specific AIF projection, experimental legacy converters and attack-only TGF. They do not claim general AIFdb/corpus conformance, a complete modern ICCMA implementation, large-instance solver capacity, a published release, or empirical validity. Their assertions about final frozen execution evidence remain conditional on the root reviewer completing and supplying that evidence. I did not independently re-review source-book approval or the AATS mathematics in this documentation pass.

The AIF documents now disclose all three consequential conversion decisions: RA is generic linked with external scheme provenance; binary CA becomes symmetric incompatibility; unargued non-issue I nodes become stipulations. Those match the implementation. The legacy flag is correctly described as explicit experimental opt-in rather than certification or a validation bypass. The CAF defaults match the pinned official 1.3 schema; the AGXML segmentation boundary and LKIF exception/legacy-assumption qualifications match the current code. Native structural validation remains required after import.

## Concrete mismatch found and corrected

This pass discovered a real remaining implementation weakness behind the strict JSON admission claim. Go's JSON decoder matches struct fields case-insensitively. The initial duplicate-key guard compared exact decoded key strings, allowing a later `Nodes` array to overwrite `nodes`. A bounded probe reproduced a successful empty graph from a source containing a statement.

The guard now rejects keys equivalent under Unicode simple folding, with linear key-processing rather than pairwise scanning. Regressions cover root `nodes`/`Nodes`, nested `nodeID`/`NODEID`, and the Unicode alias `nodes`/`nodeſ`, in addition to exact and escaped duplicates. The counterexample now returns an invalid-input error. AIF and actual CLI subprocess tests passed after the correction; evidence is `aif-case-alias-before.jsonl`, `aif-case-alias-after.jsonl` and `aif-cli-final-go1.27.4.log`. Only `aif/profile.go` and `aif/profile_test.go` changed during this cross-review. The root reviewer was notified before and after the change.

No further material engine/input-documentation mismatch was found. Stable 1.0.0, full interchange fidelity and the broader validation effort remain explicitly separate decisions; the preview wording is appropriate.

## Accepted document hashes

| File | SHA-256 |
|---|---|
| `docs/import-profiles.md` | `ca68f03dcd284f2c282f2645699be0aa70b9c8da1ed2a01ed6fbebb6d141f7ab` |
| `docs/releases/framework-0.4.1.md` | `4a56600229641fb76f557e5a052bcd0bd89d2c19daa28ab0ec1b74db5e942415` |
| `PLAN.md` | `6dd6254eb24137934622cf3718af1460572c00eede5e2d96029c7c3f9f0b33be` |
| `README.md` | `d44012fb421947c2e4e44a2e12c030b71f33c209834f1b5f287b4d313d8e3c5b` |

## Original provenance inputs

I compared each file's bytes with the same path at public PR21 commit `98a2335d2c9909cf597ae85a9f359f20d092ae6c`, not merely its current Git status. All three are unchanged, and the two ZIP hashes match `sources.lock`.

| File | SHA-256 | Comparison with PR21 |
|---|---|---|
| `sources.lock` | `636751754f8ad5fd6412508d2670232c9c11a569f9d6602e96025eee49f4861e` | Identical |
| `carneades-4.zip` | `dbac157bb9505e3c48a376f7bc06418cced949144ff2dd6ed91f90fcff611d0b` | Identical |
| `GoCHR.zip` | `6f79056a58183807d907d3536820f48d0a0cfe9bc02233486f4b93f4b5d50851` | Identical |

The patched `third_party` sources are therefore separate from these unchanged provenance archives, exactly as the release and README text state.
