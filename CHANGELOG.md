# Changelog

All page numbers refer to Walton, Reed & Macagno (2008), *Argumentation Schemes*, Cambridge University Press ("the book"); section numbers are those of the User's Compendium of Schemes (ch. 9).

## 2026-11-21 — a file format that round-trips (`PLAN.md` 4.1 to 4.3)

* The YAML export (and the JSON export, which goes through it) writes each argument's parameters one per line and quoted where YAML needs it: a flow list split `liable(acme,smith)` at its comma. Statement ids in premises, conclusions, undercutters and assumptions are quoted the same way.
* Expected labels are written under `tests:`, the key the importer reads (they were written under `labels:` and lost).
* The theory is written with each scheme's weight and roles and with the named weighing functions, as they were read: constant weights, criteria and preferences were written only as comments, so `rebuttal`, `mcda-porsche` and `lex` lost their labels on a round trip.
* Scheme `roles`, which USAGE.md documents, are imported and give each premise its role (E-ROLES); more roles than premises is an import error.
* Reference cases written by hand (`PLAN.md` 4.4, `tools/vv/cases/reference.yml`): what happens to the argument and to the conclusion for linked support, a missing premise, every state of a critical question, an undercutter, a rebuttal in an issue, three proof standards, two conclusions, a cycle, four of the ways a run ends, two cases side by side, a chain of schemes and an unrelated fact.
* Property tests (4.5): 400 generated statements mean the same to Carneades' reader and to GoCHR's; renaming every scheme's variables changes none of its arguments; every reference case is evaluated by the command, written as YAML and read back in a second process with the same labels.
* Found by these: an assumption no argument uses had no statement and no label (it is now in once a run completes); a written-out argument without parameters was refused although its premises and conclusion instantiate its scheme (it is now validated by one substitution, as contract 6 says); a scheme that only gives a weighing function made the validator refuse every argument naming it. Every example upstream ships now passes `carneades check`, but `chr-leq` (unsupported) and `paraconsistency` (an inconsistent context), as designed.
* The ledger lists the checks a complete run produces by name (`tools/vv/expected_checks.txt`, PLAN.md 0.3): a check that disappears or appears fails the run until the list is rewritten in the same change. An entry for a known failure can carry a fingerprint of its message, and a listed check that fails for another reason counts as a new failure.
* The ledger of known failures is empty: all 6,774 checks pass.

## 2026-11-21 — the scheme repairs of release R1 (`PLAN.md` Phase 3)

Each repair names its page, says how the passage is read, and has a positive and an adverse reference case in `tools/vv/cases/phase3.yml` (D13); `tools/validate_schemes.py` pins the instantiated statements of each one (3.10). The new predicate names follow the rule of D6 and are recorded with their pages in `tools/vv/source_ledger.yml`. All of these are **executable interpretations** of the book until Phase 5 has checked them. Identifiers change only where a scheme had to: none of the nine repaired schemes could be evaluated through the command before.

