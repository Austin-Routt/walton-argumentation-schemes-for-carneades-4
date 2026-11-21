# Verification and validation suite

This directory checks two things together: the Carneades 4 engine whose sources
are kept in `third_party/` at the repository root (imported byte for byte from
`carneades-4.zip` and `GoCHR.zip`, see `sources.lock`) and the Walton scheme
library (`walton.yml` and the category files). Every check drives that engine;
nothing here reimplements it.

## Running it

```
tools/vv/run.sh          # Python checks, the source check, then the suite (about 40 seconds)
```

or step by step:

```
tools/vv/engine-tests.sh # the engines' own unit tests on third_party/ (a few seconds)
tools/vv/setup.sh        # check the archives against sources.lock; list engine files changed here
cd tools/vv
go test                  # about 35 seconds
```

CI runs `tools/vv/engine-tests.sh`, then `tools/vv/run.sh`, on every push and pull request
(`.github/workflows/verify.yml`).

Needs Go, `unzip`, Python 3 with PyYAML, and network access to the Go module
proxy the first time (for `gopkg.in/yaml.v2` and `go-yaml2json`, the engine's
own dependencies). The suite
was written and run with Go 1.24.7; its `go.mod`, like the engine's, asks for
1.18 or newer, and older versions than 1.24 have not been tried. Graphviz `dot`
is used by one check if it is installed.

The run ends with a summary by issue. A green run means: nothing fails that the
ledger does not already list, and nothing the ledger lists has started passing.

Run `go test` without a package argument, as shown. `go test .` runs the same
checks but prints only `ok` when the run is green, so the summary is lost.

| command | what it does |
|---|---|
| `go test` | everything |
| `go test -run TestLibrary` | the scheme files through the engine's Go API |
| `go test -run TestCLI` | the scheme files through the `carneades` command, one new process per run |
| `go test -run 'TestEngine$'` | engine behaviour, independent of the scheme files |
| `go test -run TestSource` | `walton.yml` against the map of the book's compendium (`source_ledger.yml`) |
| `go test -run TestEngineToolchain` | build, vet and test the Carneades module as shipped |
| `VV_STRICT=1 go test` | known failures count as failures; green only when the ledger is empty |
| `go run ./cmd/schemecheck` | plain-text report: every failing check, and the schemes that can re-trigger each other |
| `go run ./cmd/schemecheck -format md` | one table row per scheme |
| `go run ./cmd/schemecheck -format json` | every result, one JSON object per line |
| `go run ./cmd/schemecheck -format source` | CSV: each scheme next to its entry in the book's compendium, with the questions the book prints and the ones the scheme models |

Environment variables: `WALTON_REPO` (repository root, found automatically when
the suite sits in `tools/vv`), `CARNEADES_BIN` (use this binary instead of
building one), `VV_RESULTS` (write all results to this file as JSON lines),
`VV_SKIP_TOOLCHAIN` (skip `TestEngineToolchain`), `VV_GOMOD_EXTRA` (file with
extra `go.mod` directives for building the engine where module downloads are
restricted).

## The ledger

`known_issues.yml` lists the checks that fail today and the issue each belongs
to. It makes the suite a ratchet:

* a listed check that fails is reported as known, and the run stays green;
* a check that fails without being listed fails the run: a regression, or a
  defect nobody had seen;
* a listed check that passes fails the run until its entry is removed.

When you fix something, the run tells you which entries to delete. An entry
that matches no check that ran fails the run too, and so does a run with fewer
checks than `min_checks` in the ledger says, because a passing check leaves no
entry and could otherwise be removed without anyone noticing.

Two limits are worth knowing. Nothing stops a later change from adding entries
or lowering `min_checks`, so a change to the ledger needs the same review as a
change to the code. And an entry names a check, not a cause: a check that starts
failing for a different reason still counts as known. `VV_STRICT=1` is the
release gate; the ordinary run only says that nothing got worse.

## What is checked

**Per scheme, per file** (`TestLibrary`). Each scheme gets a ground instance:
every declared variable gets its own constant, and a variable that stands for a
whole statement gets a small proposition. The instance's premises are asserted
and the engine constructs and evaluates the argument, once with the scheme alone
(`iso/…`) and once with every scheme of the file loaded (`full/…`).

