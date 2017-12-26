package downloader

import (
	"fmt"
	"image"
	"net/http"

	"github.com/talento90/gorpo/errors"
)

type httpdownloader struct {
	client *http.Client
}

// NewHTTPDownloader creates a Downloader that get an image over the HTTP protocol.
func NewHTTPDownloader() Downloader {
	return &httpdownloader{
		client: http.DefaultClient,
	}
}

func (d *httpdownloader) DownloadImage(path string) (image.Image, string, error) {
	response, err := d.client.Get(path)

	if err != nil {
		return nil, "", errors.ENotExists("")
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, "", errors.ENotExists(fmt.Sprintf("Image %s not found", path))
	}

	img, imgType, err := image.Decode(response.Body)

	if err != nil {
		return nil, "", err
	}

	return img, imgType, nil
}
