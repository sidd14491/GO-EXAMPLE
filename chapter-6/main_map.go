/*A map is an unordered collection of key-value pairs.
Also known as an associative array, a hash table or a
dictionary, maps are used to look up a value by its as-
sociated key. Here's an example of a map in Go:

	var x map[string]int
*/
package main

import "fmt"

func main() {
	//Create a map key with string
	//var x map[string]int
	x := make(map[string]int)
	x["age"] = 10
	fmt.Println(x[`age`])
	//create a map key with integr
	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1])

	element := make(map[string]string)
	element["N"] = "Nitrogen"
	element["H"] = "Hydrogen"
	element["O"] = "Oxygen"
	element["Ne"] = "Neon"
	fmt.Println(element, len(element))
	//Verify map is return value or not in this case it's return True
	if name, ok := element["O"]; ok {
		fmt.Println(name, ok)
	}
	///Verify map is return value or not in this case it's return False
	if name, ok := element["OP"]; ok {
		fmt.Println(name, ok)
	}

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
	}
	fmt.Println(elements["H"])
}
