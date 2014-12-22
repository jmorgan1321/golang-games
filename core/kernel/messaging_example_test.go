package kernel

import (
	"fmt"
)

func Example_messaging() {
	dispatcher := testDispatcher{EventDispatcher: &BasicDispatcher{}}

	rcvr1 := testDispatcher{EventDispatcher: &BasicDispatcher{}}
	rcvr2 := testDispatcher{EventDispatcher: &BasicDispatcher{}}
	rcvr3 := testDispatcher{EventDispatcher: &DelayDispatcher{BasicDispatcher: &BasicDispatcher{}}}
	rcvr1.RegisterForEvent("event", dispatcher, func(e EventData) {
		fmt.Println("rcvr1 received message:", e.(*testEventData).Data)
	})
	rcvr2.RegisterForEvent("event", dispatcher, func(e EventData) {
		fmt.Println("rcvr2 received message:", e.(*testEventData).Data)
	})
	rcvr3.RegisterForEvent("event", dispatcher, func(e EventData) {
		fmt.Println("rcvr3 received message:", e.(*testEventData).Data)
	})
	dispatcher.TriggerEvent("event", &testEventData{Data: 5})

	fmt.Println("")
	// TODO: make this generic when Component is defined
	rcvr2.EventDispatcher.(*BasicDispatcher).DeInit()
	dispatcher.TriggerEvent("event", &testEventData{Data: 42})

	fmt.Println("")
	rcvr3.EventDispatcher.(*DelayDispatcher).DeInit()
	dispatcher.TriggerEvent("event", &testEventData{Data: 13})

	fmt.Println("")
	rcvr1.EventDispatcher.(*BasicDispatcher).DeInit()
	dispatcher.TriggerEvent("event", &testEventData{Data: 5})
	fmt.Println("...")

	fmt.Println("")

	rcvr1.RegisterForEvent("event", dispatcher, func(e EventData) {
		fmt.Println("rcvr1 received message:", e.(*testEventData).Data)
	})
	dispatcher.TriggerEvent("event", &testEventData{Data: 21})

	// Output:
	// rcvr3 received message: 5
	// rcvr2 received message: 5
	// rcvr1 received message: 5
	//
	// rcvr3 received message: 42
	// rcvr1 received message: 42
	//
	// rcvr1 received message: 13
	//
	// ...
	//
	// rcvr1 received message: 21
	//
}
