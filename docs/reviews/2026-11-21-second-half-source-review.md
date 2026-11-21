> **Dated follow-up:** this independent audit is preserved in full. Subsequent PR13 corrections and accepted aliases are recorded in `peer13-second-followup.md`; `record-dispositions.md` is the current consolidated decision register.

# Independent source audit: compendium entries 28–60

## Decision

The original green suite did not establish complete source correctness. I independently reproduced three materially wrong outcomes in the unchanged baseline: a diagram treated as its conclusion, a person's memory defeater applied to another person's independent recall, and class inclusion treated as synonymy. The defects were accepted by the earlier fixtures or not exercised by them. The role audit also found missing action, quality and event identities in the three character/retroduction records. Those have bounded repairs and independent adverse controls in this patch.

The available source permits decisions without another broad review cycle. It does not justify a blanket stable-release approval. This audit approves 35 declared compendium interpretations and their bounded examples, qualifies 15 records for specific source/context limits, and blocks exact-compendium approval of the negative-ethos contradiction conclusion. The canonical ledger advances the approved 35 to `example-passed`; the remaining 16 retain their previous stage and explicit reasons. The root reviewer must review the new repair patch and run the integrated strict suite before a release decision.

## Evidence and independence

Baseline: commit `49ad425cbf7afeebcbd9a07ee479d60681ca33ea`, the clean repository state assigned by the root reviewer. Work was performed in a separate worktree, `review/book-second`. The source is the user-supplied 453-page *Argumentation Schemes* PDF, SHA-256 `d010348014e8b80fa561b68e5fac732cce0b3905fb0b1d522dd6fd2578a69634`. Main printed pages map to PDF pages by adding ten. All source pages 328–346 were read as extracted text; the nineteen corresponding rendered PDF pages were visually reviewed, with decisive full-page inspections for the role, source-polarity and conditional-warrant questions. This report paraphrases the source and does not redistribute its pages.

The scope contains 51 records mapped under 48 ledger entry labels, including both forms of entry 41, both inherited components under 49, and obligation/prohibition under 55.3. Each record's premises, conclusion, critical-question classification, source-ledger assertions, and relevant book/phase/reference cases were checked. A passing fixture was treated as evidence about execution only; its expected answer was separately compared with the source.

For original works cited by the compendium, reading the compendium does not prove those original works were inspected. In particular, the 1995 sources for inherited consequence and established-rule forms remain unverified. The root reviewer independently obtained and checked the original 2006 appearance article, pp. 324–325; combined with this audit's compendium comparison, that specific source hold can close. Nothing here independently certifies the sunk-costs article, unavailable books or correspondence outside this scope.

## Implemented repairs and decisions

1. **Forward abduction:** the strongest diagram G1 belongs to G; `diagram_conclusion(G1,C1)` links it to its conclusion. The parameter vector appends G1. Old graph instances require actual diagram identities, not a name-only rewrite.
2. **Character and retroduction:** action X, quality Q and event E are explicit. A typed character-quality fact feeds the predictive variant, while its generic property projection can feed retroduction. A negative control prevents arbitrary color/taxonomic properties from generating character predictions. Retroduction requires evidence about the same event and quality.
3. **Constitutive rule 54.1:** coextension is a separate required premise, supplying the printed warrant's antecedent. The previous whale/mammal positive fixture is now an adverse witness; a defining triangle/three-sided-polygon pair is positive. This is a supported executable interpretation, not a claimed publisher erratum.
4. **Memory:** both original-belief exceptions now carry recaller and proposition. Same-recaller defeat, other-recaller isolation, surviving independent support and defeat of both supports are independently checked.
5. **Exceptional cases:** the aggregate default is given a precise conditional meaning. Recognition is an applicable recognized exception; an unrecognized case needs evidence of rule non-applicability; a borderline case additionally needs comparable-case justification. Four challenged-witness routes and four negative controls operationalize that policy. The core remains defeasible when undisputed; the engine does not authenticate external evidence documents.
6. **Consequence likelihood:** the supplied book is sufficient to add challengeable action/outcome likelihood-adequacy defaults to inherited positive and negative consequences. This closes the available-book CQ1 gap without imposing a fixed probability or the unrelated requirement “strongly likely.” Questioning, support recovery and wrong-context controls are included. Original-source verification stays distinct.
7. **Inherited established-rule questions:** conflicting/overriding rules and exceptional/excusing cases now have R/C-scoped exception hooks from the available compendium. The inherited rule/conclusion formulation remains qualified against the actor/action form and the original 1995 source.
8. **Migration reporting:** new ordinary premises previously had empty `old` fields and were silently skipped by the migrator's manual-action detection. Optional `trigger` metadata now reports their review requirement. No missing factual role or coextension claim is invented. Historical spelling rows that now require new roles are explicitly manual migrations.

The source contracts document the single-context and aggregate-witness boundaries of anonymous respondents, sequences, institutions, languages and pleas. These boundaries are stated as limits, not used to claim unrestricted multi-context safety. The per-record qualifications below identify where a future scope expansion requires explicit parameters and controls.

## Verification

The first twelve independent regression cases were run against the unchanged baseline and produced eight unexpected failures among 24 API/round-trip checks. Separate native CLI probes recorded the baseline’s `plausible_conclusion(butler_argument)` In, `means_same(whale,mammal)` In, and `reasonable_believe(door_locked)` Out after a shared unary history attack defeated both recallers.

After repairs, the focused run of `TestLibrary`, `TestCases`, `TestSource`, and `TestCLI` completed **8,068 checks, all passing**, under Go 1.27.4. This is a selected run, not the root's final integrated check count. The Python scheme validator reports **0 errors and 10 warnings across 14 files**; validator mutation/schema tests pass; eight migration tests pass. The source-stage advancement is separately rechecked after this report is built. The full expected-check manifest was deliberately not regenerated by a partial run. Root integration must regenerate it from the complete integrated suite on both required toolchains.

## Changes to the previous completion claims

“All checks pass” was accurate as a statement about the previous suite; it was insufficient as a claim that all reference expectations matched the book. Some newly found defects were acknowledged in ledger notes yet still described operationally as finished. The unmodeled questions were not all blocked by inaccessible books: four could be implemented as explicit adaptations from the supplied compendium while keeping original-source verification open. Conversely, source availability does not eliminate the negative-ethos strength problem or make a component slope decomposition equal to the full normative schema.

