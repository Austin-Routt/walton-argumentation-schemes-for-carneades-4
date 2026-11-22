# Independent source-closure review

Reviewed baseline: `98a2335d2c9909cf597ae85a9f359f20d092ae6c`.
Review date: 2026-11-22 UTC.
Reviewer: ChatGPT `final_sources`, a separate model-assisted reviewer who did
not author the reviewed Walton-record changes. This is a technical source
review, not human, publisher, or original-author certification.

## Decision

The newly supplied books resolve the outstanding source questions for
`practical_reasoning1`, `practical_reasoning2`, and `exeption_to_rule` within
their explicitly declared interpretations. The new omission branch of
entry 33.3 is defensible and fills real coverage. The existing six
Scare-Tactics-related forms remain acceptable as compendium interpretations.
There is no source-based reason to remove these changes.

The baseline is nevertheless not completely accurate. It retains misleading
record-specific citation locators in authoritative fields, misdescribes the
compendium's coverage of conditional threats, misidentifies the supported
premise in its Smart Car example explanation, and overstates completion of
independent acceptance. Its reported 149 approvals include nine explicitly
self-authored stamps. An additional D20 source-contract change to
`sunk_costs` also lacked a genuinely separate review of that change.

These are bounded, resolvable findings. The integration reviewer owns the
central YAML/ledger corrections. This review supplies non-author acceptance
for the four new Walton interpretations and the `sunk_costs`
reclassification, and acceptance of the citation-only corrections once
their exact final fields are inspected. A separate non-author Atkinson
review owns the five AS1/AS2 interpretations and adapter. No new broad
opinion round or acquisition of every cited original is required by the
owner's adopted D20 scope.

The machine-readable companion `final-source-acceptance.json` distinguishes
retained historical acceptance from fresh source review. It must be refreshed
after final citation fields are applied; a recommendation for an older
fingerprint is not approval of a different representation.

## Evidence actually inspected

The supplied files were read locally. The source manifest contains their
SHA-256 digests. PDF page numbers below are one-based physical page numbers;
printed page numbers are the page labels on the source itself. They should
not be interchanged. The Scare Tactics scan has changing offsets, so a
single arithmetic offset does not correctly locate every printed page.

| Source | Exact inspected locator | Main use |
|---|---|---|
| Walton, Reed and Macagno, *Argumentation Schemes* (2008) | Printed pp. 326,332–334; PDF pp. 336,342–344 | Waste/sunk-cost distinction; entry 33.3/33.4; entries 34–37 and their shared references |
| Walton, *Goal-Based Reasoning for Argumentation* (2015) | EPUB `text_part0008_split_008.html`, §3.8, `page_91`/`page_92` anchors; `text_part0008_split_009.html` and `text_part0008_split_010.html`, §3.9, `page_93`/`page_94` anchors | Smart Car analysis; configured value-based and instrumental schemes |
| Walton, *A Pragmatic Theory of Fallacy* (1995) | Printed pp. 135,142,146–148,150–153,155–157,160; PDF printed-page+20 | Case 80a; original forms and question counts cited by the closure report |
| Walton, *Scare Tactics* (2000) | Printed p. 22/PDF 35;123/135;138/150;139/151;140/152;142/154;173/184;191/201;192/202 | Fear chain; conditional/disjunctive threats; omission; practical critical questions; smoking example |
| Walton, sunk-costs article (2002) | Printed pp. 488–490/PDF 16–18 | Waste, precommitment and commitment forms |
| Repository evidence | Current formulas/templates, ledger and approval registry; original 2019 commits `0b1dd955859b9ae9b4b8d2841a4dc2fb400c6c4f`, `16c09f2b1ba1f63529e8b51763b57e8bdf44df42`; current reference cases | Exact implemented interpretation, provenance, independence and expected behavior |

Decisive PDF pages were rendered and visually inspected, including the
printed omission parenthesis and unchanged conclusion on Scare Tactics
p. 140, the Case 80a/precedent distinction on the 1995 p. 148, and the
compendium's conditional threat immediately preceding its disjunctive
variant on p. 333. Other passages were checked against extracted text and
retained page locators. The 2015 source is an EPUB with explicit printed-page
anchors, not a PDF facsimile; the evidence claim reflects that distinction.

## Fresh record judgments

### 1. `practical_reasoning1`: accept the 2015 configuration

The value-based scheme in §3.9 has the same conclusion, four ordinary
premises, three assumptions and four exceptions as the current record.
The assumptions concern the legitimacy of the value, worthiness of the
goal and possibility of the action. The exceptions concern a more effective
route to the resulting circumstances, goal or value, and adverse value
side effects. The existing side-effect template's reference to some value
includes the selected value and other values, as the printed formulation
does.

