# Independent engine and import-boundary closure review

**Baseline inspected:** `98a2335d2c9909cf597ae85a9f359f20d092ae6c` (PRs #18–#21 merged).  
**Comparison baseline:** `fcccbba53337321566caf4951eaa8addf1037c0b`.  
**Review completed:** 2026-11-22 UTC.  
**Scope:** current engine/importers and supported input boundary, before a broader corpus/ICCMA benchmarking effort. The root reviewer owns library/source approval, version locks, the overall report, and full acceptance on both Go versions.

## Conclusion

The repository's latest source work did not close the previously demonstrated importer defects. Between the earlier baseline and current master, production importer code was unchanged: the engine changes in those PRs were additional tests. The positive native-suite result therefore could coexist with the old AIF, malformed YAML, CAF and AGXML counterexamples. I reproduced that coexistence before changing code.

The local candidate repairs the concrete native defects and makes the supported boundary enforceable. It accepts strict native YAML and a deliberately limited AIF conversion policy. CAF, AGXML and LKIF require explicit CLI experimental opt-in and additional admission checks. Dung TGF accepts attack-only edges with declared endpoints. The captured Kialo exports now fail explicitly where incompatible, rather than yielding empty graphs, inverted signs, premise-free arguments or support-as-attack graphs.

This is a candidate for the root reviewer's final acceptance gate, not a claim of full AIF/AIFdb, CAF, LKIF, arg-microtexts or ICCMA conformance. No publication, tag, push or outside contact was performed. I authored these repairs; independent root review and the final full acceptance run remain separate evidence.

## Evidence reproduced before repairs

`before-import-results.json` contains 16 pinned observations using a fresh probe compiled against master before edits: the 12 exact legacy Kialo exports (six formats, two scopes), one minimal renamed native-YAML candidate, and three small AIF cases. Each records input SHA-256, compiler/runtime, counts, diagnostics and structural validation results.

* AIF's linked case assumed both premise and inferred conclusion. Its conflict case assumed both issue positions and omitted one Issue backlink. Unsupported L/YA nodes disappeared without an error.
* Renaming the Kialo candidate's unknown scheme to `linked` did not repair it: object-valued premises disappeared and an unaccepted premise produced an accepted conclusion. The capture's `direction`/`pro` attributes also had no native meaning.
* AGXML accepted EDU-only captures as empty, structurally valid graphs.
* CAF accepted the mismatched captures but replaced texts with IDs and negated assumptions, premises and conclusions. The capture's version and fields differ from native CAF; an additional native-default defect was independently found in the official CAF 1.3 schema.
* TGF treated every labeled edge as an attack, including support edges. LKIF already rejected the wrong root envelope. AIF already rejected object-valued scheme representations on the old parser.

The old issue register was correct about these reproducible implementation/profile problems. They were not refuted by the expanded positive suite.

## Repairs and decisions

### Native YAML

The importer now rejects duplicate keys and unknown semantic fields using strict decoding; validates list/map premise shape; refuses null argument, issue and scheme records; and rejects invalid or unknown statement shapes/labels. Lists contain statement-ID strings; role-bearing premises use a role-to-ID mapping. `direction: attack` and `pro: false` are not native ways to change a conclusion and are rejected. The native representation of opposition must be authored explicitly.

The previously misspelled internal `Weigth` field now has an explicit `yaml:"weight"` tag, so ordinary serialized weights are decoded. Both list premises and role maps have a positive regression, including a check that an unsupported premise cannot become unconditional acceptance.

Strict decoding exposed an additional defect in a bundled example: `examples/AGs/YAML/abduction.yml` used the ignored top-level key `labels`. Its apparent expected labels had never been assertions. The example now uses `tests`, declares the expected statement identities so CLI prevalidation can check them, and exercises its existing expected labels during native YAML/JSON roundtrips. This is a test-fixture correction, not a change to the Walton library's interpretation.

### Restricted AIF conversion

The JSON profile requires nodes and edges arrays, unique nonempty IDs, existing endpoints, and no duplicate relations. Only I, RA and CA nodes are admitted. RA has at least one I premise and exactly one I conclusion. A CA connects exactly two distinct I positions, one on each side; an I position cannot participate in several CA nodes. Other node types or incidence patterns are refused. Unknown JSON fields and duplicate JSON object keys are rejected, including duplicate nested keys and escaped-key aliases.

The conversion policy is explicit:

* RA becomes a generic linked argument. Its external scheme string is retained as `aif_scheme` provenance and **does not instantiate or validate a Walton executable scheme**.
* Binary CA becomes a symmetric incompatibility issue. This does not preserve arbitrary directed AIF conflict.
* I nodes that are neither RA conclusions nor CA issue positions become stipulations. This is an import policy, not a finding that the captured statements are true.

Both issue positions now have backlinks, and false assumption flags are respected. The resulting graph's metadata records the import/assumption policy. CLI help explains the restrictions and states that this is not general AIFdb conformance. The profile is structural/interpretive admission, not a full AIF serialization roundtrip or dialogical semantics adapter. API callers still must validate the resulting native graph before inference, as the release contract requires.

### Dung and TGF

`EqualArgSetSlices` now compares extension families in both directions. Missing or extra extensions, including an empty family versus a family containing the empty extension, are distinguished. Reordering and duplicate copies of the same extension do not matter because a semantic extension family is treated as a set.

The malformed odd-cycle fixture now uses three distinct arguments and three edges. It also checks that a directed odd cycle has no stable extension.

TGF now requires a separator, unique declared node IDs and exactly two endpoint IDs per edge. Dangling or malformed edges and all edge labels are refused. Duplicate attack edges collapse as set members. Optional node text has no Dung semantics. This prevents support labels from silently becoming attacks; it does not create a bipolar-to-Dung translation. The modern ICCMA interface and exponential enumeration limits are unchanged future capability/architecture work.

### Experimental legacy converters

`check` and `eval` refuse `-f caf`, `-f agxml` and `-f lkif` with exit 2 unless `-experimental-import` is supplied. The flag cannot bypass parser/profile checks, does not certify source fidelity, and does not make a conversion a benchmark reference. Native YAML remains the default. Typed AIF/profile refusals also keep exit 2; malformed native input keeps exit 1. Actual CLI subprocess tests cover these behaviors, including no evaluation output on refusal.

CAF admission is bounded to its 1.3-shaped profile. Unknown versions, unrecognized statement/argument fields, strict arguments without a reviewed mapping, invalid ranges/standards, duplicate IDs and dangling references are refused. The official [CAF 1.3 schema at commit 5d37c956](https://github.com/carneades/carneades-3/blob/5d37c956f18ad0c72b8d2838234ddd001f88adf6/schemas/CAF.rnc) specifies a default statement weight of 0.5, argument `pro=true`, and premise `positive=true`. Those defaults are now implemented. An absent evaluation value is treated as unknown by this converter; explicitly negative signs/zero weights retain their meaning. The shadowed pointer that could panic for a forward undercutter reference is repaired. External scheme names are provenance for generic linked conversion, not source-library bindings.

AGXML checks for actual ADUs, valid endpoints and recognized edge types. Its admitted segmentation is one EDU per ADU; joint segmentation is explicitly unsupported rather than silently misassembled. The existing false-assumption bug, duplicated conclusion backlinks, nil metadata on example support, and missing contrary text were repaired. These checks do not certify the full arg-microtexts schema or its entire annotation model.

LKIF admits one argument graph; a second graph is no longer silently discarded. IDs, endpoints and explicit premise polarity/type are checked. Exception arguments now receive the necessary conclusion backlink and target undercutter link. A behavioral control proves the exception defeats the target when supported and does not do so when unsupported. Generated exception IDs cannot overwrite declared IDs. The legacy distinction between facts/assumptions and native dialectical defaults is still not a certified semantic correspondence. LKIF stays experimental; no new corpus adapter or conformance claim was made.

## Final bounded results

| Observation | Before at master | Patched candidate |
|---|---|---|
| Exact mismatched Kialo exports | Several rejected; others silently changed or emptied graphs | 12/12 explicitly refused |
| Renamed minimal native candidate | Premise discarded; target accepted | Import error; no graph returned |
| Linked AIF assumptions | 2 | 1 (premise only) |
| Conflict AIF assumptions | 2 | 0 |
| Conflict missing Issue backlinks | 1 | 0 |
| Unsupported AIF L/YA nodes | Silently dropped | Typed unsupported refusal |
| Four previously authored conditional native fixtures | Previously passed in earlier package | 4/4 pass exact expected labels and argument instances |

The four conditional fixtures are engineering examples with explicit assumptions. They are not human semantic gold and do not validate Kialo conclusions or source truth. No large captured graph was inferred.

`after-import-results.json`, `conditional-results.json`, `import-summary.json` and `run_import_review.py` provide the corresponding machine evidence and reproducible bounded runner. The exact source hashes are checked before re-running the inputs.

## Focused verification completed

Using Go 1.27.4, all changed importer packages, Dung/TGF and the CLI passed focused tests, including actual CLI subprocess status checks. The bundled native YAML and JSON roundtrip tests passed, as did the repaired Dung example tests. See `focused-go1.27.4.log`, `native-examples-go1.27.4.log` and `dung-native-go1.27.4.log`. Duplicate JSON keys, unsupported nodes, malformed premise structures, CAF explicit-negative controls, AGXML rebuttal topology, and LKIF proved/unproved exception behavior are included.

The root reviewer is responsible for the final required full native engine, race and library suites on both supported Go versions, against the complete frozen candidate. These focused results must not be added to an older suite count or reported as a fresh complete framework benchmark. `changed-files.json` records the concrete edited engine files and hashes at handoff; root-owned version/lock updates are excluded from that ownership list.

## What remains outside this closure

There is no claim of formal proof, all-input safety, arbitrary-input resource bounds or complete source/corpus fidelity. The native parser/inference contract still requires API structural validation and handling typed failures. Full AIFdb/dialogical translation, general CAF/LKIF/microtexts conformance, Kialo annotation/adapters, modern ICCMA protocols, and large-AF solver scaling are separate validation-phase work. These are explicit scope boundaries and future capabilities, not hidden successful conversions.

The user can finalize a bounded framework baseline once independent integration review and the final acceptance gate succeed. It would be inaccurate to finalize a claim that every inherited converter or every proposed dataset already has a faithful, verified translation.

## Final cross-review addendum

The subsequent independent check of root-authored documentation exposed the JSON decoder's case-insensitive key matching: `Nodes` could overwrite `nodes` despite the initial exact duplicate-key check. The guard now rejects Unicode simple-fold-equivalent keys, with root/nested/Unicode controls. The original counterexample returned an empty graph; it now returns invalid input. AIF and actual CLI subprocess tests pass after this final repair. See `root-docs-cross-review.md`, the before/after case-alias probe logs, and `aif-cli-final-go1.27.4.log`. `changed-files.json` has been refreshed for these two files. This addendum precedes root's final candidate freeze.
