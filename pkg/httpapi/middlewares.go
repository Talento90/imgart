package httpapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/errors"
	"github.com/talento90/gorpo/pkg/log"
)

type appHandler func(http.ResponseWriter, *http.Request, httprouter.Params) appResponse

func serializeResponse(r *http.Request, response *appResponse) (string, []byte) {
	const contentType = "application/json"

	bytes, err := json.Marshal(response.body)

	if err != nil {
		response.statusCode = http.StatusInternalServerError
		bytes, _ = json.Marshal(err)
	}

	return contentType, bytes
}

func loggerMiddleware(logger log.Logger, handler appHandler) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		start := time.Now()
		response := handler(w, r, params)

		if response.err != nil {

			if err, ok := response.err.(errors.Error); ok {
				logger.Error(err, err.Cause())
			} else {
				logger.Error(response.err)
			}
		}

		logger.InfoWithFields(
			log.Fields{
				"method":      r.Method,
				"url":         r.URL,
				"status_code": response.statusCode,
				"time":        time.Now().Sub(start),
			}, "api request")
	})
}

func responseMiddleware(handler appHandler) appHandler {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
		response := handler(w, r, params)

		if response.body == nil {
			return response
		}

		contentType, bytes := serializeResponse(r, &response)

		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(response.statusCode)
		w.Write(bytes)

		return response
	}
}
