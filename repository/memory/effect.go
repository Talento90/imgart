package memory

import (
	"fmt"

	"github.com/talento90/gorpo/effect"
	"github.com/talento90/gorpo/errors"
)

type effectRepository struct {
	effects []effect.Effect
}

// NewEffectRepository creates a memory repository for Effect entity
func NewEffectRepository() effect.Repository {
	return &effectRepository{
		effects: []effect.Effect{
			effect.NewRotate(),
			effect.NewResize(),
		},
	}
}

func (r *effectRepository) GetEffects() ([]effect.Effect, error) {
	return r.effects, nil
}

func (r *effectRepository) GetEffect(id string) (effect.Effect, error) {
	for _, effect := range r.effects {
		if effect.Descriptor().ID == id {
			return effect, nil
		}
	}

	return nil, errors.ENotExists(fmt.Sprintf("Effect %s does not exist", id))
}
