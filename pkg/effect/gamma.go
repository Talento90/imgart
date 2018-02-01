package effect

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/talento90/imgart/pkg/imgart"
)

type gamma struct {
	effect
}

// NewGamma creates an Effect changes the image gamma
func NewGamma() imgart.Effect {
	return &gamma{
		effect: effect{
			id:          "gamma",
			description: "Gamma - Change the image gamma",
			parameters: imgart.Parameters{
				"gamma": imgart.Parameter{
					Description: "Percentage of the gamma correction.",
					Required:    true,
					Example:     0.75,
					Type:        "float",
				},
			},
		},
	}
}

func (r *gamma) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	gamma, err := floatBinder("gamma", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustGamma(img, gamma)

	return img, nil
}
