package redis

import (
	"fmt"
	"github.com/redis/rueidis"
	"time"
)

type Redis struct {
	cfg   Config
	Redis rueidis.Client
}

func New(cfg Config) *Redis {
	return &Redis{
		cfg: cfg,
	}
}
func (r *Redis) Setup() error {
	var (
		clientOption = rueidis.ClientOption{
			ClientName:            r.cfg.Name,
			InitAddress:           []string{r.cfg.Host},
			ClientTrackingOptions: nil,
			ShuffleInit:           false,
			DisableRetry:          false,
			RetryDelay: func(attempts int, cmd rueidis.Completed, err error) time.Duration {
				return 3 * time.Second
			},
			DisableCache:          false,
			DisableAutoPipelining: false,
		}
		err error
	)

	r.Redis, err = rueidis.NewClient(clientOption)
	if err != nil {
		return fmt.Errorf("failed to NewClient %w", err)
	}
	return nil
}

func (r *Redis) Close() {
	if r.Redis != nil {
		r.Redis.Close()
	}
}
