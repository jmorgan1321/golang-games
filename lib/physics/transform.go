package physics

import (
	"github.com/jmorgan1321/golang-games/lib/core"
	"github.com/jmorgan1321/golang-games/lib/engine/debug"
	"github.com/jmorgan1321/golang-games/lib/engine/event"
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/engine/types"
)

func init() {
	meta.Register((*Transform)(nil),
		meta.Init(func() interface{} { return &Transform{} }),
	)
}

type Transform struct {
	types.JBaseObject
	X, Y int
}

func (t *Transform) Init() {
	defer debug.Trace().UnTrace()

	input := core.Core.Manager("QmlInputMngr").(types.JManager)

	t.Owner().Register("debug_draw", input, func(e event.Data) {
		t.DebugDraw()
	})
}
func (t *Transform) Deinit() {
	defer debug.Trace().UnTrace()
}

func (t *Transform) DebugDraw() {
	defer debug.Trace().UnTrace()

	core.Core.DebugDraw.AddLine(t.Owner(), t.X, t.Y, t.X+50, t.Y+50)
}
