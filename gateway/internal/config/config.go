package config

import "os"

type Config struct {
	Env     string
	Port    string
	Tracing struct {
		Endpoint string
		Name     string
	}
	Services struct {
		Insurance struct {
			Addr string
		}
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

	cfg.Services.Insurance.Addr = os.Getenv("INSURANCE_ADDR")

	return &cfg
}
