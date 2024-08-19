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

func (Converter) FromInsuranceData(in models.InsuranceData) *proto.InsuranceData {
	return &proto.InsuranceData{
		ActiveTill:         in.ActiveTill.String(),
		Owner:              in.Owner,
		Price:              in.Price,
		MaxInsurancePayoff: in.MaxInsurancePayoff,
		MinInsurancePayoff: in.MinInsurancePayoff,
		AvgInsurancePayoff: in.AvgInsurancePayoff,
	}
}
