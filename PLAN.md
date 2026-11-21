# Plan after independent acceptance review

Updated 2026-11-21. Reviewed baseline: `49ad425cbf7afeebcbd9a07ee479d60681ca33ea` (PR #12 merged). Former proposed R1: `8c092b6d45106a3262a95629a124993a5d72c058`. This plan supersedes the prior operational instructions. The complete previous plan remains in `docs/reviews/plan-before-independent-review.md`.

## Decision and present state

Proceed with the corrected **preview**, library **0.2.0** and engine **4.3.1+walton.4**, after its clean acceptance runs. **Do not publish the former R1 commit.** Independent counterexamples exposed engine and source defects its green tests did not cover. **Do not promote the complete library to 1.0.0 because tests pass.**

The author's measurements are reproducible: clean baseline copies pass all 7,994 registered checks on Go 1.24.13 and Go 1.27.4, including the separate engine build, vet, unit and race gate. A separate full clone of former R1 passes 7,164 checks on Go 1.24.13. These results validate the reported executions, not completeness of those tests.

| Item | Current disposition |
|---|---|
| Library records | 120: 96 book-origin and 24 inherited |
| Mapping | 114 records under 109 compendium entries; six upstream-only |
| Independent source review | All 114 mapped records examined against the supplied book; attached Atkinson and Bench-Capon paper reviewed in full |
| Source stages | 88 `example-passed`, 30 `transcription-checked`, two `mapped` (70, 48 and two before the follow-up of 2026-11-21) |
| Approved records | 78 book-origin and ten inherited, within their documented interpretation and provenance |
| Remaining source holds | 32 named records, each with one decided action (`docs/reviews/2026-11-21-reviewer-follow-up.md`, section 6) |
| Previously deferred five questions | Modelled from the supplied compendium; verification of unavailable cited originals remains separate |
| Atkinson coverage | All 17 questions reviewed; eight lack a complete independent mapping in the inherited AS1 form |
| Runtime gate | Added logical reuse, streaming/bounded matching, complete context closure, atomic inference, strict nested parsing and authored-exception checks |
| Stable gate L | Open; requires every record to meet both source and executable gates |

The exact current check inventory is `tools/vv/expected_checks.txt`. The review archive holds clean-run logs, counterexamples and all per-record adjudications. Source judgments are preserved under `docs/reviews/` and in `tools/vv/source_ledger.yml`.

## Decisions retained and clarified

| Decision | Required contract |
|---|---|
| D1: engine ownership | Keep sources in `third_party/`, the original archives and `sources.lock`; no external forks. Preserve focused patches. Coordinated adapter/engine fixes may share a commit when their common contract requires it. Prepare upstream patches by source root. |
| D2: questions | Preserve premises, challengeable defaults, exceptions and rebuttals. Record positive support, alternatives and dialogue/proof-standard questions separately where required. Asking a question is not the same as proving an exception; a proof standard does not implement burden allocation. |
| D3: supported profile | Offline CLI and Go API, propagation rules, ground logical facts and supported terms, with serialized access. Web hosting, arithmetic, guards/deletions, concurrent mutation of one graph and evidence authentication remain outside acceptance. |
| D4: versions and migration | Corrected R1 is preview 0.2.0/.2; R2 is complete stable 1.0.0. Source repairs can require semantic migration before stable release. Use the migration table/tool; never invent newly required witnesses or participant identities. |
| D5: source authority | The compendium governs numbered entries; its narrative resolves ambiguities where justified. Inherited originals retain separate provenance. A book-form adaptation is not verification of an unavailable original. Record interpretations without claiming confirmed authorial errata. |
| D6: shared vocabulary | Merge only propositions with the same meaning, participants, scope and polarity. Preserve agent, event, action, quality, case and diagram identities. A typed trait may project to a property; an arbitrary property must not predict character. |
| D7: required gates | Both pinned Go 1.24.13 and Go 1.27.4 must pass engine and strict gates, and all named results must agree. Neither version nor comparison is informational. Exact inventories, hashes and mutation tests prevent silent weakening. |
| D8: facts and work | Logical facts form a set and can satisfy repeated premises without hidden inequality. Stream matching; bound applications, term nodes, candidate visits, constraints, stored nodes and issue construction. Logical quotas are not an operating-system byte-memory or universal time bound. |
| D9: outcomes | Keep the seven typed outcomes. Failed/cancelled inference leaves the original graph unchanged; success preserves existing node identities. Resolve issue closure/default priority before seeding inference and again after derived bridges. Reject incompatible existing issues rather than choosing a proof standard silently. |
| D10: opposition | Detect contradictory stipulations through explicit negation and generated/overlapping issue schemes, even without rule firing or predeclared statement nodes. Defeated defaults must not seed further defaults. |
| D11: syntax | Reject malformed nested terms and preserve valid quoted atoms on both toolchains. Keep the supported reserved-name policy. A partial parse is not valid input. |
| D12: authored arguments | Validate a single substitution and every direct linked exception witness to the actual undercutter. An inert undercutter, wrong instance, extra prerequisite or disabled weighing function is insufficient. API clients must call `validation.Validate` for authored graphs; inference does not replace validation. |
| D13: source plus behavior | Every repair needs a source locator, explicit interpretation, positive example and a discriminating adverse example. Source judgment and executable verification are separate evidence. |
| D14: independent acceptance | Review source passages and expectations independently of the implementer's tests; rerun clean copies and record disagreement. AI-assisted review supplies technical evidence, not institutional or authorial certification. No new broad opinion round is required; reopen concrete holds/regressions. |

## Decisions D15 to D23 (follow-up of 2026-11-21)

Made by the reviewer of 2026-11-21 at the owner's request, after the implementer's response left them to a reviewer or to the owner. The reasons, and the action each decision gives every held row, are in `docs/reviews/2026-11-21-reviewer-follow-up.md`. Where a sentence elsewhere in this plan says otherwise, these decisions govern; the owner can overrule any of them.

| Decision | Required contract |
|---|---|
| D15: release | `8c092b6` is withdrawn. `r1` is the corrected candidate with engine 4.3.1+walton.4, tagged at the master commit that contains that engine after someone who did not write its changes has checked them and both Go jobs are green. |
| D16: questions and burden | The library implements the book's formal account of critical questions (pp. 33, 373–375): each is sorted once and evaluation is static. The dialogue of p. 96 is not implemented and not claimed. Questions about the burden or standard of proof are answered by how the application builds the graph. No row is held for this. |
| D17: variants | A record states one variant of its entry; the ledger names it and what is left out, and that is a complete record. 9.2 is the species reading with the plain conclusion. 10.2 takes a part premise before the stable release. |
| D18: participants | A statement carries a participant when the print names it in that statement; an exception carries enough of them that one instance cannot undercut another. First- and second-person schemes keep one proponent and one respondent. Inherited records keep upstream's arguments and declare the limit. |
| D19: polarity | The engine weighs positions and has no degrees: a negated recommendation reads "should not", "more plausible" is an argument for a statement and "rejected as less plausible" an argument for its negation. |
| D20: source of record | The compendium. An inherited record filed under an entry is complete when it is compared with that entry, every printed question has a disposition, differences are declared and its examples pass. Agreement with the older work it cites is recorded when that work is at hand and is not a condition of the stable release. |
| D21: legacy records | `practical_reasoning1`, `practical_reasoning2` and `exeption_to_rule` stay as Carneades wrote them, are marked legacy and are not counted as source-checked. No record for the fuller model of Atkinson and Bench-Capon. The stable gate is 117 records at `example-passed` and three legacy records identical to upstream; `tools/check_source_approvals.py --require-stable` is to learn that class. |
| D22: outside contacts | Upstream patches are regenerated from the released engine and the two drafts for the authors are merged. Both are sent only when the owner says so; neither holds a release. |
| D23: entries 33.3 and 33.4 | They take the reading of 33.1 and 33.2: a likelihood adequate for the argument, with no fixed threshold. |

## Source stages and stable gate L

The stages are cumulative evidence states:

| Stage | Required evidence |
|---|---|
| `mapped` | Identity and provenance located; content unapproved |
| `transcription-checked` | Premises, conclusion, variables and differences compared with accessible text |
| `interpretation-approved` | Independent source judgment accepts the executable interpretation and declared scope |
| `questions-sorted` | Every printed question has an explicit justified disposition; missing treatment is a hold; verified zero differs from unknown |
| `example-passed` | Prior requirements plus independently examined examples and successful supported execution |

`tools/check_source_approvals.py` requires a current recorded approval for every row at or beyond interpretation approval. Its fingerprint covers the complete scheme, statically referenced language, ledger semantics and global issue schemes. It verifies that recorded example IDs exist. This detects stale decisions; it does not automate philosophical judgment, authenticate reviewer identity or verify application facts. `--fingerprints` is inventory only and creates no approvals.

The incoming independent review in PR #13 (`96daf6c`) is merged and preserved under `docs/reviews/peer-13/`. Its reviewer/date/decision validation is retained; a matching attribution snapshot is required in each approval registry entry. `tools/apply_source_review_metadata.py` can reconcile already adjudicated metadata, but refuses stale semantic hashes and cannot promote a row. The complete source and release differences are resolved in `docs/reviews/pr13-reconciliation.md`; its 90 approvals do not supersede the current 70 justified approvals and 50 explicit holds.

For 1.0.0, run both ordinary gates and:

```
python3 tools/check_source_approvals.py --repo . --require-stable
```

All 120 records must reach `example-passed`, except the three legacy records of D21, which must be identical to upstream and marked. Do not silently shrink the release scope or count qualified interpretations as complete. A selectively approved preview is useful but does not satisfy the complete-library gate.

## Remaining implementation sequence

### Defaults for the remaining technical decisions

These are implementation directions, not a menu requiring owner approval. Prefer explicit participant, case, inquiry and institutional parameters when a held record must support several contexts in one graph. Preserve the existing limited branch as an honestly named variant while adding a separately tested source alternative; do not silently change its polarity or strength. Keep the current engine's static evidence profile explicit, and implement any question-and-answer dialogue protocol as a separate layer with tests for suspension, answers and withdrawal. Treat numerical likelihood calibration as application input unless the source supplies a particular rule. Keep the inherited AS1 record distinct from any later transition/value model. Never infer missing migration witnesses or claim to have verified an unseen original.

Some per-record actions list more than one technically defensible implementation. The defaults above settle the preferred route; another route requires its own documented source justification and adverse example. No further general review vote or owner approval round is needed.

### A. Corrected preview

1. Integrate the independent engine/source repairs, migrations and adjudications; preserve earlier evidence.
2. Regenerate categories and inspect every expected-check addition/removal. A release run must not set `VV_UPDATE_CHECKS`. Keep the known-failure list empty only when every strict check passes.
3. Commit the candidate. From clean full clones, run `tools/vv/engine-tests.sh` and `VV_STRICT=1 tools/vv/run.sh` with each pinned Go binary. Compare every named result. Record commit, actual version, environment, clean status, exit codes and log hashes. A linked-worktree VCS-stamping failure is infrastructure evidence, not a product defect or a waived gate.
4. Publish `r1` only at the accepted corrected candidate, using `docs/releases/r1-candidate.md`. The old `8c092b6` tagging command is superseded. Distinguish patched source assets from the original engine ZIPs. Without authenticated GitHub write access, deliver the tested bundle, patches and exact publication commands; do not claim a remote release exists.

### B. The source holds (50 at the acceptance candidate, 32 after the follow-up of 2026-11-21)

Decisions D16 to D21 settle the choices the items below left open: item 1 is no longer a condition of the stable release (D20), item 4 is decided for static evaluation (D16), and item 5 is decided for the legacy class (D21). The action for each of the 32 rows is in `docs/reviews/2026-11-21-reviewer-follow-up.md`, section 6.

Resolve these in targeted batches using each record's adjudication, not another general review round:

1. **Inherited originals:** obtain cited original pages and attributable 2019 correspondence. The compendium verifies its own form and questions, not an unseen original. Unknown question count remains unknown.
2. **Participant/context scope:** complete qualified person, dialogue, domain, causal application and event identities. Use two-instance isolation examples before approving broader use.
3. **Branches/modalities:** distinguish a selected branch from a complete entry; resolve reduced plausibility versus explicit negation (`ethotic2` remains blocked), generalization strength and comparative probability without strengthening the source silently.
4. **Practical/ignorance questions:** implement the intended burden behavior. A request for support and a proved exception have different effects. Issue proof standards alone do not settle dialogue obligations.
5. **Atkinson and Bench-Capon:** retain `practical_reasoning1` as a qualified inherited AS1 form. Its source is now available and reviewed; access is no longer the blocker. Resolve the eight incomplete mappings in `docs/reviews/atkinson-2007-decisions.json`. A separately named AS2 extension requires agents, joint actions, states, values and preference semantics. Alternatively, approve a precisely justified AS1 scope without claiming full-paper coverage. CQ10 is positive support; goal worth differs from realizability; alternative existence differs from greater effectiveness.
6. For each resolved row, update source reasoning and migration, add discriminating cases, obtain independent judgment, record a new approval fingerprint and rerun acceptance. Reopen accepted records only if their meaning or dependencies change.

### C. Stable R2

Complete final interface cleanup/migration, all source holds and the stable approval gate; rerun both toolchains cleanly. Publish 1.0.0 only after those conditions hold. No date is inferred from an assumption that unavailable sources or disputed interpretations will resolve themselves.

### D. Upstream and correspondence

Prepare source-root-specific patches, minimal counterexamples and explanations of preserved upstream behavior. Prepare a focused source-question memorandum. Keep outside messages unsent until the owner explicitly authorizes contact; upstream acceptance is not a prerequisite for the local preview.

## Completion criteria

This independent-review task is complete when the local corrected candidate, reproducible evidence, revised plan, complete source dispositions and one review archive are delivered. Remote publication is access-dependent. Complete stable content approval is the explicit record-specific work above. A passing test total does not erase either distinction.
