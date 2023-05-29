package coolCaptcha

import (
	"errors"
	"math/rand"
	"strings"
)

const (
	uppercaseEnglishCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericCharacters          = "0123456789"
	charactersLength           = 4
)

// getRandomCodeItems
// @Description: Random characters will be generated according to the codeType set by the user
// @receiver c
// @return result
// @return err
func (c *Config) getRandomCodeItems() (codeItems []string, err error) {
	var characters string

	switch c.CodeType {
	case UppercaseEnglishCharacters:
		characters = uppercaseEnglishCharacters
	case NumericCharacters:
		characters = numericCharacters
	case MixedCharacters:
		characters = numericCharacters + uppercaseEnglishCharacters
	default:
		err = errors.New("for unknown codeType parameters, please use one of the built-in types: UppercaseEnglishCharacters, NumericCharacters, MixedCharacters. Please read the documentation for more information")
	}

	if err != nil {
		return
	}

	for i := 0; i < charactersLength; i++ {
		codeItems = append(codeItems, string(characters[rand.Intn(len(characters))]))
	}
	return
}

func (c *Config) getLastCode() (code string, codeItems []string, err error) {
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

	code = strings.Join(codeItems, "")
	return
}