| check | passes when |
|---|---|
| `engine-validator` | the engine's own validator reports nothing for the scheme |
| `chr-translation` | GoCHR accepts the rule the scheme is translated into |
| `static/id-round-trips` | the scheme id comes back from the engine's term reader as it was written, and has no number or variable in it that the rule engine would compute or bind |
| `static/range-restricted` | every variable used after the premises occurs in a premise |
| `static/no-reserved-functor` | no conclusion or assumption uses `is/2`, `=/2`, `:=/2` |
| `static/not-self-feeding` | the scheme does not derive one of its own premises |
| `…/assumptions-questionable` | each assumption of the scheme, called into question and unsupported, gives the scheme's argument the weight 0 |
| `…/infer-twice` | a second `Infer` on the finished graph adds no argument and changes no label |
| `…/infer-ok`, `…/terminates` | argument construction returns without error and without hitting the rule limit |
| `…/fires` | one argument per conclusion is constructed, with the expected parameters |
| `…/ground` | no statement in the resulting graph contains a variable |
| `…/conclusions-in` | every conclusion is labelled in |
| `iso/premises` | the argument's premises are the scheme's premises followed by its assumptions, and the assumptions are in |
| `…/exceptions-undercut` | asserting any one exception undercuts the scheme's own argument |
| `…/premises-needed` | leaving out any one premise stops the scheme: it constructs no argument |
| `iso/texts`, `iso/negation-texts` | every statement has a text produced by a language template |

`exceptions-undercut` and `premises-needed` judge the scheme's own argument, not
the label of its conclusion. With other schemes loaded a conclusion can keep the
support of another argument, which is how convergent arguments are meant to
behave. Both need a control: the plain run has to finish and build the scheme's
argument, and their own runs have to finish. Where that is not so they fail with
"not judged", because a graph the engine left half-built would otherwise count
as a successful defeat.

**Per scheme, through the command** (`TestCLI`): the scheme file plus the
instance's premises as assumptions is given to `carneades eval -t yaml` in a new
process, and the conclusions must come out labelled in (`cli/scenario`). Per
file: `carneades check` reports nothing and `carneades eval` prints something.

**Against the book** (`TestSource`). `source_ledger.yml` maps the schemes of
the compendium in Walton, Reed and Macagno (2008, ch. 9) to the records of
`walton.yml`: 106 by the book's own numbers, 109 entries here because two
numbers hold several schemas, with the page of each entry and the number of critical questions
it prints. The checks are structural: every scheme has exactly one place in the
map (`source/mapped`), and every entry has a record (`source/has-record`). The
file holds numbers and pages only, none of the book's text.

**Engine** (`TestEngine`, `TestEngineToolchain`, and the `command` group of
`TestCLI`): small experiments stating what practical use needs, such as "two
schemes that derive each other's premise terminate", "the engine can read the
YAML it writes", "an argument written by hand that names a scheme is subject to
the scheme's exceptions". Checks whose name starts with `baseline-` state what
works today and has to keep working.

**Containment.** The engine cannot be interrupted from outside. Four probes
that can run away (a counter that never stops, a term that keeps growing, a
fact of the wrong arity, labelling after an issue gains a position) therefore
run in a process of their own, with a time
limit and a memory cap of 2 GB, and are killed when they exceed either. If any
other check times out, the run stops there and fails: a worker that could not
be stopped is still using the engine's global state, and nothing measured after
it could be trusted.

## Adding a check

Add it to `internal/vvlib` so that it returns a `Result`. State the behaviour
that is wanted, not the behaviour that exists; if it fails today, add an entry
to the ledger under an issue.

## Limits

* The suite tests that a scheme is usable by the engine and behaves as its own
  text says. Against the book it checks only that every scheme has a place in
  the compendium, that the record filed there is meant to be that entry, and
  how many assumption and exception statements it has; whether a
  scheme's wording and its questions say what the book says has to be judged by
  a person with the book.
* One generic instance per scheme. Interactions that need particular data
  (cyclic `causes` facts, for example) are reported by `schemecheck` as cycles
  of the feeds graph, not exercised.
