package effect

import (
	"testing"
)

func TestIntegerBinder(t *testing.T) {
	params := map[string]interface{}{"key": 1}
	_, err := integerBinder("key", params)

	if err != nil {
		t.Error("Should return an int", err)
	}
}
