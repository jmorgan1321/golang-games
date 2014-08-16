package kernel

import (
	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/support"
	"reflect"
)

func LoadConfig(file string) GameObject {
	debug.Trace()
	defer debug.UnTrace()

	data, err := support.OpenFile(file)
	if err != nil {
		support.LogFatal("Failed to open Config file: " + file)
	}

	holder, err := support.ReadData(data)
	if err != nil {
		support.LogFatal("Failed to read in Config file: " + file)
		return nil
	}

	var goc *Goc
	Serialize2(&goc, holder)
	return goc
}

func factoryFunc(typename string) interface{} {
	var obj interface{}
	// TODO: test this assertion
	// We search first so that users can't "overwrite" our core types
	switch typename {
	default:
		if CoreFactoryFunc != nil {
			obj = CoreFactoryFunc(typename)
		}
		if obj == nil {
			support.LogFatal("unknown object type, %s, passed to factory", typename)
		}
	case "Goc":
		obj = &Goc{}
	}
	return obj
}

// Serailze takes a pointer to an object that it will fill out and a
//
// See the Serialize Example for more.
func Serialize2(obj, data interface{}) {
	m := data.(map[string]interface{})

	typename, found := m["Type"]
	if !found {
		support.LogFatal("Unsupported type found in serialization: %s", typename)
	}

	// TODO: Check that anything with a recognized objtypename is a pointer to that type
	objtypename := typename.(string)
	tmp := factoryFunc(objtypename)
	reflect.ValueOf(obj).Elem().Set(reflect.ValueOf(tmp))
	obj = tmp

	// TODO: Error when field, k, in json file isn't member of obj
	for k, v := range m {
		// ignore "Type" since we already processed it
		if k == "Type" {
			continue
		}

		switch vtype := v.(type) {
		default:
			// Assume that anything we don't manually look for is a basic type.
			trySerializeBasicKV(obj, k, v)
		case []interface{}:
			switch k {
			case "Components":
				for _, u := range vtype { // iterate though all components
					data := u.(map[string]interface{})
					var comp Component
					Serialize2(&comp, data)
					obj.(GameObject).AddComp(comp)
				}
			case "Gocs":
				for _, u := range vtype { // iterate though all Gocs
					data := u.(map[string]interface{})
					var goc *Goc
					Serialize2(&goc, data)
					obj.(Space).AddGoc(goc)
				}
			}
		}
	}
}

func trySerializeBasicKV(obj interface{}, memberName string, data interface{}) {
	field := reflect.ValueOf(obj).Elem().FieldByName(memberName)

	switch datatype := data.(type) {
	default:
		support.LogFatal("unknown datatype: %v", datatype)
	case string:
		field.SetString(datatype)
	case int64, float64:
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(int64(datatype.(float64)))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(uint64(datatype.(float64)))
		case reflect.Float32, reflect.Float64:
			field.SetFloat(datatype.(float64))
		}
	case bool:
		field.SetBool(datatype)
	}
}
