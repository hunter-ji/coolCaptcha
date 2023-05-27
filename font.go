package coolCaptcha

import (
	"github.com/fogleman/gg"
)

const fontPoints = 120

type fontConfig struct {
	Text  string
	X     float64
	Y     float64
	AX    float64
	AY    float64
	Color string
}

func (c *Config) setFontFace(dc *gg.Context) (err error) {
	// load font
	face, err := loadFontFace()
	if err != nil {
		return
	}

	dc.SetFontFace(face)
	return
}

func (c *Config) writeText(dc *gg.Context, font fontConfig) {
	dc.SetHexColor(font.Color)
	dc.DrawStringAnchored(font.Text, font.X, font.Y, font.AX, font.AY)
}
