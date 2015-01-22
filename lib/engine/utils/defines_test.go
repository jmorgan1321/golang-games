package utils

import (
	"testing"
)

func Test_Defines(t *testing.T) {
	// TODO: figure out why this test fails

	if int(ES_Success) != 0 {
		t.Error("ES_Success was wrong")
	}

	if int(ES_Restart) != 1 {
		t.Error("ES_Restart was wrong")
	}

	if Epsilon != 0.025 {
		t.Error("Epsilon was wrong")
	}
}
