package components

import (
	"github.com/jmorgan1321/golang-games/core/types"
)

type EventHandlerComponent struct {
	listeners map[string][]coreType.EventHandlerPair
}

func (ehc *EventHandlerComponent) RegisterForEvent(event string, s coreType.Space, fn coreType.EventHandlerFunc) {
	// TODO: error check
	c, _ := s.GetComponent("EventHandlerComponent")
	sehc := c.(*EventHandlerComponent)
	l := sehc.listeners[event]
	l = append(l, coreType.EventHandlerPair{nil, fn})
	sehc.listeners[event] = l
}
func (ehc *EventHandlerComponent) TriggerEvent(event string, data coreType.EventData) {
	for _, v := range ehc.listeners[event] {
		v.Fn(data)
	}
}

func (ehc *EventHandlerComponent) Construct() error {
	ehc.listeners = make(map[string][]coreType.EventHandlerPair)

	return nil
}
func (ehc *EventHandlerComponent) Destruct() error {
	return nil
}

func (*EventHandlerComponent) IsComponent() {}
