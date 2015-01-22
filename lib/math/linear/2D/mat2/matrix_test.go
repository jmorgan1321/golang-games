package mat2

import (
	"testing"

	"github.com/jmorgan1321/golang-games/lib/engine/test"
)

func TestMatrix_String(t *testing.T) {
	c := test.Checker(t)

	m := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c.Expect(test.EQ, "[[1 2 3] [4 5 6] [7 8 9]]", m.String())
}

func TestMatrix_Row(t *testing.T) {
	c := test.Checker(t)
	m := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c.Expect(test.EQ, Row{1, 2, 3}, m.Row(0))
	c.Expect(test.EQ, Row{4, 5, 6}, m.Row(1))
	c.Expect(test.EQ, Row{7, 8, 9}, m.Row(2))
}

func TestMatrix_Col(t *testing.T) {
	c := test.Checker(t)
	m := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c.Expect(test.EQ, Col{1, 4, 7}, m.Col(0))
	c.Expect(test.EQ, Col{2, 5, 8}, m.Col(1))
	c.Expect(test.EQ, Col{3, 6, 9}, m.Col(2))
}

func TestMatrix_SetRow(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.EQ, Row{3, 14, 19}, Zero.SetRow(0, Row{3, 14, 19}).Row(0))
	c.Expect(test.EQ, Row{9, 11, 01}, Zero.SetRow(1, Row{9, 11, 01}).Row(1))
	c.Expect(test.EQ, Row{10, 15, 25}, Zero.SetRow(2, Row{10, 15, 25}).Row(2))
}

func TestMatrix_SetCol(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.EQ, Col{1, 4, 7}, Zero.SetCol(0, Col{1, 4, 7}).Col(0))
	c.Expect(test.EQ, Col{2, 5, 8}, Zero.SetCol(1, Col{2, 5, 8}).Col(1))
	c.Expect(test.EQ, Col{3, 6, 9}, Zero.SetCol(2, Col{3, 6, 9}).Col(2))
}

func TestMatrix_Transpose(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.EQ, Matrix{1, 4, 7, 2, 5, 8, 3, 6, 9}, Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}.Transpose())
}

func TestMatrix_Det(t *testing.T) {
	c := test.Checker(t)
	c.Expect(test.FloatEQ, 1, Identity.Det())
	c.Expect(test.FloatEQ, 5, Matrix{5, 0, 0, 0, 1, 0, 0, 0, 1}.Det())
	c.Expect(test.FloatEQ, -4, Matrix{1, 2, 3, 0, 4, 5, 2, 0, 0}.Det())
}

func TestMatrix_Inverse(t *testing.T) {
	c := test.Checker(t)

	r, err := Identity.Inverse()
	c.Expect(test.Ok, err)
	c.Expect(test.EQ, Identity, r)
	r, err = Matrix{7, 2, 1, 0, 3, -1, -3, 4, -2}.Inverse()
	c.Expect(test.Ok, err)
	c.Expect(test.EQ, Matrix{-2, 8, -5, 3, -11, 7, 9, -34, 21}, r)
	r, err = Zero.Inverse()
	c.Expect(test.Err, err)
	c.Expect(test.EQ, Zero, r)
	r, err = Matrix{0, 0, 0, 1, 1, 1, 2, 3, 4}.Inverse()
	c.Expect(test.Err, err)
	c.Expect(test.EQ, Zero, r)
}
func TestMatrix_Add(t *testing.T) {
	c := test.Checker(t)
	m1 := Matrix{10, 20, 30, 40, 50, 60, 70, 80, 90}
	m2 := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := Matrix{11, 22, 33, 44, 55, 66, 77, 88, 99}
	c.Expect(test.EQ, r, m1.Add(m2))
	c.Expect(test.EQ, r, m2.Add(m1))
}

func TestMatrix_Sub(t *testing.T) {
	c := test.Checker(t)
	m1 := Matrix{10, 20, 30, 40, 50, 60, 70, 80, 90}
	m2 := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c.Expect(test.EQ, Matrix{9, 18, 27, 36, 45, 54, 63, 72, 81}, m1.Sub(m2))
	c.Expect(test.EQ, Matrix{-9, -18, -27, -36, -45, -54, -63, -72, -81}, m2.Sub(m1))
}

func TestMatrix_Mul(t *testing.T) {
	c := test.Checker(t)
	m2 := Matrix{
		0, 2, 4,
		6, 8, 7,
		5, 3, 1,
	}
	m1 := Matrix{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	c.Expect(test.EQ, Matrix{27, 27, 21, 60, 66, 57, 93, 105, 93}, m1.Mul(m2))
	c.Expect(test.EQ, Matrix{36, 42, 48, 87, 108, 129, 24, 33, 42}, m2.Mul(m1))
}
