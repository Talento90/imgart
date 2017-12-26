package httpapi

import (
	"net/http"

	"github.com/talento90/gorpo/errors"
)

type appResponse struct {
	statusCode int
	body       interface{}
	err        error
}

type appError struct {
	ErrorCode errors.Type `json:"error_code"`
	ErrorType string      `json:"error_type"`
	Message   string      `json:"message"`
}

func response(statusCode int, body interface{}) appResponse {
	return appResponse{
		statusCode: statusCode,
		body:       body,
		err:        nil,
	}
}

func errResponse(err error) appResponse {
	statusCode := http.StatusInternalServerError
	appError := appError{
		ErrorCode: errors.Internal,
		ErrorType: errors.Internal.String(),
		Message:   err.Error(),
	}

	if e, ok := err.(*errors.Error); ok {
		appError.ErrorCode = e.ErrorType
		appError.ErrorType = e.ErrorType.String()

		switch e.ErrorType {
		case errors.NotExist:
			statusCode = http.StatusNotFound
		case errors.Validation:
			statusCode = http.StatusUnprocessableEntity
		case errors.Malformed:
			statusCode = http.StatusBadRequest
		}
	}

	return appResponse{
		statusCode: statusCode,
		body:       appError,
		err:        err,
	}
}