* **3.1 `ad_hominem`** (42, p. 336): adds the premise `advocates_conclusion(A,S,C)`, which the sibling schemes 43 and 45 open with, so the attacked argument is the person's own. Adverse: another person's argument, another argument of the same person.
* **3.2 `regulative_rule`** (55.3, p. 343): the book's entry pairs a prohibition with an obligation ("prohibited (obligatory)"); the record is the obligation reading (`obligatory/1` now says "%s is obligatory.") and `regulative_rule_prohibition` is added beside it (`prohibited/1`, `must_not_do/2`: a duty not to act, not the negation of a duty to act). Both add `addressee_of(A,X)`: "S must do A" does not follow from "A is obligatory" unless S is someone the rule applies to, which the book does not print. The excuse names its subject (`excuse_or_overriding_duty/2`, replacing `excuse/0`); the warrant `do_obligator` stays general.
* **3.3 `practical_necessary_condition`, `practical_sufficient_condition`** (22.2, 22.3, pp. 323-324): the goal premise names the agent (`goal_of_agent/2`, replacing `goal_is_action/1`); the last question's "other goals" are not named, so `conflicting_goals/2` replaces `conflict_goals/3` and the parameter `A2` goes. New names, because a fact left at the old arity would be another statement.
* **3.4 `sorites_slippery_slope`** (51, pp. 340-341): the sequence premise names its two ends and its property, `reapplication_sequence(A2,A,P)`, so that a sequence for another property or between other individuals does not serve, and `A` is bound.
* **3.5 `pop_scheme`** (4.1, p. 311): concludes the statement `A` itself instead of `true(A)` (`true/1` is a boolean to the engine; L-TRUE). The second conclusion, `should_accept(A)`, stays until R2 splits the two readings.
* **3.6 `group_2`** (21.2, p. 322): concludes `has(M,Q)`, the statement variant 1 uses for the group; `is/2`, which the rule engine executes as an assignment, is renamed `is_a/2` in `rule_physical_world_1` (54.1, p. 342), where it is a premise.
* **3.7 `exeptional`** (56, p. 344): the major premise is stated for the case at hand (`rule_can_waived(A)`, the variable `X` goes), and the assumption that repeated the minor premise goes: it made the rule feed itself, and a stipulated premise can never be questioned. The scheme's three conditional questions are left to Phase 5.
* **3.8 three entries get records of their own**: `precedent_slippery_slope` (50, p. 340; its three questions restate the premises and are not modelled, p. 374); `from_precedent` (57, p. 344: the minor premise "a has F but does not have G" is `has_property(A,F)` with `¬has_property(A,G)`; its legitimacy is an assumption, CQ2; CQ1 and CQ3 are exceptions); `value_based_practical_reasoning` (22.4, p. 324, the three-premise form: CQ6 is an assumption, CQ1, CQ3, CQ4, CQ5 and CQ7 are exceptions, CQ2 restates premise 2). The inherited `exeption_to_rule` and `practical_reasoning1`, filed under 57 and 22.4 before, keep their ids and are listed as taken over from Carneades.
* **3.9 the negation issue scheme**: `issue_schemes: negation: ["S", "¬S"]` in `walton.yml`, copied into every category file by `tools/build_category_files.py`, so arguments for a statement and for its negation are weighed in one issue (D10).
* **3.10 `tools/validate_schemes.py`**: the engine's rules are errors (range restriction, `is`/`true`/`false`, an assumption repeating a premise, the engine's `%` count, deletions and guards, duplicate YAML keys); a conclusion repeating a premise and cycles of schemes are warnings; the 17 gaps of the third review's probes are closed, and its probes and mutation tests are ported as `tools/test_validate_schemes.py`, which CI runs; one instance per corrected, repaired or added scheme is written out by hand.
* **3.11** two citations that differ from the compendium are recorded in the source ledger, not changed: `sorites_slippery_slope` (p. 151; the compendium has p. 251) and `exeptional` (p. 148; the compendium has p. 147).
* Engine, with the repairs: the validator rejects `is`, `true` and `false` as predicate names (contract 7; X-IS); `Infer` reports a graph that assumes a statement and its negation, or two positions of one issue, as an inconsistent context (contract 3), as `check` and `eval` already did; upstream's example `paraconsistency`, which does that on purpose, is expected to be refused.
* Gate S of the plan: `carneades check` reports nothing for any of the 14 files, every one of the 118 schemes is evaluated through the command from `walton.yml` and from its category file, and the ledger holds only E-EXPORT and E-ROLES. Gate: 6,692 checks, 3 known failures under those 2 issues.

## 2026-11-21 — engine fixes, first part (`PLAN.md` phases 1 and 2)

No scheme file changed. Each fix is one commit in `third_party/`, with tests in the engine and the suite's ledger updated in the same commit; CI also runs both engines' own unit tests now (`tools/vv/engine-tests.sh`).

