package httpapi

import (
	"net/http"

	"github.com/talento90/imgart/errors"
)

type appResponse struct {
	statusCode int
	body       interface{}
	err        error
}

type appError struct {
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
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
		ErrorType: errors.Internal.String(),
		Message:   err.Error(),
	}

	// Convert to Application Error if err is a Cancelled or DeadlineExceeded Error
	if ctxErr := errors.IsContextError(err); ctxErr != nil {
		err = ctxErr
	}

	if e, ok := err.(*errors.Error); ok {
		appError.ErrorType = e.Type.String()

		switch e.Type {
		case errors.NotFound:
			statusCode = http.StatusNotFound
		case errors.Validation:
			statusCode = http.StatusUnprocessableEntity
		case errors.Malformed:
			statusCode = http.StatusBadRequest
		case errors.AlreadyExists:
			statusCode = http.StatusConflict
		case errors.Timeout:
			statusCode = http.StatusRequestTimeout
		case errors.Cancelled:
			statusCode = http.StatusNoContent
		}
	}

	return appResponse{
		statusCode: statusCode,
		body:       appError,
		err:        err,
	}
}
