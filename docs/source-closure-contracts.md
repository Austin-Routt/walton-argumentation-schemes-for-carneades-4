# Source closure contracts: entries 9.2, 10.2 and 16


Current approval status (2026-11-21): implementation-stage statements below saying review is pending are historical. The new compendium variants now have non-author acceptance in `reviews/closure-current-register.md` and individual acceptance reports. Inherited-original holds remain separate; the contracts are retained.

These contracts interpret the supplied Walton, Reed and Macagno (2008)
compendium. Its bibliographical predecessors are not claimed as independently
inspected. The implementation author proposes the new forms and tests;
another reviewer must approve them before their source stages advance.

## Entry 9.2: distinct inclusion and degree alternatives (p.316)

The existing `part_in_whole` record remains the positive species branch.
`part_in_whole_physical` separately represents the physical-part branch.
Physical parts and taxonomic species use different predicates, and neither
relation is silently inferred from the other. These are the source's
presumptive argument patterns. The presence of a part relation alone does
not certify that every real-world property transfers to every part; the
application must select an appropriate scheme and justify its premises.

The parenthetical lesser-degree conclusions are separately executable in
`part_in_whole_species_degree`, `part_in_whole_physical_degree`, and
`part_in_whole_smaller_degree`. The last record additionally requires a
supported claim that the part is smaller in the relevant measure. All three
require `degree_comparable(Y,X,A)`: the same contextually identified property
and comparison scale must apply to the part and whole. This is an explicit
application precondition for interpreting "less A", not a numbered source CQ.

Their conclusions are qualitative `less_property_than(Y,A,X)` arguments,
alongside property possession. They do not compute numerical magnitudes,
infer that the property is absent, or turn lower degree into a defeater of
property possession. A negative argument against the qualitative comparison
can oppose it through the ordinary negation issue. The source supplies no
separate numbered question list for entry 9.2.

Separate records preserve the alternative readings. Loading all variants may
produce several arguments for the same property or comparison when all their
premises hold. This is visible provenance, not an additional independent
measurement; applications should enable the alternatives they intend to use.

## Entry 10.2: arbitrary finite partitions and conditional direction (p.317)

The printed source uses x1 through xn. Its second premise says "only if".
Read as a material conditional, that is a necessary condition: the whole
having A entails that some part has A. A positive part does not establish the
whole's property by that necessary condition alone.

The historical `whole_into_parts` identifier and parameter order remain,
but its positive three-part form now also requires the ordinary premise
`selected_part_sufficient_for_whole(X,X1,A)`. This states the sufficient
direction for the same whole, selected part and property. It is never supplied
as a default or inferred from the printed only-if premise. Existing graphs
need this explicitly justified witness; the migration table flags it for
manual work. The record is a declared application adaptation, not a claim
that the book says if where it prints only-if. The following general forms
also support arbitrary finite partitions and the negative branch.

### Positive selected-part form

`whole_into_parts_selected(X,S,Y,A)` requires:

1. `whole_partition(X,S)`: S identifies the complete finite, nonempty set of
   parts of X for this particular application.
2. `partition_member(Y,S)`: the selected Y belongs to that same set.
3. `some_part_sufficient_for_whole(X,S,A)`: this application supplies the
   sufficient direction of the property warrant.
4. `has_property(Y,A)`.

It concludes `has_property(X,A)`. The sufficient warrant is deliberately an
ordinary premise, not an automatically assumed answer and not an inference
from the printed necessary condition. This is a declared strengthening of
the application warrant needed for that positive conclusion; it is not a
claim that "only if" in the book means "if". Any selected member can be the
witness. The record has no hard-coded three-part arity.

### Negative only-if form

`whole_into_parts_negative(X,S,A)` requires the complete partition,
`whole_requires_property_part(X,S,A)`, and
`no_partition_part_property(X,S,A)`. The latter is an explicit supported
universal negative over the same complete part set. It concludes
`¬has_property(X,A)`, retaining the source's necessary direction.

No rule derives that universal negative from absent facts, from one negative
part, from an incomplete enumeration, or from a different partition. The
whole, partition and property identifiers must all agree. Likewise, a
stipulated sufficient condition does not substitute for the necessary
condition in this branch.

The part set identifier is an opaque, application-qualified term. The engine
does not infer its cardinality or certify the external completeness of the
asserted partition. An application can use a two-part set, a seven-part set,
or a larger finite set without changing the scheme; it supplies membership
and complete-set evidence. This is the same explicit aggregate-evidence
style already used for `all_parts_property/2`. Runtime quotas remain in force.

These two general forms are semantic additions, not automatic replacements
for a saved three-part graph. Both positive forms require an explicit sufficient warrant. The migration table calls for review of the
partition, membership and directional warrant. It never manufactures those
premises.

## Entry 16: static evidence burden for the vagueness challenge (pp.319-320)

The implementer's `certainty_no_need(A,F)` wording correctly asks whether
the degree of vagueness is a problem for the dialogue. The historical name is
retained for migration compatibility; it does not assert certainty.

For this library's static evidence graph, the proponent of the rejection
must support the conditional warrant, the context's precision requirement,
and the property-specific inadequacy. Missing any required premise prevents
construction of that rejection argument. The exception asserts affirmative
evidence that the particular vagueness is no problem in this dialogue.
Supported exception evidence defeats that argument; merely asking whether
the exception is true does not establish it. A challenge about another
argument or another property cannot supply this exception.

This is an explicit static burden policy. The displayed CQ3 asks for a
reason but does not itself prescribe a dialogue protocol in which uttering
the question automatically suspends an already supported argument. No
automatic question-utterance suspension is claimed. Applications with that
dialogue policy must translate their evidence state accordingly or use a
separately specified dialogue controller.

## Evidence and approval boundary

`closure_parts.yml` exercises the new alternatives, missing premises,
relational/scale/partition isolation, arbitrary selected members, both
conditional directions, and the refusal of closed-world universal negatives.
`closure_vagueness.yml` exercises the current wording and static burden.
Existing book examples remain regression checks for the legacy forms.
These tests are proposed independently checkable evidence; this document does
not grant its author permission to approve the authored changes.
