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

	loggerHandler := loggerMiddleware(logger)

	router.GET("/api/v1/images", loggerHandler(responseMiddleware(imgCtrl.ImageHandler)))

	router.GET("/api/v1/effects/:id", loggerHandler(responseMiddleware(effectCtrl.GetEffectByID)))
	router.GET("/api/v1/effects", loggerHandler(responseMiddleware(effectCtrl.GetAllEffects)))

	return http.Server{
		Addr:    ":4005",
		Handler: router,
	}
}
