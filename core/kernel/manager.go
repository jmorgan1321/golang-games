package kernel

type Manager interface {
	StartUp(config GameObject)
	ShutDown()
	BeginFrame()
	EndFrame()
}
