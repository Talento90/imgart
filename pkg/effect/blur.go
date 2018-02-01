package effect

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/talento90/imgart/pkg/imgart"
)

type blur struct {
	effect
}

// NewBlur creates an Effect that blurrs an image
func NewBlur() imgart.Effect {
	return &blur{
		effect: effect{
			id:          "blur",
			description: "Blur - Gaussian Blur",
			parameters: imgart.Parameters{
				"sigma": imgart.Parameter{
					Description: "How much the image will be blurred.",
					Required:    true,
					Example:     0.5,
					Type:        "float",
				},
			},
		},
	}
}

func (r *blur) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	sigma, err := floatBinder("sigma", params)

	if err != nil {
		return nil, err
	}

	img = imaging.Blur(img, sigma)

	return img, nil
}
