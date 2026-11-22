#!/usr/bin/env python3
"""Tests of tools/aats.py: the AATS definitions (pp. 860-861), AS2 and its
critical questions (pp. 861-862), the farmer (pp. 863-869) and the trains
(pp. 870-872) of Atkinson and Bench-Capon (2007), and the comparison contract.

Three things are checked against something other than the code under test:
  * the paper: its arguments Arg1-Arg8, its objections, its formulation
    questions and every choice it reports (decisions in the models);
  * a second implementation, written here: the formal critical questions of
    p. 862 as set comprehensions, and each choice recomputed from the facts and
    argument identifiers of the generated Carneades cases alone (which the
    suite then checks the engine builds, exactly);
  * random models: both implementations agree on every one.
Run: python3 tools/test_aats.py
"""
import copy
import random
import re
import sys
import unittest
from pathlib import Path

import yaml

sys.path.insert(0, str(Path(__file__).resolve().parent))
import aats  # noqa: E402

FARMER = aats.MODELS / "farmer.yml"
TRAINS = aats.MODELS / "trains.yml"


def farmer():
    return aats.Model(FARMER)


def situation_from(raw_model, situation):
    """A Model built from a mapping, for mutation tests."""
    path = Path("/nonexistent")
    m = aats.Model.__new__(aats.Model)
    m.path = path
    doc = copy.deepcopy(raw_model)
    m.name, m.source = doc["name"], doc["source"]
    m.agents, m.joint, m.audiences = doc["agents"], doc["joint_actions"], doc.get("audiences", {})
    return aats.Situation(m, situation)


def raw(path):
    return yaml.safe_load(Path(path).read_text(encoding="utf-8"))


def raw_situation(path, name):
    doc = raw(path)
    return doc, copy.deepcopy(next(s for s in doc["situations"] if s["name"] == name))


class Definitions(unittest.TestCase):
    """The AATS of pp. 860-861: what a model must satisfy."""

    def setUp(self):
        self.doc, self.sit = raw_situation(FARMER, "fig2-q0")

    def rejects(self, mutate, text):
        doc, sit = copy.deepcopy(self.doc), copy.deepcopy(self.sit)
        mutate(doc, sit)
        with self.assertRaises(aats.ModelError) as e:
            situation_from(doc, sit)
        self.assertIn(text, str(e.exception))

    def test_control(self):
        s = situation_from(self.doc, self.sit)
        self.assertEqual(s.options(), ["row_alone", "row_chicken", "row_dog", "row_seeds"])

    def test_initial_state_in_Q(self):
        self.rejects(lambda d, s: s.update(initial="q99"), "initial state")

    def test_tau_targets_in_Q(self):
        self.rejects(lambda d, s: s["transitions"].append({"from": "q0", "by": "j5", "to": "q99"}), "both states must be in Q")

    def test_tau_is_a_function(self):
        self.rejects(lambda d, s: s["transitions"].append({"from": "q0", "by": "j9", "to": "q4"}), "tau is a function")

    def test_goal_must_change(self):
        def m(d, s):
            s["transitions"][3]["promotes"] = {"progress": ["d_left"]}
        self.rejects(m, "does not change")

    def test_effect_value_is_a_value(self):
        def m(d, s):
            s["transitions"][3]["promotes"] = {"glory": ["c_right"]}
        self.rejects(m, "is not a value of any agent")

    def test_every_agent_has_values(self):
        def m(d, s):
            s["values"] = {"farmer": ["progress"]}
        self.rejects(m, "every agent needs")

    def test_states_are_told_apart(self):
        def m(d, s):
            s["states"]["q9b"] = list(s["states"]["q9"])
        self.rejects(m, "same interpretation")

    def test_preconditions_agree_with_tau(self):
        def m(d, s):
            s["preconditions"] = {"row_chicken": ["q4"]}
        self.rejects(m, "cannot be performed in q0 (rho)")

    def test_assurance_cannot_remove_an_own_action(self):
        def m(d, s):
            s["assurances"] = ["j9"]
        self.rejects(m, "exclude every joint action")

    def test_disjoint_actions(self):
        doc = copy.deepcopy(raw(FARMER))
        doc["agents"]["dog"].append("chicken_waits")
        p = Path(self._tmp("disjoint.yml"))
        p.write_text(yaml.safe_dump(doc), encoding="utf-8")
        with self.assertRaises(aats.ModelError) as e:
            aats.Model(p)
        self.assertIn("pairwise disjoint", str(e.exception))

    def test_joint_action_has_one_action_per_agent(self):
        doc = copy.deepcopy(raw(FARMER))
        del doc["joint_actions"]["j3"]["dog"]
        p = Path(self._tmp("joint.yml"))
        p.write_text(yaml.safe_dump(doc), encoding="utf-8")
        with self.assertRaises(aats.ModelError) as e:
            aats.Model(p)
        self.assertIn("exactly one of its own actions", str(e.exception))

    def _tmp(self, name):
        import tempfile
        d = Path(tempfile.mkdtemp())
        return d / name


