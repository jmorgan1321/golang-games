package utils

import (
	"testing"
)

func TestEpsilonCompare(t *testing.T) {
	if EpsilonCompare(0, 0) == false {
		t.Error("number within epsilon are equal")
	}
	if EpsilonCompare(0, 0+Epsilon) {
		t.Error("number outside of pos epsilon are NOT be equal")
	}
	if EpsilonCompare(0, 0-Epsilon) {
		t.Error("number outside of neg epsilon are NOT be equal")
	}
}
