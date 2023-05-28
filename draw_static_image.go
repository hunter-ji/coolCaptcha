package coolCaptcha

import (
	"image"
	"math/rand"

	"github.com/fogleman/gg"
)

// drawStaticImage
// @Description: The set parameters are drawn into an image, then the base64 data and code are returned
// @receiver c
// @return imageBase64Data: The base64 data of the graphic captcha can generate an image on the front end
// @return code: Randomly generated characters that are compared to the verification code entered by the user. When custom code is used, uppercase code is output.
// @return err
func (c *Config) drawStaticImage(codeItems []string) (imageOriginData image.Image, err error) {
	// create a new image
	dc := gg.NewContext(c.Width, c.Height)
	dc.SetHexColor(c.BackgroundHexColor)
	dc.Clear()

	// load font
	face, err := loadFontFace()
	dc.SetFontFace(face)

	// write random code and set lines
	randomColorIndex := rand.Perm(len(c.LineHexColors))
	for index, character := range codeItems {
		c.writeText(dc, fontConfig{character, float64(50 + index*70), 50,
			randomFloat64(0.3, 0.7), randomFloat64(0.3, 0.7), c.FontHexColor, true})

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
