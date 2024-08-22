package service

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/clients"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
)

type payments struct {
	insuranceClient clients.InsuranceClient
}

func (p payments) Payoff(ctx context.Context, payoff models.Payoff) error {
	//TODO implement me
	panic("implement me")
}
