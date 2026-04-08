package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

// Generates GIF animations of random Lissajous figures.
func Generate(out io.Writer, opts ...Option) {
	cfg := &Config{
		cycles:     5,
		resolution: 0.001,
		size:       100,
		nframes:    64,
		delay:      8,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	// Generate a palette containing a black background and one color for each
	// cycle.
	palette := make([]color.Color, cfg.cycles+1)
	palette[0] = color.Black
	for j := 1; j <= cfg.cycles; j++ {
		theta := 2 * math.Pi * float64(j-1) / float64(cfg.cycles)
		palette[j] = color.RGBA{
			R: uint8(128 + 127*math.Sin(theta)),
			G: uint8(128 + 127*math.Sin(theta+2*math.Pi/3)),
			B: uint8(128 + 127*math.Sin(theta+4*math.Pi/3)),
			A: 0xFF,
		}
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: cfg.nframes}
	phase := 0.0 // phase difference
	for range cfg.nframes {
		rect := image.Rect(0, 0, 2*cfg.size+1, 2*cfg.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cfg.cycles)*2*math.Pi; t += cfg.resolution {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				cfg.size+int(x*float64(cfg.size)+0.5),
				cfg.size+int(y*float64(cfg.size)+0.5),
				uint8(1+int(t/(2*math.Pi))),
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, cfg.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

type Config struct {
	cycles     int     // number of complete x oscillator revolutions
	resolution float64 // angular resolution
	size       int     // image canvas covers [-size...+size]
	nframes    int     // number of animation frames
	delay      int     // delay between frames in 10ms units
}

type Option func(*Config)

func WithCycles(cycles int) Option {
	return func(c *Config) {
		c.cycles = cycles
	}
}

func WithResolution(resolution float64) Option {
	return func(c *Config) {
		c.resolution = resolution
	}
}

func WithSize(size int) Option {
	return func(c *Config) {
		c.size = size
	}
}

func WithNFrames(nframes int) Option {
	return func(c *Config) {
		c.nframes = nframes
	}
}

func WithDelay(delay int) Option {
	return func(c *Config) {
		c.delay = delay
	}
}