Re-sourcing this inherited record to the 2015 configured Carneades scheme
is justified. It is not an implementation of all seventeen questions of
the 2007 AS1 merely because the book says it follows that paper. Keeping
the exact 2015 configuration separate from the five new AS records is a
correct separation of claims.

The Smart Car case is a constructed instance of that value-based reading.
The source itself permits alternative analyses. Its Figure 3.16 discussion
says V1 supports premise E, the bridge from lower fuel use to reduced
environmental harm, or alternatively the goal A. It does not say V1
supports factual premise B, that the particular Smart Car uses less fuel.
The case's executable facts are reasonable; its original source explanation
needs that precision correction. A successful test cannot make an inaccurate
description of a figure accurate.

### 2. `practical_reasoning2`: accept both corrections

The source's new instrumental configuration has three ordinary premises,
two assumptions and five exceptions. The goal-possibility assumption is
about achieving G; the old generic `possible/1` template called it an
action. `possible_goal(G)` accurately removes that conflation.

The printed side-effect exception requires effects that should be taken
into account. It does not restrict those effects to demoting values.
`side_effects_to_consider(A,S1)` accurately restores the instrumental
wording. Keeping the value-specific side-effect template for the preceding
value-based scheme is justified.

The reference Smart Car instance is an explicitly instrumental reading
of the source's goal of buying a new car. State and goal names are
application encodings, not claimed literal quotations. The `variants: true`
reference-case mechanism separately withdraws ordinary premises and
tests assumptions/exceptions under an isolated selected scheme; the root's
fresh integrated execution remains the execution evidence.

### 3. `exeption_to_rule`: accept as a declared case-based interpretation

Printed p. 148 contains both Case 80a and a formal argument from precedent.
Those are distinct evidence objects. Case 80a transfers an extension
granted to one student, on a family-illness ground, to a similar student's
request. The formal schema printed afterward derives the need to recognize
an exception or revise a rule from a legitimate contrary case.

The current record follows the case-based transfer in Case 80a. It does
not implement the later formal conclusion about changing a rule. That
formal conclusion belongs to `from_precedent`. The ledger correctly
labels the present record an interpretation and explains this distinction.

The correction from `relevant_differences(R,C)` to
`relevant_differences(C)` is appropriate: R is a rule, not the current
comparison case. The retained context contract means C is compared with
the application's current case; a multi-case application must keep its
case identities/context explicit.

The three printed questions on p. 148 directly accompany the formal
precedent schema. Applying them to Case 80a is an adapted mapping, not an
exact printed question set for a separately printed exception-to-rule
schema. The ledger explains the adaptation: rule applicability; legitimate
ratio of the cited case; and differences relevant to whether the same
exception category applies. That is acceptable within the declared
interpretation. It should continue to be described as an adaptation.

The 2019 history changes this affected record's ID and title without
changing its logical premises or conclusion. Commit `0b1dd95` also changes
unrelated diagram/template material, so claims about what the entire
commit changed should remain record-specific. The README is the owner's
accepted public provenance for private correspondence. It establishes
what this project records about that correspondence, not independent
authentication of the email or Walton's approval of the present executable
representation. Keeping the historical misspelling avoids a gratuitous
compatibility change.

### 4. `negative_consequences_reasoning_omission`: accept the interpretation

Scare Tactics p. 139 expressly extends consequence reasoning to omissions.
It also connects that account to the prudential inference whose conclusion
recommends performing the action. Page 140 prints the don't-bring-about
parenthesis in the first premise while retaining a single negative
conclusion in the compressed displayed form. Reading that shorthand
literally as an identical conclusion for both alternatives would not
faithfully express the preceding omission account.

The new branch explicitly represents failure to perform A producing an
outcome bad for the same agent, and concludes a recommendation to perform
A. It is a defensible contextual reconstruction of the omission branch,
not a literal transcription of the unchanged displayed conclusion. The
existing `claim: interpretation` and the ledger's bracketed reconstruction
are therefore important and should remain.

The separate omission-likelihood default avoids conflating the likelihood
of doing A with the likelihood of not doing A. The smoking example on
p. 192 supports the chosen reconstruction and the presence of opposing
side effects. The two competing recommendations are local positions in
the existing `should_do` issue. This is not a general proof that denying
an obligation entails an obligation to abstain, or a claim that all
deontic operators collapse to Boolean negation.

The three current cases have suitable discriminating expectations:
support for quitting, suspension of the omission-likelihood default,
and opposed recommendations when quitting's adverse effect is also
represented. Equal-weight engine rejection of both opposed positions is
a statement about the declared engine policy, not a source claim that
health harms and weight gain have equal real-world importance.

### 5. `sunk_costs`: accept the D20 source-contract reclassification