* The engine keeps global state and writes to stdout and stderr. The suite
  resets and captures both, and never runs checks in parallel.
* A green strict run says that these finite cases pass. It is not a proof for
  every input, and no check here can tell whether a scheme says what the book
  says: a flipped negation or two swapped arguments in a scheme pass every
  check. That needs examples written from the book, with the expected result
  written by hand (plan items 3.10, 4.4 and 5.7).
* Checks that state one of two outcomes the project could choose, to make an
  input work or to reject it with a message, follow docs/contracts.md:
  `premise-order-does-not-matter` passes (the input works);
  `hyphenated-constants-are-not-computed` was restated as
  `arithmetic-in-a-statement-is-invalid-input` (contract 1);
  `explicit-argument-gets-critical-questions` was restated as
  `explicit-argument-without-questions-is-reported` (contract 6);
  `argument-from-scheme-and-parameters` passes (the request is carried out);
  `is-as-predicate-in-conclusion` was restated as
  `reserved-predicate-names-are-rejected` (contract 7);
  `scheme-roles-are-imported` still states the stronger outcome. Where the decision is to reject, the check is
  restated in a change of its own, after the contract that says so.
* A ledger entry names a check, not a cause. Nothing here notices when a known
  failure starts failing for another reason; a fingerprint of each known
  failure's message would, and is not built yet.
* Not ported yet from the two other reviews: the second review's six acceptance
  cases; the third review's per-scheme checks of a withdrawn assumption and of
  a second inference (this suite has both only at engine level), its Go
  contracts for concurrency and for the store GoCHR returns when no rule fired,
  its 20 probes of `tools/validate_schemes.py` and its two mutation tests.
* `experiments/` holds four experiments, none applied: a prototype that gives
  the rule engine set semantics, trial edits to the scheme files on a copy of
  the repository, the two combined with the engine patch of a second review,
  and a prototype of the `questioned:` list for critical questions. See
  `experiments/README.md`.

## History

* 2026-11-20: first version, 5,176 checks, 349 known failures under 20 issues.
* 2026-11-21: compared with two independent reviews and with the book. Added
  the checks only those reviews had (a critical question under `assumptions:`,
  a second `Infer` with an exception, compound parameters and expected labels
  in the export, issue membership, unreadable statements, premises and
  exceptions with the whole file loaded) and the source ledger: 5,935 checks,
  370 known failures under 25 issues.
* 2026-11-21, second round: the two other reviews examined this suite. Taken
  from them: probes that can run away now run in a process of their own, and a
  timeout anywhere else stops the run; `setup.sh` refuses an archive whose
  checksum differs; exceptions and premises are judged by the scheme's own
  argument, which withdrew issue L-OVERLAP; new checks for a term that keeps
  growing, a fact of the wrong arity, premise order (issue E-MATCH), shared
  bindings in argument validation, `Inconsistent()` before labelling and after
  a withdrawal, and the `questioned:` contract for critical questions.
  5,944 checks, 368 known failures under 25 issues.
* 2026-11-21, third round: both other reviews examined the second version of
  this suite. Taken from them: a check that adds an exception or leaves a
  premise out needs a finished control run and a finished run of its own (26
  checks that used to pass on a half-built graph now fail as "not judged");
  entries 57 and 22.4 of the compendium count as having no record, because the
  records filed there hold other schemes; the id check asks whether an id
  comes back unchanged, which withdrew issue L-ID; the count of assumption and
  exception statements is no longer called "questions modelled". New checks:
  a negation through a statement variable, a user predicate called `neg`, and
  labelling after an issue gains a position. A run with fewer checks than
  `min_checks` fails. 5,947 checks, 390 known failures under 24 issues.

## Reference cases

`tools/vv/cases/*.yml` hold graphs written by hand against the library, each
with the arguments and labels it must produce (`internal/vvlib/cases.go`
documents the format). They are the positive and adverse examples the plan asks
for with every scheme repair (D13), and grow into the reference cases of
`PLAN.md` 4.4 and the book instances of 5.7. Each case is one check, `case`,
under its file. Write YAML lists of terms one per line: a flow list splits a
term at its comma.
