package circle

import (
	"github.com/egustafson/sandbox/go/inheritance-packages/shape"
)

type Circle struct {
	shape.ShapeRec
	radius int
}

func NewCircle(r int) *Circle {
	c := Circle{radius: r}
	c.Name = "circle"
	return &c
}

func (c *Circle) GetShapeRec() *shape.ShapeRec {
	return &c.ShapeRec
}

func (c *Circle) GetWidth() int {
	return c.radius * 2
}

func (c *Circle) GetHeight() int {
	return c.radius * 2
}
