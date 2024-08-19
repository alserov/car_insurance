package async

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/insurance/internal/async"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
)

func NewContractClient(p async.Producer) clients.ContractClient {
	return &contract{
		p: p,
	}
}

type contract struct {
	p async.Producer
}

func (c contract) CreateInsurance(ctx context.Context, ins models.Insurance) error {
	if err := c.p.Produce(ctx, ins); err != nil {
		return fmt.Errorf("failed to produce create insurance message: %w", err)
	}

	return nil
}

func (c contract) Payoff(ctx context.Context, payoff models.Payoff) error {
	if err := c.p.Produce(ctx, payoff); err != nil {
		return fmt.Errorf("failed to produce payoff message: %w", err)
	}

	return nil
}
