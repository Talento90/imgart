package effect

import (
	"image"
	"testing"

	"github.com/talento90/gorpo/pkg/errors"
)

func TestContrastTransform(t *testing.T) {
	constrast := NewContrast()
	params := map[string]interface{}{"percentage": 20.0}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := constrast.Transform(img, params)

	if err != nil {
		t.Error("Should not return any error", err)
	}
}

func TestContrastTransformMissingPercentage(t *testing.T) {
	constrast := NewContrast()
	params := map[string]interface{}{}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := constrast.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}
