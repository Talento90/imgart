package httpapi

import (
	"encoding/json"
	"net/http"
)

func toJSON(w http.ResponseWrite, statusCode int) {

	w.Header().Set("Content-Type", "image/application/json")
	w.WriteHeader(statusCode)

	json, err := json.Marshal(gomage.GetAllEffects)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.write(err)
	} else {
		w.Write(json)
	}
}
