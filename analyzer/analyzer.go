package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "jsoninterface",
	Doc:      "Check if serialized structs contain an interface",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := inspector.New(pass.Files)
	inspector.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.StructType:
			for _, field := range n.Fields.List {
				if field.Tag == nil {
					continue
				}
				if !strings.Contains(field.Tag.Value, "json") {
					continue
				}

				fieldType, ok := field.Type.(*ast.Ident)
				if !ok {
					continue
				}
				// skip basic types
				if fieldType.Obj == nil {
					continue
				}

				typeSpec, ok := fieldType.Obj.Decl.(*ast.TypeSpec)
				if !ok {
					continue
				}
				_, ok = typeSpec.Type.(*ast.InterfaceType)
				if !ok {
					continue
				}

				pass.Reportf(field.Pos(), "interface field %s is exported as json attribute", field.Names[0].Name)
			}
		}
	})
	return nil, nil
}
