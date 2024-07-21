package jsontag

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/eltonjr/json-interface-linter/internal/analyzer"
)

var Analyzer = &analysis.Analyzer{
	Name:     "jsontag",
	Doc:      "Check if structs tagged as json contain an interface",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	analyzer.InitExcluders()
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
