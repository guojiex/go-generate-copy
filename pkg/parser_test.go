package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func Test1(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "demo_obj.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	for _, decl := range f.Decls {
		if stc, ok := decl.(*ast.GenDecl); ok && stc.Tok == token.TYPE {
			for _, spec := range stc.Specs {
				if tp, ok := spec.(*ast.TypeSpec); ok {
					log.Println("[struct name]", tp.Name)
					if stp, ok := tp.Type.(*ast.StructType); ok {
						if !stp.Struct.IsValid() {
							continue
						}
						for _, field := range stp.Fields.List {
							log.Println("field:", field.Names, field.Type)
						}
					}
				}
			}
		}
	}
}
