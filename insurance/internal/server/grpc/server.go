package grpc

import (
	"context"
	"fmt"
	"github.com/alserov/car_insurance/insurance/internal/app"
	mw "github.com/alserov/car_insurance/insurance/internal/middleware/grpc"
	"github.com/alserov/car_insurance/insurance/internal/service"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	proto "github.com/alserov/car_insurance/insurance/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

func NewServer(srvc service.Service) app.Server {
	s := grpc.NewServer(
		mw.ChainUnaryServer(
			mw.WithErrorHandler(),
			mw.WithRecovery(),
		),
	)

	proto.RegisterInsuranceServer(s, &server{})

	return &server{
		grpcServer: s,
		srvc:       srvc,
	}
}

type server struct {
	proto.UnimplementedInsuranceServer

	grpcServer *grpc.Server

	srvc service.Service

	conv utils.Converter
}

func (s server) Serve(port string) error {
	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	if err = s.grpcServer.Serve(l); err != nil {
		return err
	}

	return nil
}

func (s server) CreateInsurance(ctx context.Context, insurance *proto.NewInsurance) (*emptypb.Empty, error) {
	if err := s.srvc.CreateInsurance(ctx, s.conv.ToInsurance(insurance)); err != nil {
		return nil, fmt.Errorf("service failed: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (s server) Payoff(ctx context.Context, payoff *proto.NewPayoff) (*emptypb.Empty, error) {
	if err := s.srvc.Payoff(ctx, s.conv.ToPayoff(payoff)); err != nil {
		return nil, fmt.Errorf("service failed: %w", err)
	}

	return &emptypb.Empty{}, nil
}
