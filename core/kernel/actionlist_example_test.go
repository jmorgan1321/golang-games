package kernel

import "github.com/jmorgan1321/golang-game/lib/engine/support"

func init() {
	RegisterType((*alistTestJumpAction)(nil), func() interface{} {
		return &alistTestJumpAction{}
	})
	RegisterType((*alistTestMoveAction)(nil), func() interface{} {
		return &alistTestMoveAction{}
	})
	RegisterType((*alistTestSpinAction)(nil), func() interface{} {
		return &alistTestSpinAction{}
	})
	RegisterType((*alistTestPlaySoundAction)(nil), func() interface{} {
		return &alistTestPlaySoundAction{}
	})
	RegisterType((*alistTestPrintNewLineAction)(nil), func() interface{} {
		return &alistTestPrintNewLineAction{}
	})
	RegisterType((*alistTestTransformComp)(nil), func() interface{} {
		return &alistTestTransformComp{}
	})
}

func Example_actionList() {
	exampleObjectFileData := []byte(`
{
"Type": "Goc",
"Components": [
    {
        "Type": "kernel.ActionList",
        "Sequence": [
            {
                "Type": "kernel.alistTestMoveAction",
                "X": 20,
                "Y": 30
            },

            {"Type": "kernel.alistTestPrintNewLineAction"},

            {
                "Type": "kernel.ParallelAction",
                "Actions": [
                    {
                        "Type": "kernel.alistTestJumpAction"
                    },
                    {
                        "Type": "kernel.alistTestPlaySoundAction",
                        "Sound": "jump.wav"
                    }
                ]
            },

            {"Type": "kernel.alistTestPrintNewLineAction"},

            {
                "Type": "kernel.alistTestSpinAction",
                "Dir": "counterclock"
            }
        ]
    },

    {"Type": "kernel.alistTestTransformComp" }
]
}`)

	goc := Goc{}
	holder, _ := support.ReadData(exampleObjectFileData)
	SerializeInPlace(&goc, holder)
	alist := goc.GetComp("ActionList").(*ActionList)

	for !alist.IsFinished() {
		alist.Run()
	}

	// Output:
	// move -> (10, 15)
	// move -> (20, 30)
	//
	// jump
	// play sound jump.wav
	//
	// spin counterclock
}
