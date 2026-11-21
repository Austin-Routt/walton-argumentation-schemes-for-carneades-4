# Follow-up of the reviewer of 2026-11-21 to the implementer's response

Reviewer: "Fable", who wrote the plan before the independent reviews and the review merged as pull request #13 (`docs/reviews.md`, "The independent review of 2026-11-21 (D14)"). Reviewed: master `e6d2d6f` (pull request #15 merged), `docs/reviews/implementer-response.md`, and every change to the engine, the library, the ledger and the suite since `96daf6c`. "The third review" below is the ChatGPT team that wrote the acceptance candidate; "the implementer" wrote the response and engine 4.3.1+walton.3.

The owner asked three things: is the response accurate, do I agree with the changes, and can the remaining gaps and decisions be closed. The answers are in sections 1 to 3, 4 and 5 to 7.

## Summary

* **My verdict of this morning is withdrawn.** I wrote that Release 1 at `8c092b6` passed and could be published. It should not be. The third review's counterexamples are real defects of that engine; I reproduced each of them (section 1). My probes had covered two derived positions and two stipulated ones, not a stipulated position against a derived one, and I had accepted the check on authored arguments as described instead of attacking it.
* **The response is accurate with three exceptions**: two in its table of engine claims and one about why rows are held (section 2).
* **I agree with the changes** to the engine and to the 19 records, each checked against the printed page (sections 3 and 4).
* **Reviewing the implementer's engine changes turned up three more engine defects**, two of them older than the changes, **and one check that fails at random** (section 3): the argument graph does not read `¬¬S` as `S` although the rule store does; issue construction depends on run order; the new key lookup misses keys that are not in normal form. They are fixed in the branch `engine/negation-identity` (engine 4.3.1+walton.4), which someone else has to check because I wrote it.
* **Decisions D15 to D23 close what the response left to the owner or to a reviewer** (section 5). Under them 18 held rows are approved here (88 of 120) and each of the other 32 has one decided action (section 6).

## 1. The old Release 1 engine against the candidate

Clean clones, Go 1.24.7, the commands a user runs. Probe files are in the review archive.

| Probe | Engine of `8c092b6` | Candidate (`e6d2d6f`) |
|---|---|---|
| One fact, premises `person(X)`, `person(Y)` | no argument | `pair(alex,alex)` in |
| Issue scheme `[accepts(X), rejects(X)]`, both stipulated | exit 0, both in, no issue | exit 3, inconsistent context |
| Assumption `q(f(g(a))` with a bracket missing | parse error printed, exit 0, a conclusion derived from it | exit 1 |
| Authored argument naming a scheme, with an undercutter that nothing can reach | `check` exit 0, conclusion in although the exception is stipulated | exit 1 |
| Three-way join over 160 facts | 6.4 s and 790 MB before "limit reached" | 0.9 s and 139 MB |
| Library: `is_required(b)` stipulated, `¬is_required(b)` derived | both in, no issue | one issue, the stipulation in, the derived position out |

Eight of the third review's nine regression tests for the argument engine compile against the engine of `8c092b6` (the ninth uses an interface that engine does not have). Seven of the eight fail (cancellation before an empty theory, a generated issue accepting two positions, an invalid issue pattern accepted, default priority, a malformed assumption in an empty theory, issue closure not atomic, a defeated default seeding another), and the control passes.

## 2. The response, claim by claim

