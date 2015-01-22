package test

import (
	"errors"
	"testing"

	"github.com/jmorgan1321/golang-games/lib/engine/utils"
)

// func TestAssert(t *testing.T) {
// 	c := NewCheckerF(t, "with test %v, testing %v:", 0, "summary")
// 	c.Assert2(EQ, true, true, "true equals true should pass")
// 	c.Assert2(EQ, true, false)
// }

// func TestExpect(t *testing.T) {
// 	c := NewCheckerF(t, "with test %v, testing %v:", 0, "summary")
// 	c.Expect2(EQ, true, true, "true equals true should pass")
// 	c.Expect2(EQ, true, false, "true equals false should fail")
// 	c.Expect2(EQ, 32, 21, "32 equals 21 should fail")
// 	c.Expect2(EQ, true, false)
// }

func TestValidators(t *testing.T) {
	tests := []struct {
		summary string
		v       validator
		args    []interface{}
		result  bool
	}{
		{
			summary: "true EQ true should pass",
			v:       EQ,
			args:    []interface{}{true, true},
			result:  true,
		},

		{
			summary: "42 EQ 42 should pass",
			v:       EQ,
			args:    []interface{}{42, 42},
			result:  true,
		},

		{
			summary: "`string` EQ `string` should pass",
			v:       EQ,
			args:    []interface{}{`string`, `string`},
			result:  true,
		},

		{
			summary: "`string1` EQ `string2` should fail",
			v:       EQ,
			args:    []interface{}{`string1`, `string2`},
			result:  false,
		},

		{
			summary: "bool EQ `string2` should fail",
			v:       EQ,
			args:    []interface{}{true, `string2`},
			result:  false,
		},

		{
			summary: "21.5 FloatEQ 21.5 should pass",
			v:       FloatEQ,
			args:    []interface{}{21.5, 21.5},
			result:  true,
		},

		{
			summary: "321.5 FloatEQ 321.5 should pass",
			v:       FloatEQ,
			args:    []interface{}{321.5, 321.5 + utils.Epsilon},
			result:  false,
		},

		{
			summary: "true NE true should fail",
			v:       NE,
			args:    []interface{}{true, true},
			result:  false,
		},

		{
			summary: "100 NE 100 should fail",
			v:       NE,
			args:    []interface{}{100, 100},
			result:  false,
		},

		{
			summary: "`string` NE `string` should fail",
			v:       NE,
			args:    []interface{}{`string`, `string`},
			result:  false,
		},

		{
			summary: "`string1` NE `string2` should pass",
			v:       NE,
			args:    []interface{}{`string1`, `string2`},
			result:  true,
		},

		{
			summary: "bool NE `string2` should pass",
			v:       NE,
			args:    []interface{}{true, `string2`},
			result:  true,
		},

		{
			summary: "True true should pass",
			v:       True,
			args:    []interface{}{true},
			result:  true,
		},

		{
			summary: "True (42==42) should pass",
			v:       True,
			args:    []interface{}{42 == 42},
			result:  true,
		},

		{
			summary: "True false should fail",
			v:       True,
			args:    []interface{}{false},
			result:  false,
		},

		{
			summary: "True (42==41) should fail",
			v:       True,
			args:    []interface{}{42 == 41},
			result:  false,
		},

		{
			summary: "False false should pass",
			v:       False,
			args:    []interface{}{false},
			result:  true,
		},

		{
			summary: "False (3.21 == 1.0) should pass",
			v:       False,
			args:    []interface{}{3.21 == 1.0},
			result:  true,
		},

		{
			summary: "False true should fail",
			v:       False,
			args:    []interface{}{true},
			result:  false,
		},

		{
			summary: "False ('string1'==`string1`) should fail",
			v:       False,
			args:    []interface{}{`string1` == `string1`},
			result:  false,
		},

		{
			summary: "Ok (error)(nil) should pass",
			v:       Ok,
			args:    []interface{}{nil},
			result:  true,
		},

		{
			summary: "Ok with non-nil error should fail",
			v:       Ok,
			args:    []interface{}{errors.New("error")},
			result:  false,
		},

		{
			summary: "Err (error)(nil) should fail",
			v:       Err,
			args:    []interface{}{nil},
			result:  false,
		},

		{
			summary: "Err with non-nil error should pass",
			v:       Err,
			args:    []interface{}{errors.New("error")},
			result:  true,
		},

		{
			summary: "Panic with function raising panic should pass",
			v:       Panic,
			args:    []interface{}{func() { panic("anything") }},
			result:  true,
		},

		{
			summary: "Panic with function that doesn't panic should fail",
			v:       Panic,
			args:    []interface{}{func() {}},
			result:  false,
		},

		{
			summary: "PanicEQ with function raising panic with correct message should pass",
			v:       PanicEQ,
			args:    []interface{}{5, func() { panic(5) }},
			result:  true,
		},

		{
			summary: "PanicEQ with function raising panic with incorrect message should fail",
			v:       PanicEQ,
			args:    []interface{}{"panic", func() { panic("incorrect") }},
			result:  false,
		},

		{
			summary: "PanicEQ with function that doesn't panic should fail",
			v:       PanicEQ,
			args:    []interface{}{"panic", func() {}},
			result:  false,
		},
	}
	for i, tt := range tests {
		if tt.v.validate(tt.v, tt.args...) != tt.result {
			t.Errorf("with test %v, testing %v", i, tt.summary)
			tt.v.printFailure(tt.v, tt.args...)
		}
	}
}
