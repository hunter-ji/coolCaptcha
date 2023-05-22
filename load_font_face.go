package coolCaptcha

import (
	_ "embed"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed blowbrush.ttf
var fontFile []byte

const fontPoints = 120

// loadFontFace
// @Description: load custom font file and return font face
// @return face
// @return err
func loadFontFace() (face font.Face, err error) {
	f, err := truetype.Parse(fontFile)
	if err != nil {
		return
	}

	face = truetype.NewFace(f, &truetype.Options{
		Size: fontPoints,
	})
	return
}
