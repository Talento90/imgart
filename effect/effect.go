package effect

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

// Repository to store effects
type Repository interface {
	GetEffects() ([]Effect, error)
	GetEffect(id string) (Effect, error)
}

// Service interface with effect operations
type Service interface {
	// GetEffects return all available effects
	GetEffects() ([]Effect, error)
	// GetEffect returns an effect by the given id
	GetEffect(id string) (Effect, error)
}

// NewService creates a service
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

type service struct {
	repository Repository
}

func (es *service) GetEffects() ([]Effect, error) {
	return es.repository.GetEffects()
}

func (es *service) GetEffect(id string) (Effect, error) {
	return es.repository.GetEffect(id)
}
