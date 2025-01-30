package main

import (
	"github.com/milad-rasouli/toker/internal/infra/config"
	"github.com/milad-rasouli/toker/internal/infra/logger"
	"github.com/milad-rasouli/toker/internal/infra/redis"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to setup config %s\n", err)
	}

	logger, closeLogger := logger.New(cfg.Logger)
	defer closeLogger()

	rdis := redis.New(cfg.Redis)
	err = rdis.Setup()
	if err != nil {
		log.Fatalf("failed to setup redis %s\n", err)
	}

}
