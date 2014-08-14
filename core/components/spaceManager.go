package components

import (
	"github.com/jmorgan1321/golang-games/core/types"
)

type SpaceManagerComponent struct {
	spaceMap map[string]coreType.Space
}

func (smc *SpaceManagerComponent) Construct() error {
	smc.spaceMap = make(map[string]coreType.Space)
	return nil
}

// AddSubSpace inserts the coreType.Space into the SpaceManagerComponent, using coreType.Space's
// name as the key.
func (smc *SpaceManagerComponent) AddSubSpace(s coreType.Space) {
	// TODO: check for duplicates
	smc.spaceMap[s.Name()] = s
}

func (smc *SpaceManagerComponent) GetSubSpace(name string) (coreType.Space, error) {
	return smc.spaceMap[name], nil
}

func (*SpaceManagerComponent) IsComponent() {}
