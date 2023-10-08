package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"testing"
	"text/template"
)

func Test1(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "demo_obj.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	//err = ast.Print(fset, f)
	//if err != nil {
	//	return
	//}
	data := RenderData{
		PackageName: f.Name.Name,
		Cmd:         "demo",
	}
	for _, decl := range f.Decls {
		if stc, ok := decl.(*ast.GenDecl); ok && stc.Tok == token.TYPE {
			for _, spec := range stc.Specs {
				if tp, ok := spec.(*ast.TypeSpec); ok {
					node := RenderNode{
						StructName: tp.Name.Name,
					}
					if stp, ok := tp.Type.(*ast.StructType); ok {
						if !stp.Struct.IsValid() {
							continue
						}

						for _, field := range stp.Fields.List {
							if len(field.Names) > 0 {
								log.Println("field:", field)
								node.Fields = append(node.Fields, Field{
									FieldName: field.Names[0].Name,
								})
							}
						}
						data.Nodes = append(data.Nodes, node)
					}
				}
			}
		}
	}
	log.Println(data)

	var tmplFile = "clone.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	// create a new file
	file, _ := os.Create("demo.go")
	defer file.Close()
	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
