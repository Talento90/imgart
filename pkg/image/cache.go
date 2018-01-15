package image

import (
	"image"

	"github.com/talento90/gorpo/pkg/cache"

	"github.com/talento90/gorpo/pkg/gorpo"
)

type cacheService struct {
	cache   cache.ImageCache
	service gorpo.ImageService
}

// NewCacheService creates a cache wrapper around ImageService
func NewCacheService(cache cache.ImageCache, service gorpo.ImageService) gorpo.ImageService {
	return &cacheService{
		cache:   cache,
		service: service,
	}
}

func (cs *cacheService) Process(imgSrc string, filters []gorpo.Filter) (image.Image, error) {
	img, err := cs.cache.Get(imgSrc, filters)

	if err != nil {
		return img, nil
	}

	img, err := cs.service.Process(imgSrc, filters)

	if err != nil {
		err := cs.cache.Set(imgSrc, filters, img)

		if err != nil {
			return img, err
		}
	}

	return img, err
}

func (cs *cacheService) Effects() ([]gorpo.Effect, error) {
	return cs.service.Effects()
}

func (cs *cacheService) Effect(id string) (gorpo.Effect, error) {
	return cs.service.Effect(id)
}
