# Source contracts for compendium entries 28–60

This document records decisions made by the independent source review of
2026-11-21. Page numbers refer to Walton, Reed and Macagno (2008),
*Argumentation Schemes*. These are executable interpretations of the supplied
book, not publisher-confirmed errata or certifications of other works cited by
the book. The row-by-row approval ledger distinguishes those questions.

## Roles that a graph must preserve

The names below designate particular case entities. A convenient short spelling
does not permit a graph to reuse one entity for two different roles.

| Record | Required identity and interpretation |
| --- | --- |
| `abductive_forward` | `G` is a set of diagrams, `G1` is its strongest diagram, and `C1` is that diagram's conclusion. The premises explicitly link the three. A diagram is not its conclusion. |
| `abductive_action_to_character` | `A` is the person, `X` the performed action, and `Q` the particular quality. The typed `character_quality(A,Q)` supports character prediction; its additional projection `has_property(A,Q)` can feed retroduction. An ordinary property alone does not assert a character trait. |
| `abductive_character_to_action` | `character_action(A,Q)` is the conditional prediction that future conduct is likely to fit Q. It neither asserts that a particular future action occurs nor identifies the agent of a past event. |
| `retroductive` | `E` is the same event in the candidate-agent premise, event-quality premise, conclusion and critical questions. Evidence about another event cannot be substituted. The reference class concerns this event and quality. |
| `memory` | The original-belief undercutters identify both recaller P and proposition M. One person's defective belief history does not defeat another person's independent support. |
| `rule_physical_world_1` | The graph supplies both the printed inclusion premise and the antecedent of the synonymy warrant, `referentially_coextensive(W,D)`. Inclusion alone does not establish coextension. The library does not infer coextension from inclusion. |

For all shared statement predicates, an application must supply ground terms
whose intended entities remain identical across the schemes that share them.
In particular, action terms used with `should_do(P,A)` describe the same action
by the same P. Negative-consequences reasoning supports `¬should_do(P,A)`,
which is evaluated against `should_do(P,A)` in the library's negation issue.
This is a policy for weighing recommendations, not a proof that the empirical
premises or value judgements supplied by the application are true.

## Exceptional cases: conditional backing, not three interchangeable routes

Entry 56 (p. 344) asks a conditional second question about unrecognized cases
and a conditional third question about borderline cases. The selected policy
for the aggregate default `exception_backed(C)` is:

1. A recognized exception must actually apply to case C. If C is unrecognized,
   evidence must establish why the governing rule does not apply to C.
2. If C is borderline, comparable cases must additionally justify its treatment.
   Recognition alone does not discharge this comparison requirement.
3. All classifications, evidence and comparisons refer to the same case and
   governing rule. Use a ground case term such as `case(hearing_17,rule_3)`.
   A bare person name is unsuitable when more than one case or rule is involved.

Let R mean recognition as an applicable exception, B mean borderline status,
E mean supported non-applicability of the rule, and K mean comparable cases
supporting the claimed treatment. After the aggregate has been challenged,
an explicit witness may establish it under:

`(R or (not R and E)) and (not B or (B and K))`.

The negative statuses in this formula require evidence or an explicit case
stipulation. Absence of a fact is not its negation. Unknown recognition or
borderline status cannot automatically complete a challenged witness. A
recognized, non-borderline exception; an evidenced, unrecognized,
non-borderline exception; and the corresponding two borderline forms with
comparison support are the four complete routes.

The core scheme remains defeasible: an undisputed default can apply without
an independently produced witness. A reviewer who requests the grounds adds
`exception_backed(C)` to `questioned:`. It then needs an explicit stipulation
or supporting argument, as contract 5 requires. Directly stipulating the
aggregate asserts the whole conditional condition above. Merely having one
of R, E, or K does not justify stipulating it.

