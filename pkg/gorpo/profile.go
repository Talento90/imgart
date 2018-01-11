package gorpo

import (
	"time"
)

// Profile that represents a set of effects
type Profile struct {
	ID      string
	Created time.Time
	Updated time.Time
	Effects []Effect
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
