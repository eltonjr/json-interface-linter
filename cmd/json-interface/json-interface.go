package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/eltonjr/json-interface-linter/jsontag"
)

func main() {
	singlechecker.Main(jsontag.Analyzer)
}
