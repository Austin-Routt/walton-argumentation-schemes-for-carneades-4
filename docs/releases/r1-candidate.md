# R1 — corrected preview

Library **0.2.0**; engine **4.3.1+walton.2**. This candidate supersedes the old R1 proposal at `8c092b6`. Publish only the commit recorded by the final clean acceptance evidence, not the old pin. This file does not assert that a remote tag or release already exists.

## Why a corrected candidate is needed

Independent clean runs reproduce baseline master's7,994 passing checks on both Go 1.24.13 and 1.27.4. They also reproduce the former R1's 7,164 checks from a separate full clone. Independent counterexamples nevertheless exposed cancellation, matching, logical fact reuse, issue-context, nested-parser and authored-exception defects. Source inspection found role and warrant defects that green reference examples had accepted. Those cases are now persistent regressions.

## Engine changes

* Stream candidate matching and check cancellation/work quotas during unsuccessful joins; do not allocate the full Cartesian product before applying a limit.
* Let one logical fact satisfy repeated premises without imposing hidden inequality. Ordinary GoCHR default semantics remain unchanged.
* Check empty/no-match cancellation, malformed nested terms and empty-theory assumptions.
* Resolve complete issue closure before seeding rules and after derived bridges. Reject inconsistent stipulations and incompatible issue structures; a defeated default cannot seed another accepted default.
* Stage mutable inference state and publish only on success. Failure is atomic; successful publication preserves existing node identities.
* Require every authored exception's direct linked witness under the same substitution, with actual registered graph nodes and the correct target. API clients still need structural validation.
* Document rule, term, matching, constraint, stored-node and issue quotas without claiming a process-memory or unrestricted-service guarantee.

## Library and provenance

All 114 book-mapped records were independently reviewed. Repairs cover generalization identity, complete model warrants, agent-scoped practical reasoning, sufficient-condition quantification, diagram/conclusion identity, typed character traits, event-scoped retroduction, explicit coextension, recaller-scoped memory and conditional exceptional-case backing. Five formerly deferred questions were modelled from the available compendium with their actual provenance. Migrations report cases requiring new evidence instead of fabricating it.

The supplied Atkinson and Bench-Capon paper was reviewed in full. Its17 questions have individual dispositions; eight still lack complete mappings in the inherited AS1 record. CQ10 is positive support, and AS2 requires a fuller transition/value model. A falsely attributed example was replaced with the paper's actual example.

**70/120 records** have explicit current source approvals:65 book-origin and five inherited. **50 remain held** (48 transcription-checked, two mapped). This is a usable preview with declared interpretations, not a complete stable 1.0.0 library. The source-approval integrity checker rejects stale/missing approvals; `--require-stable` deliberately refuses this incomplete content state.

## Verification

Both pinned Go versions are required CI jobs. The complete inventory contains 8,206 distinct checks; duplicate emitted identities, duplicate manifest entries and malformed comparator records fail. The final package includes clean-checkout engine/race/vet logs, strict results, exact per-check comparison, the expected-check inventory, source-approval mutation tests and the validator output. No known executable failure is waived. Informational validator warnings are explained in the review report; a warning count is not a quality score.

## Assets and migration

The source tree/bundle contains the patched engines. `carneades-4.zip` and `GoCHR.zip` remain byte-for-byte upstream baselines; downloading just those archives does not install these fixes. Keep `sources.lock` and the original assets for reproducible comparisons. Use the reviewed source snapshot or bundle to build the corrected engine.

Read `docs/using.md`, `docs/contracts.md`, `tools/migration_r2.yml` and the migration tool's manual-action output before updating a graph. Source repairs can add explicit witnesses or change participant roles; old graph facts cannot always be renamed automatically.

Remote publication requires authenticated repository write access. Do not present a prepared local tag, patch or command as a completed GitHub release. Upstream patches and the source-question memorandum are prepared for review and remain unsent pending explicit authorization to contact others.
