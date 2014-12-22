package kernel

import (
	"encoding/json"
	"fmt"

	"github.com/jmorgan1321/golang-games/core/support"
)

func Example_serialize() {
	// Must set up factory functions for any types that the factory will know
	// about.  These functions will typically go into a game's driver or some
	// other game specific library (since the factory needs to know about
	// every type it can possibly create).
	//
	CoreFactoryFunc = func(name string) interface{} {
		switch name {
		case "CameraComponent":
			return &serTestComp1{}
		case "ColliderComponent":
			return &serTestComp2{}
		case "serTestSpace":
			return &serTestSpace{}
		}
		return nil
	}

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

	space := serTestSpace{}
	SerializeInPlace(&space, holder)

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
