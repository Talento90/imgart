package httpapi

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool        `json:"success,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewApiResponse(success bool, message string, data interface{}) ApiResponse {
	return ApiResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func toJSON(w http.ResponseWriter, statusCode int, response ApiResponse) {
	jsonBody, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		unexpectedError, _ := json.Marshal(NewApiResponse(false, "Unexpected error occured", nil))
		w.Write(unexpectedError)
	} else {
		w.WriteHeader(statusCode)
		w.Write(jsonBody)
	}

	w.Header().Set("Content-Type", "application/json")
}
