package main

import "fmt"

func main() {
	// integer type
	fmt.Println("1 + 1 =", 1+1)
	//flotating type
	fmt.Println("1 + 1 =", 1.0+1.0)
	/* Booleans type example
	three logical operator used with boolean value
	a> &&(and) b> || (or) c> ! (not)
	*/
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
