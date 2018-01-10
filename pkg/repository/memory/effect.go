package memory

import (
	"fmt"

	"github.com/talento90/gorpo/pkg/effect"
	"github.com/talento90/gorpo/pkg/errors"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type effectRepository struct {
	effects []gorpo.Effect
}

// NewImageRepository creates a memory repository for Effect entity
func NewImageRepository(imgRepo gorpo.ImageRepository) gorpo.EffectRepository {
	return &effectRepository{
		effects: []gorpo.Effect{
			effect.NewRotate(),
			effect.NewResize(),
			effect.NewOverlay(imgRepo),
			effect.NewBlur(),
			effect.NewBrightness(),
			effect.NewGamma(),
			effect.NewContrast(),
			effect.NewCrop(),
		},
	}
}

func (r *effectRepository) GetEffects() ([]gorpo.Effect, error) {
	return r.effects, nil
}

func (r *effectRepository) GetEffect(id string) (gorpo.Effect, error) {
	for _, effect := range r.effects {
		if effect.ID() == id {
			return effect, nil
		}
	}

	return nil, errors.ENotExists(fmt.Sprintf("Effect %s does not exist", id), nil)
}
