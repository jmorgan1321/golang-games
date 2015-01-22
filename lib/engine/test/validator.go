package test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/jmorgan1321/golang-games/lib/engine/utils"
)

type Chkr struct {
	summary string
	*testing.T
}

type option func(*Chkr)

// Summary is used by Chkr to display useful message about failing tests.
func Summary(format string, args ...interface{}) option {
	return func(c *Chkr) {
		c.summary = fmt.Sprintf(format, args...)
	}
}

// Checker returns a new test Chkr that tests assertions on a code base.
func Checker(t *testing.T, opts ...option) *Chkr {
	c := &Chkr{
		T: t,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Assert fails the test case (not allowing any other tests to run) if the
// validator, v, fails.
//
// For example, only the first test will run:
// 		c.Assert(test.EQ, true, false, "message to print with test %v", 1)
// 		c.Assert(test.EQ, true, false, "message to print with test %v", 2)
//
//
func (c Chkr) Assert(v validator, args ...interface{}) {
	if c.runTest(v, args...) {
		c.FailNow()
	}
}

// Expect fails the test case (allowing other tests to run) if the
// validator, v, fails.
//
// For example, both thest will run:
// 		c.Expect(test.EQ, true, false, "message to print with test %v", 1)
// 		c.Expect(test.EQ, true, false, "message to print with test %v", 2)
//
func (c Chkr) Expect(v validator, args ...interface{}) {
	if c.runTest(v, args...) {
		c.Fail()
	}
}

func (c Chkr) runTest(v validator, args ...interface{}) (testFailed bool) {
	if len(args) < v.argc() {
		panic(fmt.Sprintf("invalid number of args to %v", v.name()))
	}
	if v.validate(v, args[:v.argc()]...) {
		return false
	}
	_, file, line, _ := runtime.Caller(2)
	fmt.Printf("%v:%v\n", filepath.Base(file), line)
	if c.summary != "" {
		fmt.Println("\t" + c.summary)
	}
	if len(args) > v.argc()+1 {
		fmt.Printf("\t\t%v\n", fmt.Sprintf(args[v.argc()].(string), args[v.argc()+1]))
	} else if len(args) > v.argc() {
		fmt.Printf("\t\t%v\n", args[v.argc()].(string))
	}
	v.printFailure(v, args[:v.argc()]...)
	fmt.Println()

	return true
}

type validator interface {
	argc() int
	printFailure(validator, ...interface{})
	validate(validator, ...interface{}) bool
	name() string
}

type basic_validator struct {
	argc_         int
	printFailure_ func(validator, ...interface{})
	validate_     func(validator, ...interface{}) bool
	name_         string
}

func (b basic_validator) argc() int { return b.argc_ }
func (b basic_validator) printFailure(v validator, args ...interface{}) {
	b.printFailure_(v, args...)
}
func (b basic_validator) validate(v validator, args ...interface{}) bool {
	return b.validate_(v, args...)
}
func (b basic_validator) name() string { return b.name_ }

type panic_validator struct {
	basic_validator
	message interface{}
}

var (
	// EQ tests than two expresions are equal according to reflect.DeepEqual.
	//
	// Ex:
	//  	c.Assert(test.EQ, true, 1==1, "this passes")
	//  	c.Assert(test.EQ, 42, 42.0, "this faile")
	//
	EQ = &basic_validator{
		name_: "EQ",
		argc_: 2,
		printFailure_: func(v validator, args ...interface{}) {
			fmt.Printf("\t\texpected: (%T)%v\n", args[0], args[0])
			fmt.Printf("\t\treceived: (%T)%v\n", args[1], args[1])
		},
		validate_: func(v validator, args ...interface{}) bool {
			return reflect.DeepEqual(args[0], args[1])
		},
	}

	// FloatEQ tests than two numeric values are within utils.Epsilon of each
	// other.
	//
	// Ex:
	//  	c.Assert(test.FloatEQ, 42, 42.0, "this passes")
	//  	c.Assert(test.FloatEQ, 42, 42 + utils.Epsilon, "this fails")
	//  	c.Assert(test.FloatEQ, 42, 42 - utils.Epsilon, "this fails")
	//
	FloatEQ = &basic_validator{
		name_: "FloatEQ",
		argc_: 2,
		printFailure_: func(v validator, args ...interface{}) {
			fmt.Printf("\t\texpected: (%T)%v\n", args[0], args[0])
			fmt.Printf("\t\treceived: (%T)%v\n", args[1], args[1])
		},
		validate_: func(v validator, args ...interface{}) bool {
			make_float := func(iface interface{}) float32 {
				switch i := iface.(type) {
				case float32:
					return i
				case float64:
					return float32(i)
				case int:
					return float32(i)
				default:
					panic(i)
				}
			}
			return utils.EpsilonCompare(make_float(args[0]), make_float(args[1]))
		},
	}

	// NE tests than two expresions are not equal according to reflect.DeepEqual.
	//
	// Ex:
	//  	c.Assert(test.NE, true, false, "this passes")
	//  	c.Assert(test.NE, 42, 42.0, "this passes")
	//
	NE = &basic_validator{
		name_: "NE",
		argc_: 2,
		printFailure_: func(v validator, args ...interface{}) {
			fmt.Printf("\t\targ1: (%T)%v\n", args[0], args[0])
			fmt.Printf("\t\targ2: (%T)%v\n", args[1], args[1])
		},
		validate_: func(v validator, args ...interface{}) bool {
			return !reflect.DeepEqual(args[0], args[1])
		},
	}

	// True tests that an expression is true.
	//
	// Ex:
	// 		c.Assert(test.True, func()bool{ return 1==1 })
	//
	True = &basic_validator{
		name_: "True",
		argc_: 1,
		printFailure_: func(v validator, args ...interface{}) {
			fmt.Println("\t\tvalue was false")
		},
		validate_: func(v validator, args ...interface{}) bool {
			return args[0].(bool) == true
		},
	}

	// False tests that an expression is false.
	//
	// Ex:
	// 		c.Assert(test.False, 1==0)
	//
	False = &basic_validator{
		name_: "False",
		argc_: 1,
		printFailure_: func(v validator, args ...interface{}) {
			fmt.Println("\t\tvalue was true")
		},
		validate_: func(v validator, args ...interface{}) bool {
			return args[0].(bool) == false
		},
	}

	// Ok tests that any error passed in is nil.
	//
	// Ex:
	// 		v, err := foo()
	// 		c.Assert(test.Ok, err)
	//
	Ok = &basic_validator{
		name_: "Ok",
		argc_: 1,
		printFailure_: func(v validator, args ...interface{}) {
			fmt.Println("\t\tunexpected error:", args[0].(error))
		},
		validate_: func(v validator, args ...interface{}) bool {
			return args[0] == nil
		},
	}

	// TODO: is this necessary? Can't I just use NE, nil, err, instead? Or make
	//       NotNil?

	// Err tests that any error passed in is not nil.
	//
	// Ex:
	// 		v, err := foo()
	// 		c.Assert(test.Err, err)
	//
	Err = &basic_validator{
		name_: "Err",
		argc_: 1,
		printFailure_: func(v validator, args ...interface{}) {
			fmt.Println("\t\tdidn't recieve expected error")
		},
		validate_: func(v validator, args ...interface{}) bool {
			if args[0] == nil {
				return false
			}
			return args[0].(error) != nil
		},
	}

	// c.Assert(test.ErrEQ, MyCustomErr, errors.New("error"), "wrong type error should fail")

	// Panic ensures that a panic was raised during execution of an arbitrary
	// function.
	//
	// Ex:
	// 		c.Expect(test.Panic, func(){ funcThatCausesPanic() })
	//
	Panic = &basic_validator{
		name_: "Panic",
		argc_: 1,
		printFailure_: func(v validator, args ...interface{}) {
			fmt.Println("\t\texpected panic wasn't raised")
		},
		validate_: func(v validator, args ...interface{}) (ret bool) {
			defer func() {
				ret = false
				if r := recover(); r != nil {
					ret = true
				}
			}()

			f := args[0].(func())
			f()

			return
		},
	}

	// PanicEQ ensures that a specific panic was called during execution of an
	// anonymous function.
	//
	// Ex:
	// 		c.Expect(test.PanicEQ, "panic message", func(){ funcThatCausesPanic() })
	// 		c.Expect(test.PanicEQ, 3, func(){ panic(3) })
	//
	PanicEQ = &panic_validator{
		basic_validator: basic_validator{
			name_: "PanicEQ",
			argc_: 2,
			printFailure_: func(v validator, args ...interface{}) {
				fmt.Println("\t\tfunction did not cause anticipated panic")
				p := v.(*panic_validator)
				if p.message != nil {
					fmt.Printf("\t\texpected: (%T)%v\n", args[0], args[0])
					fmt.Printf("\t\treceived: (%T)%v\n", p.message, p.message)
					p.message = nil
				}
			},
			validate_: func(v validator, args ...interface{}) (ret bool) {
				defer func() {
					ret = false
					if r := recover(); r != nil {
						v.(*panic_validator).message = r
						ret = reflect.DeepEqual(args[0], r)
					}
				}()

				f := args[1].(func())
				f()

				return
			},
		},
	}
)
