package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/milad-rasouli/toker/internal/entity"
	"github.com/milad-rasouli/toker/internal/infra/redis"
	repo "github.com/milad-rasouli/toker/internal/repository/url"
	"github.com/redis/rueidis"
	"go.uber.org/zap"
)

const luaSetUrlTTL = `
redis.call("JSON.SET", KEYS[1], ARGV[1], ARGV[2])
redis.call("EXPIRE", KEYS[1], ARGV[3])
return 1
`

type UrlRepository struct {
	logger    *zap.Logger
	redis     *redis.Redis
	urlScript *rueidis.Lua
}

func NewUrlRepository(logger *zap.Logger, redis *redis.Redis) *UrlRepository {
	return &UrlRepository{
		logger:    logger.Named("UrlRepository"),
		redis:     redis,
		urlScript: rueidis.NewLuaScript(luaSetUrlTTL),
	}
}

func (m *UrlRepository) generateKey(url string) string {
	return "repo:" + url
}

func (m *UrlRepository) SaveUrl(ctx context.Context, url entity.URL) error {
	key := m.generateKey(url.URL)

	jsonData, err := url.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	resp, err := m.urlScript.Exec(ctx, m.redis.Redis, []string{key}, []string{".", string(jsonData), repo.MemoryTTL}).ToInt64()
	if err != nil {
		return fmt.Errorf("failed to store JSON data with TTL: %w", err)
	}

	if resp != 1 {
		return errors.New("unexpected response from Lua script execution")
	}

	return nil
}

func (m *UrlRepository) GetUrl(ctx context.Context, urlAddress string) (entity.URL, error) {
	key := m.generateKey(urlAddress)

	cmd := m.redis.Redis.B().JsonGet().Key(key).Path(".").Build()
	resp, err := m.redis.Redis.Do(ctx, cmd).ToString()

	if err != nil {
		return entity.URL{}, fmt.Errorf("failed to retrieve JSON data: %w", err)
	}

	if resp == "" {
		return entity.URL{}, repo.ErrNotFound
	}

	var url entity.URL
	if err := url.FromString([]byte(resp)); err != nil {
		return entity.URL{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return url, nil
}
