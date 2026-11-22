# Framework baseline and next-stage validation plan

Updated 2026-11-21 (America/Chicago; execution and some reviewer attestations are dated 2026-11-22 UTC). The independent closure started from public PR21 `98a2335d2c9909cf597ae85a9f359f20d092ae6c` and was integrated by PR22 at `d9a25857277140a5d3f10aaf1dfa4ade0f15b91f`. The subsequent integration review produces **library 0.4.1 / engine 4.3.1+walton.8**: the source interpretations are unchanged, eleven live attributions are corrected, and the two CAES HTTP routes now enforce validation and inference failure consistently. See `docs/reviews/post-integration-readiness-review.md`. The exact tested candidate commit, clean states, commands, hashes and results are in the accompanying acceptance evidence. This plan supersedes the operational status of the preceding plan, preserved in `docs/reviews/plan-before-final-baseline.md`. Historical review reports retain their original conclusions and dates.

## Decision and current evidence

Freeze this corrected **preview as the framework baseline for the next validation effort**, once its exact commit passes the clean acceptance gates below. Do not publish old R1 pin `8c092b6`, PR15, PR21 `98a2335`, or PR22 `d9a2585` as though they contain the later accepted corrections. Complete non-author review and exact-commit gates before tagging; a full Git commit and hashes are sufficient to begin protocol preparation without a remote tag. Source closure and baseline verification are complete only within the explicit contracts; they do not establish empirical validity, broad interchange conformance or benchmark-scale capacity. Stable version 1.0.0 remains a later release decision based on the agreed validation evidence.

The owner decisions remain settled: explicit finite part/whole and source-alternative variants; separate action and omission likelihood assessments with manual semantic migration; the compendium as source of record for numbered entries (D20); separate AS1 and finite AATS/AS2 representations (D24); and an optional finite dialogue layer. No broad opinion round is required. Reopen reproducible counterexamples or new source evidence. The independent source and implementation acceptance of this round is recorded in `docs/reviews/final-source-independent-review.md`, `final-atkinson-independent-review.md`, and `final-baseline-independent-review.md`.

| Item | Current state |
|---|---|
| Records | 149: 120 book-origin, 24 inherited and 5 from Atkinson and Bench-Capon (2007) |
| Approvals | 149 scoped approvals with explicit non-author evidence. The closure replaced nine disclosed self-stamps, independently accepted the D20 sunk-cost reclassification, refreshed two further citations, and corrected eight historical attributions. The post-integration review also corrects eleven live D20 attributions while preserving all 149 semantic fingerprints and example lists. |
| Holds | none (`docs/reviews/final-closure.md`) |
| Original 50 holds | all closed (35 by the closure review, 12 under D20, 3 in the final closure) |
| Additions since the 120-record library | 29: the earlier 23 variants, five AS1/AS2 records and the omission branch of 33.3; all accepted within their declared scope. |
| Compendium mapping | 138 records under 109 entries; 6 upstream-only; 5 publication-only |
| Entry approval coverage | every one of the 109 represented entries has an approved record |
| Registered inventory | 10,428 distinct checks (9,989 before the final closure), plus separate Python/native/unit/vet/race gates |
| Validator | 0 errors; 10 warnings: 5 unused inherited templates and 5 dependency-cycle groups |
| Source-stage gate L | 149 of 149 rows at example-passed with current fingerprints and non-author attestations; run by `tools/vv/run.sh`. This is not the empirical or complete-release gate. |

The current complete acceptance register is `docs/reviews/final-source-acceptance.json`, with live stages in `tools/vv/source_ledger.yml`. `docs/reviews/closure-current-register.json` and its readable companion preserve the earlier 143-record closure and original-hold outcomes. The exact test inventory is `tools/vv/expected_checks.txt`. Historical reports retain their original counts and do not describe current status.

## Completed closure work

