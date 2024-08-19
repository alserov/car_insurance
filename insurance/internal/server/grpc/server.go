package grpc

import (
	"context"
	"fmt"
	mw "github.com/alserov/car_insurance/insurance/internal/middleware/grpc"
	"github.com/alserov/car_insurance/insurance/internal/server"
	"github.com/alserov/car_insurance/insurance/internal/service"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	proto "github.com/alserov/car_insurance/insurance/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

func NewServer(srvc service.Service) server.Server {
	s := grpc.NewServer(
		mw.ChainUnaryServer(
			mw.WithErrorHandler(),
			mw.WithRecovery(),
		),
	)

	proto.RegisterInsuranceServer(s, &grpcServer{})

	return &grpcServer{
		grpcServer: s,
		srvc:       srvc,
	}
}

type grpcServer struct {
	proto.UnimplementedInsuranceServer

	grpcServer *grpc.Server

	srvc service.Service

	conv utils.Converter
}

func (s grpcServer) Shutdown() error {
	s.grpcServer.GracefulStop()
	return nil
}

func (s grpcServer) Serve(port string) error {
	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	if err = s.grpcServer.Serve(l); err != nil {
		return err
	}

	return nil
}

func (s grpcServer) GetInsuranceData(ctx context.Context, owner *proto.InsuranceOwner) (*proto.InsuranceData, error) {
	data, err := s.srvc.GetInsuranceData(ctx, owner.Addr)
	if err != nil {
		return nil, fmt.Errorf("service failed: %w", err)
	}

	return s.conv.FromInsuranceData(data), nil
}

func (s grpcServer) CreateInsurance(ctx context.Context, insurance *proto.NewInsurance) (*emptypb.Empty, error) {
	if err := s.srvc.CreateInsurance(ctx, s.conv.ToInsurance(insurance)); err != nil {
		return nil, fmt.Errorf("service failed: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (s grpcServer) Payoff(ctx context.Context, payoff *proto.NewPayoff) (*emptypb.Empty, error) {
	if err := s.srvc.Payoff(ctx, s.conv.ToPayoff(payoff)); err != nil {
		return nil, fmt.Errorf("service failed: %w", err)
	}

	return &emptypb.Empty{}, nil
}
