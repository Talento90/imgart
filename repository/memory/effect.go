package memory

import (
	"fmt"

	"github.com/talento90/gorpo/effect"
	"github.com/talento90/gorpo/errors"
	"github.com/talento90/gorpo/gorpo"
)

type effectRepository struct {
	effects []gorpo.Effect
}

// NewEffectRepository creates a memory repository for Effect entity
func NewEffectRepository(downloader gorpo.ImageRepository) gorpo.EffectRepository {
	return &effectRepository{
		effects: []gorpo.Effect{
			effect.NewRotate(),
			effect.NewResize(),
			effect.NewOverlay(downloader),
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
		if effect.Descriptor().ID == id {
			return effect, nil
		}
	}

	return nil, errors.ENotExists(fmt.Sprintf("Effect %s does not exist", id), nil)
}
