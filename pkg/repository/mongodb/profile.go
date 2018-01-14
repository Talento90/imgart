package mongodb

import (
	"time"

	"github.com/talento90/gorpo/pkg/errors"
	"github.com/talento90/gorpo/pkg/gorpo"
	"gopkg.in/mgo.v2"
)

type profileRepository struct {
	database   string
	collection string
	session    *mgo.Session
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	if err == mgo.ErrNotFound {
		return errors.ENotExists("Profile does not exists", err)
	}

	return errors.EInternal("Error occured", err)
}

// NewProfileRepository returns a profile mongo repository
func NewProfileRepository(config Configuration, session *mgo.Session) gorpo.ProfileRepository {
	return &profileRepository{
		database:   config.Database,
		collection: "profiles",
		session:    session,
	}
}

func (r *profileRepository) GetAll(limit int, skip int) (*[]gorpo.Profile, error) {
	session := r.session.Copy()
	defer session.Close()
	c := session.DB(r.database).C(r.collection)

	profiles := make([]gorpo.Profile, 0, limit)
	err := c.Find(nil).Skip(skip).Limit(limit).All(&profiles)

	return &profiles, handleError(err)
}

func (r *profileRepository) Get(id string) (*gorpo.Profile, error) {
	session := r.session.Copy()
	defer session.Close()

	c := session.DB(r.database).C(r.collection)

	profile := &gorpo.Profile{}
	err := c.FindId(id).One(profile)

	return profile, handleError(err)
}

func (r *profileRepository) Create(profile *gorpo.Profile) error {
	session := r.session.Copy()
	defer session.Close()
	c := session.DB(r.database).C(r.collection)

	profile.Created = time.Now().UTC()
	profile.Updated = time.Now().UTC()

	err := c.Insert(profile)

	return handleError(err)
}

func (r *profileRepository) Update(profile *gorpo.Profile) error {
	session := r.session.Copy()
	defer session.Close()
	c := session.DB(r.database).C(r.collection)

	profile.Updated = time.Now().UTC()

	err := c.UpdateId(profile.ID, profile)

	return handleError(err)
}

func (r *profileRepository) Delete(id string) error {
	session := r.session.Copy()
	defer session.Close()
	c := session.DB(r.database).C(r.collection)

	err := c.RemoveId(id)

	return handleError(err)
}
