package boot

import (
	"context"
	"github.com/milad-rasouli/toker/internal/entity"
	"github.com/milad-rasouli/toker/internal/service"
	"go.uber.org/zap"
)

type Boot struct {
	logger     *zap.Logger
	urlService service.UrlService
}

func NewBoot(urlService service.UrlService,
	logger *zap.Logger) *Boot {
	return &Boot{
		logger:     logger.Named("Boot"),
		urlService: urlService,
	}
}
func (b *Boot) Boot() error {
	var (
		err error
		lg  = b.logger.With(zap.String("method", "Boot"))
	)
	url := entity.URL{
		URL:    "google.com",
		Detail: "written in C++ and Java mostly.",
	}

	err = b.urlService.CreateUrl(context.TODO(), url)
	if err != nil {
		return err
	}
	instance, err := b.urlService.GetUrl(context.TODO(), url.URL)
	if err != nil {
		return err
	}
	lg.Info("Got ", zap.Any("instance", instance))
	return nil
}
