package main

import "fmt"

type Car struct {
	make   string
	model  string
	year   int
	engine Engine
}

type Engine struct {
	horsepower int
	Type       string
}

func (c *Car) displayInfo() {
	fmt.Printf("Car make is %s\nCar model is %s\nCar manufacturing year is %d\ncar engine  horsepower is %d\nCar engine type is %s\n", c.make, c.model, c.year, c.engine.horsepower, c.engine.Type)
}
func main() {
	vehicle := Car{
		make:  "tata",
		model: "xz+",
		year:  2021,
		engine: Engine{
			horsepower: 2000,
			Type:       "Diesel",
		},
	}
	vehicle.displayInfo()
}
