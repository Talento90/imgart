package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/downloader"
	"github.com/talento90/gorpo/effect"
	"github.com/talento90/gorpo/image"
	"github.com/talento90/gorpo/log"
)

// CreateServer creates an http server
func CreateServer(logger log.Logger, downloader downloader.Downloader, effectService effect.Service, imgService image.Service) http.Server {
	router := httprouter.New()

	imgCtrl := newImagesController(imgService)
	effectCtrl := newEffectsController(effectService)

	router.GET("/api/v1/images", loggerMiddleware(logger, responseMiddleware(imgCtrl.ImageHandler)))

	router.GET("/api/v1/effects/:id", loggerMiddleware(logger, responseMiddleware(effectCtrl.GetEffectByID)))
	router.GET("/api/v1/effects", loggerMiddleware(logger, responseMiddleware(effectCtrl.GetAllEffects)))

	return http.Server{
		Addr:    ":4005",
		Handler: router,
	}
}
