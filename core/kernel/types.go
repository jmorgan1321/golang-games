package kernel

import (
	"reflect"
	"time"
)

// Spaces are containers of Gocs and Systems and can be treated like levels in
// a game context.
//
type Space interface {
	Init()
	DeInit()
	Update(dt float32)
	AddGoc(goc *Goc)
	EventDispatcher
	// GocManager
	// SystemManager
	// GameObject
}

// Managers are long-running game operations/processes that handle operations
// outside of a space.
//
// An example would be audio; mulitple spaces can play sounds, so they would
// delegate that operation to the core's one audio manager.
//
type Manager interface {
	StartUp(config GameObject)
	ShutDown()
	BeginFrame()
	EndFrame()

	EventDispatcher
}

// Components hold data for GameObjects, which differentiates game
// objects from each other.
type Component interface {
	SetOwner(GameObject)
	// Owner() GameObject
	Initializer
}

type Initializer interface {
	Init()
	DeInit()
}

// GameObjects are containers for Components.
type GameObject interface {
	ComponentManager
	// EventDispatcher
	// TypedObject
}

// Goc (aka Game Object Composition) is a concrete implementation of GameObject.
type Goc struct {
	Name       string
	Components []Component
}

func (goc *Goc) AddComps(cmps ...Component) {
	for _, c := range cmps {
		goc.Components = append(goc.Components, c)
		c.SetOwner(goc)
	}
}

// GetComp searches this goc for a component with the matching and returns it
// as a generic component, which can then be casted to it's actual type.
//
// ie:
//  alist := goc.GetComp("ActionListComponent").(*ActionListComponent)
func (goc *Goc) GetComp(name string) Component {
	for _, comp := range goc.Components {
		// TODO: possibly switch this to use package.Name
		ctype := reflect.ValueOf(comp).Type().Elem()
		if ctype.Name() == name {
			return comp
		}
	}
	// TODO: better error message when component isn't found
	return nil
}

// GetCompInterface finds the first component of goc that implements the given
// interface.  That component is returned as a generic component, which can
// then be casted to it's actual type.
//
// ie:
//  input := goc.GetCompInterface(new(InputHandlerComponent)).(*PlayerInputComponent)
func (goc *Goc) GetCompInterface(iface interface{}) Component {
	for _, comp := range goc.Components {
		ctype := reflect.TypeOf(comp)
		itype := reflect.TypeOf(iface).Elem()
		if ctype.Implements(itype) {
			return comp
		}
	}
	return nil
}

func (goc *Goc) Unmarshal(data interface{}) {
	m := data.(map[string]interface{})

	// Unmarshall all "normal" fields
	for k, v := range m {
		if k == "Type" || k == "Components" {
			continue
		}

		f := reflect.ValueOf(goc).Elem().FieldByName(k)
		SerializeInPlace(f, v)
	}

	// Special case handling for `Components`
	for _, v := range m["Components"].([]interface{}) {
		compData := v.(map[string]interface{})

		typename, _ := compData["Type"]
		comp := factoryFunc(typename.(string)).(Component)
		SerializeInPlace(comp, compData)
		goc.AddComps(comp)
	}
}

type ComponentManager interface {
	GetComp(name string) Component
	AddComps(...Component)
}
type OwnerMngr struct {
	owner GameObject
}

func (o *OwnerMngr) SetOwner(g GameObject) {
	o.owner = g
}
func (o *OwnerMngr) Owner() GameObject {
	return o.owner
}

/************************************





***********************************/

type JObject interface {
	JEventDispatcher
	JOwnedObject
	JTypedObject
	JInitializer
}

type JGameObject interface {
	JObject
	JComponentManager
}

type JSpace interface {
	JGameObject
	JSystemManager
	JGocManager
	JUpdater
}

type JSystem interface {
	JObject
	JUpdater
}

type JManager interface {
	StartUp(config JGameObject)
	ShutDown()
	BeginFrame()
	EndFrame()
}

type JComponent interface {
	JObject
}

type JEventDispatcher interface {
	// JComponent
	RegisterForEvent(event string, sender JEventDispatcher, handler EventHandler)
	TriggerEvent(event string, data EventData)
	AddListener(event string, handler EventHandler) *tracker
}

