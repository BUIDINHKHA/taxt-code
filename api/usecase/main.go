package usecase

import (
	"megabot/config"
	"megabot/pkg/logger"
)

type Config struct {
	log *logger.Logger
	cfg *config.Environment
}

func NewConfig(
	log *logger.Logger,
	cfg *config.Environment,
) *Config {
	return &Config{
		log: log,
		cfg: cfg,
	}
}
