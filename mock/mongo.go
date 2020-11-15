package mock

import (
	"sync"
	"time"

	"github.com/talento90/imgart/errors"
	"github.com/talento90/imgart/imgart"
)

type profileRepo struct {
	mutex      *sync.Mutex
	repository []imgart.Profile
}

// NewProfileRepository returns a mock implemation of ProfileRepository interface
func NewProfileRepository() imgart.ProfileRepository {
	return &profileRepo{
		mutex:      &sync.Mutex{},
		repository: []imgart.Profile{},
	}
}

func (r *profileRepo) GetAll(limit int64, skip int64) (*[]imgart.Profile, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	p := r.repository[skip : limit+skip]

	return &p, nil
}

func (r *profileRepo) Get(id string) (*imgart.Profile, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, v := range r.repository {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.ENotExists("Profile does not exists", nil)
}

func (r *profileRepo) Create(profile *imgart.Profile) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	profile.Created = time.Now().UTC()
	profile.Updated = time.Now().UTC()

	r.repository = append(r.repository, *profile)

	return nil
}

func (r *profileRepo) Update(profile *imgart.Profile) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, v := range r.repository {
		if profile.ID == v.ID {
			profile.Updated = time.Now().UTC()
			r.repository[i] = *profile
			return nil
		}
	}

	return errors.ENotExists("Profile does not exist", nil)
}

func (r *profileRepo) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, v := range r.repository {
		if v.ID == id {
			r.repository = append(r.repository[:i], r.repository[i+1:]...)
			return nil
		}
	}

	return errors.ENotExists("Profile does not exist", nil)
}
