package service

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/contract/internal/clients"
	"github.com/alserov/car_insurance/contract/internal/db"
	"github.com/alserov/car_insurance/contract/internal/service/models"
	"github.com/google/uuid"
)

type Service interface {
	GetNewInsuranceCommits(ctx context.Context) ([]models.OutboxItem, error)
	DeleteCommits(ctx context.Context, ids []string) error
	CreateInsurance(ctx context.Context, ins models.NewInsurance) error
	Payoff(ctx context.Context, ins models.Payoff) error
}

type Clients struct {
	InsuranceClient clients.InsuranceClient
	ContractClient  clients.ContractClient
}

func NewService(cls Clients, outboxRepo db.Outbox) Service {
	return &service{
		outboxRepo:     outboxRepo,
		contractClient: cls.ContractClient,
	}
}

type service struct {
	outboxRepo db.Outbox

	contractClient clients.ContractClient
}

func (s service) DeleteCommits(ctx context.Context, ids []string) error {
	for _, id := range ids {
		err := s.outboxRepo.Delete(ctx, id)
		if err != nil {
			return fmt.Errorf("failed to delete commit: %w", err)
		}
	}

	return nil
}

func (s service) GetNewInsuranceCommits(ctx context.Context) ([]models.OutboxItem, error) {
	items, err := s.outboxRepo.Get(ctx, models.Pending, models.GroupNewInsurances)
	if err != nil {
		return nil, fmt.Errorf("failed to get outbox items: %w", err)
	}

	return items, nil
}

func (s service) CreateInsurance(ctx context.Context, ins models.NewInsurance) error {
	if err := s.contractClient.Insure(ctx, ins); err != nil {
		return fmt.Errorf("failed to create insurance: %w", err)
	}

	if err := s.outboxRepo.Create(ctx, models.OutboxItem{
		ID:      uuid.NewString(),
		GroupID: models.GroupNewInsurances,
		Status:  models.Pending,
	}); err != nil {
		return fmt.Errorf("failed to write to outbox: %w", err)
	}

	return nil
}

func (s service) Payoff(ctx context.Context, pay models.Payoff) error {
	if err := s.contractClient.Payoff(ctx, pay); err != nil {
		return fmt.Errorf("failed to payoff: %w", err)
	}

	return nil
}
