package api

import (
	"context"
	"crypto/ecdsa"
	"github.com/alserov/car_insurance/contract/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func MustSetupContract(contractAddr string) (*Api, *ethclient.Client) {
	cl, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic("failed to dial: " + err.Error())
	}

	contrAddr := MustDeploy(cl, contractAddr)

	conn, err := NewApi(common.HexToAddress(contrAddr.Hex()), cl)
	if err != nil {
		panic(err)
	}

	return conn, cl
}

func GetAccountAuth(ctx context.Context, cl *ethclient.Client, accAddr string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(accAddr)
	if err != nil {
		return nil, utils.NewError("invalid address", utils.BadRequest)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, utils.NewError("failed to get pub key", utils.BadRequest)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, utils.NewError("failed to get nonce", utils.BadRequest)
	}

	gasPrice, err := cl.SuggestGasPrice(ctx)
	if err != nil {
		return nil, utils.NewError("failed to gas price: "+err.Error(), utils.Internal)
	}

	chainID, err := cl.ChainID(ctx)
	if err != nil {
		return nil, utils.NewError("failed to chain id price: "+err.Error(), utils.Internal)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, utils.NewError("failed to get auth: "+err.Error(), utils.Internal)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

// MustDeploy Deploying a smart contract with an admin account
func MustDeploy(cl *ethclient.Client, contractAddr string) common.Address {
	auth, err := GetAccountAuth(context.Background(), cl, contractAddr)
	if err != nil {
		panic("failed to get account auth: " + err.Error())
	}

	//deploying smart contract
	deployedContractAddress, _, _, err := DeployApi(auth, cl) //api is redirected from api directory from our contract go file
	if err != nil {
		panic("failed to deploy: " + err.Error())
	}

	return deployedContractAddress
}

func CheckTransactionReceipt(cl *ethclient.Client, _txHash string) error {
	txHash := common.HexToHash(_txHash)
	tx, err := cl.TransactionReceipt(context.Background(), txHash)
	if tx.Status == 1 {
		return nil
	}

	if err != nil {
		return utils.NewError("failed to complete tx: "+err.Error(), utils.BadRequest)
	}

	return utils.NewError("failed to complete tx: invalid address or amount", utils.BadRequest)
}
