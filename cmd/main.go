package main

import (
	"github.com/milad-rasouli/toker/cmd/boot"
	"log"

	"github.com/milad-rasouli/toker/internal/infra/config"
	"github.com/milad-rasouli/toker/internal/infra/logger"
	"github.com/milad-rasouli/toker/internal/infra/redis"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to setup config %s\n", err)
	}

	logger, closeLogger := logger.New(cfg.Logger)
	defer closeLogger()

	redis := redis.New(cfg.Redis)
	err = redis.Setup()
	if err != nil {
		log.Fatalf("failed to setup redis %s\n", err)
	}

	boot, err := boot.WireApp(cfg, logger, redis)
	if err != nil {
		log.Fatalf("failed to setup app %s\n", err)
	}
	err = boot.Boot()
	if err != nil {
		log.Fatalf("failed to boot app %s\n", err)
	}
}
