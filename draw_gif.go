package coolCaptcha

import (
	"image"
	"math/rand"

	"github.com/fogleman/gg"
)

func (c *Config) drawGifImage(texts []fontConfig, lines []lineConfig) (imageOriginData image.Image, err error) {
	// create a new image
	dc := gg.NewContext(c.Width, c.Height)
	dc.SetHexColor(c.BackgroundHexColor)
	dc.Clear()

	err = c.setFontFace(dc)
	if err != nil {
		return
	}

	// write random code and set lines
	randomColorIndex := rand.Perm(len(c.LineHexColors))
	for index, character := range texts {
		c.writeText(dc, character)

		// set 3 lines with random color
		if index < charactersLength-1 {
			c.setStaticLine(dc, c.LineHexColors[randomColorIndex[index]])
		}
	}

	imageOriginData = dc.Image()

	if c.DevMode {
		c.devModeHandler(dc)
	}

	return
}
