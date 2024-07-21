package marshal

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"

	"github.com/eltonjr/json-interface-linter/internal/analyzer"
)

var Analyzer = &analysis.Analyzer{
	Name:     "marshal",
	Doc:      "Check if marshaled structs contain an interface",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	analyzer.InitExcluders()
	initMarshalers()

	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		n := node.(*ast.CallExpr)

		o := typeutil.Callee(pass.TypesInfo, n)
		if o == nil {
			return
		}

		// probably a builtin function like make or new
		if o.Pkg() == nil {
			return
		}

		fnName := fmt.Sprintf("%s.%s", o.Pkg().Path(), o.Name())
		marshaler, found := isMarshaler(fnName, defaultMarshalers)
		if !found {
			return
		}

		// should we allow many arguments?
		arg := n.Args[marshaler.argument]

		tav := pass.TypesInfo.Types[arg]
		analyzer.CheckType(pass, n.Pos(), tav.Type)
	})
	return nil, nil
}
