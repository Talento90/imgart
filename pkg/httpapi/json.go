package httpapi

import (
	"encoding/json"
	"net/http"
)

func toJSON(w http.ResponseWriter, body interface{}, statusCode int) {
	json, err := json.Marshal(body)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(json)
		w.WriteHeader(statusCode)
	}

	w.Header().Set("Content-Type", "image/application/json")
}