class FarmerFigure1(unittest.TestCase):
    """p. 864, Fig. 1 and pp. 864-866: the arguments and objections the paper
    lists are those the model yields."""

    def setUp(self):
        self.s = farmer().situation("fig1-q0")
        self.args = self.s.arguments()

    def find(self, sign, joint, value):
        hits = [a for a in self.args if (a.sign, a.joint, a.value) == (sign, joint, value)]
        self.assertEqual(len(hits), 1, (sign, joint, value))
        return hits[0]

    def test_positive_arguments_are_arg1_to_arg8(self):
        paper = {("j2", "chicken_happy"),                          # Arg1
                 ("j3", "dog_happy"),                              # Arg2
                 ("j4", "dog_happy"),                              # Arg3
                 ("j5", "progress"),                               # Arg4
                 ("j6", "progress"), ("j6", "dog_happy"),          # Arg5: P and DH
                 ("j7", "progress"), ("j7", "friendship"),         # Arg6: P and F
                 ("j8", "progress"), ("j8", "friendship"), ("j8", "chicken_happy"),  # Arg7
                 ("j9", "progress")}                               # Arg8
        self.assertEqual({(a.joint, a.value) for a in self.args if a.sign == "+"}, paper)

    def cqs(self, arg):
        return [o["cq"] for o in self.s.objections(arg)]

    def test_obj1a_obj1b(self):
        arg1 = self.find("+", "j2", "chicken_happy")
        self.assertTrue(self.s.cq17_unanswered(arg1))                  # Obj1a (CQ17)
        self.assertIn(9, self.cqs(arg1))                               # Obj1b: S demoted in q2

    def test_obj2a_obj2b_obj2bi_obj2d_obj2f(self):
        arg2 = self.find("+", "j3", "dog_happy")
        dev = dict(self.s.deviations(arg2))
        self.assertFalse(dev["j2"])            # Obj2a: q2, DH not promoted, S demoted
        self.assertTrue(dev["j4"])             # Obj2b answered by Obj2bi: q4 still promotes DH
        o17 = {o["joint"]: o for o in self.s.objections(arg2) if o["cq"] == 17}
        self.assertEqual(o17["j2"]["demotes"], ["seeds"])
        self.assertEqual(o17["j4"]["demotes"], ["chicken", "seeds"])
        self.assertIn(9, self.cqs(arg2))                               # Obj2d: C demoted in q3
        o11 = [o for o in self.s.objections(arg2) if o["cq"] == 11]
        self.assertIn(("j5", "progress"), {(o["joint"], o["value"]) for o in o11})   # Obj2f

    def test_obj3a(self):
        self.assertIn(9, self.cqs(self.find("+", "j4", "dog_happy")))

    def test_obj4a_obj4b_obj4c(self):
        arg4 = self.find("+", "j5", "progress")
        o = self.s.objections(arg4)
        o17 = [x for x in o if x["cq"] == 17]
        self.assertEqual([(x["joint"], x["answered"], x["demotes"]) for x in o17], [("j6", True, ["chicken"])])  # Obj4a
        self.assertIn(("j9", "progress"), {(x["joint"], x["value"]) for x in o if x["cq"] == 7})       # Obj4b
        self.assertNotIn("j9", {x["joint"] for x in o if x["cq"] == 6})   # see test_cq6_needs_the_same_proposition
        self.assertIn(("j7", "friendship"), {(x["joint"], x["value"]) for x in o if x["cq"] == 11})    # Obj4c

    def test_obj5a_obj7a_obj6a(self):
        self.assertIn(9, self.cqs(self.find("+", "j6", "progress")))      # Obj5a
        self.assertIn(9, self.cqs(self.find("+", "j8", "progress")))      # Obj7a
        arg6 = self.find("+", "j7", "progress")
        self.assertEqual([(x["joint"], x["demotes"]) for x in self.s.objections(arg6) if x["cq"] == 17],
                         [("j8", ["seeds"])])                              # Obj6a (CQ17, CQ9)

    def test_cq10_arg6(self):
        arg6 = self.find("+", "j7", "progress")
        self.assertIn(("friendship"), {o["value"] for o in self.s.objections(arg6) if o["cq"] == 10})

    def test_cq6_needs_the_same_proposition(self):
        # p. 865 calls Obj4b CQ6, reading "putting a possession on the right
        # bank" as one goal (transportation, p. 864). On the AATS a goal is a
        # proposition: CQ6 holds for another joint action changing s_right.
        doc, sit = raw_situation(FARMER, "fig1-q0")
        sit["states"]["q5b"] = ["c_left", "d_left", "b_right", "s_right", "repeated"]
        sit["transitions"].append({"from": "q0", "by": "j0", "to": "q5b", "promotes": {"progress": ["s_right"]}})
        s = situation_from(doc, sit)
        arg4 = [a for a in s.arguments() if (a.joint, a.value) == ("j5", "progress")][0]
        self.assertIn("j0", {o["joint"] for o in s.objections(arg4) if o["cq"] == 6})


