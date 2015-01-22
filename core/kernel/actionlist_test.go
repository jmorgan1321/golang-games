package kernel

import (
	"fmt"
	"strings"

	"github.com/jmorgan1321/golang-game/lib/engine/test"

	"testing"
)

// test types for testing action list
type (
	alistTestMoveAction struct {
		actInList
		out  *string
		X, Y int
	}
	alistTestJumpAction struct {
		actInList
		out *string
	}
	alistTestPlaySoundAction struct {
		actInList
		out   *string
		Sound string
	}
	alistTestSpinAction struct {
		actInList
		out *string
		Dir string
	}
	alistTestPrintNewLineAction struct {
		actInList
		out *string
	}

	alistTestTransformComp struct {
		OwnerMngr
		x, y int
		Z    int
	}
)

func print(msg string, dest *string) {
	if dest != nil {
		*dest += msg + "\n"
	} else {
		fmt.Println(msg)
	}
}

func (a *alistTestMoveAction) Run(obj GameObject) action {
	done := true
	trans := obj.GetComp("alistTestTransformComp").(*alistTestTransformComp)
	if a.X != trans.x {
		trans.x += 10
		done = false
	}

	if a.Y != trans.y {
		trans.y += 15
		done = false
	}

	if done {
		return a.Next()
	}

	print(fmt.Sprintf("move -> (%v, %v)", trans.x, trans.y), a.out)
	return a
}
func (a *alistTestJumpAction) Run(obj GameObject) action {
	print("jump", a.out)
	return a.Next()
}
func (a *alistTestPlaySoundAction) Run(obj GameObject) action {
	print("play sound "+a.Sound, a.out)
	return a.Next()
}
func (a *alistTestSpinAction) Run(obj GameObject) action {
	print("spin "+a.Dir, a.out)
	return a.Next()
}
func (a *alistTestPrintNewLineAction) Run(obj GameObject) action {
	print("", a.out)
	return a.Next()
}

func TestActionList(t *testing.T) {
	output := ""
	tests := []struct {
		actions []action
		output  string
	}{
		// basic case
		{
			actions: []action{
				&alistTestMoveAction{out: &output, X: 20, Y: 30},

				InParallel(
					&alistTestJumpAction{out: &output},
					&alistTestPlaySoundAction{out: &output, Sound: "jump.wav"},
				),
				&alistTestSpinAction{out: &output, Dir: "counterclock"},
			},
			output: `
move -> (10, 15)
move -> (20, 30)
jump
play sound jump.wav
spin counterclock
`,
		},

		// test parallel actions only run actions that aren't finished
		{
			actions: []action{
				InParallel(
					&alistTestMoveAction{out: &output, X: 20, Y: 30},
					&alistTestSpinAction{out: &output, Dir: "counterclock"},
				),
				&alistTestPrintNewLineAction{out: &output},
				InParallel(
					&alistTestJumpAction{out: &output},
					&alistTestPlaySoundAction{out: &output, Sound: "jump.wav"},
					&alistTestMoveAction{out: &output, X: 60, Y: 90},
					&alistTestSpinAction{out: &output, Dir: "clockwise"},
				),
				&alistTestPrintNewLineAction{out: &output},
				&alistTestSpinAction{out: &output, Dir: "counterclock"},
			},
			output: `
move -> (10, 15)
spin counterclock
move -> (20, 30)

jump
play sound jump.wav
move -> (30, 45)
spin clockwise
move -> (40, 60)
move -> (50, 75)
move -> (60, 90)

spin counterclock
`,
		},
	}
	for i, tt := range tests {
		output = ""
		alist := &ActionList{}
		goc := Goc{}
		goc.AddComps(&alistTestTransformComp{})
		goc.AddComps(alist)
		alist.Enqueue(tt.actions...)

		for !alist.IsFinished() {
			alist.Run()
		}
		test.AssertEQ(t, strings.TrimSpace(tt.output), strings.TrimSpace(output),
			fmt.Sprintf("test %v: sequence didn't run correctly", i))
	}
}
