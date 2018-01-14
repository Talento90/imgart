package http

import (
	"fmt"
	"image"
	"net/http"

	"github.com/talento90/gorpo/pkg/gorpo"

	"github.com/talento90/gorpo/pkg/errors"
)

type httpdownloader struct {
	client *http.Client
}

// NewImageRepository creates a Downloader that get an image over the HTTP protocol.
func NewImageRepository() gorpo.ImageRepository {
	return &httpdownloader{
		client: http.DefaultClient,
	}
}

func (d *httpdownloader) Get(path string) (image.Image, string, error) {
	response, err := d.client.Get(path)

	if err != nil {
		return nil, "", errors.EInternal("Error trying to download image", err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, "", errors.ENotExists(fmt.Sprintf("Image %s not found", path), nil)
	}

	img, imgType, err := image.Decode(response.Body)

	if err != nil {
		return nil, "", errors.EInternal("Error decoding image", err)
	}

	return img, imgType, nil
}
