package memory

import (
	"merlin/effects"
	"merlin/merlin"
)

type EffectRepository struct {
	effects []merlin.Effect
}

func NewEffectRepository() EffectRepository {
	return EffectRepository{
		effects: []merlin.Effect{effects.NewRotate()},
	}
}

func (r *EffectRepository) GetEffects() []merlin.Effect {
	return r.effects
}

func (r *EffectRepository) GetEffect(id string) merlin.Effect {
	for _, effect := range r.effects {
		if effect.Descriptor().Id == id {
			return effect
		}
	}

	return nil
}
