package gorpo

// Profile that represents a set of effects
type Profile struct {
	ID      string
	Effects []Effect
}

// ProfileRepository stores profiles
type ProfileRepository interface {
	GetAll() ([]Profile, string, error)
	Get(id string) (Profile, error)
	Create(profile Profile) (Profile, error)
	Update(profile Profile) error
	Delete(id string) error
}
