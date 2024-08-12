package async

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
)

func NewContractClient(addr, newInsuranceTopic, payoffTopic string) clients.ContractClient {
	return &contract{}
}

type contract struct {
	p sarama.SyncProducer
}

func (c contract) CreateInsurance(ctx context.Context, ins models.Insurance) error {
	//TODO implement me
	panic("implement me")
}

func (c contract) Payoff(ctx context.Context, receiverAddr string, mult float32) error {
	//TODO implement me
	panic("implement me")
}
