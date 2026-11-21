# Plan (final)

Adopted on 2026-11-21 as the plan of this repository; implementation has started (see **Implementation status** at the end). The text is the first review's `PLAN.final.md`, which merged the third-round decisions of all three reviews, with the owner's instructions of the same day applied: no external forks and no owner approvals, so the engine is kept as source in this repository (D1) and names are settled by a rule (D6). Baseline: `master` at 7ad5617; engine snapshots carneades-4 d19d543 and GoCHR e1ccd78.

Three independent reviews examined the repository and the two engine archives on 2026-11-20. In a second round each reviewer examined the other two and the source text: Walton, Reed and Macagno, *Argumentation Schemes* (2008), "the book" below. In a third round each reviewer examined the others' second-round work and made the decisions the earlier drafts had left to the owner. All three now say that no further broad review is needed and that implementation should start. This plan holds the decisions the three sets agree on. Where they still differed, a measurement, the book or the repository's own README decided, and the plan says so. Page and section numbers are the book's. It replaces `PLAN.md` (2026-11-06) and every plan proposed since.

## Where things stand

Measured with `tools/vv` (5,947 checks). Issue ids are those of its ledger, `tools/vv/known_issues.yml`.

| | result |
|---|---|
| schemes of `walton.yml` the `carneades` command can evaluate | 0 of 114 |
| schemes the command can evaluate from their category file | 54 of 114 |
| schemes the engine handles correctly through its Go API, one at a time | 103 of 114 |
| the same with the whole of `walton.yml` loaded | 101 of 114 |
| failing checks | 390, under 24 issues: 16 in the engine, 6 where an engine limit meets a scheme, 2 in the scheme files |
| entries of the book's compendium (ch. 9) | 109 in this plan's ledger, which splits the two numbers that hold several schemas (11.1 and 40). The book prints no total. The README's 106 follows from the same ledger in two ways: without that split, or without the three numbers that print a scheme of their own and also have sub-entries (4, 26, 35) |
| records of `walton.yml` | 114: 110 are filed under 108 of the 109 entries, 4 come from Carneades and have no entry |
| entries without a record of their own | 3: 50 has none; 57 and 22.4 are filed under records taken over from Carneades that hold other schemes |
| records with no assumption and no exception | 59; the book prints no question for 45 of them, and some for 13 |
| records with fewer assumptions and exceptions than the book prints questions | 31. These counts size the review; they do not say which questions a record answers |

Two engine defects are worse than a wrong label. Labelling never returns once a second `Infer` has added a position to an issue (E-ISSUE). And a scheme that re-derives a fact, or two schemes that derive each other's premises, run to 100,000 rule applications and starve every scheme listed after them (X-LOOP).

## Decisions

All three reviewers made these decisions independently in the third round. "All three" below means the three decision sets agree. Where they do not, the row says who decided what and what settled it. The owner can overrule any of them.

