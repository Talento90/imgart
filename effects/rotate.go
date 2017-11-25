package effects

import "image"

type Rotate struct {
	EffectDescriptor
}

func NewRotate() Effect {
	return Rotate{
		EffectDescriptor{
			Id:          "rotate",
			Description: "This effect rotate an image",
			Parameters: EffectParameters{
				"teste": EffectParameter{Value: 1, Required: true, Example: ""},
			},
		},
	}
}

func (r Rotate) Id() string {
	return r.EffectDescriptor.Id
}

func (r Rotate) Validate() error {
	return nil
}

func (r Rotate) Transform(img image.Image, params EffectParameters) (image.Image, error) {
	return img, nil
}
