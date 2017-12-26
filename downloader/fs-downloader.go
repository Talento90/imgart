package downloader

import (
	"fmt"
	"image"
	"os"

	"github.com/talento90/gorpo/errors"
)

type fsdownloader struct{}

// NewFSDownloader creates a Downloader that get an image over the File System.
func NewFSDownloader() Downloader {
	return &fsdownloader{}
}

func (d *fsdownloader) DownloadImage(path string) (image.Image, string, error) {
	file, err := os.Open(path)

	if err == os.ErrNotExist {
		return nil, "", errors.ENotExists(fmt.Sprintf("Image path: %s not exists", path))
	}

	defer file.Close()

	img, imgType, err := image.Decode(file)

	if err != nil {
		return nil, "", err
	}

	return img, imgType, nil
}
