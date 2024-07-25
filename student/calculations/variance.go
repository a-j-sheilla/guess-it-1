package calculations

func Variance(value []float64) float64 {
	x := float64(len(value))
	ave := Average(value)
	// return ave
	var sumdif float64
	for i := range value {
		// var dif float64
		dif := value[i] - ave
		sumdif += dif * dif
	}
	variance := sumdif / x
	return variance
}
