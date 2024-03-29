package fileutil

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/zonewave/pkgs/mocks/aferomock"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type Suite struct {
	suite.Suite
	ctrl  *gomock.Controller
	afero *aferomock.MockAfero
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (s *Suite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.afero = aferomock.NewMockAfero(s.ctrl)
}

func (s *Suite) TearDownSuite() {
	s.ctrl.Finish()
}

// The SetupTest method will be run before every test in the suite.
func (s *Suite) SetupTest() {
}

// The TearDownTest method will be run after every test in the suite.
func (s *Suite) TearDownTest() {
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
