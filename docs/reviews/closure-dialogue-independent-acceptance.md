# Independent acceptance of dialogue, burden and evaluative interpretations

Date: 2026-11-21. Reviewer: ChatGPT integrating independent reviewer.
The scope reviewer authored the dialogue adapter and three practical scope
repairs; the source reviewer authored the two presumption branches; the scope
reviewer authored the final importance-template clarification. This reviewer
did not author those implementations. The separately authored assumption-origin
engine repair has its own non-author acceptance record.

## Decision

Approve the eleven interpretations below within the stated supported profile.
Source approval is bound to each current scheme/ledger semantic fingerprint
and attributed in the approval registry. It does not certify unseen cited
originals, autonomous answer evaluation, unrestricted dialogue or all possible
graphs. The final complete clean-copy matrix remains a separate gate.

| Record | Source inspected | Reason and scope of acceptance |
|---|---|---|
| `practical_inference` | Compendium p.323; narrative p.96 | Goal, means, acting agent, efficiency and consequences retain their identities. The five appropriate questions can suspend this argument through the adapter; accepted independent answers remove only that question's attack. Static exception policies remain declared separately. |
| `practical_necessary_condition` | pp.95–96,323–324 | The selected action and alternative-set witness are distinct. Ordinary premises state the necessary-condition selection and practicality. The five-question dialogue contract now has executable suspension and restoration. |
| `practical_sufficient_condition` | p.96 and pp.323–324 | The narrative's each-one-sufficient statement governs the declared correction to the compendium's inconsistent alternatives wording. No confirmed authorial erratum is claimed. The same five-question protocol preserves agent and set identity. |
| `value_based_practical_reasoning` | p.324, with p.96's dialogue discussion | Values, alternatives, efficiency and consequences now belong to the acting agent. All seven questions have explicit challenge handling; value/goal facts and answer adequacy remain supplied application assessments. |
| `two_person_practical` | pp.325–326, with p.97 | The actor and adviser remain distinct; alternatives and additional actions concern the actor. All five questions are tested. The adviser reports a necessary or sufficient condition, rather than arbitrarily creating one. |
| `position_to_know` | Compendium p.309; Walton2002 indexed pp.46–48 | The domain, assertion and source fit the cited general formulation. Its three questions receive the source's explicit burden-shifting behavior through the adapter. The dishonest-source exception is a declared static policy. Indexed text was inspected; no full-book facsimile certification is claimed. |
| `negative_normal_expectations` | Compendium p.327 | Explicit absence of the expected proposition combines only with the same situation-qualified expectation. `not_normal(A)` names that expectation context. New mismatched-context controls prevent mixing situations; burden/probandum/standard allocation is explicit application metadata. |
| `negative_practical` | Compendium p.327 | The B-more-serious ranking supports acting on the false presumption, not asserting the proposition false. Search adequacy is challengeable; the ultimate issue, burden holder and standard are supplied separately. |
| `negative_practical_true` | Compendium p.327 | The opposite C-more-serious ranking supports the true presumption, not truth. The optional explicit binary-choice issue is an application protocol; no global choice/negation equivalence is installed. Both opposed rankings can preserve arguments without accepting either choice. |
| `end_and_means_2` | Compendium p.325 | Reclassifying a means as an end plus the same-context importance warrant supports a high-importance assessment. The clarified language denotes an evaluative judgment, not an obligation to perform an action. |
| `end_and_means_reverse` | Compendium p.325 | Reclassifying an end as a means plus the ordinary importance warrant supports rejection of that assessment. Lack of positive evidence alone supplies no negative argument. The supported positive and negative assessments compete only for the same X,C. |

I inspected the supplied printed p.96 facsimile and rendered pp.323–327,
including the printed parentheses and the three ignorance questions. The
Walton2002 source was re-opened at the accessible page-marked full text:
https://www.academia.edu/37757995/_Douglas_N_Walton_Legal_Argumentation_and_Eviden_pdf
Printed pp.46–47 give the formulation, three questions, temporary burden shift
and restoration after a satisfactory answer. The already retained source
review supplies p.309 and p.97 comparison evidence.

## Adapter code and adversarial review

I inspected the complete adapter and actual-engine tests. Exact targeting
uses argument ID, scheme and ordered ground parameters. Open questions add
supported attacks on that target's undercutter, preserving other arguments
and original attacks. Every snapshot is rebuilt from the caller's immutable
base. Old labels are removed before evaluation. Shared or issue-competing
undercutters are refused because this protocol cannot guarantee suspension
for them.

The author repaired two issues raised during review: legitimate quoted
identities are accepted without losing spaces or escapes, and answer evidence
that depends indirectly on the challenged conclusion or undercutter is
rejected. The initial quoting limitation was rejection of valid quoted terms;
it was not a demonstrated accepted collision between two quoted identities.
The dependence check is intentionally conservative and can reject a statement
with both dependent and independent proofs. It also does not recompute answer
acceptability for every combination of other open challenges. Adequacy and
appropriateness are explicit external judgments; accepting a label alone is
not presented as automated source reasoning.

Burden metadata names the argument, dialogue, ultimate probandum, burden
holder, issue and proof standard. The ultimate probandum may be a broader
dialogue claim. The program validates reference consistency, not the logical
pertinence of that broader issue or the institution's allocation decision.
Tests exercise all four ignorance-family records. This closes the declared
representation gap without claiming that an evidence threshold alone decides
who has to prove what. The inherited `ignorance` record still lacks its cited
1995 original comparison and is not promoted by this decision.

The p.325 clarification is justified by its evaluative context. Explicit
rejection of a high-importance assessment is tested with the reverse
reclassification warrant. It cannot be inferred from absence of positive
proof, and it does not generate an action prohibition. This differs from the
anti-model source, which requires an explicit recommendation to refrain.

## Independent execution

The integrating reviewer ran the thirteen actual-engine Python test groups
with Go1.24.13: all passed. They cover all thirty CQs across six families,
states before/after questions, accepted answers and reopening, independent
support, retained original defeaters, multiple questions, malformed journals,
quotes, indirect dependence, the Python CLI and all four burden families.
These groups are part of the required runner and are separate from registered
Go result counts.

The focused book/family/source-branch/context/scope selection passed482/482
registered API/CLI checks. After the importance clarification, its new
absence-of-evidence control passed2/2. Prior accepted positive and reverse
assessment cases remain in the full combined gate. All checks use the
strengthened observation-only CLI oracle.

| Evidence | SHA-256 |
|---|---|
| `practical-audit/dialogue-independent.log` | `0df07951a59a74792685c06250775d33cf3dc0886dffeeb792bf5be897b723b8` |
| `practical-audit/source-cases-independent.log` | `d42f4e9ea3d347441c69be0893076172ab7e76c8238ac7b9f4bd3ecc1629ae71` |
| `practical-audit/source-cases-independent.jsonl` | `649a4d3d8c8ef4ab6fba297e3e7bc66756d8f0de37a9cd10f911c386e5eefe3b` |
| `practical-audit/assessment-independent.log` | `e1429315463bc41d62d7f5471d3e6e0db4dddd6ae96ff779eca7d5d87a156af6` |

The first focused run follows engine repair e6e4015 and precedes the later
importance wording; the last targeted control uses the integrated clarification.
The final acceptance evidence identifies one frozen commit and reruns all of
these together on both required Go versions.
