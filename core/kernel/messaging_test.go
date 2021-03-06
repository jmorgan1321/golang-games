package kernel

import (
	"github.com/jmorgan1321/golang-games/core/test"
	"testing"
)

type testDispatcher struct {
	EventDispatcher
	MsgData, OtherData int
	MsgReceived        bool
}
type testEventData struct {
	Data  int
	Delay float32
}

func (e *testEventData) GetDelay() float32 {
	return e.Delay
}

func TestMessaging(t *testing.T) {
	dispatcher := testDispatcher{EventDispatcher: &BasicDispatcher{}}

	// test dispatching to one listenre
	rcvr1 := testDispatcher{EventDispatcher: &BasicDispatcher{}}
	rcvr1.RegisterForEvent("event1", dispatcher, func(e EventData) {
		rcvr1.MsgData = e.(*testEventData).Data
	})

	msg := testEventData{Data: 5}
	dispatcher.TriggerEvent("event1", &msg)
	test.ExpectEQ(t, msg.Data, rcvr1.MsgData, "message handler wasn't triggered")

	// dispatching to multiple listeners/events
	rcvr2 := testDispatcher{EventDispatcher: &BasicDispatcher{}}
	rcvr2.RegisterForEvent("event1", dispatcher, func(e EventData) {
		rcvr2.MsgData = e.(*testEventData).Data
	})
	rcvr2.RegisterForEvent("event2", dispatcher, func(e EventData) {
		rcvr2.OtherData = e.(*testEventData).Data
	})

	rcvr1.MsgData = 0
	rcvr2.MsgData = 0
	dispatcher.TriggerEvent("event1", &msg)
	test.ExpectEQ(t, msg.Data, rcvr1.MsgData, "message handler wasn't triggered")
	test.ExpectEQ(t, msg.Data, rcvr2.MsgData, "message handler wasn't triggered")
	test.ExpectNEQ(t, msg.Data, rcvr2.OtherData, "message handler shouldn't have triggered")

	dispatcher.TriggerEvent("event2", &msg)
	test.ExpectEQ(t, msg.Data, rcvr2.OtherData, "message handler wasn't triggered")
}

type testMsgSpace struct {
	EventDispatcher
}

func TestDelayedMessaging(t *testing.T) {
	dispatcher := testDispatcher{EventDispatcher: &DelayDispatcher{BasicDispatcher: &BasicDispatcher{}}}

	space := &testMsgSpace{EventDispatcher: &BasicDispatcher{}}
	dispatcher.SetOwner(space)
	// TODO: make this generic when Component is defined
	dispatcher.EventDispatcher.(*DelayDispatcher).Init()

	rcvr := testDispatcher{EventDispatcher: &BasicDispatcher{}}
	rcvr.RegisterForEvent("event", dispatcher, func(e EventData) {
		rcvr.MsgReceived = true
	})
	msg := testEventData{Data: 5}
	dispatcher.TriggerEvent("event", &msg)
	test.ExpectEQ(t, true, rcvr.MsgReceived, "message handler shouldn't have triggered")

	rcvr.MsgReceived = false
	var delay float32 = 100.0
	msg.Delay = delay
	dispatcher.TriggerEvent("event", &msg)

	// No time passed
	test.ExpectEQ(t, false, rcvr.MsgReceived, "message handler shouldn't have triggered")

	// Too little time passed
	var tooLittleDt float32 = 1
	dtRemaining := delay - tooLittleDt
	space.TriggerEvent("FrameUpdateEvent", &FrameUpdateEvent{tooLittleDt})
	test.ExpectEQ(t, false, rcvr.MsgReceived, "message handler shouldn't have triggered")

	// Enough time passed
	enoughDt := dtRemaining
	space.TriggerEvent("FrameUpdateEvent", &FrameUpdateEvent{enoughDt})
	test.ExpectEQ(t, true, rcvr.MsgReceived, "message handler wasn't triggered")
}

