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

	router.GET("/api/v1/images", logRequest(logger, imgCtrl.ImageHandler))

	router.GET("/api/v1/effects/:id", effectCtrl.GetEffectById)
	router.GET("/api/v1/effects", effectCtrl.GetAllEffects)

	return http.Server{
		Addr:    ":4005",
		Handler: router,
	}
}
