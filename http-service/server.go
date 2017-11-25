package httpservice

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateServer() {
	router := httprouter.New()

	NewImagesController(router)
	NewEffectsControler(router)

	log.Fatal(http.ListenAndServe(":4005", router))
}
