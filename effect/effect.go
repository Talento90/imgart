package effect

import "github.com/talento90/gorpo/gorpo"

// Descriptor struct has a detailed description about the effect
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
