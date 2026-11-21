#!/usr/bin/env python3
"""Compare the results of tools/vv under two Go toolchains (PLAN.md Phase 6).

    python3 tools/vv/compare_results.py A.jsonl B.jsonl

Each file holds one JSON object per unique check, with a boolean pass value
(VV_RESULTS of the suite). Invalid or duplicate records are rejected. The two
runs agree when they ran the same checks and each check passed or failed in
both. Exit status 0 when they agree, 1 when they do not, 2 on bad input.
"""
import json
import sys


def load(path):
    out = {}
    first_line = {}
    with open(path, encoding="utf-8") as f:
        for n, line in enumerate(f, 1):
            line = line.strip()
            if not line:
                continue
            try:
                r = json.loads(line)
            except ValueError as e:
                raise ValueError(f"{path}:{n}: {e}") from e
            if not isinstance(r, dict):
                raise ValueError(f"{path}:{n}: result must be a JSON object")
            for field in ("file", "scheme", "check"):
                if not isinstance(r.get(field), str):
                    raise ValueError(f"{path}:{n}: {field} must be a JSON string")
            if type(r.get("pass")) is not bool:
                raise ValueError(f"{path}:{n}: pass must be a JSON boolean")
            key = (r["file"], r["scheme"], r["check"])
            if key in out:
                raise ValueError(f"{path}:{n}: duplicate check key {key!r} (first at line {first_line[key]})")
            first_line[key] = n
            out[key] = r["pass"]
    return out


def main(argv):
    if len(argv) != 3:
        print(__doc__.strip(), file=sys.stderr)
        return 2
    try:
        a, b = load(argv[1]), load(argv[2])
    except (OSError, ValueError) as e:
        print(e, file=sys.stderr)
        return 2
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
