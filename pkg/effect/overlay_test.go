package effect

import (
	"image"
	"testing"

	"github.com/talento90/gorpo/pkg/errors"
	"github.com/talento90/gorpo/pkg/mock"
)

func TestOverlayTransform(t *testing.T) {
	overlay := NewOverlay(mock.NewImageRepository())
	params := map[string]interface{}{
		"position": []interface{}{100.0, 50.0},
		"url":      "http://test.com/image.png",
		"opacity":  70,
	}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := overlay.Transform(img, params)

	if err != nil {
		t.Error("Should not return any error", err)
	}
}

func TestOverlayMissingPosition(t *testing.T) {
	overlay := NewOverlay(mock.NewImageRepository())
	params := map[string]interface{}{
		"url":     "http://test.com/image.png",
		"opacity": 70,
	}

	img := image.NewRGBA(image.Rect(0, 0, 100, 50))
	_, err := overlay.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}

func TestOverlayMissingUrl(t *testing.T) {
	overlay := NewOverlay(mock.NewImageRepository())
	params := map[string]interface{}{
		"position": []interface{}{100.0, 50.0},
		"opacity":  70,
	}

	img := image.NewRGBA(image.Rect(0, 0, 100, 50))
	_, err := overlay.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}
