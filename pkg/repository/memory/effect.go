package memory

import (
	"github.com/talento90/merlin/pkg/merlin"
)

type EffectRepository struct {
	effects []merlin.Effect
}

func NewEffectRepository() EffectRepository {
	return EffectRepository{
		effects: []merlin.Effect{effects.NewRotate()},
	}
}

func (r *EffectRepository) GetEffects() ([]merlin.Effect, error) {
	return r.effects, nil
}

func (r *EffectRepository) GetEffect(id string) (merlin.Effect, error) {
	for _, effect := range r.effects {
		if effect.Descriptor().Id == id {
			return effect, nil
		}
	}

	return nil, nil
}
