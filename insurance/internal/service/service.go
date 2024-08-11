package service

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
)

type Service interface {
	CreateInsurance(ctx context.Context, insData models.Insurance, carImage []byte) error
	Payoff(ctx context.Context, receiverAddr string, carImage []byte) error
}

type Clients struct {
	Recognition clients.RecognitionClient
	Contract    clients.ContractClient
}

func NewService(cls Clients) Service {
	return &service{
		recognitionCl: cls.Recognition,
		contractCl:    cls.Contract,
	}
}

type service struct {
	recognitionCl clients.RecognitionClient
	contractCl    clients.ContractClient
}

func (s service) CreateInsurance(ctx context.Context, insData models.Insurance, carImage []byte) error {
	if err := s.recognitionCl.CheckIfCarIsOK(ctx, carImage); err != nil {
		return fmt.Errorf("failed to create insurance: %w", err)
	}

	if err := s.contractCl.CreateInsurance(ctx, insData); err != nil {
		return fmt.Errorf("failed to create insurance: %w", err)
	}

	return nil
}

func (s service) Payoff(ctx context.Context, receiverAddr string, carImage []byte) error {
	mult, err := s.recognitionCl.CalcDamageMultiplier(ctx, carImage)
	if err != nil {
		return fmt.Errorf("failed to payoff: %w", err)
	}

	if err = s.contractCl.Payoff(ctx, receiverAddr, mult); err != nil {
		return fmt.Errorf("failed to payoff: %w", err)
	}

	return nil
}
