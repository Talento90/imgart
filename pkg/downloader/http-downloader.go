package downloader

import (
	"fmt"
	"image"
	"net/http"

	"github.com/talento90/gorpo/pkg/gorpo"
)

type httpdownloader struct {
	client *http.Client
}

// NewHTTPDownloader creates a Downloader that download images over the HTTP protocol.
func NewHTTPDownloader() gorpo.Downloader {
	return &httpdownloader{
		client: http.DefaultClient,
	}
}

func (d *httpdownloader) DownloadImage(path string) (image.Image, string, error) {
	response, err := d.client.Get(path)

	if err != nil {
		return nil, "", gorpo.ENotExists("")
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, "", gorpo.ENotExists(fmt.Sprintf("Image %s not found", path))
	}

	img, imgType, err := image.Decode(response.Body)

	if err != nil {
		return nil, "", err
	}

	return img, imgType, nil
}
