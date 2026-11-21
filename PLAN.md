# Plan after the PR15 and PR16 closure review

Updated 2026-11-21. Reviewed public snapshots: original PR15 `e6d2d6f427c6ec0447cae514810c256903b15779` and later PR16 `ea46ab56c5ebd6fe76f409053cc2f008b13f38cc`. The corrected continuation is **library 0.3.0 / engine 4.3.1+walton.5**. This plan supersedes earlier operational instructions; the preceding plan is preserved in `docs/reviews/plan-before-pr15-closure.md`.

## Decision and current evidence

Proceed with the corrected **preview** only at the exact commit in the final clean acceptance evidence. Do not publish old R1 pin `8c092b6` or the PR15 baseline as this corrected candidate. Do not label the complete library 1.0.0. Fifteen inherited-source holds remain, despite repaired engine behavior and passing executable examples.

The owner decisions are settled: add explicit finite part/whole and source-alternative variants; correct the likelihood assessments in 33.3 and 33.4 without conflating action and omission using semantic migration; retain qualified AS1 separately from a future full AS2 model; implement question-induced suspension in a separate finite dialogue layer. No broad opinion round or majority vote is needed. Reopen concrete counterexamples or newly available source evidence.

| Item | Current state |
|---|---|
| Records | 143: 119 book-origin and 24 inherited |
| Independent approvals | 128: 119 book-origin and 9 inherited |
| Holds | 15 inherited: 13 transcription-checked and 2 mapped |
| Original 50 holds | 35 closed; 15 retained with exact next actions |
| Additions | 23 variants, all independently accepted within declared scope |
| Compendium mapping | 137 records under 109 entries; 6 upstream-only |
| Entry approval coverage | 103 entries have at least one approved record; 2, 3, 26, 29, 33.1 and 33.2 retain inherited-source holds |
| Registered inventory | 9,989 distinct checks (9,979 at the closure candidate `fcccbba`, ten added by the implementer's second response), plus separate Python/native/unit/vet/race gates |
| Validator | 0 errors; 11 warnings: 6 unused inherited templates and 5 dependency-cycle groups |
| Stable gate L | Open; every one of the 143 rows must reach example-passed |

The current complete register and every original-hold outcome are `docs/reviews/closure-current-register.json` and its readable companion. The exact test inventory is `tools/vv/expected_checks.txt`. Historical reports retain their original counts and do not describe current status.

## Completed closure work

1. Independently reproduced PR15's 8,230 checks in a clean Go 1.24.13 clone and checked hosted results for both required versions. The exact earlier defeated-default counterexample is a real repaired defect; the implementer's dismissal tested a different premise. Its table has nine rows, not eight.
2. Retained useful ground lookup, larger finite budgets, substitution-chain termination and normalized exception comparison. Repaired noncanonical API lookup identities, cross-layer double negation, authored-premise order and explicit default-to-stipulation promotion, with non-author review.
3. Strengthened every case expectation through both CLI round trips without re-inference. Preserved existing assumptions/questions and supplied-fact origins. API observation now recognizes implicit linked arguments consistently with exported graphs. The fresh reviewer's probes and all intermediate failures are retained.
4. Repaired participant/context leakage, conditional reversal, conclusion strength and missing branches. Added six part/whole variants, four supplemental source forms and thirteen source alternatives. Regenerated categories and all four worked examples.
5. Implemented `tools/dialogue_snapshot.py`: exact argument targeting, question/accepted-answer/reopen states, retained original defeaters and independent support, strict journals and conservative answer-dependency checks. All 30 CQs across six families run against the actual engine. Explicit allocation metadata covers four ignorance families.
6. Integrated public PR16 history and its failed-join cancellation fixture, corrected the universal timing claim, and repaired fresh-variable issue matching under the documented per-anchor alternatives contract. Repeated inference is stable and all matching alternatives participate in consistency checks. PR16 D23 also exposed the remaining 33.4 threshold defect; its omission-specific repair has a new non-author approval.
7. Recorded non-author source judgments, source locations, adverse examples and current fingerprints. Atkinson's 17 questions split into 7 direct, 5 partial and 5 absent treatments: **ten** lack full-source coverage. The historical eight-item worklist omitted partial CQ8–9.

## Binding supported profile and decisions

The runtime profile is offline CLI/Go API, propagation rules, ground logical facts, supported terms and serialized use. Logical facts form a set; one fact may satisfy repeated premises without hidden inequality. Resolve complete issue/default closure before rule seeding and after derived bridges. Failure or cancellation is atomic; successful publication retains existing node identities. Finite quotas are explicit ceilings, with typed diagnosable refusal; larger budgets are not unrestricted capacity, universal wall time or byte-memory guarantees.

API callers structurally validate authored graphs before inference. Use `AddAssumption` to promote an existing default; direct public-field changes require maintaining origin invariants. Preserve premises, defaults, supported exceptions, rebuttals and dialogue questions as distinct operations. The finite adapter requires application judgments of question appropriateness and answer adequacy; it is not autonomous natural-language assessment or a complete dialogue calculus.

Use named participants/cases when their identities matter. Legacy unqualified first/second-person forms have one declared context; use named variants or valid context-qualified terms for broader graphs. Finite complete-partition and universal-negative witnesses are supplied evidence, not conclusions from an incomplete list or missing information. Do not install a global equivalence between practical choices and negations of the opposite choice.

The compendium governs numbered book interpretations; its narrative may justify a disclosed correction. Inherited originals retain separate provenance. A later reprint or newly approved book-form record does not certify an unseen original. External truth/authentication, numerical calibration, institutional burden allocation, full AS2 deliberation, web-service hosting and concurrent mutation of one graph remain outside this accepted profile.

## Decision register

The register of decisions that tools, contracts, ledger notes and reviews cite as `PLAN.md Dn`. D1 to D14 are as clarified by the first review before pull request #16; D15 to D23 were proposed by the second review in #16 (full text in `docs/reviews/peer-16/` and in the history of this file) and adjudicated by the first review's closure; the last two columns say what this repository does and the implementer's view (`docs/reviews/implementer-second-response.md`).

### D1 to D14

| Decision | Required contract |
|---|---|
| D1: engine ownership | Keep sources in `third_party/`, the original archives and `sources.lock`; no external forks. Preserve focused patches. Coordinated adapter/engine fixes may share a commit when their common contract requires it. Prepare upstream patches by source root. |
| D2: questions | Preserve premises, challengeable defaults, exceptions and rebuttals. Record positive support, alternatives and dialogue/proof-standard questions separately where required. Asking a question is not the same as proving an exception; a proof standard does not implement burden allocation. |
| D3: supported profile | Offline CLI and Go API, propagation rules, ground logical facts and supported terms, with serialized access. Web hosting, arithmetic, guards/deletions, concurrent mutation of one graph and evidence authentication remain outside acceptance. |
| D4: versions and migration | Corrected R1 is preview 0.2.0/.2; R2 is complete stable 1.0.0. Source repairs can require semantic migration before stable release. Use the migration table/tool; never invent newly required witnesses or participant identities. |
| D5: source authority | The compendium governs numbered entries; its narrative resolves ambiguities where justified. Inherited originals retain separate provenance. A book-form adaptation is not verification of an unavailable original. Record interpretations without claiming confirmed authorial errata. |
| D6: shared vocabulary | Merge only propositions with the same meaning, participants, scope and polarity. Preserve agent, event, action, quality, case and diagram identities. A typed trait may project to a property; an arbitrary property must not predict character. |
| D7: required gates | Both pinned Go 1.24.13 and Go 1.27.4 must pass engine and strict gates, and all named results must agree. Neither version nor comparison is informational. Exact inventories, hashes and mutation tests prevent silent weakening. |
| D8: facts and work | Logical facts form a set and can satisfy repeated premises without hidden inequality. Stream matching; bound applications, term nodes, candidate visits, constraints, stored nodes and issue construction. Logical quotas are not an operating-system byte-memory or universal time bound. |
| D9: outcomes | Keep the seven typed outcomes. Failed/cancelled inference leaves the original graph unchanged; success preserves existing node identities. Resolve issue closure/default priority before seeding inference and again after derived bridges. Reject incompatible existing issues rather than choosing a proof standard silently. |
| D10: opposition | Detect contradictory stipulations through explicit negation and generated/overlapping issue schemes, even without rule firing or predeclared statement nodes. Defeated defaults must not seed further defaults. |
| D11: syntax | Reject malformed nested terms and preserve valid quoted atoms on both toolchains. Keep the supported reserved-name policy. A partial parse is not valid input. |
| D12: authored arguments | Validate a single substitution and every direct linked exception witness to the actual undercutter. An inert undercutter, wrong instance, extra prerequisite or disabled weighing function is insufficient. API clients must call `validation.Validate` for authored graphs; inference does not replace validation. |
| D13: source plus behavior | Every repair needs a source locator, explicit interpretation, positive example and a discriminating adverse example. Source judgment and executable verification are separate evidence. |
| D14: independent acceptance | Review source passages and expectations independently of the implementer's tests; rerun clean copies and record disagreement. AI-assisted review supplies technical evidence, not institutional or authorial certification. No new broad opinion round is required; reopen concrete holds/regressions. |

### D15 to D23

| Decision | Proposed in #16 | Closure review | In this repository | Implementer |
|---|---|---|---|---|
| D15: release | withdraw `8c092b6`; tag the corrected candidate with engine walton.4 once a non-author has checked it | agree, at an exact accepted commit | the candidate is engine 4.3.1+walton.5, library 0.3.0; tagging is the owner's | agree; walton.5 supersedes #17's walton.4 |
| D16: questions | static evaluation of each question, as the book formalizes them; no dialogue layer | accept as the static profile; a separate finite dialogue adapter added | both: records are static; `tools/dialogue_snapshot.py` is an optional adapter for six families | agree |
| D17: variants | one variant per record, declared in the ledger; 10.2 gets a part premise | agree, and implement the omitted branches | 9.2 has four records, 10.2 three; the positive 10.2 branch needs an explicit sufficiency warrant | agree; the sufficiency premise is a stricter reading than the print, declared |
| D18: participants | as printed; exceptions name enough participants | preserve identities a conclusion needs even when elliptical | identities added where instances could leak (e.g. 19, 28, 30, 31.1, 45, 54, 58, 59.1) | agree with the closure's form |
| D19: polarity | "should not" = ¬should_do; "less plausible" = argument for the negation | reject as general equivalences; local binary issues allowed | ethos and anti-model use qualitative or restraint statements (`ethos_less_plausible`, `should_refrain`) | agree with the closure: an ad hominem lowers plausibility, it does not prove the contrary |
| D20: source of record | the compendium; older originals not a condition of the stable release | reject as discharging original-source obligations | not adopted: 13 inherited records stay held for originals | **owner's decision**; either is honest if the ledger states the scope |
| D21: legacy records | three records legacy, outside the stable count; checker to learn it | reject a retrospective narrowing of the gate | not adopted: they stay held; the checker has no legacy class | **owner's decision**, together with D20 |
| D22: outside contacts | send only when the owner says so | agree | nothing sent | agree |
| D23: 33.3 and 33.4 | likelihood adequate for the argument | agree; 33.4's omission has its own statement | adopted: `consequence_likelihood_adequate`, `omission_consequence_likelihood_adequate` | agree |

## Source stages and stable gate L

The stages are cumulative: `mapped` locates provenance; `transcription-checked` compares accessible source content; `interpretation-approved` records independent judgment; `questions-sorted` supplies every printed CQ's justified disposition; `example-passed` adds actual positive and discriminating adverse execution. Tests alone cannot advance source stages. Verified zero printed questions differs from an unknown count.

The approval checker binds schemes, statically referenced language, selected ledger semantics and global issue schemes to reviewer stamps and existing case IDs. It does not authenticate reviewers, application evidence, dynamic templates, engine or adapter code. Final commit-matched acceptance binds those components together.

```sh
python3 tools/check_source_approvals.py --repo .
python3 tools/check_source_approvals.py --repo . --require-stable
```

The first must pass; the second must refuse the fifteen current holds. This is the intended complete-library gate, not an executable failure waiver or a silent scope reduction. All 143 records are still required for full 1.0.0.

## Remaining sequence

### A. Corrected preview

1. Review every expected-check addition/removal and freeze a committed candidate. Updating the inventory is development work, never acceptance.
2. In clean full clones of that exact commit, run `tools/vv/engine-tests.sh` and `VV_STRICT=1 tools/vv/run.sh` on pinned Go 1.24.13 and 1.27.4. Require unique Boolean results, exact manifest agreement and check-by-check cross-version agreement. Record actual versions, environment, exits, clean before/after states and hashes. Neither job is informational.
3. Produce source archive, Git bundle, cumulative repository patch and separately named patched engine sources. Verify clean application and resulting tree equality. Root `carneades-4.zip` and `GoCHR.zip` remain unchanged original provenance inputs with `sources.lock`; they alone do not install patches.
4. Publish only the exact accepted candidate using `docs/releases/r1-candidate.md`. Without authenticated GitHub write access, deliver the tested artifacts and exact commands. A local bundle/tag or prepared command is not a completed remote release.

### B. Thirteen original-source obligations

Obtain the actual cited 1995 pages, 1997 expert-opinion pages, correct slippery-slope edition/pages, relevant 2008 witness passage, 2015 instrumental-reasoning passage and attributable 2019 correspondence. The current register identifies every affected record and exact comparison. Relevant 2002/2006 indexed text and the 2003 article have already resolved some access claims; they are not blanket substitutes for other originals. Preserve any original-form differences even when the compendium adaptation is approved.

### C. Two available-source model/adaptation holds

Keep `practical_reasoning1` as qualified AS1 in the preview outside full source approval. A full AS2 implementation is a separate finite-model extension with agents, component/joint actions, states, admissible transitions, changed goals, transition-relative values, preferences, all 17 CQ dispositions and source reference cases. Adding five absent flat exception names is insufficient. That larger model is explicitly deferred beyond this preview, not waiting for another owner choice.

Keep `sunk_costs` as a qualified legacy quantity/feasibility/future-loss composite. Its source is accessible; it is not the article's printed p.489 precommitment form. Use the approved `sunk_costs_precommitment` for that form. Complete source certification of the composite's selected practical-question policy remains a named interpretation obligation, not a source-access claim.

### D. StableR2 and outside contacts

After those fifteen rows are resolved, independently review all semantic migrations and rerun stable gate L and both clean toolchains before releasing 1.0.0. No completion date is inferred for unavailable evidence.

Source-root-specific upstream patches and reproductions are prepared and mechanically verified. Maintainers still decide compatibility and adoption of downstream policy. The source-question memorandum is prepared. Outside messages remain unsent without explicit owner authorization; upstream acceptance is not required for local preview validation.

## Review completion

Deliver one archive containing the exact candidate, complete current report, all per-record decisions, independent execution evidence, verified patches and the preserved earlier reviews/source materials. Distinguish completed local acceptance, access-blocked remote publication and complete-library source holds. No additional general review round is required for those distinctions.
