package image

import (
	"image"

	"github.com/talento90/gorpo/pkg/gorpo"
)

// NewService creates an ImageService
func NewService(imageRepo gorpo.ImageRepository, effectRepo gorpo.EffectRepository) gorpo.ImageService {
	return &service{
		imageRepo:  imageRepo,
		effectRepo: effectRepo,
	}
}

type service struct {
	imageRepo  gorpo.ImageRepository
	effectRepo gorpo.EffectRepository
}

func (s *service) Process(imgSrc string, filters []gorpo.Filter) (image.Image, string, error) {
	img, imgType, err := s.imageRepo.Get(imgSrc)

	if err != nil {
		return nil, imgType, err
	}

	for _, filter := range filters {
		effect, err := s.effectRepo.GetEffect(filter.ID)

		if err != nil {
			return nil, imgType, err
		}

		img, err = effect.Transform(img, filter.Parameters)

		if err != nil {
			return nil, imgType, err
		}
	}

	return img, imgType, nil
}

func (s *service) Effects() ([]gorpo.Effect, error) {
	return s.effectRepo.GetEffects()
}

func (s *service) Effect(id string) (gorpo.Effect, error) {
	return s.effectRepo.GetEffect(id)
}
