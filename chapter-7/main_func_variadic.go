/* Variadic function
	By using ... before the type name of the last parame-
ter you can indicate that it takes zero or more of those
parameters.
*/

package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	fmt.Println(add(1, 2, 3))
}