class Formulation(unittest.TestCase):
    """p. 864 (bridge, muzzled dog, q_n with n > 9, another initial state),
    p. 865 (Obj1c, Obj2c, Obj2e), p. 867 (the CQ3 objection to Arg10)."""

    def setUp(self):
        self.m = farmer()
        self.fig1 = self.m.situation("fig1-q0")
        self.arg8 = {"state": "q0", "agent": "farmer", "action": "row_chicken", "joint": "j9",
                     "result": "q9", "goal": "c_right", "value": "progress"}

    def claim(self, sit, **change):
        c = dict(self.arg8)
        c.update(change)
        return sit.claim_problems(c)

    def test_arg8_is_sound(self):
        self.assertEqual(self.claim(self.fig1), [])

    def test_cq13_bridge(self):
        self.assertIn(13, self.claim(self.fig1, joint={"farmer": "cross_bridge", "chicken": "chicken_waits", "dog": "dog_waits"},
                                     action="cross_bridge"))

    def test_cq12_cq14_cq15_muzzled_dog(self):
        muzzled = ["b_left", "c_left", "d_left", "s_left", "d_muzzled"]
        self.assertIn(12, self.claim(self.fig1, state=muzzled))
        self.assertIn(14, self.claim(self.fig1, result=["d_left", "s_left", "b_right", "c_right", "d_muzzled"]))
        self.assertIn(15, self.claim(self.fig1, goal="d_muzzled"))

    def test_cq2_consequence_beyond_q9(self):
        # j9 is claimed to reach q8: the transition (CQ2) and the goal's
        # change (CQ3) are denied; delta is not asked of a transition that is not there
        self.assertEqual(self.claim(self.fig1, result="q8"), [2, 3])

    def test_cq1_other_initial_state(self):
        self.assertIn(1, self.claim(self.fig1, state="q1"))

    def test_cq16_values_dropped_by_the_reformulation(self):
        arg1 = {"state": "q0", "agent": "farmer", "action": "row_alone", "joint": "j2", "result": "q2",
                "goal": "s_eaten", "value": "chicken_happy"}
        self.assertEqual(self.fig1.claim_problems(arg1), [])
        doc, sit = raw_situation(FARMER, "fig2-q0")
        sit["values"]["chicken"] = ["chicken_happy"]
        sit["values"] = {"farmer": sit["values"]["farmer"], "chicken": ["hen_comfort"], "dog": ["dog_comfort"]}
        sit["states"]["q2"] = ["c_left", "d_left", "b_right", "s_eaten"]
        sit["transitions"].append({"from": "q0", "by": "j2", "to": "q2", "demotes": {"seeds": ["s_eaten"]}})
        reformulated = situation_from(doc, sit)
        self.assertIn(16, reformulated.claim_problems(arg1))          # Obj1c
        self.assertIn(4, reformulated.claim_problems(arg1))

    def test_cq4_obj2c(self):
        arg2 = {"state": "q0", "agent": "farmer", "action": "row_alone", "joint": "j3", "result": "q3",
                "goal": "c_eaten", "value": "dog_happy"}
        self.assertEqual(self.fig1.claim_problems(arg2), [])
        doc, sit = raw_situation(FARMER, "fig1-q0")
        t = next(t for t in sit["transitions"] if t["by"] == "j3")
        del t["promotes"]                                              # eating the chicken does not promote DH
        self.assertEqual(situation_from(doc, sit).claim_problems(arg2), [4])

    def test_cq3_arg10_and_repetition(self):
        arg10 = {"state": "q9", "agent": "farmer", "action": "farmer_waits", "joint": "j0", "result": "q9_again",
                 "goal": "repeated", "value": "progress"}
        fig3 = self.m.situation("fig3-q9")
        self.assertNotIn(3, fig3.claim_problems(arg10))               # non-movement counted as repetition
        doc, sit = raw_situation(FARMER, "fig3-q9")
        sit["transitions"][0] = {"from": "q9", "by": "j0", "to": "q9"}  # staying in q9 changes nothing
        literal = situation_from(doc, sit)
        self.assertIn(3, literal.claim_problems(dict(arg10, result="q9")))


