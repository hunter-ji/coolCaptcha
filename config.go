package coolCaptcha

type Config struct {
	Width              int
	Height             int
	BackgroundHexColor string
	FontHexColor       string
	LineHexColors      []string
	Code               string
	DevMode            bool
}

var (
	width              = 300
	height             = 120
	backgroundHexColor = "#ec6a52"
	fontHexColor       = "#312E2E"
	lineHexColors      = []string{"#93aec1", "#9dbdba", "#f8b042", "#f3b7ad"}
)

type Options func(*Config)

// SetWidth
// @Description: Set the width of the captcha image, which is 300 by default
// @param width
// @return Options
func SetWidth(width int) Options {
	return func(c *Config) {
		c.Width = width
	}
}

// SetHeight
// @Description: Set the height of the captcha image, which is 120 by default
// @param height
// @return Options
func SetHeight(height int) Options {
	return func(c *Config) {
		c.Height = height
	}
}

// SetBackgroundHexColor
// @Description: Set the background color of the captcha image
// @param backgroundHexColor: Background color, only hex can be used
// @return Options
func SetBackgroundHexColor(backgroundHexColor string) Options {
	return func(c *Config) {
		c.BackgroundHexColor = backgroundHexColor
	}
}

// SetFontHexColor
// @Description: Set the font color of the captcha image
// @param fontHexColor: Font color, only hex can be used
// @return Options
func SetFontHexColor(fontHexColor string) Options {
	return func(c *Config) {
		c.FontHexColor = fontHexColor
	}
}

// SetLineHexColors
// @Description: Set the line color of the captcha image, a minimum of 3 colors need to be set, and the line will randomly get 3 colors from them to draw
// @param lineHexColors: Font colors, only hex can be used, a minimum of 3 colors need to be set
// @return Options
func SetLineHexColors(lineHexColors []string) Options {
	return func(c *Config) {
		c.LineHexColors = lineHexColors
	}
}

// SetDevMode
// @Description: In the development mode, the generated image is saved as a local file for easy viewing// @param devMode
// @return Options
func SetDevMode(devMode bool) Options {
	return func(c *Config) {
		c.DevMode = devMode
	}
}

func defaultConfig(c *Config) *Config {
	c.Width = width
	c.Height = height
	c.BackgroundHexColor = backgroundHexColor
	c.FontHexColor = fontHexColor
	c.LineHexColors = lineHexColors
	return c
}

func New(options ...Options) *Config {
	c := &Config{}
	c = defaultConfig(c)
	for _, op := range options {
		op(c)
	}
	return c
}

// CustomCode
// @Description: Users can use their own generated characters as verification codes
// @receiver c
// @param code: Customize the generated verification code, but the length limit of 4 digits needs to be met
// @return *Config
func (c *Config) CustomCode(code string) *Config {
	c.Code = code
	return c
}
