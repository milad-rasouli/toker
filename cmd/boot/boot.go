package boot

import (
	httpServer "github.com/milad-rasouli/toker/internal/app/http_server"
	"go.uber.org/zap"
)

type Boot struct {
	logger  *zap.Logger
	hServer *httpServer.HttpServer
}

func NewBoot(h *httpServer.HttpServer,
	logger *zap.Logger) *Boot {
	return &Boot{
		logger:  logger.Named("Boot"),
		hServer: h,
	}
}
func (b *Boot) Boot() error {
	var (
		err error
		lg  = b.logger.With(zap.String("method", "Boot"))
	)
	err = b.hServer.Start()
	if err != nil {
		lg.Error("failed to start http server", zap.Error(err))
		return err
	}
	return nil
}
