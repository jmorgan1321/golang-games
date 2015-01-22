package vec2

import (
	"math"

	"github.com/jmorgan1321/golang-games/lib/engine/meta"
)

func init() {
	meta.Register((*Vector)(nil),
		meta.Init(func() interface{} { return &Vector{} }),
	)
}

// Vector represents a force or axis that spans infinitely along in a direction.
type Vector struct {
	X, Y float32
}

// Add takes a vector a returns a new vector representing the sum of each vectors components.
func (v Vector) Add(u Vector) Vector {
	return Vector{v.X + u.X, v.Y + u.Y}
}

// Sub takes a vector a returns a new vector representing the difference of each vectors components.
func (v Vector) Sub(u Vector) Vector {
	return v.Add(u.Neg())
}

// Neg returns a new vector travelling in the exact opposite direction.
func (v Vector) Neg() Vector {
	return Vector{-v.X, -v.Y}
}

// Mul returns a new vector equal to the old vector scaled by a certain amount.
func (v Vector) Mul(r float32) Vector {
	return Vector{v.X * r, v.Y * r}
}

// Div returns a new vector equal to the old vector scaled by a certain amount.
func (v Vector) Div(r float32) Vector {
	if r == 0 {
		panic("div by zero")
	}
	return v.Mul(1 / r)
}

// Dot returns a velue between [0..1] or [perpendicular to parallel].
func (v Vector) Dot(u Vector) float32 {
	return (v.X*u.X + v.Y*u.Y) / v.Mag()
}

// Mag returns the magnitured of v.
func (v Vector) Mag() float32 {
	return float32(math.Sqrt(float64(v.Mag2())))
}

// Mag2 returns the magnitude before squaring it.  This is a useful shortcut
// that can be used to optimize v.Dot(u) iff v and u are both already normed.
//
func (v Vector) Mag2() float32 {
	return v.X*v.X + v.Y*v.Y
}

// Norm returns a "normalized" vector where the magnitured of v is equal to 1.
func (v Vector) Norm() Vector {
	return v.Div(v.Mag())
}
