package redis

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/db"
	"time"
)

func NewOutbox(rd *rd.Client, tickPeriod time.Duration) db.Outbox {
	return &redis{
		tickPeriod: tickPeriod,
	}
}

type redis struct {
	tickPeriod time.Duration
}

const (
	Initialized = iota
	Succeeded
	Failed
)

func (r redis) Create(ctx context.Context, key string, val any) error {
	//TODO implement me
	panic("implement me")
}

func (r redis) Delete(ctx context.Context, key string, val any) error {
	//TODO implement me
	panic("implement me")
}

func (r redis) ProcessNotCommittedInsurances(ctx context.Context) {
	ticker := time.NewTicker(r.tickPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
		case <-ticker.C:

		}
	}
}

func (r redis) ProcessNotCommittedPayoffs(ctx context.Context) {
	ticker := time.NewTicker(r.tickPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
		case <-ticker.C:

		}
	}
}
