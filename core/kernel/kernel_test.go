package kernel

import (
	"fmt"
	"testing"

	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/support"
	"github.com/jmorgan1321/golang-games/core/test"
	"github.com/jmorgan1321/golang-games/core/utils"
)

func init() {
	debug.SilentMode = true
}

func TestKernel_New(t *testing.T) {
	core := New()
	test.ExpectEQ(t, core.State, Running, "Core defaults to running.")
	test.ExpectEQ(t, 0, core.CurrFrame, "Core starts on frame 0.")

	// TODO: decide if I should test internals... hmmm
	test.ExpectEQ(t, core.managers, []Manager(nil), "Core starts with no managers.")
}

func TestKernelState(t *testing.T) {
	test.ExpectEQ(t, "Kernel State: Stopped", Stopped.String(), "Stopped.String() failed.")
	test.ExpectEQ(t, "Kernel State: Terminated", Terminated.String(), "Terminated.String() failed.")
	test.ExpectEQ(t, "Kernel State: Rebooting", Rebooting.String(), "Rebooting.String() failed.")
	test.ExpectEQ(t, "Kernel State: Running", Running.String(), "Running.String() failed.")
}

// Global helper vars for tests
var (
	actOut = []string{}
)

type frameControlMgr struct {
	count int
	core  *Core
}

func (tfm *frameControlMgr) StartUp(cfg GameObject) {}
func (tfm *frameControlMgr) ShutDown()              {}
func (tfm *frameControlMgr) BeginFrame() {
	tfm.count++
	if tfm.count == 4 {
		tfm.core.State = Terminated
	}
}
func (tfm *frameControlMgr) EndFrame() {
}

func TestCore_Run(t *testing.T) {
	core := New()

	core.State = Rebooting
	test.ExpectEQ(t, utils.ES_Restart, core.Run(), "The core should have returned a restart value.")
	test.ExpectEQ(t, 0, core.CurrFrame, "No frame ran.")

	core.State = Terminated
	test.ExpectEQ(t, utils.ES_Success, core.Run(), "The core should have return quit value.")
	test.ExpectEQ(t, 0, core.CurrFrame, "No frame ran.")

	core.State = Stopped
	test.ExpectEQ(t, utils.ES_Success, core.Run(), "The core should return quit")
	test.ExpectEQ(t, 1, core.CurrFrame, "One frame ran.")

	core.RegisterManager(&frameControlMgr{core: core})
	core.StartUp(nil)

	core.State = Running
	test.ExpectEQ(t, utils.ES_Success, core.Run(), "The core should return quit")
	test.ExpectEQ(t, 5, core.CurrFrame, "Core ran until it was stopped.")
}

/*
Manager Tests
*/

type TestManager struct {
	name, comp string
}
type TestMgr1InitComponent struct {
	OwnerMngr
	Test, Type string
}
type TestMgr2InitComponent struct {
	OwnerMngr
	Test, Type string
}

func (tm *TestManager) StartUp(cfg GameObject) {
	actOut = append(actOut, tm.name+" StartUp()")

	// test hack to make sure each manager got proper config
	switch d := cfg.GetComp(tm.comp).(type) {
	case *TestMgr1InitComponent:
		if d.Test == tm.name {
			actOut = append(actOut, tm.name+" got config")
		}
	case *TestMgr2InitComponent:
		if d.Test == tm.name {
			actOut = append(actOut, tm.name+" got config")
		}
	}
}

func (tm *TestManager) ShutDown() {
	actOut = append(actOut, tm.name+" ShutDown()")
}
func (tm *TestManager) BeginFrame() {
	actOut = append(actOut, tm.name+" BeginFrame()")
}
func (tm *TestManager) EndFrame() {
	actOut = append(actOut, tm.name+" EndFrame()")
}

func TestCore_Managers(t *testing.T) {
	core := New()

	// Test Setup
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

	// Testing Core.StartUp()
	actOut = nil
	expOut := []string{
		"Mgr1 StartUp()",
		"Mgr1 got config",
		"Mgr2 StartUp()",
		"Mgr2 got config",
	}
	core.StartUp(cfg)
	test.AssertEQ(t, len(expOut), len(actOut), "Length of actOut and expOut aren't equal. "+fmt.Sprintf("\n\t\tactOut: %v", actOut))
	for i, s := range actOut {
		test.ExpectEQ(t, expOut[i], s, fmt.Sprintf("With test %d: Initialization order was wrong.", i))
	}

	// Testing Core.Run()
	actOut = nil
	expOut = []string{
		"Mgr1 BeginFrame()",
		"Mgr2 BeginFrame()",
		"Mgr2 EndFrame()",
		"Mgr1 EndFrame()",
	}
	core.State = Stopped
	core.Run()
	test.AssertEQ(t, len(expOut), len(actOut), "Length of actOut and expOut aren't equal.")
	for i, s := range actOut {
		test.ExpectEQ(t, expOut[i], s, fmt.Sprintf("With test %d: Frame Update order was wrong.", i))
	}

	// Testing Core.ShutDown()
	actOut = nil
	expOut = []string{
		"Mgr2 ShutDown()",
		"Mgr1 ShutDown()",
	}
	core.ShutDown()
	test.AssertEQ(t, len(expOut), len(actOut), "Length of actOut and expOut aren't equal.")
	for i, s := range actOut {
		test.ExpectEQ(t, expOut[i], s, fmt.Sprintf("With test %d: DeInitialization order was wrong.", i))
	}
}

/*
Space Tests
*/

type TestSpace struct {
	name string
}

func (ts *TestSpace) Init() {
	actOut = append(actOut, ts.name+" Init()")
}
func (ts *TestSpace) DeInit() {
	actOut = append(actOut, ts.name+" DeInit()")
}

func (ts *TestSpace) Update(dt float32) {
	actOut = append(actOut, ts.name+" Update()")
}

func (ts *TestSpace) AddGoc(*Goc) {
}

func TestCore_Spaces(t *testing.T) {
	core := New()

	// Space Setup
	Spc1 := &TestSpace{name: "Spc1"}
	Spc2 := &TestSpace{name: "Spc2"}

	core.RegisterSpace(Spc1)
	core.RegisterSpace(Spc2)

	// Testing core.StartUp()
	actOut = nil
	expOut := []string{
		"Spc1 Init()",
		"Spc2 Init()",
	}

	core.StartUp(nil)
	test.AssertEQ(t, len(expOut), len(actOut), "Length of actOut and expOut aren't equal.")
	for i, s := range actOut {
		test.ExpectEQ(t, expOut[i], s, fmt.Sprintf("With test %d: Initialization order was wrong.", i))
	}

	// Testing Core.Run()
	actOut = nil
	expOut = []string{
		"Spc1 Update()",
		"Spc2 Update()",
	}
	core.State = Stopped
	core.Run()
	test.AssertEQ(t, len(expOut), len(actOut), "Length of actOut and expOut aren't equal.")
	for i, s := range actOut {
		test.ExpectEQ(t, expOut[i], s, fmt.Sprintf("With test %d: Frame Update order was wrong.", i))
	}

	// Testing core.ShutDown()
	actOut = nil
	expOut = []string{
		"Spc2 DeInit()",
		"Spc1 DeInit()",
	}

	core.ShutDown()
	test.AssertEQ(t, len(expOut), len(actOut), "Length of actOut and expOut aren't equal.")
	for i, s := range actOut {
		test.ExpectEQ(t, expOut[i], s, fmt.Sprintf("With test %d: DeInitialization order was wrong.", i))
	}
}
