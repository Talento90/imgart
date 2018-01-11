package profile

import (
	"github.com/talento90/gorpo/pkg/gorpo"
)

type service struct {
	repository gorpo.ProfileRepository
}

// NewService returns a profile service
func NewService(repository gorpo.ProfileRepository) gorpo.ProfileService {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll(limit int, skip int) (*[]gorpo.Profile, error) {
	return s.repository.GetAll(limit, skip)
}

func (s *service) Get(id string) (*gorpo.Profile, error) {
	return s.repository.Get(id)
}

func (s *service) Create(profile *gorpo.Profile) error {
	// TODO: check if effects are valid
	return s.repository.Create(profile)
}

func (s *service) Update(profile *gorpo.Profile) error {
	// TODO: check if effects are valid
	return s.repository.Update(profile)
}

func (s *service) Delete(id string) error {
	return s.repository.Delete(id)
}
