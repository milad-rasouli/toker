//go:build wireinject
// +build wireinject

package boot

import (
	"github.com/google/wire"
	app "github.com/milad-rasouli/toker/internal/app"
	configInfra "github.com/milad-rasouli/toker/internal/infra/config"
	redisInfra "github.com/milad-rasouli/toker/internal/infra/redis"
	Repo "github.com/milad-rasouli/toker/internal/repository"
	Svc "github.com/milad-rasouli/toker/internal/service"
	"go.uber.org/zap"
)

func WireApp(
	cfg *configInfra.Config,
	logger *zap.Logger,
	redis *redisInfra.Redis,
) (*Boot, error) {
	panic(wire.Build(
		Repo.ProviderSet,
		Svc.ProviderSet,
		app.ProviderSet,
		wire.NewSet(NewBoot),
	))
}
