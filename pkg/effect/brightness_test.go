package effect

import (
	"image"
	"testing"

	"github.com/talento90/gorpo/pkg/errors"
)

func TestBrightnessTransform(t *testing.T) {
	brightness := NewBrightness()
	params := map[string]interface{}{"percentage": 0.5}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := brightness.Transform(img, params)

	if err != nil {
		t.Error("Should not return any error", err)
	}
}

func TestBrightnessTransformMissingPercentage(t *testing.T) {
	brightness := NewBrightness()
	params := map[string]interface{}{}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := brightness.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}
