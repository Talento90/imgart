package httpapi

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

func CreateServer(logger *log.Logger, downloader gorpo.Downloader, effectService gorpo.EffectService) http.Server {
	router := httprouter.New()

	imgCtrl := newImagesController(downloader)
	effectCtrl := newEffectsController(effectService)

	router.GET("/api/v1/images", LogHandler(logger, imgCtrl.ImageHandler))

	router.GET("/api/v1/effects/:id", ResponseHandler(effectCtrl.GetEffectById))
	router.GET("/api/v1/effects", ResponseHandler(effectCtrl.GetAllEffects))

	return http.Server{
		Addr:    ":4005",
		Handler: router,
	}
}
