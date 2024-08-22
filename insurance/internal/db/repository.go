package db

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
)

type Repository interface {
	GetInsuranceData(ownerAddr string) (models.InsuranceData, error)
	CreateInsuranceData(insData models.InsuranceData) error
	UpdateInsuranceStatus()
}

type Outbox interface {
	Create(ctx context.Context, item models.OutboxItem) error
	Get(ctx context.Context, status int, groupID int) ([]models.OutboxItem, error)
	Delete(ctx context.Context, status int, groupID int) error
}
