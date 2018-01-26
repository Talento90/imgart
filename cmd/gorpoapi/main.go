package main

import (
	"net/http"

	"github.com/talento90/gorpo/pkg/cache"
	"github.com/talento90/gorpo/pkg/config"
	"github.com/talento90/gorpo/pkg/gorpo"
	"github.com/talento90/gorpo/pkg/health"
	"github.com/talento90/gorpo/pkg/httpapi"
	"github.com/talento90/gorpo/pkg/image"
	"github.com/talento90/gorpo/pkg/log"
	"github.com/talento90/gorpo/pkg/profile"
	httprepository "github.com/talento90/gorpo/pkg/repository/http"
	"github.com/talento90/gorpo/pkg/repository/memory"
	"github.com/talento90/gorpo/pkg/repository/mongo"
	"github.com/talento90/gorpo/pkg/repository/redis"
)

func mongoSession() (*mongo.Session, error) {
	c, err := config.GetMongoConfiguration()

	if err != nil {
		return nil, err
	}

	return mongo.NewSession(c)
}

func redisClient() (*redis.Client, error) {
	c, err := config.GetRedisConfiguration()

	if err != nil {
		return nil, err
	}

	return redis.NewClient(c)
}

func main() {
	logConfig, err := config.GetLogConfiguration()

	if err != nil {
		panic(err)
	}

	logger, err := log.NewLogger(logConfig)

	if err != nil {
		panic(err)
	}

	redisClient, err := redisClient()

	if err != nil {
		logger.Panic(err)
	}

	mongoSession, err := mongoSession()

	if err != nil {
		logger.Panic(err)
	}

	var imgService gorpo.ImageService
	{
		redisCache := redis.New(redisClient)
		imgRepository := httprepository.NewImageRepository()
		effectRepo := memory.NewImageRepository(imgRepository)
		imgCache := cache.NewImage(redisCache)

		imgService = image.NewService(imgRepository, effectRepo)
		imgService = image.NewCacheService(imgCache, imgService)
		imgService = image.NewLogService(logger, imgService)
	}

	var profileService gorpo.ProfileService
	{
		profileRepository := mongo.NewProfileRepository(mongoSession)
		profileService = profile.NewService(profileRepository)
		profileService = profile.NewLogService(logger, profileService)
	}

	health := health.New()
	health.RegisterChecker("redis", redisClient)
	health.RegisterChecker("mongo", mongoSession)

	serverDeps := &httpapi.ServerDependencies{
		Logger:         logger,
		ImgService:     imgService,
		ProfileService: profileService,
		Health:         health,
	}

	serverConfig, err := config.GetServerConfiguration()

	if err != nil {
		logger.Panic(err)
	}

	server := httpapi.NewServer(&serverConfig, serverDeps)

	defer server.Close()

	logger.Info("Starting gorpo API")

	http.ListenAndServe(server.Addr, server.Handler)
}
