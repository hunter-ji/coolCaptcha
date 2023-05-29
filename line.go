package coolCaptcha

import (
	"math/rand"

	"github.com/fogleman/gg"
)

type lineConfigItem struct {
	X    float64
	Y    float64
	IsUp bool
}

type lineConfig struct {
	Start  lineConfigItem
	End    lineConfigItem
	Zigzag lineConfigItem
	Width  float64
	Color  string
}

func (c *Config) genLineCoordinates() (line lineConfig) {
	width := float64(c.Width)
	height := float64(c.Height)

	line.Start.X = 0.0
	line.Start.Y = rand.Float64()*height/2 + height/2

	line.End.X = width
	line.End.Y = rand.Float64() * height / 2

	line.Zigzag.X = rand.Float64() * width / 2
	line.Zigzag.Y = rand.Float64() * height

	line.Start.IsUp = rand.Intn(2) == 1
	line.End.IsUp = rand.Intn(2) == 1
	line.Zigzag.IsUp = rand.Intn(2) == 1

	return
}

// lineWidth
// @Description: set random line width
// @return float64
func (c *Config) lineWidth() float64 {
	return randomFloat64(float64(c.Height/12), float64(c.Height/10))
}

func (c *Config) drawLine(dc *gg.Context, line lineConfig) {
	dc.MoveTo(line.Start.X, line.Start.Y)
	dc.QuadraticTo(line.Zigzag.X, line.Zigzag.Y, line.End.X, line.End.Y)
	dc.SetHexColor(line.Color)
	dc.SetLineWidth(line.Width)
	dc.Stroke()
}

func (c *Config) setStaticLine(dc *gg.Context, color string) {
	line := c.genLineCoordinates()
	line.Width = c.lineWidth()
	line.Color = color
	c.drawLine(dc, line)
}
