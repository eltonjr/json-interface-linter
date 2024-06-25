package main

import (
	"flag"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	flags := parseFlags()
	ja := jsontagAnalyzer(flags)
	ma, err := marshalAnalyzer(flags)
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
