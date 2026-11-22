# Independent Atkinson AS1/AS2 review and closure

Reviewed on 2026-11-22 UTC against repository base commit `98a2335d2c9909cf597ae85a9f359f20d092ae6c` and the corrective working-tree changes described below.

**Reviewer:** ChatGPT, a separate model-assisted reviewer from Claude, the author of the five records and original adapter. This is independent review of those existing records, not an authenticated human endorsement. The adapter corrections made during this audit are authored by this reviewer and need a separate integration/code review; source approval of unchanged records does not certify those new patches independently.

## Disposition

**Accept the five publication-origin scheme records as explicitly declared interpretations of Atkinson and Bench-Capon (2007), with the scope below. Do not accept the unqualified claims that the adapter previously handled every valid AATS model, that every figure/text variation preserves the reported choices, or that AS1 inputs can directly execute the comparison contract.**

Several implementation and input-integrity defects and one material factual overstatement were found and corrected, together with a minor display-order inconsistency. The source interpretations themselves did not require changes to `walton.yml` or the five fingerprinted ledger contracts. The corrections preserve all 17 paper situations and 32 recorded choice expectations.

The accepted framework profile is:

1. AS1 and two negative variants, with explicit mappings for all 17 critical questions and a declared Carneades burden policy.
2. AS2 and its negative variant over an explicitly supplied AATS model, with a goal-labelled refinement of the paper's value effects and a conclusion projected to the deliberating agent's individual action.
3. An external, one-step, per-value worst-case comparison policy, using unit goal counts and lexicographic value order. This policy agrees with the declared reconstruction of the example choices. It is not an algorithm uniquely prescribed by the paper, a general VAF evaluator, or a planner.
4. Application of that comparison policy to AS1 requires an additional AATS model and an explicit correspondence to the AS1 argument. AS1 premises alone do not determine the required joint actions, possible outcomes, assurances or counting of goals.

This closes the source-review decision for these five records within that profile. It does not establish general real-world validity, calibrated preferences, correctness of an externally authored model, or correctness of all future engine inputs.

## Evidence examined

The primary source was the owner-supplied 20-page PDF, *Practical reasoning as presumptive argumentation using action based alternating transition systems*, Artificial Intelligence 171 (2007), pp. 855–874. Its SHA-256 is `dad4124c052c48b5c82de02494aac6971de113bb8452ae9fc6e764aa35e43cbc`.

Printed pp. 858–871 were rendered and visually inspected. In particular, the conjunction and inequality symbols in the formal CQ definitions, the state subscripts, the plus/minus annotations in Figs. 1–8, and the train transition table were read from images. Extracted text was used for navigation, not as the sole basis for mathematical transcription.

Repository evidence included all five live scheme records and their language templates; their complete source-ledger question mappings; `tools/aats.py`; both YAML models; `tools/test_aats.py`; the handwritten and generated practical-reasoning cases; contract 8; the usage guide; `final-closure.md`; and `atkinson-2007-example.md`. The supplied handoff was treated as a list of claims to check, not as an authority replacing the source.

Machine-readable findings, exact source-acceptance recommendations, counterexamples, rendered evidence and test logs are supplied in the review package's `reviews/atkinson/` directory.

## Findings and repairs

### AAT-01 — Different effects on the same state pair produced extra engine arguments

**Severity: high. Confirmed; repaired.**

The old validator accepted two joint actions `ja` and `jb` with the same transition `s0 → s1` while only `ja` promoted `benefit` through `goal`. `Situation.arguments()` predicted only:

`practical_reasoning_as2(alice,act_a,ja,s0,s1,goal,benefit)`.

The exported effect fact was `transition_promotes(s0,s1,goal,benefit)`, which intentionally contains no joint-action parameter. Together with both transition facts, the actual engine correctly constructed an additional argument for `act_b` and labelled both actions in. The adapter's choice remained only `act_a`. This was an inconsistent adapter/model contract, not evidence that the engine's premise join was wrong.

The counterexample, the pre-fix prediction and the actual evaluated graph are retained as `same_edge_different_effects.yml`, `effect_alias.expected.json`, `effect_alias.graph.yml` and `effect_alias.actual.yml`.

The source's δ is a function of the source state, target state and value (p. 861). The project's refinement additionally attaches a goal, but still omits the joint action. The repair therefore rejects differing effect sets for one source/target pair. It does not silently merge them, change the scheme arity or alter the paper models. Authors needing action-dependent costs must represent those distinctions in the state/goal model or use a separately specified extension.

