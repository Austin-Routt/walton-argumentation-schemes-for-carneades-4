#!/usr/bin/env bash
# Build, vet and test both engines as shipped in third_party/.
#
#   tools/vv/engine-tests.sh
#
# For each module: go build ./..., go vet ./..., go test ./... (PLAN.md 1.7,
# E-BUILD), then the concurrency tests under the race detector. Carneades runs against the GoCHR of third_party/ (a replace
# directive in a copy of its go.mod; the module's own files are not touched).
# TestCouchDB skips itself unless a CouchDB server listens on 127.0.0.1:5984.
# VV_GOMOD_EXTRA, as for the suite, names a file of extra go.mod lines
# (replace directives for an offline machine).
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo="${WALTON_REPO:-$(cd "$here/../.." && pwd)}"
tp="$repo/third_party"

tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

echo "-- GoCHR"
(cd "$tp/GoCHR" && go build ./... && go vet ./... && go test -count=1 ./...)

echo "-- carneades-4 (against third_party/GoCHR)"
cp "$tp/carneades-4/go.mod" "$tmp/go.mod"
cp "$tp/carneades-4/go.sum" "$tmp/go.sum"
go mod edit -replace "github.com/hfried/GoCHR=$tp/GoCHR" "$tmp/go.mod"
if [ -n "${VV_GOMOD_EXTRA:-}" ]; then cat "$VV_GOMOD_EXTRA" >> "$tmp/go.mod"; fi
m=(-mod=mod -modfile="$tmp/go.mod")
(cd "$tp/carneades-4" && go build "${m[@]}" ./... && go vet "${m[@]}" ./... && go test "${m[@]}" -count=1 ./...)

# The race job (docs/contracts.md, contract 1): the concurrency tests of both
# engines under the race detector, which needs cgo and a C compiler.
if [ "$(go env CGO_ENABLED)" = "1" ] && command -v "$(go env CC)" >/dev/null 2>&1; then
  echo "-- race: the concurrency tests under -race"
  (cd "$tp/GoCHR" && go test -race -count=1 -run 'Concurrent' ./src/engine/CHR/)
  (cd "$tp/carneades-4" && go test "${m[@]}" -race -count=1 -run 'Concurrent|WaitingRun' ./src/engine/caes/ ./src/engine/caes/encoding/yaml/)
else
  echo "-- race: skipped, cgo or a C compiler is missing" >&2
  exit 1
fi
