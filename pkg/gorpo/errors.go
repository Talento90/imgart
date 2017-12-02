package gorpo

import "fmt"

type Error struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Errors  []error `json:"errors,omitempty"`
}

func NewError(code int, message string, errors ...error) Error {
	return Error{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("ErrorCode: %d Message: %s", e.Code, e.Message)
}

var downloadError = NewError(100, "Error downloading image")
var processImageError = NewError(100, "Error processing image")
