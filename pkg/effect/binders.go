package effect

import (
	"fmt"
	"image/color"

	"github.com/disintegration/imaging"

	"github.com/talento90/gorpo/pkg/gorpo"
)

var colorsList = []string{"black", "opaque", "transparent", "white"}
var colorMapping = map[string]color.Color{
	"black":       color.Black,
	"opaque":      color.Opaque,
	"transparent": color.Transparent,
	"white":       color.White,
}

var filtersList = []string{"lanczos", "catmull-rom", "mitchell-netravali", "bs-pline", "linear", "box", "nearest-neighbor"}
var filterMapping = map[string]imaging.ResampleFilter{
	"lanczos":            imaging.Lanczos,
	"catmull-rom":        imaging.CatmullRom,
	"mitchell-netravali": imaging.MitchellNetravali,
	"bs-pline":           imaging.BSpline,
	"linear":             imaging.Linear,
	"box":                imaging.Box,
	"nearest-neighbor":   imaging.NearestNeighbor,
}

func extractParameter(key string, params map[string]interface{}) (interface{}, error) {
	if value, ok := params[key]; ok {
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

func floatBinder(key string, params map[string]interface{}) (float64, error) {
	value, err := extractParameter(key, params)

	if err != nil {
		return 0, err
	}

	valueFloat, ok := value.(float64)

	if !ok {
		return 0, gorpo.EValidation(fmt.Sprintf("Parameter %s needs to be a float", key))
	}

	return valueFloat, nil
}

func colorBinder(key string, params map[string]interface{}) (color.Color, error) {
	value, err := extractParameter(key, params)

	if err != nil {
		return nil, err
	}

	colorKey, ok := value.(string)

	if !ok {
		return nil, gorpo.EValidation(fmt.Sprintf("Parameter %s needs to be a string", key))
	}

	color, ok := colorMapping[colorKey]

	if !ok {
		return nil, gorpo.EValidation(fmt.Sprintf("Value %s not supported", colorKey))
	}

	return color, nil
}

func filterBinder(key string, params map[string]interface{}) (imaging.ResampleFilter, error) {
	value, err := extractParameter(key, params)

	if err != nil {
		return imaging.ResampleFilter{}, err
	}

	filterKey, ok := value.(string)

	if !ok {
		return imaging.ResampleFilter{}, gorpo.EValidation(fmt.Sprintf("Parameter %s needs to be a string", key))
	}

	filter, ok := filterMapping[filterKey]

	if !ok {
		return imaging.ResampleFilter{}, gorpo.EValidation(fmt.Sprintf("Value %s not supported", filterKey))
	}

	return filter, nil
}
