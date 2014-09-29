package main

import (
	"fmt"
)

func Foo(x interface{}) {
	switch t := x.(type) {
	default:
		fmt.Println("unknown type")
	case int:
		fmt.Printf("%d\n", t)
	case string:
		fmt.Printf("%s\n", t)
	}
}
func main() {
	Foo(5)
	Foo("Five")
	Foo(true)
}
