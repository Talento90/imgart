package cache

import (
	"crypto/md5"
	"image"

	"github.com/talento90/gorpo/pkg/gorpo"
)

// ImageCache caches images by the given URL
type ImageCache interface {
	Get(url string, filters []gorpo.Filter) (image.Image, error)
	Set(url string, filters []gorpo.Filter, value image.Image) error
}

func NewImageCache(cache Cache) ImageCache {
	return &imageCache{
		cache: cache,
	}
}

type imageCache struct {
	cache Cache
}

func generateHash(url string, filters []gorpo.Filter) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(url))

	if err != nil {
		return "", err
	}

	for _, v := range filters {
		h.Write([]byte(v.ID))

		for _, p := range v.Parameters {
			bytes, ok := p.([]byte)

			if !ok {
				return "", err
			}

			_, err := h.Write(bytes)

			if err != nil {
				return "", err
			}
		}
	}

	//hex.EncodeToString(hasher.Sum(nil))
	return string(h.Sum(nil)), nil
}

func (c *imageCache) Get(url string, filters []gorpo.Filter) (image.Image, error) {

	_, err := generateHash(url, filters)

	return nil, err
}

func (c *imageCache) Set(url string, filters []gorpo.Filter, value image.Image) error {

	_, err := generateHash(url, filters)

	return err
}
