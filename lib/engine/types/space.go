package types

import (
	"time"

	"github.com/jmorgan1321/golang-games/lib/engine/event"
	"github.com/jmorgan1321/golang-games/lib/engine/meta"
)

func init() {
	meta.Register((*BasicSpace)(nil),
		meta.Init(func() interface{} { return &BasicSpace{} }),
	)
}

type JSpace interface {
	JGameObject
	JSystemManager
	JGameObjectManager
	Update(time.Duration)
}

type BasicSpace struct {
	event.InstantDispatcher
	JBasicGameObjectMgr
	JBasicSystemMgr
	JBasicCompMgr
	JBaseObject
}

func (s *BasicSpace) Update(dt time.Duration) {
	for _, sys := range s.JBasicSystemMgr.Objects {
		sys.(JSystem).Update(dt)
	}
}
