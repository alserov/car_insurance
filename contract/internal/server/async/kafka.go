package async

import (
	"context"
	"github.com/alserov/car_insurance/contract/internal/async"
	"github.com/alserov/car_insurance/contract/internal/logger"
	"github.com/alserov/car_insurance/contract/internal/service"
	"github.com/alserov/car_insurance/contract/internal/service/models"
	"time"
)

type server struct {
	service service.Service
}

type Consumers struct {
	NewInsuranceCons async.Consumer[models.NewInsurance]
	PayoffCons       async.Consumer[models.Payoff]
}

func StartServer(ctx context.Context, c Consumers, srvc service.Service) {
	s := server{
		service: srvc,
	}

	s.consumeNewInsurances(ctx, c.NewInsuranceCons)
	s.consumePayoffs(ctx, c.PayoffCons)

	<-ctx.Done()
}

func (s server) consumeNewInsurances(ctx context.Context, cons async.Consumer[models.NewInsurance]) {
	bgCtx := context.Background()
	log := logger.ExtractLogger(ctx)

	for ins := range cons.Consume(ctx) {
		ctx, cancel := context.WithTimeout(bgCtx, time.Second*5)

		if err := s.service.CreateInsurance(ctx, ins); err != nil {
			log.Error("failed to create insurance", logger.WithArg("error", err.Error()))
		}

		cancel()
	}
}

func (s server) consumePayoffs(ctx context.Context, cons async.Consumer[models.Payoff]) {
	bgCtx := context.Background()
	log := logger.ExtractLogger(ctx)

	for pay := range cons.Consume(ctx) {
		ctx, cancel := context.WithTimeout(bgCtx, time.Second*5)

		if err := s.service.Payoff(ctx, pay); err != nil {
			log.Error("failed to payoff", logger.WithArg("error", err.Error()))
		}

		cancel()
	}
}
