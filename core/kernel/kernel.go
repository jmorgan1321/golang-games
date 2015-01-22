package kernel

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/utils"
)

// FrameUpdateEvent is dispatched by the core, once per frame.
//
// It can be used by game entities that need to know when the frame is
// has ticked (and for how long).
type FrameUpdateEvent struct {
	Dt float32
}

// TODO: remove CoreFactoryFunc for CoreFactoryMap
//
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

var coreFactoryMap = map[string]func() interface{}{}

// RegisterType allows users to extend the factory by adding in types that
// they want to serialize in.
//
// Ie:
// func init() {
//     kernel.RegisterType((*ActionList)(nil), func()interface{ return &ActionList{}})
// }
//
func RegisterType(iface interface{}, f func() interface{}) {
	coreFactoryMap[reflect.TypeOf(iface).Elem().String()] = f
}

type TypeId struct {
	Type string
}

func GetTypeName(iface interface{}) string {
	return reflect.TypeOf(iface).Elem().String()
}

// A Core is used to drive every system of the game.  It ticks once a frame a
// causing all other game components to fire.
//
type Core struct {
	managers []Manager
	spaces   []Space
	GameData
	RootSpace Space
}

// GameData stores information about the game, like frame information and game
// state.
type GameData struct {
	CurrFrame int
	State     State
	TimeStamp time.Time
}

func New() *Core {
	defer debug.Trace().UnTrace()

	core := &Core{}
	core.State = Running
	return core
}

// StartUp passes an initliazer to all managers on start up, before intializing
// all registered spaces.
//
func (c *Core) StartUp(config GameObject) {
	defer debug.Trace().UnTrace()

	for _, m := range c.managers {
		m.StartUp(config)
	}

	for _, s := range c.spaces {
		s.Init()
	}
}

// Shutdown deinits all spaces and then deinits all managers (in LIFO order)
func (c *Core) ShutDown() {
	defer debug.Trace().UnTrace()

	for i := len(c.spaces) - 1; i >= 0; i-- {
		c.spaces[i].DeInit()
	}

	for i := len(c.managers) - 1; i >= 0; i-- {
		c.managers[i].ShutDown()
	}
}

func (c *Core) Run() utils.ReturnCode {
	defer debug.Trace().UnTrace()

	framesPerSec := time.Duration(int(1e9) / 30)
	clk := time.NewTicker(framesPerSec)
	prevTime := time.Now()

UpdateLoop:
	for {
		select {
		case currTime := <-clk.C:
			if c.State != Running && c.State != Stopped {
				break UpdateLoop
			}

			c.GameData.CurrFrame++

			for _, mgr := range c.managers {
				mgr.BeginFrame()
				defer mgr.EndFrame()
			}

			for _, spc := range c.spaces {
				dt := float32(currTime.Sub(prevTime).Seconds())
				spc.Update(dt)
			}

			// TODO: make a public stopped channel, once gamecore is public
			if c.State == Stopped {
				break UpdateLoop
			}

			prevTime = currTime
		}
	}

	switch c.State {
	case Rebooting:
		return utils.ES_Restart
	}

	return utils.ES_Success
}

func (c *Core) Step() {
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

// 	debug.Trace()
// 	defer debug.UnTrace()

// 	// framesPerSec := time.Duration(int(1e9) / 30)
// 	// clk := time.NewTicker(framesPerSec)
// 	// prevTime := time.Now()

// 	// UpdateLoop:
// 	// for {
// 	// select {
// 	// case currTime := <-clk.C:
// 	// if c.State != Running && c.State != Stopped {
// 	// 	// break UpdateLoop
// 	// }

// 	c.GameData.CurrFrame++

// 	for _, mgr := range c.managers {
// 		mgr.BeginFrame()
// 		defer mgr.EndFrame()
// 	}

// 	for _, spc := range c.spaces {
// 		dt := float32(currTime.Sub(prevTime).Seconds())
// 		spc.Update(dt)
// 	}

// 	// // TODO: make a public stopped channel, once gamecore is public
// 	// if c.State == Stopped {
// 	// 	break UpdateLoop
// 	// }

// 	// prevTime = currTime
// 	// }
// 	// }

// 	// switch c.State {
// 	// case Rebooting:
// 	// 	return utils.ES_Restart
// 	// }

// 	// return utils.ES_Success
// }

func (c *Core) RegisterSpace(s Space) {
	c.spaces = append(c.spaces, s)
}

func (c *Core) RegisterManager(m Manager) {
	c.managers = append(c.managers, m)
}

func (c *Core) RegisterManagers(mgrs ...Manager) {
	c.managers = append(c.managers, mgrs...)
}
func (c *Core) Manager(name string) Manager {
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
