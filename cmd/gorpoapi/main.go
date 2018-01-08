package main

import (
	"net/http"

	"github.com/talento90/gorpo/config"
	"github.com/talento90/gorpo/effect"
	"github.com/talento90/gorpo/gorpo"
	"github.com/talento90/gorpo/httpapi"
	"github.com/talento90/gorpo/image"
	"github.com/talento90/gorpo/log"
	httprepository "github.com/talento90/gorpo/repository/http"
	"github.com/talento90/gorpo/repository/memory"
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

	imgRepository := httprepository.NewImageRepository()
	effectRepo := memory.NewImageRepository(imgRepository)
	effectService := effect.NewService(effectRepo)

	var imgService gorpo.ImageService
	{
		imgService = image.NewService(imgRepository, effectRepo)
		imgService = image.NewLogService(logger, imgService)
	}

	serverDeps := &httpapi.ServerDependencies{
		Logger:        logger,
		ImgRepository: imgRepository,
		EffectService: effectService,
		ImgService:    imgService,
	}

	serverConfig, err := config.GetServerConfiguration()

	if err != nil {
		logger.Panic(err)
	}

	server := httpapi.NewServer(&serverConfig, serverDeps)

	logger.Info("Starting gorpo API")

	http.ListenAndServe(server.Addr, server.Handler)
}
