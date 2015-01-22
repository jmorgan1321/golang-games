package input

import (
	"os"

	"github.com/jmorgan1321/golang-games/lib/core"
	"github.com/jmorgan1321/golang-games/lib/engine/debug"
	"github.com/jmorgan1321/golang-games/lib/engine/event"
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/engine/types"
	"github.com/jmorgan1321/golang-games/lib/input/keycodes"
)

func init() {
	meta.Register((*Config)(nil),
		meta.Init(func() interface{} { return &Config{} }),
	)
}

func NewManager() types.JManager {
	return &QmlInputMngr{}
}

// type InputHandlerComponent struct {
//     kernel.BasicDispatcher
// }

// func (c *InputHandlerComponent) Init() {
//     krnl.RootSpace.TriggerEvent("register_input_comp", c)
// }
// func (c *InputHandlerComponent) DeInit() {
// }
// func (c *InputHandlerComponent) SetOwner(obj kernel.GameObject) {
// }

type Config struct {
	types.JBaseObject
}

func (c *Config) Init() {
	defer debug.Trace().UnTrace()
}
func (c *Config) Deinit() {
	defer debug.Trace().UnTrace()
}

type QmlInputMngr struct {
	types.JBaseObject
	event.InstantDispatcher
}

func (q *QmlInputMngr) StartUp(config types.JGameObject) {
	defer debug.Trace().UnTrace()
	// cfg := config.Comp("input.Config").(*input.Config)

	q.Register("register_input_comp", core.Core.RootSpace, func(e event.Data) {
		q.registerInputHandler(e.(*InputHandler))
	})
}
func (q *QmlInputMngr) ShutDown() {
	defer debug.Trace().UnTrace()
}
func (q *QmlInputMngr) BeginFrame() {
	defer debug.Trace().UnTrace()
}
func (q *QmlInputMngr) EndFrame() {
	defer debug.Trace().UnTrace()
}
func (q *QmlInputMngr) registerInputHandler(c *InputHandler) {
	defer debug.Trace().UnTrace()
	debug.MsgF("registering %#v\n", c)
}
func (q *QmlInputMngr) ScriptName() string {
	return "inputMngr"
}

// HandleKey handles keyboard events
func (q *QmlInputMngr) HandleKey(key int, shift_mod, ctrl_mod, alt_mod bool) {
	defer debug.Trace().UnTrace()

	switch keycodes.Keycode(key) {
	case keycodes.Key_None:
		// do nothing
	case keycodes.Key_Alt, keycodes.Key_Shift, keycodes.Key_Control:
		debug.Msg(keycodes.Keycode(key))
		debug.Msg()
	default:
		// TODO(jemorgan): move this logic into input system
		switch key {
		case keycodes.Key_Right:
			q.Dispatch("move_right", nil)
			return
		case keycodes.Key_Left:
			q.Dispatch("move_left", nil)
			return
		case keycodes.Key_Up:
			q.Dispatch("move_up", nil)
			return
		case keycodes.Key_Down:
			q.Dispatch("move_down", nil)
			return
		case keycodes.Key_D:
			q.Dispatch("debug_draw", nil)
		case keycodes.Key_F:
			q.Dispatch("collision", nil)
		}

		if shift_mod && key != keycodes.Key_Shift {
			debug.Msg("shift+")
		}
		if ctrl_mod && key != keycodes.Key_Control {
			debug.Msg("ctrl+")
		}
		if alt_mod && key != keycodes.Key_Alt {
			debug.Msg("alt+")
		}
		debug.Msg(keycodes.Keycode(key))
		debug.Msg()
	}

	if shift_mod && key == keycodes.Key_Q {
		os.Exit(1)
	}
	// if key == keycodes.Key_T {
	//     ctrl.Timer.Set("running", true)
	// }
}

func (q *QmlInputMngr) HandleMouseDown(xPos, yPos int) {
	defer debug.Trace().UnTrace()

	debug.Msg("mouse:", xPos, yPos)
}
