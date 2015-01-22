// +build debug

package debug

import (
	"fmt"
	"runtime"
	"strings"
)

type dummy struct{}

func Trace() dummy {
	IndentationLevel.Increment()

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("what??")
		return dummy{}
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		fmt.Println("what??")
		return dummy{}
	}

	s := strings.Split(fn.Name(), "/")
	fmt.Printf("%s%s()\n", IndentationLevel, s[len(s)-1])

	return dummy{}
}

func (dummy) UnTrace() {
	IndentationLevel.Decrement()
}