The earlier random tests did not expose this: their second derivation read the adapter's own predicted argument identifiers. It could reproduce the adapter's choice without discovering that the emitted facts allowed extra arguments. The new oracle independently joins emitted facts to reconstruct exact argument sets and choice scores without reading predicted identifiers. It runs on all 17 paper situations and 300 random models. A separate Go test sends adapter facts into the actual vendored engine and checks four explicitly authored expected arguments, including a goal becoming false. An inconsistent model must be rejected before graph generation.

### AAT-02 — The claim that all figure/text variations preserve choices was false

**Severity: medium; material to the validation claim. Confirmed; corrected and tested.**

Fig. 5 labels the transition from q11 back to q10 with one −P sign. The text and footnote 5 on p. 868 describe two harmful goals: moving the dog left and revisiting q10. A single sign can indicate direction without encoding magnitude, so this need not be called a false figure. It does mean that treating every printed sign as one unit is a different operational interpretation.

Under the implemented value order with P above F:

| Interpretation at q11 | Row chicken | Row dog | Selected action |
|---|---|---|---|
| Text: two harmful goals for rowing the dog | P = −1, F = 0 | P = −2, F = +1 | Row chicken |
| One unit for the single graphical −P sign | P = −1, F = 0 | P = −1, F = +1 | Row dog |

The previous blanket invariance assertion was therefore wrong. The existing invariance tests never made this Fig. 5 substitution. The model appropriately follows the textual two-goal reading, and the new regression deliberately proves that the alternative interpretation changes the choice. The analogous Fig. 7 substitution does preserve its reported choices; it now has its own check.

This finding does not invalidate the declared reconstruction. It rules out claiming that all graphical readings are equivalent or that the expected choices determine the policy uniquely. The original handoff cannot be accepted without this correction.

### AAT-03 — The validator incorrectly required distinct proposition sets for distinct states

**Severity: medium. Confirmed; repaired.**

Page 860 defines an interpretation function π from states to sets of propositions; it does not require π to be injective. The old implementation rejected two different state ids with identical proposition sets, even though they can have different transitions or future behaviour.

The repair permits those states. A state named by id remains unambiguous. When a claimed state is supplied only as a proposition list matching several ids, diagnostics now require an explicit id instead of arbitrarily choosing the first match. The source-based acceptance test covers both state ids and the ambiguous-list rejection. None of the paper models depends on the removed restriction.

### AAT-04 — An unnamed legal joint-action tuple was misclassified as CQ13

**Severity: medium. Confirmed; repaired.**

The paper defines the full joint-action set as the Cartesian product of the agents' action sets (p. 860), while τ is partial. The YAML format gives names only to the combinations used by the model. Those names cannot redefine membership in the Cartesian product.

Before repair, an explicit tuple in which the farmer waits, the chicken eats seeds and the dog waits was rejected as CQ13 merely because the tuple had no name in the farmer model. Each component is a declared action, so the tuple is legal. Its claimed transition is absent, which is CQ2 under the paper's distinction on p. 859.

Diagnostics now check membership from the component actions. An unnamed legal tuple without a transition yields CQ2. A genuinely unknown component, such as a bridge-crossing action absent from the action set, still yields CQ13. A named tuple supplied as a mapping is also tested. This changes diagnostics, not the existing paper graphs or their choice results.

### AAT-05 — AS1 comparison is conditional on an additional model

**Severity: medium documentation/claim boundary. Confirmed; made explicit.**

`tools/aats.py decide` reads a `walton-aats/v1` model. It does not read an AS1 graph, and no AS1-to-AATS translator is present. A ledger entry naming `value-based-choice` records an available external method, not proof that an arbitrary AS1 instance supplies its inputs.

Contract 8 and the usage guide now list the additional required model and mapping. This is a legitimate bounded interpretation of the ledger's question sorting: alternative actions are represented, and a conditional comparison method is identified. It would be inaccurate to call the bare AS1 records a complete deliberation procedure or to claim the engine automatically ranks their competing conclusions by an audience.

### AAT-06 — The comparison rule is an explicit project policy

**Disposition: accepted with bounded claims. No algorithmic change.**

The paper supplies qualitative value-order reasoning, worst-case concerns, assurances and some magnitude-sensitive examples. It does not specify the exact implemented tuple-valued maximin algorithm. In particular, unit counts for distinct goals, minimization separately for each value, lexicographic ordering and returning every tied action are project decisions.

