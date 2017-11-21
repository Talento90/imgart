package gomage

import (
	"image"
)

type ImageParameters map[string]string

type Effect interface {
	Description() ImageParameters

	Transform(img image.Image, params ImageParameters) image.Image
}
