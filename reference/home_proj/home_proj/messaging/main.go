package main

import (
	"fmt"
)

type MessageHandler func(Object, Message)
type Message interface {
	// MessageId int
}

type UpdateData struct {
	dt float32
}

type CollisionInfo struct {
	id1, id2 int
}

type MessageReceiver struct {
	O Object
	F MessageHandler
}

type Object interface {
	GetDispatcher() Dispatcher
}

type Obj struct {
	name string
	Dispatcher
}

func (o *Obj) HandleCollision(c *CollisionInfo) {
	fmt.Printf("%s received collision info: %#v\n", o.name, c)
}

type UISystem struct {
	name string
	Dispatcher
}

func (ui *UISystem) Update(c *CollisionInfo) {
	fmt.Printf("%s received collision info: %#v\n", ui.name, c)
}

type Dispatcher interface {
	Dispatch(event string, m Message)
	// DispatchDelayed(event string, m Message, delay float32)
	Register(obj Object, event string, f MessageHandler)
	GetDispatcher() Dispatcher
}

type BasicDispatcher struct {
	EventMap map[Object][]MessageReceiver
	Owner    Object
}

func (d *BasicDispatcher) GetDispatcher() Dispatcher {
	return d
}

func (d *BasicDispatcher) Dispatch(event string, m Message) {
	recievers := d.EventMap[d.Owner]
	for _, r := range recievers {
		r.F(r.O, m)
	}
}

// func (d *BasicDispatcher) DispatchDelayed(event string, m Message, delay float32) {
// }

func (d *BasicDispatcher) Register(obj Object, event string, f MessageHandler) {
	dis := obj.GetDispatcher().(*BasicDispatcher)

	rcvr, found := dis.EventMap[obj.(*Obj)]
	if !found {
		rcvr = make([]MessageReceiver, 0)
	}
	rcvr = append(rcvr, MessageReceiver{d.Owner, f})
	dis.EventMap[obj.(*Obj)] = rcvr
}

func main() {
	fmt.Println("hello")
	defer fmt.Println("goodbye")

	// funcs["OnCollision"](o, &CollisionInfo{1, 2})
	o1 := &Obj{"bill", &BasicDispatcher{EventMap: make(map[Object][]MessageReceiver)}}
	o1.Dispatcher.(*BasicDispatcher).Owner = o1

	o2 := &Obj{"bob", &BasicDispatcher{}}
	o2.Dispatcher.(*BasicDispatcher).Owner = o2
	o2.Register(o1, "OnCollision", func(o Object, m Message) {
		c := m.(*CollisionInfo)
		o.(*Obj).HandleCollision(c)
	})

	ui := &UISystem{"uiSys", &BasicDispatcher{}}
	ui.Dispatcher.(*BasicDispatcher).Owner = ui
	ui.Register(o1, "OnCollision", func(o Object, m Message) {
		c := m.(*CollisionInfo)
		o.(*UISystem).Update(c)
	})

	o1.Dispatch("OnCollision", &CollisionInfo{1, 2})
}