| # | Decision | Agreement |
|---|---|---|
| D1 | **Engine sources: kept in this repository, one commit per fix.** The two archives are imported byte for byte into `third_party/carneades-4` and `third_party/GoCHR` (commit c7d0eae; 846 files compared equal; the two Windows executables left out). Every engine fix is one commit that touches one engine and comes with its check in `tools/vv`, so that it can go upstream as it stands (`git format-patch -- third_party/GoCHR`). The engine is built against GoCHR e1ccd78, which every measurement used, not the ca9745c that carneades-4's `go.mod` names. `sources.lock` keeps both upstream commits, the SHA-256 of the two zips and the import commit; `tools/vv/setup.sh` refuses a mismatch and lists every engine file changed here. | All three chose forks; the owner chose not to create them. Same properties, one repository |
| D2 | **Critical questions: the book's three kinds, and an explicit questioned state.** An assumption holds until it is questioned, and then needs support like an ordinary premise; an exception counts only once it is shown (pp. 33, 373-375). A graph can list questioned statements (`questioned:` in YAML, and through the API). Assumptions are not turned into exceptions, and a counter-argument does not question anything by itself. | All three |
| D3 | **First supported profile: offline and sequential.** The `carneades` command and the Go API, one graph at a time, arguments generated from schemes. Schemes are propagation rules; ground facts form a set; constants are names, integers or quoted strings, with no arithmetic. The web server, concurrent use and CHR deletions or guards are outside, and the engine refuses them instead of the README warning about them. | All three |
| D4 | **Two releases, one migration.** R1 is a preview: the library runs. It repairs the nine schemes that have to change and adds what is missing. None of the nine can be evaluated through the command today, so repairing them breaks no file that works there, and no other identifier is touched. R2 is the first stable release: content checked against the book, and one coordinated migration of identifiers (misspelt ids and predicates, the four ids with a hyphen, file names) with an old-to-new table and aliases. | Third and first. The second puts every rename into R1. Settled by the README, see below |
| D5 | **The compendium of 2008 is the reference for every numbered entry.** The 24 records taken over from Carneades keep their wording and ids and are labelled with their own sources. Entries 50, 57 and 22.4 get records in the compendium's form. Where the book contradicts itself, its narrative chapters decide, and the reading is recorded as supported by the book, not as a confirmed erratum. | All three; the second would only flag 22.4 |
| D6 | **Scheme content of R1** as in Phase 3: the structure of each repair is agreed. Names of new predicates follow one rule: reuse a statement the library already has when it says the same thing (`advocates_conclusion/3`, `has/2`); otherwise take the name measured in the first review's trial (`goal_of_agent/2`, `conflicting_goals/2`, `reapplication_sequence/3`, `addressee_of/2`, `excuse_or_overriding_duty/2`, `is_a/2`); a new record's names describe its statement in the book's words. Each name is recorded with its page in the source ledger. A name found wrong later is changed in the one migration of R2 (D4). Both practical-reasoning records stay. | All three on structure; names by rule, on the owner's instruction |
| D7 | **One gate, after a parity register.** `tools/vv` is the gate, run by CI on every push. Every check the other two harnesses have is listed in `docs/parity.md` against the check that covers it, and is ported before that harness's evidence is set aside. The third review's runner is not run side by side: it unpacks and audits the two archives, so once fixes land in `third_party/` it would keep testing the unmodified engine. | Second and first on `tools/vv`. The third keeps its own runner as the front door during consolidation. See below |
| D8 | **Loops: set semantics in the rule store.** `RuleStore.Unique` is the starting design, with an index, and is not merged unchanged. Its contract comes first: ground facts form a set, arity is part of a fact's identity, distinct arguments for one conclusion are kept, deletions are outside the profile. Limits bound term size, match work, time and memory, not only the count of rule applications. | All three |
| D9 | **A run ends in one of seven ways:** complete, invalid input, inconsistent context, unsupported, limit reached, cancelled, internal failure. A failed run returns a typed error, leaves the graph unchanged and writes nothing to stdout. A partial graph is never presented as a result. | All three |
| D10 | **Negation is explicit and structural.** One encoding under an internal name users cannot write, applied as a statement enters the store, so that it works through statement variables. The library supplies the opposition of `S` and `¬S` as an issue scheme, once E-ISSUE is fixed. Contradictory assumptions are an inconsistent context, reported and not resolved silently. | All three |
| D11 | **Operator names:** the validator rejects `is`, `true` and `false` as predicate names. `in`, `or`, `div` and `mod` stay allowed. | Measured; the second would reject all of them. See below |
| D12 | **Arguments in an input file that name a scheme are never left outside the scheme silently.** A written-out argument is validated against its scheme with one substitution and has to carry the scheme's critical questions; if it does not, it is reported. A request to instantiate (a scheme and parameters, nothing else) is carried out or refused as unsupported. | All three on the rule. See below |
| D13 | **Scheme edits run beside the engine work.** Each change names its page, records how the passage is read, and comes with a positive and an adverse example written before the edit. | All three |
| D14 | **No further broad review.** Each change is reviewed against its tests. At gate S and before R1, someone who did not write the change reruns the strict suite from a clean checkout and checks the reference cases against the book by hand. | All three |

### Where the three decision sets differed

