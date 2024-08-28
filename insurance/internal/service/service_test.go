package service

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/mocks"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(serviceSuite))
}

type serviceSuite struct {
	suite.Suite

	ctrl          *gomock.Controller
	contractCl    *mocks.MockContractClient
	recognitionCl *mocks.MockRecognitionClient
	repo          *mocks.MockRepository
	outbox        *mocks.MockOutbox

	ctx context.Context

	srvc Service
}

func (s *serviceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.recognitionCl = mocks.NewMockRecognitionClient(s.ctrl)
	s.contractCl = mocks.NewMockContractClient(s.ctrl)
	s.repo = mocks.NewMockRepository(s.ctrl)
	s.outbox = mocks.NewMockOutbox(s.ctrl)
	s.ctx = logger.WrapLogger(context.Background(), logger.NewLogger(logger.Zap, "local"))

	s.srvc = NewService(Clients{
		Recognition:    s.recognitionCl,
		ContractClient: s.contractCl,
	}, s.outbox, s.repo)
}

func (s *serviceSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *serviceSuite) TestProducePendingInsuranceItems() {
	outboxItems := []models.OutboxItem{
		{
			ID:      "uuid",
			GroupID: models.GroupInsurance,
			Status:  models.Pending,
			Val: models.Insurance{
				ID:         "uuid",
				SenderAddr: "x001",
				Amount:     100,
				CarImage:   nil,
				ActiveTill: time.Now().Add(time.Hour),
			},
		},
	}

	s.outbox.EXPECT().
		Get(gomock.Any(), gomock.Eq(models.Pending), gomock.Eq(models.GroupInsurance)).
		Times(1).
		Return(outboxItems, nil)

	s.contractCl.EXPECT().
		CreateInsurance(gomock.Any(), gomock.Eq(outboxItems[0].Val)).
		Times(1).
		Return(nil)

	s.Require().NoError(s.srvc.ProducePendingInsuranceItems(s.ctx))
}

func (s *serviceSuite) TestProducePendingPayoffItems() {
	outboxItems := []models.OutboxItem{
		{
			ID:      "uuid",
			GroupID: models.GroupInsurance,
			Status:  models.Pending,
			Val: models.Payoff{
				CarImage:     nil,
				ReceiverAddr: "x001",
				Multiplier:   1.7,
			},
		},
	}

	s.outbox.EXPECT().
		Get(gomock.Any(), gomock.Eq(models.Pending), gomock.Eq(models.GroupPayoff)).
		Times(1).
		Return(outboxItems, nil)

	s.contractCl.EXPECT().
		Payoff(gomock.Any(), gomock.Eq(outboxItems[0].Val)).
		Times(1).
		Return(nil)

	s.Require().NoError(s.srvc.ProducePendingPayoffItems(s.ctx))
}

func (s *serviceSuite) TestUpdateInsuranceStatus() {
	s.outbox.EXPECT().
		Delete(gomock.Any(), gomock.Eq("x001")).
		Times(1).
		Return(nil)

	s.repo.EXPECT().
		UpdateInsuranceStatus(gomock.Any(), gomock.Eq("x001"), uint(models.Active)).
		Times(1).
		Return(nil)

	s.Require().NoError(s.srvc.ActivateInsurance(s.ctx, "x001"))
}

func (s *serviceSuite) TestGetInsuranceData() {
	data := models.InsuranceData{
		ID:                 "x001",
		Status:             models.Active,
		ActiveTill:         time.Now().Add(time.Hour),
		Price:              100,
		MaxInsurancePayoff: 200,
		MinInsurancePayoff: 101,
		AvgInsurancePayoff: 150,
	}

	s.repo.EXPECT().
		GetInsuranceData(gomock.Any(), gomock.Eq(data.ID)).
		Times(1).
		Return(data, nil)

	res, err := s.srvc.GetInsuranceData(s.ctx, data.ID)
	s.Require().NoError(err)
	s.Require().Equal(data, res)
}

func (s *serviceSuite) TestCreateInsurance() {
}

func (s *serviceSuite) TestPayoff() {
}
