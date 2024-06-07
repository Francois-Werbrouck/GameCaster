package test_test

import (
	"github.com/stretchr/testify/suite"
)

const DEFAULT_TESTING_PORT = "13455"

type TestSuite struct {
	suite.Suite
}

func (tests *TestSuite) SetupAllSuite() {

}

func (tests *TestSuite) TestAssert() {
	tests.Greater(2, 1)
}
