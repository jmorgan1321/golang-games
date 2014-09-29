package main

import (
	"fmt"
	"github.com/jmorgan1321/golang-games/slides/foo"
)

func main() {
	// in another pkg, 'last_foo' can't be seen. Values are accessed through package
	fmt.Println(foo.Bar, foo.Baz)
	// fmt.Println(foo.max_few)
}
