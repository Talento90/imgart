package httpapi

import (
	"net/http"

	"github.com/talento90/gorpo/errors"
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
	if appErr, ok := err.(errors.Error); ok {

		if errors.IsNotExists(appErr) {
			return http.StatusNotFound
		}

		if errors.IsEValidation(appErr) {
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
