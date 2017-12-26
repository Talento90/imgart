package main

import (
	"net/http"
	"os"

	"github.com/talento90/gorpo/downloader"
	"github.com/talento90/gorpo/effect"
	"github.com/talento90/gorpo/httpapi"
	"github.com/talento90/gorpo/image"
	"github.com/talento90/gorpo/log"
	"github.com/talento90/gorpo/repository/memory"
)

func main() {
	logConfig := log.Configuration{
		Level:  "debug",
		Output: os.Stdout,
	}

	logger, _ := log.NewLogger(logConfig)

	logger.Info("Starting gorpo API")

	httpDownloader := downloader.NewHTTPDownloader()

	effectRepo := memory.NewEffectRepository(httpDownloader)
	effectService := effect.NewService(effectRepo)
	imgService := image.NewService(httpDownloader, effectRepo)

	dependencies := &httpapi.ServerDependencies{
		Logger:        logger,
		Downloader:    httpDownloader,
		EffectService: effectService,
		ImgService:    imgService,
	}

	server := httpapi.NewServer(dependencies)

	http.ListenAndServe(server.Addr, server.Handler)
}
