package cache

import "time"

// Cache interface
type Cache interface {
	// Get value by a given key
	Get(key string) ([]byte, error)
	// Set value
	Set(key string, value []byte, expiration time.Duration) error
}
