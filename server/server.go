package server

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type S struct {
	Config Config
}

type Config struct {
	DiagPort    int `envconfig:"DIAG_PORT" default:"8081" required:"true"`
	RESTAPIPort int `envconfig:"PORT" default:"8080" required:"true"`
}

func New(logger *zap.SugaredLogger) (*S, error) {
	conf := Config{}
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}

	return &S{
		Config: conf,
	}, nil
}
