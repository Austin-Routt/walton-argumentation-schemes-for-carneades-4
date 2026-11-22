# Independent review of PR22 source and approval provenance

Review date: 21 November 2026, America/Chicago. Machine execution records use UTC.
Reviewer: ChatGPT `integration_provenance`, a separate non-author reviewer of the integration changes.

Inspected public commit: `d9a25857277140a5d3f10aaf1dfa4ade0f15b91f`.
Inspected tree: `fcea512fd99551b453ec283158ade395262d4b7c`.
Previously accepted candidate: `da1dac542ffd336b7e059f00162bd92a9f4e5d6a`.

## Decision

The substantive library and source-scope decisions have converged. PR22 does not change a library record, executable scheme formula, source-ledger contract or semantic fingerprint. The implementer's historical correction is supported: the three identified classes of earlier changes were already present at `49ad425`, the explicitly pinned baseline of the two cited non-author compendium reviews.

One narrow correction remains in the live approval data. The integration amended `d20-approvals.md`, but all eleven affected live ledger review notes still say the implementer wrote neither the record **nor its last change**. The matching registry still attributes independence generically to the implementer's D20 check. The actual independent evidence chain is more specific than that attribution. This is not a reason to reopen the eleven accepted source interpretations; it is a reason to correct their current attribution consistently before freezing/tagging the next baseline.

`recommended-attribution-updates.json` supplies the eleven proposed current non-author retained-acceptance stamps, record-specific scope, historical evidence, corrected notes and integration constraints. This review authorizes retaining those existing interpretations and examples after the evidence-chain audit. It does not purport to perform eleven new full primary-source comparisons.

The new server change is outside this review's code assignment. The root reviewer owns its implementation/runtime acceptance and CI verification. Tagging should follow that non-author review and the narrow metadata correction, rather than preceding them. Validation protocol design can begin immediately at an explicit commit; a remote tag is not technically required to define or pilot the study.

## What was checked

I inspected the complete change list from the accepted `da1dac5` to current `d9a2585`, the integration report, the current plan/release/import documents, the D20 report, the live source ledger, the approval registry and the current source-acceptance register. I traced the relevant changes through Git history and compared the named historical reports to their pinned source snapshot and the actual affected formulas. I ran only the source-stage integrity checker; no full runtime suite, server experiment, source acquisition or external contact was performed by this reviewer.

The following files are byte-identical between the accepted candidate and public PR22:

| File | Consequence |
|---|---|
| `walton.yml` | No library source/formula/language/issue change in the integration |
| `tools/vv/source_ledger.yml` | No source-contract or stage change; the attribution problem also remains here |
| `docs/reviews/independent-approvals.json` | No approval fingerprint/stamp change |
| `docs/reviews/final-source-acceptance.json` | Previous record-by-record recommendations retained unchanged |
| `sources.lock` | Original-source lock unchanged |
| `carneades-4.zip`, `GoCHR.zip` | Original provenance archives unchanged |

The checker returned:

```text
Source approval integrity passed for 149 records; stable source-stage gate passed.
```

This result establishes matching source-stage metadata and fingerprints. It does not establish the truth of an authorship claim embedded in those fields. The current finding is a concrete example of that limit rather than a failure of the checker's stated design.

Exact hashes, the command and output, and the clean starting state are in `provenance-checks.json`.

## The historical correction is supported

The three relevant classes of changes are visible in the history before `49ad425cbf7afeebcbd9a07ee479d60681ca33ea`:

| Commit | Author shown in history | Change |
|---|---|---|
| `481bc40` | Claude | Adds the compendium entry/page citation to each of the eleven records' source lines |
| `2bc28ee` | Claude | Adds `credible(W)` to `expert_opinion`'s assumptions |
| `e8c00ab` | Claude | Corrects `witness_testimony` from `in_domain(D,S)` to `in_domain(S,D)` |

All three commits are ancestors of `49ad425`. The first-half JSON explicitly records that reviewed commit for `expert_opinion`, `witness_testimony` and `ignorance`. Its saved baseline formulas contain the credibility default and corrected domain argument order. It discusses the credibility default using the compendium's own Carneades treatment on pp. 374–375. Thus those two executable modifications were present and actually addressed, not merely assumed from chronology.

The second-half source report explicitly identifies `49ad425` as its baseline and states that it examined premises, conclusions, critical-question dispositions, source-ledger assertions and relevant cases for the 51-record scope. Its individual sections include the remaining eight D20 records. Therefore the implementer's citation-change assertion is supported by the recorded review scope and Git source snapshot.

The complete before/after record-field changes are retained in `historical-implementer-changes.json`, and the current stamps and relevant first-half evidence are in `d20-attribution-evidence.json`.

## What those reviews accepted—and what they did not

The historical reviews were not eleven blanket, unqualified approvals of all sources. Their qualifications matter and should remain visible.

