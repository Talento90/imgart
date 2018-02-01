package httpapi

import (
	"github.com/talento90/imgart/pkg/errors"
	"github.com/talento90/imgart/pkg/imgart"
)

type createProfileModel struct {
	ID      string          `json:"id"`
	Filters []imgart.Filter `json:"filters"`
}

type updateProfileModel struct {
	Filters []imgart.Filter `json:"filters"`
}

func (m *createProfileModel) toProfile() (*imgart.Profile, error) {
	if m.ID == "" {
		return nil, errors.EValidation("id is missing", nil)
	}

	if len(m.Filters) == 0 {
		return nil, errors.EValidation("filters are empty", nil)
	}

	return &imgart.Profile{ID: m.ID, Filters: m.Filters}, nil
}

func (m *updateProfileModel) toProfile(profile *imgart.Profile) (*imgart.Profile, error) {
	if len(m.Filters) == 0 {
		return nil, errors.EValidation("effects are empty", nil)
	}

	profile.Filters = m.Filters

	return profile, nil
}
