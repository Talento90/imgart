package effect

import (
	"context"
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/talento90/imgart/imgart"
)

type rotate struct {
	effect
}

// NewRotate creates an Effect that rotates an image
func NewRotate() imgart.Effect {
	return &rotate{
		effect: effect{
			id:          "rotate",
			description: "Rotate - rotates an image",
			parameters: imgart.Parameters{
				"angle": imgart.Parameter{
					Description: "Rotation angle in degreesº",
					Required:    true,
					Example:     -90,
					Type:        "integer",
				},
				"bgcolor": imgart.Parameter{
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

func (r *rotate) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	angle, err := floatBinder("angle", params)

	if err != nil {
		return nil, err
	}

	bgColor, err := colorBinder("bgcolor", params)

	if err != nil {
		bgColor = color.Transparent
	}

	img = imaging.Rotate(img, angle, bgColor)

	return img, ctx.Err()
}
