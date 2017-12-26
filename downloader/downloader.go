package downloader

import (
	"image"
)

// Downloader is the interface for downloading images
type Downloader interface {
	DownloadImage(path string) (image.Image, string, error)
}