The plan should retain stage separation: transcription, interpretation, CQ sorting and example validation are different judgements. Do not redefine interpretation approval to mean only “passes tests.” The appropriate next step is a bounded integration review of these repairs, the root's independent engine findings, and the remaining named qualifications. A new general round of unspecific reviewer commentary is unnecessary; specific source acquisitions or representation decisions remain necessary for records still outside a source-certified release.

## Per-record adjudications
### 28 — `cause_to_effect`: qualified

**Source:** printed p./pp. 328; PDF p./pp. 338. **Stage:** transcription-checked → transcription-checked.

The two-premise causal inference and effect conclusion correspond to entry 28. The inherited form makes one occurred event and causal relation explicit; the printed may/might modality is implemented as defeasible support, not a probability value.

**Critical questions.** CQ1–CQ2 challenge the supplied causal generalization; CQ3 is interference(E1). The latter is broad: it must mean an interference with this causal application, not every possible effect of E1.

**Examples and controls.** The book fixture and exception variant demonstrate one cause/effect application, not effect-specific isolation for several effects of one cause.

**Recorded mapping:** CQ1: premise `causes(E1,E2)`; CQ2: premise `causes(E1,E2)`; CQ3: exception `interference(E1)`.

**Remaining condition.** Retain an inherited-form qualification and the single causal-context interpretation. The stale ledger note saying no meta.source is historical: the current record does have a compendium source. No separate original book was independently verified here.

### 29 — `correlation_to_cause`: qualified

**Source:** printed p./pp. 328, 329; PDF p./pp. 338, 339. **Stage:** transcription-checked → transcription-checked.

The direction correlation(E1,E2) → causes(E1,E2) follows entry 29. It is defeasible, not a demonstration that correlation alone establishes causation.

**Critical questions.** CQ1 is the correlation premise; CQ3 is a common-cause exception. CQ2 is implemented by a challengeable explanatory theory. That is a selected evidential policy for non-coincidence, not a literal equivalence between the two.

**Examples and controls.** The positive example and common-cause variant exercise the selected policy. They cannot validate the empirical causal theory.

**Recorded mapping:** CQ1: premise `correlated(E1,E2)`; CQ2: assumption `explanatory_theory(E1,E2)`; CQ3: exception `causes_both(E1,E2)`.

**Remaining condition.** Original cited 1995 work not independently read. Preserve the explanatory-theory policy qualification and require an application to justify causal evidence.

### 30 — `sign`: qualified

**Source:** printed p./pp. 329; PDF p./pp. 339. **Stage:** transcription-checked → transcription-checked.

The sign and indicated event are correctly separated in the ordinary premises and conclusion. The comparative alternative defeater identifies the sign alone.

**Critical questions.** CQ1 is pair-scoped strong_correlation(A,B). CQ2 more_reliable(A) compresses the competing-event comparison; with several candidate events for one sign it may suppress every candidate, including an independently better one.

**Examples and controls.** Existing positive/defeat and different-sign examples cover a single interpretation per sign. They do not establish candidate-specific defeat for multiple readings of the same sign.

**Recorded mapping:** CQ1: assumption `strong_correlation(A,B)`; CQ2: exception `more_reliable(A)`.

**Remaining condition.** Before unrestricted multi-interpretation use, scope the comparator to the candidate (A,B) or impose and test a one-candidate-per-sign case contract. Keep below blanket example-passed approval.

### 31.1 — `abductive_backward`: qualified

**Source:** printed p./pp. 329, 330; PDF p./pp. 339, 340. **Stage:** transcription-checked → transcription-checked.

The selected account, data set and best-explanation conclusion follow 31.1. The available-accounts set A is named, but A1 is not explicitly linked to membership in that set.

**Critical questions.** CQ1 is sufficient_account(A1,D); CQ2 is the comparative best-account premise. CQ3–CQ4 share inquiry_incomplete(D), an explicit aggregate policy about whether to conclude now.

**Examples and controls.** The book case and questioned-account/incomplete-inquiry cases are adequate for one fixed candidate universe. They do not reject a best-account witness taken from a different account set for the same data.

**Recorded mapping:** CQ1: assumption `sufficient_account(A1,D)`; CQ2: premise `successful_account(A1,D)`; CQ3: exception `inquiry_incomplete(D)`; CQ4: exception `inquiry_incomplete(D)`.

**Remaining condition.** Document and enforce one candidate universe per D or add an explicit membership/comparison link before approving unrestricted multi-inquiry use.

### 31.2 — `abductive_forward`: approve

**Source:** printed p./pp. 329, 330; PDF p./pp. 339, 340. **Stage:** transcription-checked → example-passed.

Repaired. Printed pp. 329–330 distinguish diagram G1 from conclusion C1. The old unary successful_diagram allowed the diagram itself to be concluded as a hypothesis. Selection now names G and G1 and an explicit diagram_conclusion(G1,C1) supplies the conclusion.

**Critical questions.** CQ1 remains a challengeable account-adequacy aggregate; CQ2 is selection of the strongest diagram within G; CQ3–CQ4 retain the inquiry-completeness aggregate. These are declared defeasible interpretations, not numeric graph-strength calculations.

**Examples and controls.** Independent positive, wrong-set, wrong-conclusion-link, and missing-premise controls now pass; the baseline independently returned plausible_conclusion(butler_argument), demonstrating the original role error.

**Recorded mapping:** CQ1: assumption `sufficient_account(C1,D)`; CQ2: premise `successful_diagram(G,G1)`; CQ3: exception `inquiry_incomplete(D)`; CQ4: exception `inquiry_incomplete(D)`.

### 31.3 — `abductive_action_to_character`: approve

**Source:** printed p./pp. 330; PDF p./pp. 340. **Stage:** transcription-checked → example-passed.

Repaired. The person A, performed action X and particular quality Q are explicit. The typed conclusion character_quality(A,Q) preserves what kind of property the source asserts; its additional has_property projection makes the trait available to generic trait reasoning.

**Critical questions.** CQ1 is the named quality in the action premise. CQ2 is quality_defined(Q); CQ3 is action_fits_quality(A,X,Q). A challenge concerning one quality or action does not suspend another.

**Examples and controls.** The independent action→typed trait→prediction chain, a separate action-fit challenge, and distinct-quality definition challenge pass. The source conclusion remains defeasible evidence of character, not certainty from a single act.

**Recorded mapping:** CQ1: premise `action_character(A,X,Q)`; CQ2: assumption `quality_defined(Q)`; CQ3: assumption `action_fits_quality(A,X,Q)`.

