package mat2

import (
	"errors"
	"fmt"

	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/engine/utils"
)

func init() {
	meta.Register((*Matrix)(nil),
		meta.Init(func() interface{} { return &Matrix{} }),
	)
}

// Matrix is used to store information about a linear space.  Assuming we're
// using row-major matrixes, the the cols store the following information:
// direction, right, position, and up (not in that order).
//
// We use 3x3 matrixes in 2D space so that we can store affine transformation
// information inside of the matrix.
//
type Matrix [9]float32

func (m Matrix) String() string {
	return fmt.Sprintf("[%v %v %v]", m[:3], m[3:6], m[6:9])
}

func (m Matrix) Row(i int) Row {
	b := i * 3
	return Row{m[b], m[b+1], m[b+2]}
}
func (m Matrix) SetRow(i int, r Row) Matrix {
	b := i * 3
	m[b+0] = r[0]
	m[b+1] = r[1]
	m[b+2] = r[2]
	return m
}
func (m Matrix) Col(i int) Col {
	return Col{m[i], m[i+3], m[i+6]}
}
func (m Matrix) SetCol(i int, c Col) Matrix {
	m[i+0] = c[0]
	m[i+3] = c[1]
	m[i+6] = c[2]
	return m
}

func (m Matrix) Transpose() Matrix {
	return Construct([]Row{
		MakeRow(m.Col(0)),
		MakeRow(m.Col(1)),
		MakeRow(m.Col(2)),
	})
}

func (m Matrix) Det() float32 {
	return m[0]*(m[4]*m[8]-m[7]*m[5]) - m[3]*(m[1]*m[8]-m[7]*m[2]) + m[6]*(m[1]*m[5]-m[4]*m[2])
}

func (m Matrix) Inverse() (Matrix, error) {
	det := m.Det()
	if utils.EpsilonCompare(det, 0.0) {
		return Zero, errors.New("Matrix could not be inverted")
	}
	r := Matrix{
		m[4]*m[8] - m[5]*m[7], m[2]*m[7] - m[1]*m[8], m[1]*m[5] - m[2]*m[4],
		m[5]*m[6] - m[3]*m[8], m[0]*m[8] - m[2]*m[6], m[2]*m[3] - m[0]*m[5],
		m[3]*m[7] - m[4]*m[6], m[1]*m[6] - m[0]*m[7], m[0]*m[4] - m[1]*m[3],
	}
	inv := 1 / det
	for i := 0; i < 9; i++ {
		r[i] *= inv
	}
	return r, nil
}

func (m Matrix) Add(n Matrix) Matrix {
	r := Matrix{}
	for i := 0; i < 9; i++ {
		r[i] = m[i] + n[i]
	}
	return r
}

func (m Matrix) Sub(n Matrix) Matrix {
	r := Matrix{}
	for i := 0; i < 9; i++ {
		r[i] = m[i] - n[i]
	}
	return r
}

func (m Matrix) Mul(n Matrix) Matrix {
	mr0, mr1, mr2 := m.Row(0), m.Row(1), m.Row(2)
	nc0, nc1, nc2 := n.Col(0), n.Col(1), n.Col(2)

	m[0] = mr0[0]*nc0[0] + mr0[1]*nc0[1] + mr0[2]*nc0[2]
	m[1] = mr0[0]*nc1[0] + mr0[1]*nc1[1] + mr0[2]*nc1[2]
	m[2] = mr0[0]*nc2[0] + mr0[1]*nc2[1] + mr0[2]*nc2[2]

	m[3] = mr1[0]*nc0[0] + mr1[1]*nc0[1] + mr1[2]*nc0[2]
	m[4] = mr1[0]*nc1[0] + mr1[1]*nc1[1] + mr1[2]*nc1[2]
	m[5] = mr1[0]*nc2[0] + mr1[1]*nc2[1] + mr1[2]*nc2[2]

	m[6] = mr2[0]*nc0[0] + mr2[1]*nc0[1] + mr2[2]*nc0[2]
	m[7] = mr2[0]*nc1[0] + mr2[1]*nc1[1] + mr2[2]*nc1[2]
	m[8] = mr2[0]*nc2[0] + mr2[1]*nc2[1] + mr2[2]*nc2[2]

	return m
}
