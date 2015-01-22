package core

import (
	"reflect"
	"time"

	"github.com/jmorgan1321/golang-games/lib/engine/debug"
	"github.com/jmorgan1321/golang-games/lib/engine/serialization"
	"github.com/jmorgan1321/golang-games/lib/engine/support"
	"github.com/jmorgan1321/golang-games/lib/engine/types"
	"github.com/jmorgan1321/golang-games/lib/factory"

	"gopkg.in/qml.v1"
)

var (
	Core *core
)

type DebugDrawer interface {
	types.JManager
	AddLine(goc types.JGameObject, x1, y1, x2, y2 int)
}

type core struct {
	managers []types.JManager
	spaces   []types.JSpace

	// Game State
	RootSpace types.JGameObject
	Factory   *factory.Factory
	DebugDraw DebugDrawer

	Context *qml.Context

	// Data
	TimeStamp time.Time
	CurrFrame int
}

func New() *core {
	return &core{}
}

// LoadConfig is needed as a backdoor to factory not being inititalized
func LoadConfig(file string) types.JGameObject {
	defer debug.Trace().UnTrace()

	data, err := support.OpenFile(file)
	if err != nil {
		debug.Fatal("Failed to open Config file: " + file)
	}

	holder, err := support.ReadData(data)
	if err != nil {
		debug.Fatal("Failed to read in Config file: " + file)
		return nil
	}

	goc := types.JGoc{}
	serialization.SerializeInPlace(&goc, holder)
	return &goc
}

func (c *core) StartUp(config types.JGameObject) {
	defer debug.Trace().UnTrace()

	for _, m := range c.managers {
		m.StartUp(config)
	}

	for _, s := range c.spaces {
		s.Init()
	}
}

func (c *core) ShutDown() {
	defer debug.Trace().UnTrace()

	for i := len(c.spaces) - 1; i >= 0; i-- {
		c.spaces[i].Deinit()
	}

	for i := len(c.managers) - 1; i >= 0; i-- {
		c.managers[i].ShutDown()
	}
}

func (c *core) Step() {
	defer debug.Trace().UnTrace()

	c.TimeStamp = time.Now()
	dt := time.Since(c.TimeStamp)
	// if dt < 16*time.Millisecond { return }
	debug.Msg()
	debug.Msg("frame:", c.CurrFrame)
	debug.Log("dt:", dt)
	c.CurrFrame++

	for _, mgr := range c.managers {
		mgr.BeginFrame()
		defer mgr.EndFrame()
	}

	for _, spc := range c.spaces {
		spc.Update(dt)
	}

	debug.Msg()
	// }
}

func (c *core) RegisterSpaces(spaces ...types.JSpace) {
	c.spaces = append(c.spaces, spaces...)
}
func (c *core) RegisterManagers(mgrs ...types.JManager) {
	for _, m := range mgrs {
		c.Context.SetVar(m.ScriptName(), m)
		c.managers = append(c.managers, m)
	}
}
func (c *core) Manager(name string) types.JManager {
	for _, mgr := range c.managers {
		// TODO: possibly switch this to use package.Name
		mtype := reflect.ValueOf(mgr).Type().Elem()
		if mtype.Name() == name {
			return mgr
		}
	}
	// TODO: better error message when component isn't found
	return nil
}
