package serialization

import (
	"fmt"
	"reflect"

	"github.com/jmorgan1321/golang-games/lib/engine/debug"
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
)

func FactoryFunc(typename string) interface{} {
	// defer debug.Trace().UnTrace()

	if ti, found := meta.TypeMap[typename]; found {
		return ti.Init()
	}

	var obj interface{}
	// TODO: test this assertion
	// We search first so that users can't "overwrite" our core types
	switch typename {
	default:
		debug.FatalF("unknown object type, %s, passed to factory", typename)
		// case "Goc":
		// 	obj = &Goc{}
	}
	return obj
}

type Serializer interface {
	Unmarshal(data interface{})
	// Marshal() string
}

func SerializeInPlace(obj, data interface{}) {
	// defer debug.Trace().UnTrace()

	// defer to special handler, if type has one
	if ser, ok := obj.(Serializer); ok {
		ser.Unmarshal(data)
		return
	}

	switch dh := data.(type) {
	default:
		panic(dh)
	case map[string]interface{}:
		for k, v := range dh {
			switch k {
			case "Type":
				continue
			default:
				// basic type (probably)
				field := reflect.ValueOf(obj).Elem().FieldByName(k)
				SerializeInPlace(field, v)
			}
		}
	case []interface{}:
		panic("[]interface{}")

	// basic types:
	case string:
		obj.(reflect.Value).SetString(data.(string))
	case bool:
		obj.(reflect.Value).SetBool(data.(bool))
	case int64, float64:
		switch obj.(reflect.Value).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			obj.(reflect.Value).SetInt(int64(data.(float64)))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			obj.(reflect.Value).SetUint(uint64(data.(float64)))
		case reflect.Float32, reflect.Float64:
			obj.(reflect.Value).SetFloat(data.(float64))
		default:
			panic("unknown numeric type")
		}
	}
}

func SerPlace(obj, data interface{}) {

}

func Serialize3(obj, data interface{}) {
	switch dh := data.(type) {
	case map[string]interface{}:
		fmt.Println(dh)
		panic("test")
		if typ, found := dh["Type"]; found {
			if obj == nil {
				o := FactoryFunc(typ.(string))
				panic(o)
			}
		}
		for k, v := range dh {
			fmt.Println(k, v)

			switch k {
			case "Type":
				continue
			default:
				fmt.Println("\t", k, v)

				// basic type (probably)
				// field := reflect.ValueOf(obj).Elem().FieldByName(k)
				// SerPlace(field, v)
			}
		}
	}
}
