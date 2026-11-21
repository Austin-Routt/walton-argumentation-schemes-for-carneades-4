#!/usr/bin/env python3
"""Compare the results of tools/vv under two Go toolchains (PLAN.md Phase 6).

    python3 tools/vv/compare_results.py A.jsonl B.jsonl

Each file holds one JSON object per check (VV_RESULTS of the suite). The two
runs agree when they ran the same checks and each check passed or failed in
both. Exit status 0 when they agree, 1 when they do not, 2 on bad input.
"""
import json
import sys


def load(path):
    out = {}
    with open(path, encoding="utf-8") as f:
        for n, line in enumerate(f, 1):
            line = line.strip()
            if not line:
                continue
            try:
                r = json.loads(line)
            except ValueError as e:
                raise SystemExit(f"{path}:{n}: {e}")
            out[(r.get("file", ""), r.get("scheme", ""), r.get("check", ""))] = bool(r.get("pass"))
    return out


def main(argv):
    if len(argv) != 3:
        print(__doc__.strip(), file=sys.stderr)
        return 2
    a, b = load(argv[1]), load(argv[2])
    only_a = sorted(set(a) - set(b))
    only_b = sorted(set(b) - set(a))
    differ = sorted(k for k in set(a) & set(b) if a[k] != b[k])
    per_scheme = {k for k in a if not k[1].startswith("(")}
    print(f"{argv[1]}: {len(a)} checks, {sum(a.values())} pass")
    print(f"{argv[2]}: {len(b)} checks, {sum(b.values())} pass")
    print(f"per-scheme checks in the first run: {len(per_scheme)}")
    for title, keys in (("only in the first run", only_a), ("only in the second run", only_b), ("different outcome", differ)):
        if keys:
            print(f"{title}: {len(keys)}")
            for k in keys[:20]:
                print("  " + "|".join(k))
    if only_a or only_b or differ:
        print("the two runs do not agree")
        return 1
    print("the two runs agree on every check")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
