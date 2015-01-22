package nor2

import (
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/math/linear/2D/vec2"
)

func init() {
	meta.Register((*Normal)(nil),
		meta.Init(func() interface{} { return &Normal{} }),
	)
}

// Normal represents a vector parallel to a surface in 2D space.  It is mostly
// useful for graphics applications and is it's own type because it treated
// differently from regular vectors during certain transformations.
//
type Normal struct {
	X, Y float32
}

func (n Normal) Mag() float32 {
	return vec2.Vector{n.X, n.Y}.Mag()
}
func (n Normal) Mag2() float32 {
	return vec2.Vector{n.X, n.Y}.Mag2()
}
func (n Normal) Norm() Normal {
	return convert(vec2.Vector{n.X, n.Y}.Norm())
}

// Convert returns a Normal from various other math types, as approproate.
// Convert panics if the conversion doesn't make sense.
//
func convert(iface interface{}) Normal {
	switch i := iface.(type) {
	default:
		panic(i)
	case vec2.Vector:
		return Normal{i.X, i.Y}
	case *vec2.Vector:
		return Normal{i.X, i.Y}
	}
}
