package image

import (
	"image"
	"time"

	"github.com/talento90/gorpo/log"

	"github.com/talento90/gorpo/gorpo"
)

type logService struct {
	logger  log.Logger
	service gorpo.ImageService
}

// NewLogService creates a log wrapper around ImageService
func NewLogService(logger log.Logger, service gorpo.ImageService) gorpo.ImageService {
	return &logService{
		logger:  logger,
		service: service,
	}
}

func (ls *logService) Process(imgSrc string, filters []gorpo.Filter) (image.Image, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(log.Fields{"imgSrc": imgSrc, "time": time.Now().Sub(start)}, "ImageService:Process")
	}(time.Now())

	return ls.service.Process(imgSrc, filters)
}
