package jsontag

import (
	"flag"
	"go/ast"

	"github.com/eltonjr/json-interface-linter/analyzer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func Analyzer(flags flag.FlagSet) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "jsontag",
		Doc:      "Check if structs tagged as json contain an interface",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    flags,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.StructType)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		n := node.(*ast.StructType)

		if !analyzer.HasJSONTag(n) {
			return
		}

		analyzer.CheckStructType(pass, n)
	})
	return nil, nil
}
