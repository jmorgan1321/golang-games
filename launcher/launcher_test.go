package main

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type LauncherSuite struct{}

var _ = Suite(&LauncherSuite{})

// func (s *LauncherSuite) Test_IfGamePathNotFoundLauncherExits(c *C) {
// }

// func (s *LauncherSuite) Test_RestartingGame(c *C) {
// }

// func (s *LauncherSuite) Test_OpeningBrowser(c *C) {
// }