### 31.4 — `abductive_character_to_action`: approve

**Source:** printed p./pp. 330; PDF p./pp. 340. **Stage:** transcription-checked → example-passed.

Repaired. The ordinary premise requires a typed character_quality(A,Q). The conclusion remains the source conditional prediction about future conduct with that quality; it does not identify any specific past actor/event.

**Critical questions.** The quality is explicit in CQ1 and CQ2. The counterexample defeater is scoped to (A,Q). The definition remains a defeasible default that can be challenged.

**Examples and controls.** A wrong-quality counterexample leaves the other prediction supported. A rose having the ordinary property red produces no character prediction. That adverse control prevents an unsafe over-unification with generic has_property.

**Recorded mapping:** CQ1: premise `character_quality(A,Q)`; CQ2: assumption `quality_defined(Q)`; CQ3: exception `not_character_action(A,Q)`.

### 31.5 — `retroductive`: approve

**Source:** printed p./pp. 331; PDF p./pp. 341. **Stage:** transcription-checked → example-passed.

Repaired. Event E is now present in candidate-agent and event-quality premises, the conclusion, the fit question and reference class. The source distinguishes this backward attribution from forward character prediction.

**Critical questions.** CQ1 identifies the quality through event_quality(E,Q); CQ2 concerns that Q; CQ3 concerns A/E/Q; CQ4 concerns other agents who could have brought about this E and possess Q.

**Examples and controls.** Independent wrong-event and wrong-quality cases do not produce an attribution. A questioned reference class for one event does not defeat attribution for another. The action-to-character projection can feed this record while event-specific evidence remains necessary.

**Recorded mapping:** CQ1: premise `event_quality(E,Q)`; CQ2: assumption `quality_defined(Q)`; CQ3: assumption `action_fits_quality(A,E,Q)`; CQ4: assumption `reference_class(A,E,Q)`.

### 32.1 — `evidence_hypothesis_verification`: approve

**Source:** printed p./pp. 331, 332; PDF p./pp. 341, 342. **Stage:** transcription-checked → example-passed.

The conditional H→observed A and observed-A premise support H in the printed defeasible verification schema. This is intentionally defeasible confirmation, not a deductively valid affirmation-of-the-consequent rule.

**Critical questions.** CQ1–CQ2 test the two ordinary premises. CQ3 is the chosen sufficient_correlation(H,A) default, representing exclusion/adequacy of competing explanations when challenged.

**Examples and controls.** The independently checked book substitution preserves hypothesis/event roles. Missing-premise variants and the shared-CQ cases exercise the default. Approval is for this declared qualitative interpretation.

**Recorded mapping:** CQ1: premise `hypothesis_event(H,A)`; CQ2: premise `event_true(A)`; CQ3: assumption `sufficient_correlation(H,A)`.

### 32.2 — `evidence_hypothesis_falsification`: approve

**Source:** printed p./pp. 331, 332; PDF p./pp. 341, 342. **Stage:** transcription-checked → example-passed.

H and event A remain in the same order as verification. An observed-false event supports explicit ¬H, matching 32.2.

**Critical questions.** The book prints the shared alternative-explanation question after both variants. Keeping sufficient_correlation(H,A) in the same order is an explicit shared-CQ interpretation; the textual question is awkward for falsification and is not silently rewritten as an erratum.

**Examples and controls.** The book example, adverse correlation case and opposite hypothesis conclusion preserve polarity and roles. The static hand-written instance also fixes the expected negated conclusion.

**Recorded mapping:** CQ1: premise `hypothesis_event(H,A)`; CQ2: premise `event_false(A)`; CQ3: assumption `sufficient_correlation(H,A)`.

### 33.1 — `positive_consequences`: qualified

**Source:** printed p./pp. 332, 333; PDF p./pp. 342, 343. **Stage:** transcription-checked → transcription-checked.

The inherited action/outcome form supports should_be_performed(A), corresponding to 33.1. This review adds a book-supported likelihood-adequacy default; the previous missing-CQ status was not forced by lack of the 1995 source.

**Critical questions.** CQ1 is now consequence_likelihood_adequate(A,G), meaning adequate assessment for the claimed argument strength, not necessarily strongly likely. CQ2 remains the supported consequence premise. CQ3 is a rebuttal for ¬should_be_performed(A).

**Examples and controls.** Independent questioning, explicit recovery, wrong-action and wrong-outcome controls pass. Existing positive/negative consequence examples exercise the common negation issue.

**Recorded mapping:** CQ1: assumption `consequence_likelihood_adequate(A,G)`; CQ2: premise `brings_about(A,G)`; CQ3: rebuttal `¬should_be_performed(A)`.

**Remaining condition.** The compendium-supported hook closes this operational gap. Preserve the original cited-source hold: the 1995 work was not independently read.

### 33.2 — `negative_consequences`: qualified

**Source:** printed p./pp. 332, 333; PDF p./pp. 342, 343. **Stage:** transcription-checked → transcription-checked.

The inherited action/outcome form supports ¬should_be_performed(A), corresponding to 33.2. The new likelihood default repairs the available-book omission without claiming to have seen the older cited book.

**Critical questions.** CQ1 is the same action/outcome adequacy assessment used by 33.1. CQ2 remains brings_about(A,G), and CQ3 permits an opposing positive-consequence argument in one issue. No fixed probability threshold is introduced.

**Examples and controls.** Questioning suppresses the default until supported; explicit support restores it. Questions about another action/outcome do not suspend this argument.

**Recorded mapping:** CQ1: assumption `consequence_likelihood_adequate(A,G)`; CQ2: premise `brings_about(A,G)`; CQ3: rebuttal `should_be_performed(A)`.

**Remaining condition.** The original cited 1995 work remains unverified; this is an explicit adaptation of the inherited schema from the supplied compendium.

### 33.3 — `negative_consequences_reasoning`: approve

**Source:** printed p./pp. 332, 333; PDF p./pp. 342, 343. **Stage:** transcription-checked → example-passed.

Agent P and action A are in the correct roles. The selected bring-about reading concludes ¬should_do(P,A), so a practical recommendation for the same agent/action is its opponent in the negation issue.

**Critical questions.** CQ1 is outcome_likely(A,B); CQ2 is the agent-conditioned consequence premise; CQ3 is a counterargument supporting should_do(P,A). The printed parenthetical alternative concerning not bringing A about is not claimed as a separately implemented variant.

