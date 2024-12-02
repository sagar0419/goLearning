package main

import (
	animal "github.com/sagar0419/goLearning/InterfaceInGo/animal"
	interfaceGo "github.com/sagar0419/goLearning/InterfaceInGo/interfaceBasic"
	reuse "github.com/sagar0419/goLearning/InterfaceInGo/reusablePrinter"
	simpleshape "github.com/sagar0419/goLearning/InterfaceInGo/simpleShape"
)

func main() {
	interfaceGo.GreetInterface()
	interfaceGo.GreeterInterface()
	animal.AnimalSpealGreet()
	simpleshape.ShapeInterface()
	reuse.ReusablePrinter()
}