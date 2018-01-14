package cache

import "time"

type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, expiration time.Duration) error
}
