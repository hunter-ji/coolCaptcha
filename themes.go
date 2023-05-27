// @Title themes.go
// @Description
// @Author 红尘 2023/5/27 21:18

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
			BackgroundHexColor: "#c4e1f6",
			FontHexColors:      []string{"#f7dfd6", "#f0d1c3", "#ecae9b", "#ce755f"},
			LineHexColors:      []string{"#0b9b8a", "#f596a1", "#fadeeb", "#f9c975"},
		},
		{
			BackgroundHexColor: "#f1f1f2",
			FontHexColors:      []string{"#ffc0cc", "#dbd2cc", "#c0b5b9", "#f1d8db"},
			LineHexColors:      []string{"#f6d9e1", "#dceef8", "#fcdda3", "#fbf2d2"},
		},
		{
			BackgroundHexColor: "#b63231",
			FontHexColors:      []string{"#e5e4f0", "#d5d3e5", "#f0afcd", "#ea7da2"},
			LineHexColors:      []string{"#d6bfba", "#e1b1ab", "#da6c64", "#afaea4"},
		},
		{
			BackgroundHexColor: "#f1b3cd",
			FontHexColors:      []string{"#315098", "#8ca8be", "#afbdb0", "#e4c0be"},
			LineHexColors:      []string{"#fdf4f3", "#f8d6dd", "#f7d3e4", "#f0c2a8"},
		},
		{
			BackgroundHexColor: "#86abba",
			FontHexColors:      []string{"#deebec", "#bed9dd", "#aecfd0", "#73b3b2"},
			LineHexColors:      []string{"#8b86be", "#deb0bd", "#ecb761", "#cbd690"},
		},
		{
			BackgroundHexColor: "#fcd97d",
			FontHexColors:      []string{"#e1eff5", "#c6e4e8", "#abdad9", "#ecd1e2"},
			LineHexColors:      []string{"#ece2eb", "#f4e0ea", "#efc9d6", "#efb1c7"},
		},
		{
			BackgroundHexColor: "#b5d5d0",
			FontHexColors:      []string{"#c34a5c", "#eecbd1", "#29425c", "#9ea544"},
			LineHexColors:      []string{"#f7f9ed", "#f3f7e6", "#dbeadd", "#c4dcd3"},
		},
		{
			BackgroundHexColor: "#ebcda7",
			FontHexColors:      []string{"#a52e45", "#2b5278", "#61787b", "#bf5c45"},
			LineHexColors:      []string{"#f1cce1", "#dfabcf", "#8fd2e6", "#5d7ab5"},
		},
	}
}
