package core_event

// FrameUpdateEvent is dispatched by the core, once per frame.
//
// It can be used by game entities that need to know when the frame
// has ticked (and for how long).
type FrameUpdateEvent struct {
	Dt float32
}
