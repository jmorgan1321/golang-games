package components

import (
	"github.com/jmorgan1321/golang-games/core/types"
)

type SpaceManagerComponent struct {
	spaceMap map[string]types.Space
}

func (smc *SpaceManagerComponent) Construct() error {
	smc.spaceMap = make(map[string]types.Space)
	return nil
}

// AddSubSpace inserts the types.Space into the SpaceManagerComponent, using types.Space's
// name as the key.
func (smc *SpaceManagerComponent) AddSubSpace(s types.Space) {
	// TODO: check for duplicates
	smc.spaceMap[s.Name()] = s
}

func (smc *SpaceManagerComponent) GetSubSpace(name string) (types.Space, error) {
	return smc.spaceMap[name], nil
}

func (*SpaceManagerComponent) IsComponent() {}
