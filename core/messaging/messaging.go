package event

// import (
// 	"container/list"

// 	"github.com/jmorgan1321/golang-games/core/types"
// 	"github.com/jmorgan1321/golang-games/core/utils"
// )

// type Data interface{}
// type DelayedData interface {
// 	Data
// 	GetDelay() float32
// }
// type Handler func(Data)

// type Dispatcher interface {
// 	Register(event string, sender Dispatcher, handler Handler)
// 	Dispatch(event string, data Data)
// 	// addListener(event string, handler Handler) *tracker
// 	types.JComponent
// }

// type InstantDispatcher struct {
// 	JBaseObject
// 	ListenerMap map[string]*list.List
// 	TrackerList []*tracker
// }

// func (d *InstantDispatcher) DeInit() {
// 	d.unhook()
// }

// type tracker struct {
// 	L *list.List
// 	E *list.Element
// }

// func (d *InstantDispatcher) unhook() {
// 	for _, t := range d.TrackerList {
// 		t.L.Remove(t.E)
// 	}
// }

// func (d *InstantDispatcher) addTracker(tracker *tracker) {
// 	d.TrackerList = append(d.TrackerList, tracker)
// }

// func (d *InstantDispatcher) addListener(event string, handler Handler) *tracker {
// 	if d.ListenerMap == nil {
// 		d.ListenerMap = make(map[string]*list.List)
// 	}

// 	l := d.ListenerMap[event]
// 	if l == nil {
// 		l = list.New()
// 	}

// 	e := l.PushFront(handler)
// 	d.ListenerMap[event] = l

// 	return &tracker{L: l, E: e}
// }

// func (d *InstantDispatcher) Dispatch(event string, data Data) {
// 	l := d.ListenerMap[event]
// 	for e := l.Front(); e != nil; e = e.Next() {
// 		f := e.Value.(Handler)
// 		f(data)
// 	}
// }

// func (d *InstantDispatcher) Register(event string, sender Dispatcher, handler Handler) {
// 	node := sender.addListener(event, handler)
// 	d.addTracker(node)
// }

// // TODO: get rid of this when component interface is defined.
// func (d *InstantDispatcher) SetOwner(owner Dispatcher) {
// 	d.owner = owner
// }

// type DelayedEvent struct {
// 	Event    string
// 	Msg      Data
// 	FireTime float32
// }
// type DelayDispatcher struct {
// 	*InstantDispatcher
// 	EventQueue []*DelayedEvent
// }

// func (d *DelayDispatcher) DispatchDelayed(fue *FrameUpdateEvent) {
// 	// copy current queue, so that any messages that cause a delayed message to
// 	// be fired don't mess us up.
// 	tmpQueue := d.EventQueue[:len(d.EventQueue)]
// 	d.EventQueue = nil
// 	for _, e := range tmpQueue {
// 		e.FireTime -= fue.Dt
// 		d.dispatchDelayed(e.Event, e.Msg, e.FireTime)
// 	}
// }

// func (d *DelayDispatcher) Init() {
// 	// TODO: register with messaging system instead of with Owner.
// 	d.Register("FrameUpdateEvent", d.owner, func(e Data) {
// 		d.DispatchDelayed(e.(*FrameUpdateEvent))
// 	})
// }

// func (d *DelayDispatcher) Dispatch(event string, data Data) {
// 	if de, ok := data.(DelayedData); !ok {
// 		d.InstantDispatcher.Dispatch(event, data)
// 	} else {
// 		d.dispatchDelayed(event, data, de.GetDelay())
// 	}
// }

// func (d *DelayDispatcher) dispatchDelayed(event string, data Data, delay float32) {
// 	if delay < utils.Epsilon {
// 		d.InstantDispatcher.Dispatch(event, data)
// 		return
// 	}

// 	delayedMsg := DelayedEvent{Event: event, Msg: data, FireTime: delay}
// 	d.EventQueue = append(d.EventQueue, &delayedMsg)
// }
