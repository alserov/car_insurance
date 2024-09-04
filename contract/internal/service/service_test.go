package service

import (
	"context"
	"github.com/alserov/car_insurance/contract/internal/logger"
	"github.com/alserov/car_insurance/contract/internal/mocks"
	"github.com/alserov/car_insurance/contract/internal/service/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type ServiceSuite struct {
	suite.Suite

	ctrl        *gomock.Controller
	insuranceCl *mocks.MockInsuranceClient
	contractCl  *mocks.MockContractClient
	outbox      *mocks.MockOutbox

	ctx context.Context

	srvc Service
}

func (s *ServiceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.insuranceCl = mocks.NewMockInsuranceClient(s.ctrl)
	s.contractCl = mocks.NewMockContractClient(s.ctrl)
	s.outbox = mocks.NewMockOutbox(s.ctrl)

	s.srvc = NewService(Clients{
		InsuranceClient: s.insuranceCl,
		ContractClient:  s.contractCl,
	}, s.outbox)

	s.ctx = logger.WrapLogger(context.Background(), logger.NewLogger(logger.Zap, "local"))
}

func (s *ServiceSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *ServiceSuite) TestGetNewInsuranceCommits() {
	data := []models.OutboxItem{
		{
			ID:      "some id",
			GroupID: models.GroupNewInsurances,
			Status:  models.Pending,
			Val: models.NewInsurance{
				Sender:     "x001",
				Amount:     100,
				ActiveTill: time.Now(),
			},
		},
		{
			ID:      "some id 1",
			GroupID: models.GroupNewInsurances,
			Status:  models.Pending,
			Val: models.NewInsurance{
				Sender:     "x002",
				Amount:     100,
				ActiveTill: time.Now(),
			},
		},
	}

	s.outbox.EXPECT().
		Get(gomock.Any(), gomock.Eq(models.Pending), gomock.Eq(models.GroupNewInsurances)).
		Times(1).
		Return(data, nil)

	res, err := s.srvc.GetNewInsuranceCommits(s.ctx)
	s.Require().NoError(err)
	s.Require().Equal(data, res)
}

func (s *ServiceSuite) TestDeleteCommits() {
	data := []string{
		"id 1", "id2",
	}

	for _, id := range data {
		s.outbox.EXPECT().
			Delete(gomock.Any(), gomock.Eq(id)).
			Times(1).
			Return(nil)
	}

	err := s.srvc.DeleteCommits(s.ctx, data)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestCreateInsurance() {
	data := models.NewInsurance{
		Sender:     "x001",
		Amount:     100,
		ActiveTill: time.Now(),
	}

	s.contractCl.EXPECT().
		Insure(gomock.Any(), gomock.Eq(data)).
		Times(1).
		Return(nil)

	s.outbox.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	err := s.srvc.CreateInsurance(s.ctx, data)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestPayoff() {
	data := models.Payoff{
		Receiver: "x001",
		Mult:     1.55,
	}

	s.contractCl.EXPECT().
		Payoff(gomock.Any(), gomock.Eq(data)).
		Times(1).
		Return(nil)

	err := s.srvc.Payoff(s.ctx, data)
	s.Require().NoError(err)
}
