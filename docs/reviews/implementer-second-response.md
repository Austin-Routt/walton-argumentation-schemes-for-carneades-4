# Implementer's second response: the two follow-up reviews

2026-11-21, evening. Branch `review/implementer-second-response`, based on the
first review's closure candidate `fcccbba` (library 0.3.0, engine
4.3.1+walton.5), which contains master `ea46ab5` (pull request #16).

Inputs: the second review's follow-up (sections 11 to 16 of its report, pull
request #16, merged, pull request #17, open, and its evidence archive); the
first review's closure package (report, candidate `fcccbba` as a Git bundle,
evidence); the compendium (Walton, Reed and Macagno 2008); Atkinson and
Bench-Capon (2007). Every finding below was rerun or read again here; the
engines were built from `ea46ab5` (walton.3), #17's head `e48da87` (walton.4)
and `fcccbba` (walton.5) with Go 1.24.7, a third toolchain besides the two the
gate pins.

## 1. Are the follow-ups accurate?

Both are accurate in substance. Their corrections of my first response are
right and are recorded in `implementer-response.md` ("Corrections").

**The second review's findings**, rerun:

| Finding | On master (walton.3) | #17 (walton.4) | `fcccbba` (walton.5) |
|---|---|---|---|
| F1: a scheme concluding ¬S on a negative S | `¬¬q(a)` in, `q(a)` out | `q(a)` in | `q(a)` in |
| F1, inside a term: `q(¬X)` with `X = ¬a`, and `q(a) ⇒ t` | `q(¬¬a)` in; `q(a)`, `t` out, though the argument for `t` was built | the same as master | `q(a)`, `t` in |
| F2: four statements `q(a)` … `¬¬¬q(a)`, 40 runs | refused 16 times, one issue 24 times | refused 40 times | refused 40 times |
| F2: issue scheme `[leads(X), trails(Y)]`, anchor `leads(x)`, four `trails`, 40 API runs | 9 different issues; 27 change on a second inference | `leads(x) trails(a)` every time | all five positions every time |
| F2 with `leads(x)` and `trails(c)` both stipulated | sometimes refused, sometimes accepted | accepted, issue `{leads(x), trails(a)}`: the contradiction is missed | refused, inconsistent context |
| F3: API key `accepts( a )` | no issue | issue, keys kept | issue, keys normalized (contract 4) |
| F4: the interrupt check | fixed on master by #16 | — | kept |

So F1 to F4 are real, and #17 fixes F1 and F3 and makes F2 repeatable. But
#17's F2 fix keeps the first of the matching statements in sorted order, so a
stipulated alternative outside that choice escapes the consistency check, as
the first review said; and it folds only the outer double negation, while the
rule store reads `q(¬¬a)` as `q(a)` (the argument for `t` was built on every
engine). Engine walton.5 is right on both. #17 is superseded; its tests are
kept (section 3).

The second review's D15 to D23 are discussed in section 2 and in the decision
register of `PLAN.md`.

**The first review's closure**, checked:

* Its engine findings reproduce (the table above, the default-seed probe and
  the noncanonical-key probe on all three engines). Its candidate passes the
  full local gate here: 9,979 of 9,979 checks, the engines' own tests,
  validator 0 errors and 11 warnings, source-approval integrity for 143
  records, and the 13 dialogue-snapshot tests.
* I read the engine diff (`caes.go`, `normalize.go`, `negation.go`,
  `validation.go`, `yaml.go`, `transaction.go`). The issue lookup keeps an
  index of parsed keys and confirms each candidate by matching, so typed
  equality is kept; later patterns no longer carry one candidate's bindings
  to the next. I found no defect. One behaviour was undocumented: after
  inference, a program's own key `accepts( a )` becomes `accepts(a)` (same
  statement object). Contract 4 now says so.
* Its library work (23 new records, 24 changed) was written and approved
  inside one review system, by separate agents. I checked all of it against
  the printed pages (section 4).
* Two defects of the candidate's documents, fixed here: `PLAN.md` had lost
  the decision register D1 to D23, which the contracts, the ledger and the
  reviews cite as `PLAN.md Dn`; and README, plan, release notes, change log
  and user guide had words fused to numbers ("provides24", "records:119",
  "gateL") in over a hundred places.

## 2. Where they agree and differ

They agree on: withdrawing `8c092b6`; every engine defect above; the 19
record changes before #16 and the seven records whose holds named real
defects (22.4, 28, 30, 31.1, 45, 58, 59.1, now repaired in 0.3.0); D16's
static evaluation; D22; D23.