**Examples and controls.** The worked practical-inference/opposing-negative-consequences graph and phase-5 paired-scheme examples establish actual shared-statement opposition. Different agents/actions must remain distinct. Numerical likelihood and proof-standard policy remain application responsibilities.

**Recorded mapping:** CQ1: assumption `outcome_likely(A,B)`; CQ2: premise `agent_bring_about(P,A,B)`; CQ3: rebuttal `should_do(P,A)`.

### 33.4 — `negative_consequences_prudential`: approve

**Source:** printed p./pp. 332, 333; PDF p./pp. 342, 343. **Stage:** transcription-checked → example-passed.

The contemplated omission of X and its bad consequence yield reconsideration and a prudential recommendation to do X. This conclusion is intentionally more specific than an unconditional practical ought.

**Critical questions.** Likelihood and evidential CQ roles are retained. other_consequences(X) is a broad defeater for overlooked consequences, not a separately weighted opposite conclusion; this is an explicit conservative policy.

**Examples and controls.** The book example and exception variant are adequate for the printed single respondent. Use a concrete action/case context or a separate graph when multiple respondents could be confused.

**Recorded mapping:** CQ1: assumption `strongly_occur(X,C)`; CQ2: premise `bad_consequence(X,C)`; CQ3: exception `other_consequences(X)`.

### 34 — `pragmatic_alternative`: approve

**Source:** printed p./pp. 333; PDF p./pp. 343. **Stage:** transcription-checked → example-passed.

The disjunctive practical premise and bad outcome yield the respondent-directed ought. Shared ought_to_bring_about has the same anonymous respondent reading as the threat/fear/danger family.

**Critical questions.** The compendium prints no separate CQs here. The two ordinary premises must be supported; no extra question is asserted to be printed.

**Examples and controls.** The book example and missing-premise variants match both source premises and conclusion. Approval assumes one respondent context, as the source does.

**Recorded mapping:** no printed critical-question list for this entry.

### 35 — `threat`: approve

**Source:** printed p./pp. 333; PDF p./pp. 343. **Stage:** transcription-checked → example-passed.

The three source roles remain distinct: adverse consequence of the respondent action, the proponent being able to produce the consequence, and the expressed intention. The conclusion opposes bringing about the action.

**Critical questions.** No separate CQs are printed for 35. The representation is a schema of threat argumentation and does not endorse the threat or establish a general moral prohibition.

**Examples and controls.** The example requires all three premises. The polarity fixture protects ¬ought_to_bring_about; ordinary danger and practical alternative interactions keep the same respondent/action context.

**Recorded mapping:** no printed critical-question list for this entry.

### 35.1 — `ad_baculum`: approve

**Source:** printed p./pp. 333; PDF p./pp. 343. **Stage:** transcription-checked → example-passed.

The disjunctive threat variant is distinct from entry 35: failure to do the requested action triggers the threatened bad outcome, and its conclusion recommends that action.

**Critical questions.** No separate CQs are printed. The threatening conditional is one ordinary premise, not an inferred causal fact.

**Examples and controls.** The book substitution and missing-premise variants distinguish this positive ought from entry 35’s negative ought. One dialogue respondent/proponent context is required.

**Recorded mapping:** no printed critical-question list for this entry.

### 36 — `fear`: approve

**Source:** printed p./pp. 333, 334; PDF p./pp. 343, 344. **Stage:** transcription-checked → example-passed.

The omitted action leads to a bad event; prevention is recommended; the selected action is the only way to prevent it. The conclusion follows the source direction.

**Critical questions.** No separate CQs are printed. The prevention recommendation and sole-means relation remain ordinary premises; they are not inferred solely from fear.

**Examples and controls.** All four printed premises are represented in the worked case and omission variants. The scheme shares the anonymous respondent practical conclusion under the explicit single-context contract.

**Recorded mapping:** no printed critical-question list for this entry.

### 37 — `danger`: approve

**Source:** printed p./pp. 334; PDF p./pp. 344. **Stage:** transcription-checked → example-passed.

A dangerous consequence of doing X supports the on-balance recommendation not to do X. Explicit negation of ought_to_bring_about is the selected practical opposition policy.

**Critical questions.** No separate CQs are printed. The badness/danger and consequence premises remain distinct.

**Examples and controls.** The source-directed example and a dedicated polarity mutation protect the negative conclusion. The graph can still contain opposing reasons and resolve them under its proof standard.

**Recorded mapping:** no printed critical-question list for this entry.

### 38 — `help`: approve

**Source:** printed p./pp. 334; PDF p./pp. 344. **Stage:** transcription-checked → example-passed.

Beneficiary A, potential helper B and action X are correctly separated. The conclusion should_do(B,X) is about the helper, not the beneficiary.

**Critical questions.** CQ1 is whether the proposed concrete action really helps A. CQ2 and CQ3 restate the ability and acceptable-cost ordinary premises; duplicating them as defaults/exceptions was properly removed.

**Examples and controls.** The source substitution has distinct people and an explicit action. Missing-premise cases demonstrate the need for ability and acceptable cost. A concrete action term carries its intended helper/context when several acts have similar names.

**Recorded mapping:** CQ1: assumption `action_will_help(X,A)`; CQ2: premise `can_carry_out(B,X)`; CQ3: premise `not_too_costly(B,X)`.

### 39 — `distress`: approve

**Source:** printed p./pp. 334, 335; PDF p./pp. 344, 345. **Stage:** transcription-checked → example-passed.

The relief premise explicitly identifies the distressed person A, the relieving actor B and action X. The shared conclusion concerns B doing X.

**Critical questions.** CQ1 is the distress premise; CQ2 is the relief relation; CQ3 is a challengeable possibility assumption; CQ4 is the actor/action side-effect exception.

**Examples and controls.** The book instance and phase-5 scoped cases distinguish beneficiary from actor, test unsupported possibility, and test side-effect defeat. These repairs are source-supported.

**Recorded mapping:** CQ1: premise `distress(A)`; CQ2: premise `relieve(B,X,A)`; CQ3: assumption `possible_bring_about(B,X)`; CQ4: exception `great_side_effects(B,X)`.

### 40a — `commitment1`: approve

**Source:** printed p./pp. 335; PDF p./pp. 345. **Stage:** transcription-checked → example-passed.

Evidence of commitment to X and the commitment link to Y support commitment to Y. It does not assert that Y itself is true.

