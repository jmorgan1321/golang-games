package kernel

import (
	"testing"

	"github.com/jmorgan1321/golang-game/lib/engine/test"
)

// types for testing GetCompInterface
type (
	gocTestCompInterface interface {
		getTestCompType() string
	}

	gocTestComp1 struct {
		OwnerMngr
	}
	gocTestComp2 struct {
		OwnerMngr
	}
	gocTestComp3 struct {
		OwnerMngr
	}
)

func (g *gocTestComp1) getTestCompType() string { return "gocTestComp1" }
func (g *gocTestComp2) getTestCompType() string { return "gocTestComp2" }

func TestGocsCanGetCompByInterface(t *testing.T) {
	tests := []struct {
		c []Component
		b bool
		s string
	}{
		{nil, false, ""},
		{[]Component{&gocTestComp3{}}, false, ""},
		{[]Component{&gocTestComp3{}, &gocTestComp1{}}, true, "gocTestComp1"},
		{[]Component{&gocTestComp2{}, &gocTestComp3{}}, true, "gocTestComp2"},
		{[]Component{&gocTestComp2{}, &gocTestComp1{}}, true, "gocTestComp2"},
	}

	for i, tt := range tests {
		goc := &Goc{}
		goc.AddComps(tt.c...)
		cmp := goc.GetCompInterface(new(gocTestCompInterface))

		if cmp == nil && tt.b {
			test.Expect(t, false, "test %v: failed to comp of interface")
		} else if cmp != nil && !tt.b {
			test.Expect(t, false, "test %v: found erroneous comp of interface")
		} else if cmp != nil {
			test.Expect(t, tt.s == cmp.(gocTestCompInterface).getTestCompType(),
				"test %v: wrong type found", i)
		}
	}
	goc := &Goc{}
	goc.AddComps(&gocTestComp3{})

	cmp3 := goc.GetCompInterface(new(gocTestCompInterface))
	test.Assert(t, cmp3 == nil, "unexpectedly returned result")

	goc.AddComps(&gocTestComp1{})
	cmp1 := goc.GetCompInterface(new(gocTestCompInterface)).(*gocTestComp1)
	test.ExpectEQ(t, "gocTestComp1", cmp1.getTestCompType(),
		"failed to get component by interface (1)")

	goc2 := &Goc{}
	goc2.AddComps(&gocTestComp2{})
	goc2.AddComps(&gocTestComp3{})
	cmp2 := goc2.GetCompInterface(new(gocTestCompInterface)).(*gocTestComp2)
	test.ExpectEQ(t, "gocTestComp2", cmp2.getTestCompType(),
		"failed to get component by interface (2)")
}
