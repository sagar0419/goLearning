package main

import (
	animal "github.com/sagar0419/goLearning/InterfaceInGo/animal"
	emptyinterface "github.com/sagar0419/goLearning/InterfaceInGo/emptyInterface"
	interfaceGo "github.com/sagar0419/goLearning/InterfaceInGo/interfaceBasic"
	MultiStruct "github.com/sagar0419/goLearning/InterfaceInGo/multiStruct"
	reuse "github.com/sagar0419/goLearning/InterfaceInGo/reusablePrinter"
	simpleshape "github.com/sagar0419/goLearning/InterfaceInGo/simpleShape"
)

func main() {
	interfaceGo.GreetInterface()
	interfaceGo.GreeterInterface()
	animal.AnimalSpealGreet()
	simpleshape.ShapeInterface()
	reuse.ReusablePrinter()
	MultiStruct.MultiStruct()
	emptyinterface.EmptyInterface()
}
