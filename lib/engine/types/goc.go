package types

import (
	"reflect"

	"github.com/jmorgan1321/golang-games/lib/engine/event"
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
	"github.com/jmorgan1321/golang-games/lib/engine/serialization"
)

func init() {
	meta.Register((*JGoc)(nil),
		meta.Init(func() interface{} { return &JGoc{} }),
	)
}

type JGoc struct {
	JBaseObject
	JBasicCompMgr
	event.InstantDispatcher
}

func (goc *JGoc) Init() {
	for _, c := range goc.Objects {
		c.(JComponent).SetOwner(goc)
		c.(JComponent).Init()
	}
}
func (goc *JGoc) Deinit() {
	for _, c := range goc.Objects {
		c.(JComponent).Deinit()
	}
}

func (goc *JGoc) Unmarshal(data interface{}) {
	m := data.(map[string]interface{})

	// Unmarshall all "normal" fields
	for k, v := range m {
		if k == "Type" || k == "Components" {
			continue
		}

		if k == "Name" {
			goc.SetName(v.(string))
			continue
		}

		f := reflect.ValueOf(goc).Elem().FieldByName(k)
		serialization.SerializeInPlace(f, v)
	}

	// Special case handling for `Components`
	for _, v := range m["Components"].([]interface{}) {
		compData := v.(map[string]interface{})

		typename, _ := compData["Type"]
		comp := serialization.FactoryFunc(typename.(string)).(JComponent)
		serialization.SerializeInPlace(comp, compData)
		// comp.SetOwner(goc)
		goc.AddComps(comp)
	}
}

// type JGocManager interface {
// 	AddGocs(...JGoc)
// 	RemGoc(string)
// 	Goc(string) *JGoc
// }
