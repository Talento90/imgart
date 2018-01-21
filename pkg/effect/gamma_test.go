package effect

import (
	"image"
	"testing"

	"github.com/talento90/gorpo/pkg/errors"
)

func TestGammaTransform(t *testing.T) {
	gamma := NewGamma()
	params := map[string]interface{}{"gamma": 0.9}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := gamma.Transform(img, params)

	if err != nil {
		t.Error("Should not return any error", err)
	}
}

func TestGammaTransformMissingGamma(t *testing.T) {
	gamma := NewGamma()
	params := map[string]interface{}{"gammax": 0}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := gamma.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}
