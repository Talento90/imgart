package effect

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/talento90/gorpo/gorpo"
)

type overlay struct {
	imgRepository gorpo.ImageRepository
	effect
}

// NewOverlay creates an Effect that overlay an image with other image
func NewOverlay(imgRepository gorpo.ImageRepository) gorpo.Effect {
	return &overlay{
		imgRepository: imgRepository,
		effect: effect{
			id:          "overlay",
			description: "Overlay - Overlay image",
			parameters: gorpo.Parameters{
				"position": gorpo.Parameter{
					Description: "Position for the overlay image",
					Required:    true,
					Example:     "[1,2]",
					Type:        "array[int]",
				},
				"url": gorpo.Parameter{
					Description: "Url for overlay image",
					Required:    true,
					Example:     "http://image.png",
					Type:        "string",
				},
				"opacity": gorpo.Parameter{
					Description: "Opacity for overlay image",
					Required:    false,
					Example:     90,
					Type:        "float",
					Default:     100,
				},
			},
		},
	}
}

func (o *overlay) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	position, err := pointBinder("position", params)

	if err != nil {
		return nil, err
	}

	url, err := urlBinder("url", params)

	if err != nil {
		return nil, err
	}

	opacity, err := floatBinder("opacity", params)

	if err != nil {
		opacity = 100
	}

	overlayImg, _, err := o.imgRepository.Get(url.String())

	if err != nil {
		return nil, err
	}

	img = imaging.Overlay(img, overlayImg, position, opacity)

	return img, nil
}
