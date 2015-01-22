package types

import (
	"time"

	"github.com/jmorgan1321/golang-games/lib/engine/types/base"
)

type JSystem interface {
	Name() string
	SetName(string)
	// JInitializer
	Init()
	Deinit()

	// JOwnedObject
	Owner() JGameObject
	SetOwner(JGameObject)

	// JUpdater
	Update(time.Duration)
}

type JSystemManager interface {
	AddSystems(...JSystem)
	RemSystem(string)
	RemISystem(interface{})
	System(string) JSystem
	ISystem(interface{}) JSystem
}

type JBasicSystemMgr struct {
	base.JObjectMgr
}

func (m *JBasicSystemMgr) AddSystems(systems ...JSystem) {
	for _, s := range systems {
		m.Add(s)
	}
}
func (m *JBasicSystemMgr) RemSystem(name string) {
	m.Rem(name)
}
func (m *JBasicSystemMgr) RemISystem(iface interface{}) {
	m.RemInterface(iface)
}
func (m *JBasicSystemMgr) System(name string) JSystem {
	return m.Get(name).(JSystem)
}
func (m *JBasicSystemMgr) ISystem(iface interface{}) JSystem {
	return m.GetInterface(iface).(JSystem)
}
