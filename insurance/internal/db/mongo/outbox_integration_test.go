package mongo

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestIntegration(t *testing.T) {
	suite.Run(t, new(mongoRepoIntegrationSuite))
}

type mongoRepoIntegrationSuite struct {
	suite.Suite
}