Different values' minima may be attained by different joint outcomes. The resulting score tuple is a conservative per-value summary, not necessarily the valuation of one reachable state. Two synonyms encoded as distinct goals can increase a score twice. Neither fact makes the declared policy internally inconsistent, but both matter when building real-world models. Authors must justify goal decomposition and assurance choices; test agreement does not calibrate them empirically.

The revised contract states these choices and their limits. Accepting this reproducible policy as the project's initial profile is reasonable. Describing it as a general VAF semantics, a uniquely faithful implementation of every literal printed definition, or a proof of optimal planning would not be.

### AAT-07 — Duplicate YAML keys could silently replace model content

**Severity: medium input-integrity defect. Found during the root integration review; repaired.**

The prior use of `yaml.safe_load` silently retained the last value of a repeated mapping key. A repeated state id or `transitions` section could therefore change the effective model without reporting that the submitted input was ambiguous. The adapter now uses the repository's duplicate-key-aware safe loader and raises `ModelError` with the repeated key and source lines. This also wraps malformed YAML and common invalid model shapes in typed model errors, rather than silently splitting a scalar into characters or exposing an incidental Python exception.

The control checks intentional YAML merge-key inheritance with explicit overrides: the supplied farmer and train files remain valid. Negative tests cover duplicate top-level and nested state keys, malformed YAML, non-mapping agent data, a non-list situations collection, scalar proposition lists, non-mapping transitions/effects and scalar decision choices. These are checks of the supported input format, not a new general YAML parser.

### AAT-08 — Mixed delta display did not match its documented order

**Severity: low. Corrected and tested.**

For a transition with both signs, the prior reverse lexical sort returned `-+` while the docstring promised `+-`. Sign-membership tests and choice computations were unaffected. The display now returns `+-` in the documented order, with a direct Fig. 6 regression.

### AAT-09 — An unknown joint reference caused a cascading KeyError

**Severity: medium input-error handling. Found by the separate source reviewer during final cross-review; repaired.**

Changing a transition's `by` field to `unknown_joint` was detected by the validator, but its last option check called a method that assumed every joint reference was already valid. The resulting raw `KeyError` masked the intended model diagnostic. The validator now tests option existence using only known joint relations. The full model API raises `ModelError`; the CLI exits 2 with the appropriate unknown-joint diagnostic and no traceback. Both interfaces have an explicit regression.

### AAT-10 — An audience typo silently recommended every action

**Severity: medium; affects decision output. Found by the separate source reviewer during final cross-review; repaired.**

`--audience definitely_unknown` previously became a one-value order that was then filtered to no values. All action scores were empty and every option was returned as best, with success status. The adapter now validates named and ad-hoc orders against the model's declared value universe and rejects undeclared values, duplicate values and empty orders. A model with no named audiences requires an explicit order when using `decide`.

Valid partial orders remain supported. Values declared for another modeled agent also remain legitimate: the comparison contract still filters them to the deliberator's values, and an intentional order consisting only of such values can legitimately tie all actions. The repair distinguishes this documented case from a misspelling. Regressions cover the API, CLI exit 2 with no choice output, valid subset orders and the other-agent control.

### AAT-11 — Malformed claimed-argument shapes raised incidental Python errors

**Severity: low input-error handling. Found during the same cross-review; repaired.**

Inputs such as `claim_problems([])`, a list-valued joint action, or a nested-list state description produced `AttributeError` or `TypeError`. Claimed arguments now require the documented mapping fields and typed atoms, action tuples and proposition lists. These malformed shapes raise `ModelError`; ordinary source-level disagreements such as an unknown but well-formed action, goal or value still receive their critical-question classification.

## Adjudication of the paper-review claims

### The formal definitions

| Claim | Independent disposition |
|---|---|
| CQ1 is printed as wrong initial state **and** unavailable individual action | Correct; verified visually on p. 862. It is narrower than the informal truth-of-circumstances question and the p. 864 wording. The AS2 scheme separately requires the actual state and action precondition, so its acceptance condition is stricter than that printed conjunction. |
| CQ3 and CQ15 are printed for a proposition becoming true, although AS2 permits a proposition becoming false | Correct. The polarity reversal for false-making goals is an explicit, defensible repair. It is not a literal transcription of those two CQ formulas. |
| CQ6 prints qy instead of the alternative result qz and collapses to CQ5 for a valid AS2 instance | Correct. The rendered page has qy. Because AS2 already requires the goal proposition to change between qx and qy, the extra condition adds nothing. Reading qz restores the informal same-goal alternative and the example use. |
| Literal CQ8 requires δ to be both + and − for the same arguments | Correct. The p. 861 codomain gives a single sign. Goal-labelled effects are a declared extension motivated by Fig. 6 and the magnitude discussion; they are not the same mathematical δ as printed. |
| CQ5–CQ7 and CQ11 can overlap CQ17 | Correct as a formal observation: the printed condition distinguishes joint actions without requiring a different individual action. The comparator groups outcomes by individual action, so it does not treat another agent's choice as the deliberating agent's freely selectable action. |

