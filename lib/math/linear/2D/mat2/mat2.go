package mat2

import (
	"fmt"

	"github.com/jmorgan1321/golang-games/lib/math/linear/2D/pnt2"
	"github.com/jmorgan1321/golang-games/lib/math/linear/2D/vec2"
)

var (
	Identity = Construct([]float32{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	})
	Zero = Matrix{}
)

// Col represents a col of matrix
type Col [3]float32

func MakeCol(iface interface{}) Col {
	switch iface := iface.(type) {
	default:
		panic(fmt.Sprintf("invalid type passed into MakeCol: %T", iface))
	case vec2.Vector:
		return Col{iface.X, iface.Y, 0}
	case *vec2.Vector:
		return Col{iface.X, iface.Y, 0}
	case pnt2.Point:
		return Col{iface.X, iface.Y, 1}
	case *pnt2.Point:
		return Col{iface.X, iface.Y, 1}
	case Row:
		return Col{iface[0], iface[1], iface[2]}
	}
	return Col{}
}

// Row represents a row of a matrix.
type Row [3]float32

func MakeRow(iface interface{}) Row {
	switch iface := iface.(type) {
	default:
		panic(fmt.Sprintf("invalid type passed into MakeRow: %T", iface))
	case vec2.Vector:
		return Row{iface.X, iface.Y, 0}
	case *vec2.Vector:
		return Row{iface.X, iface.Y, 0}
	case pnt2.Point:
		return Row{iface.X, iface.Y, 1}
	case *pnt2.Point:
		return Row{iface.X, iface.Y, 1}
	case Col:
		return Row{iface[0], iface[1], iface[2]}
	}
	return Row{}
}

// Construct allows us to create matrices in various ways.
//
// Ie, we can use:
//      []float32:
//          m := Construct([]float32{
//              1, 0, 0,
//              0, 1, 0,
//              0, 0, 1
//          })
//
//      []Col:
//          m := Construct([]Col{ c1, c2, c3 })
//
//      []Row:
//          m := Construct([]Row{ r1, r2, r3 })
//
//      and more. Check TestMatrix_Construct and
//      TestMatrix_ConstructPanicsOnInvalidTypes for more.
//
func Construct(iface interface{}) Matrix {
	switch iface := iface.(type) {
	default:
		panic(fmt.Sprintf("invalid type passed to Construct: %T", iface))
	case []float32:
		if len(iface) != 9 {
			panic("Construct(a []float32) must have len(a)==9")
		}
		m := Matrix{}
		for i := 0; i < 9; i++ {
			m[i] = iface[i]
		}
		return m
	case [9]float32:
		return Matrix(iface)
	case Matrix:
		return iface
	case *Matrix:
		return *iface
	case []Row:
		if len(iface) != 3 {
			panic("Construct(a []Row) must have len(a)==3")
		}
		m := Matrix{}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				m[i*3+j] = iface[i][j]
			}
		}
		return m
	case [3]Row:
		m := Matrix{}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				m[i*3+j] = iface[i][j]
			}
		}
		return m
	case []Col:
		if len(iface) != 3 {
			panic("Construct(a []Col) must have len(a)==3")
		}
		m := Matrix{}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				m[j*3+i] = iface[i][j]
			}
		}
		return m
	case [3]Col:
		m := Matrix{}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				m[j*3+i] = iface[i][j]
			}
		}
		return m
	}
	return Matrix{}
}
