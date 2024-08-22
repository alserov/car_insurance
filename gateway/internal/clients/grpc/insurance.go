package grpc

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/clients"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
)

func NewInsuranceClient() clients.InsuranceClient {
	return &insurance{}
}

type insurance struct {
}

func (i insurance) CreateInsurance(ctx context.Context, insurance models.Insurance) error {
	//TODO implement me
	panic("implement me")
}

func (i insurance) Payoff(ctx context.Context, payoff models.Payoff) error {
	//TODO implement me
	panic("implement me")
}

func (i insurance) GetInsuranceData(ctx context.Context) (models.InsuranceData, error) {
	//TODO implement me
	panic("implement me")
}