The original D20 author changed the record's compendium classification
from 25 to 24 and changed its feasibility/future-loss question numbering
from unspecified source additions to entry 24's CQ1/CQ2. The executable
premises were inherited, but those source-contract changes were authored
by the person who approved them. Their fingerprint is affected. A blanket
assertion that all twelve D20 records' last changes were made by someone
else therefore needs this qualification.

An independent source comparison now supports the change. Compendium
entry 24 and the 2002 article p. 488 give the waste form; entry 25 and the
article p. 489 give the precommitment choice form. The legacy record's sunk
cost and too-high-to-waste predicates are a quantity-oriented adaptation
of waste. Feasibility and reassessment of future costs/benefits are the
two entry 24 questions, with the selected future-loss exception a declared
static policy. `sunk_costs_precommitment` remains the distinct printed
precommitment form. This fresh acceptance closes the narrow independence
gap without changing that executable interpretation.

## Scare Tactics comparison and corrections

| Record | Verified source | Judgment |
|---|---|---|
| `negative_consequences_reasoning` | p. 140, reasoning-from-negative-consequences form | Existing action branch acceptable; point the record-specific citation here |
| `negative_consequences_reasoning_omission` | pp. 139–140; smoking illustration p. 192 | Accept reconstructed omission branch as interpretation |
| `negative_consequences_prudential` | p. 123 | Prudential form matches; no citation move needed |
| `pragmatic_alternative` | p. 142 | Disjunctive practical alternative matches the compendium reading |
| `ad_baculum` | p. 142 | Disjunctive threat matches; p. 140 is the conditional variant |
| `fear` | p. 22 | Source has two linked inferences; compendium/record flatten the intermediate ought-to-prevent conclusion into a premise |
| `danger` | p. 173 | Negative recommendation is qualified by balancing; declared defeasible issue evaluation preserves that reading |

The compendium's p. 123 reference is shared by the entry 33 group, where
it correctly locates 33.4. Its p. 140 reference is shared by the threat
section and concerns the conditional form. Correcting the individual
records' verified locators should preserve those inherited shared
references as provenance. This is citation disambiguation, not a claimed
confirmed publisher or author erratum.

The closure report originally says the conditional threat on p. 140 is
absent from the compendium. That is too strong. Compendium entry 35,
printed p. 333, is a conditional threat form and adds an explicit premise
about the speaker's capability. It is closely related to the p. 140 form;
it is not the disjunctive entry 35.1. No additional record is needed merely
to repair that documentary assertion.

The fear/threat forms are not exempt from critical assessment because no
questions immediately follow their displayed forms. Scare Tactics
pp. 191–192 explicitly apply practical-reasoning questions. Page 138 prints
five, and p. 191 refers to four; p. 138's observation that the second is
normally not in doubt offers an explanation for that reduction, but it
should be presented as a reading, not an authenticated authorial
correction. The current ledger appropriately discloses the broader
source discussion while following the compendium's absence of numbered
questions for entries 34–37. Under D20, that is a scope boundary rather
than an undisclosed completeness defect. A future full account of those
dialogue contexts would need explicit question/answer behavior.

## Other supplied originals and D20

The closure report's comparisons of the 1995 example, established-rule,
exceptional-case, ignorance, ethos, bias, consequences, threat and verbal
slope passages agree with the inspected source passages. Correlation-to-
cause has seven questions on p. 142, while the compendium uses three;
the notes correctly disclose that difference. The negative-normal and
negative-practical subtypes are not printed on the original pp. 150–151
as distinct schemes. Their compendium-based source claim should remain
distinct from that original-page comparison.

The 2002 article pp. 488–490 indeed places waste, precommitment and the
older/newer commitment forms together, supporting the closure report's
reported comparisons. It does not turn the legacy quantity form into a
literal transcription of the precommitment display.

D20 is explicitly adopted by the owner in the active decision register.
It makes the compendium the source of record for inherited records filed
under its entries. It does not prove the unseen originals agree. The
unavailable expert-opinion, slippery-slope and witness sources remain
unverified comparisons and future research; they are not release
blockers for this expressly scoped compendium product. Historical
reviews rejecting D20 as then-unapproved should remain historical
evidence and should not be copied over the later owner decision.

## Independence and approval integrity

At the reviewed baseline, all 149 ledger rows are `example-passed`:
120 book-origin, 24 upstream-origin and 5 publication-origin. The final-closure
document's 119 book-origin statement is stale after adding the omission
record. Exactly nine review stamps explicitly admit that the implementer
also wrote the change. Running the old integrity tool succeeded for all 149
because it checked matching content/metadata, not the truth of claimed
independence. That success did not satisfy D14.