1. Independently reproduced PR15's 8,230 checks in a clean Go 1.24.13 clone and checked hosted results for both required versions. The exact earlier defeated-default counterexample is a real repaired defect; the implementer's dismissal tested a different premise. Its table has nine rows, not eight.
2. Retained useful ground lookup, larger finite budgets, substitution-chain termination and normalized exception comparison. Repaired noncanonical API lookup identities, cross-layer double negation, authored-premise order and explicit default-to-stipulation promotion, with non-author review.
3. Strengthened every case expectation through both CLI round trips without re-inference. Preserved existing assumptions/questions and supplied-fact origins. API observation now recognizes implicit linked arguments consistently with exported graphs. The fresh reviewer's probes and all intermediate failures are retained.
4. Repaired participant/context leakage, conditional reversal, conclusion strength and missing branches. Added six part/whole variants, four supplemental source forms and thirteen source alternatives. Regenerated categories and all four worked examples.
5. Implemented `tools/dialogue_snapshot.py`: exact argument targeting, question/accepted-answer/reopen states, retained original defeaters and independent support, strict journals and conservative answer-dependency checks. All 30 CQs across six families run against the actual engine. Explicit allocation metadata covers four ignorance families.
6. Integrated public PR16 history and its failed-join cancellation fixture, corrected the universal timing claim, and repaired fresh-variable issue matching under the documented per-anchor alternatives contract. Repeated inference is stable and all matching alternatives participate in consistency checks. PR16 D23 also exposed the remaining 33.4 threshold defect; its omission-specific repair has a new non-author approval.
7. Recorded non-author source judgments, source locations, adverse examples and fingerprints. The earlier comparison of the inherited practical record with Atkinson found seven direct, five partial and five absent treatments. It remains historical evidence; that record now claims the verified 2015 Carneades configuration, while five separate publication records represent the 2007 AS1/AS2 interpretations.
8. Independently reviewed the newly supplied originals and all nine final-closure source changes, plus the D20 sunk-cost reclassification. Corrected source locators, the conditional-threat and Smart Car descriptions, and reviewer attribution; made independence an explicit integrity requirement.
9. Repaired state-pair effect aliasing, model-input loss and CQ diagnostics in the AATS adapter. Replaced its circular expected-argument check with a fact-derived comparison and authored actual-engine controls. Recorded Fig. 5 sensitivity and the requirement for an explicit model before comparing AS1 alternatives.
10. Repaired the known native YAML/AIF and Dung import/test defects, and enforced an explicit experimental boundary for legacy CAF/AGXML/LKIF conversion. Corrected discovered legacy defaults and structural faults without claiming full format conformance.
11. Independently checked PR22's server profile gate, then repaired two uncovered HTTP wrapper defects: discarded inference errors in the form route and missing authored-argument validation in the raw route. Both use one validation/inference boundary with request cancellation and non-success HTTP statuses. Actual-handler tests, paired live-server probes and separate non-author review cover this correction; hosted-service acceptance remains excluded.
12. Corrected the eleven D20 ledger/registry/acceptance attributions and independently checked the integrated metadata. Rebased the supplied Kialo pilot's machine context against immutable current Git objects; kept its twelve proposals unadjudicated and its source-only packets free of proposed labels. This prepares validation; it is not a completed human or runtime study.

## Binding supported profile and decisions

The runtime profile is offline CLI/Go API, propagation rules, ground logical facts, supported terms and serialized use. Logical facts form a set; one fact may satisfy repeated premises without hidden inequality. Resolve complete issue/default closure before rule seeding and after derived bridges. Failure or cancellation is atomic; successful publication retains existing node identities. Finite quotas are explicit ceilings, with typed diagnosable refusal; larger budgets are not unrestricted capacity, universal wall time or byte-memory guarantees.

API callers structurally validate authored graphs before inference. Use `AddAssumption` to promote an existing default; direct public-field changes require maintaining origin invariants. Preserve premises, defaults, supported exceptions, rebuttals and dialogue questions as distinct operations. The finite adapter requires application judgments of question appropriateness and answer adequacy; it is not autonomous natural-language assessment or a complete dialogue calculus.

