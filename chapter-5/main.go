package main

import "fmt"

// Example for For,if ,if -- else, if --else if, switch
func main() {
	// 1st way to declare for loop
	i := 5
	fmt.Println("1st way to declare for loop")
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	// 2nd way
	fmt.Println(" way to declare for loop")
	for i := 11; i <= 15; i++ {
		fmt.Println(i)
	}

	// syntax for If and else inside loop
	i = 20
	for i <= 25 {
		if i%2 == 0 {
			fmt.Println("Even ", i)
		} else {
			fmt.Println("Odd", i)
		}
		i += 1
	}

	// syntax for If -- else if -- else inside loop
	for i := 30; i <= 40; i++ {
		if i%3 == 0 {
			fmt.Println("3 divide ", i)
		} else if i%4 == 0 {
			fmt.Println("4 divide", i)
		} else if i%5 == 0 {
			fmt.Println("5 Divided", i)
		}
	}

	// Switch example
	k := 2
	switch k {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println(i)
	}

}
