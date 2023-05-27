package coolCaptcha

import (
	_ "embed"
	"sync"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed blowbrush.ttf
var fontFile []byte

var (
	loadFontFaceOnce sync.Once
	fontFace         font.Face
	loadFontFaceErr  error
)

// loadFontFace
// @Description: load custom font file and return font face
// @return face
// @return err
func loadFontFace() (font.Face, error) {
	loadFontFaceOnce.Do(func() {
		f, err := truetype.Parse(fontFile)
		if err != nil {
			loadFontFaceErr = err
			return
		}

		fontFace = truetype.NewFace(f, &truetype.Options{
			Size: fontPoints,
		})
	})

	return fontFace, loadFontFaceErr
}
