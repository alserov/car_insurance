package grpc

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/clients"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
	"github.com/alserov/car_insurance/gateway/internal/utils"
	proto "github.com/alserov/car_insurance/insurance/pkg/grpc"
	"google.golang.org/grpc"
)

func NewInsuranceClient(cc *grpc.ClientConn) clients.InsuranceClient {
	cl := proto.NewInsuranceClient(cc)

	return &insurance{
		cl: cl,
	}
}

type insurance struct {
	cl proto.InsuranceClient
}

func (i insurance) CreateInsurance(ctx context.Context, insurance models.Insurance) error {
	_, err := i.cl.CreateInsurance(ctx, &proto.NewInsurance{
		SenderAddr: insurance.SenderAddr,
		Amount:     insurance.Amount,
		CarImage:   insurance.CarImage,
	})
	if err != nil {
		return utils.FromGRPCError(err)
	}

	return nil
}

func (i insurance) Payoff(ctx context.Context, payoff models.Payoff) error {
	_, err := i.cl.Payoff(ctx, &proto.NewPayoff{
		ReceiverAddr: payoff.ReceiverAddr,
		CarImage:     payoff.CarImage,
	})
	if err != nil {
		return utils.FromGRPCError(err)
	}

	return nil
}

func (i insurance) GetInsuranceData(ctx context.Context, addr string) (models.InsuranceData, error) {
	data, err := i.cl.GetInsuranceData(ctx, &proto.InsuranceOwner{Addr: addr})
	if err != nil {
		return models.InsuranceData{}, utils.FromGRPCError(err)
	}

	return models.InsuranceData{
		Status:             0,
		ActiveTill:         data.ActiveTill,
		Owner:              data.Owner,
		Price:              data.Price,
		MaxInsurancePayoff: data.MaxInsurancePayoff,
		MinInsurancePayoff: data.MinInsurancePayoff,
		AvgInsurancePayoff: data.AvgInsurancePayoff,
	}, nil
}
