package mongodb

import (
	"time"

	"github.com/talento90/gorpo/pkg/errors"
	"github.com/talento90/gorpo/pkg/gorpo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type profileRepository struct {
	collection *mgo.Collection
}

func handleError(err error) error {
	if err == mgo.ErrNotFound {
		return errors.ENotExists("Profile does not exists", err)
	}

	return errors.EInternal("Error occured", err)
}

// NewProfile returns a profile mongo repository
func NewProfile(db *mgo.Database) gorpo.ProfileRepository {
	return &profileRepository{
		collection: db.C("profiles"),
	}
}

func (r *profileRepository) GetAll(limit int, skip int) (*[]gorpo.Profile, error) {
	profiles := make([]gorpo.Profile, 0, limit)

	err := r.collection.Find(bson.M{}).Skip(skip).Limit(limit).All(profiles)

	return &profiles, err
}

func (r *profileRepository) Get(id string) (*gorpo.Profile, error) {
	profile := &gorpo.Profile{}

	err := r.collection.FindId(id).One(profile)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (r *profileRepository) Create(profile *gorpo.Profile) error {
	profile.Created = time.Now().UTC()
	profile.Updated = time.Now().UTC()

	err := r.collection.Insert(profile)

	return handleError(err)
}

func (r *profileRepository) Update(profile *gorpo.Profile) error {
	err := r.collection.UpdateId(profile.ID, profile)

	return handleError(err)
}

func (r *profileRepository) Delete(id string) error {
	err := r.collection.RemoveId(id)

	return handleError(err)
}
