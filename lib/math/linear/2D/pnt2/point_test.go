package pnt2

import (
	"testing"

	"github.com/jmorgan1321/golang-games/lib/engine/test"
	"github.com/jmorgan1321/golang-games/lib/math/linear/2D/vec2"
)

func TestPoint_DefaultState(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.EQ, Point{}, Origin)
}

func TestPnt2_convert(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.EQ, Point{2, 2}, convert(vec2.Vector{2, 2}))
	c.Expect(test.EQ, Point{1, 21}, convert(&vec2.Vector{1, 21}))
	c.Expect(test.PanicEQ, "unknown", func() { convert("unknown") })
}

func TestPoint_Add(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.EQ, Point{2, 3}, Point{1, 2}.Add(vec2.Vector{1, 1}))
}

func TestPoint_Sub(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.EQ, vec2.Vector{0, -1}, Point{1, 2}.Sub(Point{1, 1}))
}
