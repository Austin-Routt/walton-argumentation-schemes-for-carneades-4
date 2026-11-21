# Plan after the PR15 and PR16 closure review

Updated 2026-11-21. Reviewed public snapshots: original PR15 `e6d2d6f427c6ec0447cae514810c256903b15779` and later PR16 `ea46ab56c5ebd6fe76f409053cc2f008b13f38cc`. The corrected continuation is **library 0.3.0 / engine 4.3.1+walton.5**. This plan supersedes earlier operational instructions; the preceding plan is preserved in `docs/reviews/plan-before-pr15-closure.md`.

## Decision and current evidence

Proceed with the corrected **preview** only at the exact commit in the final clean acceptance evidence. Do not publish old R1 pin `8c092b6` or the PR15 baseline as this corrected candidate. Do not label the complete library 1.0.0. Fifteen inherited-source holds remain, despite repaired engine behavior and passing executable examples.

The owner decisions are settled: add explicit finite part/whole and source-alternative variants; correct the likelihood assessments in33.3 and33.4 without conflating action and omission using semantic migration; retain qualified AS1 separately from a future full AS2 model; implement question-induced suspension in a separate finite dialogue layer. No broad opinion round or majority vote is needed. Reopen concrete counterexamples or newly available source evidence.

| Item | Current state |
|---|---|
| Records |143:119 book-origin and24 inherited |
| Independent approvals |128:119 book-origin and9 inherited |
| Holds |15 inherited:13 transcription-checked and2 mapped |
| Original50 holds |35 closed;15 retained with exact next actions |
| Additions |23 variants, all independently accepted within declared scope |
| Compendium mapping |137 records under109 entries;6 upstream-only |
| Entry approval coverage |103 entries have at least one approved record;2,3,26,29,33.1,33.2 retain inherited-source holds |
| Registered inventory |9,979 distinct checks, plus separate Python/native/unit/vet/race gates |
| Validator |0 errors;11 warnings:6 unused inherited templates and5 dependency-cycle groups |
| Stable gateL |Open; every one of the143 rows must reach example-passed |

The current complete register and every original-hold outcome are `docs/reviews/closure-current-register.json` and its readable companion. The exact test inventory is `tools/vv/expected_checks.txt`. Historical reports retain their original counts and do not describe current status.

## Completed closure work

1. Independently reproduced PR15's8,230 checks in a clean Go1.24.13 clone and checked hosted results for both required versions. The exact earlier defeated-default counterexample is a real repaired defect; the implementer's dismissal tested a different premise. Its table has nine rows, not eight.
2. Retained useful ground lookup, larger finite budgets, substitution-chain termination and normalized exception comparison. Repaired noncanonical API lookup identities, cross-layer double negation, authored-premise order and explicit default-to-stipulation promotion, with non-author review.
3. Strengthened every case expectation through both CLI round trips without re-inference. Preserved existing assumptions/questions and supplied-fact origins. API observation now recognizes implicit linked arguments consistently with exported graphs. The fresh reviewer's probes and all intermediate failures are retained.
4. Repaired participant/context leakage, conditional reversal, conclusion strength and missing branches. Added six part/whole variants, four supplemental source forms and thirteen source alternatives. Regenerated categories and all four worked examples.
5. Implemented `tools/dialogue_snapshot.py`: exact argument targeting, question/accepted-answer/reopen states, retained original defeaters and independent support, strict journals and conservative answer-dependency checks. All30 CQs across six families run against the actual engine. Explicit allocation metadata covers four ignorance families.
6. Integrated public PR16 history and its failed-join cancellation fixture, corrected the universal timing claim, and repaired fresh-variable issue matching under the documented per-anchor alternatives contract. Repeated inference is stable and all matching alternatives participate in consistency checks. PR16 D23 also exposed the remaining33.4 threshold defect; its omission-specific repair has a new non-author approval.
7. Recorded non-author source judgments, source locations, adverse examples and current fingerprints. Atkinson's17 questions split into7 direct,5 partial and5 absent treatments: **ten** lack full-source coverage. The historical eight-item worklist omitted partial CQ8–9.

## Binding supported profile and decisions

The runtime profile is offline CLI/Go API, propagation rules, ground logical facts, supported terms and serialized use. Logical facts form a set; one fact may satisfy repeated premises without hidden inequality. Resolve complete issue/default closure before rule seeding and after derived bridges. Failure or cancellation is atomic; successful publication retains existing node identities. Finite quotas are explicit ceilings, with typed diagnosable refusal; larger budgets are not unrestricted capacity, universal wall time or byte-memory guarantees.

