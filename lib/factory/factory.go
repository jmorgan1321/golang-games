package factory

import (
	"fmt"

	"github.com/jmorgan1321/golang-games/lib/engine/debug"
	"github.com/jmorgan1321/golang-games/lib/engine/event"
	"github.com/jmorgan1321/golang-games/lib/engine/serialization"
	"github.com/jmorgan1321/golang-games/lib/engine/support"
	"github.com/jmorgan1321/golang-games/lib/engine/types"
)

func New() *Factory {
	return &Factory{}
}

// type TypeId struct {
//  Type string
// }

// func GetTypeName(iface interface{}) string {
//  return reflect.TypeOf(iface).Elem().String()
// }

type Factory struct {
	types.JBaseObject
	event.InstantDispatcher

	NameMap map[string]types.JGameObject
	ObjList []types.JGameObject
}

func (f *Factory) StartUp(config types.JGameObject) {
	defer debug.Trace().UnTrace()

}
func (f *Factory) ShutDown() {
	defer debug.Trace().UnTrace()

}
func (f *Factory) BeginFrame() {
	defer debug.Trace().UnTrace()

}
func (f *Factory) EndFrame() {
	defer debug.Trace().UnTrace()
	f.endOfFrameCleanup()
}
func (f *Factory) ScriptName() string { return "factory" }

func (f *Factory) endOfFrameCleanup() {
	defer debug.Trace().UnTrace()
}
func (f *Factory) MarkObjectForDeletion(types.JGameObject) {
	defer debug.Trace().UnTrace()

}

type FactoryOpt func(types.JGameObject)

// s := f.Create(f1, Name("space"))
// goc1 := f.Create(f2, Owner(s), Name("goc1"))
// goc2 := f.Create(f2, Owner(s), Dispatcher(d))
func (f *Factory) Create(file string, opts ...FactoryOpt) types.JGameObject {
	defer debug.Trace().UnTrace()

	// read in game object data

	data, err := support.OpenFile(file)
	if err != nil {
		debug.Fatal("OpenFile (", file, "):", err)
	}

	holder, err := support.ReadData(data)
	if err != nil {
		debug.Fatal("ReadData (", file, "):", err)
	}

	m := holder.(map[string]interface{})
	typename, found := m["Type"]
	if !found {
		debug.Fatal("'Type' member not found in holder")
	}
	obj := serialization.FactoryFunc(typename.(string)).(types.JGameObject)
	serialization.SerializeInPlace(obj, holder)

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
	return obj
	// h := types.JGameObject{id: len(f.ObjList)}
	// f.NameMap[obj.Name()] = h
	// return h
}

func Name(name string) FactoryOpt {
	return func(obj types.JGameObject) {
		obj.SetName(name)
	}
}

// func AddComp(comp JComponent) FactoryOpt {
//  return func(obj JGameObject) {
//      obj.AddComps(comp)
//  }
// }