They differ on four decisions. The repository now follows the closure on all
four; two are the owner's to settle.

| Decision | Second review (#16) | First review (closure) | Mine |
|---|---|---|---|
| D17, variants | one variant per record, declared | implement the omitted branches (9.2, 10.2 and 13 printed alternatives) | closure; more coverage, each branch declared |
| D19, polarity | "less plausible" is an argument for the negation; "should not" is ¬should_do | no general equivalence | closure: bad character lowers the weight of what a person says; an argument for ¬S would accept the contrary whenever nothing else is argued, the fallacy itself |
| D20, source of record | the compendium; older originals not a release condition | originals remain obligations | **owner**. Both are honest if the ledger says what was compared; I would take D20 with a recorded "original not yet compared" for the 13 rows |
| D21, legacy records | three legacy records outside the stable count | no retrospective narrowing | **owner**, with D20; the checker has no legacy class yet |

Approvals: 88 of 120 at `ea46ab5`, 128 of 143 at `fcccbba`. Of the 15 held
now, 13 wait for original sources (D20) and two (`practical_reasoning1`,
`sunk_costs`) for model work. Under D20 and D21 most of the 13 could be
approved against their entries; that is the one choice left that changes the
count.

## 3. What this branch adds to `fcccbba`

* From #17, kept with walton.5: the random comparison of key lookup and scan
  on 400 graphs (`issue_lookup_random_test.go`), the folded-instance
  validation test, and the checks `derived-double-negation-is-the-statement`
  and `issue-construction-is-repeatable` (the latter rewritten for walton.5:
  the overlapping graph is refused, every time).
* New: `derived-nested-double-negation-is-the-statement`,
  `authored-double-negation-collision-is-refused` (a lone `¬¬q(a)` reads as
  `q(a)`; keys `q(a)` and `¬¬q(a)` together are refused), and three library
  cases in `tools/vv/cases/negation.yml` (26 and 13.2 in a chain, 32.2, and a
  positive control). 9,989 checks.
* The decision register restored in `PLAN.md`, with each of D15 to D23's
  status here; the key normalization stated in contract 4; the fused numbers
  separated; corrections to my first response.
* The ledger rows of the 46 records in section 4 record my check.

## 4. The library changes against the print

I compared each new or changed record with its printed entry: pp. 312
(4.5–4.7), 314 (5), 315 (6.4, 7), 316–317 (9.2, 10.2), 318 (13.2), 321
(19), 322 (20), 323–325 (22.1, 22.4, 22.6, 23), 327 (26.2), 328 (28), 329
(30, 31.1), 332 (33.3, 33.4), 336–337 (41, 45), 339 (49), 342–343 (54.2,
54.3, 55.1), 345 (58, 59.1). I concur with all 46. The variants follow the
print's parentheses (4.5 "is (or should be)", 4.6 "x (or accepting A)", 4.7
"accept A (acquire P or buy O)", 20 "(small)", 22.6 "(bad)", 26.2 "(true)");
the identity changes keep one instance from defeating or feeding another, and
follow the printed questions (28 CQ3 "in the given case", 30 CQ2 relative to
the event, 22.1 and 22.4 "my", "I"); 7's CQ3 no longer asks for a case "more
on point", which the print does not say; 41's qualitative forms keep "more
plausible" and "less plausible" as weights. Three readings rather than
transcriptions, in five records, each now noted in the ledger:

* 9.2, the three degree variants conclude that y is A and that it is less A
  than X; the print offers the lesser degree in place of the plain
  conclusion.
* 10.2, the positive branch requires a sufficiency warrant. The print affirms
  the whole from one part under an only-if premise, a presumptive form; the
  added premise is stricter and declared.
* 26.2, the true branch's template says "I should act", the print "I act".

I did not check entry 3 (witness testimony), which stays held.

## 5. Left open

* **Owner:** D20 and D21 (section 2); tagging and publishing the preview at
  the commit that merges this branch; sending the upstream patches and the
  source memorandum, prepared by the first review and not sent.
* **Reviewer other than me:** the checks and tests this branch adds, and the
  edits to `PLAN.md` and contract 4.
* **Stable release:** the 15 holds, each with one action in
  `docs/reviews/closure-current-register.md`; a full AS2 record if
  `practical_reasoning1` is to be source-approved.
