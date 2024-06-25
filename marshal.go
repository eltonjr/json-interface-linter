package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

func marshalAnalyzer(flags flag.FlagSet) (*analysis.Analyzer, error) {
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
		Run:      runMarshal,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    flags,
	}, nil
}

func runMarshal(pass *analysis.Pass) (interface{}, error) {
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

		// dont allow marshal of interfaces
		tav := pass.TypesInfo.Types[arg]
		if types.IsInterface(tav.Type) {
			pass.Reportf(arg.Pos(), "interface value %s is exported as json", tav.Type.String())
			return
		}

		argType := tav.Type.Underlying()
		s, ok := argType.(*types.Struct)
		if !ok {
			return
		}

		for i := 0; i < s.NumFields(); i++ {
			field := s.Field(i)
			if !field.Exported() {
				continue
			}

			if types.IsInterface(field.Type()) {
				pass.Reportf(arg.Pos(), "interface field %s is exported as json attribute", field.Name())
			}
		}
	})
	return nil, nil
}
