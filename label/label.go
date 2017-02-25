package label

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
)

var (
	dpi = float64(72)
)

// ImageLabel to draw on a last-supper image.
type ImageLabel struct {
	Text        string
	Color       color.RGBA
	MaxFontsize float64
}

// Draw draws the given label in the biggest possible
// way to the given image.
func Draw(img *image.RGBA, label ImageLabel) {
	// Read the font data. (This should probably happen in main)
	fontBytes, err := ioutil.ReadFile("./Roboto-Regular.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	imgBounds := img.Bounds()

	label.MaxFontsize = float64(imgBounds.Dy())

	textImage := textBox(label, f, imgBounds.Dx())
	textImageBounds := textImage.Bounds()
	w, h := textImageBounds.Dx(), textImageBounds.Dy()
	x := imgBounds.Dx()/2 - w/2
	y := imgBounds.Dy()/2 - h/2
	draw.DrawMask(img, image.Rect(x, y, x+w, y+h), image.NewUniform(label.Color), image.ZP, textImage, textImageBounds.Min, draw.Over)
}

func textBox(label ImageLabel, f *truetype.Font, maxWidth int) image.Image {
	bg := image.NewUniform(color.Alpha{0})
	fg := image.NewUniform(label.Color)
	width := textWidth(label.Text, f, label.MaxFontsize)
	for width > maxWidth {
		label.MaxFontsize *= 0.9
		width = textWidth(label.Text, f, label.MaxFontsize)
	}
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(f)
	c.SetFontSize(label.MaxFontsize)
	boundingBox := f.Bounds(c.PointToFixed(label.MaxFontsize))
	bbDelta := boundingBox.Max.Sub(boundingBox.Min)
	height := int(bbDelta.Y+32) >> 6

	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(canvas, canvas.Bounds(), bg, image.ZP, draw.Src)
	c.SetDst(canvas)
	c.SetSrc(fg)
	c.SetClip(canvas.Bounds())
	// Draw the text.
	extent, err := c.DrawString(label.Text, fixed.Point26_6{X: 0, Y: boundingBox.Max.Y})
	if err != nil {
		log.Println(err)
		return nil
	}

	return canvas.SubImage(image.Rect(0, 0, int(extent.X)>>6, height))
}

func textWidth(s string, f *truetype.Font, fontsize float64) int {
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(f)
	c.SetFontSize(fontsize)

	width, err := c.DrawString(s, freetype.Pt(0, 0))
	if err != nil {
		log.Println(err)
		return len(s)
	}
	return int(width.X+32)>>6 + 1
}
