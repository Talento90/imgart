package effects

import (
	"github.com/talento90/go-mage"
)

type EffectManager struct {
	effects []Effect
}

func NewEffectManager() EffectManager {
	return EffectManager{
		effects: []Effect{},
	}
}

func (m *EffectManager) GetEffects() []Effect {
	return m.effects
}

func (m *EffectManager) GetEffect(id string) Effect {

	for _, effect := m.effects {
		if (effect.Id == id)
			return effect
	}

	return nil
}
