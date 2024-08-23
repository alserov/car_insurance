package workers

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"time"
)

type Outbox interface {
	Start(ctx context.Context)
}

const (
	outboxTickPeriod = time.Second * 60 * 30
)

type outbox struct {
	repo       db.Outbox
	contractCl clients.ContractClient

	log logger.Logger
}

func NewOutboxWorker(repo db.Outbox, contractCl clients.ContractClient, log logger.Logger) *outbox {
	return &outbox{
		repo:       repo,
		contractCl: contractCl,
		log:        log,
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

				items, err := o.repo.Get(jobCtx, models.Pending, models.GroupInsurance)
				if err != nil {
					o.log.Error("failed to get pending items from outbox", logger.WithArg("error", err.Error()))
					return
				}

				for _, val := range items {
					item, ok := val.Val.(models.Insurance)
					if !ok {
						continue
					}

					if err = o.contractCl.CreateInsurance(jobCtx, item); err != nil {
						o.log.Error("client failed to create insurance", logger.WithArg("error", err.Error()))
						continue
					}
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

				items, err := o.repo.Get(jobCtx, models.Pending, models.GroupPayoff)
				if err != nil {
					o.log.Error("failed to get pending items from outbox", logger.WithArg("error", err.Error()))
					return
				}

				for _, val := range items {
					item, ok := val.Val.(models.Payoff)
					if !ok {
						continue
					}

					if err = o.contractCl.Payoff(jobCtx, item); err != nil {
						o.log.Error("client failed to payoff", logger.WithArg("error", err.Error()))
						continue
					}
				}
			}()
		case <-ctx.Done():
			return
		}
	}
}
