package kernel

// import (
// 	"reflect"

// 	"github.com/jmorgan1321/golang-games/core/support"
// )

// func factoryFunc(typename string) interface{} {
// 	if f, found := coreFactoryMap[typename]; found {
// 		return f()
// 	}

// 	var obj interface{}
// 	// TODO: test this assertion
// 	// We search first so that users can't "overwrite" our core types
// 	switch typename {
// 	default:
// 		if CoreFactoryFunc != nil {
// 			obj = CoreFactoryFunc(typename)
// 		}
// 		if obj == nil {
// 			support.LogFatal("unknown object type, %s, passed to factory", typename)
// 		}
// 	case "Goc":
// 		obj = &Goc{}
// 	}
// 	return obj
// }

// type Serializer interface {
// 	Unmarshal(data interface{})
// 	// Marshal() string
// }

// func SerializeInPlace(obj, data interface{}) {
// 	// defer to special handler, if type has one
// 	if ser, ok := obj.(Serializer); ok {
// 		ser.Unmarshal(data)
// 		return
// 	}

// 	switch dh := data.(type) {
// 	default:
// 		panic(dh)
// 	case map[string]interface{}:
// 		for k, v := range dh {
// 			switch k {
// 			case "Type":
// 				continue
// 			default:
// 				// basic type (probably)
// 				field := reflect.ValueOf(obj).Elem().FieldByName(k)
// 				SerializeInPlace(field, v)
// 			}

// 		}
// 	case []interface{}:
// 		panic("[]interface{}")

// 	// basic types:
// 	case string:
// 		obj.(reflect.Value).SetString(data.(string))
// 	case bool:
// 		obj.(reflect.Value).SetBool(data.(bool))
// 	case int64, float64:
// 		switch obj.(reflect.Value).Kind() {
// 		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 			obj.(reflect.Value).SetInt(int64(data.(float64)))
// 		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 			obj.(reflect.Value).SetUint(uint64(data.(float64)))
// 		case reflect.Float32, reflect.Float64:
// 			obj.(reflect.Value).SetFloat(data.(float64))
// 		default:
// 			panic("unknown numeric type")
// 		}
// 	}
// }
