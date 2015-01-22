package graphics

import (
	"github.com/jmorgan1321/golang-games/lib/core"
	"github.com/jmorgan1321/golang-games/lib/engine/debug"
	"github.com/jmorgan1321/golang-games/lib/engine/event"
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/engine/types"
	"github.com/jmorgan1321/golang-games/lib/physics"

	"gopkg.in/qml.v1"
	"gopkg.in/qml.v1/gl/2.0"
)

func init() {
	meta.Register((*Config)(nil),
		meta.Init(func() interface{} { return &Config{} }),
	)

	meta.Register((*GoRect)(nil),
		meta.Init(func() interface{} { return &GoRect{} }),
	)

}

func NewManager() types.JManager {
	return &QmlGrfxMngr{}
}

type Config struct {
	types.JBaseObject

	H, W   int
	Window *qml.Window
	Engine *qml.Engine
}

func (c *Config) Init() {
	defer debug.Trace().UnTrace()
}
func (c *Config) Deinit() {
	defer debug.Trace().UnTrace()
}

type QmlGrfxMngr struct {
	types.JBaseObject
	event.InstantDispatcher

	engine *qml.Engine
	window *qml.Window
	r      *GoRect
}

func (q *QmlGrfxMngr) StartUp(config types.JGameObject) {
	defer debug.Trace().UnTrace()

	types := []qml.TypeSpec{
		{
			Init: func(r *GoQmlRect, obj qml.Object) {
				defer debug.Trace().UnTrace()
				r.Object = obj
			},
		},
	}
	qml.RegisterTypes("GoExtensions", 1, 0, types)

	cfg := config.Comp("graphics.Config").(*Config)
	q.engine = cfg.Engine
	q.window = cfg.Window
	q.window.Set("width", cfg.W)
	q.window.Set("height", cfg.H)
}
func (q *QmlGrfxMngr) ShutDown() {
	defer debug.Trace().UnTrace()
}
func (q *QmlGrfxMngr) BeginFrame() {
	defer debug.Trace().UnTrace()
}
func (q *QmlGrfxMngr) EndFrame() {
	defer debug.Trace().UnTrace()

	if q.r != nil {
		trans := q.r.Owner().Comp("physics.Transform").(*physics.Transform)
		q.r.Set("x", trans.X)
		q.r.Set("y", trans.Y)
		// 	// qml.Changed(q.r, fieldAddr)
	}
}
func (q *QmlGrfxMngr) ScriptName() string {
	return "grfxMngr"
}

type GoRect struct {
	types.JBaseObject
	GoQmlRect
}

func (c *GoRect) Init() {
	defer debug.Trace().UnTrace()

	grfx := core.Core.Manager("QmlGrfxMngr").(*QmlGrfxMngr)
	grfx.r = c
	component, err := grfx.engine.LoadFile("file3.qml")
	if err != nil {
		panic(err)
	}
	c.Object = component.Create(nil)
	obj := c.Object.Interface().(*GoQmlRect)
	obj.R = c
	obj.Test = "what"

	c.Set("parent", grfx.window.Root())
}
func (c *GoRect) Deinit() {
	defer debug.Trace().UnTrace()
}

type GoQmlRect struct {
	qml.Object
	R    *GoRect
	Test string
}

func (c *GoQmlRect) Paint(p *qml.Painter) {
	defer debug.Trace().UnTrace()

	gl := GL.API(p)

	trans := c.R.Owner().Comp("physics.Transform").(*physics.Transform)
	c.Set("x", trans.X)
	c.Set("y", trans.Y)
	width := float32(c.Int("width"))
	height := float32(c.Int("height"))

	gl.Enable(GL.BLEND)
	gl.BlendFunc(GL.SRC_ALPHA, GL.ONE_MINUS_SRC_ALPHA)
	gl.Color4f(1.0, 1.0, 1.0, 0.8)
	gl.Begin(GL.QUADS)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(0, height)
	gl.End()

	gl.LineWidth(2.5)
	gl.Color4f(0.0, 0.0, 0.0, 1.0)
	gl.Begin(GL.LINES)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(0, height)
	gl.End()
}
