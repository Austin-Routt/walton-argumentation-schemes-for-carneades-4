> **Independent implementation acceptance, 2026-11-21:** see `reviews/record-dispositions.md`, `reviews/atkinson-2007.md`, `reviews/independent-engine-cross-review.md`, `reviews/pr13-reconciliation.md` and the revised `PLAN.md`. The historical convergence below established a development plan; later counterexamples required additional repairs. PR #13 arrived during this audit and is preserved below and under `reviews/peer-13/`. Its recommendation to publish the old R1 pin and its 90-row approval population are superseded by the reconciled evidence. The corrected preview has 70 approved records and 50 explicit content holds.

# The three reviews, 2026-11-20 and 21

Three independent reviews examined this repository, the two engine archives and the source text, Walton, Reed and Macagno, *Argumentation Schemes* (2008): "Fable" (first review), Claude (second) and ChatGPT (third). Each round, every reviewer checked the others' work against the engine and the book. Their full reports, evidence and patches are in their own archives, which are not part of this repository.

| Round | What each review delivered |
|---|---|
| 1, 2026-11-20 | Independent audits of the engine and the library, each with its own harness (`tools/vv`, `tools/go/schemecheck`, `tools/enginecheck`) |
| 2, 2026-11-21 | Each checked the other two and the book; the plans turned to "engine first, change a scheme only where it is wrong" |
| 3, 2026-11-21 | Each made the decisions the earlier drafts had left to the owner. The first review then merged the three decision sets into `PLAN.final.md`, which is `PLAN.md` here with the owner's instructions applied |

## The verdict

The reviews have converged and no further broad review is needed; all three say so. They agree on the diagnosis, on the order of work, on the source text as the reference, and on 10 of 14 decisions as written; the other 4 were settled by a measurement, the book or the README (`PLAN.md`, "Where the three decision sets differed"). What replaces further rounds is narrower: every change is checked by `tools/vv` in CI, and at gate S and before release R1 someone who did not write the changes reruns the strict suite from a clean checkout and checks the reference cases against the book by hand.

## Third-round documents, as checked by the second review

The second review reran or read what each claim rests on. Measurements were repeated with the archived engines and Go 1.24.7.

| Claim | Verdict |
|---|---|
| First review: 5,947 checks, 5,557 pass, 390 known failures under 24 issues | Reproduced exactly |
| First review: the nine scheme repairs give 122 failing checks and 108 of 114 schemes by command on the engine as shipped; 34 failing and 114 of 114 with the prototype engine and the `questioned:` list | Reproduced exactly; the one newly failing check is the user predicate `neg`, as stated |
| First review: the suite on the engine sources in `third_party/` | All 5,947 verdicts identical to the run on the unpacked archives |
| First review: of 28 claims of the other two it checked, 21 confirmed, 3 qualified | Its three qualifications of the second review hold: 106 and 109 are two ways of counting; `in`, `or`, `div`, `mod` work as predicates and only `is/2` misbehaves; the compendium cites two page numbers at second hand |
| Third review: the second review's sufficient-condition scheme still says "necessary" (22.3, p. 323), so its round-2 statement that the library already uses the correct text was wrong | Correct: `alternative_action/2` says "necessary" and the sufficient scheme uses it |
| Third review: the second review's six-edit trial leaves roles, polarity and the sorites sequence open, and its `pop_scheme` emits both readings | Correct; the plan's nine repairs and the R2 split address them |
| Third review: the second review's acceptance cases discard the error `Infer` returns | Correct for the round-2 package it saw; fixed in the second review's round 3 |
| Third review: an opposition issue alone does not make contradictory assumptions consistent | Correct (both reviews ran the fixture): both positions are labelled in. Now contract 4 |
| Third review: one coordinated identifier migration, with previews before it | Adopted: R1 changes only the nine schemes that do not run today; every rename goes to R2 |
| README, 2019-10-08: "Argument From Precedent" was renamed "Argument From Exception To A Rule" after correspondence with Douglas Walton | Confirmed in the README; it settles the id of the inherited record (`exeption_to_rule` stays) |

## The owner's instructions of 2026-11-21

No engine forks and no owner approvals. The engine sources are kept in `third_party/` with one commit per fix (D1), and names follow the rule of D6. The second review implements the plan in this repository; CI runs the gate on every push.

## The independent review of 2026-11-21 (D14)

The first reviewer wrote the adopted plan and the first version of `tools/vv`, and none of the changes of pull requests #2 to #12. At the owner's request it acted as the reviewer D14 asks for: it reran the gate from clean clones, compared the cases and the records with the book by hand, and decided each row of `tools/vv/source_ledger.yml` (`review:`).

### What was rerun

| commit | engines' unit tests | strict run of `tools/vv` | validator |
|---|---|---|---|
| 8c092b6 (R1) | pass, also under the race detector | 7,164 of 7,164 checks pass, ledger empty | 0 errors, 25 warnings |
| 49ad425 (`master`) | pass | 7,994 of 7,994 checks pass, ledger empty | 0 errors, 10 warnings |

