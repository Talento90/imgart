package mock

import (
	"image"

	"github.com/talento90/gorpo/pkg/gorpo"
)

// NewImageRepository returns a mock implementation of ImageRepository
func NewImageRepository() gorpo.ImageRepository {
	return &httpMock{}
}

type httpMock struct{}

func (h *httpMock) Get(path string) (image.Image, string, error) {
	return image.NewRGBA(image.Rect(0, 0, 100, 50)), "jpeg", nil
}
