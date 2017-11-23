package gomage

import "image"

type Rotate struct {
	Id          string           `json:"id"`
	Description string           `json:"description"`
	Parameters  EffectParameters `json:"parameters"`
}

func NewRotate() Effect {
	return Rotate{
		Id:          "rotate",
		Description: "This effect rotate an image",
		Parameters: EffectParameters{
			"teste": EffectParameter{Value: 1, Required: true, Example: ""},
		},
	}
}

func (r Rotate) Validate() error {
	return nil
}

func (r Rotate) Transform(img image.Image, params EffectParameters) (image.Image, error) {
	return img, nil
}
