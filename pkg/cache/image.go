package cache

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
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
	arrBytes := []byte{}

	arrBytes = append(arrBytes, url...)

	for _, filter := range filters {
		jsonBytes, err := json.Marshal(filter)

		if err != nil {
			return "", err
		}

		arrBytes = append(arrBytes, jsonBytes...)
	}

	hash := md5.Sum(arrBytes)

	return string(hash[:]), nil
}

func (c *imageCache) Get(url string, filters []gorpo.Filter) (image.Image, error) {
	hash, err := generateHash(url, filters)

	if err != nil {
		return nil, err
	}

	imgBytes, err := c.cache.Get(hash)

	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(imgBytes)

	img, _, err := image.Decode(r)

	if err != nil {
		return nil, err
	}

	return img, err
}

func (c *imageCache) Set(url string, filters []gorpo.Filter, value image.Image) error {
	_, err := generateHash(url, filters)

	if err != nil {
		return err
	}

	//c.cache.Set(hash, value, time.Second)

	return err
}
