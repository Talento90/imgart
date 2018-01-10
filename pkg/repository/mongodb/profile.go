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

func (r *profileRepository) GetAll() ([]gorpo.Profile, error) {
	profiles := new([]gorpo.Profile)

	err := r.collection.Find(bson.M{}).Skip(5).Limit(5).All(profiles)

	return profiles, err
}

func (r *profileRepository) Get(id string) (Profile, error) {

}

func (r *profileRepository) Create(profile Profile) (Profile, error) {

}

func (r *profileRepository) Update(profile Profile) error {

}

func (r *profileRepository) Delete(id string) error {

}
