package core

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// assert fails the test and displays 'msg', if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// expect fails the test and displays 'msg', if the condition is false.
func expect(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.Fail()
	}
}

// ok fails the test and displays 'err', if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// expectEQ fails the test and displays 'msg', if exp is not equal to act.
func expectEQ(tb testing.TB, exp, act interface{}, msg string) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, msg, exp, act)
		tb.Fail()
	}
}

// expectNEQ fails the test and displays 'msg', if exp is equal to act.
func expectNEQ(tb testing.TB, exp, act interface{}, msg string) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, msg, exp, act)
		tb.Fail()
	}
}

// assertEQ fails the test and displays 'msg', if exp is not equal to act.
func assertEQ(tb testing.TB, exp, act interface{}, msg string) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, msg, exp, act)
		tb.FailNow()
	}
}

// TODO: I should separate things out based on their interfaces.
//       So, RegisterForEvent will take an EventHandler

// Spaces just hold things and know how to delegate to the things they contain.
type Space interface {
	// EventHandler
	EventHandler() *EventHandlerComponent
	RegisterForEvent(string, Space, EventHandlerFunc)
	TriggerEvent(string, EventData)

	// SubSpaceManager - Too specific? SpaceManagerComponent, would allow components to be here as well.
	//                   Components shouldn't have subspaces, though.
	AddSubSpace(Space)
	GetSubSpace(string) (Space, error)
	// RemSubSpace(string)

	// Initializer
	Construct() error
	Destruct() error

	// Identifier
	Name() string
	SetName(string)
	Type() string

	// ComponentManager
	GetComponent(string) (Component, error)
}

type Component interface{}

type BasicSpace struct {
	*IdentifierComponent
	*SpaceManagerComponent
	*EventHandlerComponent
	// TODO: move this into GOC
	*ComponentManagerComponent
}

func (bs *BasicSpace) Construct() error {
	bs.IdentifierComponent = &IdentifierComponent{}

	bs.SpaceManagerComponent = &SpaceManagerComponent{}
	bs.SpaceManagerComponent.Construct()

	bs.EventHandlerComponent = &EventHandlerComponent{}
	bs.EventHandlerComponent.Construct()

	bs.ComponentManagerComponent = &ComponentManagerComponent{}
	bs.ComponentManagerComponent.Construct()
	return nil
}

func (bs *BasicSpace) Destruct() error {
	return nil
}

type ComponentManagerComponent struct {
}

func (cmc *ComponentManagerComponent) Construct() error {
	return nil
}
func (cmc *ComponentManagerComponent) GetComponent(string) (Component, error) {
	return nil, nil
}

// IdentifierComponent exists to allow a GOC to be identify itself.
type IdentifierComponent struct {
	name string
	ty   string
}

func (id *IdentifierComponent) SetName(s string) {
	id.name = s
}
func (id *IdentifierComponent) Name() string {
	return id.name
}
func (id *IdentifierComponent) Type() string {
	return id.ty
}

type SpaceManagerComponent struct {
	spaceMap map[string]Space
}

func (smc *SpaceManagerComponent) Construct() error {
	smc.spaceMap = make(map[string]Space)
	return nil
}

// AddSubSpace inserts the Space into the SpaceManagerComponent, using Space's
// name as the key.
func (smc *SpaceManagerComponent) AddSubSpace(s Space) {
	// TODO: check for duplicates
	smc.spaceMap[s.Name()] = s
}

func (smc *SpaceManagerComponent) GetSubSpace(name string) (Space, error) {
	return smc.spaceMap[name], nil
}

// create will return a new space with the given name, inside of the parent
// (if one is provided).  If jsonData is empty, the new space will be an
// empty space.
//
// create(name string, parentSpace Space, jsonData string) Space
type Core struct {
	ResourceMgr *ResourceManager
}

var core Core = Core{
	ResourceMgr: &ResourceManager{},
}

type ResourceManager struct {
}

func (r *ResourceManager) GetFileData(filename string) (string, error) {
	mockFileSystem := map[string]string{
		"LevelSpaceFile": "{}",
		"GocFile":        "{}",
	}

	return mockFileSystem[filename], nil
}

