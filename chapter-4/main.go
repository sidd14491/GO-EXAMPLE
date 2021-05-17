package main

import "fmt"

// Examples for assgining variable
func main() {
	// 1st way to declare and assigne value to variable
	var x string = "Hello World"
	// 2nd way firsh declare a variable and then assign value to variable
	var y string
	y = "Hello India"
	fmt.Println(x)
	fmt.Println(y)

	// Value of variable can change throughouth the lifetime of program
	var z string
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)

	var op string
	op = "Hello "
	fmt.Println(op)
	// 1st way
	op = op + "World "
	fmt.Println(op)
	//2nd way
	op += "India"
	fmt.Println(op)

	// In go we use "==" symbol of equality
	var out string = "Hello"
	var out1 string = "World"
	fmt.Println(out == out1)

	/* Go also support shorter type to declare a variable
	value := "Hello"
	var value = "Hello World"
	*/
	dog1_var := "Max"
	fmt.Println("My dog name is", dog1_var)
	var dog2_var = "Tommy"
	fmt.Println("My dog name is", dog2_var)

	// Syntax constant variable
	const pi = 3.14
	fmt.Println(pi)

	// Syntax for define multiple variable
	var (
		a = 10
		b = 12
		c = 13
	)
	fmt.Println(a, b, c)
}
