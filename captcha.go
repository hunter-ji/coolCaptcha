package coolCaptcha

import (
	_ "embed"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"math/rand"
	"os"

	"github.com/samber/lo"
)

func (c *Config) GenerateImage() (imageBase64Data string, code string, err error) {
	err = c.checkConfig()
	if err != nil {
		return
	}

	code, codeItems, err := c.getLastCode()
	if err != nil {
		return
	}

	originImage, err := c.drawStaticImage(codeItems)
	if err != nil {
		return
	}

	imageBase64Data, err = convertImageToBase64(originImage)
	if err != nil {
		return
	}
	return
}

func (c *Config) GenerateGif() (code string, err error) {
	err = c.checkConfig()
	if err != nil {
		return
	}

	code, codeItems, err := c.getLastCode()
	if err != nil {
		return
	}

	fmt.Println("codeItems: ", codeItems)

	// gen random line config
	var lineConfigs []lineConfig
	for i := 0; i < 4; i++ {
		line := c.genLineCoordinates()
		line.Width = lineWidth()

		lineConfigs = append(lineConfigs, line)
	}

	var textConfigs []fontConfig
	for i := 0; i < charactersLength; i++ {
		textConfig := fontConfig{
			Character: codeItems[i],
			X:         float64(50 + i*70),
			Y:         randomFloat64(30, 70),
			AX:        randomFloat64(0.3, 0.7),
			AY:        randomFloat64(0.3, 0.7),
			Color:     "",
			IsUp:      rand.Intn(2) == 1,
		}

		textConfigs = append(textConfigs, textConfig)
	}

	imageCount := 40

	// gen random themes
	randomThemeIndex := rand.Perm(len(themes))
	for len(randomThemeIndex) < imageCount {
		randomThemeIndex = append(randomThemeIndex, rand.Perm(len(themes))...)
	}

	var images []image.Image
	var theme Theme
	for i := 0; i < imageCount; i++ {
		for lineIndex, line := range lineConfigs {
			// start
			lineConfigs[lineIndex].Start.Y = genRandomPoint(line.Start.Y, 15)

			// end
			lineConfigs[lineIndex].End.Y = genRandomPoint(line.End.Y, 15)

			// zigzag
			lineConfigs[lineIndex].Zigzag.X = genRandomPoint(line.Zigzag.X, 20)
			lineConfigs[lineIndex].Zigzag.Y = genRandomPoint(line.Zigzag.Y, 20)
		}

		for textIndex, text := range textConfigs {
			textConfigs[textIndex].X = genRandomPoint(text.X, 6)
			textConfigs[textIndex].Y = genRandomPoint(text.Y, 8)
		}

		if i%4 == 0 {
			theme = themes[randomThemeIndex[i]]
			fontColorIndex := rand.Perm(len(theme.FontHexColors))
			lineColorIndex := rand.Perm(len(theme.LineHexColors))

			for lineIndex := range lineConfigs {
				lineConfigs[lineIndex].Color = theme.LineHexColors[lineColorIndex[lineIndex]]
				lineConfigs[lineIndex].Color = themes[0].LineHexColors[lineIndex]
			}

			for textIndex, text := range textConfigs {
				// textConfigs[textIndex].Color = theme.FontHexColors[fontColorIndex[textIndex]]
				text.Color = theme.FontHexColors[fontColorIndex[textIndex]]
			}
		}

		originImage, err := c.drawGifImage("#f1f1f2", textConfigs, lineConfigs)
		if err != nil {
			break
		}

		images = append(images, originImage)
	}

	if err != nil {
		return
	}

	outGif := &gif.GIF{
		LoopCount: 0,
	}
	delay := 20
	for _, imageItem := range images {
		bounds := imageItem.Bounds()

		paletteImage := image.NewPaletted(bounds, palette.Plan9)
		draw.Draw(paletteImage, paletteImage.Rect, imageItem, bounds.Min, draw.Over)

		outGif.Image = append(outGif.Image, paletteImage)
		outGif.Delay = append(outGif.Delay, delay)
	}

	f, err := os.Create("./out.gif")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = f.Close() }()
	err = gif.EncodeAll(f, outGif)
	if err != nil {
		return
	}

	return
}

func genRandomPoint(originNum float64, coefficient float64) float64 {
	return lo.If(rand.Intn(2) == 1, originNum+(coefficient*rand.Float64())).Else(originNum - (coefficient * rand.Float64()))
}
