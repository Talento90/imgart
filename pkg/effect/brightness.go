package effect

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type brightness struct {
	effect
}

// NewBrightness creates an Effect changes the image brightnness
func NewBrightness() gorpo.Effect {
	return &brightness{
		effect: effect{
			id:          "brightness",
			description: "Brightness - Change the image brightness",
			parameters: gorpo.Parameters{
				"percentage": gorpo.Parameter{
					Description: "Percentage of the brightness.",
					Required:    true,
					Example:     0.5,
					Type:        "float",
				},
			},
		},
	}
}

func (r *brightness) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	percentage, err := floatBinder("percentage", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustBrightness(img, percentage)

	return img, nil
}
