package effect

import (
	"context"
	"image"
	"testing"

	"github.com/talento90/imgart/errors"
)

func TestResizeTransform(t *testing.T) {
	tt := []struct {
		name   string
		params map[string]interface{}
		err    errors.Type
	}{
		{
			name: "transform successfully",
			params: map[string]interface{}{
				"width":  200.0,
				"height": 300.0,
				"filter": "lanczos",
			}},
		{
			name: "missing width",
			params: map[string]interface{}{
				"height": 300.0,
				"filter": "lanczos",
			}},
		{
			name: "missing height",
			params: map[string]interface{}{
				"width":  200.0,
				"filter": "lanczos",
			}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			img := image.NewRGBA(image.Rect(0, 0, 100, 50))
			resize := NewResize()

			_, err := resize.Transform(context.Background(), img, tc.params)

			if tc.err != "" {
				if err == nil || !errors.Is(tc.err, err) {
					t.Error("Expected validation error", err)
				}
			}
		})
	}
}
