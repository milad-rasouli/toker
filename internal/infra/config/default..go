package config

import (
	"github.com/milad-rasouli/toker/internal/infra/logger"
	"github.com/milad-rasouli/toker/internal/infra/redis"
)

func Default() *Config {
	return &Config{
		App: app{
			Name: "toker",
			Port: ":5001",
		},
		Redis: redis.Config{
			Name: "tocker",
			Host: "127.0.0.1:6379",
		},
		Logger: logger.Config{
			Level: "info",
		},
	}
}
