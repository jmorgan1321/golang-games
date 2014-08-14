package kernel

import (
	"reflect"
	// "github.com/jmorgan1321/golang-games/core/managers"
	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/support"
)

type GameObject interface{}
type Goc struct {
	Comps []Component
}

func (goc *Goc) AddComp(c Component) {
	goc.Comps = append(goc.Comps, c)
}

type Component interface{}
type Manager interface{}

var CoreTempFactoryFunc func(string) Component

type Core struct {
	// ResourceMgr *managers.ResourceManager
	// Factory     *managers.ObjectManager
	config GameObject
}

func (c *Core) Startup() {
	debug.Trace()
	defer debug.UnTrace()
}
func (c *Core) Shutdown() {
	debug.Trace()
	defer debug.UnTrace()
}
func (c *Core) Run() {
	debug.Trace()
	defer debug.UnTrace()
}

func New(configFile string) *Core {
	debug.Trace()
	defer debug.UnTrace()

	core := &Core{
	// ResourceMgr: &managers.ResourceManager{},
	// Factory:     &managers.ObjectManager{},
	}

	// core.ResourceMgr.Construct()
	// core.Factory.Construct()

	core.config = loadConfig(configFile)
	return core
}

func loadConfig(file string) GameObject {
	debug.Trace()
	defer debug.UnTrace()

	data, err := support.OpenFile(file)
	if err != nil {
		return nil
	}

	holder, err := support.ReadData(data)
	if err != nil {
		return nil
	}

	var goc Goc
	// TODO: move this somewhere else
	Serialize(&goc, holder)
	support.Log("goc: %#v", goc)

	return nil
}

// TODO: Clean up Serialize()

func Serialize(obj, data interface{}) {
	debug.Trace()
	defer debug.UnTrace()

	m := data.(map[string]interface{})
	support.Log("map: %v", m)
	for k, v := range m {
		support.Log("working with field: %s", k)

		if k == "Type" {
			switch v {
			case "Goc":
				// create a goc to put into obj interface?
				continue
			}
		}

		// support.Log("curr obj: %#v", obj)
		currField := reflect.ValueOf(obj).Elem().FieldByName(k)
		// support.Log("currField: %#v", currField)

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
			support.Log("[]interface{}: %v", vv)

			switch k {
			case "Comps":
				for _, u := range vv { // iterate though all components
					support.Log("u: %v", vv)
					compData := u.(map[string]interface{})
					val, ok := compData["Type"]
					if !ok {
						// TODO: handle error
						support.LogError("Component::Type not found.", nil)
					}

					comp := CoreTempFactoryFunc(val.(string))
					support.Log("Comp: %v", comp)
					Serialize(comp, u)
					obj.(*Goc).AddComp(comp)
				}
			}
		}
	}
}