| Question | Second review | Third review | This plan | What settled it |
|---|---|---|---|---|
| Where renames go | All in R1, so that ids migrate once | One coordinated migration after the content is settled; previews before | R2 | The README says misspelt identifiers are kept "so that existing argument graphs keep working". With renames in R1 those graphs migrate twice, ids at R1 and predicates at R2; with renames in R2, once. R1 still gives a usable library early, which is what the second review wanted |
| The four ids with a hyphen | Rename in R1 | They work; normalise them in the coordinated migration | R2, with the others | Two of the four work through the command today. The check now asks what matters, whether an id comes back from the term reader unchanged, and all 114 pass it |
| Which script is the gate | `tools/vv` | Its own runner during consolidation; one orchestrator in the end | `tools/vv`, after a parity register; the other runner's checks are ported, not run beside it | The flaws the third review found in `tools/vv` are fixed: a checksum mismatch fails, runaway probes run in a process of their own, and a check on a half-built graph no longer passes. It is the only harness that runs every scheme through the API and the command. The third review's runner tests the archives, not the engine in `third_party/` |
| Where the engine lives | Forks | Forks | `third_party/` in this repository | The owner's instruction of 2026-11-21: no forks to create. One commit per fix and the lock file keep what the forks were for |
| Who approves names | The owner, in review | The owner, in review | The rule in D6 | The owner's instruction of 2026-11-21: no approvals. Names are recorded with their page; a wrong one is fixed in the R2 migration |
| Arguments written by hand | Complete them from the scheme, and instantiate on request | Validate a written-out graph; refuse a request until its contract exists; do not rewrite an authored graph | Validate and report; instantiate or refuse | The third review's objection: completing a written-out argument changes a graph its author wrote out on purpose |
| Operator names | Reject `in`, `or`, `div`, `mod` too; rename `in/2` in `abduction` | Define a policy | Reject only what the engine executes | Measured: only `is/2` misbehaves, and only in a conclusion. `abduction` is a Carneades record that upstream ships with `in/2` |
| `pop_scheme` | One record that concludes `A` and `should_accept(A)` | Two variants, one for each reading | `A` replaces `true(A)` in R1; two variants in R2 | The book offers the two conclusions as alternatives (p. 311). Splitting the record changes its interface, so it goes with the migration |
| Citations | Correct two page numbers from the compendium | Record the difference; check the originals first | Record | The compendium cites both works at second hand |
| Entry 57 and the id `precedent` | Give the record its old id back | Keep the 2019 rename until its source is checked | Add 57 as its own record; the inherited record keeps the id `exeption_to_rule` | The README records the rename (2019-10-08) as following Douglas Walton's correspondence of 19 September 2019. That note is the owner's own record; the id stays, and 57 gets a record of its own |

### What the owner still has to do

Nothing. The forks are replaced by `third_party/` (D1), names are settled by the rule in D6, and the README's own note settles the id of the inherited `precedent` record. The second review (Claude) implements the plan in this repository; the gates in D14 ask someone who did not write a change to rerun it, which either of the other two reviews can do from a clean checkout.

## Rules this plan follows

* **Engine first.** A green check means little while the engine can drop a scheme, hit its limit or lose an argument without saying so.
* **Change a scheme only where the scheme is wrong, and only with its row in the source ledger written.** The row gives the page, what the book has, how it is read and what kind of claim the change makes (D5). Negated premises and the composition/division pair are what the book says (6.4, 11.1, 26; 9.1, 10.1). They stay.
* **A scheme is judged by its own argument.** Whether a conclusion keeps other support is a property of the whole theory. A check that leaves a premise out or adds an exception counts only if the plain run finished and built the argument.
* **A check changes only with its contract.** Several items below allow two outcomes: make an input work, or reject it with a message. Six checks of the suite demand the first today (its README lists them). Where the second is decided, the contract of Phase 0 says so first and the check is restated in a change of its own. The ledger carries `min_checks`, so a check that disappears fails the run.
* **No prototype is merged whole.** Each fix is selected by its regression check. The combined prototype engine breaks one check that passes today: it takes a user's fact `neg(p(a))` for a negation.
* **A repair is not a transcription.** The scheme repairs of R1 make schemes run and bind the right participant. Where a role, a polarity or a conditional question is still open, the ledger row says "executable interpretation", and R1 is called a preview for that reason.

## Phase 0: baseline, parity and contracts

Starts now. No behaviour changes.