| Record or group | Retained source interpretation |
|---|---|
| `expert_opinion` | The six-question compendium/Carneades interpretation, including challengeable credibility; original-source status remains separately scoped |
| `witness_testimony` | Domain-based access and sincerity as declared interpretations; the later account-specific bias repair has separate acceptance |
| `ignorance` | Explicit would-be-known and not-known premises; burden/probandum allocation is not supplied by a proof-standard threshold alone |
| `correlation_to_cause` | An explanatory-theory default is a selected non-coincidence policy, not empirical causal proof |
| `positive_consequences`, `negative_consequences` | Action/outcome-specific likelihood adequacy; no universal fixed probability requirement |
| `ethotic1` | Defeasible support for the proposition is the inherited evidential policy, distinct from a literal graded-plausibility or numeric model |
| `ethotic2` | The current reduced-plausibility repair supersedes the previously blocked contradictory-proposition conclusion |
| The two inherited slippery-slope records | Components of a construction, distinct from the full compendium normative form supplied separately |
| `established_rule` | Inherited rule/conclusion application with scoped exceptions, distinct from the separate actor/action form |

In particular, the original second-half report **blocked** the old `ethotic2` conclusion. It must not be cited as though it approved today's result unchanged. The later `closure-supplemental-independent-acceptance.md` explicitly accepts the reduced-plausibility correction and the `witness_testimony` account-scoped bias correction. Those later links belong in the current evidence chain.

D20 changes which original-source comparisons are required for the owner's chosen product scope. It does not retroactively turn an unexamined book into an examined one or erase the declared interpretation limits. Where later supplied originals have since been compared, the current ledger preserves that evidence. This review neither retracts those comparisons nor makes new claims about originals it has not inspected.

## The live attribution correction is incomplete

The current `d20-approvals.md` retains the opening assertion that the implementer did not write the inherited executable records and adds a parenthetical qualification explaining his three earlier changes. The parenthetical is a useful correction, but replacing the opening assertion with a direct description of the evidence chain would be clearer.

More importantly, all eleven current ledger review notes still include this statement:

```text
Reviewer: the implementer, who did not write this record (Carneades did) nor its last change.
```

The reviewer's `by` value remains Claude, and the common `independence_basis` describes his D20 comparison of inherited records and changes made by other reviewers without identifying the separate historical checks of his own changes. The same basis appears in the matching registry and current acceptance recommendations.

The root's preceding review also retained this overbroad attribution. The new integration report correctly calls attention to that earlier mistake, but a prose-only correction leaves current machine-readable evidence misleading. The remedy is a narrow attribution update, not a new source-stage advancement, source reinterpretation or fingerprint refresh.

## Proposed current attestation and integration instructions

The supplied eleven-record update map identifies the current reviewer as **ChatGPT independent integration provenance reviewer (retained source acceptance)**. This reviewer did not implement the records, make the D20 source decisions, author commit `6647eae`, or author the parent's eventual metadata integration. The attestation is based on the historical evidence-chain audit described above and retention of the already accepted scoped interpretations.

For each record, the map:

1. Preserves the accepted semantic fingerprint, examples, stage and interpretation.
2. Identifies the actual first-half or second-half compendium review at `49ad425`.
3. Includes the later supplemental acceptance for the witness and negative-ethos corrections where required.
4. Distinguishes Claude's owner-directed D20 adjudication from independent review of Claude's own earlier edits.
5. Replaces the false ledger-note sentence with accurate historical attribution.
6. Explicitly says that this is retained source acceptance after an attribution audit, not eleven new complete source comparisons.

Synchronize `tools/vv/source_ledger.yml`, `docs/reviews/independent-approvals.json`, and the affected fields of `docs/reviews/final-source-acceptance.json`. Keep the existing `retained_prior_acceptance` classification. Add a dated attribution-update note rather than rewriting the acceptance register's original reviewed commit or historical raw-source hashes. Preserve historical reports and their original qualifications.

Copy this report into the repository as `docs/reviews/integration-provenance-review.md` so the proposed evidence references resolve. Correct the opening D20 attribution paragraph. Then run source integrity, compare all 149 semantic fingerprints, and obtain the promised separate check of the actual integrated metadata. No formula, source locator, CQ mapping or expected example output should change through this operation.

## Release ordering and readiness

The user-provided next-step list places tagging before the non-author check of the new integration change. The established plan instead requires publication at the exact accepted commit. The useful order is:

1. Complete independent acceptance of the server admission change and the narrow attribution update.
2. Record the actual accepted commit, source/engine versions and applicable execution evidence.
3. Tag or publish that exact preview when authorized and technically available.
4. Begin or continue the preregistered validation tasks against that pin.

The current archive's previous acceptance remains evidence for engine walton.6 at `da1dac5`, not automatic evidence for walton.7. The release notes correctly disclose that distinction; root verification of current CI and the new server behavior supplies the new component evidence. A future attribution-only follow-up commit must likewise be identified honestly if that becomes the accepted pin.

Remote tag permission is a distribution issue, not a scientific prerequisite for defining tasks, freezing samples or running a local pilot at a full commit hash. It should not delay useful protocol work. Conversely, the existence of a tag must not be used to imply completed empirical validation.

There is no new source-content dispute requiring another broad opinion round. The remaining item found here is a concrete, bounded metadata correction. General human annotation validity, input-conversion fidelity, capacity, external semantics and source authentication remain separate validation claims, as the plan already states.

## Limits of this review

This review checks repository history, recorded source judgments and current approval integrity. It does not independently authenticate reviewer identities, rerun every historical local command claimed by the implementer, or re-read every primary source. The reported local all-example rerun is consistent with the prior evidence but is not newly executed by this reviewer. Runtime, server and hosted-CI checks belong to the root review. The evidence chain and current non-author role are technical AI-assisted evidence; human application-reference work remains a separate activity.
