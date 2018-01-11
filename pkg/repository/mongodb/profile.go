package mongodb

import (
	"github.com/talento90/gorpo/pkg/gorpo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type profileRepository struct {
	collection *mgo.Collection
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

func (r *profileRepository) Create(profile gorpo.Profile) (*gorpo.Profile, error) {
	err := r.collection.Insert(profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *profileRepository) Update(profile Profile) error {

	r.collection.Update(mgo.M)
}

func (r *profileRepository) Delete(id string) error {
	err := r.collection.RemoveId(id)

	return err
}
