package main

import (
	"log"

	"github.com/jmorgan1321/golang-games/lib/core"
	"github.com/jmorgan1321/golang-games/lib/engine/types"
	"github.com/jmorgan1321/golang-games/lib/factory"
	"github.com/jmorgan1321/golang-games/lib/graphics"
	"github.com/jmorgan1321/golang-games/lib/input"

	"gopkg.in/qml.v1"
)

/*
* Call qml.Run from function main providing a function with the logic below
* Create an engine for loading and running QML content (see NewEngine)
* Make Go values and types available to QML (see Context.SetVar and RegisterType)
* Load QML content (see Engine.LoadString and Engine.LoadFile)
* Create a new window for the content (see Component.CreateWindow)
* Show the window and wait for it to be closed (see Window.Show and Window.Wait)
 */

func init() {

}

func main() {
	err := qml.Run(runGame)
	if err != nil {
		log.Fatal(err)
	}
}

func runGame() error {
	engine := qml.NewEngine()
	component, err := engine.LoadFile("file2.qml")
	if err != nil {
		log.Fatal(err)
	}
	win := component.CreateWindow(nil)

	// init control object (used for communicating with the QML code)
	// and pass it to the QML code.
	core.Core = core.New()
	core.Core.Context = engine.Context()
	core.Core.Context.SetVar("gameCore", core.Core)

	core.Core.RegisterManagers(
		factory.New(),
		input.NewManager(),
		graphics.NewManager(),
		graphics.NewDebugDrawer(),
	)
	core.Core.Factory = core.Core.Manager("Factory").(*factory.Factory)
	// TODO(jemorgan): remove implementation detail
	core.Core.DebugDraw = core.Core.Manager("QmlDebugDrawer").(core.DebugDrawer)
	core.Core.RootSpace = &types.BasicSpace{}

	// needed as a backdoor to factory not being inititalized
	cfg := core.LoadConfig("Pong.jrm")
	grfx := cfg.Comp("graphics.Config").(*graphics.Config)
	grfx.Window = win
	grfx.Engine = engine
	core.Core.StartUp(cfg)

	player := core.Core.Factory.Create("player.jrm")
	player.Init()

	win.Show()
	Timer := win.Root().ObjectByName("t1")
	Timer.Set("running", true)
	win.Wait()

	core.Core.ShutDown()
	return nil
}
