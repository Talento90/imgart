package config

import (
	"os"
	"strconv"
	"time"

	"github.com/talento90/imgart/httpapi"
	"github.com/talento90/imgart/log"
	"github.com/talento90/imgart/repository/mongo"
	"github.com/talento90/imgart/repository/redis"
)

// GetLogConfiguration get logger configurations
func GetLogConfiguration() (log.Configuration, error) {
	config := log.Configuration{
		Level:  getEnv("LOG_LEVEL", "debug"),
		Output: os.Stdout,
	}

	return config, config.Validate()
}

// GetServerConfiguration get server configurations
func GetServerConfiguration() (httpapi.Configuration, error) {
	config := httpapi.Configuration{
		Address:      getEnv("PORT", "4005"),
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 4,
	}

	return config, config.Validate()
}

// GetMongoConfiguration returns the mongo configuration
func GetMongoConfiguration() (mongo.Configuration, error) {
	config := mongo.Configuration{
		Database: getEnv("MONGO_DATABASE", "imgart"),
		MongoURL: getEnv("MONGO_URL", "localhost:27017"),
	}

	return config, config.Validate()
}

// GetRedisConfiguration returns the redis configuration
func GetRedisConfiguration() (redis.Configuration, error) {
	db, err := strconv.Atoi(getEnv("REDIS_SERVICE_DB", "0"))

	if err != nil {
		db = 0
	}

	config := redis.Configuration{
		Address:  getEnv("REDIS_URL", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		Database: db,
	}

	return config, config.Validate()
}

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
