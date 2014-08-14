package main


}

// package core

// import (
// 	"github.com/jmorgan1321/golang-games/core/debug"
// 	"testing"

// 	"github.com/jmorgan1321/golang-games/core/components"
// 	"github.com/jmorgan1321/golang-games/core/kernel"
// 	"github.com/jmorgan1321/golang-games/core/types"
// )

// var core *kernel.Core

// func TestCore(t *testing.T) {
// 	core = kernel.New()
// 	objMgr := core.ObjectMgr

// 	s := objMgr.CreateSpace(nil, "levelSpace", "_")
// 	debug.ExpectEQ(t, "levelSpace", s.Name(), "Name not set correctly.")
// 	spaceFromMgr, _ := objMgr.GetByName("levelSpace")
// 	debug.ExpectEQ(t, s, spaceFromMgr, "levelSpace not stored correctly.")

// 	g1 := objMgr.CreateSpace(s, "goc1", "GocFile")
// 	debug.ExpectEQ(t, "goc1", g1.Name(), "Name not set correctly.")
// 	gocFromMgr, _ := objMgr.GetByName("goc1")
// 	debug.ExpectEQ(t, g1, gocFromMgr, "goc1 not stored correctly.")
// 	gocFromSpace, _ := s.GetSubSpace(g1.Name())
// 	debug.ExpectEQ(t, g1, gocFromSpace, "goc1 not retrieved from levelSpace.")

// 	g2 := objMgr.CreateSpace(s, "goc2", "GocFile")
// 	debug.ExpectEQ(t, "goc2", g2.Name(), "Name not set correctly.")
// 	goc2FromMgr, _ := objMgr.GetByName("goc2")
// 	debug.ExpectEQ(t, g2, goc2FromMgr, "goc2 not stored correctly.")
// 	goc2FromSpace, _ := s.GetSubSpace(g2.Name())
// 	debug.ExpectEQ(t, g2, goc2FromSpace, "goc2 not retrieved from levelSpace.")
// 	debug.ExpectNEQ(t, gocFromSpace, goc2FromSpace, "gocs should be different.")

// 	testInt := 0
// 	test2Int := 0

// 	g2.RegisterForEvent("MyEvent", g1, func(e coreType.EventData) {
// 		testInt = e.(int)
// 	})
// 	g2.RegisterForEvent("AnotherEvent", g1, func(e coreType.EventData) {
// 		test2Int = e.(int)
// 	})

// 	g1.TriggerEvent("MyEvent", 5)
// 	g1.TriggerEvent("AnotherEvent", 10)

// 	debug.ExpectEQ(t, 5, testInt, "Messaging didn't work")
// 	debug.ExpectEQ(t, 10, test2Int, "Registering multiple messages didn't work")

// 	ehc, _ := g1.GetComponent("EventHandlerComponent")
// 	ehc.(*components.EventHandlerComponent).TriggerEvent("MyEvent", 6)
// 	debug.ExpectEQ(t, 6, testInt, "GetComponent() failed for static component")
// }
