package httpapi

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type AppResponse struct {
	Code int
	Body interface{}
}

func Response(code int, body interface{}) AppResponse {
	return AppResponse{
		Code: code,
		Body: body,
	}
}

type AppHandle func(http.ResponseWriter, *http.Request, httprouter.Params) AppResponse

func SerializeBody(r *http.Request, response *AppResponse) (string, []byte) {
	var err error
	var bytes []byte
	var contentType = "application/json"

	if accept := r.Header.Get("Accept"); strings.Contains(accept, "application/xml") {
		bytes, err = xml.Marshal(response.Body)
		contentType = "application/xml"
	} else {
		bytes, err = json.Marshal(response.Body)
	}

	if err != nil {
		response.Code = http.StatusInternalServerError
		bytes, _ = json.Marshal(errors.New(""))
	}

	return contentType, bytes
}

func LogHandler(logger *log.Logger, handler httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		start := time.Now()

		handler(w, r, params)

		logger.Printf("%s %s %s %s ms\n", r.Method, r.RemoteAddr, r.URL, time.Now().Sub(start))
	})
}

func ResponseHandler(handler AppHandle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		response := handler(w, r, params)

		contentType, bytes := SerializeBody(r, &response)

		w.WriteHeader(response.Code)
		w.Header().Set("Content-Type", contentType)
		w.Write(bytes)
	})
}