API callers structurally validate authored graphs before inference. Use `AddAssumption` to promote an existing default; direct public-field changes require maintaining origin invariants. Preserve premises, defaults, supported exceptions, rebuttals and dialogue questions as distinct operations. The finite adapter requires application judgments of question appropriateness and answer adequacy; it is not autonomous natural-language assessment or a complete dialogue calculus.

Use named participants/cases when their identities matter. Legacy unqualified first/second-person forms have one declared context; use named variants or valid context-qualified terms for broader graphs. Finite complete-partition and universal-negative witnesses are supplied evidence, not conclusions from an incomplete list or missing information. Do not install a global equivalence between practical choices and negations of the opposite choice.

The compendium governs numbered book interpretations; its narrative may justify a disclosed correction. Inherited originals retain separate provenance. A later reprint or newly approved book-form record does not certify an unseen original. External truth/authentication, numerical calibration, institutional burden allocation, full AS2 deliberation, web-service hosting and concurrent mutation of one graph remain outside this accepted profile.

## Source stages and stable gateL

The stages are cumulative: `mapped` locates provenance; `transcription-checked` compares accessible source content; `interpretation-approved` records independent judgment; `questions-sorted` supplies every printed CQ's justified disposition; `example-passed` adds actual positive and discriminating adverse execution. Tests alone cannot advance source stages. Verified zero printed questions differs from an unknown count.

The approval checker binds schemes, statically referenced language, selected ledger semantics and global issue schemes to reviewer stamps and existing case IDs. It does not authenticate reviewers, application evidence, dynamic templates, engine or adapter code. Final commit-matched acceptance binds those components together.

```sh
python3 tools/check_source_approvals.py --repo .
python3 tools/check_source_approvals.py --repo . --require-stable
```

The first must pass; the second must refuse the fifteen current holds. This is the intended complete-library gate, not an executable failure waiver or a silent scope reduction. All143 records are still required for full1.0.0.

## Remaining sequence

### A. Corrected preview

1. Review every expected-check addition/removal and freeze a committed candidate. Updating the inventory is development work, never acceptance.
2. In clean full clones of that exact commit, run `tools/vv/engine-tests.sh` and `VV_STRICT=1 tools/vv/run.sh` on pinned Go1.24.13 and1.27.4. Require unique Boolean results, exact manifest agreement and check-by-check cross-version agreement. Record actual versions, environment, exits, clean before/after states and hashes. Neither job is informational.
3. Produce source archive, Git bundle, cumulative repository patch and separately named patched engine sources. Verify clean application and resulting tree equality. Root `carneades-4.zip` and `GoCHR.zip` remain unchanged original provenance inputs with `sources.lock`; they alone do not install patches.
4. Publish only the exact accepted candidate using `docs/releases/r1-candidate.md`. Without authenticated GitHub write access, deliver the tested artifacts and exact commands. A local bundle/tag or prepared command is not a completed remote release.

### B. Thirteen original-source obligations

Obtain the actual cited1995 pages,1997 expert-opinion pages, correct slippery-slope edition/pages, relevant2008 witness passage,2015 instrumental-reasoning passage and attributable2019 correspondence. The current register identifies every affected record and exact comparison. Relevant2002/2006 indexed text and the2003 article have already resolved some access claims; they are not blanket substitutes for other originals. Preserve any original-form differences even when the compendium adaptation is approved.

### C. Two available-source model/adaptation holds

Keep `practical_reasoning1` as qualified AS1 in the preview outside full source approval. A full AS2 implementation is a separate finite-model extension with agents, component/joint actions, states, admissible transitions, changed goals, transition-relative values, preferences, all17 CQ dispositions and source reference cases. Adding five absent flat exception names is insufficient. That larger model is explicitly deferred beyond this preview, not waiting for another owner choice.

Keep `sunk_costs` as a qualified legacy quantity/feasibility/future-loss composite. Its source is accessible; it is not the article's printed p.489 precommitment form. Use the approved `sunk_costs_precommitment` for that form. Complete source certification of the composite's selected practical-question policy remains a named interpretation obligation, not a source-access claim.

### D. StableR2 and outside contacts

After those fifteen rows are resolved, independently review all semantic migrations and rerun stable gateL and both clean toolchains before releasing1.0.0. No completion date is inferred for unavailable evidence.

Source-root-specific upstream patches and reproductions are prepared and mechanically verified. Maintainers still decide compatibility and adoption of downstream policy. The source-question memorandum is prepared. Outside messages remain unsent without explicit owner authorization; upstream acceptance is not required for local preview validation.

## Review completion

Deliver one archive containing the exact candidate, complete current report, all per-record decisions, independent execution evidence, verified patches and the preserved earlier reviews/source materials. Distinguish completed local acceptance, access-blocked remote publication and complete-library source holds. No additional general review round is required for those distinctions.
