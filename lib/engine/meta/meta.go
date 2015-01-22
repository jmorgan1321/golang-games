package meta

import "reflect"

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
func Register(iface interface{}, opts ...option) {
	ti := typeInfo{}

	for _, opt := range opts {
		opt(&ti)
	}

	TypeMap[reflect.TypeOf(iface).Elem().String()] = ti
}

func Init(f func() interface{}) option {
	return func(ti *typeInfo) {
		ti.Init = f
	}
}

var TypeMap = map[string]typeInfo{}

type typeInfo struct {
	Init func() interface{}
}

type option func(*typeInfo)
