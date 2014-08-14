package coreType

type Component interface {
	IsComponent()

	// interfaces Components satisfy
	Initializer
	Identifier
	EventHandler
}
