package input

import (
	"github.com/jmorgan1321/golang-games/lib/core"
	"github.com/jmorgan1321/golang-games/lib/engine/debug"
	"github.com/jmorgan1321/golang-games/lib/engine/event"
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/engine/types"
	"github.com/jmorgan1321/golang-games/lib/physics"
)

func init() {
	meta.Register((*InputHandler)(nil),
		meta.Init(func() interface{} { return &InputHandler{} }),
	)
}

type InputHandler struct {
	types.JBaseObject
}

func (c *InputHandler) Init() {
	defer debug.Trace().UnTrace()

	core.Core.RootSpace.Dispatch("register_input_comp", c)

	input := core.Core.Manager("QmlInputMngr").(*QmlInputMngr)

	c.Owner().Register("move_left", input, func(e event.Data) {
		c.movePlayer(-1, 0)
	})
	c.Owner().Register("move_right", input, func(e event.Data) {
		c.movePlayer(1, 0)
	})
	c.Owner().Register("move_down", input, func(e event.Data) {
		c.movePlayer(0, 1)
	})
	c.Owner().Register("move_up", input, func(e event.Data) {
		c.movePlayer(0, -1)
	})
}
func (c *InputHandler) Deinit() {
	defer debug.Trace().UnTrace()
}

func (c *InputHandler) movePlayer(x, y int) {
	defer debug.Trace().UnTrace()

	trans := c.Owner().Comp("physics.Transform").(*physics.Transform)
	trans.X += (10 * x)
	trans.Y += (10 * y)

	debug.Msg(trans.X, trans.Y)
}
