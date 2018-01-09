package gorpo

import (
	"image"
)

// Filter represents an effect with all parameters
type Filter struct {
	// ID of the effect
	ID string `json:"id"`
	// Parameters to apply
	Parameters map[string]interface{} `json:"parameters"`
}

// ImageService interface has the logic for processing images by the given a list of filters
type ImageService interface {
	// Effects returns all available effects for the image
	Effects() ([]Effect, error)
	// Effect returns a effect by the given ID
	Effect(id string) (Effect, error)
	// Process an image with a set of filters
	Process(imgSrc string, filters []Filter) (image.Image, error)
}

// ImageRepository interface layer to get images
type ImageRepository interface {
	// Get an image by the given path
	Get(path string) (image.Image, string, error)
}
