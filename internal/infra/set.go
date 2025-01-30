package infra

import (
	"github.com/google/wire"
	"github.com/milad-rasouli/toker/internal/infra/config"
)

var ProviderSet = wire.NewSet(
	config.New,
)
