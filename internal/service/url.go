package service

import (
	"context"
	"errors"
	"github.com/milad-rasouli/toker/internal/entity"
	"github.com/milad-rasouli/toker/internal/infra/config"
	urlRepo "github.com/milad-rasouli/toker/internal/repository/url"
	"go.uber.org/zap"
)

var ErrNotFound = errors.New("url not found")

type UrlService interface {
	CreateOrGetUrl(ctx context.Context, url string) (*entity.URL, error)
}
type urlService struct {
	logger  *zap.Logger
	env     *config.Config
	urlRepo urlRepo.Repository
}

func NewUrlService(logger *zap.Logger, env *config.Config, repo urlRepo.Repository) UrlService {
	return &urlService{
		logger:  logger.Named("UrlService"),
		env:     env,
		urlRepo: repo,
	}
}
func (u *urlService) CreateOrGetUrl(ctx context.Context, url string) (*entity.URL, error) {
	var (
		err error
		lg  = u.logger.With(zap.String("method", "UrlService.CreateUrl"))
	)
	lg.Info("called with", zap.Any("url", url))

	err = u.urlRepo.SaveUrl(ctx, entity.URL{
		URL:    url,
		Detail: "TODO",
	})
	if err != nil {
		lg.Error("failed to save url", zap.Error(err))
		return nil, err
	}
	return nil, nil
}

//
//func (u *urlService) GetUrl(ctx context.Context, address string) (*entity.URL, error) {
//	var (
//		err error
//		lg  = u.logger.With(zap.String("method", "UrlService.GetUrl"))
//		url entity.URL
//	)
//	lg.Info("called with", zap.Any("url", address))
//	url, err = u.urlRepo.GetUrl(ctx, address)
//	if err != nil {
//		if errors.Is(err, urlRepo.ErrNotFound) {
//			return nil, ErrNotFound
//		}
//		lg.Error("failed to fetch url", zap.Error(err))
//		return nil, err
//	}
//
//	return &url, nil
//}
