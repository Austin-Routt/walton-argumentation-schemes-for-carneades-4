# Independent review of Atkinson and Bench-Capon (2007)


> Current closure correction (2026-11-21): seven direct, five partial and five absent mappings leave **ten** questions without full-source coverage. The historical eight-item/blank-hook worklist below omitted partial CQ8–9 and is preserved as history, not the current completeness count. Current stages and obligations are in `closure-current-register.md`.

**Review date:** 2026-11-21. **Repository examined:** commit `49ad425cbf7afeebcbd9a07ee479d60681ca33ea`. **Record:** `practical_reasoning1`. **Decision:** the original-source comparison is completed, the incorrect reference-case attribution is repaired, and focused regression tests pass. Full source-fidelity approval of the inherited record is withheld. Its formulas remain the legacy Carneades adaptation of AS1; they do not implement the paper's complete action-based alternating transition system (AATS) account or all 17 critical questions.

This review was performed independently of the author of the merged scheme and engine changes. The reviewer subsequently authored the narrowly scoped ledger, validation and fixture changes described here. These changes therefore still require the main integration gate's clean rerun. They are not represented as independently reviewed by a second person after their creation.

## 1. Evidence and scope

The supplied source is Katie Atkinson and Trevor Bench-Capon, *Practical reasoning as presumptive argumentation using action based alternating transition systems*, **Artificial Intelligence 171 (2007), 855–874**, DOI `10.1016/j.artint.2007.04.009`. The supplied PDF has 20 pages and SHA-256:

```
dad4124c052c48b5c82de02494aac6971de113bb8452ae9fc6e764aa35e43cbc
```

The complete paper was read. The decisive printed pages 858–862 and 866 were also inspected as rendered PDF pages, including the mathematical signs in the formal CQ definitions. Printed page 855 is PDF page 1: add 854 to a PDF page number to obtain the printed page. Text extraction alone is insufficient for these formulas because membership and inequality signs can be lost.

Exactly one record of the current master library cites Atkinson: `practical_reasoning1`. The practical-reasoning category file repeats that record; it is not a second source-derived scheme. `value_based_practical_reasoning` is the separate 2008 compendium entry 22.4. `practical_reasoning2` cites Walton (2015), so this paper cannot independently approve it.

The inherited record's variables, premises, conclusion, assumptions, exceptions and absent `roles` field were compared with `third_party/carneades-4/examples/AGs/YAML/walton.yml`. All six fields match. No engine or scheme formula is changed by this review. The source-comparison evidence comes from the supplied paper, not from the fact that existing tests pass.

### Reading map

| Printed pages | PDF pages | Material relevant to the decision |
|---|---|---|
| 855–857 | 1–3 | Defeasible practical justification, goals and values; motivation for a formal model |
| 858 | 4 | AS1, its separation of resulting state, goal and value, and CQ1–16 |
| 859 | 5 | The additional joint-action CQ17; responses, subjective values, and negative AS1 variants |
| 860–861 | 6–7 | AATS states, agents, actions, preconditions, transitions, state propositions and transition-relative values; AS2 |
| 861–863 | 7–9 | Problem formulation, epistemic reasoning and action choice; formal CQs, especially positive CQ10 |
| 863–869 | 9–15 | Farmer example, alternative actions, animal choices, value preferences and changing circumstances |
| 870–872 | 16–18 | Two-trains example, social constraints and other agents' choices |
| 872–874 | 18–20 | Scope and conclusions, limitations and bibliography |

The source is now available and has been checked. “Cannot access Atkinson and Bench-Capon (2007)” is no longer an accurate reason to leave this comparison undone. Availability does not resolve the implementation differences identified by the comparison.

## 2. What the current record actually represents

AS1 makes a presumptive argument for performing an action in given circumstances because of the resulting state, a goal realized there, and a value promoted by that goal. The four premises and action conclusion of `practical_reasoning1` reproduce that broad structure. The paper expressly separates the state, the goal within it and the value that gives a reason to want the goal (p. 858).

