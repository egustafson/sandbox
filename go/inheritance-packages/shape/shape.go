package shape

// An 'abstract' base class -- somewhat.

type ShapeRec struct {
	Name string // the attribute at the base of the class hierarchy
}

type Shape interface {
	GetShapeRec() *ShapeRec
	GetWidth() int
	GetHeight() int
}
