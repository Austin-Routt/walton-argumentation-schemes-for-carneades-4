# Independent closure review of all 50 source holds

**Baseline:** `e6d2d6f427c6ec0447cae514810c256903b15779`. **Date:** 2026-11-21. **Reviewer:** ChatGPT independent closure scope reviewer. This reviewer did not author the baseline schemes or the implementer’s engine changes. The reviewer wrote new tests for unchanged schemes and this adjudication. Integration changes subsequently made by other reviewers require their own final evidence and are not pre-approved by the recommendations here.

## Decision and boundary

The old split of 70 approved and 50 held records was a useful conservative checkpoint, not a permanent verdict. Three unchanged interpretations can now advance on actual source comparison and discriminating cases: `from_example`, `verbal_classification_arbitrariness`, and `definition_to_verbal_classification`. Ten other holds are concrete context-identity defects whose repairs can be independently reviewed. Further held records omit printed alternatives or require a dialogue/model contract; these are implementation tasks, not unavailable-book problems. Some inherited records still require genuinely uninspected original sources. Full stable gate L remains stricter than a usable, declared preview.

Every one of the original 50 holds was read against its current formula, source-ledger row, disposition and applicable supplied source. The complete machine-readable decision is in [`closure-hold-adjudication.json`](closure-hold-adjudication.json). Each row contains the baseline reason, new adjudication, exact next action, discriminating tests, source locations and remaining original-source obligation. Its baseline status must not be mistaken for a final integrated result. A later integration can close a row only with recorded review and execution evidence.

A source-certified executable interpretation need not be a literal transcription. It must identify the source it targets, disclose differences, justify its mapping and pass its evidence-based tests. Conversely, renaming a broad title to a narrower title does not implement an omitted source branch. Missing original-source evidence is not supplied by a green runtime check, a later reprint or a publisher’s table of contents.

## Three unchanged interpretations independently approved

### `from_example`

The supplied book’s entry 6.1, printed p.314, distinguishes the example from the generalization. The existing repaired conclusion `usually_properties(P1,P2)` retains the generalization’s two properties and no longer makes each example a different conclusion. Its accepted meaning is qualitative and defeasible. It asserts neither a universal implication nor a numeric frequency. Typicality and impairing circumstances refer to the identified example assessment; applications must identify a new assessment when the contextual target generalization changes. The scheme does not assess a dataset statistically.

The prior hold asked for a policy-level test that one exceptional instance does not automatically refute a usually claim. Four independently written cases now check precisely that boundary: a penguin counterexample does not automatically refute a bird/flight generalization, an accepted contrary assessment defeats the actual generalization, questioning the default connecting example to generalization removes support, and a rebuttal about swimming does not affect flight. These are constructed applications, not claimed printed examples. They demonstrate the declared CQ4 policy; they do not calibrate an empirical frequency.

**Decision:** approve the unchanged declared interpretation at `example-passed`. The separately cited historical 1995 reference is part of the compendium’s bibliography; this book-origin record is certified against the supplied 2008 entry, without a claim that the older book was read.

### `verbal_classification_arbitrariness`

Printed p.320 gives the conditional warrant, the relevant dialogue’s need for a nonarbitrary definition, and the arbitrary definition in that argument. The current scheme keeps these premises. Its third critical question becomes a supported harmlessness exception. This is a static evidence policy: the criticism requires its warrant and need premise; a proved harmlessness circumstance defeats it; merely questioning an exception does not prove the exception. The printed page does not prescribe automatic suspension on the utterance of this question. That differs from the explicit practical-reasoning discussion on printed p.96.

Four new cases cover supported harmlessness, question-only harmlessness, an unestablished need premise and harmlessness of another argument. They preserve argument/property scope without adding a premise that merely repeats the desired rejection conclusion.

**Decision:** approve the unchanged interpretation at `example-passed`, limited to this declared static profile. The record does not claim to execute a complete dialogue protocol or certify the 1996 book cited in the supplied compendium.

### `definition_to_verbal_classification`

The supplied p.319 entry 15 provides the definition-specific form and two critical questions. The current record retains the object, definition and classification roles and both appropriate definition challenges. Existing positive, supported biased-definition and wrong-definition controls were rerun. Definitions are identified versions within the relevant classification context, not generic labels whose disputed content varies silently between cases.