These are source tensions or representation decisions, not a license to treat every plausible correction as uniquely authored intent. The documentation should preserve both the print and the adopted operational reading, as the current ledger does.

### All 17 critical-question mappings

| CQ | Source/example assessment | Accepted operational treatment |
|---|---|---|
| 1 | Named on p. 864; formal conjunction on p. 862 | Actual-state/precondition premises; separate printed-CQ diagnostics. |
| 2 | Wrong transition or consequence; p. 859 distinguishes it from others' choices | Transition premise and diagnostics, including the newly corrected unnamed-tuple boundary. |
| 3 | Arg10 repetition challenge, p. 867 | Changed proposition, with the declared false-making extension. |
| 4 | Obj2c, p. 865 | Promotion/demotion condition; source and project goal attribution kept distinct. |
| 5 | No explicitly named adverse instance in the worked examples | Same-target alternative; synthetic cases supply actual coverage. The printed transition diagrams/table do not contain the necessary duplicate target from one state. |
| 6 | Obj4b and q10, pp. 865, 868 | Same transportation goal. The `right_k` interpretation is defensible and disclosed; retaining the earlier accusation that Obj4b is mislabelled would be wrong. |
| 7 | Not explicitly named; overlaps same-value alternatives present in the example | Inferred same-value alternatives; not presented as an explicit paper label. |
| 8 | Named at q14, p. 868; Fig. 6 has +P and −P | Separate harmful-goal counterargument in the project refinement. |
| 9 | Many explicit harmful-side-effect examples | Negative variant/rebuttal. At q14 the row-alone remark is a negative argument, without a displayed positive AS2 instance to attack; calling this a loose label is reasonable. |
| 10 | Not explicitly named in the example, but multi-value actions occur; p. 862 says it endorses the action | Additional support, not an undercutter. |
| 11 | Obj2f, Obj4c and q10 explicitly named | Competing actions or the AS1 preclusion variant, with external comparison and its declared scope. |
| 12 | Muzzled dog in the initial state, p. 864 | Model-state membership or AS1 possibility default. |
| 13 | Bridge action, p. 864 | Action/joint-action membership, separated from a missing transition. |
| 14 | Muzzled dog in the consequent state, p. 864 | Result-state membership or AS1 possibility default. |
| 15 | Not explicitly named; applying the muzzled-dog remark to a goal is a constructed inference | Realizability condition; this supplies an additional test, not a quotation of a paper example. |
| 16 | Obj1c and Obj2e, p. 865 | Value legitimacy/agent value. The AS2 record's agent-specific value restriction is stated in its ledger. |
| 17 | Explicit farmer/trains uses and p. 859 explanation | Other-agent deviations; the project accepts promotion despite deviation when the value remains promoted, and treats a possible harmful deviation as a reason to refrain absent an assurance. This asymmetry is an interpretation supported by the narratives. |

The inventory of 13 explicitly named CQ numbers and four unnamed ones is correct for the worked examples. It does not mean the paper contains 17 independently labelled examples or that every question has been validated by an observed real-world case. CQ5 needs a constructed case; CQ15's muzzle-goal instance is an added interpretation.

Obj6ai's relabelling from CQ2 to CQ17 is justified within the paper's explicit AATS distinction: both transitions exist, and the uncertainty concerns the chicken's choice. At an AS1-only level, subsuming it under the transition claim remains the abstraction described on p. 859. The two statements are compatible once their levels are distinguished.

### Figure and text differences

