package main

import (
	"fmt"
)

//beg OMIT
type Foo struct {
	px  *int
	val int
}

func main() {
	foo := Foo{}

	fmt.Println(foo.px, foo.val)
	*foo.px = 1
}

//end OMIT