Use named participants/cases when their identities matter. Legacy unqualified first/second-person forms have one declared context; use named variants or valid context-qualified terms for broader graphs. Finite complete-partition and universal-negative witnesses are supplied evidence, not conclusions from an incomplete list or missing information. Do not install a global equivalence between practical choices and negations of the opposite choice.

The compendium governs numbered book interpretations; its narrative may justify a disclosed correction. Inherited originals retain separate provenance. A later reprint or newly approved book-form record does not certify an unseen original. External truth/authentication, numerical calibration, institutional burden allocation, general VAF semantics, multi-step planning, web-service hosting and concurrent mutation of one graph remain outside this accepted profile.

## Decision register

The register of decisions that tools, contracts, ledger notes and reviews cite as `PLAN.md Dn`. D1 to D14 are as clarified by the first review before pull request #16; D15 to D23 were proposed by the second review in #16 (full text in `docs/reviews/peer-16/` and in the history of this file) and adjudicated by the first review's closure; the last two columns say what this repository does and the implementer's view (`docs/reviews/implementer-second-response.md`).

### D1 to D14

| Decision | Required contract |
|---|---|
| D1: engine ownership | Keep sources in `third_party/`, the original archives and `sources.lock`; no external forks. Preserve focused patches. Coordinated adapter/engine fixes may share a commit when their common contract requires it. Prepare upstream patches by source root. |
| D2: questions | Preserve premises, challengeable defaults, exceptions and rebuttals. Record positive support, alternatives and dialogue/proof-standard questions separately where required. Asking a question is not the same as proving an exception; a proof standard does not implement burden allocation. |
| D3: supported profile | Offline CLI and Go API, propagation rules, ground logical facts and supported terms, with serialized access. Web hosting, arithmetic, guards/deletions, concurrent mutation of one graph and evidence authentication remain outside acceptance. |
| D4: versions and migration | The current corrected preview is library 0.4.1 / engine 4.3.1+walton.8 (PR22's import profile plus consistent HTTP validation/inference failures); R2 remains stable 1.0.0 after its separate release requirements. Source repairs can require semantic migration before stable release. Use the migration table/tool; never invent newly required witnesses or participant identities. |
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
| D15: release | withdraw `8c092b6`; tag the corrected candidate with engine walton.4 once a non-author has checked it | agree, at an exact accepted commit | the current candidate is engine 4.3.1+walton.8, library 0.4.1; finish independent acceptance and exact-commit gates before publishing; retain separate authorization for publication | agree; walton.5 supersedes #17's walton.4 |
| D16: questions | static evaluation of each question, as the book formalizes them; no dialogue layer | accept as the static profile; a separate finite dialogue adapter added | both: records are static; `tools/dialogue_snapshot.py` is an optional adapter for six families | agree |
| D17: variants | one variant per record, declared in the ledger; 10.2 gets a part premise | agree, and implement the omitted branches | 9.2 has four records, 10.2 three; the positive 10.2 branch needs an explicit sufficiency warrant | agree; the sufficiency premise is a stricter reading than the print, declared |
| D18: participants | as printed; exceptions name enough participants | preserve identities a conclusion needs even when elliptical | identities added where instances could leak (e.g. 19, 28, 30, 31.1, 45, 54, 58, 59.1) | agree with the closure's form |
| D19: polarity | "should not" = ¬should_do; "less plausible" = argument for the negation | reject as general equivalences; local binary issues allowed | ethos and anti-model use qualitative or restraint statements (`ethos_less_plausible`, `should_refrain`) | agree with the closure: an ad hominem lowers plausibility, it does not prove the contrary |
| D20: source of record | the compendium; older originals not a condition of the stable release | reject as discharging original-source obligations | **adopted by the owner, 2026-11-21**: 12 inherited records approved against their entries; the cited originals are recorded as not compared (`docs/reviews/d20-approvals.md`) | agree |
| D21: legacy records | three records legacy, outside the stable count; checker to learn it | reject a retrospective narrowing of the gate | no legacy exclusion: all three now have supplied sources, explicit interpretations and non-author acceptance; `practical_reasoning1` claims the 2015 configuration and separate records implement AS1/AS2 | **owner**, for `practical_reasoning1` only |
| D22: outside contacts | send only when the owner says so | agree | **owner, 2026-11-21: not to be sent**; independent research follows once the gaps are closed | agree |
| D23: 33.3 and 33.4 | likelihood adequate for the argument | agree; 33.4's omission has its own statement | adopted: `consequence_likelihood_adequate`, `omission_consequence_likelihood_adequate` | agree |

### D24 (owner, 2026-11-21)

The owner requested both AS1 and AS2. The accepted implementation comprises the source-reviewed AS1 family and a finite, goal-labelled AATS/AS2 interpretation. All seventeen questions have explicit dispositions; the comparison contract is a declared project policy, not a uniquely derived algorithm or a general VAF solver. AS1 graphs require an additional, independently justified model before that contract can choose actions. The independent review repaired state-pair effect aliasing and input/diagnostic boundaries, and documented Fig. 5's outcome sensitivity (`docs/contracts.md`, contract 8). `practical_reasoning1` remains the verified Walton (2015) Carneades configuration. D21's three source holds are closed within these claims.

### D25 to D27 (independent closure decisions)

| Decision | Accepted boundary |
|---|---|
| D25: approval evidence | Require explicit Boolean non-author acceptance and its evidence basis in the ledger and matching registry. A hash-valid self-review does not pass. Metadata reconciliation cannot reverse a withdrawal or change the reviewer automatically. The tool cannot authenticate people or prove the attestation true. |
| D26: input profiles | Native YAML is the primary structured profile. AIF is a deliberately narrow I/RA/binary-CA projection with disclosed stipulation, linked-argument and symmetric-issue policies. CAF, AGXML and LKIF require explicit experimental opt-in, on the command line and in the web service; passing structural checks is not proof of conversion fidelity. Dung TGF contains unlabelled attack edges only; labelled Kialo support is rejected. |
| D27: baseline versus validation | Accept the exact corrected preview for the next planned verification/validation/benchmarking stages. No empirical, annotation, performance, ICCMA-protocol, general AIFdb or Kialo-to-scheme claim is inferred from source closure or passing regression counts. Those are next-stage work, with separate eligibility and acceptance criteria. |

## Source stages and stable gate L

The stages are cumulative: `mapped` locates provenance; `transcription-checked` compares accessible source content; `interpretation-approved` records independent judgment; `questions-sorted` supplies every printed CQ's justified disposition; `example-passed` adds actual positive and discriminating adverse execution. Tests alone cannot advance source stages. Verified zero printed questions differs from an unknown count.

The approval checker binds schemes, statically referenced language, selected ledger semantics and global issue schemes to reviewer stamps, explicit non-author attestations and existing case IDs. It does not authenticate reviewers, application evidence, dynamic templates, engine or adapter code. Final commit-matched acceptance binds those components together.

```sh
python3 tools/check_source_approvals.py --repo .
python3 tools/check_source_approvals.py --repo . --require-stable
```

Both must pass for the corrected final candidate, and `tools/vv/run.sh` runs the second. The earlier v1 integrity pass did not itself establish the nine admitted self-reviews were independent. A record added later has to reach `example-passed` with an approval before it can be merged.

## Remaining sequence

### A. Freeze and verify the corrected baseline

1. Review every expected-check addition/removal and freeze a committed candidate. Updating the inventory is development work, never acceptance.
2. In clean full clones of that exact commit, run `tools/vv/engine-tests.sh` and `VV_STRICT=1 tools/vv/run.sh` on pinned Go 1.24.13 and 1.27.4. Require unique Boolean results, exact manifest agreement and check-by-check cross-version agreement. Record actual versions, environment, exits, clean before/after states and hashes. Neither job is informational.
3. Produce source archive, Git bundle, cumulative repository patch and separately named patched engine sources. Verify clean application and resulting tree equality. Root `carneades-4.zip` and `GoCHR.zip` remain unchanged original provenance inputs with `sources.lock`; they alone do not install patches.
4. Deliver the exact candidate, reports, evidence and patches. Publishing is optional and separate: use `docs/releases/framework-0.4.1.md` and the archive's commit-matched instructions only after publication is authorized. A local branch, bundle or prepared command is not a completed remote tag, pull request or release.

### B. Original sources (closed)

Under D20 the twelve inherited records filed under compendium entries are approved against those entries. On 2026-11-21 the owner supplied Walton (1995), Walton (2015), Walton (2002) and Walton, *Scare Tactics* (2000): `practical_reasoning2` and `exeption_to_rule` were compared, corrected and approved; the six records citing *Scare Tactics* and nineteen others approved against the compendium have their original compared in a ledger note; 33.3's printed omission reading, which *Scare Tactics* (pp. 139-140) explains, became `negative_consequences_reasoning_omission` (`docs/reviews/final-closure.md`). The name `exeption_to_rule` comes from Walton's correspondence of 2019-09-19, recorded in the README's Updates of 2019-10-08; the e-mail is private, and the owner confirms the README as its official record. Works cited but not supplied (*Appeal to Expert Opinion*, *Slippery Slope Arguments*, *Witness Testimony Evidence* and the 2003 ICAIL paper, and the compendium's references for book-origin records) are research under D20, not conditions.

