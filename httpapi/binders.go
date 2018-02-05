package httpapi

import (
	"strconv"
)

func intBinder(value interface{}, defaultVal int) int {
	s, ok := value.(string)

	if !ok {
		return defaultVal
	}

	i, err := strconv.Atoi(s)

	if err != nil {
		return defaultVal
	}

	return i
}
