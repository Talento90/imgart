package config

import (
	"os"
	"time"

	"github.com/talento90/gorpo/pkg/httpapi"
	"github.com/talento90/gorpo/pkg/log"
	"github.com/talento90/gorpo/pkg/repository/mongodb"
)

// GetLogConfiguration get logger configurations
func GetLogConfiguration() (log.Configuration, error) {
	config := log.Configuration{
		Level:  "debug",
		Output: os.Stdout,
	}

	return config, config.Validate()
}

// GetServerConfiguration get server configurations
func GetServerConfiguration() (httpapi.Configuration, error) {
	config := httpapi.Configuration{
		Address:      ":4005",
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 4,
	}

	return config, config.Validate()
}

// GetMongoConfiguration returns the mongodb configuration
func GetMongoConfiguration() (mongodb.Configuration, error) {
	config := mongodb.Configuration{
		Database: "gorpo",
		MongoURL: "localhost:27017",
		//MongoURL: os.Getenv("MONGO_SERVICE_NAME"),
	}

	return config, config.Validate()
}
