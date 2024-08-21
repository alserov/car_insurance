package mongo

import (
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	suite.Run(t, new(mongoRepoIntegrationSuite))
}

type mongoRepoIntegrationSuite struct {
	suite.Suite

	insuranceValues []models.Insurance

	repo db.Repository
}

func (s *mongoRepoIntegrationSuite) SetupTest() {
	s.insuranceValues = []models.Insurance{
		{
			SenderAddr: "x01",
			Amount:     100,
			ActiveTill: time.Now().Add(time.Hour),
		},
		{
			SenderAddr: "x02",
			Amount:     100,
			ActiveTill: time.Now().Add(time.Hour * 2),
		},
	}

	//conn := MustConnect("localhost:")
}

func (s *mongoRepoIntegrationSuite) TestCreate(t *testing.T) {

}

func (s *mongoRepoIntegrationSuite) TestGet(t *testing.T) {

}

func (s *mongoRepoIntegrationSuite) TestDelete(t *testing.T) {

}

//func (s *mongoRepoIntegrationSuite) newMongoInstance() testcontainers.Container {
//	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
//		ContainerRequest: testcontainers.ContainerRequest{
//			Image:        "postgres",
//			ExposedPorts: []string{"5432/tcp"},
//			Env: map[string]string{
//				"POSTGRES_USER":     p.user,
//				"POSTGRES_PASSWORD": p.password,
//				"POSTGRES_DB":       p.db,
//			},
//			WaitingFor: wait.ForAll(
//				wait.ForLog("database system is ready to accept connections"),
//				wait.ForListeningPort("5432/tcp"),
//			),
//		},
//		Started: true,
//	})
//	require.NoError(p.T(), err)
//
//	return container
//}
