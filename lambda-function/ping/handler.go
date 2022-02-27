package ping

import (
	"context"
)

type Config struct {
	Version     string `env:"VERSION"`
	Environment string `env:"ENVIRONMENT"`
	AWSRegion   string `env:"AWS_REGION"`
}

type response struct {
	Message     string `json:"message"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
	AWSRegion   string `json:"aws-region"`
}

func (c *Config) Ping(ctx context.Context) (response, error) {
	resp := response{
		Message:     "pong",
		Version:     c.Version,
		Environment: c.Environment,
		AWSRegion:   c.AWSRegion,
	}

	return resp, nil
}
