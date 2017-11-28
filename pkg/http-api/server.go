package httpapi

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

func CreateServer(effectService gorpo.EffectService) Server {
	router := httprouter.New()

	NewImagesController(router)
	NewEffectsControler(effectService)


	return http.Server{
		Addr: ":4005",
		Handler: router
	} 
}