1. Import the engine sources and write the lock file (D1). Add `tools/vv` and its CI workflow; add `__pycache__/` to `.gitignore`. CI builds with Go 1.24 and, without blocking, the latest Go.
2. Write the parity register (D7), `docs/parity.md`: every check of the second review's `tools/go` and of the third review's `tools/enginecheck`, with the check of `tools/vv` that covers it or the note "to port". Known gaps today: the second review's six acceptance cases and its run with all premises at once; the third review's per-scheme checks of a withdrawn assumption and of a second inference, its Go contracts for concurrency and for the store GoCHR returns when no rule fired, its race job, its 20 probes of `tools/validate_schemes.py` and its two mutation tests; fresh-process round trips. Port them, then retire the other two harnesses.
3. Harden the ledger: an expected-check list in place of a count, and a fingerprint of each known failure's message, so that a failure which changes its cause is noticed.
4. Bring the README up to date. Give the mapping to the compendium from the ledger instead of a single count. Add that the file holds 114 records, that three compendium entries have no record of their own, that upstream ships 24 records, not 25, that `roles` is not read by the importer, and that `eval` refuses the full library today.
5. Write down seven contracts, each one page, in `docs/contracts.md`:
   * the supported profile (D3), and the rules for terms and names (end of this file);
   * what counts as one fact and one argument (D8): ground statements form a set, and an argument is identified by its scheme and its canonical parameters;
   * how a run can end (D9);
   * negation and opposition (D10);
   * critical questions and their states (D2): undisputed, questioned without support, questioned and supported, exception shown, rebutted in an issue, question withdrawn. Questioning suspends a scheme's default; it asserts no negation and labels nothing out. A statement the graph itself stipulates stays accepted until the stipulation is withdrawn, so the origin of an assumption has to be kept, also through export. The state survives a second `Infer`;
   * arguments in an input file that name a scheme (D12);
   * names the validator rejects (D11).
6. Extend `tools/vv/source_ledger.yml` to one row per record: book or upstream; the source cited when it is not the book; each printed question sorted as in 5.1; the kind of claim (transcription, correction supported by the book itself, executable interpretation, policy of the application, not yet checked against its source); and how far the row has been checked (mapped, transcription checked, interpretation approved, questions sorted, example passed).

Exit: a clean checkout runs `tools/vv/run.sh` green in CI. Today that is 5,947 checks with 390 known failures; item 2 adds checks.

## Phase 1: the engine says when it fails

In `third_party/`, one commit per fix. No scheme file changes.

1.1 Exit status: `check` and a refused `eval` exit non-zero; `eval` stops ignoring the error `Infer` returns (E-EXIT). Proposed codes: 1 for rejected input, 2 for an incomplete inference, 3 for an internal error.
1.2 Limits (E-BUDGET): GoCHR returns a limit status, `Infer` a typed error, and the graph stays unchanged. Nothing is written to stdout. A run counts as interrupted only if work remains at the limit. The limit bounds the work, not only the count of rule applications: a theory whose terms keep growing exhausts 2 GB in two seconds today. A run can be stopped from outside, so that a caller's timeout ends it.
1.3 Unreadable input is an error and leaves the graph unchanged: a statement the parser cannot read, rule text GoCHR accepts although it is malformed, an `AddRule` error Carneades discards (E-PARSE). With it, GoCHR returns the store it was given when no rule fired, not an empty one.
1.4 The importer keeps no state between calls (E-GLOBALS). Copying the two maps stops the leak; concurrent use needs every package-level variable gone, which is outside R1.
1.5 The validator checks an argument against its scheme with one substitution across premises, assumptions, conclusion and parameters, after 1.4, which today hides its faults (E-ARGVAL). The prototype fix accepts premise `p(a)` with conclusion `q(b)` for the scheme `p(X)`, `q(X)`.
1.6 `Inconsistent()` answers from the assumptions as they are, not from a cache that only labelling fills (E-INCONS). `check` and `eval` report an inconsistent context as its own outcome: today a graph that assumes both `p(a)` and `¬p(a)` labels both in, with the issue between them in place, and exits 0.
1.7 The module builds, vets and tests as shipped: two stale packages, vet findings, a test that needs CouchDB, a label test that only prints (E-BUILD).

Exit: those seven issues leave the ledger.

## Phase 2: the engine runs what the schemes say

