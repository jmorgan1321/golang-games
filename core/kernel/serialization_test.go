package kernel

import (
	"reflect"
	"testing"

	"github.com/jmorgan1321/golang-games/core/support"
	"github.com/jmorgan1321/golang-games/core/test"
)

type serTestComp1 struct {
	OwnerMngr
	X, Y, Z int
}
type serTestComp2 struct {
	OwnerMngr
	Shape string
}

func TestSerialize_Goc(t *testing.T) {
	// test setup
	CoreFactoryFunc = testFactoryFunc

	data := []byte(
		`{
        "Type": "Goc",
	    "Name": "camera",
        "Components": [
          {
            "Type": "CameraComponent",
            "X": 100,
            "Y": 50
          },
          {
            "Type": "ColliderComponent",
            "Shape": "aabb"
          }
        ]
      }`)

	holder, _ := support.ReadData(data)

	goc := Goc{}
	SerializeInPlace(&goc, holder)
	test.ExpectEQ(t, "camera", goc.Name, "goc.Name was set correctly")

	cmp1 := goc.GetComp("serTestComp1").(*serTestComp1)
	test.ExpectEQ(t, 100, cmp1.X, "cmp1.X was set correctly")
	test.ExpectEQ(t, 50, cmp1.Y, "cmp1.Y was set correctly")
	test.ExpectEQ(t, 0, cmp1.Z, "cmp1.Z was set correctly")

	cmp2 := goc.GetComp("serTestComp2").(*serTestComp2)
	test.ExpectEQ(t, "aabb", cmp2.Shape, "cmp2.Shape was set correctly")
}

type serTestSpace struct {
	Name string
	Gocs []*Goc
}

func (*serTestSpace) Init() {
}
func (*serTestSpace) DeInit() {
}
func (*serTestSpace) Update(dt float32) {
}
func (ts *serTestSpace) AddGoc(goc *Goc) {
	ts.Gocs = append(ts.Gocs, goc)
}
func (ts *serTestSpace) Unmarshal(data interface{}) {
	m := data.(map[string]interface{})

	// Unmarshall all "normal" fields
	for k, v := range m {
		if k == "Type" || k == "Gocs" {
			continue
		}

		f := reflect.ValueOf(ts).Elem().FieldByName(k)
		SerializeInPlace(f, v)
	}

	// TODO: pull this out into GocManager (so every type of
	//       space doesn't have to re-implement this).
	// Special case handling for `Gocs`
	for _, v := range m["Gocs"].([]interface{}) {
		objData := v.(map[string]interface{})
		typename, _ := objData["Type"]
		obj := factoryFunc(typename.(string)).(*Goc)
		SerializeInPlace(obj, objData)
		ts.AddGoc(obj)
	}
}

func testFactoryFunc(name string) interface{} {
	switch name {
	case "serTestComp1", "CameraComponent":
		return &serTestComp1{}
	case "serTestComp2", "ColliderComponent":
		return &serTestComp2{}
	case "serTestSpace":
		return &serTestSpace{}
	}
	return nil
}