| Current field | Source comparison | Independent disposition |
|---|---|---|
| `current_circumstances(S1)` | AS1 current circumstances | Accept as an AS1 premise abstraction |
| `would_bring_about(A,S1,S2)` | Action's stated resulting circumstances | Accept as a stipulated AS1 transition; it does not independently establish other agents' cooperation |
| `would_be_realized(G,S2)` | Goal in the resulting state | Accept as an AS1 abstraction; AS2 also represents change relative to the initial state |
| `would_promote_value(G,V)` | AS1 goal/value link | Accept within a fixed deliberative context; it omits AS2's agent and transition arguments |
| `should_be_performed(A)` | AS1 action recommendation | Accept as a context-dependent recommendation; an application must encode distinct action instances when agents or contexts differ |
| `legitimate_value(V)` default | Related to CQ16 | Source supports asking about legitimacy; the default/questioned burden is an inherited implementation policy |
| `possible(A)` default | Related to CQ13 | Source supports questioning possibility; joint-action membership in AS2 is richer than this atom |
| `worthy_goal(G)` default | No matching printed CQ | Retain as an explicit inherited policy. Worthiness does not answer CQ15's realizability question |
| Three `*_more_effectively` exceptions | Related to CQ5–7 | Stronger and narrower than the paper's questions about alternative means; not complete mappings |
| `side_effects(A,S1)` exception | Related to CQ8 and CQ9 | Merges different kinds of value demotion into an unconditional undercutter; not the paper's general value-comparison treatment |

The record has five variables, `[A,S1,S2,G,V]`. None identifies the deciding agent, a joint action, another agent's component action, or a preference ordering. There are no premise-role labels. Omitting labels alone is not a defect; omitting the semantic participants limits what can be expressed and distinguished.

AS2 is more specific. Values belong to agents; the value function evaluates a transition between states; a joint action includes each participant's component; a goal proposition changes between the initial and resulting states. The current goal/value relation is global for a given `G,V` pair. Reusing an unqualified goal name across two transitions or agents can consequently reuse a value claim where AS2 would distinguish the instances. Treating constants as fully qualified case identifiers can avoid some accidental sharing, but it is an application convention, not automatic AATS validation. These restrictions follow from pp. 860–862 and the current predicate arguments.

## 3. All 17 critical questions

“Direct AS1 treatment” below means the existing abstraction can represent the relevant premise, default or additional supporting argument. It does not certify the complete formal AATS semantics. “Partial policy” means the inherited record reacts to a narrower or stronger condition. “No separate mapping” means there is no distinct executable representation or documented model-construction contract for the printed question.

| CQ | Source question, paraphrased | Current treatment and decision | Printed pages |
|---|---|---|---|
| 1 | Are the believed circumstances correct? | Direct AS1 premise `current_circumstances(S1)`. AS2 additionally exposes the initial state and action preconditions. | 858, 862 |
| 2 | Does the action have the alleged consequences in these circumstances? | Direct AS1 premise `would_bring_about(A,S1,S2)`. Do not use this alone to claim the distinct CQ17 is implemented. | 858–859, 862 |
| 3 | Do those consequences realize the desired goal? | Direct AS1 premise `would_be_realized(G,S2)`. AS2's state-change detail remains outside this abstraction. | 858, 861–862 |
| 4 | Does the goal promote the stated value? | Direct AS1 premise `would_promote_value(G,V)`, subject to the agent/transition limitation above. | 858, 861–862 |
| 5 | Is there another way to achieve the same consequences? | Partial: `bring_about_more_effectively(S2,A)` requires greater effectiveness, which the printed question does not. Full alternative comparison remains unresolved. | 858–859, 862 |
| 6 | Is there another way to achieve the same goal? | Partial: `realize_more_effectively(G,A)` imposes an additional criterion. Existence of an alternative alone does not decide between actions. | 858–859, 862 |
| 7 | Is there another way to promote the same value? | Partial: `promote_more_effectively(V,A)` is a stronger application policy, without an alternative witness or full comparison. | 858–859, 862 |
| 8 | Does a side effect demote the same value? | Partial: aggregate `side_effects(A,S1)` merges CQ8 and CQ9 and directly undercuts. The paper discusses negative AS1 arguments and their evaluation. | 858–859, 862 |
| 9 | Does a side effect demote a different value? | Partial: the same aggregate exception drops the other value and the possibility of a contrary preference. | 858–859, 862, 868 |
| 10 | Does the action also promote another value? | Direct supporting treatment within AS1: another `G,V` instantiation yields another argument for the same action. It must not be turned into an exception. New tests confirm separate arguments and correct polarity. | 858, 862, 866 |
| 11 | Does this action prevent another action that would promote a different value? | No separate mapping. Needs the incompatible alternative and its value justification, followed by comparison. | 858–859, 862, 865 |
| 12 | Are the described initial circumstances possible? | No separate state-membership or possibility mapping. A documented valid-model precondition could discharge this, but none is established by the current scheme. | 858, 862 |
| 13 | Is the action possible? | Direct AS1 default `possible(A)` under the inherited questioning policy; not a complete check of joint-action membership. | 858, 862 |
| 14 | Are the described consequences possible? | No separate successor-state membership or possibility mapping. A valid-model contract is an alternative to an additional scheme field. | 858, 862 |
| 15 | Can the goal be realized? | No separate realizability mapping. `worthy_goal(G)` says something different. | 858, 862 |
| 16 | Is the stated value legitimate? | Direct AS1 default `legitimate_value(V)` under the inherited questioning policy. | 858, 862, 865 |
| 17 | Will the other agents execute their required parts of the joint action? | No separate mapping. The paper explicitly adds this beyond CQ2 and distinguishes a nonexistent transition from failure to choose the intended joint action. | 859, 861–862, 865–866 |

