package coreType

// Spaces just hold things and know how to delegate to the things they contain.
type Space interface {
	EventHandler
	// SubSpaceManager - Too specific? SpaceManagerComponent, would allow components to be here as well.
	//                   Components shouldn't have subspaces, though.
	SubSpaceManager
	Initializer
	Identifier
	ComponentManager
}

type SubSpaceManager interface {
	AddSubSpace(Space)
	GetSubSpace(string) (Space, error)
	// RemSubSpace(string)
}

// type BasicSpace struct {
// 	*IdentifierComponent
// 	*SpaceManagerComponent
// 	*EventHandlerComponent
// 	// TODO: move this into GOC
// 	*ComponentManagerComponent
// }

// func (bs *BasicSpace) Construct() error {
// 	bs.IdentifierComponent = &IdentifierComponent{}

// 	bs.SpaceManagerComponent = &SpaceManagerComponent{}
// 	bs.SpaceManagerComponent.Construct()

// 	bs.EventHandlerComponent = &EventHandlerComponent{}
// 	bs.EventHandlerComponent.Construct()

// 	bs.ComponentManagerComponent = &ComponentManagerComponent{}
// 	bs.ComponentManagerComponent.Construct()
// 	bs.ComponentManagerComponent.Owner = bs

// 	return nil
// }

// func (bs *BasicSpace) Destruct() error {
// 	return nil
// }
