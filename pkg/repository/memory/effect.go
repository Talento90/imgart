package memory

import (
	"fmt"

	"github.com/talento90/gorpo/pkg/effect"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type effectRepository struct {
	effects []gorpo.Effect
}

// NewEffectRepository creates a memory repository for Effect entity
func NewEffectRepository() gorpo.EffectRepository {
	return &effectRepository{
		effects: []gorpo.Effect{
			effect.NewRotate(),
			effect.NewResize(),
		},
	}
}

func (r *effectRepository) GetEffects() ([]gorpo.Effect, error) {
	return r.effects, nil
}

func (r *effectRepository) GetEffect(id string) (gorpo.Effect, error) {
	for _, effect := range r.effects {
		if effect.Descriptor().Id == id {
			return effect, nil
		}
	}

	return nil, gorpo.ENotExists(fmt.Sprintf("Effect %s does not exists", id))
}
