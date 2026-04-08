package surface

import "math"

type Function func(x, y float64) float64

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggBox(x, y float64) float64 {
	normalize := func(x, a float64) float64 {
		x = math.Abs(x)
		x -= (float64(int64(x/a)) * a)
		x -= (a / 2)
		return x
	}

	r := math.Hypot(normalize(x, 5), normalize(y, 5))
	if r > 3 {
		r = 3
	}
	return r * r / 30
}

func moguls(x, y float64) float64 {
	k := 1.3
	return (math.Sin(x+k*y) + math.Sin(k*x-y)) / 40
}

func saddle(x, y float64) float64 {
	return (x*x - y*y + 0.1*x*y) / 600
}

var funcs = map[string]Function{
	"f":      f,
	"eggbox": eggBox,
	"moguls": moguls,
	"saddle": saddle,
}