// ObjectManager creates, destroys, and stores Game Objects.
type ObjectManager struct {
	spaceMap map[string]Space
}

func (o *ObjectManager) Construct() error {
	o.spaceMap = make(map[string]Space)
	return nil
}

func (o *ObjectManager) GetByName(name string) (Space, error) {
	return o.spaceMap[name], nil
}

func (o *ObjectManager) CreateSpace(parent Space, name, filename string) Space {
	spaceJSON, _ := core.ResourceMgr.GetFileData(filename)
	s := o.createSpaceFromString(spaceJSON)
	s.SetName(name)
	o.spaceMap[name] = s
	if parent != nil {
		parent.AddSubSpace(s)
	}
	return s
}

func (o *ObjectManager) createSpaceFromString(data string) Space {
	s := &BasicSpace{}
	// TODO: error check
	s.Construct()
	return s
}

// EventData is anything that can be passed to a message.
// TODO: Should I just use interface{}?
type EventData interface{}

type EventHandlerPair struct {
	Space
	Fn EventHandlerFunc
}

type EventHandlerFunc func(EventData)

type EventHandlerComponent struct {
	listeners map[string][]EventHandlerPair
}

// TODO: remove this for GetComp("EventHandler")
func (ehc *EventHandlerComponent) EventHandler() *EventHandlerComponent {
	return ehc
}

func (ehc *EventHandlerComponent) RegisterForEvent(event string, s Space, fn EventHandlerFunc) {
	l := s.EventHandler().listeners[event]
	l = append(l, EventHandlerPair{nil, fn})
	s.EventHandler().listeners[event] = l
}
func (ehc *EventHandlerComponent) TriggerEvent(event string, data EventData) {
	for _, v := range ehc.listeners[event] {
		v.Fn(data)
	}
}

func (ehc *EventHandlerComponent) Construct() error {
	ehc.listeners = make(map[string][]EventHandlerPair)

	return nil
}

func TestCore(t *testing.T) {
	objMgr := &ObjectManager{}
	objMgr.Construct()

	s := objMgr.CreateSpace(nil, "levelSpace", "_")
	expectEQ(t, "levelSpace", s.Name(), "Name not set correctly.")
	spaceFromMgr, _ := objMgr.GetByName("levelSpace")
	expectEQ(t, s, spaceFromMgr, "levelSpace not stored correctly.")

	g1 := objMgr.CreateSpace(s, "goc1", "GocFile")
	expectEQ(t, "goc1", g1.Name(), "Name not set correctly.")
	gocFromMgr, _ := objMgr.GetByName("goc1")
	expectEQ(t, g1, gocFromMgr, "goc1 not stored correctly.")
	gocFromSpace, _ := s.GetSubSpace(g1.Name())
	expectEQ(t, g1, gocFromSpace, "goc1 not retrieved from levelSpace.")

	g2 := objMgr.CreateSpace(s, "goc2", "GocFile")
	expectEQ(t, "goc2", g2.Name(), "Name not set correctly.")
	goc2FromMgr, _ := objMgr.GetByName("goc2")
	expectEQ(t, g2, goc2FromMgr, "goc2 not stored correctly.")
	goc2FromSpace, _ := s.GetSubSpace(g2.Name())
	expectEQ(t, g2, goc2FromSpace, "goc2 not retrieved from levelSpace.")
	expectNEQ(t, gocFromSpace, goc2FromSpace, "gocs should be different.")

	testInt := 0
	test2Int := 0

	g2.RegisterForEvent("MyEvent", g1, func(e EventData) {
		testInt = e.(int)
	})
	g2.RegisterForEvent("AnotherEvent", g1, func(e EventData) {
		test2Int = e.(int)
	})

	g1.TriggerEvent("MyEvent", 5)
	g1.TriggerEvent("AnotherEvent", 10)

	expectEQ(t, 5, testInt, "Messaging didn't work")
	expectEQ(t, 10, test2Int, "Registering multiple messages didn't work")

	// ehc, _ := g1.GetComponent("EventHandlerComponent")
	// ehc.(*EventHandlerComponent).TriggerEvent("MyEvent", 6)
	// expectEQ(t, 5, testInt, "GetComponent() failed.")
}
