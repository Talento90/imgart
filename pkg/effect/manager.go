package effects

import (
	"github.com/talento90/merlin/merlin"
)

type EffectManager struct {
	effects []merlin.Effect
}

func NewEffectManager() EffectManager {
	return EffectManager{
		effects: []merlin.Effect{NewRotate()},
	}
}

func (m *EffectManager) GetEffects() []merlin.Effect {
	return m.effects
}

func (m *EffectManager) GetEffect(id string) merlin.Effect {
	for _, effect := range m.effects {
		if effect.Descriptor().Id == id {
			return effect
		}
	}

	return nil
}
