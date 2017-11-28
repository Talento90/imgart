package gorpo

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

type EffectRepository interface {
	GetEffects() ([]Effect, error)
	GetEffect(id string) (Effect, error)
}

type EffectService interface {
	GetEffects() ([]Effect, error)
	GetEffect(id string) (Effect, error)
}

type effectService struct {
	repository *EffectRepository
}

func NewEffectService(repository *EffectRepository) EffectService {
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