* Exit status by outcome: `check` and `eval` exit 1 for invalid input, 3 for an inconsistent context, 4 at a limit, 6 for an internal failure (a panic); `eval` prints no evaluation unless the run completed (E-EXIT, contract 3).
* Issues: a statement that joins an issue becomes one of its positions, and labelling returns after a second inference (E-ISSUE).
* The importer keeps no state between calls, and the validator checks an argument against its scheme with one substitution (E-GLOBALS, E-ARGVAL). Export writes the theory (first half of `PLAN.md` 4.1).
* `Inconsistent()` answers from the current assumptions; `check` and `eval` report a graph that assumes a statement and its negation (E-INCONS).
* Argument construction treats facts as a set: composition with division, and `exeptional`, terminate instead of running to the rule limit and starving every later scheme (X-LOOP; GoCHR `RuleStore.Unique`).
* Limits: reaching the rule limit, or deriving a term over 500 nodes, is an error (`caes.ErrLimitReached`, exit 4) that leaves the graph unchanged; nothing is written to stdout; `InferContext` stops a run from outside (E-BUDGET).
* Unreadable rule text is an error: GoCHR refuses a premise, conclusion or goal it cannot read in full (it read `broken(` as `broken` and `p(X) q(Y)` as `p(X)`), and `Infer` reports the scheme as invalid input instead of skipping it. When no rule fires, GoCHR returns the goals instead of an empty store (E-PARSE).
* Build health: both engine modules build, vet and test as shipped, and CI runs all three (`tools/vv/engine-tests.sh`). The GopherJS front end compiles again; `dung/encoding/yaml`, a 2015 copy of the YAML codec that nothing imported and that no longer compiled, is removed (the archive keeps it); vet findings are fixed (among them GoCHR's parse-error messages, which were garbled, and the export of a criteria's hard constraints, written as `%!s(int=0)`); the CouchDB test skips itself without a server instead of ending the test binary; the YAML and JSON tests assert their labels, with the three examples whose weights export loses listed under E-EXPORT (E-BUILD).
* Matching: argument construction runs on a new solver in GoCHR for programs with set semantics and propagation rules only (`propagate.go`), with a real propagation history: each scheme applies once per tuple of facts, a premise that is a bare variable matches wherever it stands, and a fact with a premise's functor at another arity is harmless. The general solver tracked what it had tried by positions in per-functor lists, which both cases upset, so a rule fired on the same facts until the limit or never (E-MATCH).
* Negation: `¬S` is matched by rule heads like any statement, also when it reaches a rule through a statement variable, and `¬¬S` is read as `S` (GoCHR `ExplicitNegation`, which Carneades sets; without it GoCHR filed negated statements with its built-ins, where no rule looks). The four schemes with a negated premise (`ignorance`, `from_antimodel`, `opposition-descriptive_2`, `opposition-descriptive_3`) now fire. A user predicate called `neg` stays ordinary. Negated statements are worded with the language ("It is not the case that …") instead of printed as terms (X-NEG, E-NEGTEXT). The suite's `static/no-negated-premise` is removed, as the plan said (228 results fewer).
* Repeated inference: a second `Infer` on a graph adds nothing. Arguments are identified by scheme and canonical parameters (the engine compared `[a,b]` with `[a, b]` and built every argument with two or more parameters again, assuming its assumptions again), and the undercutters built for exceptions, which have no scheme, no longer make it panic (E-IDEMP). The suite runs a second inference for every scheme in every file (`iso/infer-twice`, `full/infer-twice`, ported from the third review).
* Names and terms: a scheme's parameters are substituted once, so a scheme variable called `A1` no longer captures the fresh variable the rule engine gives an unbound `A` (E-CAPTURE). A quoted string keeps its spaces (`"it rains"` was read as `"it`). A statement with arithmetic the rule engine would compute, such as an unquoted `1990-05-17`, is invalid input instead of the number 1968; a hyphen between names, as in `opposition-normative`, and a negative number stay constants. A scheme with CHR deletions or guards is outside the supported profile and is refused as unsupported (exit 2), as the plan decided (D3); upstream's example `chr-leq` is such a scheme, and its tests now expect that (E-TERMS, contract 1).
* Critical questions under a scheme's `assumptions:` can be called into question: a graph lists them under `questioned:`, and a questioned default is no longer assumed, so the argument needs support for it like an ordinary premise; a statement the graph stipulates stays assumed. The engine keeps the defaults a scheme instantiation adds apart from the stipulations, and export writes `questioned:`, `defaults:` and `assumptions:` apart (X-ASSUME, D2, contract 5). The suite questions every assumption of every scheme (`assumptions-questionable`, ported from the third review): `exeptional` fails it because its assumption repeats a premise (`PLAN.md` 3.7), and the practical pair because `P` is unbound (3.3).
* Arguments that name a scheme: a written-out argument whose scheme has exceptions and which has no undercutter is reported by the validator, not completed (a written-out argument is never rewritten); an argument with only `scheme:` and `parameters:` is a request to instantiate the scheme and is carried out on import, with the scheme's premises, conclusion and questions (E-EXPLICIT, D12, contract 6).
* Gate: 6,407 checks; 291 known failures under 8 issues (390 of 5,947 under 24 before).

