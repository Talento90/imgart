package cache

import "image"

// ImageCache
type ImageCache interface {
	Get(key string) (image.Image, error)
	Set(key string, value image.Image) error
}
