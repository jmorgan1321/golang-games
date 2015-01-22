package event

type Data interface{}
type Handler func(Data)

type Dispatcher interface {
	Register(event string, sender Dispatcher, handler Handler)
	Dispatch(event string, data Data)
	addListener(event string, handler Handler) *tracker
}
