package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type appHandler func(http.ResponseWriter, *http.Request, httprouter.Params) appResponse

func serializeResponse(r *http.Request, response *appResponse) (string, []byte) {
	const contentType = "application/json"

	bytes, err := json.Marshal(response.Body)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		bytes, _ = json.Marshal(err)
	}

	return contentType, bytes
}

func loggerMiddleware(logger *log.Logger, handler appHandler) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		start := time.Now()
		response := handler(w, r, params)

		logger.Printf("%s %s %d %s\n", r.Method, r.URL, response.StatusCode, time.Now().Sub(start))
	})
}

func responseMiddleware(handler appHandler) appHandler {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
		response := handler(w, r, params)

		contentType, bytes := serializeResponse(r, &response)

		// w.Header().Add("X-Count", strconv.Itoa(len(effects)))

		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(response.StatusCode)
		w.Write(bytes)

		return response
	}
}