## 2026-11-21 — the reviewed plan, the engine sources and the verification gate

No scheme file changed.

* `PLAN.md` is replaced by the plan agreed by three independent reviews (`docs/reviews.md`), with the owner's instructions applied: the engine is kept as source in this repository and names follow a rule.
* `third_party/carneades-4` and `third_party/GoCHR`: the engine sources, imported byte for byte from `carneades-4.zip` and `GoCHR.zip` (`sources.lock`). Engine fixes are made here, one per commit.
* `tools/vv`: the verification suite, 5,947 checks; `tools/vv/known_issues.yml` lists the 390 that fail today. CI runs `tools/vv/run.sh` on every push (`.github/workflows/verify.yml`).
* `docs/contracts.md` (what the engine and the library promise), `docs/parity.md` (the reviews' other checks against the suite).
* README: counts corrected (upstream ships 24 records, not 25; the compendium is 106 or 109 entries depending on the count; 114 records), known limitations updated.
* Correction to the entry of 2026-11-06: `exeptional` (56, p. 344) was not "checked and left unchanged" correctly: its major premise has a variable no premise binds and its assumption repeats the minor premise, which makes the rule feed itself (`PLAN.md` 3.7). And the alternatives premise of `practical_sufficient_condition` still says "necessary", as the compendium prints it (22.3, p. 323); the narrative on p. 96 says "sufficient" (`PLAN.md` 5.3).

## 2026-11-06 — branch `fix/validate-and-correct-schemes`

Baseline: `master` at a7cde94. Checked by loading every file with the Carneades 4 engine (master d19d543): **6 of the 14 scheme files did not load at all** (`walton.yml`, `A_reasoning/2`, `A_reasoning/5`, `B_source-based-arguments/2`, `C_applying-rules-to-cases/2`, `C_applying-rules-to-cases/4`), and in the 8 that loaded the engine's own validator reported 103 scheme problems and 27 language problems; many statements rendered as empty or garbled text. After the changes all 14 files load, `tools/validate_schemes.py` reports 0 errors, and the engine's validator reports only the five inherited range-restriction cases and `true/1` listed under Known limitations in the README.

The changes are ordered from "the file cannot be loaded" to "cosmetic"; each tier is a separate commit, so the later tiers can be dropped if they are not wanted.

### 1. YAML syntax — files could not be loaded (commit "Fix YAML syntax so every file loads")

16 one-line edits: `- id:x` without a space (2), titles containing `: ` that must be quoted (4), list items written `-smaller_part_property(...)` without a space, which YAML reads as a string instead of a list (8), and `assumption:` for `assumptions:` (2), which Carneades ignores silently so the assumptions were lost.

### 2. Statement syntax — statements Carneades' term parser rejects (commit "Make every statement parse and declare its variables")

Carneades drops, with only a message on stderr, every premise, conclusion, assumption or exception it cannot parse, and leaves variables it cannot bind in the statement. Fixed: primed variables `P'`, `W'`, `X'`, `R'`, `Y'` (21, renamed `P2` etc. in the opposition schemes), an unbalanced parenthesis in `needed_nonarbitrary((A,F)`, and three schemes using a variable they did not declare (`part_in_whole`, `verbal_classification_vagueness`, `verbal_slippery_slope`).

### 3. Scheme logic — schemes that do not say what the book says (commit "Correct scheme logic against the book's compendium")

* `division_generic` (10.1, p. 317) was a copy of composition and reasoned from the parts to the whole; now whole → parts, with divisional heredity (CQ1) as the assumption.
* `composition_generic` (9.1, p. 316): the assumption restated the premise instead of CQ1 (compositional heredity).
* `from_example` (6.1, p. 314) and `common_folk_ad_populum` (4.10, p. 313): premise and conclusion were swapped.
* `moral_justification_ad_populum_excuse` (4.6, p. 312) had no conclusion; it was listed as a premise.
* `evidence_hypothesis_falsification` (32.2, pp. 331–332) concluded `¬A`, the negation of its own premise, instead of `¬H`.
* `danger` (37, p. 334) concluded that the danger, not the action, should not be brought about.
* `treat` (35, p. 333): premise 2 is "I am in a position to bring about B", not "position to know"; the title "Argument From Treat" is "Argument From Threat".
* `sacrifice` (20, p. 322): `great_sacrifice(S,A,X)` used an undeclared `X` where the value `V` is meant.
* `abductive_backward` (31.1, pp. 329–330): the premise "A1 is the account that explains D most successfully" was missing; `data_set` had a spurious second argument.
* `opposition-normative` (11.2, p. 318): the major premise is about evaluations, and the minor premise must instantiate its antecedent (`W`). The book prints `W'` in the minor premise; with `W'` the scheme would conclude its own premise false, so this is treated as a misprint.
* `rhetorical_oppositions_normative` (12.1, p. 318): the arguments of `if_treatment_required` were in the wrong order.
* `negative_consequences_prudential`: `strongly_occur` names both the action and the consequence.
* `commitment1`: the conclusion predicate was misspelt (`commined_to`), so it had no template.
* Titles: `opposition-descriptive_3` was labelled "Descriptive Scheme 2"; `verbal_classification_vagueness` carried the definition scheme's title.
* Three schemes existed only in category files and were added to `walton.yml`: `whole_into_parts` (10.2), `negative_consequences_reasoning` (33.3; it had the duplicate id `negative_consequences`), `full_slippery_slope` (53; source added: Walton, *Slippery Slope Arguments*, 1992, pp. 200, 203).

Checked against the book and left unchanged: `fear` (premise 3 is in the book), `bias_ad_hominem`, `full_slippery_slope`'s assumption (= CQ3), `exeptional` (56), and the critical-question coverage of the 24 schemes taken over from Carneades' own `walton.yml`.

### 4. Language templates — what Carneades prints (commit "Make every language template usable and correct its text")

Carneades looks a template up by `name/arity` and fills it with Go's `Sprintf`.

* 29 keys had no arity (`background`, `strongly_occur`, …) and were never found, so the statement text was empty; 0-ary predicates now use `name/0`, the others `name/n`.
* Arity mismatches between key and statements: `cited_bad_consequences`, `conflict_goals`, `position_to_know_group`, `same_meaning`; templates with more `%s` than arguments (`favor/1`, `popular_practise/1`, `other_consequences_presumption/2`) printed `%!s(MISSING)`.
* Malformed verbs: `"% is %s."` (`is/2`), `"%is not required."`, `"% more effectively"`.
* Arguments printed in the wrong roles (`assert_bad`, `help_needed`, `commitment_arises`, `if_treatment_required`, `treatment_required`, `contrary_treatment_required`, `opposite_action`), e.g. "the treatment ⟨subject⟩ is required for the subject ⟨treatment⟩".
* Duplicate keys `biased/1` and `good/1` (YAML keeps the last definition); the first was removed.
* Templates added for the corrected and added schemes; `position_to_know/1`, whose text was "I am in position to bring about %s", renamed `position_to_bring_about/1`.
* Spelling, grammar and stray spaces in about 80 texts ("wil l", "commiued", "Oflicial rulelebooks", " 's", …). This part is cosmetic but was done in the same commit.
* Follow-up in the final commit: ten templates that printed an argument twice via `%[n]s` were rewritten to print each argument once, because Carneades' validator counts `%` characters.

### 5. Titles and sources (commit "Fix typos in titles and sources") — cosmetic

OCR-style errors in titles and references ("lnformal Logic", "Trtatise", "Wayrw Grennan", "Physical Worl", "Ignorence", "Ad Baculum Treat", "Argumentationmustern"), the missing title of `rule_mental_world` and the missing source of `retroductive` (31.5, p. 331).

### 6. Generated category files and tooling

* `tools/validate_schemes.py` (first commit on the branch) and `tools/build_category_files.py` with `categories.yml`.
* The 13 category files are regenerated from `walton.yml`. Before, they were hand-maintained copies that had drifted: `B_source-based-arguments/1` had no language section at all, seven rule schemes appeared in both `C_…/2` and `C_…/4`, `opposition_descriptive_2/3` used different ids from `walton.yml`, and seven schemes of `walton.yml` were in no category file. Membership is otherwise unchanged; `A_reasoning/3`'s second `practical_necessary_condition` is `practical_sufficient_condition` and its second `negative_consequences` is `negative_consequences_reasoning`.

### Not changed

Misspelt identifiers (scheme ids, predicate names, the file name `5_casual-reasoning.yml`) are kept because renaming them changes the interface for existing argument graphs; see `PLAN.md`.
