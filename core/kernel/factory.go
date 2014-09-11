package kernel

type ObjectManager struct {
}

func (o *ObjectManager) StartUp(config GameObject) {
}

func (o *ObjectManager) ShutDown() {
}

func (o *ObjectManager) BeginFrame() {
}

func (o *ObjectManager) EndFrame() {
	// endOfFrameCleanUp()
}
