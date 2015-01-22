package event

import "container/list"

type InstantDispatcher struct {
	// coretype.JBaseObject
	ListenerMap map[string]*list.List
	TrackerList []*tracker
}

func (d *InstantDispatcher) Init() {}
func (d *InstantDispatcher) Deinit() {
	d.unhook()
}

type tracker struct {
	L *list.List
	E *list.Element
}

func (d *InstantDispatcher) unhook() {
	for _, t := range d.TrackerList {
		t.L.Remove(t.E)
	}
}

func (d *InstantDispatcher) addTracker(tracker *tracker) {
	d.TrackerList = append(d.TrackerList, tracker)
}

func (d *InstantDispatcher) addListener(event string, handler Handler) *tracker {
	if d.ListenerMap == nil {
		d.ListenerMap = make(map[string]*list.List)
	}

	l := d.ListenerMap[event]
	if l == nil {
		l = list.New()
	}

	e := l.PushFront(handler)
	d.ListenerMap[event] = l

	return &tracker{L: l, E: e}
}

func (d *InstantDispatcher) Dispatch(event string, data Data) {
	l := d.ListenerMap[event]
	for e := l.Front(); e != nil; e = e.Next() {
		f := e.Value.(Handler)
		f(data)
	}
}

func (d *InstantDispatcher) Register(event string, sender Dispatcher, handler Handler) {
	node := sender.addListener(event, handler)
	d.addTracker(node)
}
