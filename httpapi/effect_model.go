package httpapi

import "github.com/talento90/gorpo/gorpo"

type effectModel struct {
	ID          string           `json:"id"`
	Description string           `json:"description"`
	Parameters  gorpo.Parameters `json:"parameters"`
}

func newEffectModel(e gorpo.Effect) effectModel {
	return effectModel{
		ID:          e.ID(),
		Description: e.Description(),
		Parameters:  e.Parameters(),
	}
}
