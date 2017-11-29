package httpapi

import (
	"encoding/json"
	"net/http"
)

func toJSON(w http.ResponseWriter, body interface{}, statusCode int) {
	json, err := json.Marshal(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(statusCode)
		w.Write(json)
	}

	w.Header().Set("Content-Type", "application/json")
}
