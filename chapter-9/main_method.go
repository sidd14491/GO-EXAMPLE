package main

import "fmt"

type Rectangle struct {
	x1, y1 float64
}

func (r Rectangle) area() float64 {
	return r.x1 * r.y1
}

func main() {
	value := Rectangle{12.1, 13.13}
	fmt.Println(value.area())
}
