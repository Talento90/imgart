package main

import (
	"net/http"
	"os"

	"github.com/talento90/gorpo/pkg/downloader"
	"github.com/talento90/gorpo/pkg/gorpo"
	"github.com/talento90/gorpo/pkg/httpapi"
	"github.com/talento90/gorpo/pkg/log"
	"github.com/talento90/gorpo/pkg/repository/memory"
)

func main() {
	logConfig := log.Configuration{
		Level:  "debug",
		Output: os.Stdout,
	}

	logger, _ := log.NewLogger(logConfig)

	logger.Info("Starting gorpo API")

	httpDownloader := downloader.NewHTTPDownloader()

	effectRepo := memory.NewEffectRepository()
	effectService := gorpo.NewEffectService(effectRepo)
	imgService := gorpo.NewImageService(httpDownloader, effectRepo)

	server := httpapi.CreateServer(logger, httpDownloader, effectService, imgService)

	http.ListenAndServe(server.Addr, server.Handler)
}
