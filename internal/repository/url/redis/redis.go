package redis

import (
	"context"
	"fmt"
	"github.com/milad-rasouli/toker/internal/entity"
	"github.com/milad-rasouli/toker/internal/infra/redis"
	"go.uber.org/zap"
)

type MemoryUrlRepository struct {
	logger *zap.Logger
	redis  *redis.Redis
}

func NewMemoryUrlRepository(logger *zap.Logger, redis *redis.Redis) *MemoryUrlRepository {
	return &MemoryUrlRepository{
		logger: logger,
		redis:  redis,
	}
}
func (m *MemoryUrlRepository) generateKey(url string) string {
	return "repo" + url
}

func (m *MemoryUrlRepository) SaveUrl(ctx context.Context, url entity.URL) error {
	key := m.generateKey(url.URL)

	jsonData, err := url.ToJSON()
	if err != nil {
		return fmt.Errorf("failed ToJSON %w", err)
	}

	cmd := m.redis.Redis.B().JsonSet().Key(key).Path("$").Value(jsonData).Build()
	err = m.redis.Redis.Do(ctx, cmd).Error()
	if err != nil {
		return fmt.Errorf("failed Do %w", err)
	}
}

func (m *MemoryUrlRepository) GetUrl(ctx context.Context, urlAddress string) (entity.URL, error) {
	key := m.generateKey(urlAddress)

	cmd := m.redis.Redis.B().JsonGet().Key(key).Path("$").Build()
	resp, err := m.redis.Redis.Do(ctx, cmd).ToString()
	if err != nil {
		return entity.URL{}, fmt.Errorf("failed Do %w", err)
	}

	var url entity.URL
	err = url.FromString([]byte(resp))
	if err != nil {
		return entity.URL{}, fmt.Errorf("failed FromString %w", err)
	}

	return url, nil
}
