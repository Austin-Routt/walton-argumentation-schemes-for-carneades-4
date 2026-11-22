# Atkinson and Bench-Capon (2007): every critical question in the examples, and every discrepancy

2026-11-21. A second, line-by-line check of sections 5-7 (pp. 861-872) against the models in `tools/aats_models/` and the definitions in `tools/aats.py`. The formal definitions were read from the rendered printed pages, not from the text extraction, which drops the "≠" signs. Every row below is checked by `tools/test_aats.py` (classes `PaperQuestions`, `FarmerFigure1`, `Formulation`, `Formal`, `Contract`).

Reviewer: Claude, the implementer, who also wrote the models and the tests. The owner's independent research is the second check (D14).

## 1. The seventeen questions in the farmer example

The paper says the example gives at least one instance of each of the seventeen questions (p. 864; "the complete set", p. 863). It names thirteen: CQ1, 2, 3, 4, 6, 8, 9, 11, 12, 13, 14, 16 and 17. CQ5, CQ7, CQ10 and CQ15 are never named. They are inferred below from the context and the formal definitions of p. 862.

| CQ | Where the paper names it | What the model shows | Status |
|---|---|---|---|
| 1 | p. 864: an argument based on an initial state other than [BCDS, _] | printed as a conjunction: the believed state is not the actual one **and** the action cannot be done in the actual one. Arguing from q18 that the farmer should go home meets it; arguing from a wrong state for an action that can also be done in q0 does not, and fails CQ2 instead | named; the printed definition is narrower than the informal question (section 2) |
| 2 | p. 864: a joint action claimed to reach q_n with n > 9 | the claimed state is not τ(q0, jn) | correct |
| 2 | Obj6ai (p. 866): "the chicken won't eat the seeds" | j7 → q7 and j8 → q8 both exist; which one follows is the chicken's choice | **relabelled CQ17** (section 3) |
| 3 | p. 867: the objection to Arg10 that staying in q9 does not repeat it | with non-movement counted as repetition, as the paper decides, the goal `repeated` changes | correct |
| 4 | Obj2c (p. 865) | in a formulation where eating the chicken does not promote DH, δ(q0, q3, DH) is not + | correct |
| 5 | not named | no two joint actions from one state reach the same state, in any of Figs. 1-8 or any row of Table 1 | **inferred: cannot arise** in either example. The nearest the text comes is the bridge (p. 864), a different way to the final consequence (everything across), which fails CQ13 first. Tested on constructed models |
| 6 | Obj4b (p. 865); q10 (p. 868) | P is promoted "through the goal of transportation" (p. 864), modelled as right_k, at least k possessions on the right bank; rowing the seeds and rowing the chicken from q0 both make right_1 true | correct with that goal, and with CQ6 read with qz (section 2) |
| 7 | not named | every value of the farmer has one kind of goal, so each CQ6 above is also CQ7; also Arg5 against Arg2 (both promote DH) and Obj4c (rowing the dog also promotes P) | **inferred**: the paper names the narrower CQ6 |
| 8 | q14 (p. 868): rowing the chicken promotes P and also demotes P | +P by right_2 and −P by `repeated` on one transition | correct; needs values attached to goals (section 2) |
| 9 | Obj1b, Obj2d, Obj3a, Obj5a, Obj7a; with CQ17 in Obj2a, Obj2b, Obj6a; q11 (Arg14, Arg15) and q15 (p. 868) | another value demoted by the same transition | correct, with one loose use: at q14 (p. 868) "rowing alone falls to the objection ... (CQ9)", but nothing supports rowing alone there, so there is no argument for CQ9 to question; the objection is the negative version of p. 859, an argument against rowing alone |
| 10 | not named | Arg5 (P and DH), Arg6 (P and F), Arg7 (P, F and CH): each value supports the action the others support; in q10 rowing the dog promotes "F as well" (p. 868) | **inferred**; p. 862 itself says CQ10 endorses the action |
| 11 | Obj2f, Obj4c (p. 865); q10 (p. 868) | another joint action promotes another value | correct |
| 12 | p. 864: the dog muzzled in the initial state | the described state is not in Q | correct; the reason the paper gives, that the representation has no such proposition, is CQ15's condition |
| 13 | p. 864: crossing by a bridge | no such joint action | correct |
| 14 | p. 864: the dog muzzled in a consequent state | the resulting state is not in Q | correct; as for CQ12 |
| 15 | not named | a goal that holds in no state: the muzzled dog as the goal of an argument | **inferred** from the muzzled-dog remark, applied to a goal |
| 16 | Obj1c, Obj2e (p. 865) | after the reformulation CH and DH are not values of the farmer | correct |
| 17 | Obj1a, Obj2a, Obj2b, Obj4a, Obj6a (pp. 865-866); the trains (pp. 870-872) | another joint action with the same part for the agent leads elsewhere; answered when the value is promoted there too (Obj2bi; the eastbound train in q7) or by an assurance (a social law) | correct |