| Claim | Finding |
|---|---|
| Master is `e6d2d6f`, automated tests green on both Go versions | Confirmed: the three jobs of the run for `e6d2d6f` succeeded (Go 1.24.13, Go 1.27.4, the comparison of the two). |
| All 8,230 checks pass | Confirmed from a clean clone with Go 1.24.7: engine tests pass, the strict run exits 0 with 8,230 of 8,230, the validator reports 0 errors and 10 warnings, the clone is unchanged. |
| "Seven of the eight" engine claims are real defects | **Inaccurate in its count.** The table above that sentence has nine rows, and eight of them say "reproduced". |
| "A defeated default seeds further defaults: same behaviour on both" | **Not what I measure.** The third review's test for it fails on the engine of `8c092b6` and passes on the candidate, and contract 5 now describes the fix. By my run it is a ninth real defect, fixed. |
| `8c092b6` should not be published | Agreed (section 1). |
| The candidate's work budgets refused ordinary inputs; walton.3 fixes that | Confirmed. The two checks pass, and I measured the lookup against the scan (section 3). |
| 70 of 120 approved; the 25 rows the third review holds and I approved are held "for questions of scope, not errors" | The count is confirmed. The description holds for 18 of the 25. **For seven the hold names a defect I did not see**: an exception or a link identified by too few of its participants, so that one instance can defeat another. The exception of `sign` (30) is keyed by the sign alone, so one competing explanation undercuts every candidate event; 28, 59.1, 45, 58 and 22.4 have the same shape, and 31.1 does not link the chosen account to the set it is chosen from. I withdraw my approval of those seven. |
| The eight changes the third review made and approved itself agree with the print | Confirmed against pp. 315, 329 to 331, 342, 344 and 346 (section 4). They now have two checks by people who did not write them. |
| The Atkinson and Bench-Capon paper: same 17 questions, same conclusion as my comparison | Confirmed; the reference case now uses the paper's own example (p. 866). |
| Pull request #14: four of its six wordings were already in the candidate, entry 16 and the "I and you" note were taken over, #14 closed | Confirmed in `walton.yml`, `docs/using.md` and finding 14 of `docs/reviews.md`. |

## 3. The engine changes

**The third review's changes (4.3.1+walton.2)**: streaming match with quotas, one fact for several premises, strict parser, issue closure before inference, staged and atomic inference, the check on authored exceptions. I read them and ran their tests against both engines. I agree with all of them.

**The implementer's changes (4.3.1+walton.3)**, which it asked someone else to review:

| Change | How I checked it | Result |
|---|---|---|
| `terms.Substitute` marks every variable of a binding chain | 20,000 random terms and bindings against a copy of the old function | Identical wherever the old one terminates. Agreed. |
| Positions looked up by key when the remaining patterns are ground | 4,000 random graphs with five issue schemes, lookup against scan | Identical on all 3,989 graphs for which the scan itself is stable; the other 11 differ between two runs of the scan (F2). Agreed for keys in normal form; see F3. |
| `MAXISSUESTEPS` and `MAXMATCHSTEPS` at 50,000,000 | The four-premise join over 120 facts needs about 1.7 million visits | Reasonable as safety valves. Agreed. |
| The exception check normalizes both keys | Read | Agreed. |

**Four findings.** F1 to F3 are fixed in `engine/negation-identity`; F4 is a defect of one check and is fixed here:

* **F1. The graph does not read `¬¬S` as `S`.** Contract 4 says it does, and the rule store does. A scheme that concludes `¬S`, applied to a negative statement, concludes a statement `¬¬S` of its own. The store reads it as `S` and builds the next argument from `S`; the graph labels `¬¬S` in and leaves `S`, and what that argument concludes, out. With the library and nothing written by hand: `asserts(bob,¬is_required(al))`, `bad_moral_character(bob)` and `can_be_required(al,cy)` give `¬¬is_required(al)` in, `is_required(al)` out, and the `alternative_normative` argument built on it out. Written by hand, `¬¬q(a)` behaves as a rival of `q(a)`: assuming both is reported as an inconsistent context. Both engines do this; the suite had no check for it.
* **F2. Issue construction depends on run order.** It ranges over a Go map. A graph with `q(a)`, `¬q(a)`, `¬¬q(a)` and `¬¬¬q(a)` under the library's negation scheme ends as invalid input ("connects distinct existing issues") in 13 runs of 40 on the candidate and as one issue in the other 27; the engine of `8c092b6` always exits 0 but builds two issues in 8 runs of 40. An issue scheme whose later pattern has a variable of its own gives different positions from run to run.
* **F3. The key lookup misses keys that are not in normal form.** For a graph built through the Go API with keys such as `q( c )` and `¬q( c )` the scan found the issue and the lookup finds none. The command and the YAML reader write keys in normal form, so they are not affected.

