package effect

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type rotate struct {
	gorpo.EffectDescriptor
}

func NewRotate() gorpo.Effect {
	return &rotate{
		EffectDescriptor: gorpo.EffectDescriptor{
			Id:          "rotate",
			Description: "This effect rotate an image",
			Parameters: gorpo.EffectParameters{
				"angle": gorpo.EffectParameter{
					Description: "Rotation angle in degreesº",
					Required:    true,
					Example:     "-90",
					Type:        "integer",
				},
				"bgcolor": gorpo.EffectParameter{
					Description: "Color of uncovered zones",
					Required:    false,
					Example:     "black",
					Type:        "string",
				},
			},
		},
	}
}

func (r *rotate) Descriptor() gorpo.EffectDescriptor {
	return r.EffectDescriptor
}

func (r *rotate) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	angle, _ := params["angle"]

	a, _ := angle.(float64)

	img = imaging.Rotate(img, a, color.Black)

	return img, nil
}
