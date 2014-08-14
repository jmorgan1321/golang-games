package support

import (
	"encoding/json"
	"github.com/jmorgan1321/golang-games/core/debug"
	"io/ioutil"
	"os"
)

func OpenFile(filename string) ([]byte, error) {
	debug.Trace()
	defer debug.UnTrace()

	f, err := os.Open(filename)
	if err != nil {
		return nil, LogError("error:", err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, LogError("error:", err)
	}

	return data, nil
}

func ReadData(data []byte) (interface{}, error) {
	debug.Trace()
	defer debug.UnTrace()

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, LogError("error:", err)
	}
	return v, nil
}
