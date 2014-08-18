package utils

// EpsilonCompare returns true is x and y are within the core's standard Epsilon
// of each other.  Ie, any value withing the range (-epsilon, epsilon) will be
// considered equal to 0,
func EpsilonCompare(x, y float32) bool {
	return x+Epsilon > y && x-Epsilon < y
}
