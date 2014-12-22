package kernel

import (
	"fmt"
	"github.com/jmorgan1321/golang-games/core/support"
)

func Example_core() {
	actOut = nil
	core := New()

	Mgr1 := &TestManager{name: "Mgr1", comp: "TestMgr1InitComponent"}
	Mgr2 := &TestManager{name: "Mgr2", comp: "TestMgr2InitComponent"}

	CoreFactoryFunc = func(name string) interface{} {
		var comp interface{}
		switch name {
		default:
			support.LogFatal("unknown component: %s", name)
		case "TestMgr1InitComponent":
			comp = &TestMgr1InitComponent{}
		case "TestMgr2InitComponent":
			comp = &TestMgr2InitComponent{}
		}
		return comp
	}

	core.RegisterManager(Mgr1)
	core.RegisterManager(Mgr2)

	cfg := LoadConfig("C:/Users/jmorgan/Sandbox/golang/rsrc/test/json/objects/kernel/TestCore.jrm")

	Spc1 := &TestSpace{name: "Spc1"}
	Spc2 := &TestSpace{name: "Spc2"}
	core.RegisterSpace(Spc1)
	core.RegisterSpace(Spc2)

	core.StartUp(cfg)

	core.State = Stopped
	core.Run()

	core.ShutDown()

	// sort.Reverse(sort.StringSlice(actOut))
	for _, s := range actOut {
		fmt.Println(s)
	}
	// Output:
	//
	// Mgr1 StartUp()
	// Mgr1 got config
	// Mgr2 StartUp()
	// Mgr2 got config
	// Spc1 Init()
	// Spc2 Init()
	// Mgr1 BeginFrame()
	// Mgr2 BeginFrame()
	// Spc1 Update()
	// Spc2 Update()
	// Mgr2 EndFrame()
	// Mgr1 EndFrame()
	// Spc2 DeInit()
	// Spc1 DeInit()
	// Mgr2 ShutDown()
	// Mgr1 ShutDown()
}
