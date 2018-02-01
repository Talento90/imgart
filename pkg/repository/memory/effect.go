package memory

import (
	"fmt"

	"github.com/talento90/imgart/pkg/effect"
	"github.com/talento90/imgart/pkg/errors"
	"github.com/talento90/imgart/pkg/imgart"
)

type effectRepository struct {
	effects []imgart.Effect
}

// NewImageRepository creates a memory repository for Effect entity
func NewImageRepository(imgRepo imgart.ImageRepository) imgart.EffectRepository {
	return &effectRepository{
		effects: []imgart.Effect{
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

func (r *effectRepository) GetEffects() ([]imgart.Effect, error) {
	return r.effects, nil
}

func (r *effectRepository) GetEffect(id string) (imgart.Effect, error) {
	for _, effect := range r.effects {
		if effect.ID() == id {
			return effect, nil
		}
	}

	return nil, errors.ENotExists(fmt.Sprintf("Effect %s does not exist", id), nil)
}
