package conv

import "fmt"

type Kilogram float64
type Pound float64
type Ounce float64

const (
	KiloToPound = 2.20462
	KiloToOunce = 35.274
)

func (k Kilogram) ToP() Pound {
	return Pound(k * KiloToPound)
}

func (k Kilogram) ToO() Ounce {
	return Ounce(k * KiloToOunce)
}

func (k Kilogram) String() string {
	return fmtWeight(k) + " kg"
}

func (p Pound) ToK() Kilogram {
	return Kilogram(p / KiloToPound)
}

func (p Pound) ToO() Ounce {
	return p.ToK().ToO()
}

func (p Pound) String() string {
	return fmtWeight(p) + " lbs"
}

func (o Ounce) ToK() Kilogram {
	return Kilogram(o / KiloToOunce)
}

func (o Ounce) ToP() Pound {
	return o.ToK().ToP()
}

func (o Ounce) String() string {
	return fmtWeight(o) + " oz"
}

func fmtWeight[W Kilogram | Pound | Ounce](w W) string {
	return fmt.Sprintf("%.3f", w)
}
