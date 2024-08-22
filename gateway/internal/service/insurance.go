package service

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/clients"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
)

type insurance struct {
	insuranceClient clients.InsuranceClient
}

func (i insurance) CreateInsurance(ctx context.Context, insurance models.Insurance) error {
	//TODO implement me
	panic("implement me")
}

func (i insurance) GetInsuranceData(ctx context.Context, addr string) (models.InsuranceData, error) {
	//TODO implement me
	panic("implement me")
}