The additional `sunk_costs` contract correction above makes ten new
non-author acceptances appropriate in this closure round. The five
Walton judgments are here; the five AS1/AS2 judgments are provided by
the separate Atkinson reviewer. Citation-only refreshes are separately
inspected and re-fingerprinted.

The other retained approvals are backed by named historical evidence.
Eight early repaired records originally carried their authors' stamps:
`abductive_forward`, `abductive_action_to_character`,
`abductive_character_to_action`, `retroductive`, `from_model`,
`rule_physical_world_1`, `exceptional`, and `memory`. Their genuine
non-author second acceptance is explicitly recorded in
`docs/reviews/implementer-response.md` §2. Current attribution should
identify that second acceptance rather than label the original author
independent. The later branch, parts, context, dialogue, family and
supplemental acceptance reports explicitly separate implementation and
review authors. The eleven unchanged D20 records have the implementer's
documented non-author source review; `sunk_costs` is handled separately.

An explicit Boolean `independent_of_author` plus a nonempty
`independence_basis`, bound to the registry snapshot, is a worthwhile
bookkeeping repair. It prevents a disclosed self-review from being treated
as independent merely because all hashes match. It still cannot
authenticate a reviewer or establish the truth of the attestation.

I independently reviewed the integrating reviewer's checker changes.
Initial adversarial probes showed that metadata reconciliation could
overwrite an explicit withdrawal or a false independence claim with the
registry's approval, and could write an invalid reviewer/date. The
integrator repaired those cases. The repeated probes now refuse them
before writing and preserve file bytes. A further typed-snapshot probe
found that integer 1 in the registry could compare equal to Boolean true.
The integrator repaired that case by validating the registry stamp as well
as the ledger stamp. All five independent probes now fail appropriately,
before any write, and all 23 approval-checker tests pass. The companion
findings record exact reviewed code hashes and the individual results.

## Migration acceptance

The last three migration rows are correctly manual. `possible/1` remains
the action-possibility predicate, so a graph-wide rename would corrupt
ordinary action defaults. Only the goal-typed occurrence in the
instrumental record should become `possible_goal/1` after its context is
checked. The source locator for that assumption is printed p. 93, not p. 94.

Likewise `side_effects/2` remains valid for the value-based record. Its
instrumental use must be reassessed against the broader generic side-effect
question; it cannot be renamed everywhere. The old two-argument
`relevant_differences/2` statement needs the actual cited/current-case
comparison recovered before the rule argument is dropped. An automatic
arity edit would hide that judgment.

An independently constructed mixed graph containing all three old forms
was preserved byte-for-byte by the migrator and produced all three manual
warnings. The integrator's targeted regression additionally covers positive,
negated and questioned occurrences. The complete 17-test migration module
passes. This establishes preservation and reporting, not automatic semantic
equivalence of old and new facts.

## Conditions for integration

1. Correct the three record-specific citation fields and their category
   copies, while preserving inherited shared references in notes.
2. Correct the conditional-threat and Smart Car explanations, the 120
   book-origin count, and D21's stale statement that three records remain
   held after their sources have now been supplied and compared.
3. Replace the nine self-stamps and the `sunk_costs` contract self-stamp
   with the actual non-author decisions. Preserve honest historical
   attribution for the other approvals, including the eight prior
   second-check acceptances.
4. Require explicit independence attestation in both ledger and registry,
   with no automatic reconciliation of contradictory or withdrawn
   decisions. Bind fresh citations/contract changes to their new hashes.
5. Run the full integrated strict gates and cross-version comparison on
   the exact resulting source tree. This report supplies source judgment;
   it does not substitute for that execution or the separate engine/AS2
   reviews.

After those concrete actions, the source-closure decisions can be finalized
for the declared profile. Remaining empirical validation, benchmark
performance, user annotation agreement and real-world argument adequacy
belong to the next validation effort. They should not be advertised as
already established by source approval or by the number of passing checks.

## Final integrated source inspection

The integrating reviewer has applied the source citations, category copies,
manual migration locator, 149 review recommendations and v2 approval
bookkeeping. I independently ran `python tools/check_source_approvals.py
--require-stable` against that integrated tree: all 149 records passed.
I also compared every current semantic fingerprint and registry review
snapshot with the independent recommendation file, and compared the three
corrected category records with their master copies. There were no
discrepancies. The omission row explicitly preserves the former shared
p. 123 reference as provenance.

`integrated-source-acceptance-check.json` records that comparison and hashes
of the master library, ledger, registry and category file. This closes the
source and review-accounting conditions above, conditional on preserving
those reviewed representations. The integrating reviewer's full clean-copy
engine/suite evidence and the separate AS2 code acceptance still govern
the final release-candidate decision; the source-stage gate alone is not
the complete acceptance claim.
