package effect

import (
	"context"
	"image"

	"github.com/disintegration/imaging"
	"github.com/talento90/imgart/imgart"
)

type brightness struct {
	effect
}

// NewBrightness creates an Effect changes the image brightnness
func NewBrightness() imgart.Effect {
	return &brightness{
		effect: effect{
			id:          "brightness",
			description: "Brightness - Change the image brightness",
			parameters: imgart.Parameters{
				"percentage": imgart.Parameter{
					Description: "Percentage of the brightness.",
					Required:    true,
					Example:     0.5,
					Type:        "float",
				},
			},
		},
	}
}

func (r *brightness) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	percentage, err := floatBinder("percentage", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustBrightness(img, percentage)

	return img, ctx.Err()
}
