package downloaders

import (
	"image"
	"net/http"

	"github.com/talento90/merlin/merlin"
)

type HTTPDownloader struct {
	client *http.Client
}

func NewHTTPDownloader() merlin.Downloader {
	return &HTTPDownloader{
		client: http.DefaultClient,
	}
}

func (d *HTTPDownloader) DownloadImage(path string) (image.Image, string, error) {
	response, err := d.client.Get(path)

	if err != nil {
		return nil, "", err
	}

	defer response.Body.Close()

	img, imgType, err := image.Decode(response.Body)

	if err != nil {
		return nil, "", err
	}

	return img, imgType, nil
}