**Critical questions.** CQ1’s positive part restates the commitment premise; contrary evidence is an exception. CQ2’s linkage exception names the same person, X and Y.

**Examples and controls.** The book case and adverse linkage/contrary-evidence cases preserve participant and proposition roles. Different links can be challenged independently.

**Recorded mapping:** CQ1: exception `commitment_contrary_evidence(A,X)`; CQ2: exception `commitment_linkage_exception(A,X,Y)`.

### 40b — `commitment2`: approve

**Source:** printed p./pp. 335; PDF p./pp. 345. **Stage:** transcription-checked → example-passed.

The second version preserves its dialogue-specific conditional: a prior commitment to X yields a presumptive commitment when Y becomes an issue. The distinct conclusion was correctly retained.

**Critical questions.** The commitment premise and contrary-evidence exception handle CQ1. CQ2 is the same person/X/Y linkage exception as version 1.

**Examples and controls.** The source fixture covers the dialogue-specific antecedent rather than merely copying version 1. Omission and exception variants establish the declared qualitative interpretation.

**Recorded mapping:** CQ1: exception `commitment_contrary_evidence(A,X)`; CQ2: exception `commitment_linkage_exception(A,X,Y)`.

### 41 — `ethotic1`: qualified

**Source:** printed p./pp. 336; PDF p./pp. 346. **Stage:** transcription-checked → transcription-checked.

The inherited positive-ethos record concludes S itself, whereas the compendium describes increased plausibility/acceptability. Defeasible support for S can be a declared evidential policy, but it is not a literal transcription of the graded conclusion.

**Critical questions.** Moral character is an ordinary premise, relevance is a default, and claimed weight belongs to the issue’s evidential policy. The default does not provide a numerical weight calibration.

**Examples and controls.** The book fixture proves the inherited strong output S, not agreement with the graded conclusion in the printed source.

**Recorded mapping:** CQ1: premise `good_moral_character(P)`; CQ2: assumption `character_is_relevant(S)`; CQ3: standard.

**Remaining condition.** Retain the declared-strength and unavailable-1995-source qualifications. A stable claim of exact compendium coverage needs a distinct graded formulation or an explicitly accepted argument-weight policy.

### 41 — `ethotic2`: block

**Source:** printed p./pp. 336; PDF p./pp. 346. **Stage:** transcription-checked → transcription-checked.

The inherited negative-ethos record concludes explicit ¬S. Reduced plausibility or rejection of a speaker’s assertion does not in general establish the contradictory proposition. This is a substantive strength/polarity difference from the supplied compendium.

**Critical questions.** Relevance and issue proof standards cannot by themselves make the inference from discrediting testimony to ¬S source-equivalent. The source asks whether the claimed weight is warranted.

**Examples and controls.** The existing fixture expects ¬bridge_safe, so its success confirms the inherited choice rather than independently validating it against the source.

**Recorded mapping:** CQ1: premise `bad_moral_character(P)`; CQ2: assumption `character_is_relevant(S)`; CQ3: standard.

**Remaining condition.** Do not approve as exact entry 41. Retain as an explicitly qualified legacy variant outside a fully source-certified release, or replace/add a scoped reduced-plausibility formulation and cases. The cited 1995 work remains unverified.

### 42 — `ad_hominem`: approve

**Source:** printed p./pp. 336; PDF p./pp. 346. **Stage:** transcription-checked → example-passed.

The source’s stated generic conclusion is strong rejection of the argument. The record identifies the attacked person, argument and conclusion; it does not infer the negation of that conclusion.

**Critical questions.** CQ1 restates the character-attack premise (also discussed by the book at p. 374). CQ2 is argument-relevant character. CQ3 is implemented conservatively by an exception for stronger supporting evidence. This is a selected strong-rejection policy, not numerical downweighting.

**Examples and controls.** Distinct-argument relevance and stronger-evidence tests support this declared policy. The old ledger note leaving strength to phase 5.3 is replaced with this explicit decision.

**Recorded mapping:** CQ1: premise `bad_character(A)`; CQ2: assumption `character_relevant(S)`; CQ3: exception `stronger_than_character(S)`.

### 43 — `pragmatic_inconsistency`: approve

**Source:** printed p./pp. 336, 337; PDF p./pp. 346, 347. **Stage:** transcription-checked → example-passed.

The advocate, argument and proposition are explicit. The contrary action is represented with the proposition’s explicit negation, and the conclusion rejects the argument rather than declaring its conclusion false.

**Critical questions.** CQ1 is strong advocacy/commitment; CQ2 is the contrary-action ordinary premise; CQ3 is relevance of the inconsistency to this person and argument.

**Examples and controls.** The book and adverse-question cases preserve person/argument/proposition roles. The removed duplicate-description default was correctly redundant.

**Recorded mapping:** CQ1: assumption `strong_advocate(A,S,X)`; CQ2: premise `opposite_action(A,S,¬X)`; CQ3: assumption `inconsistency_relevant(A,S)`.

### 44 — `inconsistent_commitment`: approve

**Source:** printed p./pp. 337; PDF p./pp. 347. **Stage:** transcription-checked → example-passed.

Evidence of commitment and opposed commitment supports inconsistency of the person’s commitments. The record no longer defeats itself merely because its second premise is present.

**Critical questions.** CQ1 and CQ2 test the two evidence premises. CQ3 is the challengeable commitments_conflict(A,S) relation, so apparent opposition alone need not suffice once challenged.

**Examples and controls.** The baseline book case, conflict-question case and static instance distinguish evidence for a conflict from a self-undercutting exception.

**Recorded mapping:** CQ1: premise `initial_commitment(A,S)`; CQ2: premise `opposed_commitment(A,S)`; CQ3: assumption `commitments_conflict(A,S)`.

### 45 — `circumstantial_ad_hominem`: qualified

**Source:** printed p./pp. 337, 338; PDF p./pp. 347, 348. **Stage:** transcription-checked → transcription-checked.

The practical commitment correctly opposes the argument’s conclusion C, and the final claim is weaker decreased plausibility of argument S. However, credibility_questioned(A) and character_issue(A) omit S.

**Critical questions.** CQ1 and the explained-inconsistency exception are correctly proposition/argument-scoped. CQ3 is still person-only; one argument’s dependence on credibility can be shared unintentionally with another argument by that person. CQ4 selects the weaker conclusion.

