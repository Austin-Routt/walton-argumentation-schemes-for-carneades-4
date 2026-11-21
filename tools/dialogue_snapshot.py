#!/usr/bin/env python3
"""Build a fresh, finite dialogue snapshot from an evaluated Carneades graph.

A question adds supported evidence for the target argument's undercutter.
An explicitly accepted answer closes that question; it does not assert the
argument's conclusion or erase any original undercutter. Every snapshot is
rebuilt from the same unmodified base. See docs/dialogue-snapshots.md.
"""
import argparse
import copy
import hashlib
import json
from pathlib import Path
import re
import sys

import yaml
sys.path.insert(0, str(Path(__file__).resolve().parent))
from validate_schemes import load_unique, parse_statement

SCHEMA = 'walton-dialogue-snapshot/v1'
# These are question inventories, not mechanically generated truth judgements.
CQ_COUNTS = {'practical_inference': 5, 'practical_necessary_condition': 5,
             'practical_sufficient_condition': 5, 'value_based_practical_reasoning': 7,
             'two_person_practical': 5, 'position_to_know': 3}
IGNORANCE = {'ignorance', 'negative_normal_expectations', 'negative_practical',
             'negative_practical_true'}
ID = re.compile(r'[a-z][a-z0-9_]*\Z')

class DialogueError(ValueError):
    pass

def require(ok, message):
    if not ok:
        raise DialogueError(message)

def mapping(value, label):
    require(isinstance(value, dict), label + ' must be a mapping')
    return value

def token(value, label):
    require(isinstance(value, str) and ID.fullmatch(value), label + ' must be a lowercase identifier')
    return value

def text(value, label):
    require(isinstance(value, str) and bool(value.strip()), label + ' must be nonempty text')
    return value

def keys(obj, required, optional, label):
    mapping(obj, label)
    require(set(required) <= set(obj), label + ': missing ' + ', '.join(sorted(set(required)-set(obj))))
    require(set(obj) <= set(required)|set(optional), label + ': unknown fields ' + ', '.join(sorted(set(obj)-set(required)-set(optional))))

def ground(value, label):
    text(value, label)
    normalized, masked, i, literal = [], [], 0, 0
    while i < len(value):
        ch = value[i]
        if ch.isspace():
            masked.append(' '); i += 1; continue
        if ch in ("'", '"'):
            quote, start = ch, i
            i += 1
            while i < len(value):
                if value[i] == "\\":
                    i += 2
                elif value[i] == quote:
                    i += 1; break
                else:
                    require(value[i] not in '\r\n', label + ': unescaped newline in quoted term')
                    i += 1
            else:
                raise DialogueError(label + ': unterminated quoted term')
            require(i <= len(value), label + ': incomplete escape in quoted term')
            normalized.append(value[start:i])
            masked.append(' quoted_literal_' + str(literal) + ' '); literal += 1
        else:
            normalized.append(ch); masked.append(ch); i += 1
    try:
        _, _, _, variables = parse_statement('parameter(' + ''.join(masked) + ')')
    except ValueError as exc:
        raise DialogueError(label + ': invalid supported ground term: ' + str(exc)) from exc
    require(not variables, label + ': free variables are not allowed')
    return ''.join(normalized)

def parameters(value, label):
    require(isinstance(value, list), label + ' must be a list')
    return [ground(x, label) for x in value]

def read(path):
    obj, duplicates = load_unique(Path(path).read_text(encoding='utf-8'))
    require(not duplicates, str(path) + ': duplicate YAML keys')
    return mapping(obj, str(path))

def approved_in_base(base, sid):
    if sid in base.get('assumptions', []):
        return True
    statement = base['statements'].get(sid)
    return isinstance(statement, dict) and statement.get('label') == 'in'

