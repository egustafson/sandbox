package square

import (
	"github.com/egustafson/sandbox/go/inheritance-packages/shape"
)

type Square struct {
	shape.ShapeRec
	width int
}

func NewSquare(width int) *Square {
	//s := Square{Name: "square", width: width}
	s := Square{width: width}
	s.Name = "square"
	return &s
}

func (s *Square) GetShapeRec() *shape.ShapeRec {
	return &s.ShapeRec
}

func (s *Square) GetWidth() int {
	return s.width
}

func (s *Square) GetHeight() int {
	return s.width // a square is the same width and height
}
