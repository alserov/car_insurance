package mux

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/mocks"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
	"github.com/alserov/car_insurance/gateway/internal/tracing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMuxHandlersSuite(t *testing.T) {
	suite.Run(t, new(MuxHandlersSuite))
}

type MuxHandlersSuite struct {
	suite.Suite

	srvc    *service.Service
	handler *handler

	ctrl        *gomock.Controller
	insuranceCl *mocks.MockInsuranceClient

	ctx context.Context
}

func (s *MuxHandlersSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.insuranceCl = mocks.NewMockInsuranceClient(s.ctrl)

	s.srvc = service.NewService(service.Clients{InsuranceClient: s.insuranceCl})

	s.handler = &handler{
		insurance: insurance{
			service: s.srvc.Insurance,
		},
	}

	exp, err := stdouttrace.New()
	s.Require().NoError(err)

	tracer, _ := tracing.NewTracer(exp, "test")

	s.ctx = context.Background()
	s.ctx = logger.WrapLogger(s.ctx, logger.NewLogger(logger.Zap, "local"))
	s.ctx = tracing.WrapTracer(s.ctx, tracer)
}

func (s *MuxHandlersSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *MuxHandlersSuite) TestCreateInsurance() {
	data := models.Insurance{
		SenderAddr: "x001",
		Amount:     1000,
		CarImage:   []byte("img"),
	}

	s.insuranceCl.EXPECT().
		CreateInsurance(gomock.Any(), gomock.Eq(data)).
		Times(1).
		Return(nil)

	b, err := json.Marshal(data)
	s.Require().NoError(err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b)).WithContext(s.ctx)

	s.Require().NoError(s.handler.insurance.CreateInsurance(w, r))
	s.Require().Equal(http.StatusCreated, w.Code)
}

func (s *MuxHandlersSuite) TestGetInsuranceData() {
	data := models.InsuranceData{
		Status:             1,
		ActiveTill:         time.Date(2024, 3, 3, 3, 3, 3, 3, time.UTC).String(),
		Owner:              "x001",
		Price:              100,
		MaxInsurancePayoff: 199,
		MinInsurancePayoff: 101,
		AvgInsurancePayoff: 150,
	}

	s.insuranceCl.EXPECT().
		GetInsuranceData(gomock.Any(), gomock.Eq(data.Owner)).
		Times(1).
		Return(data, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?addr=%s", data.Owner), nil).WithContext(s.ctx)

	s.Require().NoError(s.handler.insurance.GetInsuranceData(w, r))
	s.Require().Equal(http.StatusOK, w.Code)

	b, err := io.ReadAll(w.Body)
	s.Require().NoError(err)

	var body models.InsuranceData
	s.Require().NoError(json.Unmarshal(b, &body))

	s.Require().Equal(data, body)
}

func (s *MuxHandlersSuite) TestPayoff() {
	data := models.Payoff{
		ReceiverAddr: "x001",
		CarImage:     []byte("img"),
	}

	s.insuranceCl.EXPECT().
		Payoff(gomock.Any(), gomock.Eq(data)).
		Times(1).
		Return(nil)

	b, err := json.Marshal(data)
	s.Require().NoError(err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b)).WithContext(s.ctx)

	s.Require().NoError(s.handler.insurance.Payoff(w, r))
	s.Require().Equal(http.StatusCreated, w.Code)
}
