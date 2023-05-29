package coolCaptcha

import (
	"fmt"
	"image/gif"
	"os"
	"time"

	"github.com/fogleman/gg"
)

// devModeHandler
// @Description: In the development mode, the generated image is saved as a local file for easy viewing
// @receiver c
// @param dc
func (c *Config) devModeHandler(dc *gg.Context) {
	outputFolderPath := "output"
	if _, err := os.Stat(outputFolderPath); os.IsNotExist(err) {
		err := os.Mkdir(outputFolderPath, 0700)
		if err != nil {
			panic(err)
		}
	}

	outputFilePath := fmt.Sprintf("%s/%s.png", outputFolderPath, time.Now().Format("2006-01-02 15:04:05"))
	if saveErr := dc.SavePNG(outputFilePath); saveErr != nil {
		panic(saveErr)
	}

}

// devModelHandlerForGif
// @Description: In the development mode, the generated image is saved as a local file for easy viewing
// @receiver c
// @param gifData
func (c *Config) devModelHandlerForGif(gifData *gif.GIF) {
	outputFolderPath := "output"
	if _, err := os.Stat(outputFolderPath); os.IsNotExist(err) {
		err := os.Mkdir(outputFolderPath, 0700)
		if err != nil {
			panic(err)
		}
	}

	outputFilePath := fmt.Sprintf("%s/%s.gif", outputFolderPath, time.Now().Format("2006-01-02 15:04:05"))

	f, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = f.Close() }()
	err = gif.EncodeAll(f, gifData)
	if err != nil {
		return
	}

}