**Examples and controls.** Existing cases establish the selected weak output and explanation defeater for one target argument, not same-person multi-argument isolation.

**Recorded mapping:** CQ1: premise `committed_opposite(A,C)`; CQ2: exception `inconsistency_explained(A,S)`; CQ3: assumption `character_issue(A)`; CQ4: standard.

**Remaining condition.** Before unrestricted use, bind character_issue and credibility_questioned to (A,S), with adverse wrong-argument controls and migration, or enforce one attacked argument per person/context.

### 46 — `bias`: approve

**Source:** printed p./pp. 338; PDF p./pp. 348. **Stage:** transcription-checked → example-passed.

The corrected conclusion is a comparative likelihood of incomplete consideration, not a categorical claim that the argument is false or that no contrary evidence was considered.

**Critical questions.** CQ1 is relevance of the bias to this dialogue/conclusion; CQ2 is the ordinary bias premise. The same person is used in the general premise and application.

**Examples and controls.** The source fixture and distinct-conclusion relevance cases support the scoped comparative reading. Dialogue context is supplied through the case’s conclusion/context terms.

**Recorded mapping:** CQ1: assumption `bias_relevant(A,C)`; CQ2: premise `biased(A)`.

### 47 — `bias_ad_hominem`: approve

**Source:** printed p./pp. 338, 339; PDF p./pp. 348, 349. **Stage:** transcription-checked → example-passed.

The printed bad-bias/honest-participation sequence ends in reduced credibility for argument S. The record retains the printed morally-bad premise as an explicit premise rather than deriving it from any bias whatsoever.

**Critical questions.** The evidence of bias and its inappropriate kind are ordinary premises. No extra printed CQs are invented, and normal permissible bias does not satisfy the inappropriate-bias premise by definition.

**Examples and controls.** The book fixture and missing-premise variants validate the weak credibility conclusion. The scheme does not automatically change another argument’s numeric weight or conclude ¬S.

**Recorded mapping:** CQ1: premise `proponent_biased(A,S)`; CQ2: premise `biased_not_take_part(A,D,S)`.

### 48 — `gradualism`: approve

**Source:** printed p./pp. 339; PDF p./pp. 349. **Stage:** transcription-checked → example-passed.

The initial proposition, an intervening implication sequence, and acceptance of the endpoint conditional yield acceptance of the final proposition. The sequence predicate is an aggregate witness with endpoints, not an algorithm that searches every chain.

**Critical questions.** No CQs are printed. proper_sequence(P1,P4) is transparently marked n=0 as an additional executable-order assumption.

**Examples and controls.** The endpoint-scoped case and omission variants show how this chosen representation applies. Approval requires the application’s sequence witness, as documented; it does not assert that four variable slots limit every real argument to four steps.

**Recorded mapping:** CQ0: assumption `proper_sequence(P1,P4)`.

### 49 — `slippery_slope_base_case`: qualified

**Source:** printed p./pp. 339, 340; PDF p./pp. 349, 350. **Stage:** transcription-checked → transcription-checked.

This is one part of the inherited decomposition. It concludes negative consequences of an action, whereas entry 49’s full argument concludes that the initial action should not be brought about.

**Critical questions.** The action/event link carries the chain-link CQs in this decomposition. No separate first-step presumption or normative bridge appears in this record.

**Examples and controls.** Its book example and the inductive-chain fixture establish horrible-cost propagation to negative consequences. They do not establish the full printed normative conclusion.

**Recorded mapping:** CQ1: premise `would_realize(A,E)`; CQ2: premise `would_realize(A,E)`; CQ3: premise `would_realize(A,E)`.

**Remaining condition.** Preserve as a useful legacy component with original-source hold. Do not count the pair alone as exact completed coverage of entry 49’s full scheme.

### 49 — `slippery_slope_inductive_step`: qualified

**Source:** printed p./pp. 339, 340; PDF p./pp. 349, 350. **Stage:** transcription-checked → transcription-checked.

The inherited recursive rule propagates horrible costs backwards along a causal link. It is a component of a slope construction, not the whole printed scheme.

**Critical questions.** The link-strength, missing-step and weakest-link questions can be raised by contesting individual supplied links. The engine does not search for omitted intermediate events or calculate probabilities.

**Examples and controls.** The chain example shows composition with the base case. It ends in negative consequences, and a normative conclusion still requires an additional source-supported bridge.

**Recorded mapping:** CQ1: premise `causes(E1,E2)`; CQ2: premise `causes(E1,E2)`; CQ3: premise `causes(E1,E2)`.

**Remaining condition.** Retain inherited-component and unavailable-original-source qualifications, together with the entry-49 coverage limitation.

### 50 — `precedent_slippery_slope`: approve

**Source:** printed p./pp. 340; PDF p./pp. 350. **Stage:** transcription-checked → example-passed.

The source genuinely prints the initial case C0 in the intolerable-outcome premise and CQ3, even after describing a sequence ending at Cn. The existing C/R intolerability relation is faithful to that print.

**Critical questions.** CQ1–CQ2 test the precedent-series premise; CQ3 tests the stated intolerable exception premise. No intuitive C0→Cn correction was made.

**Examples and controls.** The book fixture and premise-removal variants are adequate for the aggregate sequence representation. Sequence evidence must justify case-to-case consistency; the library does not infer it from superficial similarity.

**Recorded mapping:** CQ1: premise `precedent_series(C,R,N)`; CQ2: premise `precedent_series(C,R,N)`; CQ3: premise `intolerable_exception(C,R)`.

### 51 — `sorites_slippery_slope`: approve

**Source:** printed p./pp. 340, 341; PDF p./pp. 350, 351. **Stage:** transcription-checked → example-passed.

The clearly held starting property, successor conditional and reapplication sequence support the possibility-like endpoint claim may_have(A,P). Property P is now carried throughout the sequence.

**Critical questions.** No separate CQs are printed. The whole chain is an ordinary aggregate premise whose witness must preserve property and endpoints.

**Examples and controls.** The source fixture plus the shared conditional with verbal slope demonstrate role-preserving composition. The endpoint is not promoted to an unconditional categorical has_property claim.

**Recorded mapping:** no printed critical-question list for this entry.

### 52 — `verbal_slippery_slope`: approve

**Source:** printed p./pp. 341; PDF p./pp. 351. **Stage:** transcription-checked → example-passed.

The same property P is preserved across the original individual, neighbor and generalized endpoint. The conclusion says one cannot truly assert the endpoint has P; it is distinct from ordinary ¬has_property.

