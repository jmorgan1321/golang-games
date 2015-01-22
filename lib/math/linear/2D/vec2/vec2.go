package vec2

var (
	Origin = Vector{0, 0}
	X      = Vector{1, 0}
	Y      = Vector{0, 1}
)

// // Convert returns a Point from various other math types, as approproate.
// // Convert panics if the conversion doesn't make sense.
// //
// func Convert(iface interface{}) Vector {
// 	switch i := iface.(type) {
// 	default:
// 		panic(i)
// 	case vec2.Vector:
// 		return Point{i.X, i.Y}
// 	case *vec2.Vector:
// 		return Point{i.X, i.Y}
// 	}
// }
