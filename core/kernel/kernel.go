package kernel

import (
	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/utils"
)

type FrameUpdateEvent struct {
	Dt float32
}

func (*FrameUpdateEvent) GetDelay() float32 { return 0 }

// CoreFactoryFunc allows users to extend the factory by adding in types that
// they want to serialize in.
//
// Ie:
// CoreFactoryFunc = func(string) interface{} {
//   switch name {
//   case "myNewObj", "MyNewObj":
//     return &MyNewObj{}
//   }
// }
var CoreFactoryFunc func(string) interface{}

// A Core is used to drive every system of the game.  It ticks once a frame a
// causing all other game components to fire.
//
type Core struct {
	managers []Manager
	spaces   []Space
	GameData
}

// GameData stores information about the game, like frame information and game
// state.
type GameData struct {
	CurrFrame int
	State     State
}

func New() *Core {
	debug.Trace()
	defer debug.UnTrace()

	core := &Core{}
	core.State = Running

	return core
}

func (c *Core) StartUp(config GameObject) {
	debug.Trace()
	defer debug.UnTrace()

	for _, m := range c.managers {
		m.StartUp(config)
	}

	for _, s := range c.spaces {
		s.Init()
	}
}

func (c *Core) ShutDown() {
	debug.Trace()
	defer debug.UnTrace()

	for i := len(c.managers) - 1; i >= 0; i-- {
		c.managers[i].ShutDown()
	}

	for i := len(c.spaces) - 1; i >= 0; i-- {
		c.spaces[i].DeInit()
	}
}

func (c *Core) Run() utils.ReturnCode {
	debug.Trace()
	defer debug.UnTrace()

UpdateLoop:
	for c.State == Running || c.State == Stopped {
		// support.Log("what the heck %s", c.State)
		c.GameData.CurrFrame++

		for _, mgr := range c.managers {
			mgr.BeginFrame()
			defer mgr.EndFrame()
		}

		for _, spc := range c.spaces {
			spc.Update()
		}

		if c.State == Stopped {
			break UpdateLoop
		}
	}

	switch c.State {
	case Rebooting:
		return utils.ES_Restart
	}

	return utils.ES_Success
}

func (c *Core) RegisterSpace(s Space) {
	c.spaces = append(c.spaces, s)
}

func (c *Core) RegisterManager(m Manager) {
	c.managers = append(c.managers, m)
}
