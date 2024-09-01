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

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.Tracing.Name = os.Getenv("TRACING_NAME")
	cfg.Tracing.Endpoint = os.Getenv("TRACING_EP")

	cfg.Services.Insurance.Addr = os.Getenv("INSURANCE_ADDR")

	return &cfg
}
