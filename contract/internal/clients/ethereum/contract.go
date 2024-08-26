package ethereum

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/contract/internal/clients"
	api "github.com/alserov/car_insurance/contract/internal/contracts"
	"github.com/alserov/car_insurance/contract/internal/service/models"
	"github.com/alserov/car_insurance/contract/internal/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func NewContractClient(api *api.Api, ethCl *ethclient.Client) clients.ContractClient {
	return &contract{
		api:   api,
		ethCl: ethCl,
	}
}

type contract struct {
	api   *api.Api
	ethCl *ethclient.Client
}

func (c contract) Payoff(ctx context.Context, pay models.Payoff) error {
	auth, err := api.GetAccountAuth(ctx, c.ethCl, pay.Receiver)
	if err != nil {
		return fmt.Errorf("failed to get account auth: %w", err)
	}

	tx, err := c.api.Payoff(auth, big.NewInt(pay.Mult))
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	if err = api.CheckTransactionReceipt(c.ethCl, tx.Hash().String()); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	return nil
}

func (c contract) Insure(ctx context.Context, ins models.NewInsurance) error {
	auth, err := api.GetAccountAuth(ctx, c.ethCl, ins.Sender)
	if err != nil {
		return fmt.Errorf("failed to get account auth: %w", err)
	}

	tx, err := c.api.Insure(auth, big.NewInt(ins.Amount), big.NewInt(ins.ActiveTill.Unix()))
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	if err = api.CheckTransactionReceipt(c.ethCl, tx.Hash().String()); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	return nil
}
