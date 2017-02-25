package hex2rgba

import "testing"
import "image/color"
import "github.com/stretchr/testify/assert"

type in struct {
	hex   string
	alpha uint8
}

type out struct {
	col color.RGBA
	err string
}

type testpair struct {
	in  in
	out out
}

func TestConvert(t *testing.T) {
	pairs := []testpair{
		{in{"ffffff", 255}, out{color.RGBA{255, 255, 255, 255}, ""}},
		{in{"000000", 120}, out{color.RGBA{0, 0, 0, 120}, ""}},
		{in{"ccc", 255}, out{color.RGBA{204, 204, 204, 255}, ""}},
		{in{"ff", 255}, out{color.RGBA{}, "hex2rgba: invalid color hex string \"ff\""}},
		{in{"zffggg", 255}, out{color.RGBA{}, "hex2rgba: invalid color hex string \"zffggg\""}},
	}

	for _, tp := range pairs {
		col, err := Convert(tp.in.hex, tp.in.alpha)

		assert.Equal(t, tp.out.col, col)

		if tp.out.err != "" {
			assert.EqualError(t, err, tp.out.err)
		} else {
			assert.NoError(t, err)
		}
	}
}
