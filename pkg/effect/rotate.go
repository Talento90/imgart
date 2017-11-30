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
				"teste": gorpo.EffectParameter{Value: 1, Required: true, Example: ""},
			},
		},
	}
}

func (r *rotate) Descriptor() gorpo.EffectDescriptor {
	return r.EffectDescriptor
}

func (r *rotate) Validate() []error {
	return nil
}

func (r *rotate) Transform(img image.Image, params gorpo.EffectParameters) (image.Image, error) {
	return img, nil
}
