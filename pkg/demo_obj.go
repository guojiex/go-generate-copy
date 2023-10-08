package parser

type Attribute struct {
	Id    int32
	ValId *int32
}

type ExampleObject struct {
	Id             int64
	Attributes     []Attribute
	Comment        string
	CommentPointer *string
}
