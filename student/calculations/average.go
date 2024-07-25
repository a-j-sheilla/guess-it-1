package calculations

func Average(value []float64) float64{
	x := len(value)
	var sum float64
	for _, num := range value{
		sum = sum + num
		//return sum
	}
	ave := sum / float64(x)
	return ave
}