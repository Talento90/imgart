package httpapi

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type Dependencies struct {
	log.Logger
	gorpo.Downloader
	gorpo.EffectService
	gorpo.ImageService
}

// CreateServer creates an http server
func CreateServer(logger *log.Logger, downloader gorpo.Downloader, effectService gorpo.EffectService, imgService gorpo.ImageService) http.Server {
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
