package mongo

import (
	"time"

	"github.com/talento90/imgart/pkg/errors"
	"github.com/talento90/imgart/pkg/imgart"
	"gopkg.in/mgo.v2"
)

type profileRepository struct {
	collection string
	session    *Session
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	if err == mgo.ErrNotFound {
		return errors.ENotExists("Profile does not exists", err)
	}

	if mgo.IsDup(err) {
		return errors.EAlreadyExists("Profile already exists", err)
	}

	return errors.EInternal("Error occured", err)
}

// NewProfileRepository returns a profile mongo repository
func NewProfileRepository(session *Session) imgart.ProfileRepository {
	return &profileRepository{
		collection: "profiles",
		session:    session,
	}
}

func (r *profileRepository) GetAll(limit int, skip int) (*[]imgart.Profile, error) {
	session := r.session.Copy()

	defer session.Close()
	c := session.DB(r.session.Database).C(r.collection)

	profiles := make([]imgart.Profile, 0, limit)
	err := c.Find(nil).Skip(skip).Limit(limit).All(&profiles)

	return &profiles, handleError(err)
}

func (r *profileRepository) Get(id string) (*imgart.Profile, error) {
	session := r.session.Copy()

	defer session.Close()

	c := session.DB(r.session.Database).C(r.collection)

	profile := &imgart.Profile{}
	err := c.FindId(id).One(profile)

	return profile, handleError(err)
}

func (r *profileRepository) Create(profile *imgart.Profile) error {
	session := r.session.Copy()
	defer session.Close()

	c := session.DB(r.session.Database).C(r.collection)

	profile.Created = time.Now().UTC()
	profile.Updated = time.Now().UTC()

	err := c.Insert(profile)

	return handleError(err)
}

func (r *profileRepository) Update(profile *imgart.Profile) error {
	session := r.session.Copy()
	defer session.Close()

	c := session.DB(r.session.Database).C(r.collection)

	profile.Updated = time.Now().UTC()

	err := c.UpdateId(profile.ID, profile)

	return handleError(err)
}

func (r *profileRepository) Delete(id string) error {
	session := r.session.Copy()
	defer session.Close()

	c := session.DB(r.session.Database).C(r.collection)

	err := c.RemoveId(id)

	return handleError(err)
}
