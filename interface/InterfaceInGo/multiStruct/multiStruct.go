package multistruct

import "fmt"

type Transport interface {
	Start() string
	Stop() string
}

type Car struct {
	Name string
}
type Bike struct {
	Name string
}
type Bus struct {
	Name string
}

func (c Car) Start() string {
	return fmt.Sprintf("the %s car has started", c.Name)
}
func (b Bike) Start() string {
	return fmt.Sprintf("the %s Bike has started", b.Name)
}
func (b Bus) Start() string {
	return fmt.Sprintf("the %s Bus has started", b.Name)

}

func (c Car) Stop() string {
	return fmt.Sprintf("the %s Car has Stopped", c.Name)
}
func (b Bike) Stop() string {
	return fmt.Sprintf("the %s Bike has Stopped", b.Name)
}
func (b Bus) Stop() string {
	return fmt.Sprintf("the %s Bus has Stopped", b.Name)
}

// Created a func that take slice of transport as input
func printMultiStructStart(s []Transport) {
	for _, allValues := range s {
		fmt.Println(allValues.Start())
	}
}
func printMultiStructStop(s []Transport) {
	for _, allValues := range s {
		fmt.Println(allValues.Stop())
	}
}

func MultiStruct() {
	fmt.Println("")
	fmt.Println("MultiStruct")
	fmt.Println("")
	c := Car{"White"}
	b := Bike{"Racing"}
	s := Bus{"Tourist"}

	// created slice of Transport
	multi := []Transport{}
	multi = append(multi, c)
	multi = append(multi, b)
	multi = append(multi, s)

	fmt.Println("Start Interface")
	printMultiStructStart(multi)
	fmt.Println("")
	fmt.Println("Stop Interface")
	printMultiStructStop(multi)
}
