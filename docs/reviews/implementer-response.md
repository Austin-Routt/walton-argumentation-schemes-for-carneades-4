# Implementer response to the two independent reviews

2026-11-21. Branch `review/implementer-response`, based on the acceptance
candidate `fccec4f` (branch `review/independent-acceptance`), which contains
master `96daf6c`.

Inputs: the first review's package (report, acceptance candidate `fccec4f`,
record dispositions, source approvals, Atkinson 2007 decisions); the second
review's report "Independent review before release 1", its merged pull request
#13 and its open pull request #14; the compendium (Walton, Reed and Macagno
2008) and Atkinson and Bench-Capon (2007), both supplied by the owner. Every
claim below was checked against the repository, the printed pages or a run of
the engines; nothing is taken from either report on trust.

## 1. Are the reviews accurate?

Both reports describe the repository and the status report correctly. Their
factual claims about the sources agree with the printed pages wherever I
checked them. They differ in judgment, not in facts (section 2).

**Engine claims of the first review**, rerun on master `96daf6c` with the
probes kept in the review notes:

| Claim | On master | On `fccec4f` |
|---|---|---|
| One fact cannot satisfy two premises of one argument | reproduced | fixed |
| A declared issue can contradict stipulated positions without refusal | reproduced | fixed |
| Malformed nested terms are accepted | reproduced | fixed (strict parser) |
| Cancellation is lost in about half of the runs | reproduced | fixed |
| A default does not yield to a stipulated issue position | reproduced | fixed |
| An authored argument can evade its scheme's exceptions | reproduced | fixed |
| A malformed issue pattern leaves the graph half changed | reproduced | fixed (atomic publication) |
| Memory grows without bound on large joins | reproduced | fixed (streaming match, quotas) |
| A defeated default seeds further defaults | same behaviour on both; it follows contracts 4 and 5 | — |

Seven of the eight are real defects of the engine the old R1 proposal pins.
That settles the main disagreement: `8c092b6` should not be published.

**Defects in the acceptance candidate itself.** Its new work budgets were
tight enough to refuse ordinary inputs:

| Input | master | `fccec4f` | this branch |
|---|---|---|---|
| Whole library with 500 / 800 / 1,200 facts | 0.43 / 1.5 / 5.8 s | refused above ~700 statements ("issue matching exhausted its candidate-visit budget") | 0.14 / 0.30 / 0.71 s |
| 400 arguments for `q(kN)`, 200 for `¬q(kN)` (1,200 statements, 200 issues) | 8.4 s | refused, exit 4 | 0.08 s, same labels and issues as master |
| Four-premise join over 120 facts with one match | 1.9 s | refused at 1,000,000 match visits | 1.6 s, `t(c1,c2,c3)` in |

On the 1,200-fact library graph master builds 12 issues and leaves both
`is_required(k1)` and `¬is_required(k1)` in. The candidate's scan, given a
large budget, and this branch's lookup build the same 14 issues and identical
labels for all 2,199 statements: master missed two issues, as the first
review's "complete issue closure" finding says.

Two minor points: the authored-exception check compared a normalized key with
an unnormalized one (now both normalized); and the candidate's change of `''`
to `”` in a parser comment is gofmt's doc-comment rewriting, not an error, so
it stays.

## 2. Where the reviews agree and differ, and my position

They agree on: the agent variables of 22.2 and 22.3; 6.1/6.2 sharing; 9.2's
degree parenthesis not being an exception; 26.2's instances; "sees", not
"sets", in 23; entry 4's CQ2 wording; 12.2's converse wording; restoring
`biased/1`; sharing 18.1/18.2 and 59.1/60; and modelling the deferred
questions from the compendium (the second review recommends that the owner
allow it, the first did it). I agree with all of these.

