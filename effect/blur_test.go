package effect

import (
	"image"
	"testing"

	"github.com/talento90/imgart/errors"
)

func TestBlurTransform(t *testing.T) {
	tt := []struct {
		name   string
		params map[string]interface{}
		err    errors.Type
	}{
		{
			name: "transform successfully",
			params: map[string]interface{}{
				"sigma": 0.9,
			},
		},
		{
			name:   "missing sigma",
			params: map[string]interface{}{},
			err:    errors.Validation,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			img := image.NewRGBA(image.Rect(0, 0, 100, 50))
			blur := NewBlur()

			_, err := blur.Transform(img, tc.params)

			if tc.err != "" {
				if err == nil || !errors.Is(tc.err, err) {
					t.Error("Expected validation error", err)
				}
			}
		})
	}
}
