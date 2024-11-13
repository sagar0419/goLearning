package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) area() {
	fmt.Printf("Area of rectangle is %d\n", (r.height * r.width))
}

func main() {
	RectangleArea := Rectangle{
		width:  40,
		height: 50,
	}
	RectangleArea.area()
}
