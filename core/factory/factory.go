package factory

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/support"
	"github.com/jmorgan1321/golang-games/lib/engine/types"
)

// RegisterType allows users to extend the factory by adding in types that
// they want to serialize in.
//
// Examples:
//     func init() {
//         factory.Register((*ActionList)(nil))
//
//         factory.Register(
//              (*SpecialAllocatorType)(nil),
//              Init(func()interface{ return specialAllocatorType.New() }),
//         )
//     }
//
func Register(iface interface{}, opts ...TypeOption) {
	ti := typeInfo{}

	for _, opt := range opts {
		opt(&ti)
	}

	CoreFactoryMap[reflect.TypeOf(iface).Elem().String()] = ti
}

func Init(f func() interface{}) TypeOption {
	return func(ti *TypeInfo) {
		ti.Init = f
	}
}

// type TypeId struct {
// 	Type string
// }

// func GetTypeName(iface interface{}) string {
// 	return reflect.TypeOf(iface).Elem().String()
// }

type Factory struct {
	NameMap map[string]types.JGameObject
	ObjList []JGameObject
}

func (f *Factory) StartUp(config types.JGameObject) {

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
func (f *Factory) MarkObjectForDeletion(types.JGameObject) {

}

type FactoryOpt func(JGameObject)

// s := f.Create(f1, Name("space"))
// goc1 := f.Create(f2, Owner(s), Name("goc1"))
// goc2 := f.Create(f2, Owner(s), Dispatcher(d))
func (f *Factory) Create(file string, opts ...FactoryOpt) types.JGameObject {
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
	return obj
	// h := types.JGameObject{id: len(f.ObjList)}
	// f.NameMap[obj.Name()] = h
	// return h
}

func Name(name string) FactoryOpt {
	return func(obj JGameObject) {
		obj.SetName(name)
	}
}

func New() *Factory {
	return &Factory{}
}

// func AddComp(comp JComponent) FactoryOpt {
//  return func(obj JGameObject) {
//      obj.AddComps(comp)
//  }
// }
