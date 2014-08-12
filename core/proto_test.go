package core

import (
	"testing"

	"github.com/jmorgan1321/golang-games/core/components"
	"github.com/jmorgan1321/golang-games/core/kernel"
	"github.com/jmorgan1321/golang-games/core/types"
	"github.com/jmorgan1321/golang-games/core/utils"
)

var core *kernel.Core

func TestCore(t *testing.T) {
	core = kernel.New()
	objMgr := core.ObjectMgr

	s := objMgr.CreateSpace(nil, "levelSpace", "_")
	utils.ExpectEQ(t, "levelSpace", s.Name(), "Name not set correctly.")
	spaceFromMgr, _ := objMgr.GetByName("levelSpace")
	utils.ExpectEQ(t, s, spaceFromMgr, "levelSpace not stored correctly.")

	g1 := objMgr.CreateSpace(s, "goc1", "GocFile")
	utils.ExpectEQ(t, "goc1", g1.Name(), "Name not set correctly.")
	gocFromMgr, _ := objMgr.GetByName("goc1")
	utils.ExpectEQ(t, g1, gocFromMgr, "goc1 not stored correctly.")
	gocFromSpace, _ := s.GetSubSpace(g1.Name())
	utils.ExpectEQ(t, g1, gocFromSpace, "goc1 not retrieved from levelSpace.")

	g2 := objMgr.CreateSpace(s, "goc2", "GocFile")
	utils.ExpectEQ(t, "goc2", g2.Name(), "Name not set correctly.")
	goc2FromMgr, _ := objMgr.GetByName("goc2")
	utils.ExpectEQ(t, g2, goc2FromMgr, "goc2 not stored correctly.")
	goc2FromSpace, _ := s.GetSubSpace(g2.Name())
	utils.ExpectEQ(t, g2, goc2FromSpace, "goc2 not retrieved from levelSpace.")
	utils.ExpectNEQ(t, gocFromSpace, goc2FromSpace, "gocs should be different.")

	testInt := 0
	test2Int := 0

	g2.RegisterForEvent("MyEvent", g1, func(e types.EventData) {
		testInt = e.(int)
	})
	g2.RegisterForEvent("AnotherEvent", g1, func(e types.EventData) {
		test2Int = e.(int)
	})

	g1.TriggerEvent("MyEvent", 5)
	g1.TriggerEvent("AnotherEvent", 10)

	utils.ExpectEQ(t, 5, testInt, "Messaging didn't work")
	utils.ExpectEQ(t, 10, test2Int, "Registering multiple messages didn't work")

	ehc, _ := g1.GetComponent("EventHandlerComponent")
	ehc.(*components.EventHandlerComponent).TriggerEvent("MyEvent", 6)
	utils.ExpectEQ(t, 6, testInt, "GetComponent() failed for static component")
}
