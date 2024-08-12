package utils

import (
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	proto "github.com/alserov/car_insurance/insurance/pkg/grpc"
)

type Converter struct{}

func (Converter) ToInsurance(in *proto.NewInsurance) models.Insurance {
	return models.Insurance{
		SenderAddr: in.SenderAddr,
		Amount:     in.Amount,
		CarImage:   in.CarImage,
	}
}

func (Converter) ToPayoff(in *proto.NewPayoff) models.Payoff {
	return models.Payoff{
		ReceiverAddr: in.ReceiverAddr,
		CarImage:     in.CarImage,
	}
}
