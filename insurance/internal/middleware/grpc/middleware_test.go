package grpc

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"testing"
)

func TestGRPCMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(GRPCMiddlewareSuite))
}

type GRPCMiddlewareSuite struct {
	suite.Suite

	ctx context.Context
}

func (s *GRPCMiddlewareSuite) SetupTest() {
	s.ctx = logger.WrapLogger(context.Background(), logger.NewLogger(logger.Zap, "local"))
}

func (s *GRPCMiddlewareSuite) TestWithErrorHandler() {
	res, err := WithErrorHandler()(s.ctx, "data", &grpc.UnaryServerInfo{}, func(ctx context.Context, req any) (any, error) {
		return nil, utils.NewError("some error", utils.Internal)
	})
	s.Require().Error(err)
	s.Require().Nil(res)

	res, err = WithErrorHandler()(s.ctx, "data", &grpc.UnaryServerInfo{}, func(ctx context.Context, req any) (any, error) {
		return nil, utils.NewError("some error", utils.BadRequest)
	})
	s.Require().Error(err)
	s.Require().Nil(res)

	res, err = WithErrorHandler()(s.ctx, "data", &grpc.UnaryServerInfo{}, func(ctx context.Context, req any) (any, error) {
		return 1, nil
	})
	s.Require().NoError(err)
	s.Require().NotNil(res)
}

func (s *GRPCMiddlewareSuite) TestWithLogger() {
	res, err := WithLogger(logger.NewLogger(logger.Zap, "local"))(s.ctx, "data", &grpc.UnaryServerInfo{}, func(ctx context.Context, req any) (any, error) {
		l := logger.ExtractLogger(ctx)
		s.Require().NotNil(l)
		s.Require().IsType(logger.NewLogger(logger.Zap, "local"), l)

		return 1, nil
	})
	s.Require().NoError(err)
	s.Require().Equal(1, res)

}

func (s *GRPCMiddlewareSuite) TestWithPanicRecovery() {
	res, err := WithRecovery()(s.ctx, "data", &grpc.UnaryServerInfo{}, func(ctx context.Context, req any) (any, error) {
		panic("panic")
		return 1, nil
	})
	s.Require().NoError(err)
	s.Require().Nil(res)

	res, err = WithRecovery()(s.ctx, "data", &grpc.UnaryServerInfo{}, func(ctx context.Context, req any) (any, error) {
		a := make([]struct{}, 0, 1)
		a[0] = struct{}{}

		return 1, nil
	})
	s.Require().NoError(err)
	s.Require().Nil(res)

	res, err = WithRecovery()(s.ctx, "data", &grpc.UnaryServerInfo{}, func(ctx context.Context, req any) (any, error) {
		a := make(chan int)
		close(a)
		close(a)

		return 1, nil
	})
	s.Require().NoError(err)
	s.Require().Nil(res)
}
