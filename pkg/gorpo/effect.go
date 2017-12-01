package gorpo

import (
	"image"
)

type Filter struct {
	Id         string                 `json:"id"`
	Parameters map[string]interface{} `json:"parameters"`
}

type EffectParameter struct {
	Required bool
	Type     string
	Example  string
}

type EffectParameters map[string]EffectParameter

type Effect interface {
	Descriptor() EffectDescriptor
	Transform(img image.Image, params map[string]interface{}) (image.Image, error)
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
