package types

// import (
//   "github.com/jmorgan1321/golang-games/core/kernel"
// )

// EventData is anything that can be passed to a message.
// TODO: Should I just use interface{}?
type EventData interface{}

type EventHandlerPair struct {
	Space // Unused
	Fn    EventHandlerFunc
}

type EventHandlerFunc func(EventData)

type EventHandler interface {
	RegisterForEvent(string, Space, EventHandlerFunc)
	TriggerEvent(string, EventData)
}

// type SubSpaceManager interface {
// 	AddSubSpace(Space)
// 	GetSubSpace(string) (Space, error)
// 	// RemSubSpace(string)
// }

// Identifier types store basic (uniquely) identifying information about themselves.
type Identifier interface {
	Name() string
	SetName(string)
	// Type() string
}

// Initalizer types can be constructed/destructed.
type Initializer interface {
	Construct() error
	Destruct() error
}

type ComponentManager interface {
	GetComponent(string) (interface{}, error)
}
