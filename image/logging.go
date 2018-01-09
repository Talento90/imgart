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

func (ls *logService) Effects() ([]gorpo.Effect, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(log.Fields{"time": time.Now().Sub(start)}, "ImageService:Effects")
	}(time.Now())

	return ls.service.Effects()
}

func (ls *logService) Effect(id string) (gorpo.Effect, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(log.Fields{"id": id, "time": time.Now().Sub(start)}, "ImageService:Effect")
	}(time.Now())

	return ls.service.Effect(id)
}
