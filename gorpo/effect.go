package gorpo

import (
	"image"
)

// Filter represents an effect with all parameters
type Filter struct {
	// ID of the effect
	ID string `json:"id"`
	// Parameters to apply
	Parameters map[string]interface{} `json:"parameters"`
}

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
	// Descriptor returns the detailed description of the effect
	Descriptor() Descriptor
	// Transform applies the specific transformation to the given image
	Transform(img image.Image, params map[string]interface{}) (image.Image, error)
}

// Descriptor struct has a detailed description about the effect
type Descriptor struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
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

// NewEffectService creates effect service
func NewEffectService(repository EffectRepository) EffectService {
	return &effectService{
		repository: repository,
	}
}

type effectService struct {
	repository EffectRepository
}

func (es *effectService) GetEffects() ([]Effect, error) {
	return es.repository.GetEffects()
}

func (es *effectService) GetEffect(id string) (Effect, error) {
	return es.repository.GetEffect(id)
}
