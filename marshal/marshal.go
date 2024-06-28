package marshal

import (
	"flag"
	"fmt"
	"go/ast"

	"github.com/eltonjr/json-interface-linter/analyzer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

func Analyzer(flags flag.FlagSet) (*analysis.Analyzer, error) {
	if f := flags.Lookup("marshalers"); f != nil {
		d, err := ReadMarshalers(f.Value.String())
		if err != nil {
			return nil, err
		}
		defaultMarshalers = d
	}

	return &analysis.Analyzer{
		Name:     "marshal",
		Doc:      "Check if marshaled structs contain an interface",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    flags,
	}, nil
}

func run(pass *analysis.Pass) (interface{}, error) {
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