Both ran in new clones from GitHub with Go 1.24.7, and both clones were unchanged afterwards. The CI logs of 8c092b6 show the same 7,164 checks passing on Go 1.24.13 and on Go 1.27.4. Every pull request from #2 to #12 had a passing run on its head commit before it was merged.

Four deliberate defects were each caught: the refusal of an inconsistent context switched off in the engine (5 checks fail), a premise taken out of `ad_hominem` (172 fail), a reference case renamed (the list of expected checks no longer matches), and the arguments of one conclusion swapped on `master` (the record's instance fails).

Of the 5,947 checks of the first review's suite, 5,488 run under the same names. The 459 others are the per-scheme check for variable capture (228; now two engine checks and the range check at load), the per-scheme check against negated premises (228; plan 2.4 allows them), and three engine checks restated as `tools/vv/README.md` says.

### What was compared with the book

For R1, the 66 cases of `tools/vv/cases/` and the seventeen records they exercise. For `master`, all 120 records with their templates, sorted questions and instances (`cases/book.yml`), entry by entry against chapter 9, with the page image where the text layer was in doubt (pp. 96, 325, 340, 341). The 24 records from Carneades were compared field by field with upstream's own file: they differ from it only where the changelog says (two exceptions in `appearance`, one assumption in `expert_opinion`, one exception in `verbal_classification`, the order of `in_domain`'s arguments in `witness_testimony`) and in the template of `biased/1`. Wyner (2016), Walton (2006) and Walton (2002) were read in open copies and bear out the three comparisons of 5.5. `practical_reasoning1` was compared with Atkinson and Bench-Capon (2007), which the owner supplied.

### What was found

Nothing that contradicts what release R1 claims. In the library on `master`:

1. `two_person_practical` (23): the template of premise 2 says "sets the situation"; the book has "sees" (pp. 97, 325). The wording is the original file's.
2. `practical_sufficient_condition` (22.3): the alternatives premise was corrected to "at least one of ... is sufficient"; p. 96, which the correction cites, prints "each one of ... is sufficient".
3. `negative_practical` (26.2): the conclusion's template reads "not true or true" and so concludes nothing. Premise 5 as the record states it goes with "not true". The instance in `cases/book.yml` has its example's consequences the wrong way round.
4. Wording in three more templates: the exception of `popular_opinion` (4), the major premise of `rhetorical_oppositions_descriptive` (12.2), the exception of `verbal_classification_vagueness` (16).
5. `practical_necessary_condition`, `practical_sufficient_condition` (22.2, 22.3): four premises in the first person carry no agent although the record names one. With two agents in one graph, one selection gives both an argument.
6. `abductive_action_to_character`, `abductive_character_to_action` (31.3, 31.4) have no variable for the character quality, and `retroductive` (31.5) none for the event. Statements about two qualities or two events are the same statement.
7. `from_example` (6.1): the conclusion carries the example, so two examples do not support one generalization. 6.2 was corrected for the same thing.
8. `part_in_whole` (9.2): the exception is the parenthesis of the printed conclusion, a variant of the conclusion and not a defeater.
9. `whole_into_parts` (10.2): only the part listed first can carry the property, and the negative reading is not stated.
10. Statements the shared wording of 5.6 left as two predicates without saying so: "person P is V" (18.1, 18.2), "a is a (morally) bad person" (47, 42), "a has property F and property G" (6.1, 6.2), "it is reasonable to believe φ" (59.1, 60).
11. The template of `biased/1` is upstream's predicate with the book's wording ("Arguer %s is biased."); D5 keeps upstream's ("%s is biased."), which also serves entry 46.
12. Records in the first or second person ("I", "you") name no participant. That is a workable convention, one proponent and one respondent in a graph, but `docs/using.md` does not state it.
13. With the whole library loaded, any `has_property` fact gives an accepted "all the parts of X have the property": the heredity question of 10.1 is an assumption, which holds until it is questioned. The validator's cycle warning names the five records involved.
14. Of the 28 commits that change the engine up to R1, two change both engines (5bb648a, 93ca9d0) and two change engine and library together (e65e715; 770a383, the version numbers). The first three need splitting before they are offered upstream (D1). There is no engine change after R1.
15. "90 of 119 records share a statement with another instead of 73" (changelog, 5.6): a recount gives 89 of 119, 72 before, and 90 of 120 on `master`. No tool in the repository computes the figure.

### Decisions on the rows

90 rows are approved and at "example passed": 83 records of the book and 7 from Carneades. 13 are held: findings 5 to 9 (eight records) need a change to what a graph has to state, and findings 1 to 4 (five records) a wording fix, proposed in a pull request of its own for the implementer to approve. 16 rows of records from Carneades stay open because the source they cite is not available. `practical_reasoning1` has its row written by the reviewer and waits for a second reviewer.

### Left to the owner

* Publish R1. GitHub refused the reviewer's push of the tag as it refused the implementer's (HTTP 403), so the three commands of `PLAN.md` are the owner's to run.
* Decide whether a record from Carneades that is filed under an entry may be closed against that entry when the book it cites cannot be had. The reviewer recommends it, with the row saying so; the five questions of 5.2 can then be modelled from the compendium.
* Offer the engine fixes upstream and write to the book's authors. The reviewer prepared the patch series and a draft letter and sent neither.
