// Simple Shape Interface

package simpleshape

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func CalculateShapes(areas []Shape) {
	for _, area := range areas {
		fmt.Println(area.Area())
	}
}

func ShapeInterface() {
	fmt.Println("")
	fmt.Println("Shape Interface")

	rct := Rectangle{
		9.5,
		8.5,
	}
	circle := Circle{6.5}

	Shapes := []Shape{}
	Shapes = append(Shapes, rct)
	Shapes = append(Shapes, circle)
	CalculateShapes(Shapes)
}
