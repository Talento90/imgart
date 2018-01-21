package effect

import (
	"image"
	"testing"

	"github.com/talento90/gorpo/pkg/errors"
)

func TestCropTransform(t *testing.T) {
	crop := NewCrop()
	params := map[string]interface{}{"rectangle": []interface{}{100.0, 50.0, 200.0, 300.0}}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := crop.Transform(img, params)

	if err != nil {
		t.Error("Should not return any error", err)
	}
}

func TestCropTransformMissingRectangle(t *testing.T) {
	crop := NewCrop()
	params := map[string]interface{}{}
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	_, err := crop.Transform(img, params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}
