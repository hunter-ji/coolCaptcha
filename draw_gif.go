package coolCaptcha

import (
	"image"

	"github.com/fogleman/gg"
)

func (c *Config) drawGifImage(backgroundHexColor string, texts []fontConfig, lines []lineConfig) (imageOriginData image.Image, err error) {
	// create a new image
	dc := gg.NewContext(c.Width, c.Height)
	dc.SetHexColor(backgroundHexColor)
	dc.Clear()

	err = c.setFontFace(dc)
	if err != nil {
		return
	}

	// write random code and set lines
	for index, character := range texts {
		c.writeText(dc, character)

		c.drawLine(dc, lines[index])
	}

	imageOriginData = dc.Image()
	return
}
