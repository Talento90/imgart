package image

import (
	"context"
	"image"

	"github.com/talento90/imgart/cache"

	"github.com/talento90/imgart/imgart"
)

type cacheService struct {
	cache   cache.Image
	service imgart.ImageService
}

// NewCacheService creates a cache wrapper around ImageService
func NewCacheService(cache cache.Image, service imgart.ImageService) imgart.ImageService {
	return &cacheService{
		cache:   cache,
		service: service,
	}
}

func (cs *cacheService) Process(ctx context.Context, imgSrc string, filters []imgart.Filter) (image.Image, string, error) {
	img, format, err := cs.cache.Get(imgSrc, filters)

	if err == nil {
		return img, format, nil
	}

	img, format, err = cs.service.Process(ctx, imgSrc, filters)

	if err == nil {
		err := cs.cache.Set(imgSrc, filters, format, img)

		if err != nil {
			return img, format, err
		}
	}

	return img, format, err
}

func (cs *cacheService) Effects() ([]imgart.Effect, error) {
	return cs.service.Effects()
}

func (cs *cacheService) Effect(id string) (imgart.Effect, error) {
	return cs.service.Effect(id)
}
