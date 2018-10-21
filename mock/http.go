package mock

import (
	"context"
	"image"

	"github.com/talento90/imgart/imgart"
)

// NewImageRepository returns a mock implementation of ImageRepository
func NewImageRepository() imgart.ImageRepository {
	return &httpMock{}
}

type httpMock struct{}

func (h *httpMock) Get(ctx context.Context, path string) (image.Image, string, error) {
	return image.NewRGBA(image.Rect(0, 0, 100, 50)), "jpeg", nil
}
