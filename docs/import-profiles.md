# Input and conversion profiles

The primary library graph format is native Carneades YAML. Other formats do
not identify executable Walton schemes merely by using the same scheme name.
Successful parsing and structural validation are separate from preserving the
meaning of a source graph. These boundaries are part of library 0.4.1 / engine
4.3.1+walton.8 and the next validation effort must test each claimed conversion
separately.

| Input | Admission and interpretation | Acceptance scope |
|---|---|---|
| Native YAML | Known fields, unique keys, statement-id premise lists or role-to-id mappings; explicit assumptions, defaults, questions and issues | Primary structured CLI/API profile in `contracts.md` |
| Restricted AIF JSON | I/RA/binary-CA graph; policy below | Tested project projection, not general AIF or AIFdb conformance |
| CAF | CAF 1.3 shape with explicit legacy policy | Experimental; opt-in required (CLI and web service) |
| AGXML | One EDU per ADU; supported relation shapes and declared endpoints | Experimental; segmentation joints and full corpus conformance excluded |
| LKIF | One argument graph; explicit premise polarity and type | Experimental; legacy assumption interpretation is not a full dialogue translation |
| Dung TGF | Declared argument ids, a `#` separator, unlabelled two-id attack edges | Small abstract-framework interface; optional node text is outside Dung semantics |

## Native YAML

Unknown semantic fields, duplicate keys, null arguments/issues/schemes and
unsupported premise shapes are rejected. For example, a list containing
`{statement: p, positive: false}` is not a native premise list and must not
silently become an argument with no premises. Use native statement identities
and explicit negation or an issue according to the intended argument; changing
an external scheme name to `linked` does not supply that interpretation.

The importer does not replace `validation.Validate`, context checks or
inference. Go API callers must follow the sequence in `contracts.md`. The CLI
and both CAES HTTP evaluation routes perform those checks. Native YAML round
trips preserve the declared theory and case state under the existing contract.

## Restricted AIF projection

The JSON object must contain `nodes` and `edges` arrays. Node and edge ids must
be nonempty and unique within their respective collections. Duplicate JSON
keys, including escaped aliases, trailing objects, unknown fields, dangling
edges and duplicate relations are refused. Only I, RA and CA nodes are
admitted. Edges connect I to RA/CA or RA/CA to I. An RA has one or more I
premises and exactly one I conclusion. A CA connects exactly two distinct I
nodes, one incoming and one outgoing; an I node cannot participate in several
CA issues in this profile.

The current converter uses I-node identifiers as native statement keys. Those
keys must also pass native validation as ground atomic formulas: `p` can be
used, while a numeric identifier such as `"1"` is refused at evaluation. A
general external identifier requires an explicit, reversible mapping with
collision and provenance checks. This mapping is future adapter work, not
something this converter supplies automatically.

The conversion applies these explicit project policies:

* An RA becomes a generic linked argument. Its external scheme string is
  retained as `aif_scheme` metadata, never as proof that a Walton instance's
  premises, parameters and critical questions were supplied.
* A binary CA becomes a **symmetric incompatibility issue**. General directed
  AIF conflict does not have this meaning automatically; an application must
  justify this projection before using its output as a semantic reference.
* An I node is stipulated only if it has no supporting argument and is not an
  issue position. This is an explicit evidence policy, not a property inferred
  from the truth of its text. Supported conclusions and disputed positions are
  not silently assumed.
* Both issue positions receive consistent backlinks. Conversion-policy
  metadata is retained in the resulting graph.

Locution/dialogue nodes such as L and YA, preferences, shared conflict
structures and wider AIFdb fields are refused. A future AIFdb adapter must
define its source semantics and loss accounting before relaxing these checks.

## Experimental legacy conversion

`carneades check` and `carneades eval` refuse CAF, AGXML and LKIF by default
with the typed `unsupported` outcome and CLI exit code 2. The web service
(`carneades server`) enforces the same profile; from walton.8, the form route
returns HTTP 422 for that refusal. CLI exit codes are not HTTP status codes.
A deliberate experiment can opt in:

```sh
carneades check -f caf -experimental-import legacy.xml
carneades eval -f caf -experimental-import -t yaml legacy.xml
carneades server -experimental-import
```

Before engine 4.3.1+walton.7 the web service called the legacy converters
without the opt-in; `src/web/profile.go` now applies the same profile.
The gate runs before conversion and inference. `FormValue` parses the
multipart request first, so the gate is not a guarantee that upload bytes
have not been read or spooled.

The flag does not bypass admission or structural validation and does not
certify fidelity. Legacy Go import APIs remain available as experimental
converters. Their checked boundaries now reject the incompatible Kialo
export shapes used in the review; that rejection is not a replacement Kialo
adapter.

### HTTP admission and inference outcomes

From walton.8, `/carneades/eval` and `/carneades/eval-arg-graph` both validate
authored arguments and check `InferContext(req.Context())` before labelling
or exporting the graph. They return HTTP 400 for invalid input, 422 for
unsupported input, inconsistent context or a reached engine limit, 408 for
cancelled inference, and 500 for an unclassified/internal failure. Raw-API
errors are JSON strings; form errors use the existing HTML template. Valid
results retain their existing output formats. The change intentionally
rejects invalid authored arguments that the old raw route evaluated.

These tests cover admission and inference failure before result output.
They do not certify production hosting, request-size protection, the Dung
web route, arbitrary external Graphviz behavior or atomic recovery after
an exporter has already written bytes. The supported product profile
remains the offline CLI and Go API described in `contracts.md`.

CAF defaults are read according to the 1.3 schema: absent statement weight is
0.5, argument polarity is pro, and premise polarity is positive. The repaired
converter no longer turns omissions of those fields into negative assumptions
or negative premises. It also avoids the forward-undercutter variable-shadow
panic. Unsupported versions, unrecognized reviewed fields, invalid endpoints
and unsupported strict arguments are refused. This does not establish full
CAF metadata, proof-standard or weighting equivalence.

AGXML requires ADUs, valid relation endpoints, and direct one-to-one EDU/ADU
segmentation in this profile. Assumption flags and argument backlinks are
preserved correctly, and ordinary example arguments have initialized
metadata. Requiring this shape deliberately excludes segmentation joints;
it is not a claim of complete `arg-microtexts` coverage.

LKIF refuses extra graphs rather than dropping them. Premise kinds and
polarities must be explicit and supported. Exception arguments now have
their conclusion backlink so a supported exception can defeat its target;
generated argument-id collisions are rejected. The legacy conversion of
assumption premises into graph stipulations still requires a separate
application-level justification. It is not a dialogue default protocol.

## Dung TGF and benchmark preparation

Every admitted edge is an attack. Edge labels are rejected, including Kialo
support/attack labels that would otherwise collapse into attacks. Duplicate
nodes, undeclared endpoints, missing separators and malformed edges are
rejected. Repeated identical attacks have set semantics. A bipolar or
structured argument graph needs an explicit, separately justified Dung
projection; deleting its edge labels is not such a justification.

The existing interface is not the modern indexed ICCMA input/certificate
protocol. The complete/preferred/stable algorithms retain exhaustive subset
enumeration and are not accepted as large-instance competition solvers. These
capability and capacity questions belong to the next benchmark plan. The
corrected extension-family equality helper and real three-node odd-cycle
fixture strengthen small-framework verification without changing that scope.

## Evidence

Native regressions are in the respective encoding packages, the CLI import
profile tests, `dung/extension_family_test.go`, the corrected odd-cycle test,
and `dung/encoding/tgf/tgf_test.go`. The review archive retains before/after
results for the twelve captured Kialo exports, three focused AIF controls and
the minimal YAML premise-loss example. No full corpus benchmark is claimed.