* **F4. The interrupt check races a limit.** `cli-interrupt-is-cancelled` interrupts a run 0.3 s after it starts and expects "cancelled". Its workload, 300 facts under a scheme over pairs, was seconds of work on the old engine. Since the limits of engine 4.3.1+walton.2 it ends as "limit reached" at 100,000 inserted constraints after 0.76 s on my machine and sooner on the test runners. Of the two runs of the first commit of this pull request, one failed that check on both Go versions (exit status 4 where 5 was expected) and the other passed; the candidate has the same race on master. This is a defect of the check, not of the engine. It is fixed in this pull request, in `tools/vv/internal/vvlib/cli.go`: the workload is now a join of three of the 300 facts with a fourth fact that does not exist, 27 million candidate visits under every limit, 7 s on my machine. What the check asserts is unchanged.

A smaller observation, not changed: the YAML the command writes lists statements in a different order from run to run; the content is the same.

## 4. The library changes

Nineteen records changed their statements since `96daf6c` and seven templates changed their wording. I compared each with the printed entry. I agree with all of them.

| Records | Change | Print |
|---|---|---|
| 6.1 `from_example` | concludes the generalization `usually_properties(F,G)` | p. 314; my finding 7 |
| 6.2, 18.2 | premises shared with 6.1 and 18.1 by meaning | pp. 315, 321 |
| 6.3 `from_model` | premise 3 is the conditional about the model's act | p. 315 |
| 9.2 `part_in_whole` | the comparison in the parenthesis is no longer an exception | p. 316; my finding 8 |
| 22.2, 22.3 | the agent in every first-person premise; "each one is sufficient" | pp. 95, 96, 323; my findings 5 and 2 |
| 26.2 `negative_practical` | concludes the presumption that the statement is not true | p. 327; my finding 3 |
| 31.2 | the strongest diagram is linked to its conclusion | pp. 329, 330 |
| 31.3, 31.4, 31.5 | the action, the quality and the event are variables | pp. 330, 331; my finding 6 |
| 54.1 `rule_physical_world_1` | premise `referentially_coextensive(W,D)` added | p. 342. The print has one premise. The added one states the antecedent of the printed warrant and keeps `is_a(W,D)`, a statement other schemes produce for plain class membership, from yielding "means the same". I agree with it as a declared interpretation. |
| 59.1, 60 | conclusion shared; the two undercutters of memory belong to the person who recalls | pp. 345, 346 |
| 15, 55.1, 33.1, 33.2 (inherited) | printed questions added as exceptions or as an assumption | pp. 319, 343, 332, 333 |
| 56 `exceptional` | one backing statement for the three conditional questions | p. 344 |

**Entry 16.** The implementer replaced the third review's wording of `certainty_no_need/2` by the wording of my pull request #14 and asked for someone else to review it. It is applied exactly as I proposed it, and it follows the printed third question, which asks why the vagueness is a problem for the dialogue. I wrote it, so this is not an independent check of the wording; the row stays on hold until someone who wrote neither confirms it (section 6).

**Entries 33.3 and 33.4.** The response notes that 33.3 assumes the outcome is "strongly likely" where 33.1 and 33.2 assume a likelihood "adequate for the strength of this consequence argument". The printed question (p. 332) asks how strong the likelihood is and sets no threshold. 33.4 has the same wording, and its own second premise allows "is likely to occur", so "strongly likely" asks for more than the premise states. Decision D23 below.

## 5. Decisions

The response lists what it leaves to a reviewer or to the owner; the owner asked me to decide. D1 to D14 are in `PLAN.md`; these continue the series. Each can be overruled by the owner.

