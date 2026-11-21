# Independent acceptance of the source-family repairs

Reviewer: integration reviewer (`closure/root-source-review`), 2026-11-21.
Author of reviewed semantic changes: closure scope implementation reviewer,
commit `7eb92502087c5761d22ed5fcf938db65cc567816`. This reviewer did not author
those source formulas. Review used the supplied compendium pp.314–315,318,
323–326, the exact changed formulas and language, all27 new cases, and the
current CLI oracle that checks adverse results as well as positive labels.

## Accepted source interpretations

| Record | Independent judgment | Discriminating evidence |
|---|---|---|
| `popular_practice` | The full conditional warrant now stands as a premise. Case C binds the popularity, warrant, grounds and conclusion. The source's current-case conclusion does not allow popularity in one setting and a justification from another to be combined. CQ2 is an explicit challengeable static default; this is a presumptive acceptability argument, not proof that popularity entails moral correctness. | The restaurant/court cases separate warrants and challenged grounds. The existing book example has the same actual case throughout. |
| `analogy` | The source CQ3 requires another similar case with a contrary truth value, without the additional more-on-point condition. The new evidence-backed hook correctly captures that question under the declared static profile. A is a situated target proposition, and C identifies its base comparison within that application. Reusing bare, unqualified names across different targets is outside this contract. The accessed2006 text corroborates the supplied2008 formulation; no facsimile review of the whole2006 book is claimed. | Relevant counterexample defeats; wrong claim/base does not; a bare question does not supply evidence; the source's negative truth-value branch works through explicit negation. |
| `from_antimodel` | The source recommends refraining. An explicit `should_refrain(B,X)` states that meaning. Denying a positive obligation is not equivalent to obliging abstention. The patch correctly adds no global issue or automatic bridge that would collapse those meanings. It is accepted as a recommendation schema, not a general deontic solver. | Explicit abstention follows from the anti-model premises; an isolated `¬should_do` fact does not produce it; another recipient's goal cannot satisfy the original case. |
| `alternative_normative` | The printed inference requires exclusive requirement alternatives. Stating that warrant explicitly is a justified executable interpretation of the ambiguous either/or wording. At most one alternative can be required; the affirmative requirement of A therefore supports not-required(B). This is not a prohibition of B. | The exclusive case succeeds; inclusive wording alone does not; missing affirmative requirement blocks the inference; no must-not conclusion appears. |

These four records are independently accepted at `example-passed`, with exact
semantic fingerprints and existing book/adverse cases in the registry. The
remaining practical records below are not silently promoted by this review.

## Practical context changes accepted; dialogue review separate

`value_based_practical_reasoning` now keeps the evaluating agent in values,
alternatives, efficiency, value comparison and negative consequences.
`practical_inference` uses the same agent-relative efficiency relation and
scopes incompatible consequences. `two_person_practical` keeps the advised
actor distinct from the adviser when evaluating alternatives and further
required actions. These identities follow the source's my-goal/available-to-X
roles. Compound action and goal terms distinguish repeated deliberation
contexts for the same agent.

The mixed-agent and mixed-case controls demonstrate these corrections. Their
canonical schemas still implement static evidential support, whereas the
book's practical narrative also specifies question-induced suspension. Full
acceptance of the intended dialogue profile needs the separately reviewed
snapshot adapter and its challenge/answer cases. It is not inferred from
the scope fixes alone.

## Execution and migration reconciliation

The integration run evaluates all27 new family cases through both API and
the strengthened CLI round-trip checker:54 registered family checks pass.
The source gate exposed four stale destinations in historical migration rows.
The integration reviewer repaired those destinations and converted old
spelling-only mappings into manual semantic changes where an actor or case
must now be supplied. A subsequent source/migration gate passes725/725.
These migration bookkeeping repairs do not change the reviewed formulas.

The original failures remain in `supplemental-audit/corrected.log` and
`source-corrected.log`; the passing source evidence is `source-final.log`.
The complete final candidate still must pass its clean dual-toolchain gates.
Source approval is grounded in the independent reasoning and cases above,
not in the absence of errors in a fingerprint checker.
