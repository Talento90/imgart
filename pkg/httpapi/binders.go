package httpapi

func intBinder(value interface{}, defaultVal int) int {
	i, ok := value.(int)

	if !ok {
		return defaultVal
	}

	return i
}
