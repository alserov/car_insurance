package grpc

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/mocks"
	"github.com/alserov/car_insurance/insurance/internal/service"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	proto "github.com/alserov/car_insurance/insurance/pkg/grpc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestGRPCServerSuite(t *testing.T) {
	suite.Run(t, new(GRPCServerSuite))
}

type GRPCServerSuite struct {
	suite.Suite

	srvr proto.InsuranceServer

	ctrl          *gomock.Controller
	repo          *mocks.MockRepository
	outbox        *mocks.MockOutbox
	recognitionCl *mocks.MockRecognitionClient
	contractCl    *mocks.MockContractClient
}

func (s *GRPCServerSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.repo = mocks.NewMockRepository(s.ctrl)
	s.outbox = mocks.NewMockOutbox(s.ctrl)
	s.recognitionCl = mocks.NewMockRecognitionClient(s.ctrl)
	s.contractCl = mocks.NewMockContractClient(s.ctrl)

	s.srvr = grpcServer{
		srvc: service.NewService(service.Clients{
			Recognition:    s.recognitionCl,
			ContractClient: s.contractCl,
		}, s.outbox, s.repo),
	}
}

func (s *GRPCServerSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *GRPCServerSuite) TestGetInsuranceData() {
	data := models.InsuranceData{
		ID:                 "x001",
		Status:             models.Pending,
		ActiveTill:         time.Now(),
		Price:              100,
		MaxInsurancePayoff: 199,
		MinInsurancePayoff: 150,
		AvgInsurancePayoff: 175,
	}

	s.repo.EXPECT().
		GetInsuranceData(gomock.Any(), gomock.Eq(data.ID)).
		Times(1).
		Return(data, nil)

	res, err := s.srvr.GetInsuranceData(context.Background(), &proto.InsuranceOwner{Addr: data.ID})
	s.Require().NoError(err)
	s.Require().Equal(utils.Converter{}.FromInsuranceData(data), res)
}

func (s *GRPCServerSuite) TestCreateInsurance() {
	data := models.Insurance{
		SenderAddr: "x001",
		Amount:     100,
		CarImage:   []byte("img"),
		ActiveTill: time.Now(),
	}

	s.recognitionCl.EXPECT().
		CheckIfCarIsOK(gomock.Any(), gomock.Eq(data.CarImage)).
		Times(1).
		Return(nil)

	s.outbox.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	s.repo.EXPECT().
		CreateInsuranceData(gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	_, err := s.srvr.CreateInsurance(context.Background(), &proto.NewInsurance{
		SenderAddr: data.SenderAddr,
		Amount:     data.Amount,
		CarImage:   data.CarImage,
	})
	s.Require().NoError(err)
}

func (s *GRPCServerSuite) TestPayoff() {
	data := models.Payoff{
		CarImage:     []byte("img"),
		ReceiverAddr: "x001",
		Multiplier:   1.5,
	}

	s.recognitionCl.EXPECT().
		CalcDamageMultiplier(gomock.Any(), gomock.Eq(data.CarImage)).
		Times(1).
		Return(data.Multiplier, nil)

	s.outbox.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	_, err := s.srvr.Payoff(context.Background(), &proto.NewPayoff{
		ReceiverAddr: data.ReceiverAddr,
		CarImage:     data.CarImage,
	})
	s.Require().NoError(err)
}
