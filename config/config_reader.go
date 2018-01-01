package config

import (
	"os"
	"time"

	"github.com/talento90/gorpo/httpapi"
	"github.com/talento90/gorpo/log"
)

//GetLogConfiguration get logger configurations
func GetLogConfiguration() (log.Configuration, error) {
	config := log.Configuration{
		Level:  "debug",
		Output: os.Stdout,
	}

	return config, config.Validate()
}

//GetServerConfiguration get server configurations
func GetServerConfiguration() (httpapi.Configuration, error) {
	config := httpapi.Configuration{
		Address:      ":4005",
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 4,
	}

	return config, config.Validate()
}
