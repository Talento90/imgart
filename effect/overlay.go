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
				"x": Parameter{
					Description: "Position x",
					Required:    true,
					Example:     0,
					Type:        "integer",
				},
				"y": Parameter{
					Description: "Position y",
					Required:    true,
					Example:     100,
					Type:        "integer",
				},
				"url": Parameter{
					Description: "Url for overlay image",
					Required:    true,
					Example:     "http://image.png",
					Type:        "string",
				},
			},
		},
	}
}

func (o *overlay) Descriptor() Descriptor {
	return o.descriptor
}

func (o *overlay) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	x, err := integerBinder("x", params)

	if err != nil {
		return nil, err
	}

	y, err := integerBinder("y", params)

	if err != nil {
		return nil, err
	}

	url, err := urlBinder("url", params)

	if err != nil {
		return nil, err
	}

	overlayImg, _, err := o.downloader.DownloadImage(url.String())

	if err != nil {
		return nil, err
	}

	img = imaging.Overlay(img, overlayImg, image.Pt(x, y), 100)

	return img, nil
}
