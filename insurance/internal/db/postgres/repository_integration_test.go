package postgres

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPostgresRepoSuite(t *testing.T) {
	suite.Run(t, new(PostgresRepoSuite))
}

type PostgresRepoSuite struct {
	suite.Suite
}

func (s *PostgresRepoSuite) SetupTest() {

}

func (s *PostgresRepoSuite) TeardownTest() {

}
