package utils

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type DefinesSpecSuite struct{}

var _ = Suite(&DefinesSpecSuite{})

func (s *DefinesSpecSuite) Test_Defines(c *C) {
	c.Assert(0, Equals, ES_Success)
	c.Assert(1, Equals, ES_Restart)
}
