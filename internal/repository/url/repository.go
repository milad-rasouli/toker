package reposotory

import (
	"context"
	"github.com/milad-rasouli/toker/internal/entity"
)

type Repository interface {
	SaveUrl(ctx context.Context, url entity.URL) error
	GetUrl(ctx context.Context, urlAddress string) (entity.URL, error)
}
