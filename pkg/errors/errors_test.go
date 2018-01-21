package errors

import (
	"errors"
	"testing"
)

func TestIsMustCheckErrorType(t *testing.T) {
	err := EValidation("Validation error", nil)

	if !Is(Validation, err) {
		t.Error("Error should be validation", err)
	}
}

func TestErrorMessage(t *testing.T) {
	err := EValidation("Validation error", nil)

	if err.Error() != "Validation error" {
		t.Error("Error message is not correct", err)
	}
}

func TestErrorCauseMessage(t *testing.T) {
	errMsg := "Some error occured"
	err := EValidation("Validation error", errors.New(errMsg))

	e, ok := err.(*Error)

	if !ok || e.Cause() != errMsg {
		t.Error("Cause message is not correct", err)
	}
}
