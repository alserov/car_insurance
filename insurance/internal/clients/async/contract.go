package async

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/insurance/internal/async"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
)

func NewContractClient(p async.Producer, c async.Consumer[models.ContractCommit]) clients.ContractClient {
	return &contract{
		p: p,
		c: c,
	}
}

type contract struct {
	p async.Producer
	c async.Consumer[models.ContractCommit]
}

func (c contract) GetCommits(ctx context.Context) chan models.ContractCommit {
	return c.c.Consume(ctx)
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
