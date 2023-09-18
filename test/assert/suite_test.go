package assert

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {

}

// The SetupTest method will be run before every test in the suite.
func (s *Suite) SetupTest() {

}

// The TearDownTest method will be run after every test in the suite.
func (s *Suite) TearDownTest() {

}

func (s *Suite) TearDownSuite() {

}
