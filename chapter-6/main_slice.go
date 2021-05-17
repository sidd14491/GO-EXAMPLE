package main

import "fmt"

func main() {
	//syntax for declare a slice
	x := make([]float64, 5)
	fmt.Println(x, len(x))

	/* syntax for declare slice with mention capacity
	x := make([]float64, 5, 10)
	*/
	y := make([]float64, 5, 10)
	fmt.Println(y, len(y), cap(y))

	/*Another way to create slices is to use the [low : high]
	expression:
	arr := [5]float64{1,2,3,4,5}
	x := arr[0:5]
	*/
	arr := []float64{1, 2, 3, 4, 5}
	x = arr[0:3]
	fmt.Println(x)
	arr1 := make([]float64, 5)
	y = arr1[0:2]
	fmt.Println(y)

	/*Append function
	append(slice, elm1, elm2)
	append(slece, slice1)
	*/
	slice1 := []float64{1, 2, 4}
	slice2 := append(slice1, 5, 6)
	fmt.Println(slice2)

	/*COPY syntax
	copy(destination, source)
	*/
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	fmt.Println(slice3, slice4)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
