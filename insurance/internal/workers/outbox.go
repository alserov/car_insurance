package workers

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/service"
	"time"
)

const (
	outboxTickPeriod = time.Second * 60 * 30
)

type outbox struct {
	contractCl clients.ContractClient

	srvc service.Service

	log logger.Logger
}

func NewOutboxWorker(srvc service.Service, cls service.Clients, log logger.Logger) *outbox {
	return &outbox{
		log:        log,
		srvc:       srvc,
		contractCl: cls.ContractClient,
	}
}

func (o outbox) Start(ctx context.Context) {
	go o.processPendingPayoffItems(ctx)
	go o.processPendingInsuranceItems(ctx)
}

// processPendingInsuranceItems
// Gets all pending insurance items and produces them to broker
func (o outbox) processPendingInsuranceItems(ctx context.Context) {
	tick := time.NewTicker(outboxTickPeriod)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			func() {
				jobCtx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				jobCtx = logger.WrapLogger(jobCtx, o.log)

				if err := o.srvc.ProducePendingInsuranceItems(jobCtx); err != nil {
					o.log.Error("failed to produce new insurances", logger.WithArg("error", err.Error()))
				}
			}()
		case <-ctx.Done():
			return
		}
	}
}

// processPendingPayoffItems
// Gets all pending payoff items and produces them to broker
func (o outbox) processPendingPayoffItems(ctx context.Context) {
	tick := time.NewTicker(outboxTickPeriod)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			func() {
				jobCtx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				jobCtx = logger.WrapLogger(jobCtx, o.log)

				if err := o.srvc.ProducePendingInsuranceItems(jobCtx); err != nil {
					o.log.Error("failed to produce new insurances", logger.WithArg("error", err.Error()))
				}
			}()
		case <-ctx.Done():
			return
		}
	}
}
