package coolCaptcha

import (
	_ "embed"
	"errors"
	"math/rand"
	"strings"

	"github.com/fogleman/gg"
	"github.com/samber/lo"
)

func (c *Config) setLine(dc *gg.Context, lineWidth float64, color string) {
	width := float64(c.Width)
	height := float64(c.Height)

	startX := 0.0
	startY := rand.Float64()*height/2 + height/2
	endX := width
	endY := rand.Float64() * height / 2

	x0 := rand.Float64() * width / 2
	y0 := rand.Float64() * height

	dc.MoveTo(startX, startY)
	dc.QuadraticTo(x0, y0, endX, endY)
	dc.SetHexColor(color)
	dc.SetLineWidth(lineWidth)
	dc.Stroke()
}

// lineWidth
// @Description: set random line width
// @return float64
func lineWidth() float64 {
	return randomFloat64(8, 10)
}

// Generate
// @Description: The set parameters are drawn into an image, then the base64 data and code are returned
// @receiver c
// @return imageBase64Data: The base64 data of the graphic captcha can generate an image on the front end
// @return code: Randomly generated characters that are compared to the verification code entered by the user. When custom code is used, uppercase code is output.
// @return err
func (c *Config) Generate() (imageBase64Data string, code string, err error) {
	if len(c.LineHexColors) < 3 {
		err = errors.New("lineHexColors requires at least three values")
		return
	}

	var codeItems []string
	configCode := strings.TrimSpace(c.Code)

	// When the user does not use the custom code,
	// random characters will be generated according to the codeType set by the user.
	if configCode == "" {
		codeItems, err = c.getRandomCodeItems()
		if err != nil {
			return
		}
	} else {
		err = checkCustomCodeFormat(configCode)
		if err != nil {
			return
		}

		codeItems = strings.Split(strings.ToUpper(configCode), "")
	}

	// create a new image
	dc := gg.NewContext(c.Width, c.Height)
	dc.SetHexColor(c.BackgroundHexColor)
	dc.Clear()

	// load font
	face, err := loadFontFace()
	dc.SetFontFace(face)

	// write random code and set lines
	for index, text := range codeItems {
		dc.SetHexColor(c.FontHexColor)
		dc.DrawStringAnchored(text, float64(50+index*70), 50, randomFloat64(0.3, 0.7), randomFloat64(0.3, 0.7))

		// set 3 lines with random color
		if index < charactersLength-1 {
			colorsLength := len(c.LineHexColors)
			index := rand.Intn(colorsLength)
			tmpLineHexColors := c.LineHexColors
			if index < colorsLength {
				c.setLine(dc, lineWidth(), tmpLineHexColors[index])
				tmpLineHexColors = lo.Filter(tmpLineHexColors, func(color string, colorIndex int) bool {
					return color != tmpLineHexColors[index]
				})
			}
		}
	}

	imageBase64Data, err = convertImageToBase64(dc.Image())
	if err != nil {
		return
	}

	code = strings.Join(codeItems, "")

	if c.DevMode {
		c.devModeHandler(dc)
	}

	return
}
