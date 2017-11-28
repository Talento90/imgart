package httpapi

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/gorpo"
)

func CreateServer(effectService gorpo.EffectService) {
	router := httprouter.New()

	NewImagesController(router)
	NewEffectsControler(router)

	log.Fatal(http.ListenAndServe(":4005", router))
}
