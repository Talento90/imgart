package merlin

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
	Descriptor() EffectDescriptor

	Validate() []error

	Transform(img image.Image, params EffectParameters) (image.Image, error)
}

type EffectDescriptor struct {
	Id          string           `json:"id"`
	Description string           `json:"description"`
	Parameters  EffectParameters `json:"parameters"`
}

type EffectManager interface {
	GetEffects() []Effect
	GetEffect(id string) Effect
}
