#!/usr/bin/env python3
"""Bring a YAML file written for release R1 (library 0.1.0) to release R2.

    python3 tools/migrate_r2.py FILE [FILE ...]        # rewrite in place
    python3 tools/migrate_r2.py --check FILE [...]     # print what would change, exit 1 if anything would
    python3 tools/migrate_r2.py --stdout FILE          # print the migrated text

The table is tools/migration_r2.yml (PLAN.md D4, 5.6). Its rows of the kinds
rename-scheme and rename-predicate are applied: a scheme id or a predicate is
replaced wherever it stands as an identifier, in statements, in scheme
definitions, in `scheme:` fields and in the keys of `language:` and
`statements:`. Natural-language text is left alone: the values of
`language:` and `statements:`, and every `meta:` block. Rows of the other
kinds (a statement with other arguments, another arity, a statement removed,
a premise added) cannot be applied without knowing the case; the tool lists
the ones whose old statement the file uses, and the reader decides.

The rename-file rows name the scheme files that were renamed; a file name
inside the text is replaced as well.
"""
import os
import re
import sys

import yaml

HERE = os.path.dirname(os.path.abspath(__file__))
TABLE = os.path.join(HERE, "migration_r2.yml")

# top-level sections whose entries are natural language after the key
KEYED_TEXT = {"language", "statements"}
# keys whose values are natural language wherever they stand
TEXT_KEYS = {"title", "source", "note", "description", "text", "summary"}


def load_table(path=TABLE):
    with open(path, encoding="utf-8") as f:
        rows = yaml.safe_load(f)["changes"]
    schemes, preds, files, manual = {}, {}, {}, []
    for r in rows:
        k = r["kind"]
        if k == "rename-scheme":
            schemes[r["old"]] = r["new"]
        elif k == "rename-predicate":
            o, n = r["old"].split("/")[0], r["new"].split("/")[0]
            preds[o] = n
        elif k == "rename-file":
            files[r["old"]] = r["new"]
        elif k in ("arguments", "arity", "replaced", "removed", "premise"):
            manual.append(r)
    return schemes, preds, files, manual


def token_re(name):
    # an identifier: not preceded or followed by a letter, digit, _ or -
    return re.compile(r"(?<![\w\-])" + re.escape(name) + r"(?![\w\-])")


class Migrator:
    def __init__(self, table=TABLE):
        self.schemes, self.preds, self.files, self.manual = load_table(table)
        names = dict(self.schemes)
        for o, n in self.preds.items():
            names.setdefault(o, n)
        self.names = [(token_re(o), n) for o, n in sorted(names.items(), key=lambda x: -len(x[0]))]
        self.filenames = [(re.compile(re.escape(o)), n) for o, n in self.files.items()]

    def ids(self, s):
        for rx, n in self.names:
            s = rx.sub(n, s)
        for rx, n in self.filenames:
            s = rx.sub(n, s)
        return s

    def line(self, ln, section, in_meta):
        """Migrate one line, given its top-level section and whether it lies in a meta block."""
        if in_meta or ln.lstrip().startswith("#"):
            return ln
        m = re.match(r"^(\s*)(-\s+)?([\w\-]+)(\s*:\s*)(.*)$", ln)
        if m and m.group(3) in TEXT_KEYS:
            return ln
        if section in KEYED_TEXT and re.match(r"^\s+\S", ln):
            # key: text -> migrate the key only (a statement or name/arity)
            km = re.match(r'^(\s+)("(?:[^"\\]|\\.)*"|\'[^\']*\'|[^:"\']+?)(\s*:\s*)(.*)$', ln)
            if km:
                return km.group(1) + self.ids(km.group(2)) + km.group(3) + km.group(4)
            return ln
        return self.ids(ln)

    def text(self, src):
        out, section, meta_indent = [], None, None
        for ln in src.split("\n"):
            stripped = ln.strip()
            indent = len(ln) - len(ln.lstrip())
            top = re.match(r"^([\w\-]+)\s*:", ln)
            if top:
                section = top.group(1)
                meta_indent = 0 if section == "meta" else None
            elif meta_indent is not None and stripped and indent <= meta_indent:
                meta_indent = None
            if meta_indent is None and re.match(r"^\s+(-\s+)?meta\s*:\s*$", ln):
                out.append(ln)
                meta_indent = indent + (2 if re.match(r"^\s+-\s+meta", ln) else 0)
                continue
            out.append(self.line(ln, section, meta_indent is not None and (section == "meta" or indent > meta_indent)))
        return "\n".join(out)

    def needs_hand(self, src):
        """Rows that cannot be applied mechanically and whose old statement the text uses."""
        hits = []
        for r in self.manual:
            old = r.get("old") or ""
            m = re.match(r"\s*¬?\s*([\w\-]+)", old)
            if not m:
                continue
            f = m.group(1)
            # the R2 name of the functor, after renaming
            f2 = self.preds.get(f, f)
            arity = old.count(",") + 1 if "(" in old else 0
            pat = re.compile(r"(?<![\w\-])(" + re.escape(f) + "|" + re.escape(f2) + r")\s*(\(([^()]|\([^()]*\))*\))?")
            for mm in pat.finditer(src):
                args = mm.group(2)
                n = 0 if not args else args.count(",") + 1 - sum(a.count(",") for a in re.findall(r"\([^()]*\)", args[1:-1]))
                if n == arity:
                    hits.append(r)
                    break
        return hits


def main(argv):
    check = "--check" in argv
    stdout = "--stdout" in argv
    files = [a for a in argv[1:] if not a.startswith("--")]
    if not files:
        print(__doc__.strip(), file=sys.stderr)
        return 2
    mig = Migrator()
    changed = 0
    for f in files:
        with open(f, encoding="utf-8") as fh:
            src = fh.read()
        new = mig.text(src)
        for r in mig.needs_hand(src):
            print(f"{f}: by hand ({r['kind']}, {r['record']}): {r.get('old')} -> {r.get('new') or '(removed)'}" + (f"  [{r['note']}]" if r.get("note") else ""), file=sys.stderr)
        if stdout:
            sys.stdout.write(new)
            continue
        if new != src:
            changed += 1
            if check:
                print(f"{f}: would change")
            else:
                with open(f, "w", encoding="utf-8", newline="\n") as fh:
                    fh.write(new)
                print(f"{f}: migrated")
    return 1 if check and changed else 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
