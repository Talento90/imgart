package effect

import (
	"image"

	"github.com/talento90/gorpo/pkg/gorpo"
)

type Rotate struct {
	descriptor gorpo.EffectDescriptor
}

func NewRotate() Effect {
	return Rotate{
		descriptor: gorpo.EffectDescriptor{
			Id:          "rotate",
			Description: "This effect rotate an image",
			Parameters: gorpo.EffectParameters{
				"teste": gorpo.EffectParameter{Value: 1, Required: true, Example: ""},
			},
		},
	}
}

func (r Rotate) EffectDescriptor() EffectDescriptor {
	return r.descriptor
}

func (r Rotate) Validate() error {
	return nil
}

func (r Rotate) Transform(img image.Image, params EffectParameters) (image.Image, error) {
	return img, nil
}
