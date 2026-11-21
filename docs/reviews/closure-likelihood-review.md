# Closure review: entry 33.3 likelihood

## Source decision

Walton, Reed and Macagno (2008), printed p.332, gives the shared CQ1 of entry
33 as a question about the strength of the likelihood of the consequence.
CQ2 on p.333 separately asks whether the evidence adequately supports the
asserted strength. The printed text supplies no fixed numerical threshold and
does not require every consequence to be strongly likely.

The implementer's criticism is accepted. The earlier `outcome_likely/2`
language imposed an unsupported fixed qualitative threshold. Entry 33.3 now
uses `consequence_likelihood_adequate(A,B)`, already used for the same CQ1 in
33.1 and 33.2. It is a challengeable application judgment about adequacy for
the argument's asserted strength. It is not a probability calculator, does not
encode CQ2's evidence by itself, and does not silently certify a separately
cited earlier source.

The agent-conditioned consequence and bad-outcome premises remain unchanged.
The selected bring-about branch still concludes `¬should_do(P,A)`. The printed
parenthetical don't-bring-about branch is not thereby added.

## Migration and evidence

This semantic change is deliberately a manual `replaced` migration. A denial
that an outcome is strongly likely need not establish that its likelihood is
inadequate for a modest defeasible claim. Conversely, asserting a threshold
does not automatically assess a claim's contextual strength. The migration
tool lists the change but leaves old positive, negative and questioned terms
untouched for explicit reassessment. Historic descriptions of the former hook
remain as history.

`closure_consequence_likelihood.yml` exercises two schemes in one graph,
the shared challenge, a supported answer, outcome-specific isolation, and the
two distinctions from old threshold judgments. Existing worked usage and
phase-5 examples use the current adequacy question.

## Review independence

This patch's author proposes the repair and the specified examples. The old
approval fingerprint is intentionally stale until another reviewer checks the
source judgment and these changes. This document grants no approval to its
own code or source changes.
