package gorpo

import "fmt"

type Error struct {
	errorType   string
	originalErr error
	Message     string  `json:"message"`
	Errors      []Error `json:"errors,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Message: %s", e.Message)
}

const validationType = "validation"
const notExistsType = "notExists"

func EValidation(msg string) error {
	return &Error{
		errorType: validationType,
		Message:   msg,
	}
}

func ENotExists(msg string) error {
	return &Error{
		errorType: notExistsType,
		Message:   msg,
	}
}

func EProcessing(msg string, err error) error {
	return &Error{
		errorType:   notExistsType,
		originalErr: err,
		Message:     msg,
	}
}

func IsEValidation(err error) bool {
	if e, ok := err.(Error); ok {
		return e.errorType == validationType
	}

	return false
}
