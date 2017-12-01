package effect

import (
	"image"

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
				"teste": gorpo.EffectParameter{Required: true, Example: "", Type: "integer"},
			},
		},
	}
}

func (r *rotate) Descriptor() gorpo.EffectDescriptor {
	return r.EffectDescriptor
}

func (r *rotate) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	return img, nil
}
