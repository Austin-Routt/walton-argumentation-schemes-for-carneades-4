# Experiments

Four experiments, none applied anywhere. They were measured against the engine
archives before the engine sources were imported into `third_party/`; to repeat
one, apply its patch to a scratch copy of the repository, never to a working
branch. Each one answers a question about the
plan by measuring it with the suite, against carneades-4 d19d543, GoCHR e1ccd78
and the scheme files at 7ad5617. 

## 1. Set semantics for argument construction (engine)

Question: can issue X-LOOP in `known_issues.yml` be fixed in the engine with a
small change, and what does the change cost?

* `set-semantics-gochr.patch` (45 lines): adds `RuleStore.Unique`. When it is
  set, a ground constraint equal to one already alive in the store is not added
  again. The default keeps CHR's multiset semantics, so GoCHR's own programs
  (gcd, primes, ...) are unaffected.
* `set-semantics-carneades.patch` (3 lines): `TheoryToRuleStore` sets it.

| | unmodified | with both patches |
|---|---|---|
| two schemes deriving each other's premise | stops at 100000 rule applications after 0.8 s; the scheme listed after them never fires | 3 arguments at once, including that scheme's |
| `exeptional`; `composition_generic` with `division_generic` | rule limit hit | terminate |
| this suite | 390 known failures | 15 of them pass, no check newly fails |
| GoCHR's own tests (107) | pass | pass |
| Carneades' own tests without CouchDB | pass | pass |
| a scheme that counts upwards for ever | gives up after 2.5 s | gives up after 17 s |

The last row is the cost: deleted constraints stay in GoCHR's index lists, so
the duplicate check gets slower as a run gets longer. Pruning deleted entries,
or an index, would remove that. How much it costs on long runs that do finish
has not been measured.

What the patch does not do, as the third review showed with bounded probes and
this suite now checks: a theory that builds ever larger terms still runs out of
memory (`growing-term-is-bounded`), and a fact with a premise's functor at
another arity still drives a rule to the limit (`wrong-arity-fact-is-harmless`).
Neither involves a repeated fact. The patch also changes what a rule sees that
consumes two copies of one fact: with `Unique` there is only one. Carneades sets
the flag for every theory, so the supported profile has to say that ground
statements form a set, and what that means for schemes with `deletions`.

The other way to get set semantics is from the Carneades side alone: a rule
`p(V1,...,Vn) \ p(V1,...,Vn) <=> true` for every predicate, ahead of the scheme
rules. A second, independent review of this repository proposed exactly that in
its engine patch. It ends the same loops. On the counting scheme of the last
row, however, the run does not stop at the rule limit: GoCHR's memory use per
attempted match grows until the process dies (stopped here at 2.3 GB under a
memory cap; without the cap an 8 GB machine killed it). A limit that exists to
bound runaway theories has to keep working, so the store-level patch above is
the one the converged plan uses.

To try it:

```
cd third_party   # in a scratch copy of the repository
patch -p1 -d GoCHR       < ../experiments/set-semantics-gochr.patch
patch -p1 -d carneades-4 < ../experiments/set-semantics-carneades.patch
cd .. && go test         # fails, listing the ledger entries that can now be removed
```

## 2. Trial edits to the scheme files (library)

`phase1_trial.py` edits a copy of the repository. It has three modes.

**Without an option: the library-only route.** Question: do scheme-file edits
alone make every file evaluable by the `carneades` command, with the engine as
it is? Every engine limit is worked around in the files: the five X-RANGE
schemes get binding premises, `true/1` and `is/2` are renamed, the four negated
premises become positive predicates, `exeptional` loses its repeated assumption,
`composition_generic` gets a conclusion that is not the premise of
`division_generic`, and the four hyphenated ids are renamed (14 schemes edited,
four ids renamed). Two entries of the regression table in
`tools/validate_schemes.py` have to follow: it names `opposition-normative` by
its old id and requires the composition/division pair. `--keep-pair` leaves that
pair alone.

Measured on 2026-11-20 with the first version of this suite (5,176 checks):

| | failing checks | `walton.yml` by command | category files by command |
|---|---|---|---|
| repository as it is | 349 | 0 of 114 | 54 of 114 |
| trial edits | 35 | 114 of 114 | 114 of 114 |
| trial edits with `--keep-pair` | 39 | 112 of 114 | 114 of 114 |
| trial edits with `--keep-pair`, engine with the patches of experiment 1 | 32 | 114 of 114 | 114 of 114 |

