package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func HasJSONTag(s *ast.StructType) bool {
	for _, field := range s.Fields.List {
		if field.Tag != nil && strings.Contains(field.Tag.Value, "json") {
			return true
		}
	}
	return false
}

func CheckStructType(pass *analysis.Pass, s *ast.StructType) {
	for _, field := range s.Fields.List {
		// lower case fields are not exported
		if !field.Names[0].IsExported() {
			continue
		}
		// fields with tag json:"-" are not exported
		if field.Tag != nil && strings.Contains(field.Tag.Value, "json:\"-\"") {
			continue
		}

		tav := pass.TypesInfo.Types[field.Type]
		CheckType(pass, field.Pos(), tav.Type)
	}
}

func CheckType(pass *analysis.Pass, pos token.Pos, t types.Type) {
	if types.IsInterface(t) {
		pass.Reportf(pos, "interface %s is exported as json", t.String())
	}

	switch t := t.Underlying().(type) {
	case *types.Struct:
		CheckStruct(pass, pos, t)
	case *types.Map:
		CheckMap(pass, pos, t)
	}
}

func CheckStruct(pass *analysis.Pass, pos token.Pos, s *types.Struct) {
	for i := 0; i < s.NumFields(); i++ {
		field := s.Field(i)
		if !field.Exported() {
			continue
		}

		// TODO check json:"-"

		CheckType(pass, pos, field.Type())
	}
}

func CheckMap(pass *analysis.Pass, pos token.Pos, m *types.Map) {
	if types.IsInterface(m.Elem().Underlying()) {
		pass.Reportf(pos, "interface %s is exported as json", m.Elem().String())
	}
}
