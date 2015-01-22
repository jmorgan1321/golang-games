package types

import "github.com/jmorgan1321/golang-games/lib/engine/event"

type JManager interface {
	StartUp(config JGameObject)
	ShutDown()
	BeginFrame()
	EndFrame()
	ScriptName() string
	event.Dispatcher
}
