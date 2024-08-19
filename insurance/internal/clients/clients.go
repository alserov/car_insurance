package clients

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
)

type RecognitionClient interface {
	CheckIfCarIsOK(ctx context.Context, image []byte) error
	CalcDamageMultiplier(ctx context.Context, image []byte) (float32, error)
}

type ContractClient interface {
	CreateInsurance(ctx context.Context, ins models.Insurance) error
	Payoff(ctx context.Context, payoff models.Payoff) error
}
