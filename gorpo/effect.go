package gorpo

import (
	"image"
)

// Parameter contains all properties of a single effect parameter
type Parameter struct {
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Type        string      `json:"type"`
	Example     interface{} `json:"example"`
	Default     interface{} `json:"default,omitempty"`
	Values      interface{} `json:"values,omitempty"`
}

// Parameters it's a map that contains all parameters of an effect
type Parameters map[string]Parameter

// Effect represents an image transformation (ex: rotate, resize, overlay...)
type Effect interface {
	ID() string

	Description() string

	Parameters() Parameters

	// Transform applies the specific transformation to the given image
	Transform(img image.Image, params map[string]interface{}) (image.Image, error)
}

// EffectRepository to store effects
type EffectRepository interface {
	GetEffects() ([]Effect, error)
	GetEffect(id string) (Effect, error)
}

// EffectService interface with effect operations
type EffectService interface {
	// GetEffects return all available effects
	GetEffects() ([]Effect, error)
	// GetEffect returns an effect by the given id
	GetEffect(id string) (Effect, error)
}
