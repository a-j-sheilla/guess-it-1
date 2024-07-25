package calculations

import (
	"math"
)

func StdDev(value []float64) float64 {
	var_ := Variance(value)
	dev := math.Sqrt(var_)
	return dev
}
