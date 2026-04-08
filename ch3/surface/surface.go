package surface

import (
	"fmt"
	"io"
	"math"
)

type Config struct {
	width, height    int      // configurable: canvas size in pixels
	cells            int      // configurable: number of grid cells
	xyrange          float64  // configurable: axis ranges
	angle            float64  // configurable: angle of x, y axes
	xyscale          float64  // derived: pixels per x or y unit
	zscale           float64  // derived: pixels per z unit
	sin              float64  // derived: sin(angle)
	cos              float64  // derived: cos(angle)
	red, green, blue uint8    // configurable: color values
	function         Function // configurable: the function to render
}

type Option func(*Config)

func WithHeight(height int) Option {
	return func(c *Config) {
		c.height = height
	}
}

func WithWidth(width int) Option {
	return func(c *Config) {
		c.width = width
	}
}

func WithCells(cells int) Option {
	return func(c *Config) {
		c.cells = cells
	}
}

func WithXYRange(xyrange float64) Option {
	return func(c *Config) {
		c.xyrange = xyrange
	}
}

func WithAngleDegree(angle float64) Option {
	return func(c *Config) {
		c.angle = angle * math.Pi / 180
	}
}

func WithRed(red uint8) Option {
	return func(c *Config) {
		c.red = red
	}
}

func WithGreen(green uint8) Option {
	return func(c *Config) {
		c.green = green
	}
}

func WithBlue(blue uint8) Option {
	return func(c *Config) {
		c.blue = blue
	}
}

func WithFunction(fn string) Option {
	return func(c *Config) {
		function := funcs[fn]
		if function != nil {
			c.function = function
			fmt.Printf("Setting function: %q\n", fn)
		} else {
			fmt.Printf("Function not found: %q\n", fn)
		}
	}
}

func NewConfig(opts ...Option) Config {
	cfg := Config{
		width:    600,
		height:   320,
		cells:    100,
		xyrange:  30.0,
		angle:    math.Pi / 6,
		red:      128,
		green:    128,
		blue:     128,
		function: f,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	cfg.xyscale = float64(cfg.width) / 2 / cfg.xyrange
	cfg.zscale = float64(cfg.height) * 0.4
	cfg.sin = math.Sin(cfg.angle)
	cfg.cos = math.Cos(cfg.angle)

	return cfg
}

// Creates an SVG rendering of a 3-D surface function.
func (cfg Config) CreateSVG(fn string, out io.Writer) error {
	header := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: #%02x%02x%02x; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>",
		cfg.red, cfg.green, cfg.blue,
		cfg.width, cfg.height)
	if _, err := fmt.Fprint(out, header); err != nil {
		return err
	}

	corner := func(i, j int) (float64, float64, bool) {
		x := cfg.xyrange * (float64(i)/float64(cfg.cells) - 0.5)
		y := cfg.xyrange * (float64(j)/float64(cfg.cells) - 0.5)
		z := cfg.function(x, y)
		if math.IsInf(z, 1) || math.IsInf(z, -1) || math.IsNaN(z) {
			return 0, 0, false
		}
		sx := float64(cfg.width)/2 + (x-y)*cfg.cos*cfg.xyscale
		sy := float64(cfg.height)/2 + (x+y)*cfg.sin*cfg.xyscale - z*cfg.zscale
		return sx, sy, true
	}

	for i := range cfg.cells {
		for j := range cfg.cells {
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)
			if ok1 && ok2 && ok3 && ok4 {
				polygon := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				if _, err := fmt.Fprint(out, polygon); err != nil {
					return err
				}
			}
		}
	}

	if _, err := fmt.Fprintf(out, "</svg>"); err != nil {
		return err
	}

	return nil
}
