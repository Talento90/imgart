package gorpo

import (
	"image"
)

// Filter represents an effect with all it's parameters
type Filter struct {
	ID         string                 `json:"id"`
	Parameters map[string]interface{} `json:"parameters"`
}

// EffectParameter contains all properties of a single effect parameter
type EffectParameter struct {
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Type        string      `json:"type"`
	Example     interface{} `json:"example"`
	Default     interface{} `json:"default,omitempty"`
	Values      interface{} `json:"values,omitempty"`
}

// EffectParameters it's a map that contains all parameters of an effect
type EffectParameters map[string]EffectParameter

// Effect represents an image transformation (ex: rotate, resize, overlay...)
type Effect interface {
	Descriptor() EffectDescriptor
	Transform(img image.Image, params map[string]interface{}) (image.Image, error)
}

// EffectDescriptor it's a struct that contains detailed information about an effect
type EffectDescriptor struct {
	ID          string           `json:"id"`
	Description string           `json:"description"`
	Parameters  EffectParameters `json:"parameters"`
}

type EffectRepository interface {
	GetEffects() ([]Effect, error)
	GetEffect(id string) (Effect, error)
}

type EffectService interface {
	GetEffects() ([]Effect, error)
	GetEffect(id string) (Effect, error)
}

type effectService struct {
	repository EffectRepository
}

func NewEffectService(repository EffectRepository) EffectService {
	return &effectService{
		repository: repository,
	}
}

func (es *effectService) GetEffects() ([]Effect, error) {
	return es.repository.GetEffects()
}

func (es *effectService) GetEffect(id string) (Effect, error) {
	return es.repository.GetEffect(id)
}
