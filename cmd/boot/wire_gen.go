// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package boot

import (
	"github.com/milad-rasouli/toker/internal/app/http_server"
	"github.com/milad-rasouli/toker/internal/app/http_server/url"
	"github.com/milad-rasouli/toker/internal/infra/config"
	"github.com/milad-rasouli/toker/internal/infra/redis"
	redis2 "github.com/milad-rasouli/toker/internal/repository/url/redis"
	"github.com/milad-rasouli/toker/internal/service"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func WireApp(cfg *config.Config, logger *zap.Logger, redis3 *redis.Redis) (*Boot, error) {
	urlRepository := redis2.NewUrlRepository(logger, redis3)
	urlService := service.NewUrlService(logger, cfg, urlRepository)
	urlHttp := url_server.NewUrlHttp(logger, urlService)
	httpServer := http_server.NewHttpServer(cfg, logger, urlHttp)
	boot := NewBoot(httpServer, logger)
	return boot, nil
}
