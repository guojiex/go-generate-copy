package parser

type Embedded struct {
}

type Attribute struct {
	Id    int32
	ValId *int32
}

type Object struct {
	Embedded
	Id             int64
	Attributes     []Attribute
	Comment        string
	CommentPointer *string
}
