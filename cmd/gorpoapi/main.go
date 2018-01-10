package main

import (
	"net/http"

	"github.com/talento90/gorpo/pkg/config"
	"github.com/talento90/gorpo/pkg/gorpo"
	"github.com/talento90/gorpo/pkg/gorpo/image"
	"github.com/talento90/gorpo/pkg/httpapi"
	"github.com/talento90/gorpo/pkg/log"
	httprepository "github.com/talento90/gorpo/pkg/repository/http"
	"github.com/talento90/gorpo/pkg/repository/memory"
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

	var imgService gorpo.ImageService
	{
		imgRepository := httprepository.NewImageRepository()
		effectRepo := memory.NewImageRepository(imgRepository)
		imgService = image.NewService(imgRepository, effectRepo)
		imgService = image.NewLogService(logger, imgService)
	}

	serverDeps := &httpapi.ServerDependencies{
		Logger:     logger,
		ImgService: imgService,
	}

	serverConfig, err := config.GetServerConfiguration()

	if err != nil {
		logger.Panic(err)
	}

	server := httpapi.NewServer(&serverConfig, serverDeps)

	logger.Info("Starting gorpo API")

	http.ListenAndServe(server.Addr, server.Handler)
}