2.1 **Issues** (E-ISSUE), first, because it can stop a run for good: a statement that joins an issue is added to its positions, and labelling returns.
2.2 **Set semantics for ground statements** (X-LOOP, D8): `RuleStore.Unique`, 48 changed lines, experiment 1, as the starting design. Its tests come before the change: several arguments for one conclusion kept, arities kept apart, a fact deleted and added again, GoCHR's own multiset programs unchanged while the flag is off. Its duplicate test needs an index, and dead entries pruned. Gate: the counting theory of `tools/vv` stops at the limit within twice the time of the unmodified engine (2.5 s today; 17 s with the patch as it is). The alternative, one extra rule per predicate, is dropped by its own author.
2.3 **Matching** (E-MATCH): a fact with a premise's functor at another arity must not make a rule fire until the limit, and a premise that is a bare variable must match wherever it stands, or be rejected by the validator.
2.4 **Negation** (X-NEG, E-NEGTEXT, D10), with a chosen and tested rule for double negation, and language templates applied inside `¬`.
2.5 **Repeated inference** (E-IDEMP): no panic on a graph with exceptions; one identity per argument, so that a second `Infer` adds nothing and does not assume again what was withdrawn. Today 98 of the 104 schemes that fire fail this: 32 panic, 66 duplicate.
2.6 **Names and terms** (E-CAPTURE, X-IS engine side, E-TERMS): fresh variable names that cannot collide with declared ones; `is/2` either usable as a predicate or rejected by the validator; a term means the same to both parsers, so that a quoted string with spaces survives and an expression such as `1990-05-17` in a statement is kept as written or rejected, not computed.
2.7 **Questioned statements** (X-ASSUME, D2), to the contract of Phase 0. The prototype of experiment 4 is 43 changed lines and breaks no check; it drops a questioned statement from the assumptions whatever its origin, which the contract changes.
2.8 **Arguments that name a scheme** (E-EXPLICIT, D12).

Exit: every engine check of those issues passes. Their checks on schemes pass once Phase 3 is in. `static/no-negated-premise` is removed from the suite with 2.4.

## Phase 3: scheme files, release R1

Runs beside Phases 1 and 2 (D13). On the engine as shipped these edits alone take the command from 0 to 108 of 114 schemes. `tools/vv/experiments/phase1_trial.py --final` holds the nine repairs; the first review's `final-nine-scheme-edits.patch` shows them. Names follow D6. Where a statement the library already has can do the binding, it is used.

3.1 `ad_hominem` (42, p. 336): add the premise that the person advocates the argument, with `advocates_conclusion/3`, which schemes 43 and 45 already use. Adverse examples: another person's argument, another argument of the same person.
3.2 `regulative_rule` (55.3, p. 343): two records, obligation and prohibition, with the polarity of the premise and the warrant. The existing record is repaired as the obligation variant, which is what the trial measures; the prohibition variant is added beside it. Each has a premise that its addressee is subject to the rule, which the book does not supply: "S must do A" does not follow from "A is obligatory" without it. The excuse names its subject; the warrant is general and stays as it is. A prohibition is a duty not to act, not the negation of a duty to act.
3.3 `practical_necessary_condition`, `practical_sufficient_condition` (22.2, 22.3, pp. 323-324): the goal premise names the agent. The "other goals" of the last question are not named, so that statement loses a place and the parameter `A2` goes; what counts as evidence for it is written down. The two statements whose arity changes get new names, because a fact left at the old arity can stall a rule (E-MATCH).
3.4 `sorites_slippery_slope` (51, pp. 340-341): the sequence premise gets its two ends and its property. Adverse examples: a sequence for another property, a sequence between other individuals.
3.5 `pop_scheme` (4.1, p. 311): `A` replaces `true(A)` (L-TRUE). The record's second conclusion stays until R2.
3.6 `group_2` (21.2, p. 322): `has(M,Q)`, the statement variant 1 uses for the group. The premises of the two variants differ, so no cycle arises. `is/2` in `rule_physical_world_1` is renamed with it; there it stands in a premise and works today.
3.7 `exeptional` (56, p. 344): the major premise is a generalization, and the record states it for the case at hand. The assumption that repeats the minor premise goes; it is what makes the rule feed itself. The scheme's three questions are conditional on one another and are modelled in Phase 5, not replaced by one unconditional assumption.
3.8 Add records for scheme 50, the precedent slippery slope argument (p. 340), for scheme 57, argument from precedent (p. 344), and for the three-premise form of 22.4 (p. 324), each with its questions sorted as D2 says (L-MISSING). The two inherited records filed there today stay, labelled with their own sources.
3.9 Add the negation issue scheme (D10), after 2.1. `tools/build_category_files.py` has to copy it into the category files, which it does not do today.
3.10 `tools/validate_schemes.py`: the engine's rules as errors for the supported profile (range restriction; `is`, `true` and `false` as predicate names; an assumption identical to a premise; templates whose count of `%` differs from the arity); a conclusion identical to a premise and a cycle across schemes are warnings; the 17 gaps the third review's probes found; and, in place of checks on predicate names, instances taken from the book for the 11 corrections of the changelog and the repairs above, with the expected statements written out, so that a flipped negation or two swapped arguments fail.
3.11 Record two citation differences without changing them: `sorites_slippery_slope` cites p. 151 where the compendium has p. 251, `exeptional` p. 148 where it has p. 147.