### C. AS1 and the finite AATS/AS2 profile (closed within D24 scope)

`practical_reasoning1` is the value-based scheme of Walton (2015, pp. 92-93). AS1, its negative forms and AS2 are separate accepted interpretations. `tools/aats.py` validates its declared finite-model profile, generates facts, exposes question diagnostics and implements comparison contract 8. Current source cases, independent fact joins and actual-engine regressions pass. Literal printed definitions, operational repairs, conditional AS1 comparison and the Fig. 5 sensitivity remain explicitly distinguished.

### D. Begin the separate validation effort

Use [the validation-entry protocol](docs/validation-entry.md) with the accompanying V0-V6 / K0-K6 programme. First freeze task definitions, supported profiles, independent references, corpus eligibility and split rules against the accepted candidate. Use Kialo as the first application corpus for structural preservation and human annotation calibration, while developing independent bounded CAES/CHR references in parallel. Do not wait for a remote tag, and do not treat Kialo pro/con links, votes, machine proposals or engine outputs as semantic ground truth.

The historical Kialo packet is retained unchanged. Its old `fcccbba` pin, engine walton.5 and 143-record source register are historical; use `tools/validation/refresh_kialo_context.py` to create a new derivative from the exact accepted commit, live 149-record ledger and approval registry. The supplied twelve proposals remain unadjudicated development material. A six-occurrence source-only exercise is for annotation time, ambiguity and manual calibration, not an accuracy estimate. The existing 60-dossier feasibility budget (12 calibration, 12 development, 36 sealed evaluation) is a planning budget, not statistical certification; additional eligible discussion families and people are still needed.

