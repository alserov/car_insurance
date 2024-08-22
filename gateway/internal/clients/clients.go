package clients

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
)

type InsuranceClient interface {
	CreateInsurance(ctx context.Context, insurance models.Insurance) error
	Payoff(ctx context.Context, payoff models.Payoff) error
	GetInsuranceData(ctx context.Context) (models.InsuranceData, error)
}
