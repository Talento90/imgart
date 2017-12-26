package effect

import (
	"image"

	"github.com/disintegration/imaging"
)

type crop struct {
	descriptor Descriptor
}

// NewCrop creates an Effect that crops the image
func NewCrop() Effect {
	return &crop{
		descriptor: Descriptor{
			ID:          "crop",
			Description: "Crop - Crops image",
			Parameters: Parameters{
				"rectangle": Parameter{
					Description: "Region to crop (x0,y0,x1,y1)",
					Required:    true,
					Example:     "[0,0,100,200]",
					Type:        "array[int]",
				},
			},
		},
	}
}

func (o *crop) Descriptor() Descriptor {
	return o.descriptor
}

func (o *crop) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	rectangle, err := rectangleBinder("rectangle", params)

	if err != nil {
		return nil, err
	}

	img = imaging.Crop(img, rectangle)

	return img, nil
}
