# Approvals against the compendium (decision D20)

**Status, later the same day:** the owner supplied the originals; the three records still held below are closed, and AS1 and AS2 have records of their own (`final-closure.md`). *Scare Tactics* (2000) was compared in a second round; the name `exeption_to_rule` is documented by the README's Updates of 2019-10-08 (Walton's correspondence of 2019-09-19, private).

2026-11-21. On this date the owner decided D20: the compendium (Walton, Reed
and Macagno 2008) is the source of record, and an inherited record filed under
one of its entries is approved against that entry. The older work such a
record cites is recorded as not compared; comparing it is research for later,
not a condition of the stable release. The owner also decided that the
upstream patches and the memorandum to the authors are not to be sent (D22).

The reviewer is the implementer, who did not write the inherited executable records. (Strictly, the implementer had made three earlier changes to them in Phases 5.5–5.9: the assumption `credible(W)` of `expert_opinion`, the argument order of `in_domain` in `witness_testimony`, and the compendium citation added to each source line. All three were in the tree, `49ad425`, that the non-author compendium reviews compared: `first-half-adjudication.json` for entries 2, 3 and 26, and `2026-11-21-second-half-source-review.md` for the other eight. Those reviews, not this one, are the non-author check of the three changes; found at integration, `final-baseline-integration.md`.) The later independent review found that the same reviewer did author the `sunk_costs` reclassification from entry 25 to 24; that source-contract change now has fresh non-author acceptance in `final-source-independent-review.md`. The other eleven D20 source decisions retain their non-author evidence. Each was
compared with its printed entry, page by page, for the premises, the
conclusion and every printed critical question; each has a passing book
example. The approval states a reading, declared in the ledger, not an
equivalence with the print.

## The twelve approvals

| Record | Entry, page | Reading declared | Older work cited, not compared |
|---|---|---|---|
| `expert_opinion` | 2, p. 310 | the six questions as the book's own Carneades model sorts them (pp. 374–375) | Walton 1997, pp. 211–225 |
| `witness_testimony` | 3, p. 310 | position to know stated for a domain (shared with 1); truth-telling as belief with assertion; bias scoped to the account | Walton, Prakken and Reed 2003; Walton 2008 (witness passage) |
| `sunk_costs` | **24**, p. 326 (was 25) | the argument from waste with the sunk costs stated; CQ1 the feasibility default, CQ2 the exception that future losses outweigh the value | Walton 2002, p. 489 |
| `ignorance` | 26, p. 327 | CQ1 exception; CQ2 and CQ3 (burden, standard) answered by how the graph is built (D16) | Walton 1995, p. 150 |
| `correlation_to_cause` | 29, pp. 328–329 | CQ2 (more than coincidence) as the default that a theory explains the causation | Walton 1995, p. 142 |
| `positive_consequences` | 33.1, pp. 332–333 | premise split into consequence and value; CQ1 likelihood-adequacy default, CQ2 premise, CQ3 rebuttal | Walton 1995, pp. 155–156 |
| `negative_consequences` | 33.2, pp. 332–333 | as 33.1, with the opposite value and conclusion | Walton 1995, pp. 155–156 |
| `ethotic1` | 41, p. 336 | "accepted as more plausible" as a defeasible argument for the statement, weighed by the proof standard (CQ3) | Walton 1995, p. 152 |
| `ethotic2` | 41, p. 336 | lowers the plausibility of the assertion; does not assert its negation | Walton 1995, p. 152 |
| `slippery_slope_base_case` | 49, pp. 339–340 | a component: the act realizes the first event and has negative consequences; the whole pattern is `slippery_slope_compendium` | Walton, *Slippery Slope Arguments*, pp. 93, 95 |
| `slippery_slope_inductive_step` | 49, pp. 339–340 | a component: horrible costs carried back one link | the same |
| `established_rule` | 55.1, p. 343 | Carneades' rule-application form; CQ1 premise, CQ2 and CQ3 exceptions scoped to the rule and conclusion; the actor form is `established_rule_actor`; the inherited template of `applicable/1` says "Argument" (D5) | Walton 1995, p. 147 |

`sunk_costs` was filed under entry 25, whose printed form is the
precommitment of `sunk_costs_precommitment` and which prints no questions.
Its premises (costs sunk into the action, too high to waste) and its two
statements from upstream (feasibility, future losses) are entry 24's argument
from waste and its two questions, so it is filed there now; the ledger's
entry table and the record's source line say so.

With these, 140 of 143 records are approved, and every compendium entry the
library represents has an approved record.

## Still held

| Record | Why | What is needed |
|---|---|---|
| `practical_reasoning2` | no compendium entry; its source was never seen by any review | Walton 2015, below: the instrumental scheme and its questions |
| `exeption_to_rule` | no compendium entry. Its citation (1995, p. 148) is the reference printed for entry 57, argument from precedent (p. 344), but the record applies the ratio of a similar case, while 57 concludes from a counterexample that the rule needs an exception or change | Walton 1995, p. 148, and, if it exists in writing, the 2019 correspondence behind the name |
| `practical_reasoning1` | no compendium entry; its source, Atkinson and Bench-Capon 2007, is at hand. AS1 with 17 questions: CQ5–9 and 11 only through inherited, stronger exceptions; CQ12, 14, 15, 17 through none | the owner's choice: approve as a qualified AS1 record with these differences declared, or build an AS2 model |

## The sources and their identifiers

DOIs were found for these; the books without one are listed by ISBN.

| Work | Identifier | Needed for |
|---|---|---|
| Walton, *Goal-Based Reasoning for Argumentation*, Cambridge University Press, 2015 | doi:10.1017/CBO9781316340554 | `practical_reasoning2` (held) |
| Walton, *A Pragmatic Theory of Fallacy*, University of Alabama Press, 1995 | no DOI; ISBN 0-8173-0798-3 (e-book 0-585-20160-9) | `exeption_to_rule` (held), p. 148; and pp. 142, 147, 150, 152, 155–156 for the D20 records |
| Atkinson and Bench-Capon, "Practical reasoning as presumptive argumentation using action based alternating transition systems", *Artificial Intelligence* 171 (2007) 855–874 | doi:10.1016/j.artint.2007.04.009 | `practical_reasoning1` (at hand) |
| Walton, "The Sunk Costs Fallacy or Argument from Waste", *Argumentation* 16 (2002) 473–503 | doi:10.1023/A:1021108016075 | `sunk_costs` (approved under D20) |
| Walton, *Witness Testimony Evidence*, Cambridge University Press, 2008 | doi:10.1017/CBO9780511619533 | `witness_testimony` (approved under D20) |
| Prakken, Reed and Walton, "Argumentation Schemes and Generalisations in Reasoning about Evidence", *Proceedings of ICAIL 2003*, pp. 32–41 (ACM) | DOI not found; open copy at scholar.uwindsor.ca/crrarpub/9 | `witness_testimony` |
| Walton, *Appeal to Expert Opinion*, Penn State University Press, 1997 | no DOI found; ISBN 978-0-271-01694-8 (cloth), 978-0-271-01695-5 (paper) | `expert_opinion`, pp. 211–225 |
| Walton, *Slippery Slope Arguments*, Clarendon Press, 1992; reprinted Vale Press, 1999 | ISBN 978-0-19-823925-3; Oxford Academic online edition (DOI not confirmed) | the two slippery-slope components, pp. 93, 95 |
