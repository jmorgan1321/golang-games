package serialization

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jmorgan1321/SpaceRep/internal/support"
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/engine/test"
)

func init() {
	meta.Register((*testUnmarshaller)(nil),
		meta.Init(func() interface{} { return &testUnmarshaller{} }),
	)
	meta.Register((*testType)(nil),
		meta.Init(func() interface{} { return &testType{} }),
	)
	meta.Register((*testJSONUnmarshaller)(nil),
		meta.Init(func() interface{} { return &testJSONUnmarshaller{} }),
	)
}

// These types will be used for testing reading in objects from file.
type (
	testType struct {
		I      int
		F      float32
		S      string
		B      bool
		Broken bool `json:"b"`
		T      *testType
	}

	testUnmarshaller struct {
		Arr []testType
	}

	testJSONUnmarshaller struct {
		Arr testType
	}
)

func TestSerialization(t *testing.T) {
	tests := []struct {
		summary string
		testFn  func(c *test.Chkr, obj, data interface{})
		data    []byte
		objPtr  interface{}
	}{
		{
			summary: "serializing uses golang's encoding/json serializer",
			objPtr:  &testType{},
			data:    []byte(`{"I": 5, "F": 3.14, "S": "string", "B": true}`),
			testFn: func(c *test.Chkr, obj, data interface{}) {
				SerializeInPlace(obj, data)

				typ := obj.(*testType)
				c.Expect(test.EQ, 5, typ.I)
				c.Expect(test.FloatEQ, 3.14, typ.F)
				c.Expect(test.EQ, "string", typ.S)
				c.Expect(test.EQ, true, typ.B)
			},
		},

		{
			summary: "tag support isn't in yet",
			objPtr:  &testType{},
			data:    []byte(`{"b": true}`),
			testFn: func(c *test.Chkr, obj, data interface{}) {
				c.Expect(test.Panic, func() { SerializeInPlace(obj, data) })
			},
		},

		{
			summary: "compound objects can be serialized in ",
			objPtr:  &testType{},
			data:    []byte(` {"T": { "I": 1} }`),
			// data:    []byte(` {"T": {"Type": "serialization.testType", "B": true } }`),
			testFn: func(c *test.Chkr, obj, data interface{}) {

				tf1 := testType{
					T: &testType{
						B: true,
					},
				}
				b, err := json.MarshalIndent(tf1, "    ", "    ")
				if err != nil {
					panic(err)
				}
				fmt.Println(string(b))
				tf2 := testType{}
				err = json.Unmarshal(b, &tf2)
				if err != nil {
					panic(err)
				}
				fmt.Println(tf2, *tf2.T)
				// SerializeInPlace(obj, data)

				// typ := obj.(*testType)
				c.Assert(test.True, false, "not finished")
				// c.Expect(test.EQ, true, typ.T.B)
			},
		},
	}
	for i, tt := range tests {
		c := test.Checker(t, test.Summary("with test %v: %v", i, tt.summary))

		holder, err := support.ReadData(tt.data)
		if err != nil {
			panic(err)
		}
		tt.testFn(c, tt.objPtr, holder)
	}
}
