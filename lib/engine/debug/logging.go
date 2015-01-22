package debug

import "fmt"

func Fatal(v ...interface{}) {
	s := fmt.Sprintln(append([]interface{}{IndentationLevel}, v...)...)
	panic(s)
}
func FatalF(msg string, v ...interface{}) {
	s := fmt.Sprintf("%s"+msg+"\n", append([]interface{}{IndentationLevel}, v...)...)
	panic(s)
}

func Log(v ...interface{}) {
	fmt.Println(append([]interface{}{IndentationLevel}, v...)...)
}