This is **seven direct AS1 treatments** (1–4, 10, 13, 16), **five partial treatments** (5–9), and **five questions without separate mappings** (11, 12, 14, 15, 17). There are **eight incomplete full-source mappings** in the revised ledger: the five without separate mappings plus the three alternative-choice questions 5–7 whose existing stronger exceptions do not settle the source question.

The numbers are classifications for this review, not a percentage of semantic correctness. Even an existing atom cannot certify the factual truth of a transition or a value judgment. Conversely, a missing dedicated atom is not automatically a defect: CQ12, CQ14 and CQ15 can in principle be discharged by a valid-model construction contract, or by appropriately supported challenges to the asserted premises. The current record and contracts do not define and test that discharge. The ledger therefore leaves these mappings open rather than assuming it. The proposed `assumption` classifications for these unresolved questions describe one possible burden policy; the paper does not prescribe that policy.

### Positive CQ10 is a substantive correction to the review framework

The paper explicitly explains that CQ10 can support performing the action; it questions the justification rather than the performance (p. 862). A ledger restricted to premise, assumption, exception and rebuttal cannot honestly classify this use as an attack. The new `support` classification names the same conclusion and is validated against the scheme's conclusions. A negated conclusion is rejected in this classification.

The current engine already preserves two arguments with different value parameters. No scheme edit is required to express the paper's positive-value example at AS1 level. The new tests verify both arguments, their positive weights, and continued support when only the additional value's default is questioned. This supplies positive evidence for the existing implementation instead of inventing a defect or making all CQs defeat the argument.

### Alternative and adverse-value questions do not have one automatic outcome

CQ5–11 concern action choice. The paper treats relevant responses as arguments to compare, not as an instruction that every alternative or negative side effect must defeat the original recommendation (pp. 859, 862). In the farmer example, preference for companionship can outweigh demotion of progress; the ordering can also change with circumstances (p. 868).

The existing stronger exceptions can be retained as application rules, but they must be identified as such. Neither their names nor passing exception tests prove source fidelity. The new `alternative` ledger classification records that an executable comparison contract is still needed. It cannot obtain a nonblank mapping merely by naming an action conclusion, and it blocks advancement to `questions-sorted` while unresolved.

## 4. The incorrect source attribution and its repair

Before this review, `practical_reasoning1/book` cited “Atkinson and Bench-Capon (2007); Hal and Carla, p. 322.” The supplied paper spans pp. 855–874 and uses farmer and trains examples; it does not supply that claimed page/example. The fixture may be a reasonable abstract practical argument, but its citation was not evidence for this record's own source.

The replacement uses **farmer Arg8, p. 866**: rowing the chicken from `q0` to `q9`, realizing the chicken's presence on the right bank and promoting progress. The names `row_chicken`, `chicken_right` and `progress` are this library's explicit abstraction of the source example. The source later selects Arg8 after stating particular beliefs about the animals and value choices. The fixture stipulates the premises of Arg8 and verifies its resulting argument; it does not claim to reproduce that entire deliberation.

The CQ10 companion uses **Arg6, p. 866**, conditional on the initial `q0` to `q7` transition: rowing the dog promotes progress and companionship. Two value instances are checked. The paper subsequently challenges this transition through CQ17 and discusses an alternative resulting state. Keeping that qualification is essential: a fixture that unconditionally claimed the paper finally recommends rowing the dog would be inaccurate.

## 5. Changes completed in this review

