# Using the library

This page shows how to evaluate a case of your own with the schemes of this
repository (PLAN.md, Phase 6). The examples under `docs/examples/` are checked
by `tools/vv` on every push: each is what the tool below makes from the
library today, and the command gives every label its `tests:` expect.

## What to load

* `walton.yml` holds every scheme (120 records) with the language templates
  they use and the negation issue scheme. It is the only hand-edited scheme
  file.
* The category files (`A_reasoning/`, `B_source-based-arguments/`,
  `C_applying-rules-to-cases/`) are generated from it; each loads on its own.

Carneades 4 has no include mechanism: a graph file has to carry the language
and the schemes it uses. Either write your case at the end of a copy of
`walton.yml` or of a category file, or start a file with only the schemes you
need:

```sh
python3 tools/graph_from_library.py expert_opinion > case.yml
python3 tools/graph_from_library.py practical_inference negative_consequences_reasoning --title "Should I take the job abroad?" > case.yml
```

The tool copies the schemes, the templates they use and the issue schemes
verbatim from `walton.yml` (or from the file given with `--library`), records
the library's version under `meta.library`, and ends with a commented outline
of the case. An unknown scheme id is an error.

## The command

The engine is kept as source in `third_party/`, with this repository's fixes;
`tools/vv/go.mod` builds it from there:

```sh
cd tools/vv
go build -o ~/bin/carneades github.com/carneades/carneades-4/src/cmd/carneades
carneades            # prints the engine's version, 4.3.1+walton.1
```

`carneades check case.yml` reads and validates a graph without evaluating it;
`carneades eval -t yaml case.yml` evaluates it and prints the graph with a
label for every statement (`in`, `out` or `undecided`). How a run can end, and
its exit status, is in `docs/contracts.md`, contract 3: 0 when complete, 1 for
invalid input, 3 for an inconsistent context, and so on. On any outcome other
than complete nothing is printed on stdout.

## A first case

`docs/examples/expert_opinion.yml` is the graph the tool makes for
`expert_opinion`, followed by this case:

```yaml
statements:
  liable(acme): Acme is liable.
assumptions:
  - expert(smith,law)
  - in_domain(liable(acme),law)
  - asserts(smith,liable(acme))
tests:
  in:
    - liable(acme)
```

The three assumptions are the premises of entry 2 (p. 310); the command builds
the argument from expert opinion and labels `liable(acme)` in. Its two
assumptions, `credible(smith)` and `based_on_evidence(asserts(smith,liable(acme)))`,
hold until they are questioned, and its two exceptions,
`untrustworthy(smith)` and `inconsistent_with_other_experts(liable(acme))`,
count only once they are shown.

## Writing a case

* **Statements** are terms: a predicate with arguments, `expert(smith,law)`.
  Constants begin with a lower-case letter, or are integers or quoted strings;
  names beginning with an upper-case letter are the schemes' variables. There
  is no arithmetic, and `is`, `true` and `false` are not predicates. Write one
  term per line in a YAML list: in a flow list (`[a, b]`) a comma inside a
  term splits it, so quote such a term (`["should_do(me,smoke)"]`).
* **Assumptions** (`assumptions:`) are what holds in the case: the premises of
  the schemes you mean to use, and any exception you can show.
* **Arguments** are built by the command from the schemes, one for each way
  the premises of a scheme hold. You can also write an argument out, under
  `arguments:`; one that names a scheme is checked against it (contract 6).
* **Critical questions** are of two kinds (D2, pp. 373-375). An assumption of
  a scheme holds until it is questioned: list it under `questioned:`, and the
  argument then needs support for it like an ordinary premise. An exception
  defeats the argument only once it is shown: assume it, or argue for it.
* **Negation** is written `¬`: `¬should_do(me,smoke)`. A statement and its
  negation are the two positions of one issue (the negation issue scheme), so
  arguments for both are weighed against each other. A graph that assumes
  both is an inconsistent context and is not evaluated.
* **Issues** with more than two positions, or with another proof standard, are
  written under `issues:` with their `positions:` and `standard:` (`PE`,
  preponderance of the evidence, the default; `CCE`, clear and convincing
  evidence; `BRD`, beyond reasonable doubt).
* **Expected labels** go under `tests:`, with the lists `in`, `out` and
  `undecided`. A statement named there has to be declared under
  `statements:`, with its text; otherwise the command refuses the file.

## Pro and con in one issue

`docs/examples/practical_reasoning.yml` puts practical inference (22.1) and
reasoning from negative consequences (33.3) in one graph. The first concludes
`should_do(me,take_job)`, the second its negation, so they meet in one issue.
With both arguments standing they weigh the same and neither position is in.
The example questions the assumption of 33.3 that the bad outcome is likely,
`outcome_likely(take_job,leave_family)`, which nothing supports: the argument
against falls, and `should_do(me,take_job)` is in.

The library shares a statement between schemes where the book means the same
thing by it (PLAN.md 5.6): `should_do/2` is the practical ought of entry 22,
of 6.3 and 6.4, 23, 25, 38 and 39; `has_property/2` the property of entries 9,
10, 21, 31.5 and 52; and "A is (plausibly) true" is the statement A itself. So
the conclusion of one scheme can be the premise of another, and arguments
from different schemes meet in one issue.

## Versions

The library's version is `meta.version` in `walton.yml` and the engine's is
printed by the command (`docs/releases/`). A graph written for release R1
(library 0.1.0) is brought up to date with

```sh
python3 tools/migrate_r2.py case.yml          # --check to see what would change
```

which renames identifiers from the table `tools/migration_r2.yml` and lists the
changes it cannot make without knowing the case (a statement with other
arguments, a premise added).

## Checking the examples

`python3 tools/test_graph_from_library.py` checks the tool against every
scheme; `go test -run TestCLI` in `tools/vv` evaluates the examples
(`docs/example-current`, `docs/example-labels`).