func TestDispatcherUnhooking(t *testing.T) {
	dispatcher := testDispatcher{EventDispatcher: &BasicDispatcher{}}

	// test dispatching to all listeners
	rcvr1 := testDispatcher{EventDispatcher: &BasicDispatcher{}}
	rcvr2 := testDispatcher{EventDispatcher: &BasicDispatcher{}}
	rcvr3 := testDispatcher{EventDispatcher: &DelayDispatcher{BasicDispatcher: &BasicDispatcher{}}}
	rcvr1.RegisterForEvent("event", dispatcher, func(e EventData) {
		rcvr1.MsgReceived = true
	})
	rcvr2.RegisterForEvent("event", dispatcher, func(e EventData) {
		rcvr2.MsgReceived = true
	})
	rcvr3.RegisterForEvent("event", dispatcher, func(e EventData) {
		rcvr3.MsgReceived = true
	})

	dispatcher.TriggerEvent("event", nil)
	test.ExpectEQ(t, true, rcvr1.MsgReceived, "message handler wasn't triggered - rcvr1 #1")
	test.ExpectEQ(t, true, rcvr2.MsgReceived, "message handler wasn't triggered - rcvr2 #1")
	test.ExpectEQ(t, true, rcvr3.MsgReceived, "message handler wasn't triggered - rcvr3 #1")

	// TODO: make this generic when Component is defined
	rcvr2.EventDispatcher.(*BasicDispatcher).DeInit()

	rcvr1.MsgReceived, rcvr2.MsgReceived, rcvr3.MsgReceived = false, false, false
	dispatcher.TriggerEvent("event", nil)
	test.ExpectEQ(t, true, rcvr1.MsgReceived, "message handler wasn't triggered - rcvr1 #2")
	test.ExpectEQ(t, false, rcvr2.MsgReceived, "message handler wasn't unhooked - rcvr2 #2")
	test.ExpectEQ(t, true, rcvr3.MsgReceived, "message handler wasn't triggered - rcvr3 #2")

	rcvr3.EventDispatcher.(*DelayDispatcher).DeInit()

	rcvr1.MsgReceived, rcvr2.MsgReceived, rcvr3.MsgReceived = false, false, false
	dispatcher.TriggerEvent("event", nil)
	test.ExpectEQ(t, true, rcvr1.MsgReceived, "message handler wasn't triggered - rcvr1 #3")
	test.ExpectEQ(t, false, rcvr2.MsgReceived, "message handler wasn't unhooked - rcvr2 #3")
	test.ExpectEQ(t, false, rcvr3.MsgReceived, "message handler wasn't unhooked - rcvr3 #3")

	rcvr1.EventDispatcher.(*BasicDispatcher).DeInit()

	rcvr1.MsgReceived, rcvr2.MsgReceived, rcvr3.MsgReceived = false, false, false
	dispatcher.TriggerEvent("event", nil)
	test.ExpectEQ(t, false, rcvr1.MsgReceived, "message handler wasn't unhooked - rcvr1 #4")
	test.ExpectEQ(t, false, rcvr2.MsgReceived, "message handler wasn't unhooked - rcvr2 #4")
	test.ExpectEQ(t, false, rcvr3.MsgReceived, "message handler wasn't unhooked - rcvr3 #4")

	rcvr1.RegisterForEvent("event", dispatcher, func(e EventData) {
		rcvr1.MsgReceived = true
	})

	rcvr1.MsgReceived, rcvr2.MsgReceived, rcvr3.MsgReceived = false, false, false
	dispatcher.TriggerEvent("event", nil)
	test.ExpectEQ(t, true, rcvr1.MsgReceived, "message handler wasn't triggered - rcvr1 #5")
	test.ExpectEQ(t, false, rcvr2.MsgReceived, "message handler wasn't unhooked - rcvr1 #5")
	test.ExpectEQ(t, false, rcvr3.MsgReceived, "message handler wasn't unhooked - rcvr1 #5")
}