1. The `practical_reasoning1` row now identifies its claim as an executable interpretation and advances from `mapped` to `transcription-checked`. It records the paper hash, exact source locations, all 17 CQ decisions, the eight incomplete mappings, and the limits of approval. It does not advance to `interpretation-approved`, `questions-sorted` or `example-passed`.
2. The schema adds `source_questions` for upstream-only records. An omitted value means unknown; explicit `0` means verified no printed questions. A complete upstream-only row must declare its original-source inventory. Compendium records continue to use the compendium count. This closes the prior path by which an upstream-only record could appear complete without any independently declared source-question count.
3. `support` and unresolved `alternative` classifications preserve the paper's actual directions of reasoning. Completeness checks refuse unresolved alternatives.
4. A validation defect is fixed: previously an unknown or empty question kind with blank `as` could bypass kind validation. At a complete stage an empty kind could also reach a string-slice panic. Kinds are now validated before the blank-mapping branch, and the article helper is defensive. Regression tests cover both early and complete stages.
5. The source worklist reports `questions_in_source`, displays unknown counts as blank, and excludes `n: 0` application additions from `questions_sorted`. Previously the count could mix printed CQs with policy additions. Consumers depending on the old CSV header `questions_in_book` must update that column name.
6. The wrongly attributed reference case is replaced with Arg8, and 12 persistent paper-specific cases are added. They cover positive CQ10 support, questioning only the extra value, wrong goal, wrong result state, wrong initial state, questioned and supported action possibility, side effects belonging to a different action or state, and the distinction between a challenged value argument and a conclusion with other support.
7. Argument-local unit tests select the real `practical_reasoning1` record from `walton.yml`, keeping its actual formulas. They verify the target argument's weight and undercutter, rather than allowing `practical_reasoning2` to mask a failure by independently supporting the same conclusion.
8. The named-check manifest and minimum increase by 24, from 7,994 to **8,018**, for the 12 additional cases and their fresh-process round trips. Additional unit tests run separately in the normal harness's unit-test stage and are not included in that registered-check count.

No new scheme variant, predicate rename, engine behavior, value preference or public-release action is introduced. In particular, `worthy_goal` is not silently repurposed to mean realizability, and the existing exceptions are not claimed to be source-prescribed burdens.

## 6. Validation actually performed

The focused review used **Go 1.27.4, linux/amd64**, the repository's locally pinned Carneades and GoCHR modules, and the strict ledger setting.

The new unit-test groups passed: original-source CQ inventory, validation before blank mappings, positive-support polarity, unresolved comparison rejection, unknown versus explicitly zero CQ count, accurate worklist counts, ten argument-local default/context controls, and two positive-value controls. These tests include a missing-CQ17 regression and a missing-CQ10 regression.

The registered focused command was:

```sh
VV_STRICT=1 VV_RESULTS=focused-results.jsonl \
  go test -count=1 -v \
  -run 'TestSource|TestCases/cases/(atkinson_2007|book)'
```

Result: **850 checks, 850 pass, zero known failures, zero unexpected failures and zero stale ledger entries**. The result comprises 24 new paper-specific checks, 240 checks for all 120 reference cases and their round trips, 120 library source-placement checks, 305 source-ledger checks, and 161 migration checks. The source worklist was also executed successfully. `git diff --check` passed.

These are focused results, not a claim that the final combined branch has passed the complete 8,018-check gate or both toolchains. The main reviewer owns that clean integration rerun. No number in this report should be used to certify an older release pin automatically.

The passing checks establish consistency with their stated contracts and finite examples. They do not prove all possible graphs correct, the empirical truth of asserted circumstances, a value ordering for every application, or the complete formal theory of the paper.

## 7. Reconciling “five gaps” with the new paper

Before this review, the source ledger recorded five nonpremise/nonstandard questions with no executable `as` mapping, across four records:

| Record | Recorded questions |
|---|---|
| `definition_to_verbal_classification` | CQ2 |
| `established_rule` | CQ2 and CQ3 |
| `negative_consequences` | CQ1 |
| `positive_consequences` | CQ1 |

That restricted count was reproducible. It did **not** establish that the entire library had only five unresolved original-source questions: `practical_reasoning1` was still `unchecked`/`mapped` with none of the paper's 17 CQs inventoried.

After this review's explicit mappings, the same blank-mapping measure is **13 across five records**: the previous five plus eight on `practical_reasoning1`. There are not 17 new wholly absent CQs. Seven already have working AS1 treatments; five have partial legacy treatments; five lack separate mappings. The three incomplete alternative treatments contribute to the eight new blank mappings. CQ8 and CQ9 remain nonblank but carry explicit fidelity limitations, illustrating why a blank-field count is not a complete measure of source correctness.

