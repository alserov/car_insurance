package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env    string
	Broker Kafka

	Databases struct {
		Redis struct {
			Addr string
		}
	}

	Contract struct {
		Addr string
	}
}

type Kafka struct {
	Addr   string
	Topics struct {
		Payoff       string
		NewInsurance string
		Commits      string
	}
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")

	cfg.Databases.Redis.Addr = fmt.Sprintf(fmt.Sprintf("redis://%s@%s:%s", os.Getenv("REDIS_USER"), os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")))

	cfg.Contract.Addr = os.Getenv("CONTRACT_ADDR")

	cfg.Broker.Addr = os.Getenv("KAFKA_ADDR")
	cfg.Broker.Topics.NewInsurance = os.Getenv("INSURANCE_TOPIC")
	cfg.Broker.Topics.Payoff = os.Getenv("PAYOFF_TOPIC")

	return &cfg
}
