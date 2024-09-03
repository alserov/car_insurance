package workers

import (
	"context"
	"github.com/alserov/car_insurance/contract/internal/clients"
	"github.com/alserov/car_insurance/contract/internal/logger"
	"github.com/alserov/car_insurance/contract/internal/service"
	"time"
)

func NewOutboxWorker(srvc service.Service, cls service.Clients, log logger.Logger) Worker {
	return &outbox{
		insuranceClient: cls.InsuranceClient,
		srvc:            srvc,
		log:             log,
	}
}

type outbox struct {
	insuranceClient clients.InsuranceClient

	srvc service.Service

	log logger.Logger
}

func (o outbox) Start(ctx context.Context) {
	go o.processPendingCommits(ctx)
}

func (o outbox) processPendingCommits(ctx context.Context) {
	tick := time.NewTicker(time.Minute * 30)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			func() {
				jobCtx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				items, err := o.srvc.GetNewInsuranceCommits(jobCtx)
				if err != nil {
					o.log.Error("failed to get pending items from outbox", logger.WithArg("error", err.Error()))
					return
				}

				var deleteCommitIDs []string

				for _, item := range items {
					if err = o.insuranceClient.Commit(jobCtx, item); err != nil {
						o.log.Error("client failed to payoff", logger.WithArg("error", err.Error()))
						continue
					}
					deleteCommitIDs = append(deleteCommitIDs, item.ID)
				}

				if err = o.srvc.DeleteCommits(jobCtx, deleteCommitIDs); err != nil {
					o.log.Error("failed to delete commits", logger.WithArg("error", err.Error()))
				}
			}()
		case <-ctx.Done():
			return
		}
	}
}
