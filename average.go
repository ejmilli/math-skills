package main

// Average calculates the average of a slice of float64 numbers
func Average(s []float64) float64 {
	var sum float64
	for _, num := range s {
		sum += num
	}
	return sum / float64(len(s))
}