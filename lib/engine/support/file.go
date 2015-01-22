package support

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jmorgan1321/golang-games/lib/engine/debug"
)

func OpenFile(filename string) ([]byte, error) {
	defer debug.Trace().UnTrace()

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ReadData(data []byte) (interface{}, error) {
	defer debug.Trace().UnTrace()

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return v, nil
}