**Critical questions.** No separate CQs are printed. Vague-neighbor and extreme-case claims are ordinary premises; an application must justify the intervening sequence and vagueness warrant.

**Examples and controls.** The book case and cross-scheme conditional fixture preserve the narrower speech/justification conclusion and property identity. The source’s vagueness discussion is not replaced with binary extensional negation.

**Recorded mapping:** no printed critical-question list for this entry.

### 53 — `full_slippery_slope`: approve

**Source:** printed p./pp. 341; PDF p./pp. 351. **Stage:** transcription-checked → example-passed.

The current aggregate formula follows the printed 53: initial presumption, linked sequence and social climate accepting successive steps, concluding initial nonacceptability. The printed display does not separately state a bad endpoint; this review does not silently add one.

**Critical questions.** CQ1–CQ2 concern strength and required proof standard, which remain application-level choices. CQ3 is the social-climate ordinary premise; duplicating it as an assumption was properly removed.

**Examples and controls.** The book example and separate-series cases establish scope of accept_all_steps(C0). Approval is for the displayed qualitative schema and declared sequence witness, not a universal claim that any accepted sequence is bad.

**Recorded mapping:** CQ1: standard; CQ2: standard; CQ3: premise `accept_all_steps(C0)`.

### 54.1 — `rule_physical_world_1`: approve

**Source:** printed p./pp. 342; PDF p./pp. 352. **Stage:** transcription-checked → example-passed.

Repaired. The warrant explicitly requires referential coextension. The old whale/mammal fixture gave only class inclusion and nevertheless produced synonymy. The additional premise makes the warrant’s missing antecedent reviewable.

**Critical questions.** The nonnecessary-truth rebuttal factor remains an undercutter scoped to W/D. The universal synonymy warrant is still explicitly supplied; the review does not assert it as an independently established linguistic theory.

**Examples and controls.** Triangle/three-sided-polygon with coextension is a positive witness. Whale/mammal inclusion alone and coextension of a different pair produce no synonymy argument. Baseline CLI evidence independently reproduces the old false oracle.

**Recorded mapping:** CQ1: exception `not_necessarily_true(W,D)`.

### 54.2 — `rule_physical_world_2`: qualified

**Source:** printed p./pp. 342; PDF p./pp. 352. **Stage:** transcription-checked → transcription-checked.

The institutional rulebook, warrant and changed-rule defeater match 54.2. The conclusion counts_as(D,W) omits institution I.

**Critical questions.** The only printed rebuttal factor is correctly scoped to I/D/W. The rulebook-reporting warrant remains global as a universal warrant, not a claim that every particular rulebook is current.

**Examples and controls.** The source and two-rules cases cover rule-change isolation. They do not separate two institutions with incompatible counts-as relations using the same D/W terms.

**Recorded mapping:** CQ1: exception `rule_changed(I,D,W)`.

**Remaining condition.** Use one institution or institution-qualified ground terms. A generally multi-institution release should carry I in the counts-as statement and corresponding consumers.

### 54.3 — `rule_mental_world`: qualified

**Source:** printed p./pp. 342; PDF p./pp. 352. **Stage:** transcription-checked → transcription-checked.

The source language L is present in the speakers’ premise but omitted from means_same(W,D). The counterexample without_correcting(W) is also broader than the particular language/definition.

**Critical questions.** The printed rebuttal concerns competent usage without correct understanding of the referent. The current W-only hook captures a single lexical context, not arbitrary languages or competing definitions.

**Examples and controls.** The book example and exception variant are adequate for one language and one definition of W.

**Recorded mapping:** CQ1: exception `without_correcting(W)`.

**Remaining condition.** Retain a single-language/definition context qualification or make language and definition explicit before unrestricted composition across languages.

### 55.1 — `established_rule`: qualified

**Source:** printed p./pp. 343; PDF p./pp. 353. **Stage:** transcription-checked → transcription-checked.

The inherited schema remains rule-and-conclusion application, rather than 55.1’s actor/action form. This review adds the available-book conflicting/overriding-rule and exceptional/excusing-case hooks.

**Critical questions.** CQ1 remains the asserted rule content; CQ2 and CQ3 now name R and its particular conclusion/application C. The valid(R) assumption remains explicitly marked as an additional n=0 policy.

**Examples and controls.** Both new exceptions defeat the same R/C application. Wrong-rule and wrong-conclusion facts do not. Positive and missing-premise cases remain valid.

**Recorded mapping:** CQ1: premise `has_conclusion(R,C)`; CQ2: exception `conflicting_or_overriding_rule(R,C)`; CQ3: exception `rule_exception_or_excuse(R,C)`; CQ0: assumption `valid(R)`.

**Remaining condition.** The two previously unmodeled book questions are operationally closed. Original 1995 source verification and exact actor/action compendium-form coverage remain qualified.

### 55.2 — `rule`: approve

**Source:** printed p./pp. 343; PDF p./pp. 353. **Stage:** transcription-checked → example-passed.

Entry 55.2 prints alternative evaluation/conduct wording. The combined evaluation_justified(E,C) atom is read as the selected evaluative/conduct judgement described by the input, not as two independent propositions that are both automatically established.

**Critical questions.** No separate CQs are printed. Both the rule conditional and particular case are ordinary premises.

**Examples and controls.** The book substitution and omission variants are adequate for the selected combined reading. Applications requiring separate evaluative and conduct statements should select and document that interpretation explicitly.

**Recorded mapping:** no printed critical-question list for this entry.

### 55.3 — `regulative_rule`: approve

**Source:** printed p./pp. 343; PDF p./pp. 353. **Stage:** transcription-checked → example-passed.

The source warrant supports obligatory action → must do. The printed parenthetical pairing reverses obligation/prohibition in the displayed conclusion; the current variant follows the warrant, with addressee scope explicit.

**Critical questions.** The source rebuttal factor is an adequate excuse or overriding duty, correctly actor/action-scoped.

**Examples and controls.** Positive, exception and polarity fixtures protect must_do(A,X). This is a supported interpretation of the supplied book, not a publisher-confirmed erratum or verification of Grennan’s original book.

**Recorded mapping:** CQ1: exception `excuse_or_overriding_duty(A,X)`.

### 55.3 — `regulative_rule_prohibition`: approve

**Source:** printed p./pp. 343; PDF p./pp. 353. **Stage:** transcription-checked → example-passed.

