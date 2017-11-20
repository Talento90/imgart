package downloaders

import (
	"go-mage/gomage"
	"image"
	"net/http"
)

type HTTPDownloader struct {
	client *http.Client
}

func NewHTTPDownloader() gomage.Downloader {
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
