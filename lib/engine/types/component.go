package types

import (
	"github.com/jmorgan1321/golang-games/lib/engine/debug"
	"github.com/jmorgan1321/golang-games/lib/engine/types/base"
)

type JComponent interface {
	// JOwnedObject
	Owner() JGameObject
	SetOwner(JGameObject)

	// JInitializer
	Init()
	Deinit()
}

type JComponentManager interface {
	AddComps(...JComponent)
	RemComp(string)
	RemIComp(interface{})
	Comp(string) JComponent
	IComp(interface{}) JComponent
}

type JBasicCompMgr struct {
	base.JObjectMgr
}

func (m *JBasicCompMgr) AddComps(comps ...JComponent) {
	defer debug.Trace().UnTrace()
	for _, c := range comps {
		m.Add(c)
	}
}
func (m *JBasicCompMgr) RemComp(name string) {
	defer debug.Trace().UnTrace()
	m.Rem(name)
}
func (m *JBasicCompMgr) RemIComp(iface interface{}) {
	defer debug.Trace().UnTrace()
	m.RemInterface(iface)
}
func (m *JBasicCompMgr) Comp(name string) JComponent {
	defer debug.Trace().UnTrace()
	return m.Get(name).(JComponent)
}
func (m *JBasicCompMgr) IComp(iface interface{}) JComponent {
	defer debug.Trace().UnTrace()
	return m.GetInterface(iface).(JComponent)
}
