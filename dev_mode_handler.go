package coolCaptcha

import (
	"fmt"
	"os"
	"time"

	"github.com/fogleman/gg"
)

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
