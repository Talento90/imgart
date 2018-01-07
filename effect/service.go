package effect

import "github.com/talento90/gorpo/gorpo"

// NewService creates a service
func NewService(repository gorpo.EffectRepository) gorpo.EffectService {
	return &service{
		repository: repository,
	}
}

type service struct {
	repository Repository
}

func (es *service) GetEffects() ([]gorpo.Effect, error) {
	return es.repository.GetEffects()
}

func (es *service) GetEffect(id string) (gorpo.Effect, error) {
	return es.repository.GetEffect(id)
}
