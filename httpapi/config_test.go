package httpapi

import (
	"testing"
	"time"
)

func TestConfiguration(t *testing.T) {
	c := Configuration{
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Second,
	}

	if err := c.Validate(); err == nil {
		t.Errorf("Expected validation error and got: %s", err)
	}
}
