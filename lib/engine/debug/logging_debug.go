// +build debug

package debug

import "fmt"

func Msg(v ...interface{}) {
	fmt.Println(append([]interface{}{IndentationLevel}, v...)...)
}

func MsgF(msg string, v ...interface{}) {
	fmt.Printf("%s"+msg+"\n", append([]interface{}{IndentationLevel}, v...)...)
}