The answer is yes. The price is that the files stop saying what the book says in
two places: the book states those four premises as negations (6.4, 11.1 and 26;
pp. 315, 317-318, 327), and its composition and division schemes are converses
over the same two statements (9.1 and 10.1, pp. 316-317). With the current suite
the renamed ids also have to be renamed in `source_ledger.yml`; the suite says
so.

**`--converged`: the route of the converged plan.** Only what is wrong in the
scheme files themselves is changed, and the book decides where it says
something: the generic ad hominem takes its binding premise from its sibling
schemes 43 and 45; the practical-reasoning pair binds its agent in the goal
premise and no longer names the "other goals"; the sorites scheme gives its
sequence premise the two ends it connects; `exeptional` loses the variable its
major premise never shared with the rest and the assumption that repeats its
minor premise; `true/1` and `is/2` are renamed. Nine schemes; `walton.yml` gains
21 lines and loses 20.
Negated premises, the composition/division pair and the ids stay. The premise
added to `regulative_rule` is a placeholder: the book has none to offer.

**`--final`: the edit set of the final plan.** The same nine schemes as
`--converged`, after three rounds of review. Where a statement the library
already has can do the binding, it is used, so that no new predicate is added:
the generic ad hominem takes `advocates_conclusion/3` from its sibling schemes
43 and 45, and `group_2` concludes with `has/2`, as variant 1 does. The second
review had proposed new predicates for both and withdrew them in its third
round. From that review's book-based patch come the regulative rule (a premise
saying that its addressee is subject to the rule, and an excuse that names its
subject) and the pop scheme (it concludes the statement itself where it
concluded `true/1`). The practical pair binds its agent in the goal premise and
gets new names for the two statements whose arity changes, because a fact left
over at the old arity can stall a rule (issue E-MATCH). The sorites sequence
names its two ends and its property, as the third review asked. `exeptional`
and `rule_physical_world_1` are edited as in `--converged`. `walton.yml` gains
22 lines and loses 22.

The third review proposed no edits of its own. It set conditions for them: bind
the participant the book means, not a convenient one, and write the reading
down before the edit. These edits bind whose argument is attacked, which agent
acts, who is subject to the rule and which sequence reaches the target. Three
questions that review also names stay open for the content phase: the
alternatives against the action selected in the practical pair, obligation
against prohibition in the regulative rule, and whether the pop scheme's two
readings become two records. The plan requires the written reading before
each edit, and calls these repairs interpretations, not transcriptions.

Of the operator names of the rule engine only `is/2` misbehaves, and only in a
conclusion: `in`, `or`, `div` and `mod` were tried as predicates in both
positions and work, so `in/2` in `abduction` stays.

**All new predicate names and texts are proposals** that test the mechanics.
`is_a/2` in particular is a placeholder.

## 3. The converged route (engine and library together)

Question: do the engine fixes prototyped by two reviews, taken together, and the
`--final` edits make all 114 schemes evaluable by the command, without
changing what the files take from the book?

The engine is built from three patches:

1. the second review's `tools/go/patches/engine-fixes.patch` (not shipped here;
   it is part of that review): rule limit reported as an error, exit status of
   `eval`, negated premises encoded as `neg(...)`, texts inside `¬`, importer
   maps copied, argument validation, the nil check in `Infer`, theory and quoted
   parameters in the export, `Inconsistent()`;
2. `set-semantics-gochr.patch` from experiment 1;
3. `merged-use-store-set-semantics.patch` (3 lines): switches the first patch's
   per-predicate rules off and `RuleStore.Unique` on.

Measured on 2026-11-21 with this suite (5,947 checks). In every row with the
three patches one check fails that passes on the engine as shipped:
`baseline-user-predicate-neg-is-ordinary`. The first patch encodes `¬p(a)` as
`neg(p(a))`, so a user's own fact `neg(p(a))` sets off a negated premise. The
third review found this; the numbers below include it. No other check newly
fails in any row.

| scheme files | engine | failing checks | `walton.yml` by command | category files by command |
|---|---|---|---|---|
| as they are | as shipped | 390 | 0 of 114 | 54 of 114 |
| as they are | patches of experiment 1 | 375 | 0 of 114 | 54 of 114 |
| as they are | three patches | 301 | 0 of 114 | 58 of 114 |
| second review's patch (six schemes) | as shipped | 150 | 106 of 114 | 108 of 114 |
| second review's patch (six schemes) | three patches | 53 | 113 of 114 | 113 of 114 |
| `--final` edits (nine schemes) | as shipped | 122 | 108 of 114 | 110 of 114 |
| `--final` edits (nine schemes) | three patches | 35 | 114 of 114 | 114 of 114 |
| `--final` edits (nine schemes) | three patches and experiment 4 | 34 | 114 of 114 | 114 of 114 |

