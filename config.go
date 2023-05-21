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

func SetWidth(width int) Options {
	return func(c *Config) {
		c.Width = width
	}
}

func SetHeight(height int) Options {
	return func(c *Config) {
		c.Height = height
	}
}

func SetBackgroundHexColor(backgroundHexColor string) Options {
	return func(c *Config) {
		c.BackgroundHexColor = backgroundHexColor
	}
}

func SetFontHexColor(fontHexColor string) Options {
	return func(c *Config) {
		c.FontHexColor = fontHexColor
	}
}

func SetLineHexColors(lineHexColors []string) Options {
	return func(c *Config) {
		c.LineHexColors = lineHexColors
	}
}

func SetDevMode(devMode bool) Options {
	return func(c *Config) {
		c.DevMode = devMode
	}
}

func DefaultConfig(c *Config) *Config {
	c.Width = width
	c.Height = height
	c.BackgroundHexColor = backgroundHexColor
	c.FontHexColor = fontHexColor
	c.LineHexColors = lineHexColors
	return c
}

func New(options ...Options) *Config {
	c := &Config{}
	c = DefaultConfig(c)
	for _, op := range options {
		op(c)
	}
	return c
}