type JTypedObject interface {
	Name() string
	SetName(string)
}
type JInitializer interface {
	Init()
	Deinit()
}
type JOwnedObject interface {
	Owner() JObject
	SetOwner(JObject)
}
type JGocManager interface {
	AddGocs(...JGoc)
	RemGoc(string)
	Goc(string) *JGoc
}
type JSystemManager interface {
	AddSystems(...JSystem)
	RemSystem(string)
	RemISystem(interface{})
	System(string) JSystem
	ISystem(interface{}) JSystem
}
type JComponentManager interface {
	AddComps(...JComponent)
	RemComp(string)
	RemIComp(interface{})
	Comp(string) JComponent
	IComp(interface{}) JComponent
}
type JUpdater interface {
	Update(time.Duration)
}

type JBaseObject struct {
	name  string
	owner JObject
}

func (o *JBaseObject) Name() string {
	return o.name
}
func (o *JBaseObject) Owner() JObject {
	return o.owner
}

type jObjectMgr struct {
	objects []interface{}
}

func (m *jObjectMgr) add(objs ...interface{}) {
	m.objects = append(m.objects, objs...)
}
func (m *jObjectMgr) rem(name string) {
	for i, obj := range m.objects {
		t := reflect.ValueOf(obj).Type().Elem()
		if t.Name() == name {
			m.objects = append(m.objects[:i-1], m.objects[i+1:]...)
			return
		}
	}
}
func (m *jObjectMgr) get(name string) interface{} {
	for _, obj := range m.objects {
		t := reflect.ValueOf(obj).Type().Elem()
		if t.Name() == name {
			return obj
		}
	}
	return nil
}
func (m *jObjectMgr) remInterface(iface interface{}) {
	for i, obj := range m.objects {
		t := reflect.ValueOf(obj).Type().Elem()
		e := reflect.TypeOf(iface).Elem()
		if t.Implements(e) {
			m.objects = append(m.objects[:i-1], m.objects[i+1:]...)
			return
		}
	}
}
func (m *jObjectMgr) getInterface(iface interface{}) interface{} {
	for _, obj := range m.objects {
		t := reflect.ValueOf(obj).Type().Elem()
		i := reflect.TypeOf(iface).Elem()
		if t.Implements(i) {
			return obj
		}
	}
	return nil
}

type JBasicSystemMgr struct {
	jObjectMgr
}

func (m *JBasicSystemMgr) AddSystems(systems ...JSystem) {
	m.add(systems)
}
func (m *JBasicSystemMgr) RemSystem(name string) {
	m.rem(name)
}
func (m *JBasicSystemMgr) RemISystem(iface interface{}) {
	m.remInterface(iface)
}
func (m *JBasicSystemMgr) System(name string) JSystem {
	return m.get(name).(JSystem)
}
func (m *JBasicSystemMgr) ISystem(iface interface{}) JSystem {
	return m.getInterface(iface).(JSystem)
}

type JBasicCompMgr struct {
	jObjectMgr
}

func (m *JBasicCompMgr) AddComps(Comps ...JComponent) {
	m.add(Comps)
}
func (m *JBasicCompMgr) RemComp(name string) {
	m.rem(name)
}
func (m *JBasicCompMgr) RemIComp(iface interface{}) {
	m.remInterface(iface)
}
func (m *JBasicCompMgr) Comp(name string) JComponent {
	return m.get(name).(JComponent)
}
func (m *JBasicCompMgr) IComp(iface interface{}) JComponent {
	return m.getInterface(iface).(JComponent)
}

// type JGoc struct {
// 	// JBasicCompMgr
// 	// JBaseObject
// }

// type JCore struct {
// 	root Space
// }

// type JFactory struct {
// }

// func (f *JFactory) Create(file string, opts ...Options) GameObjectHandle {
// 	return nil
// }

// s := f.Create(f1, Name("space"))
// goc1 := f.Create(f2, Owner(s), Name("goc1"))
// goc2 := f.Create(f2, Owner(s), Dispatcher(d))