`tools/vv/cases/independent_exceptional.yml` gives supporting rules for the
four witness routes. It also checks missing evidence, missing comparisons,
wrong-case comparisons and unsupported challenged defaults. These support
rules are application policies, not additional schemes claimed to be printed
in the book. The named evidence predicates assert evaluated evidence; the
engine does not inspect or authenticate documents behind them.

## Other aggregate and context boundaries

For inherited positive/negative consequences (33.1–33.2), the default
`consequence_likelihood_adequate(A,G)` makes CQ1 challengeable for the specific
action and outcome. Its meaning is that the assessed likelihood is adequate
for the strength of this argument. It does not require a fixed numerical
probability or silently strengthen every consequence claim to "strongly
likely". Evidence that the consequence follows (CQ2) remains the ordinary
`brings_about(A,G)` premise. An application must justify any quantitative
threshold it adopts; the generic scheme imposes none.

The inherited `established_rule` form now has exceptions for conflicting or
overriding rules and for exceptional/excusing cases, scoped to (R,C), based on
55.1 CQ2–CQ3. C must identify the particular conclusion/application being
challenged. The formulation remains upstream's rule-and-conclusion schema,
not the compendium's actor/action schema. These adaptations close the
available-book CQ gaps; they do not certify the unavailable 1995 cited work.

Several printed schemas summarize a sequence or a dialogue in one premise.
Such a premise is a proposition requiring justification, not a certificate
that the engine searched the underlying sequence or dialogue.

* `sequence`, `reapplication_sequence`, `precedent_series`, `vague_property`,
  and `leads_to_next` require a traceable sequence with the stated endpoints,
  property or rule, and every required intermediate link. An application
  should retain that witness with its case. The generic library does not
  enumerate arbitrary paths from a tag asserting that a path exists.
* Anonymous "you", "I", "this case", governing rule, institution or language
  in a schema denotes one fixed context. Independent contexts must be
  represented with distinct ground case/action/event terms or separate graphs
  when a predicate has no parameter for the omitted context. This limit is
  explicit; it is not permission to regard distinct contexts as identical.
* `counts_as(D,W)` in 54.2 is relative to the selected institution. The graph
  must not reuse D and W to combine incompatible institutions. `means_same`
  in 54.3 similarly needs one selected language or explicitly qualified terms.
* For 55.3, `must_not_do(A,X)` expresses a prohibition; `¬must_do(A,X)` merely
  denies an obligation. They are not interchangeable. The printed obligation
  and prohibition pair is reversed. The library's two variants use the
  source-supported reading of the warrant: obligatory → must do; prohibited
  → must not do. A conflict between an obligation and a prohibition needs
  an application issue/policy; the generic negation issue does not create it.
* Generic ad hominem (42) uses the printed strong conclusion, rejection of the
  argument, with a defeasible relevance assumption and an outweighing-evidence
  exception. It does not silently change the numerical weight of the target
  argument. Circumstantial and bias forms retain their weaker conclusions.

These contracts explain the supported interpretation of aggregate premises.
They do not fill an unmodelled critical question, validate a missing source,
or remove any record-specific blocker recorded by the independent audit.

## Migration

The changes in entries 31.2–31.5 and 60 require actual diagram, action, event,
quality or recaller identities. A name replacement cannot infer those facts.
The migration table therefore reports these changes for manual completion.
New ordinary premises can include a `trigger` identifying an existing premise
whose use requires a review, while leaving `old` empty because it remains in
the scheme. The migrator reports the need for coextension without adding it.
After changing a scheme's parameters or questions, reconstruct persisted
instances and their undercutters from the reviewed case facts.

## Cross-review clarification: predictive character evidence

For entry 31.4, `not_character_action(A,Q)` means relevant counterevidence sufficient to undermine this qualitative trait-based prediction. One isolated atypical act is not automatically a probability-theoretic refutation of a likely tendency. The evidence adequacy judgment is external; the library provides a scoped defeater hook. This explicit interpretation was selected in the independent integration review after inspecting printed p.330.
