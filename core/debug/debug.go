package debug

import (
	"fmt"
	"runtime"
	"strings"
)

var SilentMode = false

// TODO: show this to digipen
type IndentLevel int

func (l *IndentLevel) Increment() {
	*l++
}
func (l *IndentLevel) Decrement() {
	*l--
}
func (l *IndentLevel) String() string {
	s := ""
	for i := *l; i > 0; i-- {
		s += "\t"
	}
	return s
}

// TODO: move global into Env
var IndentationLevel *IndentLevel

func init() {
	IndentationLevel = new(IndentLevel)
}
func Trace() {
	if SilentMode {
		return
	}

	IndentationLevel.Increment()

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("what??")
		return
		// return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		fmt.Println("what??")
		return
		// return "unnamed"
	}

	s := strings.Split(fn.Name(), "/")
	fmt.Printf("%s%s()\n", IndentationLevel, s[len(s)-1])
}

func UnTrace() {
	if SilentMode {
		return
	}

	IndentationLevel.Decrement()
}
