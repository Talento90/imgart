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

// NewRotate creates an Effect that rotates an image
func NewRotate() gorpo.Effect {
	return &rotate{
		EffectDescriptor: gorpo.EffectDescriptor{
			Id:          "rotate",
			Description: "Rotate - rotates an image",
			Parameters: gorpo.EffectParameters{
				"angle": gorpo.EffectParameter{
					Description: "Rotation angle in degreesº",
					Required:    true,
					Example:     -90,
					Type:        "integer",
				},
				"bgcolor": gorpo.EffectParameter{
					Description: "Color of uncovered zones",
					Required:    false,
					Example:     "black",
					Type:        "string",
					Default:     "transparent",
					Values:      colorsList,
				},
			},
		},
	}
}

func (r *rotate) Descriptor() gorpo.EffectDescriptor {
	return r.EffectDescriptor
}

func (r *rotate) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	angle, err := floatBinder("angle", params)

	if err != nil {
		return nil, err
	}

	bgColor, err := colorBinder("bgcolor", params)

	if err != nil {
		bgColor = color.Transparent
	}

	img = imaging.Rotate(img, angle, bgColor)

	return img, nil
}
