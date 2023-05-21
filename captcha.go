package coolCaptcha

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/fogleman/gg"
	"github.com/samber/lo"
)

func (c *Config) CustomCode(code string) *Config {
	c.Code = code
	return c
}

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

func lineWidth() float64 {
	return randomFloat64(8, 10)
}

func (c *Config) Generate() (imageBase64Data string, code string, err error) {
	if len(c.LineHexColors) < 3 {
		err = errors.New("lineHexColors requires at least three values")
		return
	}

	var codeItems []string
	configCode := strings.TrimSpace(c.Code)
	if configCode != "" && len(configCode) != 4 {
		err = errors.New("the expected length of customCode is 4")
		return
	}

	if configCode == "" {
		codeItems = getRandomCodeItems()
	} else {
		codeItems = strings.Split(strings.ToUpper(configCode), "")
	}

	dc := gg.NewContext(c.Width, c.Height)
	dc.SetHexColor(c.BackgroundHexColor)
	dc.Clear()

	err = dc.LoadFontFace("./blowbrush.ttf", 120)
	if err != nil {
		return
	}

	// write text
	for index, text := range codeItems {
		dc.SetHexColor(c.FontHexColor)
		dc.DrawStringAnchored(text, float64(50+index*70), 50, randomFloat64(0.3, 0.7), randomFloat64(0.3, 0.7))

		if index < charactersLength-1 {
			colorsLength := len(c.LineHexColors)
			index := rand.Intn(colorsLength)
			if index < colorsLength {
				c.setLine(dc, lineWidth(), c.LineHexColors[index])
				c.LineHexColors = lo.Filter(c.LineHexColors, func(color string, colorIndex int) bool {
					return color != c.LineHexColors[index]
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
