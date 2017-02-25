package image

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/bmp"

	"github.com/Fahrradflucht/last-supper/label"
)

// New creates a new image and encodes it in the given format
func New(width int, height int, bgcol color.RGBA, label label.ImageLabel) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{bgcol}, image.ZP, draw.Src)
	// label.Draw(img, label)
	return img
}

// Encode encodes a given image in the given format to the given
// writer
func Encode(w io.Writer, m image.Image, format string) {
	switch format {
	case "bmp":
		bmp.Encode(w, m)
	case "gif":
		// TODO: This is quite slow for bigger images.
		gif.Encode(w, m, &gif.Options{NumColors: 256})
	case "jpg", "jpeg":
		jpeg.Encode(w, m, &jpeg.Options{Quality: 100})
	case "png":
		png.Encode(w, m)
	}
}
