package fiber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/mocks"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFiberHandlersSuite(t *testing.T) {
	suite.Run(t, new(fiberHandlersSuite))
}

type fiberHandlersSuite struct {
	suite.Suite

	ctrl        *gomock.Controller
	insuranceCl *mocks.MockInsuranceClient

	app *fiber.App

	hndlr handler
}

func (s *fiberHandlersSuite) SetupTest() {
	s.app = fiber.New()

	s.ctrl = gomock.NewController(s.T())
	s.insuranceCl = mocks.NewMockInsuranceClient(s.ctrl)

	s.hndlr = handler{
		insurance: insurance{
			service: service.NewService(service.Clients{
				InsuranceClient: s.insuranceCl,
			}).Insurance,
		},
	}

	s.app.Post("/new", s.hndlr.insurance.CreateInsurance)
	s.app.Get("/info", s.hndlr.insurance.GetInsuranceData)
	s.app.Post("/payoff", s.hndlr.insurance.Payoff)
}

func (s *fiberHandlersSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *fiberHandlersSuite) TestCreateInsurance() {
	data := models.Insurance{
		SenderAddr: "x001",
		Amount:     1000,
		CarImage:   []byte("image"),
	}

	s.insuranceCl.EXPECT().
		CreateInsurance(gomock.Any(), gomock.Eq(data)).
		Times(1).
		Return(nil)

	b, err := json.Marshal(data)
	s.Require().NoError(err)

	req := httptest.NewRequest(http.MethodPost, "/new", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	res, err := s.app.Test(req, -1)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, res.StatusCode)
}

func (s *fiberHandlersSuite) TestGetInsuranceData() {
	data := models.InsuranceData{
		Status:             1,
		ActiveTill:         time.Now().String(),
		Owner:              "x001",
		Price:              100,
		MaxInsurancePayoff: 200,
		MinInsurancePayoff: 150,
		AvgInsurancePayoff: 175,
	}

	s.insuranceCl.EXPECT().
		GetInsuranceData(gomock.Any(), gomock.Eq(data.Owner)).
		Times(1).
		Return(data, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/info?addr=%s", data.Owner), nil)
	req.Header.Set("Content-Type", "application/json")

	res, err := s.app.Test(req, -1)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, res.StatusCode)

	b, err := io.ReadAll(res.Body)
	s.Require().NoError(err)

	var resBody models.InsuranceData
	s.Require().NoError(json.Unmarshal(b, &resBody))

	s.Require().Equal(data, resBody)
}

func (s *fiberHandlersSuite) TestPayoff() {
	data := models.Payoff{
		ReceiverAddr: "x001",
		CarImage:     []byte("image"),
	}

	s.insuranceCl.EXPECT().
		Payoff(gomock.Any(), gomock.Eq(data)).
		Times(1).
		Return(nil)

	b, err := json.Marshal(data)
	s.Require().NoError(err)

	req := httptest.NewRequest(http.MethodPost, "/payoff", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	res, err := s.app.Test(req, -1)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, res.StatusCode)
}
