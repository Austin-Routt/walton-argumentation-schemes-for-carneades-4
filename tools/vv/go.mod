module waltonvv

go 1.18

require (
	github.com/bronze1man/go-yaml2json v0.0.0-20150129175009-f6f64b738964
	github.com/carneades/carneades-4 v0.0.0
	github.com/hfried/GoCHR v0.0.0
	gopkg.in/yaml.v2 v2.4.0
)

// The engine is built from its sources in this repository, third_party/ at the
// repository root (see sources.lock).
replace github.com/carneades/carneades-4 => ../../third_party/carneades-4

replace github.com/hfried/GoCHR => ../../third_party/GoCHR
