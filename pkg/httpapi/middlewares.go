package httpapi

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func logRequest(logger *log.Logger, handler httprouter.Handle) httprouter.Handle {

	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		start := time.Now()

		handler(w, r, params)

		logger.Printf("%s %s %s %s ms\n", r.Method, r.RemoteAddr, r.URL, time.Now().Sub(start))
	})
}
