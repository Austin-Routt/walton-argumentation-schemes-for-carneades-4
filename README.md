
# Walton Argumentation Schemes for Carneades 4

This repository contains an executable preview library of **Walton's Argumentation Schemes** in YAML for the [**Carneades Argumentation System** (version 4)](https://github.com/carneades/carneades-4).

The pinned original Carneades snapshot provides24 inherited records. The current `walton.yml` contains **143 records:119 book-origin and24 inherited**, with137 mapped to the109 indexed compendium entries and6 inherited records without an entry of their own. Several entries have separately executable alternatives. The source text prints no single compendium total; the repository's explicit entry index and current approval scope are in `tools/vv/source_ledger.yml` and the [current complete register](docs/reviews/closure-current-register.md).

In addition to our main file [walton.yml](walton.yml) in which a complete list is presented, there are also smaller files which are classified based on the classification system present in [**Argumentation Schemes (2008)**](https://www.cambridge.org/core/books/argumentation-schemes/9AE7E4E6ABDE690565442B2BD516A8B6). First, all the schemes are divided into 3 main categories ([**Reasoning**](A_reasoning), [**Source-based Arguments**](B_source-based-arguments), [**Applying Rules to Cases**](C_applying-rules-to-cases)), and under those categories, different files are created corresponding to different types of arguments falling into each category.

This study was conducted within the scope of the master's thesis titled **Douglas Walton's Argumentation Schemes and Applications** prepared by [**Başak Kurtuldu**](https://github.com/basakkurtuldu) under the supervision of [**Dr. Vedat Kamer**](https://istanbul.academia.edu/VedatKamer), in the [**Department of Philosophy**](http://felsefe.edebiyat.istanbul.edu.tr), [**Istanbul University**](https://www.istanbul.edu.tr). This fork of the [original repository](https://github.com/basakkurtuldu/walton-argumentation-schemes-for-carneades-4) corrects errors in the files; see [CHANGELOG.md](CHANGELOG.md) for what changed and [PLAN.md](PLAN.md) for what remains open.

## Files

* `walton.yml` – the complete list (143 schemes) and the negation issue scheme. **This is the only hand-edited scheme file.**
* `A_reasoning/`, `B_source-based-arguments/`, `C_applying-rules-to-cases/` – one file per type of argument. These files are **generated** from `walton.yml`: each contains, copied verbatim, the schemes assigned to it in `categories.yml` and exactly the language templates those schemes use, so each file loads into Carneades 4 on its own. Do not edit them; edit `walton.yml` (or the grouping in `categories.yml`) and run `python3 tools/build_category_files.py`.
* `categories.yml` – which scheme goes into which category file.
* `docs/using.md` – how to evaluate a case of your own with the library: Carneades 4 has no include mechanism, so a graph carries the schemes and templates it uses; `tools/graph_from_library.py SCHEME...` starts such a graph from `walton.yml`, and `docs/examples/` holds four worked cases that `tools/vv` evaluates on every push.
* `tools/validate_schemes.py` – checks the files against the rules Carneades 4 applies when it loads a theory (see below); `tools/test_validate_schemes.py` tests it with malformed fixtures and mutated copies of the library.
* `tools/build_category_files.py` – regenerates the category files; `--check` reports stale files without writing.
* `third_party/` – the sources of the Carneades 4 engine and of its rule engine GoCHR, imported byte for byte from `carneades-4.zip` and `GoCHR.zip` (see `sources.lock`). Engine fixes are kept here as focused, reviewable commits.
* `tools/vv/` – the verification suite: every scheme through the engine's Go API and through the `carneades` command, engine probes, reference cases written by hand (`tools/vv/cases/`), and a map of the book's compendium. CI runs it on every push.
* `PLAN.md` – the plan agreed by three independent reviews, with its implementation status; `docs/contracts.md` – what the engine and the library promise; `docs/parity.md` – the checks of the reviews' other harnesses against the suite.

## Checking the files

```
python3 tools/validate_schemes.py            # 0 errors expected; warnings are informational
python3 tools/build_category_files.py --check
```

Both need Python 3 and PyYAML. The validator follows the Carneades 4 source (`src/engine/caes/caes.go`, `src/engine/caes/encoding/yaml/yaml.go`, `src/engine/validation/validation.go`):

* every file must parse as YAML with an `argument_schemes` list, and a scheme may only use the keys Carneades reads (`id`, `meta`, `variables`, `premises`, `conclusions`, `assumptions`, `exceptions`, `deletions`, `guards`, `weight`, `roles`); other keys, such as a misspelt `assumption:`, are silently ignored by Carneades;
* every statement must be a term Carneades' parser accepts (`name`, `name(A,B)`, `¬name(A)` or a bare schema variable); a statement that does not parse is silently dropped when a scheme is instantiated;
* every variable used must be declared in `variables`, because `InstantiateScheme` binds parameters to the declared variables by position (never remove or reorder variables of a published scheme);
* every predicate needs a language template keyed `name/arity` (`name/0` for 0-ary predicates, which is how Carneades looks them up); the template is filled with Go's `fmt.Sprintf`, so the number of `%s` verbs must match the arity and each argument must be printed exactly once;
* the engine's rules for the supported profile are errors: every variable occurs in a premise; `is`, `true` and `false` are no predicate names; no assumption repeats a premise; a template has as many `%` signs as its predicate has arguments; deletions and guards are outside the profile; no YAML key occurs twice;
* the category files must agree exactly with `walton.yml`, every scheme must be in exactly one category file, and
* regression tests pin the direction of the schemes whose logic was once wrong, and one instance per corrected, repaired or added scheme, taken from the book and written out by hand, fails on a flipped negation or two swapped arguments (see `CHANGELOG.md`). A cycle of schemes that derive each other's premises is a warning.

A second, independent check runs the Carneades 4 engine itself on every scheme:

```
VV_STRICT=1 tools/vv/run.sh  # required repository, Python, harness and engine/library checks
```

`tools/vv/known_issues.yml` records known executable failures. The strict gate allows none. The independent review reproduced the previous 7,994 passing checks on both Go versions, then added cases that exposed and repaired previously untested defects. Use `tools/vv/expected_checks.txt` and the acceptance logs for the current named inventory. Passing executable tests does not automatically approve an interpretation of a source.

## Versions and releases

The corrected preview uses library **0.3.0** (`walton.yml` and all category files) and engine **4.3.1+walton.4** (`third_party/VERSION`). The old proposed R1 commit `8c092b6` is superseded and should not be tagged for publication. See [PLAN.md](PLAN.md), [the corrected release notes](docs/releases/r1-candidate.md) and [the engine contracts](docs/contracts.md).

The closure review has **128 of143 records at `example-passed`:119 book-origin and9 inherited**.13 inherited rows remain `transcription-checked` and2 `mapped`.35 of the original50 holds are closed and all23 new variants independently accepted. All book-origin rows being approved does not certify every inherited record assigned to a compendium entry: entries2,3,26,29,33.1 and33.2 still have no approved record. See the [current complete register](docs/reviews/closure-current-register.md) for every decision. The complete library remains a preview; full1.0.0 requires all143 rows to meet gateL.

Source approvals are recorded in `docs/reviews/independent-approvals.json` and bound to the reviewed content. Run:

```
python3 tools/check_source_approvals.py --repo .
python3 tools/check_source_approvals.py --repo . --require-stable
```

The first checks integrity of current approvals. The second is expected to fail while the 15 inherited-source holds remain; it is the complete-library stable gate. Fingerprint inventory does not create an approval. Both pinned Go versions (1.24.13 and 1.27.4) are required CI jobs.

## Closure verification and dialogue

The current named inventory contains **9,959 distinct registered checks**, plus separate native/unit/Python/vet/race gates. All declared API/CLI case expectations are checked without re-inference. Final acceptance requires clean copies on both pinned versions and an exact result comparison; the review archive records the exact candidate commit and all outcomes. Validator:0 errors and11 informational warnings, comprising six preserved unused templates and five dependency-cycle groups.

[Finite dialogue snapshots](docs/dialogue-snapshots.md) supplies ask, accepted-answer and reopen states for30 CQs across six families, plus explicit burden allocation for four ignorance families. It preserves independent support and original defeaters; appropriateness and answer adequacy are application judgments. It is separate from canonical static exceptions. API callers use `AddAssumption` to promote a default to a stipulation and maintain origin invariants when directly editing public fields.

Root engine ZIPs remain unchanged original provenance inputs; the current source snapshot and Git bundle carry the patches. Prepared assets and commands do not imply a remote release or upstream message.

## Known limitations and migration

* Source agreement and successful execution are separate requirements. Original cited books and the attributed 2019 correspondence remain outstanding where individually identified. The Atkinson paper is now reviewed: the inherited AS1 record has ten incomplete full-source question mappings and does not implement the paper's full AS2 transition/value model. See `docs/reviews/atkinson-2007.md`.
* A scheme default holds until challenged through `questioned:`; a graph-level stipulation remains assumed until withdrawn. Premises, assumptions, exceptions and rebuttals have distinct burdens. Evidence adequacy and application truth are supplied by the caller; the engine does not infer them from natural-language facts.
* This release supports the offline CLI/API profile. Logical work/storage quotas and cancellation are documented in `docs/contracts.md`; they do not provide a total process-memory or hostile-input service guarantee. API clients must structurally validate authored graphs before inference.
* Use `tools/migration_r2.yml` and `tools/migrate_r2.py` for previous preview graphs. Participant, event, diagram, coextension and evidence changes can require new facts; the tool reports manual work instead of inventing those facts. Keep the 2019 `exeption_to_rule` provenance hold separate from the compendium's `from_precedent` record.
* Shared predicates are reused only for equal meanings and roles. `has_property` is generic; typed `character_quality` is needed for character prediction. Personal practical statements use agent-scoped predicates. Nullary universal warrants are retained only with their declared interpretation; other implicit-context records remain qualified.
* Term syntax, supported names and outcomes are specified in `docs/contracts.md` and `docs/using.md`. Category files are generated; edit `walton.yml` and regenerate them.

#### Updates
##### 2019-10-08

According to our correspondence with **Dr. Douglas Walton** on September 19, 2019, in which he informed us about the changes he made on some of the argumentation schemes which were presented in [**Argumentation Schemes (2008)**](https://www.cambridge.org/core/books/argumentation-schemes/9AE7E4E6ABDE690565442B2BD516A8B6) , we implemented these changes in our file:

* _Argument From Precedent_ was changed to _Argument From Exception To A Rule_.
* Spelling error was corrected in _Abductive Argumentation Scheme_.
* Spelling error was corrected in _Argument From Evidence to Hypothesis_.

We wish to thank **Dr. Douglas Walton** for his comments and corrections on _Carneades YAML (CYML)_ files at this repository.
