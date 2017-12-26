package effect

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

type rotate struct {
	descriptor Descriptor
}

// NewRotate creates an Effect that rotates an image
func NewRotate() Effect {
	return &rotate{
		descriptor: Descriptor{
			ID:          "rotate",
			Description: "Rotate - rotates an image",
			Parameters: Parameters{
				"angle": Parameter{
					Description: "Rotation angle in degreesº",
					Required:    true,
					Example:     -90,
					Type:        "integer",
				},
				"bgcolor": Parameter{
					Description: "Color of uncovered zones",
					Required:    false,
					Example:     "black",
					Type:        "string",
					Default:     "transparent",
					Values:      colorsList,
				},
			},
		},
	}
}

func (r *rotate) Descriptor() Descriptor {
	return r.descriptor
}

func (r *rotate) Transform(img image.Image, params map[string]interface{}) (image.Image, error) {
	angle, err := floatBinder("angle", params)

	if err != nil {
		return nil, err
	}

	bgColor, err := colorBinder("bgcolor", params)

	if err != nil {
		bgColor = color.Transparent
	}

	img = imaging.Rotate(img, angle, bgColor)

	return img, nil
}
