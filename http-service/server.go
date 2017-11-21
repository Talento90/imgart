package httpservice

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateServer() {
	router := httprouter.New()

	RegisterImagesController(router)

	log.Fatal(http.ListenAndServe(":4005", router))
}
