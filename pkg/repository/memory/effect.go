package memory

import (
	"github.com/talento90/gorpo/pkg/gorpo"
)

type EffectRepository struct {
	effects []gorpo.Effect
}

func NewEffectRepository() EffectRepository {
	return EffectRepository{
		effects: []gorpo.Effect{effects.NewRotate()},
	}
}

func (r *EffectRepository) GetEffects() ([]gorpo.Effect, error) {
	return r.effects, nil
}

func (r *EffectRepository) GetEffect(id string) (gorpo.Effect, error) {
	for _, effect := range r.effects {
		if effect.Descriptor().Id == id {
			return effect, nil
		}
	}

	return nil, nil
}
