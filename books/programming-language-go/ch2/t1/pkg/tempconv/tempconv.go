package tempconv

import "fmt"

type Celsius float64
type Kelvins float64

const (
	AbsoluteZeroC Celsius = 273.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (k Kelvins) String() string {
	return fmt.Sprintf("%g°K", k)
}
