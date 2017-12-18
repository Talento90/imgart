package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/talento90/gorpo/pkg/gorpo"

	"github.com/julienschmidt/httprouter"
)

type appResponse struct {
	StatusCode int
	Body       interface{}
}

func response(statusCode int, body interface{}) appResponse {
	return appResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}

func getHTTPError(err error) int {
	if appErr, ok := err.(gorpo.Error); ok {

		if gorpo.IsNotExists(appErr) {
			return http.StatusNotFound
		}

		if gorpo.IsEValidation(appErr) {
			return http.StatusUnprocessableEntity
		}
	}

	return http.StatusInternalServerError
}

func errResponse(err error) appResponse {
	return appResponse{
		StatusCode: getHTTPError(err),
		Body:       err,
	}
}

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

func loggerMiddleware(logger *log.Logger) func(handler httprouter.Handle) httprouter.Handle {
	return func(handler httprouter.Handle) httprouter.Handle {
		return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			start := time.Now()

			handler(w, r, params)

			logger.Printf("%s %s %s %s\n", r.Method, r.RemoteAddr, r.URL, time.Now().Sub(start))
		})
	}
}

func responseMiddleware(handler appHandler) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		response := handler(w, r, params)

		if response.Body == nil {
			return
		}

		contentType, bytes := serializeResponse(r, &response)

		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(response.StatusCode)
		w.Write(bytes)
	})
}
