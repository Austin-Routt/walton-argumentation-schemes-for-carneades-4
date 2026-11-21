# Closing concrete context defects after PR #15

## Basis and review boundary

This patch changes ten previously held records after comparing their formulas
with Walton, Reed and Macagno (2008), printed pp.321,328–330,337–338,342 and345.
The supplied PDF is the source. These are executable identity defects, not
merely a preference to hold an otherwise unrestricted interpretation on the
safe side. The source describes a particular agent, inquiry, argument,
institution, language, observation or plea; the former atoms sometimes
discarded that identity or allowed a challenge to spread across applications.

The patch author supplies the proposed interpretation and cases. A different
reviewer must accept the source decisions and examples before promotion.
This document does not grant approval to its own changes.

## Exact decisions

| Record | Source | Repair and retained interpretation |
|---|---|---|
| `positive_value` | 19.1,p.321 | Carry the evaluating agent into `reason_retaining_goal(A,V,G)`. A reason concerns this agent's commitment, not every owner's identically named goal. |
| `negative_value` | 19.2,p.321 | Carry the agent into `reason_retracting_goal(A,V,G)`. Retain the corrected negative-value polarity. |
| `sign` | 30,p.329 | Scope the competing explanation to both sign and candidate: `more_reliable(A,B)`. A better explanation relative to one candidate does not automatically defeat another candidate. |
| `cause_to_effect` | 28,p.328 | Scope interference to the event/effect application: `interference(E1,E2)`. The event terms must identify the application, e.g. `lightning(storm1)` and `fire(storm1)`. The selected affirmative occurrence branch and application-assessed causal premise remain explicit. |
| `abductive_backward` | 31.1,pp.329–330 | Require `member_of_accounts(A1,A)`, compare the selected member within `successful_account(A1,A,D)`, and retain the candidate universe and case data in `plausible_hypothesis(A1,A,D)`. Membership cannot be invented from a best-account claim in another universe. |
| `circumstantial_ad_hominem` | 45,pp.337–338 | Scope questioned credibility and character relevance to `(A,S)`, the person and attacked argument. Retain the weaker source conclusion of decreased plausibility; no automatic rejection or destruction is added. |
| `rule_physical_world_2` | 54.2,p.342 | Preserve institution `I` in `counts_as(I,D,W)`. `correctly_report` remains the source's explicitly universal warrant about official rulebooks; it is not a factual assertion about only one book. Rule-change counterevidence remains specific to `(I,D,W)`. |
| `rule_mental_world` | 54.3,p.342 | Use `means_same_in_language(L,W,D)` and `without_correcting(L,W,D)`. Keep the source's universal `speakers_correct` warrant. This linguistic conclusion is separate from 54.1's referential-coextension conclusion. |
| `perception` | 59.1,p.345 | Scope unreliability to `not_indicator(P,I)`. The proposition `I` identifies the circumstances, and observer `P` identifies the perception. Independent perception and memory may still support the shared `reasonable_believe(I)`. |
| `plea_for_excuse` | 58,p.345 | Add pleader `P` and explicit case `C`, while preserving rule `R`, action type `T`, sanction `S` and excuse `E`. The conclusion is `exempted(P,C,R,T,S)`; justification and problematic-precedent challenges name that same plea. An action type is not overloaded to stand for all its occurrences. |

For the plea, the `excuse` assumption expresses the source's conditional
justification: recognized-category membership must be justified in this case,
or the special features must justify an unrecognized excuse. The
`not_excuse` exception means an unrecognized excuse would set a problematic
precedent; mere lack of recognition does not establish that exception.
Recognition, justification and precedent are application evidence, not facts
calculated by the engine. Questioning the justification suspends its default;
an explicit accepted answer can restore it. Evidence for one case does not
answer another case's question.

## Executable evidence

`tools/vv/cases/closure_context.yml` contains 31 manually specified cases,
covering positive controls, wrong-agent or wrong-application controls,
independent support, supported exceptions, and unanswered/answered default
questions. The harness runs the cases through both API and CLI paths. The
existing book and phase cases retain their intended scenarios, now with the
missing identities explicit. Their expected outputs were examined against
the source; they were not generated from the code under test.

The cases distinguish:

- Opposite values held by different agents for the same named goal.
- Two sign interpretations with a competitor against only one candidate.
- Two effects of one event with interference against only one effect.
- Membership and comparative success in different hypothesis universes.
- Credibility and relevance challenges to different arguments by one speaker.
- Different constitutive rules for the same position in two institutions.
- The same word with different definitions in two languages.
- Independent observers of the same proposition and one observer in two cases.
- Two pleaders in separate cases and one pleader in two occurrences.

## Compatibility

The migration table describes every changed atom and the added membership
premise. These are manual semantic migrations: no program can infer the
missing person, institution, case or candidate set from a formerly global
atom. The historical `succesful_account/2` spelling change also becomes a
manual migration because the corrected form requires the candidate set.
Existing authored argument parameters for `plea_for_excuse` must be updated
to `[P,C,S,R,T,E]`; a missing participant or case is not silently guessed.

Do not infer approval of other held records from these repairs. In particular,
language-specific identity does not implement a dialogue protocol, and
explicit event terms do not calculate causal probability or validate the
user's real-world evidence.