class Formal(unittest.TestCase):
    """The choice-of-action questions of p. 862 written again, as set
    comprehensions over tau and delta, compared with Situation.objections on
    the paper's models and on random ones."""

    @staticmethod
    def formal(s, a):
        qx, jn, qy, pa, vu, i = a.state, a.joint, a.result, a.goal, a.value, a.agent
        J = set(s.possible())
        tau = {j: s.tau[(qx, j)] for j in J}
        eff = {j: set(s.effects[(qx, j)]) for j in J}
        pi = s.states
        out = set()
        out |= {(5, jm) for jm in J - {jn} if tau[jm] == qy}
        out |= {(6, jm) for jm in J - {jn} if pa in pi[qx] ^ pi[tau[jm]]}
        out |= {(7, jm) for jm in J - {jn} if any(e.sign == "+" and e.value == vu for e in eff[jm])}
        out |= {(8, e.goal) for e in eff[jn] if e.sign == "-" and e.value == vu and e.goal != pa}
        out |= {(9, e.value) for e in eff[jn] if e.sign == "-" and e.value != vu}
        out |= {(10, e.value) for e in eff[jn] if e.sign == "+" and e.value != vu}
        out |= {(11, jm) for jm in J - {jn} if any(e.sign == "+" and e.value != vu for e in eff[jm])}
        out |= {(17, jm) for jm in J if s.joint[jm][i] == s.joint[jn][i] and jm != jn and tau[jm] != tau[jn]}
        return out

    @staticmethod
    def implemented(s, a):
        key = {5: "joint", 6: "joint", 7: "joint", 11: "joint", 17: "joint", 8: "goal", 9: "value", 10: "value"}
        return {(o["cq"], o[key[o["cq"]]]) for o in s.objections(a)}

    def check(self, s):
        for a in s.arguments():
            if a.sign == "+":
                self.assertEqual(self.implemented(s, a), self.formal(s, a), (s.name, a))

    def test_paper_models(self):
        for path in (FARMER, TRAINS):
            for s in aats.Model(path).situations:
                self.check(s)

    def test_random_models(self):
        for seed in range(300):
            self.check(random_situation(seed))

    def test_cq5_same_state_by_another_joint_action(self):
        # neither paper example has two joint actions with the same result
        # from the current state; constructed ones do
        hits = 0
        for seed in range(60):
            s = random_situation(seed, force_cq5=True)
            for a in s.arguments():
                if a.sign == "+" and a.joint in ("j0", "j1"):
                    other = "j1" if a.joint == "j0" else "j0"
                    if other in s.possible():
                        self.assertIn((5, other), self.implemented(s, a))
                        hits += 1
        self.assertGreater(hits, 10)

    def test_cq15_unrealizable_goal(self):
        s = farmer().situation("fig2-q0")
        claim = {"state": "q0", "agent": "farmer", "action": "row_chicken", "joint": "j9", "result": "q9",
                 "goal": "all_home", "value": "progress"}
        self.assertIn(15, s.claim_problems(claim))


