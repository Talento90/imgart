package effect

import (
	"image"

	"github.com/disintegration/imaging"
)

type contrast struct {
	descriptor Descriptor
}

// NewContrast creates an Effect changes the image contrast
func NewContrast() Effect {
	return &contrast{
		descriptor: Descriptor{
			ID:          "contrast",
			Description: "Contrast - Change the image contrast",
			Parameters: Parameters{
				"percentage": Parameter{
					Description: "Percentage of the contrast.",
					Required:    true,
					Example:     10,
					Type:        "float",
				},
			},
		},
	}
}

func (r *contrast) Descriptor() Descriptor {
	return r.descriptor
}

func (r *contrast) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	percentage, err := floatBinder("percentage", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustContrast(img, percentage)

	return img, nil
}
