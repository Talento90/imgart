package effect

import (
	"image"
	"testing"

	"github.com/talento90/gorpo/pkg/errors"
)

func TestRotateTransform(t *testing.T) {
	rotate := NewRotate()
	params := map[string]interface{}{"angle": 0.9, "bgcolor": "black"}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := rotate.Transform(img, params)

	if err != nil {
		t.Error("Should not return any error", err)
	}
}

func TestRotateTransformMissingParameters(t *testing.T) {
	rotate := NewRotate()
	params := map[string]interface{}{"bgcolor": "black"}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := rotate.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}
