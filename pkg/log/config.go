package log

import "io"

// Configuration contains all configurations for logger
type Configuration struct {
	// Level of logging (panic, fatal, error, warn, info, default: debug)
	Level string
	// Ouput (default: os.Stdout)
	Output io.Writer
}

// Validate validates if the configuration is valid
func (c *Configuration) Validate() error {

	return nil
}
