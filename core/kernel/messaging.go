package kernel

import (
	"container/list"
	"github.com/jmorgan1321/golang-games/core/utils"
)

type EventData interface {
	GetDelay() float32
}
type EventHandler func(EventData)

type EventDispatcher interface {
	RegisterForEvent(event string, sender EventDispatcher, handler EventHandler)
	TriggerEvent(event string, data EventData)
	AddListener(event string, handler EventHandler) *tracker
	SetOwner(EventDispatcher)
}

type BasicDispatcher struct {
	ListenerMap map[string]*list.List
	TrackerList []*tracker
	// TODO: get rid of this when component interface is defined
	owner EventDispatcher
}

func (d *BasicDispatcher) DeInit() {
	d.unhook()
}

type tracker struct {
	L *list.List
	E *list.Element
}

func (d *BasicDispatcher) unhook() {
	for _, t := range d.TrackerList {
		t.L.Remove(t.E)
	}
}

func (d *BasicDispatcher) AddTracker(tracker *tracker) {
	d.TrackerList = append(d.TrackerList, tracker)
}

func (d *BasicDispatcher) AddListener(event string, handler EventHandler) *tracker {
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

func (d *BasicDispatcher) TriggerEvent(event string, data EventData) {
	l := d.ListenerMap[event]
	for e := l.Front(); e != nil; e = e.Next() {
		f := e.Value.(EventHandler)
		f(data)
	}
}

func (d *BasicDispatcher) RegisterForEvent(event string, sender EventDispatcher, handler EventHandler) {
	node := sender.AddListener(event, handler)
	d.AddTracker(node)
}

// TODO: get rid of this when component interface is defined.
func (d *BasicDispatcher) SetOwner(owner EventDispatcher) {
	d.owner = owner
}

type DelayedEvent struct {
	Event    string
	Msg      EventData
	FireTime float32
}
type DelayDispatcher struct {
	*BasicDispatcher
	EventQueue []*DelayedEvent
}

func (d *DelayDispatcher) TriggerDelayed(fue *FrameUpdateEvent) {
	// copy current queue, so that any messages that cause a delayed message to
	// fired don't mess us up.
	tmpQueue := d.EventQueue[:len(d.EventQueue)]
	d.EventQueue = nil
	for _, e := range tmpQueue {
		e.FireTime -= fue.Dt
		d.triggerEventDelayed(e.Event, e.Msg, e.FireTime)
	}
}

func (d *DelayDispatcher) Init() {
	// TODO: register with messaging system instead of with Owner.
	d.RegisterForEvent("FrameUpdateEvent", d.owner, func(e EventData) {
		d.TriggerDelayed(e.(*FrameUpdateEvent))
	})
}

func (d *DelayDispatcher) TriggerEvent(event string, data EventData) {
	d.triggerEventDelayed(event, data, data.GetDelay())
}

func (d *DelayDispatcher) triggerEventDelayed(event string, data EventData, delay float32) {
	if delay < utils.Epsilon {
		d.BasicDispatcher.TriggerEvent(event, data)
		return
	}

	delayedMsg := DelayedEvent{Event: event, Msg: data, FireTime: delay}
	d.EventQueue = append(d.EventQueue, &delayedMsg)
}
