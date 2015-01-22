package kernel

import (
	"github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/support"
)

// LoadConfig serialize in a GameObject from a text file.
func LoadConfig(file string) GameObject {
	defer debug.Trace().UnTrace()

	data, err := support.OpenFile(file)
	if err != nil {
		support.LogFatal("Failed to open Config file: " + file)
	}

	holder, err := support.ReadData(data)
	if err != nil {
		support.LogFatal("Failed to read in Config file: " + file)
		return nil
	}

	goc := Goc{}
	SerializeInPlace(&goc, holder)
	return &goc
}
