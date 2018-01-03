package main

import (
	"net/http"

	"github.com/talento90/gorpo/config"
	"github.com/talento90/gorpo/downloader"
	"github.com/talento90/gorpo/effect"
	"github.com/talento90/gorpo/httpapi"
	"github.com/talento90/gorpo/image"
	"github.com/talento90/gorpo/log"
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

	httpDownloader := downloader.NewHTTPDownloader()
	effectRepo := memory.NewEffectRepository(httpDownloader)
	effectService := effect.NewService(effectRepo)
	imgService := image.NewService(httpDownloader, effectRepo)

	serverDeps := &httpapi.ServerDependencies{
		Logger:        logger,
		Downloader:    httpDownloader,
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
