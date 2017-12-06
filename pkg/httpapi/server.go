package httpapi

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

func CreateServer(logger *log.Logger, downloader gorpo.Downloader, effectService gorpo.EffectService, imgService gorpo.ImageService) http.Server {
	router := httprouter.New()

	imgCtrl := newImagesController(imgService)
	effectCtrl := newEffectsController(effectService)

	loggerHandler := logHandler(logger)

	router.GET("/api/v1/images", loggerHandler(responseHandler(imgCtrl.ImageHandler)))

	router.GET("/api/v1/effects/:id", loggerHandler(responseHandler(effectCtrl.GetEffectById)))
	router.GET("/api/v1/effects", loggerHandler(responseHandler(effectCtrl.GetAllEffects)))

	return http.Server{
		Addr:    ":4005",
		Handler: router,
	}
}
