package main

import "fmt"

func main() {
	//beg OMIT
	// https://golang.org/ref/spec#Function_literals
	// https://golang.org/doc/effective_go.html#channels (search 'closure')
	//
	// y is captured by reference. If you wanted to pass y by value,
	// you'd pass it as an additional arg to into your anonymous func.
	// ie,
	//     f := func(y, x) bool {...}
	//     f(y, 5)
	y := 5
	f := func(x int) bool {
		if x < y {
			return true
		}

		y++
		return false
	}

	b := f(5)
	fmt.Println(b, y)
	b = f(5)
	fmt.Println(b, y)
	//end OMIT
}
