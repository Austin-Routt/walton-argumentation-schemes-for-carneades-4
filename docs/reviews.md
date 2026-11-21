> **Independent implementation acceptance, 2026-11-21:** see `reviews/record-dispositions.md`, `reviews/atkinson-2007.md`, `reviews/independent-engine-cross-review.md` and the revised `PLAN.md`. The historical convergence below established a development plan; later counterexamples required additional repairs. The current corrected preview has 70 approved records and 50 explicit content holds.

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
