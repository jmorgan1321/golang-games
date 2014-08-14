package support

import (
	"fmt"
	"github.com/jmorgan1321/golang-games/core/debug"
)

var indent = "\t\t"

func LogError(msg string, err error) error {
	fmt.Println(debug.IndentationLevel, indent, "error:", err)
	return err
}

func LogFatal(msg string, v ...interface{}) {
	s := fmt.Sprintf("%s"+indent+msg+"\n", append([]interface{}{debug.IndentationLevel}, v...)...)
	panic(s)
}

func Log(msg string, v ...interface{}) {
	fmt.Printf("%s"+indent+msg+"\n", append([]interface{}{debug.IndentationLevel}, v...)...)
}