def random_situation(seed, force_cq5=False):
    """A random AATS with two agents, each with two actions, some states and
    random promotions and demotions attached to propositions that change."""
    rnd = random.Random(seed)
    props = ["p", "q", "r", "t"]
    n = rnd.randint(2, 5)
    states = {}
    while len(states) < n:
        s = frozenset(x for x in props if rnd.random() < 0.5)
        if s not in states.values():
            states[f"s{len(states)}"] = s
    names = list(states)
    agents = {"a": ["a1", "a2"], "b": ["b1", "b2"]}
    joint = {f"j{k}": {"a": x, "b": y} for k, (x, y) in enumerate([(x, y) for x in agents["a"] for y in agents["b"]])}
    values = {"a": ["u", "v", "w"], "b": ["w"]}
    transitions = []
    for j in joint:
        if rnd.random() < 0.85 or j == "j0" or (force_cq5 and j == "j1"):
            to = rnd.choice(names)
            if force_cq5 and j in ("j0", "j1"):
                to = names[1]
            changed = sorted(states["s0"] ^ states[to])
            t = {"from": "s0", "by": j, "to": to}
            for field in ("promotes", "demotes"):
                for v in ("u", "v", "w"):
                    goals = [g for g in changed if rnd.random() < 0.35]
                    if goals:
                        t.setdefault(field, {})[v] = goals
            transitions.append(t)
    assurances = []
    for j in joint:
        if rnd.random() < 0.15 and j != "j0" and any(t["by"] == j for t in transitions):
            others = [t["by"] for t in transitions if t["by"] != j and joint[t["by"]]["a"] == joint[j]["a"]
                      and t["by"] not in assurances]
            if others:
                assurances.append(j)
    doc = {"name": f"random{seed}", "source": "random", "agents": agents, "joint_actions": joint}
    sit = {"name": f"r{seed}", "source": "random", "agent": "a", "values": values, "initial": "s0",
           "states": {k: sorted(v) for k, v in states.items()}, "transitions": transitions, "assurances": assurances}
    return situation_from(doc, sit)


def decide_from_case(case, audience, agent):
    """The comparison contract recomputed from a generated Carneades case
    alone: the facts say which joint actions the agent may face for each of
    its actions (joint_action_part, possible_joint_action, transition); the
    argument identifiers say which goals promote or demote which value."""
    facts = case["assumptions"]
    q0 = next(re.match(r"current_state\((\w+)\)", f).group(1) for f in facts if f.startswith("current_state("))
    possible = {re.match(r"possible_joint_action\((\w+),", f).group(1) for f in facts if f.startswith("possible_joint_action(")}
    part = {}
    for f in facts:
        m = re.match(r"joint_action_part\((\w+),(\w+),(\w+)\)", f)
        if m and m.group(2) == agent:
            part[m.group(1)] = m.group(3)
    result = {}
    for f in facts:
        m = re.match(r"transition\((\w+),(\w+),(\w+)\)", f)
        if m and m.group(1) == q0 and m.group(2) in possible:
            result[m.group(2)] = m.group(3)
    net = {}
    for a in case["arguments"]:
        m = re.match(r"(practical_reasoning_as2(?:_demotion)?)\((\w+),(\w+),(\w+),(\w+),(\w+),(\w+),(\w+)\)", a)
        sign = -1 if m.group(1).endswith("_demotion") else 1
        net[(m.group(4), m.group(8))] = net.get((m.group(4), m.group(8)), 0) + sign
    own_values = [re.match(r"agent_value\(\w+,(\w+)\)", f).group(1) for f in facts if f.startswith(f"agent_value({agent},")]
    vals = [v for v in audience if v in own_values]
    options = {}
    for j in result:
        options.setdefault(part[j], []).append(j)
    scores = {a: tuple(min(net.get((j, v), 0) for j in js) for v in vals) for a, js in options.items()}
    best = max(scores.values())
    return sorted(a for a, s in scores.items() if s == best)


