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
