package effect

import (
	"image"
	"testing"

	"github.com/talento90/gorpo/pkg/errors"
)

func TestBlurTransform(t *testing.T) {
	blur := NewBlur()
	params := map[string]interface{}{"sigma": 0.9}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := blur.Transform(img, params)

	if err != nil {
		t.Error("Should not return any error", err)
	}
}

func TestBlurTransformMissingSigma(t *testing.T) {
	blur := NewBlur()
	params := map[string]interface{}{}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := blur.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}
