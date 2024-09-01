package clients

import (
	"context"
	"github.com/alserov/car_insurance/contract/internal/service/models"
)

type InsuranceClient interface {
	GetNewInsurances(ctx context.Context) <-chan models.NewInsurance
	GetPayoffs(ctx context.Context) <-chan models.Payoff
	Commit(ctx context.Context, commit models.OutboxItem) error
}

type ContractClient interface {
	Payoff(ctx context.Context, pay models.Payoff) error
	Insure(ctx context.Context, ins models.NewInsurance) error
}