## 2. The formal definitions (p. 862)

1. **CQ1** is printed as q0 ≠ qx and q0 ∉ ρ(α_i). The informal question asks whether the believed circumstances are true (p. 858), and p. 864 applies CQ1 to any other initial state. `tools/aats.py` follows the printed conjunction when it checks a claimed argument. The AS2 record states AS2's own first line, q0 = qx, and the precondition as separate premises, so a false belief about the state removes the argument even where the printed CQ1 is not met.
2. **CQ3** (pa ∉ π(qy)) and **CQ15** (pa ∉ π(q) for every q) are printed for a goal that becomes true. AS2 also allows a goal that becomes false (p. 861). For such a goal the printed CQ3 would object to every correct argument, so both are read the other way round for it.
3. **CQ6** is printed with "τ(qx, jm) is qy". AS2 already requires pa to change between qx and qy, so with qy CQ6 says exactly what CQ5 says. It is read with qz, as CQ7 and CQ11 are printed, as the informal CQ6 ("the same goal") and p. 862 ("the consequences, goal and value, respectively") require, and as Obj4b needs (rowing the chicken reaches q9, not q5). `test_printed_cq6_would_be_cq5` checks the equivalence on every argument of both models.
4. **CQ8** asks that δ(qx, qy, vu) be − on the transition where AS2 says it is +. δ is a function into {+, −, =} (p. 861), so as printed CQ8 can never hold. Fig. 6 labels one transition +P and −P, and footnote 5 counts "two harmful goals". The models therefore attach each promotion and demotion to the goal that realizes it.
5. **CQ5-CQ7 and CQ11** require only jm ≠ jn, not that the agent's own part differs. Read literally they also count the other agents' choices, which CQ17 covers. `tools/aats.py` follows the letter; the comparison contract (`docs/contracts.md`, contract 8) compares the agent's own actions.

## 3. Relabelled and corrected

* **Obj6ai (p. 866), CQ2 → CQ17.** It says the chicken won't eat the seeds, the same content as Obj1a, which the paper labels CQ17. Both transitions from q0, j7 to q7 and j8 to q8, are in the model; which one follows depends on the chicken. P. 859 assigns exactly that case to CQ17 and keeps CQ2 for a transition that does not exist. At the AS1 level, where the source of AS1 subsumed CQ17 under CQ2 (p. 859), CQ2 would be the label.
* **Obj4b is not mislabelled.** The first review of this round (`final-closure.md`, finding 2) said it was, reading the arguments' goals as the items moved. The paper's own definition of P puts the goal at transportation; with that goal, CQ6 is right. The models now encode it (`right_k`).

## 4. Figures against text

