package service

import (
	"context"
	"github.com/alserov/car_insurance/contract/internal/service/models"
)

type Service interface {
	CreateInsurance(ctx context.Context, ins models.NewInsurance) error
	Payoff(ctx context.Context, ins models.Payoff) error
}

func NewService() Service {
	return &service{}
}

type service struct {
}

func (s service) CreateInsurance(ctx context.Context, ins models.NewInsurance) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Payoff(ctx context.Context, ins models.Payoff) error {
	//TODO implement me
	panic("implement me")
}
