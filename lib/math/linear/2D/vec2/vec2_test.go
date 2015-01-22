package vec2

import (
	"testing"

	"github.com/jmorgan1321/golang-games/lib/engine/test"
)

func TestVec2_Variables(t *testing.T) {
	c := test.Checker(t, test.Summary("spec testing exported variables"))

	c.Expect(test.EQ, Vector{0, 0}, Origin, "origin")
	c.Expect(test.EQ, Vector{1, 0}, X, "x-axis")
	c.Expect(test.EQ, Vector{0, 1}, Y, "y-axis")
}

// func TestPnt2_Convert(t *testing.T) {
// 	c := test.NewChecker(t, "")
// 	c.Expect(test.EQ, Vector{2, 2}, Convert(pnt2.Point{2, 2}))
// 	c.Expect(test.EQ, Vector{1, 21}, Convert(&pnt2.Point{1, 21}))
// 	c.Expect(test.PanicEQ, "unknown", func() { Convert("unknown") })
// }
