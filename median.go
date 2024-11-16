package main

import "sort"

func Median(input []float64) float64 {
	sort.Float64s(input)
	n := len(input)
	if n%2 == 1 {
		return input[n/2]
	}

	return input[n/2-1] + input[n/2]/2
}
