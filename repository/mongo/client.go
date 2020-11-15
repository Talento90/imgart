package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client mongo client
type Client struct {
	*mongo.Client
	Database string
}

// NewClient creates a new mongo session
func NewClient(c Configuration) (*Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(c.MongoURL))

	if err != nil {
		return nil, err
	}

	if err := client.Connect(context.Background()); err != nil {
		return nil, err
	}

	return &Client{Client: client, Database: c.Database}, nil
}

// Check mongo health
func (c *Client) Check() error {
	return c.Ping(context.Background(), readpref.Primary())
}