| Item | Disposition |
|---|---|
| The text announces ten joint actions but lists j0–j10 | Correct counting discrepancy: eleven named actions. |
| The prose says right bank first while the transitions use the first list as the starting bank | Correct orientation inconsistency; the model uses a consistent left/right convention. |
| Figs. 1–2 omit waiting j0, present in Fig. 3 | Correct omission from those figures. Adding a repetition penalty for waiting preserves the tested initial decisions. This does not prove the original authors required that transition in every displayed fragment. |
| Fig. 1 j4 omits +CH although the chicken eats seeds | Correct omitted effect under the stated value definition; the corresponding sensitivity test preserves the tested choices and the relevant unanswered CQ17. |
| Figs. 5 and 7 use one −P sign where the text names two harmful goals | Correct difference in information granularity. **Fig. 5 is decision-sensitive under unit-count interpretation; Fig. 7 is not for the tested choices.** |
| Some return moves omit an additional repetition penalty in the prose | Correct observation. The model follows the narrated penalties; the extra q0 penalty is tested as invariant. This is not exhaustive path-sensitive planning validation. |
| Train standing still is said to demote progress but changes none of the listed Table 1 propositions | Correct representational gap for a changed-goal AS2 encoding. The score-only sensitivity check adds a penalty outside the goal-labelled model and preserves the recorded train choices; it is not an additional validated AATS encoding. |

## Verification performed

| Check | Result and evidential scope |
|---|---|
| Original Python suite on the fresh base | 56 tests passed; this did not prevent the new counterexamples. |
| Corrected `tools/test_aats.py` | 71 tests passed, including new state-pair, state-identity, joint-tuple, input-integrity, audience-error and figure-sensitivity checks. |
| `tools/aats.py check` | 2 models, 17 situations, 32 choice expectations; passed; generated paper cases unchanged. |
| Independent fact join | Exact AS2 argument sets and comparison results checked from facts on 17 paper situations and 300 random models; no predicted argument ids consumed by this oracle. |
| New actual-engine adapter tests | Both tests passed on Go 1.24.13 and Go 1.27.4. One uses authored expected arguments/labels; the other rejects inconsistent annotations before emission. |
| Focused live-library AS1/AS2 cases and CLI round trips | 68 checks passed, zero known/unexpected failures, using Go 1.27.4 and a freshly built vendored CLI. |
| Same-pair counterexample | Pre-fix graph built an additional `act_b` argument; fixed adapter rejects the model with an inconsistent-effects diagnostic. |
| Fig. 5 counterexample | Textual two-penalty model selects chicken; unit graphical interpretation selects dog; both outcomes recorded. |

The focused CLI was built before concurrent final engine-profile changes in the overall audit. The root integration review must run its complete clean, pinned-toolchain verification against the final combined tree. These focused results are evidence about this review's changes; they are not a substitute for that final aggregate run or a claim that the old 10,428 total is a semantic proof.

## Source acceptance fingerprints

The records below were not rewritten during this source review. Acceptance binds these exact semantic fingerprints, including their declared source contracts and used templates. Detailed registry-ready recommendations are in `reviews/atkinson/source-acceptance.json` in the review package.

| Record | Semantic fingerprint |
|---|---|
| `practical_reasoning_as1` | `sha256:dbfb6c32b24c2c31be568d53fbce03bd03fef3526b728654f72ae186521ac916` |
| `practical_reasoning_as1_demotion` | `sha256:20a0d96e5ac74a3c922c28c128ab380861a453cf58b27a2ff848c805f755bc98` |
| `practical_reasoning_as1_preclusion` | `sha256:b944cfffa12edb9e4471447c1104b47ff2c386b18c17dcb7989a7fc6942b8813` |
| `practical_reasoning_as2` | `sha256:3608e0193d17da8a500fed7c772915fa9f84e93b9c34e4cf9c58135654b73ec9` |
| `practical_reasoning_as2_demotion` | `sha256:96af2657c88049cd937f2327d1e33c926760c068335f258d79ec736aa1ad5852` |

The independent acceptance concerns source interpretation and CQ sorting. The engine does not authenticate the source, the truth of assumptions, the identity of a reviewer, or the correctness of a user-supplied AATS model. The approval fingerprint also does not itself bind the adapter's implementation or the text of an external comparison contract. Those belong to the separately reviewed and pinned release tree.

## Remaining decisions for this component

No further broad opinion round is required to decide the five source-record interpretations. Retain both AS1 and the declared AS2 model profile; retain the textual two-harmful-goal reading of Fig. 5; keep the comparator as an explicit project policy; require a separately modelled bridge for AS1 comparisons; and reject inconsistent state-pair effects.

The root reviewer must independently inspect these newly authored code corrections and bind successful final integration checks to the candidate commit. After that, this component can enter the broader validation/benchmarking effort under the stated scope. Future counterexamples, better models or a published clarification of the paper can reopen specific choices. No assertion here implies that unbounded semantic correctness has been proved.
