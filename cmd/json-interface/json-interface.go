package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/eltonjr/json-interface-linter/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
