package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
func main() {
	xv := []float64{12.12, 13.11, 14.12, 15.12, 16.6}
	fmt.Println(average(xv))
}
