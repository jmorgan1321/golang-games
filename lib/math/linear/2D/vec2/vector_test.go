package vec2

import (
	"testing"

	"github.com/jmorgan1321/SpaceRep/internal/support"
	"github.com/jmorgan1321/golang-games/lib/engine/serialization"
	"github.com/jmorgan1321/golang-games/lib/engine/test"
)

func TestVector_DefaultState(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.EQ, Vector{}, Origin)
}

func TestVector_Neg(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.EQ, Vector{-1, 0}, X.Neg())
	c.Expect(test.EQ, Vector{0, -1}, Y.Neg())
	c.Expect(test.EQ, Vector{-1, -1}, X.Add(Y).Neg())
}

func TestVector_Add(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.EQ, Vector{2, 1}, Origin.Add(Vector{2, 1}))
	c.Expect(test.EQ, Vector{2, 0}, X.Add(X))
	c.Expect(test.EQ, Vector{1, 1}, X.Add(Y))
	c.Expect(test.EQ, Vector{1, 1}, Y.Add(X))
}

func TestVector_Sub(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.EQ, Vector{-2, -1}, Origin.Sub(Vector{2, 1}))
	c.Expect(test.EQ, Vector{0, 0}, X.Sub(X))
	c.Expect(test.EQ, Vector{1, -1}, X.Sub(Y))
	c.Expect(test.EQ, Vector{-1, 1}, Y.Sub(X))
}

func TestVector_Mul(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.EQ, Origin, Origin.Mul(5))
	c.Expect(test.EQ, X, X.Mul(1))
	c.Expect(test.EQ, Vector{2, 0}, X.Mul(2))
	c.Expect(test.EQ, Vector{-5, 0}, X.Mul(-5))
	c.Expect(test.EQ, Origin, Y.Mul(0))
	c.Expect(test.EQ, Vector{0, 1}, Y.Mul(1))
}

func TestVector_Div(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.EQ, Origin, Origin.Div(5))
	c.Expect(test.PanicEQ, "div by zero", func() { Origin.Div(0) })
	c.Expect(test.EQ, Vector{.5, 0}, X.Div(2))
	c.Expect(test.EQ, Vector{-1, 0}, X.Div(-1))
	c.Expect(test.EQ, Vector{0, .2}, Y.Div(5))
	c.Expect(test.EQ, Y, Y.Div(1))
}

func TestVector_Mag(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.FloatEQ, 0.0, Origin.Mag())
	c.Expect(test.FloatEQ, 1.0, Y.Mag())
	c.Expect(test.FloatEQ, 5.0, Vector{3, 4}.Mag())
}

func TestVector_Mag2(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.FloatEQ, 0.0, Origin.Mag2())
	c.Expect(test.FloatEQ, 1.0, Y.Mag2())
	c.Expect(test.FloatEQ, 25.0, Vector{3, 4}.Mag2())
}

func TestVector_Norm(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.EQ, Y, Y.Norm())
	c.Expect(test.EQ, X, X.Norm())
	c.Expect(test.EQ, X, Vector{5, 0}.Norm())
	c.Expect(test.EQ, Vector{.70710677, .70710677}, Vector{1, 1}.Norm())
}

func TestVector_Dot(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.FloatEQ, 1.0, Y.Dot(Y))
	c.Expect(test.FloatEQ, 0.0, Y.Dot(X))
}

func TestVector_Unmarshal(t *testing.T) {
	b := []byte(`{"Type": "vec2.Vector", "X": 25, "Y": 12}`)
	holder, err := support.ReadData(b)
	if err != nil {
		panic(err)
	}
	v := Vector{}
	serialization.SerializeInPlace(&v, holder)

	c := test.Checker(t)
	c.Expect(test.EQ, Vector{25, 12}, v)
}
