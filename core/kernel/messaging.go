package kernel

import (
	"github.com/jmorgan1321/golang-games/core/utils"
)

type EventData interface {
	GetDelay() float32
}
type EventHandler func(EventData)

type EventDispatcher interface {
	RegisterForEvent(event string, sender EventDispatcher, handler EventHandler)
	TriggerEvent(event string, data EventData)
	AddListener(event string, handler EventHandler)

	SetOwner(EventDispatcher)
}

type BasicDispatcher struct {
	EventMap map[string][]EventHandler
	owner    EventDispatcher
}

func (d *BasicDispatcher) AddListener(event string, handler EventHandler) {
	if d.EventMap == nil {
		d.EventMap = make(map[string][]EventHandler)
	}

	listeners, _ := d.EventMap[event]
	listeners = append(listeners, handler)
	d.EventMap[event] = listeners
}

func (d *BasicDispatcher) TriggerEvent(event string, data EventData) {
	listeners := d.EventMap[event]
	for _, f := range listeners {
		f(data)
	}
}

func (d *BasicDispatcher) RegisterForEvent(event string, sender EventDispatcher, handler EventHandler) {
	sender.AddListener(event, handler)
	// TODO: d.AddTracker(obj) for unregistering messages
}

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
	tmpQueue := []*DelayedEvent{}
	msgsToFire := []*DelayedEvent{}
	for _, e := range d.EventQueue {
		e.FireTime -= fue.Dt
		msgShouldFire := e.FireTime < utils.Epsilon
		if msgShouldFire {
			msgsToFire = append(msgsToFire, e)
		} else {
			tmpQueue = append(tmpQueue, e)
		}
	}
	d.EventQueue = tmpQueue[:len(tmpQueue)]

	for _, e := range msgsToFire {
		d.BasicDispatcher.TriggerEvent(e.Event, e.Msg)
	}
}

func (d *DelayDispatcher) Init() {
	d.RegisterForEvent("FrameUpdateEvent", d.owner, func(e EventData) {
		d.TriggerDelayed(e.(*FrameUpdateEvent))
	})
}

func (d *DelayDispatcher) TriggerEvent(event string, data EventData) {
	shouldFire := data.GetDelay() < utils.Epsilon
	if shouldFire {
		d.BasicDispatcher.TriggerEvent(event, data)
		return
	}

	delayedMsg := DelayedEvent{Event: event, Msg: data, FireTime: data.GetDelay()}
	d.EventQueue = append(d.EventQueue, &delayedMsg)
}
