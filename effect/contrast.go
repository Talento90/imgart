package effect

import (
	"context"
	"image"

	"github.com/disintegration/imaging"
	"github.com/talento90/imgart/imgart"
)

type contrast struct {
	effect
}

// NewContrast creates an Effect changes the image contrast
func NewContrast() imgart.Effect {
	return &contrast{
		effect: effect{
			id:          "contrast",
			description: "Contrast - Change the image contrast",
			parameters: imgart.Parameters{
				"percentage": imgart.Parameter{
					Description: "Percentage of the contrast.",
					Required:    true,
					Example:     10,
					Type:        "float",
				},
			},
		},
	}
}

func (r *contrast) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	percentage, err := floatBinder("percentage", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustContrast(img, percentage)

	return img, ctx.Err()
}