**Gate S, every scheme runs** (Phases 1, 2 and 3 done): `carneades check` prints nothing for all 14 files; the command evaluates every scheme of `walton.yml` and of the category files, 114 today and 117 with 3.8; the ledger holds only E-EXPORT and E-ROLES. This is a milestone, not a claim of stability.

Measured ahead of time with prototypes (experiments 3 and 4). "Prototype engine" is the second review's patch with its per-predicate rules replaced by 2.2.

| scheme files | engine | failing checks of 5,947 | `walton.yml` by command | category files by command |
|---|---|---|---|---|
| as they are | as shipped | 390 | 0 of 114 | 54 of 114 |
| as they are | prototype engine | 301 | 0 of 114 | 58 of 114 |
| second review's six edits | as shipped | 150 | 106 of 114 | 108 of 114 |
| second review's six edits | prototype engine | 53 | 113 of 114 | 113 of 114 |
| the nine repairs of this phase | as shipped | 122 | 108 of 114 | 110 of 114 |
| the nine repairs of this phase | prototype engine | 35 | 114 of 114 | 114 of 114 |
| the nine repairs of this phase | prototype engine with the `questioned:` list | 34 | 114 of 114 | 114 of 114 |

In every row with the prototype engine one check fails that passes today, the one on a user predicate called `neg`. No other check newly fails. The 34 are that one, 22 engine checks no prototype passes yet (1.1 for `check`, 1.2 for growing terms, 1.3, 1.5, 1.6, 1.7, 2.1, 2.3, 2.4 for a negation through a statement variable, 2.5, 2.6, 2.8, 4.2, 4.3), 8 static checks that go with 2.4, and 3 for the entries of 3.8.

## Phase 4: a file format that round-trips, and reference cases

4.1 Export the theory with the graph, and quote parameters: `liable(acme,smith)` comes back today as two parameters (E-EXPORT).
4.2 Expected labels survive: the export writes `labels:`, the importer reads `tests:` (E-EXPORT). The `questioned:` list and the origin of each assumption survive too.
4.3 `roles`: import them or reject them (E-ROLES).
4.4 Reference cases written by hand, with the expected statements, arguments and labels spelled out, never generated by the code under test: linked support, a missing premise, each state of a critical question (Phase 0), an undercutter, a rebuttal inside an issue, proof standards, two conclusions, a cycle, each way a run can end, two cases side by side in one graph, several schemes that feed each other, one unrelated fact. Each case states what happens to the argument and what happens to the conclusion.
4.5 Property tests where a parser or a printer is involved: terms through both parsers, renaming of variables, export followed by import in a new process.

Exit: a graph exported and read back in a new process has the same theory, parameters, expected labels and recomputed labels, and passes the validation of 2.8.

**Gate R, release R1, a preview** (gate S and Phase 4 done): the strict run passes for the profile of D3. R1 says of itself that the engine profile is accepted and the library is not yet: its repairs are interpretations until Phase 5 has checked them.

## Phase 5: content against the book, release R2

The reading and the ledger start now; edits wait for gate S. `go run ./cmd/schemecheck -format source` prints the worklist.

