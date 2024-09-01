package mongo

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	s := mongoRepoIntegrationSuite{}

	s.container = s.newMongoInstance()
	defer func() {
		_ =
			s.container.Terminate(context.Background())
	}()

	suite.Run(t, &s)

}

type mongoRepoIntegrationSuite struct {
	suite.Suite

	insuranceValues []models.Insurance

	repo db.Outbox
	conn *mongo.Client

	container testcontainers.Container
}

func (s *mongoRepoIntegrationSuite) SetupTest() {
	s.insuranceValues = []models.Insurance{
		{
			SenderAddr: "x01",
			Amount:     100,
			ActiveTill: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			SenderAddr: "x02",
			Amount:     100,
			ActiveTill: time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	port, err := s.container.MappedPort(context.Background(), "27017")
	s.Require().NoError(err)

	s.conn = MustConnect(fmt.Sprintf("mongodb://localhost:%d", port.Int()))
	s.repo = NewOutbox(s.conn)
}

func (s *mongoRepoIntegrationSuite) TeardownTest() {
	s.Require().NoError(s.conn.Disconnect(context.Background()))
}

func (s *mongoRepoIntegrationSuite) TestCreate() {
	s.Require().NoError(s.repo.Create(context.Background(), models.OutboxItem{
		ID:      "some random uuid",
		GroupID: models.GroupInsurance,
		Status:  models.Pending,
		Val:     s.insuranceValues[0],
	}))
}

func (s *mongoRepoIntegrationSuite) TestGet() {
	s.Require().NoError(s.repo.Create(context.Background(), models.OutboxItem{
		ID:      "some random uuid",
		GroupID: models.GroupInsurance,
		Status:  models.Pending,
		Val:     s.insuranceValues[0],
	}))

	vals, err := s.repo.Get(context.Background(), models.Pending, models.GroupInsurance)
	s.Require().NoError(err)
	s.Require().GreaterOrEqual(len(vals), 1)

	s.Require().NotEmpty(vals[0].ID)
	s.Require().Equal(models.GroupInsurance, vals[0].GroupID)
	s.Require().Equal(models.Pending, vals[0].Status)
	s.Require().Equal(s.insuranceValues[0], vals[0].Val.(models.Insurance))
}

func (s *mongoRepoIntegrationSuite) TestDelete() {
	s.Require().NoError(s.repo.Delete(context.Background(), "some random uuid"))
}

func (s *mongoRepoIntegrationSuite) newMongoInstance() testcontainers.Container {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo",
			ExposedPorts: []string{"27017/tcp"},
			WaitingFor: wait.ForAll(
				wait.ForListeningPort("27017"),
			),
		},
		Started: true,
	})

	if err != nil {
		panic("failed to start container: " + err.Error())

	}

	return container
}
