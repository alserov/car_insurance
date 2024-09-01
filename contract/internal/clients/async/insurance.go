package async

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/contract/internal/async"
	"github.com/alserov/car_insurance/contract/internal/clients"
	"github.com/alserov/car_insurance/contract/internal/service/models"
)

func NewInsuranceClient(commitProd async.Producer, newInsCons async.Consumer[models.NewInsurance], payoffCons async.Consumer[models.Payoff]) clients.InsuranceClient {
	return &insurance{
		p:       commitProd,
		insCons: newInsCons,
		payCons: payoffCons,
	}
}

type insurance struct {
	p async.Producer

	insCons async.Consumer[models.NewInsurance]
	payCons async.Consumer[models.Payoff]
}

func (i insurance) Commit(ctx context.Context, comm models.OutboxItem) error {
	if err := i.p.Produce(ctx, comm); err != nil {
		return fmt.Errorf("failed to produce commit: %w", err)
	}

	return nil
}

func (i insurance) GetNewInsurances(ctx context.Context) <-chan models.NewInsurance {
	return i.insCons.Consume(ctx)
}

func (i insurance) GetPayoffs(ctx context.Context) <-chan models.Payoff {
	return i.payCons.Consume(ctx)
}
