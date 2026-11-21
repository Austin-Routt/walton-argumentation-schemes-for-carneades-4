# Authored closure repairs for seven held source interpretations

**Author:** ChatGPT closure scope implementation author, 2026-11-21. **Approval status:** no self-approvals. The formulas, evidence and proposed contracts below require independent acceptance by the integration reviewer. They address actual scope, conditional and polarity gaps identified in the 50-row independent baseline audit; earlier unrelated approvals are retained.

## Changes and source reasoning

| Record | Change | Source and deliberate boundary |
|---|---|---|
| `popular_practice` | Add case C to popularity, conditional warrant, justified grounds and acceptability. Restore the complete conditional in `to_think_acceptable(C,A)`. Scheme parameters are `[A,C]`; statement arguments are `(C,A)`. | Printed p.314 says acceptable in this case. Popularity in one situation and a warrant accepted in another cannot jointly justify an acceptability claim. CQ2 remains a challengeable static default about justified grounds. |
| `analogy` | Replace the stronger `more_on_point(A,C)` exception with `contrary_similar_case(A,C)`. | Printed p.315 and the accessed 2006 pp.96–97 require another similar case with opposite truth value; they do not require that case to be more on point. The new hook asserts supported counterexample evidence. Mere questioning does not assert it. C identifies the base case relative to the current analogy context; A is the claim in that context. Applications using multiple target contexts must use context-qualified claims/base identifiers or separate graphs. |
| `from_antimodel` | Conclude explicit `should_refrain(B,X)` instead of `¬should_do(B,X)`. | Printed p.315 recommends not doing the action. This is not identical to a denial of a positive obligation. No general deontic axiom or global issue joins the new statement to the old negation pair; doing so could incorrectly make compatible negative obligations exclusive. |
| `alternative_normative` | Require `exclusive_requirement_alternatives(A,B)` and preserve `¬is_required(B)`. | Printed p.318’s inference uses exclusive alternatives. Inclusive either/or information cannot license elimination by affirming one alternative. The result remains not-required, not forbidden or must-not. |
| `value_based_practical_reasoning` | Bind P in the values premise and all alternative/efficiency/value-ranking/consequence hooks. Reuse `alternative_means(P,G,A)` where the meaning already agrees. | Printed p.324 refers to my goal, values and alternatives. Two actors sharing goal/action names cannot share options or value preferences automatically. Repeated situations for the same actor use contextual goal/action terms such as `survival(case1)` and `treatment(case1)`. |
| `practical_inference` | Use the same agent-scoped efficiency hook and add P to incompatible consequences. | Printed p.323 and the practical narrative identify the deliberating actor. The same missing-agent defect existed here; preserving a second unscoped version would leave a known practical cross-agent failure. |
| `two_person_practical` | Scope alternatives and additionally required actions to actor A. Preserve adviser B’s distinct communication role. | Printed pp.325–326 ask about means available to the advised actor and actions that actor must perform. The adviser’s or another actor’s circumstances cannot defeat this argument. Contextual goal/action terms distinguish repeated deliberations. |

The three practical records still use the declared static evidence policy in their canonical formulas. These identity repairs do not by themselves implement printed p.96’s temporary suspension on an appropriate unanswered question. That is a separate finite dialogue snapshot adapter contract. The source ledger therefore retains a hold pending independent review of the complete intended profile.

## Compatibility decisions

Every changed atom appears in `tools/migration_r2.yml`. Added cases/actors cannot be recovered by a string rename. The source change for analogy, exclusivity and anti-model polarity is also a manual semantic migration: a saved negated judgment under the former meaning cannot be converted merely by preserving its sign.

The anti-model abstention conclusion is intentionally distinct from both `should_do` and its explicit negation. This patch adds no global issue scheme and does not invalidate all existing approvals by changing their shared issue semantics. An application that selects action options must state its choice model explicitly. It must not assume that refusal of a duty is the same proposition as a duty to refrain.

The comparison with 2006 uses the indexed text described in `closure-source-adjudication.md`, source S1. It is not claimed to be visual facsimile verification. The current work does not change that source-access description.

## Evidence and acceptance

`tools/vv/cases/closure_family.yml` contains explicit positive and adverse cases for case identity, source counterexamples, negative versus positive obligation, inclusive alternatives, agent-specific values/options/costs and contextual repeated actions. Expected facts are written independently from the source roles; the tests are not generated from the schema under test.

Existing book and phase cases were migrated to preserve their intended scenarios. The Python validator’s two hand-written actor-role oracles were updated to the explicitly identified agent. The initial focused run caught two new practical controls using the wrong input predicate name (`means_to_goal` instead of the actual `means_to_end`); those were fixture defects, repaired with the intended source premise and explicit argument-existence assertions. All original failures are retained in the review evidence. Integration must report the final corrected results, not just the first run’s aggregate.

The author grants no source approval. The integration reviewer should assess each new formula and adverse expectation, rerun final acceptance with the corrected CLI expectation harness, and promote only the profiles actually reviewed.
