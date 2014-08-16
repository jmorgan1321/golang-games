package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
)

/*
s := objMgr.CreateSpace("LevelSpace", "SpaceFile")
g1 := objMgr.CreateGoc(s, "GocFile")
g2 := objMgr.CreateGoc(s, "GocFile")
g2.RegisterForEvent("MyEvent", goc, func(e EventData){
    g2.MyEventHandler(d.(SpecificEventData))
})


t := g1.Get("TranslationComp")
t.X += 10

*/

type EventData interface{}
type EventHandler func(EventData)

type Dispatcher interface {
	TriggerEvent(event string, data EventData)
	RegisterForEvent(obj GameObject, event string, f EventHandler)
	AddListener(event string, handler EventHandler)
}

type BasicDispatcher struct {
	EventMap map[string][]EventHandler
	// Notifiers []GameObject
	Owner GameObject
}

func (d *BasicDispatcher) AddListener(event string, handler EventHandler) {
	if d.EventMap == nil {
		d.EventMap = make(map[string][]EventHandler)
	}

	listeners, _ := d.EventMap[event]
	listeners = append(listeners, handler)
	d.EventMap[event] = listeners
	// return
	// if !found {
	// }
}

func (d *BasicDispatcher) TriggerEvent(event string, data EventData) {
	listeners := d.EventMap[event]
	for _, f := range listeners {
		f(data)
	}
}

func (d *BasicDispatcher) RegisterForEvent(obj GameObject, event string, handler EventHandler) {
	obj.AddListener(event, handler)
	// TODO: d.AddTracker(obj) for unregistering messages
}

type GameObject interface {
	// Initialization
	// TODO: Construct should take a parent, maybe be called Init(parent)
	Construct()
	Destruct()

	// TODO: This is too specific.  This should be in GameObjectComposition which Goc is a concrete class of
	// Component management
	GetComp(string) (Component, error)

	Dispatcher

	// GetParent()
}

type Manager interface {
	StartUp(*Core, GameObject)
	ShutDown()

	BeginFrame()
	EndFrame()
}

type Component interface {
	// IsComponent()
}

type System interface {
	// BeginFrame()
	// EndFrame()
	StartUp()
	ShutDown()

	Update()

	Register(Component)
}

type Space interface {
	Construct()
	Destruct()

	Update()

	AddSystem(sys System)
	AddGoc(goc *Goc)
}

type LevelSpace struct {
	systems []System
	Dispatcher
}

func (ls *LevelSpace) Construct() {
	DebugTrace()
	defer DebugUnTrace()

	grfx := &GraphicsSystem{Dispatcher: &BasicDispatcher{}}
	ls.systems = append(ls.systems, grfx)

	// TODO: move this into GraphicsSystem.Init()
	//       must store parent space in GameObject
	grfx.RegisterForEvent(ls, "RegisterGraphicsComp", func(e EventData) {
		grfx.Register(e.(*GraphicsComp))
	})
	// TODO: move this into GraphicsComp.Init()
	ls.TriggerEvent("RegisterGraphicsComp", &GraphicsComp{X: 10, Y: 2, Dispatcher: &BasicDispatcher{}})
}

func (ls *LevelSpace) GetComp(string) (Component, error) {
	return nil, nil
}

func (ls *LevelSpace) Destruct() {
	DebugTrace()
	defer DebugUnTrace()
}
func (ls *LevelSpace) Update() {
	DebugTrace()
	defer DebugUnTrace()

	for _, sys := range ls.systems {
		sys.Update()
	}
}
func (ls *LevelSpace) AddSystem(sys System) {
	DebugTrace()
	defer DebugUnTrace()
}
func (ls *LevelSpace) AddGoc(goc *Goc) {
	DebugTrace()
	defer DebugUnTrace()
}

// func (ls *LevelSpace) GetComp(string) (Component, error)

type Goc struct {
	Type  string
	Comps []Component
	Dispatcher
}

func (goc *Goc) Construct() {}
func (goc *Goc) Destruct()  {}
func (goc *Goc) GetComp(name string) (Component, error) {
	// v := reflect.ValueOf(cmc.Owner)
	//
	// typ := v.Type()
	// if typ.Kind() == reflect.Ptr {
	//  typ = typ.Elem()
	// }
	//
	// // check for static components
	// c := v.Elem().FieldByName(name).Elem().Addr()
	//
	// iface := c.Interface().(Component)
	// return iface, nil

	// v := reflect.ValueOf(goc)
	// typ := v.Type().Elem()

	// TODO: check for static components

	for _, comp := range goc.Comps {
		ctype := reflect.ValueOf(comp).Type().Elem()
		if ctype.Name() == name {
			return comp, nil
		}
	}
	return nil, nil
}

