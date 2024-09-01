package service

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/clients"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
)

type Clients struct {
	InsuranceClient clients.InsuranceClient
}

func NewService(c Clients) *Service {
	return &Service{
		Insurance: insurance{
			insuranceClient: c.InsuranceClient,
		},
	}
}

type Insurance interface {
	CreateInsurance(ctx context.Context, insurance models.Insurance) error
	GetInsuranceData(ctx context.Context, addr string) (models.InsuranceData, error)
	Payoff(ctx context.Context, payoff models.Payoff) error
}

type Service struct {
	Insurance Insurance
}
