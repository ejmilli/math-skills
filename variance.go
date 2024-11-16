package main

import "sort"

func Variance(input []float64) float64 {
	var sum floa
	sort.Float64s(input)
	n:= len(input)

	for _, char:= range input{
		sum += char
	}
	v:= sum/n
	updatenum:= input-v

	var sumt float64
	for _, char:= range updatenum{
		sumt+= char
		milli := Math.Pow(sumt, 2)
		g:= milli/n-1
	}

	return g

	}
