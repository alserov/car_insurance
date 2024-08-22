package service

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/clients"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
)

type insurance struct {
	insuranceClient clients.InsuranceClient
}

func (i insurance) CreateInsurance(ctx context.Context, insurance models.Insurance) error {
	err := i.insuranceClient.CreateInsurance(ctx, insurance)
	if err != nil {
		return fmt.Errorf("insurance client error: %w", err)
	}

	return nil
}

func (i insurance) GetInsuranceData(ctx context.Context, addr string) (models.InsuranceData, error) {
	data, err := i.insuranceClient.GetInsuranceData(ctx, addr)
	if err != nil {
		return models.InsuranceData{}, fmt.Errorf("insurance client error: %w", err)
	}

	return data, nil
}

func (i insurance) Payoff(ctx context.Context, payoff models.Payoff) error {
	err := i.insuranceClient.Payoff(ctx, payoff)
	if err != nil {
		return fmt.Errorf("insurance client error: %w", err)
	}

	return nil
}
