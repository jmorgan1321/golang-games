package components

import (
	"github.com/jmorgan1321/golang-games/core/types"
)

type EventHandlerComponent struct {
	listeners map[string][]types.EventHandlerPair
}

func (ehc *EventHandlerComponent) RegisterForEvent(event string, s types.Space, fn types.EventHandlerFunc) {
	// TODO: error check
	c, _ := s.GetComponent("EventHandlerComponent")
	sehc := c.(*EventHandlerComponent)
	l := sehc.listeners[event]
	l = append(l, types.EventHandlerPair{nil, fn})
	sehc.listeners[event] = l
}
func (ehc *EventHandlerComponent) TriggerEvent(event string, data types.EventData) {
	for _, v := range ehc.listeners[event] {
		v.Fn(data)
	}
}

func (ehc *EventHandlerComponent) Construct() error {
	ehc.listeners = make(map[string][]types.EventHandlerPair)

	return nil
}

func (*EventHandlerComponent) IsComponent() {}
