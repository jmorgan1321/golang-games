package base

import (
	"reflect"
	"time"
)

type JObject interface {
	JTypedObject
	JInitializer
	JOwnedObject
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

type JUpdater interface {
	Update(time.Duration)
}

type JBaseObject struct {
	name  string
	owner JObject
}

func (o *JBaseObject) SetName(s string) {
	o.name = s
}
func (o *JBaseObject) Name() string {
	return o.name
}
func (o *JBaseObject) Owner() JObject {
	return o.owner
}
func (o *JBaseObject) SetOwner(obj JObject) {
	o.owner = obj
}

type JObjectMgr struct {
	Objects []interface{}
}

func (m *JObjectMgr) Add(objs ...interface{}) {
	m.Objects = append(m.Objects, objs...)
}
func (m *JObjectMgr) Rem(name string) {
	for i, obj := range m.Objects {
		t := reflect.ValueOf(obj).Type().Elem()
		if t.String() == name {
			m.Objects = append(m.Objects[:i-1], m.Objects[i+1:]...)
			return
		}
	}
	panic("could not find " + name)
}
func (m *JObjectMgr) Get(name string) interface{} {
	for _, obj := range m.Objects {
		t := reflect.ValueOf(obj).Type().Elem()
		if t.String() == name {
			return obj
		}
	}
	panic("could not find " + name)
	return nil
}
func (m *JObjectMgr) RemInterface(iface interface{}) {
	e := reflect.TypeOf(iface).Elem()
	for i, obj := range m.Objects {
		t := reflect.ValueOf(obj).Type().Elem()
		if t.Implements(e) {
			m.Objects = append(m.Objects[:i-1], m.Objects[i+1:]...)
			return
		}
	}
	panic("could not find " + e.Name())
}
func (m *JObjectMgr) GetInterface(iface interface{}) interface{} {
	i := reflect.TypeOf(iface).Elem()
	for _, obj := range m.Objects {
		t := reflect.ValueOf(obj).Type().Elem()
		if t.Implements(i) {
			return obj
		}
	}
	panic("could not find " + i.Name())
	return nil
}
func (m *JObjectMgr) GetAll() []interface{} {
	return m.Objects
}
