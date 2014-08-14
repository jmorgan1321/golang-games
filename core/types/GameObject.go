package coreType

type GameObject struct {
	EventHandler
	Identifier
	Initializer
}

func NewGOC(name string) *GameObject {
	return &GameObject{
		EventHandler: nil,
		Identifier:   nil,
		Initializer:  nil,
	}
}
