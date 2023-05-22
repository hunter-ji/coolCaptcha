package coolCaptcha

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"math/rand"
)

// randomFloat64
// @Description: Take a random number from the smallest and largest ranges
// @param min
// @param max
// @return float64
func randomFloat64(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

// convertImageToBase64
// @Description: convert image.Image to image base64 data
// @param imageData
// @return imageBase64Data
// @return err
func convertImageToBase64(imageData image.Image) (imageBase64Data string, err error) {
	writer := new(bytes.Buffer)
	err = png.Encode(writer, imageData)
	if err != nil {
		return
	}

	imageBase64Data = "data:image/png;base64," + base64.StdEncoding.EncodeToString(writer.Bytes())
	return
}
