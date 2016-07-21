package kvdb

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackSuite struct {
	suite.Suite
	Tact *Transaction
}

func (suite *StackSuite) SetupTest() {
	suite.Tact = &Transaction{}
}

func (suite *StackSuite) TestSomething() {
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}
