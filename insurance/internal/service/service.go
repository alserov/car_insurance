package service

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"time"
)

type Service interface {
	CreateInsurance(ctx context.Context, insData models.Insurance) error
	Payoff(ctx context.Context, payoff models.Payoff) error
	GetInsuranceData(ctx context.Context, ownerAddr string) (models.InsuranceData, error)
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

	repo   db.Repository
	outbox db.Outbox
}

func (s service) GetInsuranceData(ctx context.Context, ownerAddr string) (models.InsuranceData, error) {
	data, err := s.repo.GetInsuranceData(ownerAddr)
	if err != nil {
		return models.InsuranceData{}, fmt.Errorf("failed to get insurance data: %w", err)
	}

	return data, nil
}

func (s service) CreateInsurance(ctx context.Context, insData models.Insurance) error {
	if err := s.recognitionCl.CheckIfCarIsOK(ctx, insData.CarImage); err != nil {
		return fmt.Errorf("failed to create insurance: %w", err)
	}

	if err := s.outbox.Create(ctx, insData.SenderAddr, struct{}{}); err != nil {
		return fmt.Errorf("failed to write into outbox: %w", err)
	}

	insData.ActiveTill = time.Now().Add(models.SixMonthPeriod)

	if err := s.contractCl.CreateInsurance(ctx, insData); err != nil {
		return fmt.Errorf("failed to create insurance: %w", err)
	}

	if err := s.repo.CreateInsuranceData(models.InsuranceData{
		Status:             models.Pending,
		ActiveTill:         insData.ActiveTill,
		Owner:              insData.SenderAddr,
		Price:              insData.Amount,
		MaxInsurancePayoff: int64(float64(insData.Amount) * 1.99),
		MinInsurancePayoff: int64(float64(insData.Amount) * 1.5),
		AvgInsurancePayoff: int64(float64(insData.Amount) * 1.74),
	}); err != nil {
		return fmt.Errorf("failed to create insurance data: %w", err)
	}

	return nil
}

func (s service) Payoff(ctx context.Context, payoff models.Payoff) error {
	mult, err := s.recognitionCl.CalcDamageMultiplier(ctx, payoff.CarImage)
	if err != nil {
		return fmt.Errorf("failed to payoff: %w", err)
	}

	payoff.Multiplier = mult

	if err = s.outbox.Create(ctx, payoff.ReceiverAddr, struct{}{}); err != nil {
		return fmt.Errorf("failed to write into outbox: %w", err)
	}

	if err = s.contractCl.Payoff(ctx, payoff); err != nil {
		return fmt.Errorf("failed to payoff: %w", err)
	}

	return nil
}
