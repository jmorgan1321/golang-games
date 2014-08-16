package kernel

import (
	"reflect"
)

type GameObject interface {
	GetComp(name string) Component
	AddComp(Component)
}

type Goc struct {
	Name       string
	Components []Component
}

func (goc *Goc) AddComp(c Component) {
	goc.Components = append(goc.Components, c)
}
func (goc *Goc) GetComp(name string) Component {
	for _, comp := range goc.Components {
		ctype := reflect.ValueOf(comp).Type().Elem()
		if ctype.Name() == name {
			return comp
		}
	}
	// TODO: better error message when component isn't found
	return nil
}