5.1 **Sort every printed critical question** into one of four kinds (D2): it restates a premise and is not modelled (p. 374); it shifts the burden when asked and is an assumption; it needs backing and is an exception; it attacks the conclusion and is a rebuttal inside an issue.
5.2 **Questions the book prints and the library may not model**: 13 records have no assumption or exception although the book prints questions for them (`verbal_classification`, `verbal_classification_arbitrariness`, `negative_normal_expectations`, `abductive_forward`, `positive_consequences`, `negative_consequences`, `negative_consequences_reasoning`, `bias`, `bias_ad_hominem`, the two slippery-slope records, `appearance`, `memory`), 31 have fewer than the book prints questions. 5.1 decides which of these are gaps. The other 45 records without either have no question in the book.
5.3 **Variants and roles the R1 repairs leave open.** In the practical pair `X` stands both for the set of alternatives and for the member selected (pp. 95-96, 323), and the alternatives premise of the sufficient scheme should say "sufficient" (p. 96). `pop_scheme` becomes two records, one for each reading. `exeptional` gets its three conditional questions. `ad_hominem` has to say whether the argument is rejected or only loses credibility (p. 336). `bias` (46, p. 338) has the stray-variable pattern fixed in 3.7.
5.4 **Statements without variables**: 22, in 16 schemes, decided one by one. A statement about a participant or the case gets that variable; a general warrant such as `do_obligator` stays global. Each change comes with a test of two cases in one graph.
5.5 **The 24 records from Carneades**: compare each with the source it cites, not with the compendium. The inherited `precedent` record keeps the id it has until the 2019 correspondence is checked. Smaller points: `epistemic_ignorence` has four of the five premises of 27 (p. 328); `expert_opinion` has one of the two assumptions of the book's own Carneades example (pp. 374-375).
5.6 **Identifiers, vocabulary, wording: the one migration** (D4). Misspelt ids, predicates and file names; the four ids with a hyphen; shared predicates by meaning and role, not by equal text (337 of 394 predicates are used by one scheme; statement variables already connect 14 schemes); "Therefore," in 40 of 93 conclusion texts. One old-to-new table, with aliases or a script, and graph fixtures that are migrated in a test. `not_fit_action` has a positive text and the right logic (p. 330): rename, do not invert.
5.7 **One instance per scheme taken from the book**, with the expected result written by hand: the premises hold and the conclusion is in; each exception undercuts the argument; each premise missing, no argument.
5.8 **Readings the book supports against its own print**, each marked as such and not as a confirmed erratum: `W'` in the minor premise of 11.2 (p. 318); the glosses of the questions of 9.1 and 10.1 interchanged (pp. 316-317; the narrative on p. 113 states composition in the direction the library uses); "necessary" in the alternatives premise of 22.3 (p. 323; p. 96 says sufficient), which the library's record repeats; the reversed pair in the conclusion of 55.3 (p. 343).
5.9 **A source for every record**: entry number and page in `meta.source`. Three records have none today and one has no title.

**Gate L, release R2, the first stable release:** the strict run passes; every released record has its ledger row at "example passed"; records that are not there yet are visibly outside the release.

## Phase 6: releases and upstream

* R1 after gate R, R2 after gate L, each with its table of changes. Version the engine in `third_party/` separately from the library. Record the toolchain: the per-scheme results agree between Go 1.24.7 and 1.27.4.
* At R1 the zips can leave the tree; the lock file keeps their hashes, and the zips stay attached to the release.
* Document how the library is used: Carneades has no include mechanism, so a graph file must carry the language and the schemes it uses.
* Offer each engine commit upstream, and the scheme corrections to the original repository. Send the readings of 5.8 to the authors. Upstream last changed engine code in 2022, so do not wait for it.

## What each source contributed

