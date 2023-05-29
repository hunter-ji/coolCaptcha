package coolCaptcha

import (
	_ "embed"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"math/rand"
)

// GenerateImage
// @Description: Generate a static image
// @receiver c
// @return imageBase64Data
// @return code
// @return err
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

// GenerateGif
// @Description: Generate gif
// @receiver c
// @return gifBase64Data
// @return code
// @return err
func (c *Config) GenerateGif() (gifBase64Data string, code string, err error) {
	err = c.checkConfig()
	if err != nil {
		return
	}

	code, codeItems, err := c.getLastCode()
	if err != nil {
		return
	}

	// gen random line config
	var lineConfigs []lineConfig
	for i := 0; i < 4; i++ {
		line := c.genLineCoordinates()
		line.Width = c.lineWidth()

		lineConfigs = append(lineConfigs, line)
	}

	var textConfigs []fontConfig
	for i := 0; i < charactersLength; i++ {
		textConfig := fontConfig{
			Character: codeItems[i],
			X:         float64((c.Width / 6) + (i * c.Width / 4)),
			Y:         randomFloat64(float64(c.Height/4), float64(c.Height/3)),
			AX:        randomFloat64(0.3, 0.7),
			AY:        randomFloat64(0.3, 0.7),
			Color:     c.FontHexColor,
		}

		textConfigs = append(textConfigs, textConfig)
	}

	imageCount := 12

	randomLineIndex := rand.Perm(len(c.LineHexColors))

	var images []image.Image
	for i := 0; i < imageCount; i++ {
		for lineIndex, line := range lineConfigs {
			// start
			lineConfigs[lineIndex].Start.Y = genRandomPoint(line.Start.Y, 15)

			// end
			lineConfigs[lineIndex].End.Y = genRandomPoint(line.End.Y, 15)

			// zigzag
			lineConfigs[lineIndex].Zigzag.X = genRandomPoint(line.Zigzag.X, 20)
			lineConfigs[lineIndex].Zigzag.Y = genRandomPoint(line.Zigzag.Y, 20)

			lineConfigs[lineIndex].Color = c.LineHexColors[randomLineIndex[lineIndex]]
		}

		for textIndex, text := range textConfigs {
			textConfigs[textIndex].X = genRandomPoint(text.X, 6)
			textConfigs[textIndex].Y = genRandomPoint(text.Y, 6)
		}

		originImage, err := c.drawGifImage(c.BackgroundHexColor, textConfigs, lineConfigs)
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
	delay := 16
	for _, imageItem := range images {
		bounds := imageItem.Bounds()

		paletteImage := image.NewPaletted(bounds, palette.Plan9)
		draw.Draw(paletteImage, paletteImage.Rect, imageItem, bounds.Min, draw.Over)

		outGif.Image = append(outGif.Image, paletteImage)
		outGif.Delay = append(outGif.Delay, delay)
	}

	gifBase64Data, err = convertGifToBase64(outGif)
	if err != nil {
		return
	}

	if c.DevMode {
		c.devModelHandlerForGif(outGif)
	}

	return
}
