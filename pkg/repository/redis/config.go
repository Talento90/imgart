package redis

import "github.com/talento90/gorpo/pkg/errors"

// Configuration for Redis
type Configuration struct {
	Address  string
	Password string
	Database int
}

// Validate validates if the configuration is valid
func (c *Configuration) Validate() error {

	if c.Address == "" {
		return errors.EValidation("Missing Address", nil)
	}

	return nil
}
