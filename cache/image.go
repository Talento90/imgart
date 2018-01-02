package cache

import "image"

type ImageCache interface {
	Get(key string) (image.Image, error)
	Set(key string, value image.Image) error
}
