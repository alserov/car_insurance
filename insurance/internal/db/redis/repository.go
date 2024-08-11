package redis

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/db"
)

func NewRepository() db.Repository {
	return &redis{}
}

type redis struct {
}

func (r redis) Create(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}

func (r redis) Delete(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}