def bind_target(base, target, label, supported):
    keys(target, {'argument', 'scheme', 'parameters'}, set(), label)
    aid = text(target['argument'], label + '.argument')
    require(aid in base['arguments'], label + ': unknown argument')
    arg = mapping(base['arguments'][aid], 'argument ' + aid)
    sid = text(target['scheme'], label + '.scheme')
    require(sid in supported, label + ': unsupported scheme')
    require(arg.get('scheme') == sid, label + ': scheme does not match the actual argument')
    ps = parameters(target['parameters'], label + '.parameters')
    definitions = [s for s in base.get('argument_schemes', []) if isinstance(s, dict) and s.get('id') == sid]
    require(len(definitions) == 1 and isinstance(definitions[0].get('variables'), list) and
            len(definitions[0]['variables']) == len(ps), label + ': missing or ambiguous schema, or wrong parameter arity')
    require(parameters(arg.get('parameters', []), 'actual parameters') == ps,
            label + ': parameters do not match the actual argument')
    require(isinstance(arg.get('premises'), (list, dict)) and bool(arg['premises']) and
            arg.get('conclusion') in base['statements'],
            label + ': target must be a constructed argument, not an unexpanded request')
    return aid

def depends_on_target(base, statement, target_arg):
    """Conservative structural dependency, including issue competition.

    An independent stipulated fact is admissible. For derived answers, reject
    any path from the challenged claim/undercutter, even if another proof might
    also exist. This avoids treating a frozen old label as independent evidence.
    """
    seeds = {target_arg['conclusion'], target_arg.get('undercutter')}
    if statement in seeds:
        return True
    if statement in base.get('assumptions', []):
        return False
    edges = {}
    for arg in base['arguments'].values():
        if not isinstance(arg, dict):
            continue
        ps = arg.get('premises', [])
        ps = list(ps.values()) if isinstance(ps, dict) else list(ps or [])
        if arg.get('undercutter'):
            ps.append(arg['undercutter'])
        for premise in ps:
            edges.setdefault(premise, set()).add(arg.get('conclusion'))
    for issue in (base.get('issues') or {}).values():
        if isinstance(issue, dict):
            positions = issue.get('positions', [])
            for position in positions:
                edges.setdefault(position, set()).update(positions)
    reached, pending = set(seeds), list(seeds)
    while pending:
        for child in edges.get(pending.pop(), set()) - reached:
            reached.add(child); pending.append(child)
    return statement in reached

def check_burdens(base, journal):
    """Validate allocation separately from truth or the strength of evidence."""
    expected = {aid for aid, a in base['arguments'].items()
                if isinstance(a, dict) and a.get('scheme') in IGNORANCE}
    rows = journal.get('burdens', [])
    require(isinstance(rows, list), 'burdens must be a list')
    seen = set()
    for n, row in enumerate(rows):
        label = 'burdens[' + str(n) + ']'
        keys(row, {'argument', 'scheme', 'parameters', 'dialogue', 'probandum',
                   'burden_holder', 'issue', 'standard'}, set(), label)
        require(row['dialogue'] == journal['dialogue'], label + ': foreign dialogue')
        aid = bind_target(base, {k: row[k] for k in ('argument','scheme','parameters')}, label, IGNORANCE)
        require(aid not in seen, label + ': duplicate burden allocation')
        seen.add(aid)
        text(row['burden_holder'], label + '.burden_holder')
        p = ground(row['probandum'], label + '.probandum')
        require(p in base['statements'], label + ': unknown probandum')
        issue = (base.get('issues') or {}).get(row['issue'])
        require(isinstance(issue, dict) and p in issue.get('positions', []),
                label + ': probandum must be a position of the named issue')
        require(row['standard'] in {'PE', 'CCE', 'BRD'}, label + ': unknown proof standard')
        require(issue.get('standard', 'PE') == row['standard'], label + ': proof standard differs from the actual issue')
    require(seen == expected, 'every constructed ignorance-family argument needs its own burden/probandum/standard allocation; missing ' + ', '.join(sorted(expected-seen)))
    return rows

