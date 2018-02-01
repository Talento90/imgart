package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/talento90/imgart/pkg/errors"
	"github.com/talento90/imgart/pkg/imgart"
)

func handleError(err error) error {
	if err == nil {
		return nil
	}

	if err == redis.Nil {
		return errors.ENotExists("Item does not exists", err)
	}

	return errors.EInternal("Error occurred", err)
}

type redisCache struct {
	client *Client
}

// New creates a redis cache implamentation
func New(client *Client) imgart.Cache {
	return &redisCache{client: client}
}

func (r *redisCache) Get(key string) ([]byte, error) {
	result, err := r.client.Get(key).Bytes()

	return result, handleError(err)
}

func (r *redisCache) Set(key string, value []byte, expiration time.Duration) error {
	err := r.client.Set(key, value, expiration).Err()

	return handleError(err)
}

func (r *redisCache) Check() error {
	return r.client.Ping().Err()
}
