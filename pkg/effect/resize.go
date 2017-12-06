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
			Description: "This effect resizes an image",
			Parameters: gorpo.EffectParameters{
				"width": gorpo.EffectParameter{
					Description: "Width in px",
					Required:    true,
					Example:     "500",
					Type:        "integer",
				},
				"height": gorpo.EffectParameter{
					Description: "Height in px",
					Required:    true,
					Example:     "350",
					Type:        "integer",
				},
				"filter": gorpo.EffectParameter{
					Description: "Resample filter",
					Required:    false,
					Example:     "black",
					Type:        "string",
				},
			},
		},
	}
}

func (r *resize) Descriptor() gorpo.EffectDescriptor {
	return r.EffectDescriptor
}

func (r *resize) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	width, _ := params["width"]
	height, _ := params["height"]

	w, _ := width.(int)
	h, _ := height.(int)

	img = imaging.Resize(img, w, h, imaging.Linear)

	return img, nil
}