func (goc *Goc) AddComp(c Component) {
	goc.Comps = append(goc.Comps, c)
}

type Core struct {
	managers                  []Manager
	spaces                    map[string]Space
	shouldRestart, shouldQuit bool
	Factory                   *Factory
	ResourceMgr               *ResourceMgr
}

func (c *Core) StartUp(config GameObject) {
	DebugTrace()
	defer DebugUnTrace()

	c.managers = LoadManagers(c, config)

	c.spaces = make(map[string]Space)
	c.spaces["level"] = LoadLevel()

	for _, sp := range c.spaces {
		sp.Construct()
	}
}

func (c *Core) ShutDown() {
	DebugTrace()
	defer DebugUnTrace()

	for _, spc := range c.spaces {
		spc.Destruct()
	}

	for i := len(c.managers); i > 0; i-- {
		c.managers[i-1].ShutDown()
	}
	c.managers = nil

	c.Factory = nil
	c.ResourceMgr = nil
}

func (c *Core) Run() {
	DebugTrace()
	defer DebugUnTrace()

GameLoop:
	for {
		c.stepFrame()
		c.shouldQuit = true
		if c.shouldRestart || c.shouldQuit {
			break GameLoop
		}
	}
}

func (c *Core) stepFrame() {
	DebugTrace()
	defer DebugUnTrace()

	for _, mgr := range c.managers {
		mgr.BeginFrame()
		defer mgr.EndFrame()
	}

	for _, spc := range c.spaces {
		spc.Update()
	}
}

func LoadLevel() Space {
	return &LevelSpace{Dispatcher: &BasicDispatcher{}}
}

func LoadConfig(file string) *Goc {
	DebugTrace()
	defer DebugUnTrace()

	// goc := &Goc{
	//  Type: "Goc",
	//  Comps: []Component{
	//      &FactoryInitComp{
	//          Type:     "FactoryInitComp",
	//          DummyVar: true,
	//      },
	//      &ResourceMgrInitComp{
	//          Type:          "ResourceMgrInitComp",
	//          FavoriteColor: "bluegreen",
	//      },
	//  },
	// }

	// b, _ := json.MarshalIndent(goc, "", "\t")
	// fmt.Println(indent, "\t\t", string(b))
	filename := "D:/work/fbl_grfx_dev_p/windows/sandbox/go_tools/rsrc/json/objects/GameCoreConfig.jrm"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(indent, "\t\terror:", err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(indent, "\t\terror:", err)
	}

	// type DataHolder struct {
	//  Type  string
	//  Comps map[string]interface{}
	// }

	// var dataHolder DataHolder
	var dataHolder interface{}
	err = json.Unmarshal(data, &dataHolder)
	if err != nil {
		fmt.Println(indent, "\t\terror:", err)
	}
	fmt.Println(indent, "\t\tdata:", dataHolder)

	var goc Goc
	Serialize(&goc, dataHolder)
	fmt.Println(indent, "\t\tgoc:", goc)

	// TODO: move this into serializer?
	// set up static components
	goc.Dispatcher = &BasicDispatcher{}

	return &goc
}

func Serialize(obj, data interface{}) {
	DebugTrace()
	defer DebugUnTrace()

	m := data.(map[string]interface{})
	fmt.Println(indent, "\t\tmap:", m)
	for k, v := range m {
		fmt.Println(indent, "\t\tworking with field:", k)

		if k == "Type" {
			switch v {
			case "Goc":
				// create a goc to put into obj interface?
				continue
			}
		}

		currField := reflect.ValueOf(obj).Elem().FieldByName(k)
		// fmt.Println(indent, "\t\tcurrData:", currData)

		switch vv := v.(type) {
		case string:
			currField.SetString(vv)
		case int64, float64:
			switch currField.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				currField.SetInt(int64(vv.(float64)))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				currField.SetUint(uint64(vv.(float64)))
			case reflect.Float32, reflect.Float64:
				currField.SetFloat(vv.(float64))
			}
		case bool:
			currField.SetBool(vv)
		case []interface{}: // non basic type fields
			fmt.Println(indent, "\t\t[]interface{}", vv)

			switch k {
			case "Comps":
				fmt.Println(indent, "\t\tComps")
				for _, u := range vv { // iterate though all components
					fmt.Println(indent, "\t\tu:", u)
					compData := u.(map[string]interface{})
					val, ok := compData["Type"]
					if !ok {
						// TODO: handle error
					}

					comp := fakeFactoryFunc(val.(string))
					fmt.Println(indent, "\t\tComp", comp)
					Serialize(comp, u)
					obj.(*Goc).AddComp(comp)
				}
			}

		}
	}
}

