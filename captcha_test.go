package coolCaptcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptcha(t *testing.T) {
	// error case
	// lineHexColors requires at least three values
	options := []Options{
		SetLineHexColors([]string{"#f596a1", "#fadeeb"}),
	}
	_, code, err := New(options...).Generate()
	assert.Error(t, err)
	assert.Len(t, code, 0)

	// error case
	// the expected length of customCode is 4
	_, code, err = New().CustomCode("coo").Generate()
	assert.Error(t, err)
	assert.Len(t, code, 0)

	// success case
	// call directly, no parameters required
	_, code, err = New().Generate()
	assert.Nil(t, err)
	assert.Len(t, code, 4)

	// success case
	// configure all parameters and then call
	options = []Options{
		SetBackgroundHexColor("#c4e1f6"),
		SetFontHexColor("#312E2E"),
		SetLineHexColors([]string{"#f596a1", "#fadeeb", "#f9c975"}),
		SetWidth(300),
		SetHeight(120),
		SetCodeType(MixedCharacters),
		SetDevMode(true),
	}
	_, code, err = New(options...).CustomCode("cool").Generate()
	assert.Nil(t, err)
	assert.Equal(t, code, "COOL")

	_, code, err = New(options...).Generate()
	assert.Nil(t, err)
	assert.Len(t, code, 4)
}
