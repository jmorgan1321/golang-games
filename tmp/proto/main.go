package main

import (
	"fmt"
	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/kernel"
	"github.com/jmorgan1321/golang-games/core/support"
)

func main() {
	debug.IndentationLevel = new(debug.IndentLevel)

	fmt.Println("starting game")
	defer fmt.Println("ending game")

	debug.Trace()
	defer debug.UnTrace()

	// ResourceManager doesn't exist yet, so use whole path name.
	core := kernel.New("D:/work/fbl_grfx_dev_p/windows/sandbox/go_tools/rsrc/json/objects/GameCoreConfig.jrm")

	core.Startup()
	defer core.Shutdown()

	core.Run()
}

func init() {
	kernel.CoreTempFactoryFunc = FactoryFunc
}

type FactoryInitComp struct {
	DummyVar bool
	Type     string
}

type ResourceMgrInitComp struct {
	FavoriteColor string
	Type          string
}

// Factory func is where anything that can be created by the factory gets
// registered.
//
// This allows the factory to create arbitrary objects without it's package
// knowing anything about them (ie, the factory doesn't depend on the component
// packages).
//
// This is a quick hack until I can figure out a better way to register data
// with the factory.
//
func FactoryFunc(name string) kernel.Component {
	debug.Trace()
	defer debug.UnTrace()

	var c kernel.Component
	switch name {
	default:
		support.LogFatal("Unknown componnt passed into factory func: %s", name)
	case "FactoryInitComp":
		c = &FactoryInitComp{}
	case "ResourceMgrInitComp":
		c = &ResourceMgrInitComp{}
	}
	return c
}
