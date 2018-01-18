package effect

import (
	"image"
	"testing"

	"github.com/talento90/gorpo/pkg/errors"
)

func TestResizeTransform(t *testing.T) {
	resize := NewResize()
	params := map[string]interface{}{"width": 200, "height": 300, "filter": "lanczos"}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := resize.Transform(img, params)

	if err != nil {
		t.Error("Should not return any error", err)
	}
}

func TestResizeTransformMissingWidth(t *testing.T) {
	resize := NewResize()
	params := map[string]interface{}{"height": 300, "filter": "lanczos"}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := resize.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Error should be a validation error", err)
	}
}

func TestResizeTransformMissingHeight(t *testing.T) {
	resize := NewResize()
	params := map[string]interface{}{"width": 200, "filter": "lanczos"}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := resize.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Error should be a validation error", err)
	}
}
