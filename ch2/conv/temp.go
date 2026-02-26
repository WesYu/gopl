package conv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC       Celsius    = -273.15
	FreezingC           Celsius    = 0
	BoilingC            Celsius    = 100
	AbsoluteZeroF       Fahrenheit = -459.67
	FreezingF           Fahrenheit = 32
	BoilingF            Fahrenheit = 212.0
	AbsoluteZeroK       Kelvin     = 0
	FreezingK           Kelvin     = 273.15
	BoilingK            Kelvin     = 373.15
	CelsiusToFahrenheit            = 9 / 5
)

func Boiling() {
	fmt.Printf("boiling point = %s or %s\n", BoilingF, BoilingF.ToC())
}

func (c Celsius) ToF() Fahrenheit {
	return Fahrenheit(c*CelsiusToFahrenheit) + FreezingF
}

func (c Celsius) ToK() Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func (c Celsius) String() string {
	return fmtTemp(c) + "°C"
}

func (f Fahrenheit) ToC() Celsius {
	return Celsius(f-FreezingF) / CelsiusToFahrenheit
}

func (f Fahrenheit) ToK() Kelvin {
	return f.ToC().ToK()
}

func (f Fahrenheit) String() string {
	return fmtTemp(f) + "°F"
}

func (k Kelvin) ToC() Celsius {
	return Celsius(k) + AbsoluteZeroC
}

func (k Kelvin) ToF() Fahrenheit {
	return k.ToC().ToF()
}

func (k Kelvin) String() string {
	return fmtTemp(k) + "K"
}

func fmtTemp[T Celsius | Fahrenheit | Kelvin](t T) string {
	return fmt.Sprintf("%.3f", t)
}
