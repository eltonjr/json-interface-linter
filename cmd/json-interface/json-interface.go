package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/eltonjr/json-interface-linter/analyzer"
	"github.com/eltonjr/json-interface-linter/internal/logger"
	"github.com/eltonjr/json-interface-linter/jsontag"
	"github.com/eltonjr/json-interface-linter/marshal"
)

func main() {
	logger.RegisterFlags()
	analyzer.RegisterFlags()
	marshal.RegisterFlags()

	multichecker.Main(
		jsontag.Analyzer,
		marshal.Analyzer,
	)
}
