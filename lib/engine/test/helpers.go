package test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Assert fails the test and displays 'msg', if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// AssertEQ fails the test and displays 'msg', if exp is not equal to act.
func AssertEQ(tb testing.TB, exp, act interface{}, msg string) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, msg, exp, act)
		tb.FailNow()
	}
}

// Expect fails the test and displays 'msg', if the condition is false.
func Expect(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.Fail()
	}
}

// ExpectOk fails the test and displays 'err', if an err is not nil. Returns true if err is nil, false otherwise.
func ExpectOk(tb testing.TB, err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.Fail()
		return false
	}
	return true
}

// ExpectEQ fails the test and displays 'msg', if exp is not equal to act.
func ExpectEQ(tb testing.TB, exp, act interface{}, msg string, v ...interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		s := fmt.Sprintf(msg, v...)
		fmt.Printf("%s:%d: %s\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, s, exp, act)
		tb.Fail()
	}
}

// ExpectNEQ fails the test and displays 'msg', if exp is equal to act.
func ExpectNEQ(tb testing.TB, exp, act interface{}, msg string) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, msg, exp, act)
		tb.Fail()
	}
}