The 13 count is for this standalone patch against the stated baseline. Other independent fixes applied during integration can close some of the earlier five gaps; the final report must recount the combined ledger rather than reuse this standalone total.

The revision changes the inventory and the evidential claim, not the runtime failure count. A zero-entry `known_issues.yml` lists no failing checks under the current executable contracts; it does not erase source gaps held in a different ledger.

## 8. Approval and release decisions

| Decision | Disposition |
|---|---|
| Has the paper now been obtained and compared? | Yes, completely, with decisive rendered pages checked. |
| Can the inherited record be used as documented legacy AS1 practical reasoning? | Yes, conditionally on the explicit premises, inherited defaults and stronger exception policies, within the supported engine profile. |
| Is it an exact transcription of AS1 plus all printed CQs? | No. Defaults and stronger exceptions are implementation choices; eight full-source mappings remain unresolved. |
| Is it a full AS2/AATS implementation? | No. Agent, joint-action, state-membership, transition-value and preference details are absent. |
| Is full original-source interpretation approval warranted now? | No. The comparison justifies `transcription-checked`, with scoped approval of the identified AS1 behavior, not the later completeness stages. |
| Does a green case justify moving the record to `example-passed`? | No. That would skip the unresolved interpretation and question-mapping requirements. |
| Can this review alone certify R1 at its historical pin? | No. Release evidence must match the actual release commit and full gate. |
| Can an unrestricted stable library claim include this record as fully source-validated? | No. Gate L requires completing its declared scope and questions, or visibly keeping it outside the source-validated stable subset. |

There is no need to ask the owner to decide whether CQ10 is an attack, whether worthiness means realizability, or whether the 17-question inventory exists. The source settles those matters. Application-specific value preferences remain input data or an explicit application policy; a reviewer should not invent the user's preferences to make a scenario produce a desired result.

The other original-source books and 2019 correspondence remain separate evidence questions. This paper supplies no authority to certify them. This patch deliberately does not assign `source_questions: 0` or promote stages for `credible_source` or any other record whose original source was not independently read in this review.

## 9. Bounded next implementation task

Keep the inherited record's identity and semantics as the documented legacy adapter. If full 2007 coverage is required, implement a separately named, explicitly scoped AATS-oriented variant or adapter with its own contracts and tests. Silently rewriting `practical_reasoning1` would violate the preservation decision and change existing applications' burdens and defeat conditions.

The minimum design work is now specific:

1. Represent the deciding agent, its action component, the joint action, initial/resulting states, changed goal proposition and agent-relative transition value. The same goal token in different contexts must not leak a value assessment across them.
2. Declare whether a validated finite model discharges CQ12, CQ14 and CQ15 or whether explicit challengeable premises/defaults do so. Test this choice; do not treat a positive case as implicit validation of the model universe.
3. Distinguish CQ17's other-agent choice from CQ2's alleged transition. Supply positive cooperation evidence and a different-agent/wrong-joint-action control.
4. Generate alternative and negative AS1 arguments for CQ5–9 and CQ11, and specify when actions are incompatible. Do not assume that every pair of action recommendations conflicts.
5. Treat value preferences as explicit case inputs, allowing incomplete preferences to leave a choice unresolved. Preserve positive CQ10 support and distinct justifications. The paper's discussion does not authorize an invented universal value ordering.
6. Test the farmer example beyond the current positive abstraction: Arg6's provisional support; the CQ17 challenge leading to the alternative state; positive CQ10; Arg8 under the paper's stated beliefs/preferences; and a changed preference that changes an adverse-value outcome. Include different-context controls and a second run/export/import check.

These requirements are sufficient to start a targeted implementation; another general review round is not needed. They are not already implemented by the present patch and are not hidden behind a successful test count. The current completed deliverable is the independent source review, corrected provenance, explicit gap classification, robust ledger enforcement and passing focused regression evidence.

## 10. Reproducible review artifacts

The companion `atkinson-2007-decisions.json` records the source fingerprint, fields compared, each CQ's page locators, mapping, disposition, implementation limits and example references. Executable fixtures live in `tools/vv/cases/atkinson_2007.yml` and the corrected `practical_reasoning1/book` entry. The argument-local and schema regressions live in `tools/vv/internal/vvlib/atkinson_test.go` and `source_test.go`.

The supplied paper itself is not copied into the repository by this patch. The review's archive may carry the user's original attachment and rendered evidence separately. Page locators in this report refer to that fingerprinted source.
