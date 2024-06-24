package main

import (
	"flag"

	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/eltonjr/json-interface-linter/jsontag"
	"github.com/eltonjr/json-interface-linter/marshal"
)

func main() {
	flags := parseFlags()
	ja := jsontag.Analyzer(flags)
	ma, err := marshal.Analyzer(flags)
	if err != nil {
		panic(err)
	}

	multichecker.Main(
		ja,
		ma,
	)
}

func parseFlags() flag.FlagSet {
	flags := flag.NewFlagSet("json-interface-linter", flag.ExitOnError)

	flags.String("marshalers", "", "path to marshalers file")

	return *flags
}
