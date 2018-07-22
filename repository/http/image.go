package http

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"net/http"
	"time"

	"github.com/talento90/imgart/imgart"

	"github.com/talento90/imgart/errors"
)

const maxImageSize = 1024 * 1024 * 5

type httpdownloader struct {
	client *http.Client
}

// NewImageRepository creates a Downloader that get an image over the HTTP protocol.
func NewImageRepository() imgart.ImageRepository {
	return &httpdownloader{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (d *httpdownloader) Get(ctx context.Context, path string) (image.Image, string, error) {
	req, err := http.NewRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, "", errors.EInternal("Error trying to create http request", err)
	}

	req = req.WithContext(ctx)

	response, err := d.client.Do(req)

	if err != nil {
		if ctx.Err() != nil {
			return nil, "", ctx.Err()
		}

		return nil, "", errors.EInternal("Error trying to download image", err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, "", errors.ENotExists(fmt.Sprintf("Image %s not found", path), nil)
	}

	imgBytes := &bytes.Buffer{}

	if _, err := io.CopyN(imgBytes, response.Body, maxImageSize); err != io.EOF {

		if ctx.Err() != nil {
			return nil, "", ctx.Err()
		}

		return nil, "", errors.EValidation(fmt.Sprintf("Image size is bigger than: %d", maxImageSize), err)
	}

	img, imgType, err := image.Decode(imgBytes)

	if err != nil {
		return nil, "", errors.EInternal("Error decoding image", err)
	}

	return img, imgType, ctx.Err()
}
