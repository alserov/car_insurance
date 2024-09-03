package service

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	CreateInsurance(ctx context.Context, insData models.Insurance) error
	Payoff(ctx context.Context, payoff models.Payoff) error
	GetInsuranceData(ctx context.Context, ownerAddr string) (models.InsuranceData, error)
	ActivateInsurance(ctx context.Context, id string) error
	ProducePendingInsuranceItems(ctx context.Context) error
	ProducePendingPayoffItems(ctx context.Context) error
}

type Clients struct {
	Recognition    clients.RecognitionClient
	ContractClient clients.ContractClient
}

func NewService(cls Clients, outbox db.Outbox, repo db.Repository) Service {
	return &service{
		recognitionCl: cls.Recognition,
		contractCl:    cls.ContractClient,
		outbox:        outbox,
		repo:          repo,
	}
}

const (
	MinPriceMult = 1.5
	AvgPriceMult = 1.74
	MaxPriceMult = 1.99
)

type service struct {
	recognitionCl clients.RecognitionClient
	contractCl    clients.ContractClient

	repo   db.Repository
	outbox db.Outbox
}

func (s service) ProducePendingInsuranceItems(ctx context.Context) error {
	items, err := s.outbox.Get(ctx, models.Pending, models.GroupInsurance)
	if err != nil {
		return fmt.Errorf("failed to get items from outbox: %w", err)
	}

	for _, val := range items {
		item, ok := val.Val.(models.Insurance)
		if !ok {
			continue
		}

		if err = s.contractCl.CreateInsurance(ctx, item); err != nil {
			continue
		}
	}

	return nil
}

func (s service) ProducePendingPayoffItems(ctx context.Context) error {
	items, err := s.outbox.Get(ctx, models.Pending, models.GroupPayoff)
	if err != nil {
		return fmt.Errorf("failed to get items from outbox: %w", err)
	}

	for _, val := range items {
		item, ok := val.Val.(models.Payoff)
		if !ok {
			continue
		}

		if err = s.contractCl.Payoff(ctx, item); err != nil {
			continue
		}
	}

	return nil
}

func (s service) ActivateInsurance(ctx context.Context, id string) error {
	if err := s.outbox.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete from outbox: %w", err)
	}

	if err := s.repo.UpdateInsuranceStatus(ctx, id, models.Active); err != nil {
		return fmt.Errorf("failed to update insurance status: %w", err)
	}

	return nil
}

func (s service) GetInsuranceData(ctx context.Context, ownerAddr string) (models.InsuranceData, error) {
	data, err := s.repo.GetInsuranceData(ctx, ownerAddr)
	if err != nil {
		return models.InsuranceData{}, fmt.Errorf("failed to get insurance data: %w", err)
	}

	return data, nil
}

func (s service) CreateInsurance(ctx context.Context, insData models.Insurance) error {
	if err := s.recognitionCl.CheckIfCarIsOK(ctx, insData.CarImage); err != nil {
		return fmt.Errorf("failed to create insurance: %w", err)
	}

	insData.ActiveTill = time.Now().Add(models.SixMonthPeriod)

	itemID := uuid.NewString()
	itemStatus := models.Pending
	if err := s.outbox.Create(ctx, models.OutboxItem{
		ID:      itemID,
		GroupID: models.GroupInsurance,
		Status:  itemStatus,
		Val:     insData,
	}); err != nil {
		return fmt.Errorf("failed to write into outbox: %w", err)
	}

	if err := s.repo.CreateInsuranceData(ctx, models.InsuranceData{
		Status:             itemStatus,
		ActiveTill:         insData.ActiveTill,
		ID:                 insData.SenderAddr,
		Price:              insData.Amount,
		MaxInsurancePayoff: int64(float64(insData.Amount) * MaxPriceMult),
		MinInsurancePayoff: int64(float64(insData.Amount) * MinPriceMult),
		AvgInsurancePayoff: int64(float64(insData.Amount) * AvgPriceMult),
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

	itemID := uuid.NewString()
	itemStatus := models.Pending
	if err = s.outbox.Create(ctx, models.OutboxItem{
		ID:      itemID,
		GroupID: models.GroupPayoff,
		Status:  itemStatus,
		Val:     payoff,
	}); err != nil {
		return fmt.Errorf("failed to write into outbox: %w", err)
	}

	return nil
}
