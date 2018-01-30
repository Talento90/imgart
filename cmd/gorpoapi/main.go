package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

func httpServer(l log.Logger, rc *redis.Client, ms *mongo.Session) *http.Server {
	var imgService gorpo.ImageService
	{
		redisCache := redis.New(rc)
		imgRepository := httprepository.NewImageRepository()
		effectRepo := memory.NewImageRepository(imgRepository)
		imgCache := cache.NewImage(redisCache)

		imgService = image.NewService(imgRepository, effectRepo)
		imgService = image.NewCacheService(imgCache, imgService)
		imgService = image.NewLogService(l, imgService)
	}

	var profileService gorpo.ProfileService
	{
		profileRepository := mongo.NewProfileRepository(ms)
		profileService = profile.NewService(profileRepository)
		profileService = profile.NewLogService(l, profileService)
	}

	health := health.New("gorpo")
	health.RegisterChecker("redis", rc)
	health.RegisterChecker("mongo", ms)

	serverDeps := &httpapi.ServerDependencies{
		Logger:         l,
		ImgService:     imgService,
		ProfileService: profileService,
		Health:         health,
	}

	serverConfig, err := config.GetServerConfiguration()

	if err != nil {
		l.Panic(err)
	}

	srv := httpapi.NewServer(&serverConfig, serverDeps)

	return &srv
}

func main() {
	var gracefulShutdown = make(chan os.Signal)

	signal.Notify(gracefulShutdown, syscall.SIGTERM)
	signal.Notify(gracefulShutdown, syscall.SIGINT)

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
		logger.Panic("Error connecting to Redis", err)
	}

	mongoSession, err := mongoSession()

	if err != nil {
		logger.Panic("Error connecting to Mongo", err)
	}

	server := httpServer(logger, redisClient, mongoSession)

	go func() {
		<-gracefulShutdown
		exitCode := 0

		logger.Info("Starting graceful shutdown")

		err = server.Shutdown(context.Background())

		if err != nil {
			exitCode = 1
			logger.Error("Error closing server:", err)
		}

		mongoSession.Close()

		err = redisClient.Close()

		if err != nil {
			exitCode = 1
			logger.Error("Error closing redis:", err)
		}

		logger.Info("Graceful shutdown completed")

		os.Exit(exitCode)
	}()

	logger.Info("Server listening on port: ", server.Addr)

	http.ListenAndServe(server.Addr, server.Handler)
}
