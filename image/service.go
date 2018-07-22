package image

import (
	"context"
	"image"

	"github.com/talento90/imgart/imgart"
)

// NewService creates an ImageService
func NewService(imageRepo imgart.ImageRepository, effectRepo imgart.EffectRepository) imgart.ImageService {
	return &service{
		imageRepo:  imageRepo,
		effectRepo: effectRepo,
	}
}

type service struct {
	imageRepo  imgart.ImageRepository
	effectRepo imgart.EffectRepository
}

func (s *service) Process(ctx context.Context, imgSrc string, filters []imgart.Filter) (image.Image, string, error) {
	img, imgType, err := s.imageRepo.Get(ctx, imgSrc)

	if err != nil {
		return nil, imgType, err
	}

	for _, filter := range filters {
		effect, err := s.effectRepo.GetEffect(filter.ID)

		if err != nil {
			return nil, imgType, err
		}

		img, err = effect.Transform(ctx, img, filter.Parameters)

		if err != nil {
			return nil, imgType, err
		}
	}

	return img, imgType, ctx.Err()
}

func (s *service) Effects() ([]imgart.Effect, error) {
	return s.effectRepo.GetEffects()
}

func (s *service) Effect(id string) (imgart.Effect, error) {
	return s.effectRepo.GetEffect(id)
}
