package httpapi

import "github.com/talento90/imgart/pkg/imgart"

type effectModel struct {
	ID          string            `json:"id"`
	Description string            `json:"description"`
	Parameters  imgart.Parameters `json:"parameters"`
}

func newEffectModel(e imgart.Effect) effectModel {
	return effectModel{
		ID:          e.ID(),
		Description: e.Description(),
		Parameters:  e.Parameters(),
	}
}
