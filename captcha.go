package coolCaptcha

import (
	_ "embed"
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

	return
}
