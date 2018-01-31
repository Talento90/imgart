package effect

import "github.com/talento90/imgart/pkg/imgart"

type effect struct {
	id          string
	description string
	parameters  imgart.Parameters
}

func (e *effect) ID() string {
	return e.id
}

func (e *effect) Description() string {
	return e.description
}

func (e *effect) Parameters() imgart.Parameters {
	return e.parameters
}