| source | in this plan |
|---|---|
| all three | engine first, with the source ledger started at once; engine fixes one per commit, pinned and locked; the questioned state; the offline profile; set semantics as a declared profile; the seven ways a run can end; no prototype merged whole; negated premises, the composition/division pair and both practical-reasoning records stay; no further broad review |
| first review | `tools/vv` and its ledger; the command as the measure; set semantics in the rule store; rules for terms; the map of the compendium; the measurement that only `is/2` misbehaves among the operator names |
| second review | its engine patch as a measurement for Phases 1 and 2; six of the nine repairs from the book; the `questioned:` list; `tools/vv` as the one gate; the gate for 2.2; exit codes; the independent rerun at each gate; one instance per scheme; the reading that 106 is a way of counting |
| third review | contracts before content; the strict gate and the parity register; judging a scheme by its argument, with a finished control run; the counterexamples behind 1.2, 1.5, 1.6, 2.2, 2.3 and 2.4; contradictory assumptions as their own outcome; the two kinds of argument in an input file; one migration of identifiers; R1 as a preview; the 2019 note on `precedent`; reference cases written by hand; the stages of a ledger row |
| the book | the repairs; schemes 50, 57 and 22.4; the burden of proof behind D2; convergent arguments behind keeping both practical records; four readings against the print |
| the README | misspelt identifiers are kept so that existing graphs keep working, which decided D4; the 2019 correspondence |
| the owner, 2026-11-21 | no forks and no approvals: the engine in `third_party/` (D1) and the naming rule (D6) |

## Rules for terms (for users of the library)

* Constants start with a lower-case letter; a word starting with a capital is a variable, and a non-ground assumption blocks evaluation.
* Write a date as a quoted string: `"1990-05-17"` survives. Unquoted it is an arithmetic expression, and the argument is built for `1968`. Use underscores in names for the same reason.
* No quoted strings containing spaces: `"it rains"` is cut off at the space when the file is read.
* In YAML, write lists of terms one per line. `[p(a,b)]` is read as the two strings `p(a` and `b)`.
* Do not use `is` as a predicate name: in a conclusion the rule engine executes it.
* Do not assume both a statement and its negation: `check` and `eval` report such a graph as an inconsistent context (exit 3).

## Implementation status

Updated with each change. A box is ticked when its change is on `master` and CI is green.

- [x] 0.1 Engine sources imported into `third_party/`, `sources.lock`, `tools/vv`, CI workflow
- [x] 0.2 Parity register `docs/parity.md` written
- [ ] 0.2 Checks marked "to port" in `docs/parity.md` ported
- [ ] 0.3 Ledger with an expected-check list and failure fingerprints
- [x] 0.4 README brought up to date
- [x] 0.5 Contracts `docs/contracts.md` written (each is met when its issues leave the ledger)
- [ ] 0.6 Source ledger, one row per record
- [x] 1.1 Exit status by outcome (E-EXIT); a panic exits with 6
- [x] 1.2 Limits as a typed error, graph unchanged, nothing on stdout (E-BUDGET); a term-size limit; `InferContext`
- [x] 1.3 Unreadable input is an error (E-PARSE); GoCHR returns the goals when no rule fires
- [x] 1.4 The importer keeps no state between sequential calls (E-GLOBALS)
- [x] 1.5 The validator checks arguments with one substitution (E-ARGVAL)
- [x] 1.6 `Inconsistent()` from current assumptions; inconsistent context as an outcome (E-INCONS)
- [x] 1.7 Build health (E-BUILD): both modules build, vet and test as shipped; CI runs them
- [x] 2.1 A statement joining an issue becomes its position; labelling returns (E-ISSUE)
- [x] 2.2 Set semantics for ground statements (X-LOOP); the engines' own unit tests run in CI
- [x] 2.3 Matching (E-MATCH): a propagation solver with a real propagation history for set-semantics programs
- [x] 2.4 Negation (X-NEG, E-NEGTEXT): GoCHR `ExplicitNegation`; `static/no-negated-premise` removed from the suite
- [x] 2.5 Repeated inference (E-IDEMP); per-scheme `infer-twice` checks ported
- [x] 2.6 Names and terms: E-CAPTURE, E-TERMS; deletions and guards unsupported, arithmetic invalid (contract 1). X-IS's engine side (the validator rejects reserved names) lands with the library repair 3.6
- [ ] 2.7 Questioned statements (X-ASSUME)
- [ ] 2.8 Arguments that name a scheme (E-EXPLICIT)
- [x] 4.1, first half: export writes the theory (the parameter quoting is still open)
- [ ] Phase 3: 3.1 to 3.11; gate S
- [ ] Phase 4; gate R; release R1
- [ ] Phase 5; gate L; release R2
