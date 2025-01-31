package reposotory

import (
	"context"
	"errors"
	"github.com/milad-rasouli/toker/internal/entity"
)

const MemoryTTL = "60"

var ErrNotFound = errors.New("not found")

type Repository interface {
	SaveUrl(ctx context.Context, url entity.URL) error
	GetUrl(ctx context.Context, urlAddress string) (entity.URL, error)
}
