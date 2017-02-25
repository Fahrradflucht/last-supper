package hex2rgba

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
	"strings"
)

// Convert converts a color hex string (like #fffff)
// to a color.RGBA color.
func Convert(hex string, alpha uint8) (color.RGBA, error) {
	validHexString := regexp.MustCompile("^([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")
	if !validHexString.MatchString(hex) {
		return color.RGBA{}, fmt.Errorf("hex2rgba: invalid color hex string %q", hex)
	}

	var r, g, b uint64
	switch len(hex) {
	case 3:
		r, _ = strconv.ParseUint(strings.Repeat(hex[0:1], 2), 16, 0)
		g, _ = strconv.ParseUint(strings.Repeat(hex[1:2], 2), 16, 0)
		b, _ = strconv.ParseUint(strings.Repeat(hex[2:3], 2), 16, 0)
	case 6:
		r, _ = strconv.ParseUint(hex[0:2], 16, 0)
		g, _ = strconv.ParseUint(hex[2:4], 16, 0)
		b, _ = strconv.ParseUint(hex[4:6], 16, 0)
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), alpha}, nil
}
