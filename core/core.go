package core

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/jmorgan1321/golang-game/lib/engine/support"
	"github.com/jmorgan1321/golang-games/lib/engine/debug"
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

type Factory struct {
	NameMap map[string]JGameObjectHandle
	ObjList []JGameObject
}

func (f *Factory) StartUp(config JGameObjectHandle) {

}
func (f *Factory) ShutDown() {

}
func (f *Factory) BeginFrame() {

}
func (f *Factory) EndFrame() {
	endOfFrameCleanup()
}

func (f *Factory) endOfFrameCleanup() {
}
func (f *Factory) MarkObjectForDeletion(JGameObjectHandle) {

}

type FactoryOpt func(JGameObject)

// s := f.Create(f1, Name("space"))
// goc1 := f.Create(f2, Owner(s), Name("goc1"))
// goc2 := f.Create(f2, Owner(s), Dispatcher(d))
func (f *Factory) Create(file string, opts ...FactoryOpt) JGameObjectHandle {
	defer debug.Trace().UnTrace()

	// read in game object data

	data, err := support.OpenFile(file)
	if err != nil {
		support.LogFatal("Failed to open Config file: " + file)
	}

	holder, err := support.ReadData(data)
	if err != nil {
		support.LogFatal("Failed to read in Config file: " + file)
		return nil
	}

	m := v.(map[string]interface{})
	typename, err := m["Type"]
	if err != nil {
		log.Panic(err)
	}
	obj := factoryFunc(typename.(string))
	SerializeInPlace(obj, holder)

	// apply options to game object
	for _, opt := range opts {
		opt(obj)
	}

	// check for transform, check for dispatcher, name etc

	if obj.Name() == "" {
		obj.SetName(fmt.Sprint("obj", len(f.ObjList)))
	}

	// add game object to factory

	f.ObjList = append(f.ObjList, obj)
	h := JGameObjectHandle{id: len(f.ObjList)}
	f.NameMap[obj.Name()] = h
	return h
}

func Name(name string) FactoryOpt {
	return func(obj JGameObject) {
		obj.SetName(name)
	}
}

// func AddComp(comp JComponent) FactoryOpt {
// 	return func(obj JGameObject) {
// 		obj.AddComps(comp)
// 	}
// }
