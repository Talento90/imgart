package profile

import (
	"github.com/talento90/imgart/pkg/imgart"
)

type service struct {
	repository imgart.ProfileRepository
}

// NewService returns a profile service
func NewService(repository imgart.ProfileRepository) imgart.ProfileService {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll(limit int, skip int) (*[]imgart.Profile, error) {
	return s.repository.GetAll(limit, skip)
}

func (s *service) Get(id string) (*imgart.Profile, error) {
	return s.repository.Get(id)
}

func (s *service) Create(profile *imgart.Profile) error {
	return s.repository.Create(profile)
}

func (s *service) Update(profile *imgart.Profile) error {
	return s.repository.Update(profile)
}

func (s *service) Delete(id string) error {
	return s.repository.Delete(id)
}
