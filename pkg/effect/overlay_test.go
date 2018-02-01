package effect

import (
	"image"
	"testing"

	"github.com/talento90/imgart/pkg/errors"
	"github.com/talento90/imgart/pkg/mock"
)

func TestOverlayTransform(t *testing.T) {
	tt := []struct {
		name   string
		params map[string]interface{}
		err    errors.Type
	}{
		{
			name: "transform successfully",
			params: map[string]interface{}{
				"position": []interface{}{100.0, 50.0},
				"url":      "http://test.com/image.png",
				"opacity":  70,
			}},
		{
			name: "missing position",
			params: map[string]interface{}{
				"url":     "http://test.com/image.png",
				"opacity": 70,
			}},
		{
			name: "missing url",
			params: map[string]interface{}{
				"position": []interface{}{100.0, 50.0},
				"opacity":  70,
			}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			img := image.NewRGBA(image.Rect(0, 0, 100, 50))
			overlay := NewOverlay(mock.NewImageRepository())

			_, err := overlay.Transform(img, tc.params)

			if tc.err != "" {
				if err == nil || !errors.Is(tc.err, err) {
					t.Error("Expected validation error", err)
				}
			}
		})
	}
}
