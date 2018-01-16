package httpapi

import (
	"github.com/talento90/gorpo/pkg/errors"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type createProfileModel struct {
	ID      string         `json:"id"`
	Filters []gorpo.Filter `json:"filters"`
}

type updateProfileModel struct {
	Filters []gorpo.Filter `json:"filters"`
}

func (m *createProfileModel) toProfile() (*gorpo.Profile, error) {
	if m.ID == "" {
		return nil, errors.EValidation("id is missing", nil)
	}

	if len(m.Filters) == 0 {
		return nil, errors.EValidation("filters are empty", nil)
	}

	return &gorpo.Profile{ID: m.ID, Filters: m.Filters}, nil
}

func (m *updateProfileModel) toProfile(profile *gorpo.Profile) (*gorpo.Profile, error) {
	if len(m.Filters) == 0 {
		return nil, errors.EValidation("effects are empty", nil)
	}

	profile.Filters = m.Filters

	return profile, nil
}
