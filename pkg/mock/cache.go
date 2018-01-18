package mock

import (
	"sync"
	"time"

	"github.com/talento90/gorpo/pkg/errors"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type mockCache struct {
	mutex      *sync.Mutex
	repository map[string][]byte
}

func (c *mockCache) Get(key string) ([]byte, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if value, ok := c.repository[key]; ok {
		return value, nil
	}

	return nil, errors.ENotExists("Item does not exists", nil)
}

func (c *mockCache) Set(key string, value []byte, expiration time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.repository[key] = value

	return nil
}

// NewCache returns a new mock of Cache interface
func NewCache() gorpo.Cache {
	return &mockCache{
		mutex:      &sync.Mutex{},
		repository: map[string][]byte{},
	}
}