| Topic | Second review | First review | Position |
|---|---|---|---|
| Release | publish `8c092b6` | superseded | Superseded: seven engine defects reproduce on it. |
| Approvals | 90 (95 with #14) | 70; the overlap is 65 | Keep 70. The 25 the first review holds are held for scope, variant or burden-policy reasons, which is the safe side; a reviewer can promote them later. |
| 54.1 | approves the two-line printed form | adds `referentially_coextensive(W,D)` | First review. The printed premise "a W is a D" covers class inclusion; without the bridge, "whale is a mammal" would make "whale" mean "mammal" by default (p. 342). |
| 31.3–31.5 | `has_property` for 31.4's premise | typed `character_quality(A,Q)`; 31.3 concludes both | The reconciled form: it keeps the type and chains to 31.5's trait premise. |
| Atkinson 2007 | counts the same 17 CQs | same | Agreed on the facts (pp. 858–862, 866): CQ10 supports the action; CQ5–7 are narrowed to "more effectively"; CQ11, 12, 14, 15 and 17 have no mapping. |

**Approvals the first review both wrote and approved.** D14 asks that someone
who did not write a change review it. Eight approved records were changed and
approved by the first review alone. I did not write those changes, so I
checked each against the printed page:

| Record | Entry, page | Change | Finding |
|---|---|---|---|
| `abductive_forward` | 31.2, pp. 329–330 | the strongest diagram G1 and its conclusion C1 linked | concur: premise 3 and the conclusion name the same G_i and C_i |
| `abductive_action_to_character` | 31.3, p. 330 | agent, action and quality explicit | concur; `has_property` is a sharing addition |
| `abductive_character_to_action` | 31.4, p. 330 | quality explicit; CQ3 as scoped counterevidence | concur |
| `retroductive` | 31.5, p. 331 | the event E made explicit; CQ4 as reference class | concur: the factual, character and trait premises share the event |
| `from_model` | 6.3, p. 315 | premise 3 conditional on P carrying out A | concur |
| `rule_physical_world_1` | 54.1, p. 342 | explicit coextension premise | concur (above) |
| `exceptional` | 56, p. 344 | CQ2 only if not recognized; CQ3 for borderline cases | concur: the printed questions are conditional |
| `memory` | 60, p. 346 | CQ1, CQ2 scoped to the recaller | concur |

Each ledger row now says so. The other six changed and approved records were
written by the second review (#13, #14) and approved by the first, so they
already had two people.

## 3. What this branch changes

Engine, 4.3.1+walton.3:

* Issue construction looks positions up by key when the remaining patterns
  are ground under the bindings of the first, as the negation scheme's are.
  Statements with variables, and enumerations, keep the scan. A unit test
  compares the two on an irregular graph of some 600 statements.
* `MAXISSUESTEPS` and `MAXMATCHSTEPS` are 50,000,000: safety valves, not
  working limits (contract 3).
* `terms.Substitute` marks every variable of a binding chain, so a chain that
  ends in a variable bound to itself stops. The lookup reached this upstream
  defect through an AIF example with variables in its statement keys; it hung
  the engine's own tests.
* Two named checks: `large-issue-closure-completes`, `failing-join-completes`.
  8,230 checks.

Library and documents:

* `certainty_no_need/2` (entry 16, CQ3): the wording of #14; the record stays
  held, since I wrote this change.
* The unused `smaller_part_property/3` template is removed (its exception was
  removed from 9.2 earlier).
* `docs/using.md`: the first- and second-person schemes need one proponent and
  one respondent per graph (#14, adapted to this library).
* `docs/reviews.md`, finding 14: five commits, not three, need splitting or a
  new title before they go upstream (#14; I checked each with `git show
  --stat`).

Of #14's six wordings, four are already in the candidate (`situation_necessity`,
`doubting_general`, `relation_converse_relation`, `certainty_no_need` here),
and two belong to templates the candidate replaced (`alternative_action_sufficient/2`
in 22.3, `act_on_presumption/1` in 26.2). Its 26.2 instances are in the
candidate. #14 can be closed as superseded once this branch is merged.

## 4. Left open, and whose decision it is

* **Reviewer:** the engine changes of this branch, and the entry 16 wording,
  need review by someone other than me.
* **Reviewer:** 33.3 assumes `outcome_likely(A,B)`, "strongly likely";
  33.1–33.2 assume `consequence_likelihood_adequate(A,G)`, "adequate for the
  strength of this argument". The same CQ1 of entry 33 has two statements, and
  the stronger wording asks more than the question does. I recommend 33.3 take
  the shared statement; I did not make the change, because it would make an
  approved record stale and I cannot approve it myself.
* **Owner:** Atkinson 2007. `practical_reasoning1` is a legacy AS1 adapter.
  Full coverage needs a separate, scoped AS2 record (agent, joint action,
  states, agent-relative values; CQ12, 14, 15 and 17), or the record stays
  outside the source-validated subset.
* **Owner:** release coverage of 9.2 and 10.2 (physical-part and degree
  variants; arbitrary numbers of parts, negative branches, "only if").
* **Owner:** publication of the corrected candidate (not `8c092b6`), tags, and
  whether to send the upstream patch series and the letter. I have done none
  of these.
* **Owner:** the original sources other than the compendium and Atkinson 2007
  (the 1995, 1996 and 2006 books and the 2019 correspondence) remain
  unavailable; the holds that depend on them stand.

## Corrections (added 2026-11-21, evening)

Both follow-up reviews found errors in this response. Each was checked again
and is corrected here; the text above is left as it was written.

* **"Seven of the eight."** The table in section 1 lists nine claims and
  marks eight reproduced. The count was wrong.
* **The defeated default.** My probe stipulated nothing against the default:
  it supported a contrary by an argument, which tests a different rule. The
  first review's exact probe (a stipulated contrary) gives two arguments and
  `p`, `q`, `extra` assumed on `96daf6c`, and no argument with `p` out on
  `fccec4f`. It was a ninth real defect, fixed by the acceptance candidate.
* **"Safety valves, not working limits."** Too strong. The budgets of
  50,000,000 are finite limits; a larger finite input is still refused, with
  the documented outcome.
* **The timings.** The 500-, 800- and 1,200-fact graphs were not kept in the
  repository, so those figures cannot be checked from it. The first review's
  follow-up keeps its own fixtures and timings; use those.
* **"Held for scope, not errors."** For seven of the 25 records (entries 22.4,
  28, 30, 31.1, 45, 58, 59.1) the hold named a real defect: an exception or a
  conclusion that names too few participants, so that one instance can
  defeat or feed another. All seven are repaired in library 0.3.0.
* **"The books are unavailable."** Overbroad: relevant pages of Walton 2002
  and 2006 and an author-hosted 2003 paper were found by the first review.
  Other originals remain unavailable.
* **My key lookup** missed statements whose keys a program spelled
  differently from the engine (`accepts( a )`); files read by the command
  were not affected. Engine walton.5 fixes it.

See `docs/reviews/implementer-second-response.md`.
