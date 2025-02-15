package app

import (
	"github.com/google/wire"
	httpServer "github.com/milad-rasouli/toker/internal/app/http_server"
	urlServer "github.com/milad-rasouli/toker/internal/app/http_server/url"
)

var ProviderSet = wire.NewSet(
	httpServer.NewHttpServer,
	urlServer.NewUrlHttp,
)
