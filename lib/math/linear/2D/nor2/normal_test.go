package nor2

import (
	"testing"

	"github.com/jmorgan1321/golang-games/lib/engine/test"
	"github.com/jmorgan1321/golang-games/lib/math/linear/2D/vec2"
)

func TestNor2_convert(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.EQ, Normal{2, 2}, convert(vec2.Vector{2, 2}))
	c.Expect(test.EQ, Normal{1, 21}, convert(&vec2.Vector{1, 21}))
	c.Expect(test.PanicEQ, "unknown", func() { convert("unknown") })
}

func TestNormal_Norm(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.EQ, Normal{0, 1}, Normal{0, 1}.Norm())
	c.Expect(test.EQ, Normal{1, 0}, Normal{1, 0}.Norm())
	c.Expect(test.EQ, Normal{1, 0}, Normal{5, 0}.Norm())
	c.Expect(test.EQ, Normal{.70710677, .70710677}, Normal{1, 1}.Norm())
}

func TestNormal_Mag(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.FloatEQ, 0.0, Normal{0, 0}.Mag())
	c.Expect(test.FloatEQ, 1.0, Normal{0, 1}.Mag())
	c.Expect(test.FloatEQ, 5.0, Normal{3, 4}.Mag())
}

func TestNormal_Mag2(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.FloatEQ, 0.0, Normal{0, 0}.Mag2())
	c.Expect(test.FloatEQ, 1.0, Normal{0, 1}.Mag2())
	c.Expect(test.FloatEQ, 25.0, Normal{3, 4}.Mag2())
}