def snapshot(base, journal, through=None):
    mapping(base, 'base graph'); mapping(journal, 'journal')
    mapping(base.get('arguments'), 'base arguments'); mapping(base.get('statements'), 'base statements')
    mapping(base.get('meta', {}), 'base metadata')
    mapping({} if base.get('issues') is None else base['issues'], 'base issues')
    require(isinstance(base.get('assumptions', []), list), 'base assumptions must be a list')
    require(isinstance(base.get('argument_schemes', []), list), 'base argument_schemes must be a list')
    require(not (base.get('meta') or {}).get('dialogue_snapshot'), 'use the unmodified base, not an earlier snapshot')
    keys(journal, {'schema','dialogue','targets','events'}, {'evidence','burdens'}, 'journal')
    require(journal['schema'] == SCHEMA, 'unknown journal schema')
    dialogue = token(journal['dialogue'], 'dialogue')
    ts = mapping(journal['targets'], 'targets')
    targets, used_arguments = {}, set()
    for tid, target in ts.items():
        token(tid, 'target id')
        aid = bind_target(base, target, 'target '+tid, CQ_COUNTS)
        require(aid not in used_arguments, 'duplicate declaration of the same target argument')
        used_arguments.add(aid); targets[tid] = aid
        uc = base['arguments'][aid].get('undercutter')
        if uc is not None:
            require(isinstance(uc, str) and uc in base['statements'], 'existing undercutter must name an existing statement')
            require(sum(a.get('undercutter') == uc for a in base['arguments'].values() if isinstance(a, dict)) == 1,
                    'target undercutter is shared by another argument; exact-target suspension would be unsound')
            require(not any(isinstance(i, dict) and uc in i.get('positions', [])
                            for i in (base.get('issues') or {}).values()),
                    'target undercutter belongs to an issue; this adapter cannot guarantee its question-suspension policy')
    burdens = check_burdens(base, journal)
    evidence = mapping(journal.get('evidence', {}), 'evidence')
    for eid in evidence: token(eid, 'evidence id')
    events = journal['events']; require(isinstance(events, list), 'events must be a list')
    ids, questions, asked, consumed = set(), {}, set(), set()
    selected = {} if through == 'initial' else None
    selected_index = 0 if through == 'initial' else None
    for index, event in enumerate(events, 1):
        mapping(event, 'event'); kind = event.get('type')
        if kind == 'ask': needed = {'id','type','target','cq'}
        elif kind == 'answer': needed = {'id','type','question','evidence'}
        elif kind == 'reopen': needed = {'id','type','question','reason'}
        else: raise DialogueError('unknown event type')
        keys(event, needed, set(), 'event')
        eid = token(event['id'], 'event id')
        require(eid not in ids and eid != 'initial', 'duplicate or reserved event id')
        ids.add(eid)
        if kind == 'ask':
            tid = event['target']; require(tid in targets, 'question names an unknown target')
            cq = event['cq']; sid = ts[tid]['scheme']
            require(type(cq) is int and 1 <= cq <= CQ_COUNTS[sid], 'CQ number outside the declared source inventory')
            require((tid,cq) not in asked, 'duplicate question: use explicit reopen after an answer')
            asked.add((tid,cq)); questions[eid] = {'target':tid,'cq':cq,'open':True}
        else:
            qid = event['question']; require(qid in questions, 'answer/reopen references an unknown or future question')
            q = questions[qid]
            if kind == 'reopen':
                require(not q['open'], 'cannot reopen an already open question')
                text(event['reason'], 'reopen reason'); q['open'] = True
            else:
                require(q['open'], 'cannot answer a closed question')
                refs = event['evidence']
                require(isinstance(refs, list) and refs and all(isinstance(x,str) for x in refs) and len(set(refs)) == len(refs),
                        'an answer needs nonduplicate accepted evidence references')
                for ref in refs:
                    require(ref in evidence and ref not in consumed, 'missing, foreign or reused answer evidence reference')
                    e = evidence[ref]
                    keys(e, {'dialogue','target','question','answer','statement','accepted','accepted_by'}, set(), 'evidence '+ref)
                    require((e['dialogue'],e['target'],e['question'],e['answer']) == (dialogue,q['target'],qid,eid),
                            'evidence belongs to another dialogue, target, question or answer')
                    require(e['accepted'] is True, 'evidence adequacy has not been explicitly accepted')
                    text(e['accepted_by'], 'accepted_by')
                    statement = ground(e['statement'], 'evidence statement')
                    require(statement in base['statements'] and approved_in_base(base,statement),
                            'answer evidence must be accepted in the supplied evaluated base graph')
                    target_arg = base['arguments'][targets[q['target']]]
                    require(not depends_on_target(base, statement, target_arg),
                            'answer evidence depends on the challenged conclusion/undercutter; supply independent accepted evidence')
                    consumed.add(ref)
                q['open'] = False
        if through == eid:
            selected = copy.deepcopy(questions); selected_index = index
    require(consumed == set(evidence), 'unused or unbound evidence declarations')
    if through is None:
        selected = questions; selected_index = len(events)
    require(selected is not None, 'unknown snapshot event')
    out = copy.deepcopy(base)
    # Stale expected outputs and calculated statement labels are not evidence.
    out.pop('tests',None); out.pop('labels',None)
    for value in out['statements'].values():
        if isinstance(value,dict):
            value.pop('label',None); value.pop('value',None)
    out.setdefault('assumptions',[])
    require(isinstance(out['assumptions'],list), 'base assumptions must be a list')
    def fresh(prefix, value, container):
        name = prefix + hashlib.sha256(value.encode('utf-8')).hexdigest()
        require(name not in container, 'generated dialogue identifier collides with base graph')
        return name
    for qid, q in sorted(selected.items()):
        if not q['open']: continue
        aid = targets[q['target']]; arg = out['arguments'][aid]
        uc = arg.get('undercutter')
        if uc is None:
            uc = fresh('dialogue_uc_', dialogue+'\0'+aid, out['statements'])
            out['statements'][uc] = 'An unanswered appropriate question suspends argument '+aid+'.'
            arg['undercutter'] = uc
        marker = fresh('dialogue_question_', dialogue+'\0'+qid, out['statements'])
        out['statements'][marker] = 'Question '+qid+' (CQ'+str(q['cq'])+') remains open for target '+q['target']+' in dialogue '+dialogue+'.'
        out['assumptions'].append(marker)
        support = fresh('dialogue_attack_', dialogue+'\0'+qid, out['arguments'])
        out['arguments'][support] = {'scheme':'linked','premises':[marker],'conclusion':uc}
    base_hash = hashlib.sha256(json.dumps(base,ensure_ascii=False,sort_keys=True,separators=(',',':'),allow_nan=False).encode('utf-8')).hexdigest()
    out.setdefault('meta',{})['dialogue_snapshot'] = {'schema':SCHEMA,'dialogue':dialogue,'base_sha256':base_hash,
        'events_applied':selected_index,'open_questions':[k for k,v in sorted(selected.items()) if v['open']],
        'burden_allocations':copy.deepcopy(burdens),'requires_fresh_evaluation':True}
    return out

def main(argv=None):
    parser=argparse.ArgumentParser(description=__doc__)
    parser.add_argument('base',type=Path); parser.add_argument('journal',type=Path)
    parser.add_argument('--through',help='event id, or initial; all journal entries are validated')
    args=parser.parse_args(argv)
    try:
        result=snapshot(read(args.base),read(args.journal),args.through)
    except (OSError,ValueError,TypeError,yaml.YAMLError) as exc:
        print('dialogue snapshot: '+str(exc),file=sys.stderr);return 2
    sys.stdout.write(yaml.safe_dump(result,sort_keys=False,allow_unicode=True,width=120));return 0

if __name__=='__main__':
    raise SystemExit(main())
