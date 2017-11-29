package httpapi

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

func CreateServer(downloader gorpo.Downloader, effectService gorpo.EffectService) Server {
	router := httprouter.New()

	imgCtrl := NewImagesController(downloader)
	effectCtrl := NewEffectsController(effectService)

	router.GET("/api/v1/images", imgCtrl.ImageHandler)

	router.GET("/api/v1/effects/:id", effectCtrl.)
	router.GET("/api/v1/effects", effectCtrl.)

	return http.Server{
		Addr: ":4005",
		Handler: router
	} 
}
