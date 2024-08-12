package config

import "os"

type Config struct {
	Env      string
	Port     string
	Broker   Kafka
	Contract struct {
		Addr string
	}
}

type Kafka struct {
	Addr   string
	Topics struct {
		Payoff       string
		NewInsurance string
	}
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.Broker.Addr = os.Getenv("KAFKA_ADDR")
	cfg.Broker.Topics.NewInsurance = os.Getenv("INSURANCE_TOPIC")
	cfg.Broker.Topics.Payoff = os.Getenv("PAYOFF_TOPIC")

	return &cfg
}
