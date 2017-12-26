package effect

import (
	"image"

	"github.com/disintegration/imaging"
)

type blur struct {
	descriptor Descriptor
}

// NewBlur creates an Effect that blurrs an image
func NewBlur() Effect {
	return &blur{
		descriptor: Descriptor{
			ID:          "blur",
			Description: "Blur - Gaussian Blur",
			Parameters: Parameters{
				"sigma": Parameter{
					Description: "How much the image will be blurred.",
					Required:    true,
					Example:     0.5,
					Type:        "float",
				},
			},
		},
	}
}

func (r *blur) Descriptor() Descriptor {
	return r.descriptor
}

func (r *blur) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	sigma, err := floatBinder("sigma", params)

	if err != nil {
		return nil, err
	}

	img = imaging.Blur(img, sigma)

	return img, nil
}
