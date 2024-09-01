package workers

import (
	"context"
	"github.com/alserov/car_insurance/contract/internal/clients"
	"github.com/alserov/car_insurance/contract/internal/logger"
	"github.com/alserov/car_insurance/contract/internal/service"
	"time"
)

func NewInsuranceWorker(srvc service.Service, cls service.Clients, log logger.Logger) Worker {
	return &insurance{
		srvc:            srvc,
		insuranceClient: cls.InsuranceClient,
		log:             log,
	}
}

type insurance struct {
	srvc service.Service

	insuranceClient clients.InsuranceClient

	log logger.Logger
}

func (i insurance) Start(ctx context.Context) {
	go i.processPayoffs(ctx)
	go i.processNewInsurances(ctx)
}

func (i insurance) processNewInsurances(ctx context.Context) {
	for {
		select {
		case ins := <-i.insuranceClient.GetNewInsurances(ctx):
			func() {
				jobCtx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				if err := i.srvc.CreateInsurance(jobCtx, ins); err != nil {
					i.log.Error("failed to create new insurance", logger.WithArg("error", err.Error()))
				}
			}()
		case <-ctx.Done():
			return
		}
	}
}

func (i insurance) processPayoffs(ctx context.Context) {
	for {
		select {
		case pay := <-i.insuranceClient.GetPayoffs(ctx):
			func() {
				jobCtx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				if err := i.srvc.Payoff(jobCtx, pay); err != nil {
					i.log.Error("failed to payoff", logger.WithArg("error", err.Error()))
				}
			}()
		case <-ctx.Done():
			return
		}
	}
}
