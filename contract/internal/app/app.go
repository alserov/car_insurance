package app

import (
	"context"
	"github.com/alserov/car_insurance/contract/internal/async"
	async_cl "github.com/alserov/car_insurance/contract/internal/clients/async"
	"github.com/alserov/car_insurance/contract/internal/clients/ethereum"
	"github.com/alserov/car_insurance/contract/internal/config"
	api "github.com/alserov/car_insurance/contract/internal/contracts"
	"github.com/alserov/car_insurance/contract/internal/db/redis"
	"github.com/alserov/car_insurance/contract/internal/logger"
	"github.com/alserov/car_insurance/contract/internal/service"
	"github.com/alserov/car_insurance/contract/internal/service/models"
	"github.com/alserov/car_insurance/contract/internal/workers"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

func MustStart(cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// initializing logger
	log := logger.NewLogger(logger.Zap, cfg.Env)

	log.Info("starting server")

	// db (connecting to DBs)
	redisConn := redis.MustConnect(cfg.Databases.Redis.Addr)
	defer func() {
		_ = redisConn.Close()
	}()

	// repositories
	outboxRepo := redis.NewOutbox(redisConn)

	// *ethereum* client
	api, cl := api.MustSetupContract(cfg.Contract.Addr)

	contractClient := ethereum.NewContractClient(api, cl)

	// *async* clients
	insuranceProducer := async.NewProducer(async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.Commits)
	defer func() {
		_ = insuranceProducer.Close()
	}()

	insurancePayoffConsumer := async.NewConsumer[models.Payoff](async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.Payoff)
	defer func() {
		_ = insurancePayoffConsumer.Close()
	}()

	insuranceNewInsuranceConsumer := async.NewConsumer[models.NewInsurance](async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.NewInsurance)
	defer func() {
		_ = insuranceNewInsuranceConsumer.Close()
	}()

	insuranceClient := async_cl.NewInsuranceClient(insuranceProducer, insuranceNewInsuranceConsumer, insurancePayoffConsumer)

	cls := service.Clients{
		InsuranceClient: insuranceClient,
		ContractClient:  contractClient,
	}

	// service
	srvc := service.NewService(cls, outboxRepo)

	// initializing workers
	outboxWorker := workers.NewOutboxWorker(srvc, cls, log)
	insuranceWorker := workers.NewInsuranceWorker(srvc, cls, log)

	log.Info("server is running")

	outboxWorker.Start(ctx)
	insuranceWorker.Start(ctx)

	<-ctx.Done()

	log.Info("server stopped")
}
