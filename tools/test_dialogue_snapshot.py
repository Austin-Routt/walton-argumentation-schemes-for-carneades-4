"""Real-engine snapshot transitions plus strict journal and burden validation."""
import copy
import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

import yaml
ROOT=Path(__file__).resolve().parent.parent
sys.path.insert(0,str(ROOT/'tools'))
from dialogue_snapshot import snapshot, DialogueError, SCHEMA, read, CQ_COUNTS, IGNORANCE, ground
from graph_from_library import graph

class DialogueSnapshotTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.tmp=tempfile.TemporaryDirectory(prefix='walton-dialogue-')
        cls.work=Path(cls.tmp.name)
        configured=os.environ.get('WALTON_ENGINE_BIN')
        cls.engine=Path(configured) if configured else cls.work/'carneades'
        if not configured:
            env=dict(os.environ,GOTOOLCHAIN='local',GOFLAGS='-mod=mod')
            subprocess.run(['go','build','-o',str(cls.engine),'github.com/carneades/carneades-4/src/cmd/carneades'],
                cwd=ROOT/'tools/vv',env=env,check=True,capture_output=True,text=True)
        cls.counter=0
    @classmethod
    def tearDownClass(cls): cls.tmp.cleanup()
    @classmethod
    def evaluate(cls,value):
        cls.counter+=1; path=cls.work/('graph'+str(cls.counter)+'.yml')
        path.write_text(yaml.safe_dump(value,sort_keys=False,allow_unicode=True))
        run=subprocess.run([str(cls.engine),'eval','-t','yaml',str(path)],check=True,capture_output=True,text=True)
        return yaml.safe_load(run.stdout)
    def base(self, extra=()):
        data=yaml.safe_load(graph(['practical_inference'],str(ROOT/'walton.yml')))
        data['statements']={'adequate_answer':'Independent accepted evidence supplied for the answer.'}
        data['assumptions']=['goal_of_agent(ann,health)','means_to_end(exercise,health)','adequate_answer']+list(extra)
        return self.evaluate(data)
    def target(self,base,scheme='practical_inference',parameters=None):
        parameters=parameters or ['exercise','health','ann']
        found=[(aid,a) for aid,a in base['arguments'].items() if a.get('scheme')==scheme and a.get('parameters')==parameters]
        self.assertEqual(len(found),1)
        return {'argument':found[0][0],'scheme':scheme,'parameters':parameters}
    def journal(self,base,cq=1):
        return {'schema':SCHEMA,'dialogue':'discussion1','targets':{'target1':self.target(base)},
                'events':[{'id':'q1','type':'ask','target':'target1','cq':cq},
                          {'id':'answer1','type':'answer','question':'q1','evidence':['proof1']},
                          {'id':'reopen1','type':'reopen','question':'q1','reason':'The adequacy assessment is challenged again.'}],
                'evidence':{'proof1':{'dialogue':'discussion1','target':'target1','question':'q1','answer':'answer1',
                    'statement':'adequate_answer','accepted':True,'accepted_by':'Independent application reviewer'}}}
    def label(self,g,s):
        v=g['statements'].get(s);return v.get('label','undecided') if isinstance(v,dict) else 'undecided'
    def test_each_practical_question_suspends_restores_and_reopens(self):
        base=self.base();before=copy.deepcopy(base);self.assertEqual(self.label(base,'should_do(ann,exercise)'),'in')
        for cq in range(1,6):
            with self.subTest(cq=cq):
                j=self.journal(base,cq)
                for through,want in [('initial',True),('q1',False),('answer1',True),('reopen1',False)]:
                    made=snapshot(base,j,through);out=self.evaluate(made)
                    self.assertEqual(self.label(out,'should_do(ann,exercise)')=='in',want)
                    twice=self.evaluate(out)
                    self.assertEqual(self.label(twice,'should_do(ann,exercise)')=='in',want)
        self.assertEqual(base,before)
    def test_each_supported_family_and_its_declared_questions(self):
        cases=yaml.safe_load((ROOT/'tools/vv/cases/book.yml').read_text())['cases']
        for sid,count in CQ_COUNTS.items():
            case=next(c for c in cases if c['name']==sid+'/book')
            data=yaml.safe_load(graph([sid],str(ROOT/'walton.yml')))
            data['statements']={'adequate_answer':'Independent accepted evidence for this dialogue.'}
            data['assumptions']=list(case['assumptions'])+['adequate_answer']
            base=self.evaluate(data)
            args=[(aid,a) for aid,a in base['arguments'].items() if a.get('scheme')==sid]
            self.assertEqual(len(args),1,(sid,args))
            aid,arg=args[0]
            self.assertIn(arg['conclusion'],case['in'])
            self.assertEqual(self.label(base,arg['conclusion']),'in')
            for cq in range(1,count+1):
                with self.subTest(scheme=sid,cq=cq):
                    # Use the checked source fixture's actual fully constructed instance.
                    j={'schema':SCHEMA,'dialogue':'discussion1',
                       'targets':{'target1':{'argument':aid,'scheme':sid,'parameters':arg['parameters']}},
                       'events':[{'id':'q1','type':'ask','target':'target1','cq':cq},
                                 {'id':'answer1','type':'answer','question':'q1','evidence':['proof1']}],
                       'evidence':{'proof1':{'dialogue':'discussion1','target':'target1','question':'q1',
                         'answer':'answer1','statement':'adequate_answer','accepted':True,'accepted_by':'Application reviewer'}}}
                    self.assertNotEqual(self.label(self.evaluate(snapshot(base,j,'q1')),arg['conclusion']),'in')
                    self.assertEqual(self.label(self.evaluate(snapshot(base,j,'answer1')),arg['conclusion']),'in')
    def test_answering_one_question_leaves_the_other_open(self):
        base=self.base();j=self.journal(base)
        j['events'].insert(1,{'id':'q2','type':'ask','target':'target1','cq':2})
        out=self.evaluate(snapshot(base,j,'answer1'))
        self.assertNotEqual(self.label(out,'should_do(ann,exercise)'),'in')
    def test_other_argument_support_survives_exact_target_challenge(self):
        base=self.base(['goal_of_agent(ann,fitness)','means_to_end(exercise,fitness)']);j=self.journal(base)
        made=snapshot(base,j,'q1');out=self.evaluate(made)
        self.assertEqual(self.label(out,'should_do(ann,exercise)'),'in')
        target=self.target(base);other=self.target(base,parameters=['exercise','fitness','ann'])
        self.assertEqual(made['arguments'][other['argument']],base['arguments'][other['argument']])
        self.assertEqual(made['arguments'][target['argument']]['undercutter'],base['arguments'][target['argument']]['undercutter'])
    def test_existing_evidential_undercutter_is_not_erased_by_answer(self):
        base=self.base(['more_efficient_alternative(ann,exercise,health)']);j=self.journal(base)
        self.assertNotEqual(self.label(base,'should_do(ann,exercise)'),'in')
        asked=snapshot(base,j,'q1'); answered=snapshot(base,j,'answer1')
        for aid,arg in base['arguments'].items(): self.assertEqual(answered['arguments'][aid],arg)
        self.assertGreater(len(asked['arguments']),len(base['arguments']))
        out=self.evaluate(answered)
        self.assertNotEqual(self.label(out,'should_do(ann,exercise)'),'in')
    def test_reject_bad_journals(self):
        base=self.base();j=self.journal(base)
        snapshot(base,j)  # ensure failures below cannot be masked by an invalid control
        variants=[]
        def add(name,fn):
            b=copy.deepcopy(base);x=copy.deepcopy(j);fn(b,x);variants.append((name,b,x))
        add('unknown target',lambda b,x:x['events'][0].update(target='absent'))
        add('wrong scheme',lambda b,x:x['targets']['target1'].update(scheme='two_person_practical'))
        add('wrong parameters',lambda b,x:x['targets']['target1'].update(parameters=['exercise','health','bob']))
        add('variable parameters',lambda b,x:x['targets']['target1'].update(parameters=['Exercise','health','ann']))
        add('wrong arity',lambda b,x:x['targets']['target1'].update(parameters=['exercise']))
        add('unknown CQ',lambda b,x:x['events'][0].update(cq=6))
        add('boolean CQ',lambda b,x:x['events'][0].update(cq=True))
        add('duplicate event',lambda b,x:x['events'].append(copy.deepcopy(x['events'][0])))
        add('unknown question',lambda b,x:x['events'][1].update(question='other_question'))
        add('no answer evidence',lambda b,x:x['events'][1].update(evidence=[]))
        add('duplicate evidence refs',lambda b,x:x['events'][1].update(evidence=['proof1','proof1']))
        add('unaccepted evidence',lambda b,x:x['evidence']['proof1'].update(accepted=False))
        add('no acceptor',lambda b,x:x['evidence']['proof1'].update(accepted_by=''))
        add('foreign dialogue',lambda b,x:x['evidence']['proof1'].update(dialogue='discussion2'))
        add('foreign target',lambda b,x:x['evidence']['proof1'].update(target='target2'))
        add('foreign answer',lambda b,x:x['evidence']['proof1'].update(answer='answer2'))
        add('absent statement',lambda b,x:x['evidence']['proof1'].update(statement='unknown_evidence'))
        add('self answer',lambda b,x:x['evidence']['proof1'].update(statement='should_do(ann,exercise)'))
        add('typo field',lambda b,x:x['events'][0].update(cqq=1))
        add('duplicate target',lambda b,x:x['targets'].update(target2=copy.deepcopy(x['targets']['target1'])))
        add('shared undercutter',lambda b,x:b['arguments'].update(other=copy.deepcopy(b['arguments'][x['targets']['target1']['argument']])))
        add('contested undercutter',lambda b,x:b.update(issues={'custom':{'positions':[b['arguments'][x['targets']['target1']['argument']]['undercutter'],'adequate_answer']}}))
        for name,b,x in variants:
            with self.subTest(name=name),self.assertRaises(DialogueError): snapshot(b,x)
        with self.assertRaises(DialogueError):snapshot(snapshot(base,j,'q1'),j)
        with self.assertRaises(DialogueError):snapshot(base,j,'missing_event')
    def test_quotes_and_escapes_preserve_identity(self):
        self.assertEqual(ground(' f( "New York", x ) ', 'test'), 'f("New York",x)')
        self.assertEqual(ground(" f( 'New York', x ) ", 'test'), "f('New York',x)")
        self.assertNotEqual(ground('"New York"', 'test'), ground('"NewYork"', 'test'))
        self.assertNotEqual(ground("'New York'", 'test'), ground("'NewYork'", 'test'))
        self.assertEqual(ground(r'"A\" B"','test'),r'"A\" B"')
        with self.assertRaises(DialogueError):ground('"unterminated','test')
        with self.assertRaises(DialogueError):ground('a "adjacent"','test')
        data=yaml.safe_load(graph(['practical_inference'],str(ROOT/'walton.yml')))
        data['statements']={'adequate_answer':'Independent accepted evidence.'}
        data['assumptions']=['goal_of_agent("New York",health)','means_to_end(exercise,health)','adequate_answer']
        base=self.evaluate(data)
        aid,arg=next((i,a) for i,a in base['arguments'].items() if a.get('scheme')=='practical_inference')
        j={'schema':SCHEMA,'dialogue':'quoted_actor','targets':{'target':{'argument':aid,'scheme':'practical_inference','parameters':arg['parameters']}},
           'events':[{'id':'q','type':'ask','target':'target','cq':1}]}
        out=self.evaluate(snapshot(base,j))
        self.assertNotEqual(self.label(out,arg['conclusion']),'in')
        bad=copy.deepcopy(j);bad['targets']['target']['parameters'][2]='"NewYork"'
        with self.assertRaises(DialogueError):snapshot(base,bad)
    def test_indirect_answer_dependency_is_rejected(self):
        base=self.base();base['assumptions'].remove('adequate_answer')
        base['arguments']['circular_evidence']={'scheme':'linked','premises':['should_do(ann,exercise)'],'conclusion':'adequate_answer'}
        base=self.evaluate(base);self.assertEqual(self.label(base,'adequate_answer'),'in')
        with self.assertRaises(DialogueError):snapshot(base,self.journal(base),'answer1')
    def test_independent_derived_answer_is_accepted(self):
        base=self.base();base['assumptions'].remove('adequate_answer')
        base['statements']['observation']='Independent observation.';base['assumptions'].append('observation')
        base['arguments']['independent_evidence']={'scheme':'linked','premises':['observation'],'conclusion':'adequate_answer'}
        base=self.evaluate(base);self.assertEqual(self.label(base,'adequate_answer'),'in')
        out=self.evaluate(snapshot(base,self.journal(base),'answer1'))
        self.assertEqual(self.label(out,'should_do(ann,exercise)'),'in')
    def test_duplicate_yaml_keys_fail(self):
        p=self.work/'duplicate.yml';p.write_text('schema: one\nschema: two\n')
        with self.assertRaises(DialogueError):read(p)
    def test_cli_emits_fresh_yaml_and_no_partial_invalid_output(self):
        base=self.base();journal=self.journal(base)
        basepath=self.work/'cli-base.yml';journalpath=self.work/'cli-journal.yml'
        basepath.write_text(yaml.safe_dump(base,sort_keys=False,allow_unicode=True))
        journalpath.write_text(yaml.safe_dump(journal,sort_keys=False,allow_unicode=True))
        command=[sys.executable,str(ROOT/'tools/dialogue_snapshot.py'),str(basepath),str(journalpath),'--through','q1']
        made=subprocess.run(command,capture_output=True,text=True)
        self.assertEqual(made.returncode,0,made.stderr)
        out=self.evaluate(yaml.safe_load(made.stdout))
        self.assertNotEqual(self.label(out,'should_do(ann,exercise)'),'in')
        journal['evidence']['proof1']['accepted']=False
        journalpath.write_text(yaml.safe_dump(journal,sort_keys=False,allow_unicode=True))
        invalid=subprocess.run(command,capture_output=True,text=True)
        self.assertEqual(invalid.returncode,2)
        self.assertEqual(invalid.stdout,'')
        self.assertIn('has not been explicitly accepted',invalid.stderr)
    def test_ignorance_burden_is_separate_validated_metadata(self):
        data=yaml.safe_load(graph(['ignorance'],str(ROOT/'walton.yml')))
        data['assumptions']=['would_be_known(alarm)','¬known(alarm)']
        data['statements']={'alarm':'An alarm occurs.','¬alarm':'No alarm occurs.'}
        data['issues']={'ignorance_decision':{'positions':['alarm','¬alarm'],'standard':'PE'}}
        base=self.evaluate(data)
        aid=next(a for a,x in base['arguments'].items() if x.get('scheme')=='ignorance')
        issue=next(i for i,x in base['issues'].items() if '¬alarm' in x['positions'])
        row={'argument':aid,'scheme':'ignorance','parameters':['alarm'],'dialogue':'discussion1',
             'probandum':'¬alarm','burden_holder':'proponent','issue':issue,'standard':'PE'}
        journal={'schema':SCHEMA,'dialogue':'discussion1','targets':{},'events':[],'burdens':[row]}
        out=snapshot(base,journal)
        self.assertEqual(out['arguments'],base['arguments'])
        self.assertEqual(out['meta']['dialogue_snapshot']['burden_allocations'][0]['burden_holder'],'proponent')
        for field,value in [('dialogue','foreign'),('probandum','unknown'),('burden_holder',''),('issue','unknown'),('standard','BRD')]:
            bad=copy.deepcopy(journal);bad['burdens'][0][field]=value
            with self.subTest(field=field),self.assertRaises(DialogueError):snapshot(base,bad)
        journal['burdens']=[]
        with self.assertRaises(DialogueError):snapshot(base,journal)
    def test_every_ignorance_family_has_an_actual_checked_allocation(self):
        cases=[]
        for path in (ROOT/'tools/vv/cases').glob('*.yml'):
            document=yaml.safe_load(path.read_text())
            if isinstance(document,dict):cases.extend(document.get('cases',[]))
        for sid in sorted(IGNORANCE):
            with self.subTest(scheme=sid):
                case=next(c for c in cases if c['name']==sid+'/book')
                data=yaml.safe_load(graph([sid],str(ROOT/'walton.yml')))
                data['assumptions']=list(case['assumptions'])
                conclusion=case['in'][0]
                if sid=='negative_practical_true':
                    other=conclusion.replace('act_on_presumption_true(', 'act_on_presumption_false(',1)
                elif sid=='negative_practical':
                    other=conclusion.replace('act_on_presumption_false(', 'act_on_presumption_true(',1)
                else:other=conclusion[1:] if conclusion.startswith('¬') else '¬'+conclusion
                data['statements']={conclusion:'Target claim.',other:'Opposed position.',
                                    'ultimate_claim':'Application-declared broader ultimate claim.',
                                    'ultimate_alternative':'Application-declared alternative.'}
                data['issues']={'target_issue':{'positions':[conclusion,other],'standard':'PE'},
                                'broader_issue':{'positions':['ultimate_claim','ultimate_alternative'],'standard':'PE'}}
                base=self.evaluate(data)
                aid,arg=next((a,x) for a,x in base['arguments'].items() if x.get('scheme')==sid)
                issue=next(i for i,x in base['issues'].items() if conclusion in x['positions'])
                row={'argument':aid,'scheme':sid,'parameters':arg['parameters'],'dialogue':'allocation_review',
                     'probandum':conclusion,'burden_holder':'designated_proponent','issue':issue,'standard':'PE'}
                journal={'schema':SCHEMA,'dialogue':'allocation_review','targets':{},'events':[],'burdens':[row]}
                made=snapshot(base,journal)
                self.assertEqual(made['arguments'],base['arguments'])
                self.assertEqual(self.label(self.evaluate(made),conclusion),'in')
                for field,value in [('dialogue','foreign'),('issue','unknown'),('standard','BRD'),
                                    ('probandum','absent'),('burden_holder',''),('parameters',['wrong_case']*len(arg['parameters']))]:
                    bad=copy.deepcopy(journal);bad['burdens'][0][field]=value
                    with self.subTest(field=field),self.assertRaises(DialogueError):snapshot(base,bad)
                bad=copy.deepcopy(journal);bad['burdens']=[]
                with self.assertRaises(DialogueError):snapshot(base,bad)
                bad=copy.deepcopy(journal);bad['burdens'].append(copy.deepcopy(row))
                with self.assertRaises(DialogueError):snapshot(base,bad)
                # Logical relevance to the wider dialogue is explicitly adjudicated
                # by the application; the checker validates allocation, not that proof.
                broader=copy.deepcopy(journal)
                broader['burdens'][0].update(probandum='ultimate_claim',issue='broader_issue')
                self.assertEqual(snapshot(base,broader)['meta']['dialogue_snapshot']['burden_allocations'][0]['probandum'],
                                 'ultimate_claim')

if __name__=='__main__':unittest.main()
