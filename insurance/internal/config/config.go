package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env  string
	Port string

	Broker Kafka

	Services struct {
		RecognitionAddr string
	}

	Databases struct {
		Mongo struct {
			Addr string
		}
		Postgres struct {
			Addr string
		}
	}
}

type Kafka struct {
	Addr   string
	Topics struct {
		Payoff       string
		NewInsurance string
		Commit       string
	}
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.Services.RecognitionAddr = os.Getenv("RECOGNITION_ADDR")

	cfg.Databases.Postgres.Addr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_DB"))

	cfg.Broker.Addr = os.Getenv("KAFKA_ADDR")
	cfg.Broker.Topics.NewInsurance = os.Getenv("INSURANCE_TOPIC")
	cfg.Broker.Topics.Payoff = os.Getenv("PAYOFF_TOPIC")

	return &cfg
}
