package main

import "fmt"

// Example for array
func main() {
	//Syntax for declaring a array and type is int
	var x [5]int
	x[3] = 100
	fmt.Println(x, x[3])

	//Example of array array type float64
	var z [5]float64
	z[0] = 10
	z[1] = 20
	z[2] = 40
	z[3] = 50
	z[4] = 61
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total = total + z[i]
	}
	fmt.Println(total / 5)

	// Using len() function to find out the length of array
	var sum float64 = 0
	for i := 0; i < len(z); i++ {
		sum = sum + z[i]
	}
	/* convert int to float64 syntax
	float64(len(z))
	*/
	fmt.Println(sum / float64(len(z)))

	/* Another type of syntax of for-loop
	for _,val := range z
	*/
	var total2 float64 = 0
	for _, val := range z {
		total2 += val
	}
	fmt.Println(total2 / float64(len(z)))

	/*Go also provide shorter syntax for declaring a array
	x := [5]float64{ 98, 93, 77, 82, 83 }
	*/
}
