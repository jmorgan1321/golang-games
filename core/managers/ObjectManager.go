package managers

import (
	"encoding/json"
	"github.com/jmorgan1321/golang-games/core/components"
	"github.com/jmorgan1321/golang-games/core/types"
)

type BasicSpace struct {
	*components.IdentifierComponent
	*components.SpaceManagerComponent
	*components.EventHandlerComponent
	// TODO: move this into GOC
	*components.ComponentManagerComponent
}

func (bs *BasicSpace) Construct() error {
	bs.IdentifierComponent = &components.IdentifierComponent{}

	bs.SpaceManagerComponent = &components.SpaceManagerComponent{}
	bs.SpaceManagerComponent.Construct()

	bs.EventHandlerComponent = &components.EventHandlerComponent{}
	bs.EventHandlerComponent.Construct()

	bs.ComponentManagerComponent = &components.ComponentManagerComponent{}
	bs.ComponentManagerComponent.Construct()
	bs.ComponentManagerComponent.Owner = bs

	return nil
}

func (bs *BasicSpace) Destruct() error {
	return nil
}

// ObjectManager creates, destroys, and stores Game Objects.
type ObjectManager struct {
	spaceMap map[string]coreType.Space
}

func (o *ObjectManager) Construct() error {
	o.spaceMap = make(map[string]coreType.Space)
	return nil
}

func (o *ObjectManager) GetByName(name string) (coreType.Space, error) {
	return o.spaceMap[name], nil
}

func (o *ObjectManager) CreateSpace(parent coreType.Space, name, filename string) coreType.Space {
	// TODO: hook core back in
	// spaceJSON, _ := core.ResourceMgr.GetFileData(filename)
	spaceJSON := "{}"
	s := o.createSpaceFromString(spaceJSON)
	s.SetName(name)
	o.spaceMap[name] = s
	if parent != nil {
		parent.AddSubSpace(s)
	}
	return s
}

func (o *ObjectManager) createSpaceFromString(data string) coreType.Space {
	s := &BasicSpace{}
	// TODO: error check
	json.Unmarshal([]byte(data), s)
	// TODO: error check
	s.Construct()
	return s
}
