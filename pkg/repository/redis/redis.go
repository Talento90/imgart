package redis

import (
	"github.com/go-redis/redis"
)

// Client redis
type Client struct {
	*redis.Client
}

// NewClient creates a redis client
func NewClient(c Configuration) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.Database,
	})

	err := client.Ping().Err()

	if err != nil {
		return nil, err
	}

	return &Client{Client: client}, err
}

// Check redis health
func (c *Client) Check() error {
	return c.Ping().Err()
}
