package mock

import (
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/talento90/gorpo/pkg/errors"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type profileRepo struct {
	mutex      *sync.Mutex
	repository map[string]gorpo.Profile
}

// NewProfileRepository returns a mock implemation of ProfileRepository interface
func NewProfileRepository() gorpo.ProfileRepository {
	return &profileRepo{
		mutex:      &sync.Mutex{},
		repository: map[string]gorpo.Profile{},
	}
}

func (r *profileRepo) GetAll(limit int, skip int) (*[]gorpo.Profile, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var all = make([]gorpo.Profile, 0)
	skipCounter := 0

	for _, value := range r.repository {

		if skipCounter == skip {
			all = append(all, value)

			if len(all) == limit {
				break
			}

		} else {
			skipCounter = skipCounter + 1
		}

	}

	return &all, nil
}

func (r *profileRepo) Get(id string) (*gorpo.Profile, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if v, ok := r.repository[id]; ok {
		return &v, nil
	}

	return nil, errors.ENotExists("Profile does not exists", nil)
}

func (r *profileRepo) Create(profile *gorpo.Profile) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	profile.ID = bson.NewObjectId().Hex()
	profile.Created = time.Now().UTC()
	profile.Updated = time.Now().UTC()
	r.repository[profile.ID] = *profile

	return nil
}

func (r *profileRepo) Update(profile *gorpo.Profile) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.repository[profile.ID]; ok {
		profile.Updated = time.Now().UTC()
		r.repository[profile.ID] = *profile
	}

	return nil
}

func (r *profileRepo) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.repository[id]; ok {
		delete(r.repository, id)
		return nil
	}

	return errors.ENotExists("Profile does not exist", nil)
}
