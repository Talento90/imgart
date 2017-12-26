package effect

import (
	"image"

	"github.com/disintegration/imaging"
)

type brightness struct {
	descriptor Descriptor
}

// NewBrightness creates an Effect changes the image brightnness
func NewBrightness() Effect {
	return &brightness{
		descriptor: Descriptor{
			ID:          "brightness",
			Description: "Brightness - Change the image brightness",
			Parameters: Parameters{
				"percentage": Parameter{
					Description: "Percentage of the brightness.",
					Required:    true,
					Example:     0.5,
					Type:        "float",
				},
			},
		},
	}
}

func (r *brightness) Descriptor() Descriptor {
	return r.descriptor
}

func (r *brightness) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	percentage, err := floatBinder("percentage", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustBrightness(img, percentage)

	return img, nil
}
