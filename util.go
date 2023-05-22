package coolCaptcha

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/png"
	"math/rand"
	"regexp"
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

// CheckCustomCodeFormat
// @Description: Detect the format of the code
// @param code
// @return err
func checkCustomCodeFormat(code string) (err error) {
	reg := `^[a-zA-Z\d]{4}$`
	rgx := regexp.MustCompile(reg)
	if !rgx.MatchString(code) {
		err = errors.New("the custom code is malformed, only English letters and numbers are supported, and the length is 4")
		return
	}
	return
}
