package main

import (
	"net/http"

	"github.com/talento90/gorpo/pkg/downloader"
	"github.com/talento90/gorpo/pkg/gorpo"
	"github.com/talento90/gorpo/pkg/httpapi"
	"github.com/talento90/gorpo/pkg/logger"
	"github.com/talento90/gorpo/pkg/repository/memory"
)

func main() {
	logger := logger.NewLogger()

	logger.Println("Starting gorpo API")

	httpDownloader := downloader.NewHTTPDownloader()

	effectRepo := memory.NewEffectRepository()
	effectService := gorpo.NewEffectService(&effectRepo)

	server := httpapi.CreateServer(logger, httpDownloader, effectService)

	http.ListenAndServe(server.Addr, server.Handler)
}
