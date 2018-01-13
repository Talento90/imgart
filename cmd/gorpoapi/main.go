package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/talento90/gorpo/pkg/config"
	"github.com/talento90/gorpo/pkg/gorpo"
	"github.com/talento90/gorpo/pkg/httpapi"
	"github.com/talento90/gorpo/pkg/image"
	"github.com/talento90/gorpo/pkg/log"
	"github.com/talento90/gorpo/pkg/profile"
	httprepository "github.com/talento90/gorpo/pkg/repository/http"
	"github.com/talento90/gorpo/pkg/repository/memory"
	"github.com/talento90/gorpo/pkg/repository/mongodb"
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

	var imgService gorpo.ImageService
	{
		imgRepository := httprepository.NewImageRepository()
		effectRepo := memory.NewImageRepository(imgRepository)
		imgService = image.NewService(imgRepository, effectRepo)
		imgService = image.NewLogService(logger, imgService)
	}

	var profileService gorpo.ProfileService
	{
		profileRepository := mongodb.NewProfileRepository(mongoConfig, session)
		profileService = profile.NewService(profileRepository)
		profileService = profile.NewLogService(logger, profileService)
	}

	serverDeps := &httpapi.ServerDependencies{
		Logger:         logger,
		ImgService:     imgService,
		ProfileService: profileService,
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
