package coolCaptcha

type Theme struct {
	BackgroundHexColor string
	FontHexColors      []string
	LineHexColors      []string
}

var themes []Theme

func init() {
	themes = setThemes()
}

func setThemes() []Theme {
	return []Theme{
		{
			BackgroundHexColor: "#ec6a52",
			FontHexColors:      []string{"#f7dfd6", "#f0d1c3", "#ecae9b", "#ce755f"},
			LineHexColors:      []string{"#93aec1", "#9dbdba", "#f8b042", "#f3b7ad"},
		},
		{
			BackgroundHexColor: "#c4e1f6",
			FontHexColors:      []string{"#f7dfd6", "#f0d1c3", "#ecae9b", "#ce755f"},
			LineHexColors:      []string{"#0b9b8a", "#f596a1", "#fadeeb", "#f9c975"},
		},
		{
			BackgroundHexColor: "#b63231",
			FontHexColors:      []string{"#e5e4f0", "#d5d3e5", "#f0afcd", "#ea7da2"},
			LineHexColors:      []string{"#d6bfba", "#e1b1ab", "#da6c64", "#afaea4"},
		},
	}
}
