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
	Create(ctx context.Context, key string, val any) error
	Delete(ctx context.Context, key string, val any) error

	ProcessNotCommittedInsurances(ctx context.Context)
	ProcessNotCommittedPayoffs(ctx context.Context)
}
