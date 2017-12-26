package image

import (
	"image"

	"github.com/talento90/gorpo/downloader"
	"github.com/talento90/gorpo/effect"
)

// Service interface has the logic for processing images by the given a list of filters
type Service interface {
	Process(imgSrc string, filters []effect.Filter) (image.Image, error)
}

type service struct {
	downloader downloader.Downloader
	repository effect.Repository
}

// NewService creates an ImageService
func NewService(downloader downloader.Downloader, repo effect.Repository) Service {
	return &service{
		downloader: downloader,
		repository: repo,
	}
}

func (i *service) Process(imgSrc string, filters []effect.Filter) (image.Image, error) {
	img, _, err := i.downloader.DownloadImage(imgSrc)

	if err != nil {
		return nil, err
	}

	for _, filter := range filters {
		effect, err := i.repository.GetEffect(filter.ID)

		if err != nil {
			return nil, err
		}

		img, err = effect.Transform(img, filter.Parameters)

		if err != nil {
			return nil, err
		}
	}

	return img, nil
}
