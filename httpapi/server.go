package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/downloader"
	"github.com/talento90/gorpo/effect"
	"github.com/talento90/gorpo/image"
	"github.com/talento90/gorpo/log"
)

// ServerDependencies contains all dependencies
type ServerDependencies struct {
	Logger        log.Logger
	Downloader    downloader.Downloader
	EffectService effect.Service
	ImgService    image.Service
}

// NewServer creates an http server
func NewServer(dep *ServerDependencies) http.Server {
	router := httprouter.New()

	imgCtrl := newImagesController(dep.ImgService)
	effectCtrl := newEffectsController(dep.EffectService)

	router.GET("/api/v1/images", loggerMiddleware(dep.Logger, responseMiddleware(imgCtrl.ImageHandler)))

	router.GET("/api/v1/effects/:id", loggerMiddleware(dep.Logger, responseMiddleware(effectCtrl.GetEffectByID)))
	router.GET("/api/v1/effects", loggerMiddleware(dep.Logger, responseMiddleware(effectCtrl.GetAllEffects)))

	return http.Server{
		Addr:    ":4005",
		Handler: router,
	}
}
