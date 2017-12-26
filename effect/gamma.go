package effect

import (
	"image"

	"github.com/disintegration/imaging"
)

type gamma struct {
	descriptor Descriptor
}

// NewGamma creates an Effect changes the image gamma
func NewGamma() Effect {
	return &gamma{
		descriptor: Descriptor{
			ID:          "gamma",
			Description: "Gamma - Change the image gamma",
			Parameters: Parameters{
				"gamma": Parameter{
					Description: "Percentage of the gamma correction.",
					Required:    true,
					Example:     0.75,
					Type:        "float",
				},
			},
		},
	}
}

func (r *gamma) Descriptor() Descriptor {
	return r.descriptor
}

func (r *gamma) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	gamma, err := floatBinder("gamma", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustGamma(img, gamma)

	return img, nil
}
