package utils

import (
	"github.com/jmorgan1321/golang-games/core/test"
	"testing"
)

func TestEpsilonCompare(t *testing.T) {
	test.ExpectEQ(t, true, EpsilonCompare(0, 0), "number within epsilon are equal")
	test.ExpectEQ(t, false, EpsilonCompare(0, 0+Epsilon), "number outside of pos epsilon are NOT be equal")
	test.ExpectEQ(t, false, EpsilonCompare(0, 0-Epsilon), "number outside of neg epsilon are NOT be equal")
}
