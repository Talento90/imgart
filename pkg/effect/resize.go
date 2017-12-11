package effect

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type resize struct {
	gorpo.EffectDescriptor
}

func NewResize() gorpo.Effect {
	return &resize{
		EffectDescriptor: gorpo.EffectDescriptor{
			Id:          "resize",
			Description: "Resize - resizes an image",
			Parameters: gorpo.EffectParameters{
				"width": gorpo.EffectParameter{
					Description: "Width in px",
					Required:    true,
					Example:     500,
					Type:        "integer",
				},
				"height": gorpo.EffectParameter{
					Description: "Height in px",
					Required:    true,
					Example:     350,
					Type:        "integer",
				},
				"filter": gorpo.EffectParameter{
					Description: "Resample filter",
					Required:    false,
					Example:     "linear",
					Type:        "string",
					Default:     "linear",
					Values:      "lanczos,catmull-rom,mitchell-netravali,bs-pline,linear,box,nearest-neighbor",
				},
			},
		},
	}
}

func (r *resize) Descriptor() gorpo.EffectDescriptor {
	return r.EffectDescriptor
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