class Contract(unittest.TestCase):
    """The comparison contract value-based-choice."""

    def test_every_reported_choice(self):
        for path in (FARMER, TRAINS):
            self.assertEqual(aats.decision_problems(aats.Model(path)), [])

    def test_choices_do_not_depend_on_ranking_chicken_against_seeds(self):
        m = farmer()
        for s in m.situations:
            self.assertEqual(s.decide(m.audience("in_a_hurry")), s.decide(m.audience("seeds_first")), s.name)
            self.assertEqual(s.decide(m.audience("companionable")), s.decide(m.audience("companionable_seeds_first")), s.name)

    def test_second_derivation_from_the_engine_cases(self):
        cases = {c["name"]: c for c in yaml.safe_load(aats.CASES.read_text(encoding="utf-8"))["cases"]}
        for path in (FARMER, TRAINS):
            m = aats.Model(path)
            for s in m.situations:
                case = cases[f"practical_reasoning_as2/{m.name}-{s.name}"]
                for name in m.audiences:
                    self.assertEqual(decide_from_case(case, m.audience(name), s.agent), s.decide(m.audience(name)),
                                     (s.name, name))

    def test_second_derivation_on_random_models(self):
        for seed in range(300):
            s = random_situation(seed)
            case = yaml.safe_load(_cases_for(s))["cases"][0]
            for order in (["u", "v", "w"], ["w", "v", "u"], ["v", "u"]):
                self.assertEqual(decide_from_case(case, order, "a"), s.decide(order), (seed, order))

    def test_revisit_of_q0_does_not_change_the_choice(self):
        # p. 867 counts only "putting the chicken on the left bank" against
        # rowing the chicken back to q0, although q0 is visited again
        doc, sit = raw_situation(FARMER, "fig3-q9")
        sit["states"]["q0"] = sit["states"]["q0"] + ["repeated"]
        sit["transitions"][2]["demotes"] = {"progress": ["c_right", "repeated"]}
        s = situation_from(doc, sit)
        m = farmer()
        for name in m.audiences:
            self.assertEqual(s.decide(m.audience(name)), ["row_alone"])

    def test_standing_still_demoting_progress_changes_no_train_choice(self):
        m = aats.Model(TRAINS)
        for s in m.situations:
            before = {name: s.decide(m.audience(name)) for name in m.audiences}
            s2 = copy.copy(s)
            s2.effects = {k: list(v) for k, v in s.effects.items()}
            own = "progress_" + s.agent
            for (q, j), effs in s2.effects.items():
                if s.joint[j][s.agent].endswith("_stays"):
                    effs.append(aats.Effect("-", own, None))        # no goal: outside AS2, inside the comparison
            after = {name: s2.decide(m.audience(name)) for name in m.audiences}
            self.assertEqual(before, after, s.name)

    def test_ties_return_every_best_option(self):
        doc, sit = raw_situation(FARMER, "fig8-q16")
        sit["transitions"][0] = {"from": "q16", "by": "j1", "to": "q15_again", "promotes": {"progress": ["repeated"]}}
        s = situation_from(doc, sit)
        self.assertEqual(s.decide(["progress"]), ["row_alone", "row_chicken"])

    def test_values_outside_the_audience_are_ignored(self):
        s = farmer().situation("fig2-q0")
        # an audience that holds only the dog's happiness, which the farmer no
        # longer counts among his values (Obj2e accepted, p. 866): nothing of
        # his is left to compare, and every option is as good as any other
        self.assertEqual(s.decide(["dog_happy"]), ["row_alone", "row_chicken", "row_dog", "row_seeds"])


def _cases_for(s):
    m = type("M", (), {})()
    m.name, m.source, m.situations = "random", "random", [s]
    return aats.build_cases([m])


class Files(unittest.TestCase):
    def test_generated_cases_are_current(self):
        models = [aats.Model(p) for p in aats.model_paths()]
        self.assertEqual(aats.CASES.read_text(encoding="utf-8"), aats.build_cases(models),
                         "run python3 tools/aats.py cases --write")

    def test_check_command(self):
        self.assertEqual(aats.main(["check"]), 0)

    def test_engine_labels_follow_the_arguments(self):
        s = farmer().situation("fig1-q0")
        labels = aats.engine_labels(s)
        self.assertEqual(labels["should_perform(farmer,row_chicken)"], "in")
        self.assertEqual(labels["¬should_perform(farmer,row_alone)"], "in")      # every pro argument undercut by CQ17
        self.assertEqual(labels["should_perform(farmer,row_alone)"], "out")
        self.assertEqual(labels["should_perform(farmer,row_seeds)"], "out")      # P against C, both weigh 1
        self.assertEqual(labels["¬should_perform(farmer,row_seeds)"], "out")


if __name__ == "__main__":
    unittest.main(verbosity=1)
