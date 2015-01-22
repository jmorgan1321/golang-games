package graphics

import (
	"image/color"
	"strings"
	"time"

	"github.com/jmorgan1321/golang-games/lib/core"
	"github.com/jmorgan1321/golang-games/lib/engine/debug"
	"github.com/jmorgan1321/golang-games/lib/engine/event"
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/engine/types"
	"github.com/jmorgan1321/golang-games/lib/physics"

	"gopkg.in/qml.v1"
)

type DebugDrawEvent struct {
	Dt time.Duration
}

func init() {
	meta.Register((*GoBox)(nil),
		meta.Init(func() interface{} { return &GoBox{} }),
	)
}

func NewDebugDrawer() types.JManager {
	return &QmlDebugDrawer{}
}

type QmlDebugDrawer struct {
	types.JBaseObject
	event.InstantDispatcher

	engine *qml.Engine
	window *qml.Window
	Comp   qml.Object
	Object *GoBox
	goc    types.JGameObject
}

func (q *QmlDebugDrawer) StartUp(config types.JGameObject) {
	defer debug.Trace().UnTrace()

	qml.RegisterTypes("SqclDebugDrawExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(r *GoBox, obj qml.Object) {
			r.Object = obj
		},
	}})

	cfg := config.Comp("graphics.Config").(*Config)
	file := `
import QtQuick 2.0
import SqclDebugDrawExtensions 1.0

Item {
    property var box : Component {
        Rectangle {
            opacity: 0.5;
            width: 100; height: 100;
            border.width:5;
            border.color:"blue";
        }
    }
}
    `
	component, err := cfg.Engine.LoadString("file.qml", strings.TrimSpace(file))
	if err != nil {
		panic(err)
	}
	q.Comp = component.Create(nil)
	q.engine = cfg.Engine
	q.window = cfg.Window
}
func (q *QmlDebugDrawer) ShutDown() {
	defer debug.Trace().UnTrace()
}
func (q *QmlDebugDrawer) BeginFrame() {
	defer debug.Trace().UnTrace()
}
func (q *QmlDebugDrawer) EndFrame() {
	defer debug.Trace().UnTrace()

	if q.Object != nil {
		trans := q.goc.Comp("physics.Transform").(*physics.Transform)
		q.Object.Set("x", trans.X)
		q.Object.Set("y", trans.Y)
	}
}
func (q *QmlDebugDrawer) ScriptName() string {
	return "DebugDrawer"
}

func (q *QmlDebugDrawer) AddLine(goc types.JGameObject, x1, y1, x2, y2 int) {
	defer debug.Trace().UnTrace()

	if q.Object == nil {
		q.Object = &GoBox{x1: x1, y1: y1, x2: x2, y2: y2}
		q.Object.Init()
		q.goc = goc

		input := core.Core.Manager("QmlInputMngr").(types.JManager)
		// TODO: how to unregister specific message?
		// tracker := goc.Register("collision", input, func(e event.Data) {
		goc.Register("collision", input, func(e event.Data) {
			border := q.Object.Property("border").(qml.Object)
			red := color.RGBA{255, 0, 0, 255}
			if border.Color("color") == red {
				border.Set("color", "blue")
			} else {
				border.Set("color", "red")
			}
		})
		// goc.Unregister(tracker)
		// tracker.Unlink()
	} else {
		q.Object.Deinit()
		q.Object = nil
	}
}

// func (q *QmlDebugDrawer) line(dt Time.Duration, start, end pnt2.P) {
// }
// func (q *QmlDebugDrawer) box(dt Time.Duration, ul, lr pnt2.P) {
// }

type GoBox struct {
	types.JBaseObject
	qml.Object
	x1, y1, x2, y2 int
}

func (c *GoBox) Init() {
	defer debug.Trace().UnTrace()

	grfx := core.Core.Manager("QmlDebugDrawer").(*QmlDebugDrawer)
	c.Object = grfx.Comp.Object("box").Create(nil)

	c.Set("parent", grfx.window.Root())
}
func (c *GoBox) Deinit() {
	defer debug.Trace().UnTrace()
	debug.Msg(c.x1, c.y1, c.x2, c.y2)
	c.Destroy()
}
