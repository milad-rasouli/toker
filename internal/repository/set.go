package repository

import (
	"github.com/google/wire"
	mRepo "github.com/milad-rasouli/toker/internal/repository/url"
	redisRepo "github.com/milad-rasouli/toker/internal/repository/url/redis"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(mRepo.Repository), new(*redisRepo.UrlRepository)),
	redisRepo.NewUrlRepository,
)
