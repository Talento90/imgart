package profile

import (
	"time"

	"github.com/talento90/gorpo/pkg/log"

	"github.com/talento90/gorpo/pkg/gorpo"
)

type logService struct {
	logger  log.Logger
	service gorpo.ProfileService
}

// NewLogService creates a log wrapper around ProfileService
func NewLogService(logger log.Logger, service gorpo.ProfileService) gorpo.ProfileService {
	return &logService{
		logger:  logger,
		service: service,
	}
}

func (ls *logService) GetAll(limit int, skip int) (*[]gorpo.Profile, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"limit": limit,
				"skip":  skip,
				"time":  time.Now().Sub(start)},
			"ProfileService:GetAll")
	}(time.Now())

	return ls.service.GetAll(limit, skip)
}

func (ls *logService) Get(id string) (*gorpo.Profile, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"id":   id,
				"time": time.Now().Sub(start)},
			"ProfileService:Get")
	}(time.Now())

	return ls.service.Get(id)
}

func (ls *logService) Create(profile *gorpo.Profile) error {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"profile": profile,
				"time":    time.Now().Sub(start)},
			"ProfileService:Create")
	}(time.Now())

	return ls.service.Create(profile)
}

func (ls *logService) Update(profile *gorpo.Profile) error {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"profile": profile,
				"time":    time.Now().Sub(start)},
			"ProfileService:Update")
	}(time.Now())

	return ls.service.Update(profile)
}

func (ls *logService) Delete(id string) error {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"id":   id,
				"time": time.Now().Sub(start)},
			"ProfileService:Delete")
	}(time.Now())

	return ls.service.Delete(id)
}
