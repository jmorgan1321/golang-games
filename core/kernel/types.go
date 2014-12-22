package kernel

import "reflect"

// Spaces are containers of Gocs and Systems and can be treated like levels in
// a game context.
//
type Space interface {
	Init()
	DeInit()
	Update(dt float32)
	AddGoc(goc *Goc)
}

// Managers are long-running game operations/processes that handle operations
// outside of a space.
//
// An example would be audio; mulitple spaces can play sounds, so they would
// delegate that operation to the core's one audio manager.
//
type Manager interface {
	StartUp(config GameObject)
	ShutDown()
	BeginFrame()
	EndFrame()
}

// Components hold data for GameObjects, which differentiates game
// objects from each other.
type Component interface {
	SetOwner(GameObject)
	// Owner() GameObject

	// Init
	// Deinit
}

// GameObjects are containers for Components.
type GameObject interface {
	ComponentManager
}

// Goc (aka Game Object Composition) is a concrete implementation of GameObject.
type Goc struct {
	Name       string
	Components []Component
}

func (goc *Goc) AddComp(c Component) {
	goc.Components = append(goc.Components, c)
	c.SetOwner(goc)
}

// GetComp searches this goc for a component with the matching and returns it
// as a generic component, which can then be casted to it's actual type.
//
// ie:
//  alist := goc.GetComp("ActionListComponent").(*ActionListComponent)
func (goc *Goc) GetComp(name string) Component {
	for _, comp := range goc.Components {
		// TODO: possibly switch this to use package.Name
		ctype := reflect.ValueOf(comp).Type().Elem()
		if ctype.Name() == name {
			return comp
		}
	}
	// TODO: better error message when component isn't found
	return nil
}

func (goc *Goc) Unmarshal(data interface{}) {
	m := data.(map[string]interface{})

	// Unmarshall all "normal" fields
	for k, v := range m {
		if k == "Type" || k == "Components" {
			continue
		}

		f := reflect.ValueOf(goc).Elem().FieldByName(k)
		SerializeInPlace(f, v)
	}

	// Special case handling for `Components`
	for _, v := range m["Components"].([]interface{}) {
		compData := v.(map[string]interface{})

		typename, _ := compData["Type"]
		comp := factoryFunc(typename.(string)).(Component)
		SerializeInPlace(comp, compData)
		goc.AddComp(comp)
	}
}

type ComponentManager interface {
	GetComp(name string) Component
	AddComp(Component)
}

type OwnerMngr struct {
	owner GameObject
}

func (o *OwnerMngr) SetOwner(g GameObject) {
	o.owner = g
}
func (o *OwnerMngr) Owner() GameObject {
	return o.owner
}