| | Decision | Reason |
|---|---|---|
| D15 | **Release.** `8c092b6` is withdrawn. The preview release `r1` is the corrected candidate with engine 4.3.1+walton.4. It is tagged at the master commit that contains that engine, after someone who did not write the fix has checked it and both Go jobs are green. | Section 1; F1 is reachable with the library's own schemes, and the release notes say no known failure is waived. |
| D16 | **Questions and burden.** The library implements the book's formal account of critical questions (pp. 33, 373 to 375): each question is sorted once, as a premise, an assumption, an exception or a rebuttal, and evaluation is static. It does not implement the dialogue of p. 96, in which asking any of the five questions of practical reasoning suspends the argument until it is answered, and does not claim to. A question about the burden or the standard of proof (26, CQ2 and CQ3) is answered by how the application builds the graph. | The book's own criterion (p. 33) is whether the questioner has to bring something: "what other goals, alternatives, consequences should be considered" asks for them to be named, "what grounds are there that it is possible" does not. Carneades' own practical-reasoning records sort the questions the same way. The suite already tests both behaviours for every record. |
| D17 | **Variants.** A record states one variant of its entry; the ledger names the variant and what is left out, and that is a complete record (4.6, 20, 22.6.1 to 22.6.3, and 26.2, whose print gives the opposite ranking and conclusion in parentheses). 9.2 is the species reading with the plain conclusion; the reading of a physical part has the form of 10.1, and the comparative conclusion is left out. 10.2 takes a part premise before the stable release, so that any part of a whole with any number of parts carries the property; the negative branch is left out. | The print packs alternatives into parentheses; one record cannot state both branches without changing polarity silently. |
| D18 | **Participants.** A statement carries a participant when the print names it in that statement, and an exception carries enough of them that one instance cannot undercut another (30, 31.1, 45, 58, 59.1, 22.4, 23). Schemes in the first and second person keep the convention of one proponent and one respondent. Where the print's conclusion names no agent, institution or language (19.1, 19.2, 54.2, 54.3), the record follows the print and the application qualifies its terms. Inherited records keep the arguments Carneades gave them and declare the limit. | The third review's two-instance test is right for exceptions, where a collision defeats an unrelated argument; for conclusions the print governs. |
| D19 | **Polarity.** The engine weighs positions and has no degrees. A negated recommendation reads "should not"; "more plausible" is an argument for a statement and "rejected as less plausible" an argument for its negation. Stated once in the contracts. | It is how the engine argues against a statement; 6.4 and 41 print exactly these words. |
| D20 | **Source of record.** The compendium is the source of record. An inherited record filed under an entry is complete when it has been compared with that entry, every printed question has a disposition, the differences are declared and its examples pass. Whether the older work it cites agrees is recorded when that work is at hand and is not a condition of the stable release. | The library's claim for such a record is "Carneades' record, corresponding to entry N as follows". That claim can be checked in full today. The citation is Carneades' own and is relayed as such. |
| D21 | **Legacy records.** `practical_reasoning1`, `practical_reasoning2` and `exeption_to_rule` have no entry. They stay as Carneades wrote them, are marked legacy in the ledger and the documentation, and are not counted as source-checked. No record for the fuller model of Atkinson and Bench-Capon is added. The stable gate is 117 records at `example-passed` and three legacy records identical to upstream. They can be promoted when their sources are at hand. | The fuller model needs agents, joint actions, states and value functions; a scheme record over ground facts cannot state it, and upstream ships the same abstraction. The comparison with the paper is recorded in full. |
| D22 | **Outside contacts.** The patches for the two upstream projects are regenerated from the released engine, and my letter to the authors is merged with the third review's memorandum. Both are sent only when the owner says so, and neither holds a release. | Unchanged from D1 and from this morning. |
| D23 | **Entries 33.3 and 33.4** take the reading 33.1 and 33.2 have: the likelihood is adequate for the argument, with no fixed threshold. | Section 4. Proposed by the implementer for 33.3; 33.4 follows from the same print. |

