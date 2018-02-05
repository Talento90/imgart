package imgart

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/bmp"
)

// Filter represents an effect with all parameters
type Filter struct {
	// ID of the effect
	ID string `json:"id" bson:"filters"`
	// Parameters to apply
	Parameters map[string]interface{} `json:"parameters" bson:"parameters"`
}

// ImageService interface has the logic for processing images by the given a list of filters
type ImageService interface {
	// Effects returns all available effects for the image
	Effects() ([]Effect, error)
	// Effect returns a effect by the given ID
	Effect(id string) (Effect, error)
	// Process an image with a set of filters
	Process(imgSrc string, filters []Filter) (image.Image, string, error)
}

// ImageRepository interface layer to get images
type ImageRepository interface {
	// Get an image by the given path
	Get(path string) (image.Image, string, error)
}

// Encode image by their format (png, jpeg, bmp)
func Encode(imgFormat string, img image.Image, quality int) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error

	switch imgFormat {
	case "png":
		err = png.Encode(buf, img)
	case "jpeg":
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
	case "bmp":
		err = bmp.Encode(buf, img)
	default:
		return nil, image.ErrFormat
	}

	return buf.Bytes(), err
}
