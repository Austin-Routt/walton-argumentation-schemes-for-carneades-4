# Integration of the post-integration readiness review

2026-11-22. Integrator: Claude, the implementer, whose integration (PR22)
the review examined. This report records what was checked before PR23 merged
the review's three commits unchanged, two precision notes on its attribution
correction, and the validation pin.

## What was received

`Walton_Framework_Finalization_Review_2026-11-21` (second delivery,
131,582,730 bytes, SHA-256
`c872a78ebb6a8a1719ed7eedd0f734ed8d1b6dccd01336ea0dc5e8014a375de8`). All 3,847
manifest entries match the extracted files, with no extra file. Its bundle is
a complete history whose head, `0a1e406`, descends from PR22 `d9a2585`
through `42b6c60`, `d71456f` and `0a1e406`; the three patches applied to
`d9a2585` reproduce tree `b6cd848`, and the packaged `framework/` directory
is that tree.

## Findings about PR22, reproduced

The two HTTP defects were reproduced on PR22 with fixtures written for this
check (vocabulary different from the review's), against a running server and
the CLI built from the same tree:

| Input | PR22 form route | PR22 raw route | CLI | `0a1e406`, both routes |
|---|---|---|---|---|
| `rain` and `¬rain` both assumed | 200, both labelled `in` | 400 | exit 3 | 422 |
| scheme with an unsupported guard | 200, graph without the conclusion | 400 | exit 2 | 422 |
| argument whose premise does not instantiate its scheme | 200 (error page) | 200, conclusion labelled `in` | exit 1 | 400 |
| the same with a wrong ground instance | 200 (error page) | 200, conclusion labelled `in` | exit 1 | 400 |
| the same scheme as a rule, no guard | 200 | 200 | exit 0 | 200, same graph (parsed comparison) |

The review was right on four points against the implementer's integration:
the profile gate refuses before conversion, not "before it is read"
(multipart parsing comes first); the integration called the server boundary
done while two route defects remained (they predate PR22, but PR22 changed
that handler); the D20 correction reached `d20-approvals.md` but not the
eleven live ledger notes and registry stamps; and the recommended order
(tag, then a non-author review of the fix) was backwards.

## The HTTP repair, read

`src/web/evaluation.go` gives both routes one boundary: authored-argument
validation, then `InferContext(req.Context())`, and labelling and export only
on success, with typed failures mapped to 400/408/422/500. The CLI also calls
`CheckContext()` before inference; `InferContext` performs the same context
resolution, so the only difference is which refusal is reported when an input
has several faults. `go test -race ./src/web/` passes. As the review states,
an export that fails after writing bytes, request size and the Dung route are
outside this repair.

## The attribution correction, checked

All 149 semantic fingerprints are unchanged, and exactly the eleven D20
entries of the approval registry changed. Two precision notes:

* The old notes said the implementer wrote neither the record "nor its last
  change". Tracing each record's fingerprinted content through history, that
  was false for six records and literally true for five, whose last change
  was made by another reviewer before the D20 approvals (`f208939`):

  | Record | Last fingerprinted change | Non-author check of it on record |
  |---|---|---|
  | `expert_opinion` | `2bc28ee`, Claude: assumption `credible(W)` | `first-half-adjudication.json` (tree `49ad425`) |
  | `ignorance` | `481bc40`, Claude: compendium citation | `first-half-adjudication.json` |
  | `correlation_to_cause`, `ethotic1`, `slippery_slope_base_case`, `slippery_slope_inductive_step` | `481bc40`, Claude: compendium citation | `2026-11-21-second-half-source-review.md` (tree `49ad425`) |
  | `positive_consequences`, `negative_consequences` | `25827f2`, Codex: likelihood-adequacy assumption | `second-half-cross-review.md` (the likelihood hook); the D20 comparison |
  | `established_rule` | `25827f2`, Codex: two scoped exceptions | the D20 comparison; the hooks are also assessed in `closure-hold-adjudication.json` |
  | `ethotic2`, `witness_testimony` | `344fcc8`, Codex | `closure-supplemental-independent-acceptance.md`; the D20 comparison |

* The new stamps say, correctly, that the D20 adjudication is not an
  independent review of Claude's own earlier edits. For the last three rows
  it was a non-author comparison of Codex's changes against the compendium.
  The new stamps cite the supplemental acceptance for `ethotic2` and
  `witness_testimony`, but for `established_rule` and the two consequence
  records they cite only reviews of the tree before those changes. The checks
  exist, as listed above; the stamps do not name them. No approval is in
  doubt, and the stamps are left as their reviewer wrote them.

## Kialo context

The committed refresher, run at `0a1e406` on the archive's historical inputs,
reproduces the archive's `pilot/context-0a1e406/` byte for byte (manifest
SHA-256 `8d433d1eb7e7c3e63f809330cda0b2be3e65801801181598142cd836adeebaae`;
eleven controls rejected). The twelve-unit blank packet is byte-identical to
the historical one; the six-unit source-only packet carries no proposed
labels. No Kialo data is in this repository. The archive records the
capture's redistribution rights as unverified and its acquisition permission
as not supplied.

## The validation pin

The accepted commit is `0a1e406913a06354c9014fd708cb78859eace23f` (tree
`b6cd8484ca527b75fb8d86473d4cd2a8ab29bf94`), merged by PR23 (merge commit
`fe964cf`, the same tree; CI passed on both Go versions with exact agreement
at `0a1e406`). This report is a
later commit that changes only documentation (no library, ledger, registry,
case, engine or tool file), so the pin and the prepared Kialo context stay
bound to `0a1e406`.

`PLAN.md` refers to the V0–V6 and K0–K6 programme, which is in the review
archive (`history/previous-framework-review/`), not in this repository; keep
the archive with the pin, or add the programme under `docs/`.

## Review of this report

This report is the implementer's and has not been reviewed by a non-author
(D14). It changes no attestation.
