package fs

import (
	"context"
	"fmt"
	"image"
	"os"

	"github.com/talento90/imgart/imgart"

	"github.com/talento90/imgart/errors"
)

type fsdownloader struct{}

// NewImageRepository creates a Downloader that get an image over the File System.
func NewImageRepository() imgart.ImageRepository {
	return &fsdownloader{}
}

func (d *fsdownloader) Get(ctx context.Context, path string) (image.Image, string, error) {
	file, err := os.Open(path)

	if err == os.ErrNotExist {
		return nil, "", errors.ENotExists(fmt.Sprintf("Image path: %s not exists", path), nil)
	} else if err != nil {
		return nil, "", errors.EInternal("Error trying to get image", err)
	}

	defer file.Close()

	img, imgType, err := image.Decode(file)

	if err != nil {
		return nil, "", errors.EInternal("Error decoding image", err)
	}

	return img, imgType, ctx.Err()
}
