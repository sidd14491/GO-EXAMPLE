package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}
func zero1(yval int) int {
	yval = 10
	return yval
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
	y := 8
	op := zero1(y)
	fmt.Println(y, op)
}