The separately cited 2006 page is now compared through the indexed source text described in S1 below. It prints the ordinary classification form. This closes the access question and establishes a known difference; it does not make two distinct forms identical.

**Decision:** approve the unchanged **2008 entry-15 executable interpretation**, retaining the inherited 2006 citation as related provenance and explicitly recording that it is not an identical printed definition-specific scheme. This is a reconciled adaptation, not original-author endorsement.

## Source access findings and honest limits

The following locator notes describe what was actually accessed. No newly discovered complete book is redistributed in this patch. S1 and S2 were read as indexed text, not as inspected page facsimiles; do not describe those readings as photographic verification. The user-supplied 2008 book was available in full; decisive printed pp.314,319–320 were also rendered and visually inspected in this closure pass. Atkinson’s supplied paper was checked against the earlier full comparison and decisive passages; its formal CQ page 862 was visually inspected again.

| ID | Source and locator | Retrieval and finding |
|---|---|---|
| S1 | [Walton, *Fundamentals of Critical Argumentation* (2006), indexed text](https://www.academia.edu/37758002/Douglas_Walton_Fundamentals_of_Critical_Argumentation_Critical_Reasoning_and_Argumentation_pdf), printed pp.96–97 and128–130 | Public indexed text opened and searched, 2026-11-21. The analogy CQ concerns another similar contrary case without a stronger more-on-point requirement. Page129 presents ordinary classification and definition-related questions; the distinct definition form is supplied by the 2008 compendium. [Official publisher frontmatter](https://assets.cambridge.org/97805215/30200/frontmatter/9780521530200_frontmatter.pdf) corroborates edition metadata, not those pages. |
| S2 | [Walton, *Legal Argumentation and Evidence* (2002), indexed text](https://www.academia.edu/37757995/_Douglas_N_Walton_Legal_Argumentation_and_Eviden_pdf), printed pp.45–50 | Public indexed text opened and searched, 2026-11-21. Page46’s position-to-know form has the expected access, honesty and assertion question families. Pages47–48 explicitly discuss question-induced burden transfer. Pages49–50 reprint expert-opinion material; this does not verify the separately cited 1997 original. [Publisher metadata](https://www.psupress.org/books/titles/978-0-271-02177-5.html) confirms the work. |
| S3 | [Prakken, Reed and Walton, *Argumentation Schemes and Generalisations in Reasoning about Evidence* (2003), author-hosted PDF](https://webspace.science.uu.nl/~prakk101/pubs/icail03.pdf), §3.2, PDF pp.4–5 | Downloaded and rendered from the author’s university site. SHA-256 `a5912213c06537233c4a0658d74e339a57e4df44a4d0b5fe3eaaf28d22aa3541`. The witness form and five question themes are present. The implementation’s belief premise is a sincerity adaptation, and the account-specific bias issue remains. This completes comparison of this cited article, not the other cited witness book. |
| S4 | [Walton, *Argumentation Schemes for Presumptive Reasoning* (1996), official retailer preview](https://api.pageplace.de/preview/DT0400.9781136687068_A24375265/preview-9781136687068_A24375265.pdf) | The accessible preview is frontmatter and introductory material, not the cited pp.101–105. It does not certify those historical pages. Book-origin entries16/17 are being reviewed against the fully supplied 2008 compendium. |
| S5 | [Walton, *Appeal to Expert Opinion* (1997), publisher page](https://www.psupress.org/books/titles/0-271-01694-9.html) | Bibliographic access is not access to pp.211–225. Later reprints were found, but the original-source hold remains. |
| S6 | [Walton, *Goal-Based Reasoning for Argumentation* (2015), official frontmatter](https://assets.cambridge.org/97811071/19048/frontmatter/9781107119048_frontmatter.pdf) and [excerpt](https://assets.cambridge.org/97811071/19048/excerpt/9781107119048_excerpt.pdf) | The contents locate the instrumental/value-based section at p.92. The excerpt is introductory; the actual target scheme and CQ inventory were not inspected. Do not assign zero questions or a fabricated page-specific example. |
| S7 | [Walton, *A Pragmatic Theory of Fallacy* (1995), publisher listing](https://www.ubcpress.ca/a-pragmatic-theory-of-fallacy) and [Google Books catalogue](https://books.google.com/books/about/A_Pragmatic_Theory_of_Fallacy.html?id=q-TWAAAAMAAJ) | Listings/snippets do not supply the relevant pp.142,147–148,150,152,155–156. No original-source approval is granted from snippets or unrelated reproductions. |
| S8 | [Walton, *Witness Testimony Evidence*, official excerpt](https://assets.cambridge.org/97805217/07701/excerpt/9780521707701_excerpt.pdf) | The accessible ten-page introduction does not provide the target schema and its full question discussion. S3 resolves only the independently cited 2003 article. |
| S9 | [Walton, *Slippery Slope Arguments*, official Oxford edition metadata](https://academic.oup.com/book/49915) | Confirms a 1992 edition. The supplied compendium p.340 cites 1992, whereas both inherited records cite 1999 at pp.93,95. Preserve the discrepancy until the exact edition is identified; do not silently claim their pagination has been compared. |

The unavailable-source list is therefore broader than “the 1995,1996 and2006 books.” Conversely, the exact cited 2002 and2006 formulations above are no longer simply unavailable. For `exeption_to_rule`, a published-source comparison can settle a technical interpretation; unavailable private 2019 correspondence cannot be reconstructed or cited as if read.

## Atkinson decision: qualified AS1 now, separate AS2 extension later

**Decision made:** retain `practical_reasoning1` as the explicitly qualified legacy AS1 adapter in the preview, outside the source-approved set. Do not add an incomplete AS2 record to this release. The supplied paper is available; source access is closed. The remaining work is a real model implementation and independent validation task. Deferring that larger model from a bounded preview does not discharge the all-record stable gate.

The prior Atkinson review’s partition is supported, but its full-coverage count needs correction. Printed pp.858–859 enumerate 16 questions and then distinguish CQ17’s other-agent choice from CQ2’s transition existence. The formal definitions on p.862 and the farmer examples on p.866 confirm the key distinction. The current record offers seven direct AS1 treatments, five partial policies and five questions without separate mappings. Therefore **ten questions lack full-source coverage**: CQ5–9,11,12,14,15,17. The historical eight-item worklist included CQ5–7 and11,12,14,15,17; it excluded CQ8–9 because those have nonblank hooks, although their unconditional merged policies are also partial. Calling that narrower set “eight full-source mappings” was inaccurate in this reviewer’s first draft and is corrected here. CQ10 is supporting evidence, so another value can supply another argument for the same action. Merely counting five absent atoms misses these distinctions.

### Required contract for the deferred finite AS2 adapter

This is an implementable extension specification, not a menu requiring another owner decision:

1. **Model identity and validation.** Each finite model identifies agents, states, initial state, propositions, individual actions, joint actions, action preconditions/protocols, transitions and values. Validate references, membership and transition endpoints before generating arguments. Duplicate or out-of-universe identifiers are errors. The deciding agent and model context belong in generated statement identities.
2. **Goal realization and value assessment.** Represent the goal proposition’s relevant change across the transition. Value promotion/demotion belongs to a transition and the agent’s deliberative assessment; neither a global goal token nor a bare action token can silently carry a different agent’s values. CQ12,14,15 need independently checked model/state/goal conditions, not a repurposed worthiness default.
3. **Choice and cooperation.** Separate the existence of a transition under a particular joint action from whether the other agents will choose their components. CQ17 requires an explicit epistemic/cooperation assumption or argument about those choices. A transition lookup alone answers a different question.
4. **Alternatives and adverse values.** Generate the competing justifications for CQ5–9 and11 with actual alternative actions and values. Record incompatibility where applicable. Do not make all different recommendations contradictory, all side effects automatic defeats or all alternatives superior merely because they exist. CQ10 must preserve additional positive support.
5. **Evaluation policy.** Supply value preferences as explicit, case-specific inputs. Incomplete preferences may leave a choice unresolved. The adapter must not invent a universal value ordering to reproduce one example. Preserve separate arguments and report which comparison governs the result.
6. **Independent acceptance set.** Reproduce the farmer’s conditional Arg6 with two positive values; its CQ17 challenge and alternate state; Arg8 under the paper’s stated assumptions; and a changed value preference with the corresponding changed comparison. Add invalid state/action/goal controls, all 17 question contracts, wrong-agent/model/transition controls, and export/import/second-run checks. Include a second scenario such as the paper’s trains example to avoid fitting only the farmer narrative.

AS2 is not just five additional exception predicates. Shipping a flat named variant without the model and comparison contracts would create another unsupported completeness claim.

## Remaining feasible work is concrete

The ten context repairs recommended here are: agent-specific value reasons; cause/effect-specific interference; sign/candidate-specific competing explanations; explicit abductive membership plus universe/data-scoped conclusion; argument-specific circumstantial credibility; institution-specific counts-as; language/definition-specific synonymy and rebuttal; a pleader/case/rule/action/sanction-scoped excuse; and observer-specific perceptual unreliability. These can be tested with two contexts sharing all otherwise-equal terms, so a wrong-scope fact would have produced a false argument or defeated valid independent support before the repair. They are not cosmetic wording changes.

Further manageable tasks include the source’s omitted snob/property/purchase branches, actual versus normative group goals, action/policy excuse branches, small sacrifice, bad/less-bad ends/means variants, and the true-presumption branch of negative practical reasoning. Their exact tasks and controls are in the 50-row JSON. They must remain visible until implemented; a selected positive branch is not the full printed entry. Entry9.2/10.2 variants are independently assigned to another reviewer, with the explicit requirement that an only-if warrant cannot be reversed into sufficient support.

The practical-reasoning burden issue calls for a separate challenge/answer adapter rather than a global change to all exceptions. Printed p.96 expressly describes temporary defeat and restoration. Entries16/17’s printed why-questions do not themselves establish that same rule; the static burden chosen and tested here is therefore an interpretation with stated support. Likewise, a proof standard supplies a strength threshold, not a dialogue’s probandum or burden holder. Those must be independently represented for the ignorance family.

Several legacy records are known adaptations even after their sources are read: `sunk_costs` differs from the temporal precommitment form; the slippery-slope pair are components rather than the complete normative scheme; `ethotic2` infers contradiction where the compendium describes reduced plausibility; and `established_rule` has a rule/content form rather than the explicit actor/action form. Source acquisition alone will not implement these distinctions. Keep useful legacy IDs labelled and add separately named source-shaped forms when needed.

## Validation actually performed for this patch

The clone contains the baseline engine and unchanged scheme formulas. The new cases were run using the local Go 1.24.13 toolchain, `GOTOOLCHAIN=local`, `GOFLAGS=-mod=mod`, and `VV_STRICT=1`.

| Focused run | Registered results | Outcome |
|---|---:|---|
| `go test -count=1 -v -run '^TestCases/cases/closure_scope'` |16 API/CLI checks for8 independent cases|All pass|
| `go test -count=1 -v -run '^TestCases/cases/(book\|independent_first)'` |332 existing API/CLI checks, including all book cases and matched independent cases|All pass|

The second selection also matches `independent_book_second.yml` under Go’s test selection rules;332 is the actual result, not an estimate of the intended selector. The cases establish these declared finite contracts, not a proof over all possible graphs. No full-suite or dual-toolchain integration result is claimed by this patch; the integration owner must regenerate the check manifest and rerun final acceptance after combining changes.

Evidence SHA-256 values:

- `scope-focused.log`: `6b6042323bf7e0d4d2de63f04a82971aeb383dac382ae5e2d7f584fb9e2a4463`.
- `existing-source-cases.log`: `c78998e7d8b3a59cae3a3b763a8f8fb6fa8b2d46b44611fdba23c01df1426d65`.
- Supplied 2008 book: `d010348014e8b80fa561b68e5fac732cce0b3905fb0b1d522dd6fd2578a69634`.
- Supplied Atkinson paper: `dad4124c052c48b5c82de02494aac6971de113bb8452ae9fc6e764aa35e43cbc`.

Only the three explicitly approved unchanged schemes receive new semantic fingerprints in `independent-approvals.json`. Existing approvals are preserved. The ten proposed context repairs and further variants receive no approval from this document until their actual final formulas and independent controls have been reviewed.
