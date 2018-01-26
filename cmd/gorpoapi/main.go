package main

import (
	"net/http"

	"github.com/talento90/gorpo/pkg/health"

	"gopkg.in/mgo.v2"

	"github.com/go-redis/redis"
	"github.com/talento90/gorpo/pkg/cache"
	"github.com/talento90/gorpo/pkg/config"
	"github.com/talento90/gorpo/pkg/gorpo"
	"github.com/talento90/gorpo/pkg/httpapi"
	"github.com/talento90/gorpo/pkg/image"
	"github.com/talento90/gorpo/pkg/log"
	"github.com/talento90/gorpo/pkg/profile"
	httprepository "github.com/talento90/gorpo/pkg/repository/http"
	"github.com/talento90/gorpo/pkg/repository/memory"
	"github.com/talento90/gorpo/pkg/repository/mongodb"
	redisrepository "github.com/talento90/gorpo/pkg/repository/redis"
)

func main() {
	logConfig, err := config.GetLogConfiguration()

	if err != nil {
		panic(err)
	}

	logger, err := log.NewLogger(logConfig)

	if err != nil {
		panic(err)
	}

	mongoConfig, err := config.GetMongoConfiguration()

	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial(mongoConfig.MongoURL)

	if err != nil {
		logger.Panic(err)
	}

	defer session.Clone()

	redisConfig, err := config.GetRedisConfiguration()

	if err != nil {
		logger.Panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	})

	redisClient := redisrepository.New(client)
	profileRepository := mongodb.NewProfileRepository(mongoConfig, session)

	var imgService gorpo.ImageService
	{
		imgRepository := httprepository.NewImageRepository()
		effectRepo := memory.NewImageRepository(imgRepository)
		imgCache := cache.NewImage(redisClient)

		imgService = image.NewService(imgRepository, effectRepo)
		imgService = image.NewCacheService(imgCache, imgService)
		imgService = image.NewLogService(logger, imgService)
	}

	var profileService gorpo.ProfileService
	{
		profileService = profile.NewService(profileRepository)
		profileService = profile.NewLogService(logger, profileService)
	}

	health := health.New()
	//health.RegisterChecker("redis", redisClient)
	//health.RegisterChecker("mongo", profileRepository)

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
