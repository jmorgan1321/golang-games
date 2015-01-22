package mat2

import (
	"testing"

	"github.com/jmorgan1321/golang-games/lib/engine/test"
	"github.com/jmorgan1321/golang-games/lib/math/linear/2D/pnt2"
	"github.com/jmorgan1321/golang-games/lib/math/linear/2D/vec2"
)

func TestMat2_ExportedVariables(t *testing.T) {
	c := test.Checker(t)

	i := Matrix{}
	i[0] = 1
	i[4] = 1
	i[8] = 1
	c.Expect(test.EQ, i, Identity)
	c.Expect(test.EQ, Matrix{}, Zero)
}

func TestMat2_ConstructPanicsOnInvalidTypes(t *testing.T) {
	tests := []struct {
		f   func()
		msg string
	}{
		{
			msg: "invalid type passed to Construct: mat2.invalid",
			f:   func() { type invalid int; Construct(invalid(5)) },
		},

		{
			msg: "Construct(a []float32) must have len(a)==9",
			f:   func() { Construct([]float32{}) },
		},

		{
			msg: "Construct(a []float32) must have len(a)==9",
			f:   func() { Construct([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) },
		},

		{
			msg: "Construct(a []Row) must have len(a)==3",
			f:   func() { Construct([]Row{}) },
		},

		{
			msg: "Construct(a []Row) must have len(a)==3",
			f:   func() { Construct([]Row{{}, {}, {}, {}}) },
		},

		{
			msg: "Construct(a []Col) must have len(a)==3",
			f:   func() { Construct([]Col{}) },
		},

		{
			msg: "Construct(a []Col) must have len(a)==3",
			f:   func() { Construct([]Col{{}, {}, {}, {}}) },
		},
	}
	for i, tt := range tests {
		c := test.Checker(t, test.Summary("with test %v", i))
		c.Expect(test.PanicEQ, tt.msg, tt.f)
	}
}

func TestMat2_Construct(t *testing.T) {
	tests := []struct {
		summary string
		args    interface{}
		exp     Matrix
	}{
		{
			summary: "an slice of float32",
			args:    []float32{0, 1, 2, 3, 4, 5, 6, 7, 8},
			exp:     Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},

		{
			summary: "an [16]float32",
			args:    [9]float32{0, 1},
			exp:     Matrix{0, 1},
		},

		{
			summary: "a Matrix",
			args:    Matrix{7, 12, 15, -1, -1, -1, 0, 8, 7},
			exp:     Matrix{7, 12, 15, -1, -1, -1, 0, 8, 7},
		},

		{
			summary: "a *Matrix",
			args:    &Matrix{10, 9, 8, 7, 6, 5, 4, 3, 2},
			exp:     Matrix{10, 9, 8, 7, 6, 5, 4, 3, 2},
		},

		{
			summary: "a []Row",
			args:    []Row{{1, 2, 0}, {0, 1, 2}, {2, 0, 1}},
			exp:     Matrix{1, 2, 0, 0, 1, 2, 2, 0, 1},
		},

		{
			summary: "a [3]Row",
			args:    [3]Row{{1, 2, 0}, {0, 1, 2}, {2, 0, 1}},
			exp:     Matrix{1, 2, 0, 0, 1, 2, 2, 0, 1},
		},

		{
			summary: "a []Col",
			args:    []Col{{1, 2, 0}, {0, 1, 2}, {2, 0, 1}},
			exp:     Matrix{1, 0, 2, 2, 1, 0, 0, 2, 1},
		},

		{
			summary: "a [3]Col",
			args:    [3]Col{{1, 2, 0}, {0, 1, 2}, {2, 0, 1}},
			exp:     Matrix{1, 0, 2, 2, 1, 0, 0, 2, 1},
		},
	}
	for i, tt := range tests {
		c := test.Checker(t, test.Summary("with test %v: %v", i, tt.summary))
		c.Expect(test.EQ, tt.exp, Construct(tt.args))
	}
}

func TestMat2_MakeCol(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.EQ, Col{1, 2, 0}, MakeCol(vec2.Vector{1, 2}))
	c.Expect(test.EQ, Col{4, 7, 0}, MakeCol(&vec2.Vector{4, 7}))
	c.Expect(test.EQ, Col{3, 2, 1}, MakeCol(pnt2.Point{3, 2}))
	c.Expect(test.EQ, Col{1, 1, 1}, MakeCol(&pnt2.Point{1, 1}))
	c.Expect(test.EQ, Col{1, 2, 3}, MakeCol(Row{1, 2, 3}))
	type unknown int
	c.Expect(test.PanicEQ, "invalid type passed into MakeCol: mat2.unknown", func() { MakeCol(unknown(5)) })
}

func TestMat2_MakeRow(t *testing.T) {
	c := test.Checker(t)

	c.Expect(test.EQ, Row{1, 2, 0}, MakeRow(vec2.Vector{1, 2}))
	c.Expect(test.EQ, Row{4, 7, 0}, MakeRow(&vec2.Vector{4, 7}))
	c.Expect(test.EQ, Row{3, 2, 1}, MakeRow(pnt2.Point{3, 2}))
	c.Expect(test.EQ, Row{1, 1, 1}, MakeRow(&pnt2.Point{1, 1}))
	c.Expect(test.EQ, Row{1, 2, 3}, MakeRow(Col{1, 2, 3}))
	type unknown int
	c.Expect(test.PanicEQ, "invalid type passed into MakeRow: mat2.unknown", func() { MakeRow(unknown(5)) })
}
