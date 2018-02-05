package effect

import (
	"image"
	"testing"

	"github.com/talento90/imgart/errors"
)

func TestCropTransform(t *testing.T) {
	tt := []struct {
		name   string
		params map[string]interface{}
		err    errors.Type
	}{
		{
			name: "transform successfully",
			params: map[string]interface{}{
				"rectangle": []interface{}{100.0, 50.0, 200.0, 300.0},
			}},
		{
			name:   "missing rectangle",
			params: map[string]interface{}{},
			err:    errors.Validation,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			img := image.NewRGBA(image.Rect(0, 0, 100, 50))
			crop := NewCrop()

			_, err := crop.Transform(img, tc.params)

			if tc.err != "" {
				if err == nil || !errors.Is(tc.err, err) {
					t.Error("Expected validation error", err)
				}
			}
		})
	}
}