Then run the planned semantics, interoperability, capacity and independently annotated real-world studies with their own references and acceptance criteria. ICCMA abstract frameworks, AIFdb and natural-language corpora test different layers; none alone validates the complete structured framework. The Kialo legacy exports remain incompatible inputs, now refused explicitly; a preserving structural adapter and a compiler of reviewed scenarios are separate next-stage tasks. External generator repairs, a general AIF identifier mapping and full collection annotation are future benchmark work, not silently closed framework claims.

### E. Stable R2 and publication

Before 1.0.0, review the completed validation evidence against its preregistered criteria, review any resulting semantic migrations, and rerun source gate L, both complete toolchains and the exact result comparison on the commit intended for release. This round closes the specified source and implementation review obligations; it does not stand in for human application validation or author endorsement. Do not mark 1.0.0 merely because every ledger row says example-passed.

The upstream patches and source memorandum are not to be sent (owner, 2026-11-21, D22). No outside contact is necessary to accept this scoped baseline.

## Review completion

Deliver one archive containing the exact candidate, complete current report, all per-record decisions, independent execution evidence, verified patches and the preserved earlier reviews/source materials. Distinguish completed local acceptance, any separately authorized remote publication, and the remaining validation work and explicit scope limitations. No additional general review round is required for those distinctions.