func fakeFactoryFunc(name string) Component {
	DebugTrace()
	defer DebugUnTrace()

	var c Component
	switch name {
	default:
		panic("Unknown component passed into factory func: " + name)
	case "FactoryInitComp":
		c = &FactoryInitComp{}
	case "ResourceMgrInitComp":
		c = &ResourceMgrInitComp{}
	}
	return c
}

func LoadManagers(c *Core, config GameObject) []Manager {
	DebugTrace()
	defer DebugUnTrace()

	mgrs := []Manager{
		&ResourceMgr{},
		&Factory{},
	}
	for _, m := range mgrs {
		m.StartUp(c, config)
	}

	return mgrs
}

type FactoryInitComp struct {
	DummyVar bool
	Type     string
}

type Factory struct {
	toBeDeleted []GameObject
}

func (f *Factory) BeginFrame() {
	DebugTrace()
	defer DebugUnTrace()
}
func (f *Factory) EndFrame() {
	DebugTrace()
	defer DebugUnTrace()
}

func (f *Factory) StartUp(core *Core, config GameObject) {
	DebugTrace()
	defer DebugUnTrace()

	// TODO: error handle
	core.Factory = f
	data, _ := config.GetComp("FactoryInitComp")
	cfg := data.(*FactoryInitComp)
	if cfg.DummyVar {
		fmt.Println(indent, "\tFactory: got the config data")
	}
}
func (f *Factory) ShutDown() {
	DebugTrace()
	defer DebugUnTrace()
}

type ResourceMgrInitComp struct {
	FavoriteColor string
	Type          string
}

type ResourceMgr struct {
}

func (r *ResourceMgr) BeginFrame() {
	DebugTrace()
	defer DebugUnTrace()
}
func (r *ResourceMgr) EndFrame() {
	DebugTrace()
	defer DebugUnTrace()
}

func (r *ResourceMgr) StartUp(core *Core, config GameObject) {
	DebugTrace()
	defer DebugUnTrace()

	// TODO: error handle
	core.ResourceMgr = r
	data, _ := config.GetComp("ResourceMgrInitComp")
	cfg := data.(*ResourceMgrInitComp)
	if cfg.FavoriteColor == "bluegreen" {
		fmt.Println(indent, "\tResourecMgr: got the config data")
	}
}
func (r *ResourceMgr) ShutDown() {
	DebugTrace()
	defer DebugUnTrace()
}

type GraphicsSystem struct {
	comps []*GraphicsComp
	Dispatcher
}

func (grfx *GraphicsSystem) StartUp() {
	DebugTrace()
	defer DebugUnTrace()
}
func (grfx *GraphicsSystem) ShutDown() {
	DebugTrace()
	defer DebugUnTrace()
}
func (grfx *GraphicsSystem) Update() {
	DebugTrace()
	defer DebugUnTrace()

	for _, c := range grfx.comps {
		fmt.Println(indent, "\tGraphicsSystem: updating comp:", c)
	}
}
func (grfx *GraphicsSystem) Register(c Component) {
	DebugTrace()
	defer DebugUnTrace()

	grfx.comps = append(grfx.comps, c.(*GraphicsComp))
}

type GraphicsComp struct {
	X, Y int
	Dispatcher
}

// TODO: show this to digipen
type IndentLevel int

func (l *IndentLevel) Increment() {
	*l++
}
func (l *IndentLevel) Decrement() {
	*l--
}
func (l *IndentLevel) String() string {
	s := ""
	for i := *l; i > 0; i-- {
		s += "\t"
	}
	return s
}

var indent *IndentLevel

func DebugTrace() {
	indent.Increment()

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("what??")
		return
		// return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		fmt.Println("what??")
		return
		// return "unnamed"
	}
	fmt.Printf("%s%s()\n", indent, fn.Name())
}
func DebugUnTrace() {
	indent.Decrement()
}

func main() {
	indent = new(IndentLevel)

	fmt.Println("starting game")
	defer fmt.Println("ending game")

	DebugTrace()
	defer DebugUnTrace()

	core := &Core{}

	cfg := LoadConfig("config.txt")
	core.StartUp(cfg)
	// fmt.Printf("\n%#v\n", core)

	core.Run()
	// fmt.Printf("\n%#v\n", core)

	core.ShutDown()
	// fmt.Printf("\n%#v\n", core)
}
