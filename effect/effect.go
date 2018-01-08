package effect

import "github.com/talento90/gorpo/gorpo"

type effect struct {
	id          string
	description string
	parameters  gorpo.Parameters
}

func (e *effect) ID() string {
	return e.id
}

func (e *effect) Description() string {
	return e.description
}

func (e *effect) Parameters() gorpo.Parameters {
	return e.parameters
}
