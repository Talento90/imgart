package effect

import (
	"testing"

	"github.com/talento90/imgart/pkg/errors"
)

func TestIntegerBinder(t *testing.T) {
	params := map[string]interface{}{"key": 1.0}
	_, err := integerBinder("key", params)

	if err != nil {
		t.Error("Should return an int", err)
	}
}

func TestIntegerBinderWrongInteger(t *testing.T) {
	params := map[string]interface{}{"key": "blackmaster"}
	_, err := integerBinder("key", params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}

func TestFloatBinder(t *testing.T) {
	params := map[string]interface{}{"key": 10.5}
	_, err := floatBinder("key", params)

	if err != nil {
		t.Error("Should return a float", err)
	}
}

func TestFloatBinderWrongFloat(t *testing.T) {
	params := map[string]interface{}{"key": "blackmaster"}
	_, err := floatBinder("key", params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}

func TestColorBinder(t *testing.T) {
	params := map[string]interface{}{"key": "black"}
	_, err := colorBinder("key", params)

	if err != nil {
		t.Error("Should return a color", err)
	}
}

func TestColorBinderWrongColor(t *testing.T) {
	params := map[string]interface{}{"key": "blackmaster"}
	_, err := colorBinder("key", params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}

func TestFilterBinder(t *testing.T) {
	params := map[string]interface{}{"key": "linear"}
	_, err := filterBinder("key", params)

	if err != nil {
		t.Error("Should return a filter", err)
	}
}

func TestFilterBinderWrongFilter(t *testing.T) {
	params := map[string]interface{}{"key": "linearus"}
	_, err := filterBinder("key", params)

	if err == nil {
		t.Error("Should return an error")
	}
}

func TestUrlBinder(t *testing.T) {
	params := map[string]interface{}{"key": "http://teste.com/image.jpeg"}
	_, err := urlBinder("key", params)

	if err != nil {
		t.Error("Should return an url", err)
	}
}

func TestUrlBinderWrongUrl(t *testing.T) {
	params := map[string]interface{}{"key": "fakeurl"}
	_, err := filterBinder("key", params)

	if err == nil {
		t.Error("Should return an error")
	}
}

func TestRectangleBinder(t *testing.T) {
	params := map[string]interface{}{"key": []interface{}{1.0, 2.0, 1.0, 2.0}}
	_, err := rectangleBinder("key", params)

	if err != nil {
		t.Error("Should return an url", err)
	}
}

func TestRectangleBinderWrongRectangle(t *testing.T) {
	params := map[string]interface{}{"key": []interface{}{1.0, 2.0, 1.0}}
	_, err := rectangleBinder("key", params)

	if err == nil {
		t.Error("Should return an error")
	}
}

func TestPointBinder(t *testing.T) {
	params := map[string]interface{}{"key": []interface{}{1.0, 2.0}}
	_, err := pointBinder("key", params)

	if err != nil {
		t.Error("Should return a point", err)
	}
}

func TestPointBinderWrongPoint(t *testing.T) {
	params := map[string]interface{}{"key": []interface{}{1.0, 2.0, 1.0}}
	_, err := pointBinder("key", params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}
