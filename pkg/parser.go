package parser

type Field struct {
	FieldName string
}

type RenderNode struct {
	StructName string
	Fields     []Field
}

type RenderData struct {
	PackageName string
	Cmd         string
	Nodes       []RenderNode
}
