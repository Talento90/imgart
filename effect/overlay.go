package effect

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/talento90/gorpo/downloader"
)

type overlay struct {
	downloader downloader.Downloader
	descriptor Descriptor
}

// NewOverlay creates an Effect that overlay an image with other image
func NewOverlay(downloader downloader.Downloader) Effect {
	return &overlay{
		downloader: downloader,
		descriptor: Descriptor{
			ID:          "overlay",
			Description: "Overlay - Overlay image",
			Parameters: Parameters{
				"position": Parameter{
					Description: "Position for the overlay image",
					Required:    true,
					Example:     "[1,2]",
					Type:        "array[int]",
				},
				"url": Parameter{
					Description: "Url for overlay image",
					Required:    true,
					Example:     "http://image.png",
					Type:        "string",
				},
				"opacity": Parameter{
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

func (o *overlay) Descriptor() Descriptor {
	return o.descriptor
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

	overlayImg, _, err := o.downloader.DownloadImage(url.String())

	if err != nil {
		return nil, err
	}

	img = imaging.Overlay(img, overlayImg, position, opacity)

	return img, nil
}
