//beg show 1 OMIT
package main

import (
	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/support"
	"reflect"
)

func factoryFunc(typename string) interface{} {
	var obj interface{}
	// We search first so that users can't "overwrite" our core types
	switch typename {
	default:
		if CoreFactoryFunc != nil { // HLts2
			obj = CoreFactoryFunc(typename) // HLts2
		}
		if obj == nil {
			support.LogFatal("unknown object type, %s, passed to factory", typename)
		}
	case "Goc":
		obj = &Goc{}
	}
	return obj
}

//end show 1 OMIT

//beg show 2 OMIT

// Serailze takes a pointer to an object that it will fill out and and pointer
// to data that will populate the object.
//
// See the Serialize Example for more.
func Serialize(obj, data interface{}) {
	debug.Trace()
	defer debug.UnTrace()

	m := data.(map[string]interface{})

	typename, found := m["Type"]
	if !found {
		support.LogFatal("Unsupported type found in serialization: %s", typename)
	}

	// TODO: Check that anything with a recognized objtypename is a pointer to that type OMIT
	objtypename := typename.(string)
	tmp := factoryFunc(objtypename)
	reflect.ValueOf(obj).Elem().Set(reflect.ValueOf(tmp))
	obj = tmp

	//end show 2 OMIT
	//beg show 3 OMIT

	// TODO: Error when field, k, in json file isn't member of obj  OMIT
	for k, v := range m {
		// ignore "Type" since we already processed it
		if k == "Type" {
			continue
		}

		switch vtype := v.(type) { // HLts
		default:
			// Assume that anything we don't manually look for is a basic type.
			trySerializeBasicKV(obj, k, v)
		case []interface{}: // HLts
			switch k {
			case "Components":
				for _, u := range vtype { // iterate though all components
					data := u.(map[string]interface{})
					var comp Component
					Serialize(&comp, data)
					obj.(GameObject).AddComp(comp) // HLts
				}
			case "Gocs":
				for _, u := range vtype { // iterate though all Gocs
					data := u.(map[string]interface{})
					var goc *Goc
					Serialize(&goc, data)
					obj.(Space).AddGoc(goc) // HLts
				}
				//end show 3 OMIT
			}
		}
	}
}

//beg show 4 OMIT
func trySerializeBasicKV(obj interface{}, memberName string, data interface{}) {
	field := reflect.ValueOf(obj).Elem().FieldByName(memberName)

	switch datatype := data.(type) { // HLsy
	default:
		support.LogFatal("unknown datatype: %v", datatype)
	case string: // HLsy
		field.SetString(datatype)
	case int64, float64: // HLsy
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(int64(datatype.(float64)))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(uint64(datatype.(float64)))
		case reflect.Float32, reflect.Float64:
			field.SetFloat(datatype.(float64))
		}
	case bool: // HLsy
		field.SetBool(datatype)
	}
}

//end show 4 OMIT
