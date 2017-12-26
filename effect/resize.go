package effect

import (
	"image"

	"github.com/disintegration/imaging"
)

type resize struct {
	descriptor Descriptor
}

// NewResize creates an Effect that resizes an image
func NewResize() Effect {
	return &resize{
		descriptor: Descriptor{
			ID:          "resize",
			Description: "Resize - resizes an image",
			Parameters: Parameters{
				"width": Parameter{
					Description: "Width in px",
					Required:    true,
					Example:     500,
					Type:        "integer",
				},
				"height": Parameter{
					Description: "Height in px",
					Required:    true,
					Example:     350,
					Type:        "integer",
				},
				"filter": Parameter{
					Description: "Resample filter",
					Required:    false,
					Example:     "linear",
					Type:        "string",
					Default:     "linear",
					Values:      filtersList,
				},
			},
		},
	}
}

func (r *resize) Descriptor() Descriptor {
	return r.descriptor
}

func (r *resize) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	width, err := integerBinder("width", params)

	if err != nil {
		return nil, err
	}

	height, err := integerBinder("height", params)

	if err != nil {
		return nil, err
	}

	filter, err := filterBinder("filter", params)

	if err != nil {
		filter = imaging.Linear
	}

	img = imaging.Resize(img, width, height, filter)

	return img, nil
}