The answer is yes for the command, and no for the engine as a whole. In the
`--final` rows both validators are silent. The 35 checks that still fail with
the three patches are:

* 23 engine checks that no prototype passes: one substitution for a whole
  argument in the validator (E-ARGVAL), a growing term (E-BUDGET), stale
  packages and the CouchDB test (E-BUILD, 2), variable capture (E-CAPTURE), the
  exit status of `check` (E-EXIT), arguments written by hand (E-EXPLICIT, 2),
  expected labels in the export (E-EXPORT), a second `Infer` duplicating
  arguments and assuming a withdrawn assumption again (E-IDEMP, 2),
  `Inconsistent()` before labelling and after a withdrawal (E-INCONS), issue
  membership and the labelling that never returns (E-ISSUE, 2), wrong arity
  and premise order (E-MATCH, 2), unreadable statements (E-PARSE), roles
  (E-ROLES), terms (E-TERMS, 2), a questioned assumption (X-ASSUME), `is/2` as
  a predicate (X-IS), a negation through a statement variable (X-NEG);
* 1 check the first patch breaks: a user predicate called `neg`;
* 8 static checks that flag negated premises; they become obsolete, and should
  be removed, once the engine matches them;
* 3 for compendium entries without a record of their own (L-MISSING).

Seven of those checks exist because the other two reviews showed a fix in the
first patch to be partial (E-ARGVAL, E-INCONS, the negation through a statement
variable, the predicate `neg`) or missing in both prototypes (E-BUDGET for
growing terms, E-MATCH, the labelling that never returns).

To repeat it, from the repository root, with `OPUS_PATCH` set to the path of
the second review's `engine-fixes.patch`:

```
cd third_party   # in a scratch copy of the repository
patch -p1 < "$OPUS_PATCH"
patch -p1 -d GoCHR < ../experiments/set-semantics-gochr.patch    # reports fuzz and offsets; expected
patch -p1 < ../experiments/merged-use-store-set-semantics.patch
cd ../../..
mkdir /tmp/walton-conv && git archive HEAD | tar -x -C /tmp/walton-conv
python3 tools/vv/experiments/phase1_trial.py /tmp/walton-conv --final
python3 /tmp/walton-conv/tools/build_category_files.py /tmp/walton-conv
python3 /tmp/walton-conv/tools/validate_schemes.py /tmp/walton-conv | tail -1
cd tools/vv && WALTON_REPO=/tmp/walton-conv go test
```

The last command fails by design: it lists the 356 ledger entries that no longer
fail, and the one check the first patch breaks. In the scratch copy only
`third_party/` is changed.

## 4. A list of questioned statements (engine)

Question: can a critical question written under `assumptions:` be called into
question, as the book has it, without turning it into an exception?

The book's account (ch. 11, sec. 5, pp. 373-375): an assumption holds until the
other side calls it into question; from then on it has to be supported like an
ordinary premise. An exception is different: it counts only once it has been
shown. In the engine as shipped a scheme assumption is assumed when the scheme
is instantiated, and nothing in a file can withdraw it.

`questioned-assumptions.patch` (43 changed lines, Carneades only) adds a list
`Questioned` to the argument graph and the key `questioned:` to the YAML format.
A statement on the list is not assumed, whether a scheme lists it under its
assumptions or the graph under `assumptions:`. It is then labelled like any
other statement: in if an argument supports it, otherwise not.

```
argument_schemes:
  - id: s
    variables: [X]
    premises: [p(X)]
    assumptions: [a(X)]
    conclusions: [c(X)]
assumptions: [p(x)]
questioned: [a(x)]        # c(x) is no longer in, until something supports a(x)
```

| | engine as shipped | with the patch |
|---|---|---|
| `baseline-undisputed-assumption-holds` | pass | pass |
| `questioned-assumption-is-not-assumed` | fail (X-ASSUME) | pass |
| `questioned-assumption-holds-when-supported` | pass | pass |
| this suite | 390 known failures | 1 of them passes, no check newly fails |
| Carneades' own tests without CouchDB | pass | pass |

It applies on top of the three patches of experiment 3 as well (last row of the
table there). It is a prototype: the export writes the list, but nothing checks
a round trip yet, and the validator does not know the key.

```
cd third_party   # in a scratch copy of the repository
patch -p1 -d carneades-4 < ../experiments/questioned-assumptions.patch
cd .. && go test         # fails, listing the one ledger entry that can be removed
```
