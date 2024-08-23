package db

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
)

type Repository interface {
	GetInsuranceData(ctx context.Context, ownerAddr string) (models.InsuranceData, error)
	CreateInsuranceData(ctx context.Context, insData models.InsuranceData) error
	UpdateInsuranceStatus(ctx context.Context, id string, status uint) error
}

type Outbox interface {
	Create(ctx context.Context, item models.OutboxItem) error
	Get(ctx context.Context, status int, groupID int) ([]models.OutboxItem, error)
	Delete(ctx context.Context, id string, groupID int) error
}
