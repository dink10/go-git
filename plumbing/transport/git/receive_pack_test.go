package git

import (
	"github.com/dink10/go-git.v4/plumbing/transport/test"

	"github.com/dink10/go-git-fixtures.v3"
	. "gopkg.in/check.v1"
)

type ReceivePackSuite struct {
	test.ReceivePackSuite
	BaseSuite
}

var _ = Suite(&ReceivePackSuite{})

func (s *ReceivePackSuite) SetUpTest(c *C) {
	s.BaseSuite.SetUpTest(c)

	s.ReceivePackSuite.Client = DefaultClient
	s.ReceivePackSuite.Endpoint = s.prepareRepository(c, fixtures.Basic().One(), "basic.git")
	s.ReceivePackSuite.EmptyEndpoint = s.prepareRepository(c, fixtures.ByTag("empty").One(), "empty.git")
	s.ReceivePackSuite.NonExistentEndpoint = s.newEndpoint(c, "non-existent.git")

	s.StartDaemon(c)
}
