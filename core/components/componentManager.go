package components

import (
	"github.com/jmorgan1321/golang-games/core/types"
	"reflect"
)

type ComponentManagerComponent struct {
	Owner types.Space
}

func (cmc *ComponentManagerComponent) Construct() error {
	return nil
}
func (cmc *ComponentManagerComponent) GetComponent(name string) (interface{}, error) {
	// loop through the struct's fields and set the map
	v := reflect.ValueOf(cmc.Owner)
	typ := v.Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// check for static components
	c := v.Elem().FieldByName(name).Elem().Addr()

	// TODO: check dynmaic components if static component not found

	iface := c.Interface().(types.Component)
	return iface, nil
}

func (*ComponentManagerComponent) IsComponent() {}

type ComponentManager interface {
	GetComponent(string) (interface{}, error)
}
