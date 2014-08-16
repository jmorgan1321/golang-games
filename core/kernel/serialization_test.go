package kernel

import (
	"encoding/json"
	"fmt"
	"github.com/jmorgan1321/golang-games/core/support"
	"github.com/jmorgan1321/golang-games/core/test"
	"testing"
)

type serTestComp1 struct {
	X, Y, Z int
}
type serTestComp2 struct {
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

	var goc *Goc
	Serialize2(&goc, holder)
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
func (*serTestSpace) Update() {
}
func (ts *serTestSpace) AddGoc(goc *Goc) {
	ts.Gocs = append(ts.Gocs, goc)
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

func ExampleSerialize() {
	// Must set up factory functions for any types that the factory will know
	// about.  These functions will typically go into a game's driver or some
	// other game specific library (since the factory needs to know about
	// every type it can possibly create).
	//
	CoreFactoryFunc = testFactoryFunc

	exampleObjectFileData := []byte(
		`{
    "Type":"serTestSpace",
    "Name": "LevelSpace",
    "Gocs": [
        {
            "Type": "Goc",
            "Name": "Camera1",
            "Components": [
                {"Type": "CameraComponent", "X": 10, "Y": 10 },
                {"Type": "ColliderComponent", "Shape": "aabb"}
            ]
        },
        {
            "Type": "Goc",
            "Name": "Camera2",
            "Components": [
                {"Type": "CameraComponent", "X": 20, "Z": 30 },
                {"Type": "ColliderComponent", "Shape": "aabb"}
            ]
        },
        {
            "Type": "Goc",
            "Name": "Player",
            "Components": [
                {"Type": "ColliderComponent", "Shape": "circle"}
            ]
        }
    ]
}`)

	holder, _ := support.ReadData(exampleObjectFileData)

	// Since we want to fill out interfaces (and not concrete types), we must
	// pass a pointer to a pointer (in this case a pointer to the interface).
	//
	var space Space
	Serialize2(&space, holder)

	b, _ := json.MarshalIndent(space, "", "    ")
	fmt.Println(string(b))
	// Output:
	// {
	//     "Name": "LevelSpace",
	//     "Gocs": [
	//         {
	//             "Name": "Camera1",
	//             "Components": [
	//                 {
	//                     "X": 10,
	//                     "Y": 10,
	//                     "Z": 0
	//                 },
	//                 {
	//                     "Shape": "aabb"
	//                 }
	//             ]
	//         },
	//         {
	//             "Name": "Camera2",
	//             "Components": [
	//                 {
	//                     "X": 20,
	//                     "Y": 0,
	//                     "Z": 30
	//                 },
	//                 {
	//                     "Shape": "aabb"
	//                 }
	//             ]
	//         },
	//         {
	//             "Name": "Player",
	//             "Components": [
	//                 {
	//                     "Shape": "circle"
	//                 }
	//             ]
	//         }
	//     ]
	// }
}
