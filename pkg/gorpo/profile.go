package gorpo

import (
	"time"
)

// Profile that represents a set of effects
type Profile struct {
	ID      string    `json:"id" bson:"_id"`
	Created time.Time `json:"created" bson:"created"`
	Updated time.Time `json:"updated" bson:"updated"`
	Filters []Filter  `json:"filters" bson:"filters"`
}

// ProfileRepository stores profiles
type ProfileRepository interface {
	GetAll(limit int, skip int) (*[]Profile, error)
	Get(id string) (*Profile, error)
	Create(profile *Profile) error
	Update(profile *Profile) error
	Delete(id string) error
}

// ProfileService handles profile operations
type ProfileService interface {
	GetAll(limit int, skip int) (*[]Profile, error)
	Get(id string) (*Profile, error)
	Create(profile *Profile) error
	Update(profile *Profile) error
	Delete(id string) error
}
