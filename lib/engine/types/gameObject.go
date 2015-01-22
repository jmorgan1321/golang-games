package types

import (
	"github.com/jmorgan1321/golang-games/lib/engine/event"
	"github.com/jmorgan1321/golang-games/lib/engine/types/base"
)

type JGameObject interface {
	// JObject
	// JTypedObject
	Name() string
	SetName(string)
	// JInitializer
	Init()
	Deinit()

	// JOwnedObject
	Owner() JGameObject
	SetOwner(JGameObject)

	JComponentManager

	event.Dispatcher
}

// JGameObjectHandle wraps a game object and makes sure that the object is valid
// before making any calls into it.
//
type JGameObjectHandle struct {
	obj JGameObject
	id  int
}

func (o JGameObjectHandle) Get() JGameObject {
	// return factory.GetById(o.id)
	return nil
}
func (o JGameObjectHandle) ID() int {
	return o.id
}

type JBaseObject struct {
	name  string
	owner JGameObject
}

func (b *JBaseObject) Name() string           { return b.name }
func (b *JBaseObject) SetName(s string)       { b.name = s }
func (b *JBaseObject) Owner() JGameObject     { return b.owner }
func (b *JBaseObject) SetOwner(g JGameObject) { b.owner = g }

// type JOwnedObject interface {
// 	Owner() JGameObject
// 	SetOwner(JGameObject)
// }
//
//
type JGameObjectManager interface {
	AddGameObjects(...JGameObject)
	RemGameObject(string)
	GameObject(string) *JGoc
}

type JBasicGameObjectMgr struct {
	base.JObjectMgr
}

func (m *JBasicGameObjectMgr) AddGameObjects(GameObjects ...JGameObject) {
	for _, c := range GameObjects {
		m.Add(c)
	}
}
func (m *JBasicGameObjectMgr) RemGameObject(name string) {
	m.Rem(name)
}

// func (m *JBasicGameObjectMgr) RemIGameObject(iface interface{}) {
// 	m.RemInterface(iface)
// }
func (m *JBasicGameObjectMgr) GameObject(name string) JGameObject {
	return m.Get(name).(JGameObject)
}

// func (m *JBasicGameObjectMgr) IGameObject(iface interface{}) JGameObject {
// 	return m.GetInterface(iface).(JGameObject)
// }
