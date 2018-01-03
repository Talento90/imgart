package gorpo

import (
	"image"
)

// ImageService interface has the logic for processing images by the given a list of filters
type ImageService interface {
	Process(imgSrc string, filters []Filter) (image.Image, error)
}

// ImageRepository interface layer to get images
type ImageRepository interface {
	Get(path string) (image.Image, string, error)
}

// NewImageService creates an ImageService
func NewImageService(imgRepo ImageRepository, effectRepo EffectRepository) ImageService {
	return &imageService{
		imgRepo:    imgRepo,
		effectRepo: effectRepo,
	}
}

type imageService struct {
	imgRepo    ImageRepository
	effectRepo EffectRepository
}

func (i *imageService) Process(imgSrc string, filters []Filter) (image.Image, error) {
	img, _, err := i.imgRepo.Get(imgSrc)

	if err != nil {
		return nil, err
	}

	for _, filter := range filters {
		effect, err := i.effectRepo.GetEffect(filter.ID)

		if err != nil {
			return nil, err
		}

		img, err = effect.Transform(img, filter.Parameters)

		if err != nil {
			return nil, err
		}
	}

	return img, nil
}
