package db

import (
	"context"
	"github.com/alserov/car_insurance/contract/internal/service/models"
)

type Outbox interface {
	Create(ctx context.Context, item models.OutboxItem) error
	Get(ctx context.Context, status int, groupID int) ([]models.OutboxItem, error)
	Delete(ctx context.Context, id string) error
}
