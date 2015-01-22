package pnt2

import (
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/math/linear/2D/vec2"
)

func init() {
	meta.Register((*Point)(nil),
		meta.Init(func() interface{} { return &Point{} }),
	)
}

var (
	Origin = Point{0, 0}
)

// Point represents a location in 2D space.
type Point struct {
	X, Y float32
}

// Add returns the point translated by a vector.
func (p Point) Add(v vec2.Vector) Point {
	return Point{v.X + p.X, v.Y + p.Y}
}

// Sub returns the vector (point from p to q) that would translate p to q.
func (p Point) Sub(q Point) vec2.Vector {
	return vec2.Vector{q.X - p.X, q.Y - p.Y}
}

// TODO(jemorgan): move pnt2 methods into pnt2.go

// TODO(jemorgan): move into math2d package
//
// Convert returns a Point from various other math types, as approproate.
// Convert panics if the conversion doesn't make sense.
//
func convert(iface interface{}) Point {
	switch i := iface.(type) {
	default:
		panic(i)
	case vec2.Vector:
		return Point{i.X, i.Y}
	case *vec2.Vector:
		return Point{i.X, i.Y}
	}
}
