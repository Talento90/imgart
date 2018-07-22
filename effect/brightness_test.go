package effect

import (
	"context"
	"image"
	"testing"

	"github.com/talento90/imgart/errors"
)

func TestBrightnessTransform(t *testing.T) {
	tt := []struct {
		name   string
		params map[string]interface{}
		err    errors.Type
	}{
		{
			name: "transform successfully",
			params: map[string]interface{}{
				"percentage": 0.5,
			},
		},
		{
			name:   "missing percentage",
			params: map[string]interface{}{},
			err:    errors.Validation,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			img := image.NewRGBA(image.Rect(0, 0, 100, 50))
			brightness := NewBrightness()

			_, err := brightness.Transform(context.Background(), img, tc.params)

			if tc.err != "" {
				if err == nil || !errors.Is(tc.err, err) {
					t.Error("Expected validation error", err)
				}
			}
		})
	}
}
