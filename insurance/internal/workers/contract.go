package workers

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"time"
)

type Contract interface {
	Start(ctx context.Context)
}

type contract struct {
	contractClient clients.ContractClient

	outboxRepo db.Outbox
	repo       db.Repository

	log logger.Logger
}

func NewContractWorker(outboxRepo db.Outbox, repo db.Repository, contractClient clients.ContractClient, log logger.Logger) *contract {
	return &contract{
		outboxRepo:     outboxRepo,
		repo:           repo,
		log:            log,
		contractClient: contractClient,
	}
}

func (c contract) Start(ctx context.Context) {
	go c.processContractCommits(ctx)
}

// processContractCommits
// Gets all pending insurance items and produces them to broker
func (c contract) processContractCommits(ctx context.Context) {
	for {
		select {
		case msg := <-c.contractClient.GetCommits(ctx):
			func() {
				jobCtx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				jobCtx = logger.WrapLogger(jobCtx, c.log)

				switch msg.Group {
				case models.GroupInsurance:
					if err := c.outboxRepo.Delete(jobCtx, msg.Addr, models.GroupInsurance); err != nil {
						c.log.Error("failed to delete from outbox", logger.WithArg("error", err.Error()))
						return
					}

					if err := c.repo.UpdateInsuranceStatus(jobCtx, msg.Addr, models.Active); err != nil {
						c.log.Error("failed to update insurance status", logger.WithArg("error", err.Error()))
						return
					}
				case models.GroupPayoff:
					if err := c.outboxRepo.Delete(jobCtx, msg.Addr, models.GroupPayoff); err != nil {
						c.log.Error("failed to delete from outbox", logger.WithArg("error", err.Error()))
						return
					}
				}
			}()
		case <-ctx.Done():
			return
		}
	}
}
