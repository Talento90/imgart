package effect

import (
	"context"
	"image"

	"github.com/disintegration/imaging"
	"github.com/talento90/imgart/imgart"
)

type resize struct {
	effect
}

// NewResize creates an Effect that resizes an image
func NewResize() imgart.Effect {
	return &resize{
		effect: effect{
			id:          "resize",
			description: "Resize - resizes an image",
			parameters: imgart.Parameters{
				"width": imgart.Parameter{
					Description: "Width in px",
					Required:    true,
					Example:     500,
					Type:        "integer",
				},
				"height": imgart.Parameter{
					Description: "Height in px",
					Required:    true,
					Example:     350,
					Type:        "integer",
				},
				"filter": imgart.Parameter{
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

func (r *resize) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
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

	return img, ctx.Err()
}
