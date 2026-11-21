
# Walton Argumentation Schemes for Carneades 4

This repo contains full lists of **Walton's Argumentation Schemes** in YAML for the use for the [**Carneades Argumentation System** (version 4)](https://github.com/carneades/carneades-4).

24 of the **Walton's argumentation schemes** are already available in YAML at **Carneades 4** as [walton.yml](https://github.com/carneades/carneades-4/blob/master/examples/AGs/YAML/walton.yml). [**Argumentation Schemes (2008)**](https://www.cambridge.org/core/books/argumentation-schemes/9AE7E4E6ABDE690565442B2BD516A8B6) prints its compendium of schemes in chapter 9 without a total; counted one way it holds 106 schemes, counted another 109 (`tools/vv/source_ledger.yml` lists every entry with the record filed under it). We aimed to complete the list and created the remaining schemes in YAML: `walton.yml` holds 114 records, 110 filed under 108 compendium entries and 4 taken over from Carneades without an entry of their own. Three entries have no record of their own yet: 50, and 57 and 22.4, whose records hold other formulations (see `PLAN.md`, 3.8).

In addition to our main file [walton.yml](walton.yml) in which a complete list is presented, there are also smaller files which are classified based on the classification system present in [**Argumentation Schemes (2008)**](https://www.cambridge.org/core/books/argumentation-schemes/9AE7E4E6ABDE690565442B2BD516A8B6). First, all the schemes are divided into 3 main categories ([**Reasoning**](A_reasoning), [**Source-based Arguments**](B_source-based-arguments), [**Applying Rules to Cases**](C_applying-rules-to-cases)), and under those categories, different files are created corresponding to different types of arguments falling into each category.

This study was conducted within the scope of the master's thesis titled **Douglas Walton's Argumentation Schemes and Applications** prepared by [**Başak Kurtuldu**](https://github.com/basakkurtuldu) under the supervision of [**Dr. Vedat Kamer**](https://istanbul.academia.edu/VedatKamer), in the [**Department of Philosophy**](http://felsefe.edebiyat.istanbul.edu.tr), [**Istanbul University**](https://www.istanbul.edu.tr). This fork of the [original repository](https://github.com/basakkurtuldu/walton-argumentation-schemes-for-carneades-4) corrects errors in the files; see [CHANGELOG.md](CHANGELOG.md) for what changed and [PLAN.md](PLAN.md) for what remains open.

## Files

* `walton.yml` – the complete list (114 schemes). **This is the only hand-edited scheme file.**
* `A_reasoning/`, `B_source-based-arguments/`, `C_applying-rules-to-cases/` – one file per type of argument. These files are **generated** from `walton.yml`: each contains, copied verbatim, the schemes assigned to it in `categories.yml` and exactly the language templates those schemes use, so each file loads into Carneades 4 on its own. Do not edit them; edit `walton.yml` (or the grouping in `categories.yml`) and run `python3 tools/build_category_files.py`.
* `categories.yml` – which scheme goes into which category file.
* `tools/validate_schemes.py` – checks the files against the rules Carneades 4 applies when it loads a theory (see below).
* `tools/build_category_files.py` – regenerates the category files; `--check` reports stale files without writing.
* `third_party/` – the sources of the Carneades 4 engine and of its rule engine GoCHR, imported byte for byte from `carneades-4.zip` and `GoCHR.zip` (see `sources.lock`). Engine fixes are made here, one per commit.
* `tools/vv/` – the verification suite: every scheme through the engine's Go API and through the `carneades` command, engine probes, and a map of the book's compendium. CI runs it on every push.
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
* the category files must agree exactly with `walton.yml`, every scheme must be in exactly one category file, and
* regression tests pin the direction of the schemes whose logic was once wrong (see `CHANGELOG.md`).

A second, independent check runs the Carneades 4 engine itself on every scheme:

```
tools/vv/run.sh      # the two checks above, the engine sources against sources.lock, then tools/vv (Go 1.24, about 40 seconds)
```

`tools/vv/known_issues.yml` lists the checks that fail today and why: 293 of 6,407, under 9 issues (390 under 24 before the engine fixes of `PLAN.md` phases 1 and 2 began). A run is green when nothing fails that the ledger does not list and nothing listed has started to pass. See `tools/vv/README.md`.

## Known limitations

* **`carneades eval` evaluates nothing for a graph that contains `walton.yml`** today: the engine's validator rejects five schemes and `true/1`, and `eval` prints its problems on stderr and exits 1 (the engine as shipped exited 0). 54 of the 114 schemes can be evaluated from their category file. `PLAN.md` phases 1 to 3 fix this.
* `group_2` is never inferred (`is/2` is a GoCHR built-in). Fixed in `third_party/` (see the changelog): composition with division, and `exeptional`, no longer run to the rule limit; a limit is reported as an error; labelling returns after a second inference adds a position to an issue; negated premises (`ignorance`, `from_antimodel`, `opposition-descriptive_2/3`) match, and negations are worded with the language.
* A scheme's `assumptions` hold until they are called into question (pp. 373–375): list a statement under `questioned:` in a graph file (or `ArgGraph.Questioned`), and the default is suspended until an argument supports it. A statement the graph itself stipulates under `assumptions:` stays assumed (`docs/contracts.md`, contract 5).
* `roles` is documented by Carneades but not read by its importer.
* Rules for writing terms with this engine are at the end of `PLAN.md` (quote dates; no quoted strings with spaces; one term per line in YAML lists).
* Five schemes (`ad_hominem`, `practical_necessary_condition`, `practical_sufficient_condition`, `regulative_rule`, `sorites_slippery_slope`) have a variable that occurs only in a conclusion, assumption or exception, as in the book; Carneades' own validator reports these as problems because its inference engine needs range-restricted rules. The validator here reports them as warnings (`RANGE_RESTRICTION`).
* Misspelt identifiers are kept so that existing argument graphs keep working: scheme ids such as `exeption_to_rule`, `epistemic_ignorence`, `exeptional`, predicate names such as `recieves_evaluation`, `oficcial_rulebook`, `succesful_account`, the predicate `true/1` (which Carneades' validator rejects as a boolean), and the file name `5_casual-reasoning.yml`. See `PLAN.md`.
* Some schemes use 0-ary statements (for example `background`, `being_credible`) where the book's wording has no obvious variables.

#### Updates
##### 2019-10-08

According to our correspondence with **Dr. Douglas Walton** on September 19, 2019, in which he informed us about the changes he made on some of the argumentation schemes which were presented in [**Argumentation Schemes (2008)**](https://www.cambridge.org/core/books/argumentation-schemes/9AE7E4E6ABDE690565442B2BD516A8B6) , we implemented these changes in our file:

* _Argument From Precedent_ was changed to _Argument From Exception To A Rule_.
* Spelling error was corrected in _Abductive Argumentation Scheme_.
* Spelling error was corrected in _Argument From Evidence to Hypothesis_.

We wish to thank **Dr. Douglas Walton** for his comments and corrections on _Carneades YAML (CYML)_ files at this repository.
