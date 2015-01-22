package gamecore

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jmorgan1321/golang-games/core/debug"
)

var (
	g_gameCore *core
)

type core struct {
	RootSpace JGameObjectHandle
	managers  []JManager
	spaces    []JSpace
	// Data
	TimeStamp time.Time
}

func New() *core {
	return nil
}

// JGameObjectHandle wraps a game object and makes sure that the object is valid
// before making any calls into it.
//
type JGameObjectHandle struct {
	obj JGameObject
	id  int
}

// func (o JGameObjectHandle) TriggerEvent(...) JGameObject {
// if o.IsValid() {
//     o.Get().TriggerEvent(...)
// }
// }
func (o JGameObjectHandle) Get() JGameObject {
	// return factory.GetById(o.id)
}
func (o JGameObjectHandle) ID() int {
	return o.id
}

func LoadConfig(file string) JGameObjectHandle {
	return JGameObjectHandle{}
}

func (c *core) StartUp(config JGameObjectHandle) {
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
		c.spaces[i].DeInit()
	}

	for i := len(c.managers) - 1; i >= 0; i-- {
		c.managers[i].ShutDown()
	}
}

func (c *core) Step() {
	defer debug.Trace().UnTrace()

	if c.TimeStamp.IsZero() {
		c.TimeStamp = time.Now()
	}

	if dt := time.Since(c.TimeStamp); dt >= 16*time.Millisecond {
		fmt.Println("frame:", c.CurrFrame)
		c.CurrFrame++

		for _, mgr := range c.managers {
			mgr.BeginFrame()
			defer mgr.EndFrame()
		}

		for _, spc := range c.spaces {
			spc.Update(float32(dt))
		}
		fmt.Println()
	}
}

func (c *Core) RegisterSpaces(spaces ...JSpace) {
	c.spaces = append(c.spaces, spaces...)
}
func (c *Core) RegisterManagers(mgrs ...JManager) {
	c.managers = append(c.managers, mgrs...)
}
func (c *Core) Manager(name string) JManager {
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
