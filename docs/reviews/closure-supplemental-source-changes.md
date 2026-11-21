# Additional source corrections and explicit compendium forms

Author: integration reviewer, 2026-11-21. Independent acceptance is separate.

These changes follow the supplied *Argumentation Schemes* (2008), printed
pp.310,336,339–340,343. They do not certify the separately cited 1995 work,
the slippery-slope edition cited as1999 by the inherited file, or the
unavailable relevant passage of *Witness Testimony Evidence* (2008).

## A witness's account, not every assertion by that witness

Entry3 CQ4 attributes relevant bias to the account. The old `biased(W)`
exception could defeat an unrelated account from the same witness. The
replacement is `witness_account_biased(W,S)`, where S identifies the situated
proposition and the witness's account in this evidential application. Two
independent accounts about differently situated propositions remain distinct.
Repeated episodes about literally the same proposition require a distinct
account/proposition identity supplied by the application; this change does
not invent a complete testimony-episode ontology. Generic `biased/1` is still
used by other schemes and is not automatically expanded into all accounts.
Sincere belief remains a declared interpretation of the truth-telling premise,
not a guarantee that a sincere witness is factually correct.

## Reduced plausibility is not falsity

Entry41 distinguishes changes in plausibility from an assertion's truth.
The inherited negative record `ethotic2` previously concluded `¬S`. That
is too strong for the available printed formulation. It now concludes
`reduced_plausibility_from_ethos(P,S)` and retains its original-source hold.
The correction deliberately removes an invalid negative inference instead of
merely describing it as a legacy policy. The old negative book fixture is
updated to the weaker proposition. Separate controls show that independent
support for S can survive the credibility assessment.

The positive inherited `ethotic1` still supplies defeasible support for S.
That existing strength policy remains qualified pending its original-source
comparison; it is not described as a numerical plausibility calculation.

Two new records, `ethotic_positive_compendium` and
`ethotic_negative_compendium`, make the compendium's qualitative assessment
explicit. Each binds speaker, situated assertion and dialogue. The character
premise and assertion need support. Character relevance and the warrant for
the *claimed qualitative degree and direction* of presumption are separately
challengeable defaults. CQ3 therefore has an actual evidential route rather
than being relabelled as an issue's numerical proof standard. No numeric
probability, automatic weight update, S or ¬S is concluded by these records.
These are static evidential interpretations: an application supplies a
supported assessment or challenges its default. A dialogue protocol may add
the separately documented challenge/answer mechanism.

## The complete entry49 argument

The two old slippery-slope records construct consequence information. Neither
contains the complete entry49 first-step, recursive-chain and bad-outcome
argument with its normative conclusion. The new
`slippery_slope_compendium` includes all three and concludes the explicit
recommendation `should_not_bring_about(A)`. This is not a semantic rename of
`¬should_be_performed(A)` and no general deontic equivalence is asserted.
Applications may compare recommendations through an explicitly designed
decision policy; this record alone does not add a global exclusivity issue.

The finite sequence is an identified evidential witness in
`plausible_slope_chain(A,S,Z)`, scoped to starting proposal, sequence and final
outcome. Its template requires the actual intervening steps and plausible
links, not merely an unspecified allegation that a slope exists. CQ1 checks
that ordinary premise; CQ2 can challenge its completeness; CQ3 can challenge
the evidential adequacy of its weakest links. The engine evaluates the
supplied argument. It does not enumerate an external world, verify causal
truth or prove that a bare sequence identifier describes a valid causal chain.
This explicit aggregate interpretation is comparable to the library's other
finite-set witness forms and is subject to independent review.

This is entry49, not the existing `full_slippery_slope` record from entry53.
Both remain separately identified. The historical1992/1999 edition discrepancy
is preserved as an unresolved original-source matter.

## The actor/action form of entry55.1

`established_rule_actor` makes the actor, action, rule and application case
explicit. Both the established action-type rule and its conditional
obligation warrant require support. Relevant conflicting/overriding rules
and case-specific excuses are evidence-backed exceptions. The conclusion is
`must_do(P,A)`, shared with the existing obligation vocabulary. A must-duty
does not automatically become a merely practical `should_do` recommendation.
Action terms must identify their situated occurrence; the explicit case
argument additionally scopes the rule application and objections.

This form adds source coverage that the inherited rule-content record could
not provide. It does not silently retitle that old record or claim its1995
source was verified. Missing warrant, wrong action, other actor and other-case
controls distinguish the implemented interpretation.

## Migration and verification

Both changed inherited forms require manual semantic migration. A generic
bias fact needs the actual account; a negative proposition must be traced to
its ethotic argument before replacing that argument's conclusion. The tool
must never rewrite every negation in a graph into a credibility assessment.
Existing provenance and source holds remain. The four new compendium records
remain below interpretation approval until a reviewer who did not author
these changes examines the source, formulas and examples.

The28 hand-authored cases in `tools/vv/cases/closure_supplemental.yml` cover
qualitative versus truth conclusions, question/answer identity, alternate
dialogue or witness context, complete versus missing slope evidence, and
actor/action/rule/case isolation. Each is run through API and CLI, with
additional generated premise-deletion and exception controls where marked.
Execution logs and the independent review are recorded separately; the
presence of this document itself is not evidence that execution passed.
