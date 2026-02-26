package conv

import "fmt"

type Meter float64
type Foot float64
type Inch float64

const (
	MeterToFoot = 3.28084
	MeterToInch = 39.3701
)

func (m Meter) ToF() Foot {
	return Foot(m * MeterToFoot)
}

func (m Meter) ToI() Inch {
	return Inch(m * MeterToInch)
}

func (m Meter) String() string {
	return fmtLength(m) + " m"
}

func (f Foot) ToM() Meter {
	return Meter(f / MeterToFoot)
}

func (f Foot) ToI() Inch {
	return f.ToM().ToI()
}

func (f Foot) String() string {
	return fmtLength(f) + " feet"
}

func (i Inch) ToM() Meter {
	return Meter(i / MeterToInch)
}

func (i Inch) ToF() Foot {
	return i.ToM().ToF()
}

func (i Inch) String() string {
	return fmtLength(i) + " inch"
}

func fmtLength[L Meter | Foot | Inch](l L) string {
	return fmt.Sprintf("%.3f", l)
}
