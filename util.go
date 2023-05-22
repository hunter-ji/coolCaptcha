package coolCaptcha

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"math/rand"
)

const (
	characters       = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charactersLength = 4
)

func getRandomCodeItems() (result []string) {
	for i := 0; i < charactersLength; i++ {
		result = append(result, string(characters[rand.Intn(len(characters))]))
	}
	return
}

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
