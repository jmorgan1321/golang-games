package types

// type JTypedObject interface {
// 	Name() string
// 	SetName(string)
// }
// type JInitializer interface {
// 	Init()
// 	Deinit()
// }
// type JOwnedObject interface {
// 	Owner() JObject
// 	SetOwner(JObject)
// }
// type JGocManager interface {
// 	AddGocs(...JGoc)
// 	RemGoc(string)
// 	Goc(string) *JGoc
// }
// type JSystemManager interface {
// 	AddSystems(...JSystem)
// 	RemSystem(string)
// 	RemISystem(interface{})
// 	System(string) JSystem
// 	ISystem(interface{}) JSystem
// }
// type JComponentManager interface {
// 	AddComps(...JComponent)
// 	RemComp(string)
// 	RemIComp(interface{})
// 	Comp(string) JComponent
// 	IComp(interface{}) JComponent
// }
// type JUpdater interface {
// 	Update(time.Duration)
// }

// type JBaseObject struct {
// 	name  string
// 	owner JObject
// }

// func (o *JBaseObject) Name() string {
// 	return o.name
// }
// func (o *JBaseObject) Owner() JObject {
// 	return o.owner
// }

// type jObjectMgr struct {
// 	objects []interface{}
// }

// func (m *jObjectMgr) add(objs ...interface{}) {
// 	m.objects = append(m.objects, objs...)
// }
// func (m *jObjectMgr) rem(name string) {
// 	for i, obj := range m.objects {
// 		t := reflect.ValueOf(obj).Type().Elem()
// 		if t.Name() == name {
// 			m.objects = append(m.objects[:i-1], m.objects[i+1:]...)
// 			return
// 		}
// 	}
// }
// func (m *jObjectMgr) get(name string) interface{} {
// 	for _, obj := range m.objects {
// 		t := reflect.ValueOf(obj).Type().Elem()
// 		if t.Name() == name {
// 			return obj
// 		}
// 	}
// 	return nil
// }
// func (m *jObjectMgr) remInterface(iface interface{}) {
// 	for i, obj := range m.objects {
// 		t := reflect.ValueOf(obj).Type().Elem()
// 		e := reflect.TypeOf(iface).Elem()
// 		if t.Implements(e) {
// 			m.objects = append(m.objects[:i-1], m.objects[i+1:]...)
// 			return
// 		}
// 	}
// }
// func (m *jObjectMgr) getInterface(iface interface{}) interface{} {
// 	for _, obj := range m.objects {
// 		t := reflect.ValueOf(obj).Type().Elem()
// 		i := reflect.TypeOf(iface).Elem()
// 		if t.Implements(i) {
// 			return obj
// 		}
// 	}
// 	return nil
// }

// type JBasicSystemMgr struct {
// 	jObjectMgr
// }

// func (m *JBasicSystemMgr) AddSystems(systems ...JSystem) {
// 	m.add(systems)
// }
// func (m *JBasicSystemMgr) RemSystem(name string) {
// 	m.rem(name)
// }
// func (m *JBasicSystemMgr) RemISystem(iface interface{}) {
// 	m.remInterface(iface)
// }
// func (m *JBasicSystemMgr) System(name string) JSystem {
// 	return m.get(name).(JSystem)
// }
// func (m *JBasicSystemMgr) ISystem(iface interface{}) JSystem {
// 	return m.getInterface(iface).(JSystem)
// }

// type JBasicCompMgr struct {
// 	jObjectMgr
// }

// func (m *JBasicCompMgr) AddComps(Comps ...JComponent) {
// 	m.add(Comps)
// }
// func (m *JBasicCompMgr) RemComp(name string) {
// 	m.rem(name)
// }
// func (m *JBasicCompMgr) RemIComp(iface interface{}) {
// 	m.remInterface(iface)
// }
// func (m *JBasicCompMgr) Comp(name string) JComponent {
// 	return m.get(name).(JComponent)
// }
// func (m *JBasicCompMgr) IComp(iface interface{}) JComponent {
// 	return m.getInterface(iface).(JComponent)
// }
