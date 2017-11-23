package gomage

import (
	"image"
)

type EffectParameter struct {
	Value    interface{}
	Required bool
	Example  string
}

type EffectParameters map[string]EffectParameter

type Effect interface {
	Validate() error

	Transform(img image.Image, params EffectParameters) (image.Image, error)
}
