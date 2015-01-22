package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/engine/test"
)

func init() {
	meta.Register((*testCmp1)(nil),
		meta.Init(func() interface{} { return &testCmp1{} }),
	)
	meta.Register((*testCmpUnmarshaller)(nil),
		meta.Init(func() interface{} { return &testCmpUnmarshaller{} }),
	)
}

// These types will be used for testing reading in objects from file.
type (
	testCmp1 struct {
		JBaseObject
		I      int
		F      float32
		S      string
		B      bool
		Broken bool `json:"b"`
		T      JComponent
	}

	testCmpUnmarshaller struct {
		Arr []JComponent
		I   int
	}

	testJSONUnmarshaller struct {
		Arr []JComponent
	}
)

func (t *testCmp1) Init()   {}
func (t *testCmp1) Deinit() {}

type TT1 struct {
	TT2
	I int
}
type TT2 struct {
	S_ string `json:"S"`
}

func (t *TT2) S() string {
	return t.S_
}
func (t *TT2) SetS(s string) {
	t.S_ = s
}

type stringholder struct {
	S_ string
}

func TestGoc_Serialization(t *testing.T) {
	c := test.Checker(t)
	// tt1 := TT1{
	// 	// TT2: TT2{
	// 	// 	S: "jeff",
	// 	// },
	// 	I: 13,
	// }
	// tt1.SetS("jeff")
	// fmt.Println("name:", tt1.S())
	tt2 := &TT2{}
	tt2.SetS("jeff")
	fmt.Println("S:", tt2.S())
	b, err := json.MarshalIndent(tt2, "    ", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// tf1 := testCmp1{
	// 	T: &testCmp1{
	// 		B: true,
	// 		S: "string",
	// 	},
	// }
	// tf1.Name_ = "name"
	// tf1.SetName("john")
	// b, err := json.MarshalIndent(tf1, "    ", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b))

	// err = json.Unmarshal(b, &tf2)
	// 	var tf2 = testCmp1
	// 	holder, err := support.ReadData([]byte(`
	// {
	//     "Type": "types.TestCmp1",
	//     "B": true,
	//     "T" : {
	//         "Type": "types.TestCmp1",
	//         "S": "string"
	//     }
	// }
	//         `))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	serialization.SerializeInPlace(&tf2, holder)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// fmt.Println("data:", *tf2.(*testCmp1), *tf2.(*testCmp1).T.(*testCmp1))

	c.Assert(test.True, false, "not finished")
}

// func TestGoc_Serialization(t *testing.T) {
// 	tests := []struct {
// 		summary string
// 		testFn  func(c *test.Chkr, obj, data interface{})
// 		data    []byte
// 		objPtr  interface{}
// 	}{
// 		{
// 			summary: "compound objects can be serialized in ",
// 			objPtr:  &testCmp1{},
// 			data:    []byte(` {"T": { "I": 1} }`),

// 			testFn: func(c *test.Chkr, obj, data interface{}) {

// 				tf1 := testCmp1{
// 					T: &testCmp1{
// 						B: true,
// 					},
// 				}
// 				b, err := json.MarshalIndent(tf1, "    ", "    ")
// 				if err != nil {
// 					panic(err)
// 				}
// 				fmt.Println(string(b))
// 				tf2 := testCmp1{}
// 				err = json.Unmarshal(b, &tf2)
// 				if err != nil {
// 					panic(err)
// 				}
// 				fmt.Println(tf2, tf2.T)

// 				c.Assert(test.True, false, "not finished")
// 			},
// 		},
// 	}
// 	for i, tt := range tests {
// 		c := test.Checker(t, test.Summary("with test %v: %v", i, tt.summary))

// 		holder, err := support.ReadData(tt.data)
// 		if err != nil {
// 			panic(err)
// 		}
// 		tt.testFn(c, tt.objPtr, holder)
// 	}
// }
