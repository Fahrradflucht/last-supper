package label

import (
	"image/color"
)

// ImageLabel to draw on a last-supper image.
type ImageLabel struct {
	Text  string
	Color color.RGBA
}
