package service

import (
	"context"
	"fmt"
	api "github.com/alserov/car_insurance/contract/internal/contracts"
	"github.com/alserov/car_insurance/contract/internal/service/models"
	"github.com/alserov/car_insurance/contract/internal/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Service interface {
	CreateInsurance(ctx context.Context, ins models.NewInsurance) error
	Payoff(ctx context.Context, ins models.Payoff) error
}

func NewService(api *api.Api, ethCl *ethclient.Client) Service {
	return &service{
		insAPI: api,
		ethCl:  ethCl,
	}
}

type service struct {
	insAPI *api.Api

	ethCl *ethclient.Client
}

func (s service) CreateInsurance(ctx context.Context, ins models.NewInsurance) error {
	auth, err := api.GetAccountAuth(ctx, s.ethCl, ins.Sender)
	if err != nil {
		return fmt.Errorf("failed to get account auth: %w", err)
	}

	tx, err := s.insAPI.Insure(auth, big.NewInt(ins.Amount), big.NewInt(ins.ActiveTill.Unix()))
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	if err = api.CheckTransactionReceipt(s.ethCl, tx.Hash().String()); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	return nil
}

func (s service) Payoff(ctx context.Context, ins models.Payoff) error {
	auth, err := api.GetAccountAuth(ctx, s.ethCl, ins.Receiver)
	if err != nil {
		return fmt.Errorf("failed to get account auth: %w", err)
	}

	tx, err := s.insAPI.Payoff(auth, big.NewInt(ins.Mult))
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	if err = api.CheckTransactionReceipt(s.ethCl, tx.Hash().String()); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	return nil
}
