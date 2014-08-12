package kernel

import (
	"github.com/jmorgan1321/golang-games/core/managers"
)

type Core struct {
	ResourceMgr *managers.ResourceManager
	ObjectMgr   *managers.ObjectManager
}

func New() *Core {
	core := &Core{
		ResourceMgr: &managers.ResourceManager{},
		ObjectMgr:   &managers.ObjectManager{},
	}

	core.ResourceMgr.Construct()
	core.ObjectMgr.Construct()

	return core
}