The prohibition variant concludes must_not_do(A,X), a duty not to act. It does not merely deny a duty to act, and its addressee premise prevents introducing an unbound actor.

**Critical questions.** The same actor/action excuse-or-overriding-duty factor is appropriate for the printed combined warrant.

**Examples and controls.** The book case and exception variant retain prohibition polarity. A simultaneous positive obligation and prohibition needs an application issue/policy; the generic explicit-negation issue does not automatically identify these two deontic predicates.

**Recorded mapping:** CQ1: exception `excuse_or_overriding_duty(A,X)`.

### 56 — `exceptional`: approve

**Source:** printed p./pp. 344; PDF p./pp. 354. **Stage:** transcription-checked → example-passed.

Repaired interpretation. The three questions are conditional: applicable recognition or evidence for an unrecognized case, with comparable-case support additionally required when borderline. They are not a flat OR.

**Critical questions.** The aggregate exception_backed(C) remains a defeasible default. After challenge, the witness contract requires an explicit recognition status, borderline status and the applicable evidence/comparison route. Case C includes its governing rule/context.

**Examples and controls.** Four positive witness routes and four negative controls pass: recognition cannot bypass a borderline comparison, comparisons cannot replace required non-applicability evidence, wrong-case comparisons fail, and unknown status is not treated as clear.

**Recorded mapping:** CQ1: assumption `exception_backed(A)`; CQ2: assumption `exception_backed(A)`; CQ3: assumption `exception_backed(A)`.

### 57 — `from_precedent`: approve

**Source:** printed p./pp. 344; PDF p./pp. 354. **Stage:** transcription-checked → example-passed.

The legitimate counterexample has property F and explicit ¬G under the stated R/F/G general rule. The conclusion recognizes an exception for R/A and thereby calls for qualification; it does not silently rewrite the rule store.

**Critical questions.** CQ1 is rule inapplicability to the case; CQ2 is legitimacy of the counterexample; CQ3 is an already recognized exception requiring no rule change. Each carries the relevant R/A scope.

**Examples and controls.** The book case, negation mutation, questioning and exception variants support the displayed interpretation. The conclusion is appropriately distinct from entry 50’s refusal of exceptional status.

**Recorded mapping:** CQ1: exception `rule_inapplicable_to(R,A)`; CQ2: assumption `legitimate_case(A,R)`; CQ3: exception `recognized_exception_category(A,R)`.

### 58 — `plea_for_excuse`: qualified

**Source:** printed p./pp. 345; PDF p./pp. 355. **Stage:** transcription-checked → transcription-checked.

The special circumstances now identify the committed action. However the pleader remains implicit and excuse(E)/not_excuse(E) omit the governing rule, action and sanction application.

**Critical questions.** CQ1–CQ2 are a recognized or specially justified excuse aggregate. CQ3 is a future-precedent problem for an unrecognized excuse. The aggregate must encode this conditional meaning for the particular plea.

**Examples and controls.** The book and adverse cases establish one pleader, one rule/action/sanction context. They do not prove isolation of two pleas reusing the same emergency or sanction term.

**Recorded mapping:** CQ1: assumption `excuse(E)`; CQ2: assumption `excuse(E)`; CQ3: exception `not_excuse(E)`.

**Remaining condition.** Use case-specific excuse and sanction terms, or introduce explicit pleader/case scope before approving unrestricted multi-plea graphs.

### 59.1 — `perception`: qualified

**Source:** printed p./pp. 345; PDF p./pp. 355. **Stage:** transcription-checked → transcription-checked.

The perceiver P and perceptible property I follow 59.1, with a prima-facie evidential warrant. The defeater not_indicator(I) omits P and perceptual context.

**Critical questions.** The source undercutter asks about the circumstances in which the image is unreliable. A generic property name alone cannot distinguish one observer’s misleading circumstance from another observer’s reliable one.

**Examples and controls.** The book example and exception variant validate a single perceptual context, not independent observers of the same property.

**Recorded mapping:** CQ1: exception `not_indicator(I)`.

**Remaining condition.** Use a context-qualified I or separate graphs; for general multi-observer use, scope the unreliability defeater to observer/context and add independent-support controls like those now used for memory.

### 59.2 — `appearance`: approve

**Source:** printed p./pp. 345, 346; PDF p./pp. 355, 356. **Stage:** transcription-checked → example-passed.

The visible classification premise, defeasible category conclusion, and both critical questions match pp. 345–346. The root reviewer independently checked the original 2006 journal PDF pp. 324–325; this reviewer separately checked the supplied compendium.

**Critical questions.** misleading_appearance(O,C) undercuts misleading appearance in the proposed classification. more_justifiable_category(O,C) is a bounded existential/comparative defeater: evidence must show a better alternative category for this O/C claim, not merely some other category.

**Examples and controls.** The source fixture and exception variants cover both questions. No unbound alternative-category variable is introduced. Under this declared aggregate interpretation and the independent journal verification, the original-source hold can close.

**Recorded mapping:** CQ1: exception `misleading_appearance(O,C)`; CQ2: exception `more_justifiable_category(O,C)`.

### 60 — `memory`: approve

**Source:** printed p./pp. 346; PDF p./pp. 356. **Stage:** transcription-checked → example-passed.

Repaired. The belief-history undercutters identify recaller P and remembered proposition M. A false basis or absence of an original belief belonging to Ann does not defeat Bob’s independent recall of M.

**Critical questions.** All three printed undercutters are modeled as exceptions. CQ1–CQ2 now have P/M scope; CQ3 already had it. The source’s unusual wording of CQ2 is preserved rather than silently reinterpreted.

**Examples and controls.** Independent same-recaller and other-recaller cases pass for both history questions. One attacked support leaves the other support usable; defeating both removes both. The original baseline demonstrates the unary cross-recaller defect.

**Recorded mapping:** CQ1: exception `based_on_false_belief(P,M)`; CQ2: exception `not_originally_believed(P,M)`; CQ3: exception `recaller_doubts(P,M)`.

## Deliverables and reproducibility

`per-record-audit.json` is the machine-readable complete ledger, including source pages, dispositions, full current scheme records, question mappings and named case inputs/expectations. `per-record-audit.csv` supplies the same adjudication text in a compact comparison format. Baseline probes, before/after test logs and the patch are supplied alongside this report. The patch changes no engine source and publishes nothing. The source PDF and rendered copyrighted pages are review inputs and are not required in the deliverable archive.