1. "We have ten joint actions" (p. 863) is followed by eleven, j0 to j10.
2. The state lists are said to give the right bank first and the left second (p. 863), but the initial state is [BCDS, _] and moves put possessions into the second list: the first list is the starting (left) bank.
3. Figs. 1 and 2 omit j0, doing nothing in q0, which Fig. 3 shows in q9. Adding it, a repetition demoting P, changes no choice (tested).
4. Fig. 1 labels j4 −S −C +DH without +CH, although the chicken eats the seeds there too (Obj2b). Adding +CH changes no choice and leaves Arg1's CQ17 unanswered (tested).
5. Fig. 5 shows one −P on q11 to q10 and Fig. 7 one −P on q15 to q14, where the text counts two harmful goals (Arg15, "severely"; "in two ways", p. 868). A sign label need not encode magnitude, so this is a difference in detail, not by itself a false figure. The models follow the text. **The distinction matters in Fig. 5:** if its single −P label is interpreted as one unit of demotion, the comparison contract selects rowing the dog when P is preferred to F; the two-goal textual reading selects rowing the chicken, as the paper says. The new sensitivity test demonstrates both choices. Fig. 7's reported choices are unchanged by using one instead of two penalties.
6. Rowing the chicken back from q9 to q0 (Fig. 3) and from q18 to q16 (p. 869) revisit states, but the text counts only the move to the left bank. The models follow the text; counting the revisit in q9 changes no choice (tested).
7. The trains' progress "demoted by staying still" (p. 870) has no proposition of Table 1 that changes. It is left out; adding it changes no choice (tested).

The earlier statement that none of these differences changes a reported choice was incorrect: it overlooked Fig. 5's magnitude sensitivity. `test_figure5_single_penalty_changes_the_hurried_choice` now records the counterexample. The tests separately establish invariance for the added j0 and +CH effects, the extra q0 repetition penalty, the train standing-still penalty, and Fig. 7's penalty count. All 32 recorded choice expectations (30 narrative-based and two additional Fig. 1 checks) still pass on the declared textual reconstruction. This does not establish equivalence to every literal graphical reading or a unique decision algorithm prescribed by the paper.

## 5. Independent adapter review (2026-11-22 UTC)

A non-author model-assisted review reproduced an adapter defect outside the paper examples: different joint actions with the same source and target states could declare different effects. Because the library's effect facts depend on the state pair, exporting those models merged the effects and created arguments the adapter had not predicted. `tools/aats.py` now rejects that inconsistency, as required by the paper's state-pair δ and this project's goal-labelled refinement. Tests include a hand-authored actual-engine control and an independent join of emitted facts for all paper situations and 300 random models; this join does not read the adapter's predicted argument ids.

The same review removed a non-source restriction: p. 860 does not require π to be injective. Distinct state ids may now share their true propositions. A claim supplied only as a proposition list must use an explicit state id when that list is ambiguous. The five library records and the existing paper models are unchanged by these adapter repairs.

It also corrected the CQ2/CQ13 boundary for an explicitly supplied joint-action tuple. The paper defines J_Ag as the Cartesian product of the agents' action sets (p. 860). A legal tuple without a named or defined transition therefore fails CQ2, rather than CQ13; the input's list of named joint actions does not redefine membership in that product. An unknown component action, such as crossing a bridge outside the model's action set, still fails CQ13.

The root integration review additionally required duplicate-key rejection for model YAML. The adapter now rejects repeated literal keys and malformed collection shapes, while preserving intentional YAML merge-key overrides in both supplied models. Its mixed-sign δ display now matches the documented `+-` order. See `final-atkinson-independent-review.md` for the exact findings and acceptance scope.

A separate final cross-review found and closed three more input boundaries: an unknown joint reference now produces the intended typed model error without a cascading lookup error; a mistyped audience no longer silently recommends all actions; and malformed claimed-argument shapes produce `ModelError`. Valid partial audience orders and values declared for other modeled agents remain supported. These repairs leave the paper examples and five source fingerprints unchanged.

Acceptance is for the documented project interpretation, including its goal decomposition, assurance policy, and comparison contract. AS1 facts alone do not contain enough information to run that contract: a separately supplied AATS model and explicit mapping are required (`docs/contracts.md`, contract 8).