## 6. The 50 held rows

**Approved here, 18** (88 of 120). The rows and `docs/reviews/independent-approvals.json` carry the reason for each. `docs/reviews/record-dispositions.json` is the third review's record and is left as it wrote it; the ledger and the registry are the current state.

* Under D20, the five inherited rows held only for the older work they cite: `position_to_know` (1), `expert_opinion` (2), `definition_to_verbal_classification` (15), `positive_consequences` (33.1), `negative_consequences` (33.2).
* Under D16, D17 and D19, thirteen rows whose condition was a decision and not a change: 4.6, 6.4, 9.2, 13.2, 17, 20, 22.1, 22.2, 22.3, 22.6.1, 22.6.2, 22.6.3, 26.2. Nine of them I approved this morning and they have not changed; the other four (9.2, 22.2, 22.3, 26.2) were repaired as my findings 8, 5, 2 and 3 asked, which I checked against the print.

**Still held, 32**, each with one action:

| Row | Action |
|---|---|
| 4.5 `moral_justification_ad_populum` | State the one-respondent convention and the good-group branch in the source contract, with an example from the text; then approve. |
| 16 `verbal_classification_vagueness` | Someone who wrote neither wording confirms `certainty_no_need/2`; then approve under D16. |
| 19.1, 19.2 `positive_value`, `negative_value` | State that the goal term names the agent's goal (D18) and add a case with two agents; then approve. |
| 26.1 `negative_normal_expectations` | State one expectation context per graph and add a case with two; then approve under D16. |
| 54.2, 54.3 | State one institution, or one language, per graph, or qualified terms (D18), and add a case with two; then approve. |
| 4.7 `snob_appeal_ad_populum` | Replace the reference case that accepts an object by one that accepts a proposition. |
| 5 `popular_practice` | Qualify the practice by its context; add a case with two contexts. |
| 6.1 `from_example` | Add a case in which one counterexample does not defeat "usually". |
| 10.2 `whole_into_parts` | Part premise (D17). |
| 22.4, 23 | Carry the agent in the exceptions about alternatives and values (D18). |
| 30 `sign` | Key the competing explanation by the sign and the event (D18). |
| 31.1 `abductive_backward` | Link the chosen account to the set of accounts of the data. |
| 45 `circumstantial_ad_hominem` | Key the character and credibility statements by person and argument. |
| 58 `plea_for_excuse` | Key the excuse by the case. |
| 59.1 `perception` | Key the undercutter by the observer, as 60 now is. |
| 3 `witness_testimony`, 28 `cause_to_effect` (inherited) | Keep upstream's arguments; declare one matter per witness and one application per cause in a graph (D18, D20); then approve. |
| 7 `analogy` (inherited) | Add the entry's third question, another similar case with the opposite result, beside upstream's stronger `more_on_point`. |
| 25 `sunk_costs`, 55.1 `established_rule`, 49 (two rows) (inherited) | Declare the difference in form from the entry (D20); then approve. For 49 the two rows are the parts of upstream's decomposition; the single argument of the entry is not provided. |
| 26 `ignorance`, 29 `correlation_to_cause` (inherited) | D16 for the two questions on burden and standard; for 29 add a case that tells correlation alone from correlation with an explanation; then approve. |
| 41 `ethotic1`, `ethotic2` (inherited) | Declare D19 in the rows; then approve. |
| `practical_reasoning1`, `practical_reasoning2`, `exeption_to_rule` | Legacy (D21). |

## 7. What is open after this

* The engine fix (F1 to F3) and the change for D23 are mine, so the implementer or the third review checks them before they are merged. The new workload of the interrupt check (F4) is mine too and is merged with this review, because without it the gate fails at random; it should be read by someone else at the same time.
* I cannot create the tag: this session's connection refuses tag pushes. The commands for the owner are in the release notes once the engine fix is on master.
* The 32 rows above are the work for the stable release. None of them needs another round of general review.
