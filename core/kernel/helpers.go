package kernel

import (
	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/support"
)

func LoadConfig(file string) GameObject {
	debug.Trace()
	defer debug.UnTrace()

	data, err := support.OpenFile(file)
	if err != nil {
		support.LogFatal("Failed to open Config file: " + file)
	}

	holder, err := support.ReadData(data)
	if err != nil {
		support.LogFatal("Failed to read in Config file: " + file)
		return nil
	}

	var goc *Goc
	Serialize2(&goc, holder)
	return goc
}
