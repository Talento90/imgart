package effects

import (
	"image"

	"github.com/talento90/merlin/merlin"
)

type Rotate struct {
	descriptor merlin.EffectDescriptor
}

func NewRotate() Effect {
	return Rotate{
		descriptor: merlin.EffectDescriptor{
			Id:          "rotate",
			Description: "This effect rotate an image",
			Parameters: merlin.EffectParameters{
				"teste": merlin.EffectParameter{Value: 1, Required: true, Example: ""},
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
