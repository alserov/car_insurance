package postgres

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

func TestPostgresRepoSuite(t *testing.T) {
	suite.Run(t, new(PostgresRepoSuite))
}

type PostgresRepoSuite struct {
	suite.Suite

	conn *sqlx.DB
	repo db.Repository

	ctx context.Context

	container testcontainers.Container
}

func (s *PostgresRepoSuite) SetupTest() {
	s.container = s.newPostgresInstance()

	port, err := s.container.MappedPort(context.Background(), "5432")
	s.Require().NoError(err)

	s.conn = MustConnect(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres", "postgres", "127.0.0.1", port.Int(), "insurance"),
		"../migrations")
	s.repo = NewRepository(s.conn)

	s.ctx = logger.WrapLogger(context.Background(), logger.NewLogger(logger.Zap, "local"))
}

func (s *PostgresRepoSuite) TeardownTest() {
	s.Require().NoError(s.conn.Close())
	s.Require().NoError(s.container.Terminate(context.Background()))
}

func (s *PostgresRepoSuite) TestGetInsuranceData() {
	data := models.InsuranceData{
		ID:         "x001",
		Status:     models.Pending,
		ActiveTill: time.Time{}.UTC(),
		Price:      100,
	}

	q := `INSERT INTO insurances (id,status, active_till, price) VALUES($1, $2, $3, $4)`
	_, err := s.conn.Exec(q, data.ID, data.Status, data.ActiveTill, data.Price)
	s.Require().NoError(err)

	res, err := s.repo.GetInsuranceData(s.ctx, data.ID)
	s.Require().NoError(err)
	s.Require().Equal(data, res)
}

func (s *PostgresRepoSuite) TestCreateInsuranceData() {
	data := models.InsuranceData{
		ID:         "x001",
		Status:     models.Pending,
		ActiveTill: time.Time{}.UTC(),
		Price:      100,
	}

	err := s.repo.CreateInsuranceData(s.ctx, data)
	s.Require().NoError(err)

	q := `SELECT * FROM insurances WHERE id = $1`
	row := s.conn.QueryRowx(q, data.ID)
	s.Require().NoError(row.Err())

	var res models.InsuranceData
	s.Require().NoError(row.StructScan(&res))

	res.ActiveTill = res.ActiveTill.UTC()

	s.Require().Equal(data, res)
}

func (s *PostgresRepoSuite) TestUpdateInsuranceStatus() {
	data := models.InsuranceData{
		ID:         "x001",
		Status:     models.Pending,
		ActiveTill: time.Now(),
		Price:      100,
	}

	q := `INSERT INTO insurances (id,status, active_till, price) VALUES($1, $2, $3, $4)`
	_, err := s.conn.Exec(q, data.ID, data.Status, data.ActiveTill, data.Price)
	s.Require().NoError(err)

	err = s.repo.UpdateInsuranceStatus(s.ctx, data.ID, models.Active)
	s.Require().NoError(err)

	q = `SELECT status FROM insurances WHERE id = $1`
	row := s.conn.QueryRowx(q, data.ID)
	s.Require().NoError(row.Err())

	var res int
	s.Require().NoError(row.Scan(&res))

	s.Require().Equal(models.Active, res)
}

func (s *PostgresRepoSuite) newPostgresInstance() testcontainers.Container {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres",
			ExposedPorts: []string{"5432/tcp"},
			WaitingFor: wait.ForAll(
				wait.ForListeningPort("5432"),
				wait.ForLog("database system is ready to accept connections"),
			),
			Env: map[string]string{
				"POSTGRES_USER":     "postgres",
				"POSTGRES_PASSWORD": "postgres",
				"POSTGRES_DB":       "insurance",
			},
		},
		Started: true,
	})

	if err != nil {
		panic("failed to start container: " + err.Error())

	}

	return container
}
