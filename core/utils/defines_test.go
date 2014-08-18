package utils

import (
	"github.com/jmorgan1321/golang-games/core/test"
	"testing"
)

func Test_Defines(t *testing.T) {
	// TODO: figure out why this test fails

	test.ExpectEQ(t, 0, ES_Success, "ES_Success was wrong")
	test.ExpectEQ(t, 1, ES_Restart, "ES_Restart was wrong")

	test.ExpectEQ(t, 0.025, Epsilon, "Epsilon was wrong")
}
