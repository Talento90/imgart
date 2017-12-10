package effect

import (
	"fmt"

	"github.com/talento90/gorpo/pkg/gorpo"
)

func extractParameter(key string, params map[string]interface{}) (interface{}, error) {
	if value, ok := params["width"]; ok {
		return value, nil
	}

	return nil, gorpo.EValidation(fmt.Sprintf("Parameter %s required", key))
}

func integerBinder(key string, params map[string]interface{}) (int, error) {
	value, err := extractParameter(key, params)

	if err != nil {
		return 0, err
	}

	valueInt, ok := value.(int)

	if !ok {
		return 0, gorpo.EValidation(fmt.Sprintf("Parameter %s needs to be an integer", key))
	}

	return valueInt, nil
}
