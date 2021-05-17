package main

import "fmt"

type Circle struct {
	x float64
	y float64
}

type Rectangle struct {
	x1 float64
	y1 float64
}

// Initialize a Cicle instance 1st approac
//var val Circle
//Initialize a Cicle instance 2nd approach

func CircleArea(c *Circle) float64 {
	c.x = 12
	c.y = 13
	return c.x * c.y
}

func RectangleArea(r Rectangle) float64 {
	return r.x1 + r.y1
}

func main() {
	val := Circle{4.0, 2.0}
	val1 := Rectangle{12.0, 11.0}
	//fmt.Println(CircleArea(val))
	fmt.Println(val)
	fmt.Println(CircleArea(&val))
	fmt.Println(val)
	fmt.Println(RectangleArea(val1))
}
