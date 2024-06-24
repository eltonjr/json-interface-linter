package jsontag

import (
	"flag"
	"go/ast"
	"go/types"
	"strings"

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
		for _, field := range n.Fields.List {
			if field.Tag == nil {
				continue
			}
			if !strings.Contains(field.Tag.Value, "json") || strings.Contains(field.Tag.Value, "json:\"-\"") {
				continue
			}

			tav := pass.TypesInfo.Types[field.Type]
			if !types.IsInterface(tav.Type) {
				continue
			}

			pass.Reportf(field.Pos(), "interface field %s is exported as json attribute", field.Names[0].Name)
		}
	})
	return nil, nil
}
