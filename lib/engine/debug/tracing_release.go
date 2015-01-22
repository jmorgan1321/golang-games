// +build !debug

package debug

type dummy struct{}

func Trace() dummy {
	return dummy{}
}

func (dummy) UnTrace() {}
