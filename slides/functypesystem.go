package main

import (
	"fmt"
)

//beg show A OMIT
func main() {
	var f MessagingFunc
	f = func(i int, x interface{}) {
		fmt.Println("got called with:", i)
	}
	fmt.Println(f) // HLfun
	f(5, nil)      // HLfun
}

type MessagingFunc func(msg int, data interface{}) // HLfun

// adding function to a ... function? o_O
func (MessagingFunc) String() string { // HLfun
	return "messaging_func"
}

//end show A OMIT
