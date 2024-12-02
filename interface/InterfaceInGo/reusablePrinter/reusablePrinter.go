package reusableprinter

import "fmt"

type Printer interface {
	Details() string
}

type Book struct {
	Name   string
	Author string
}
type Car struct {
	Name         string
	Manufacturer string
}

func (b Book) Details() string {
	x := fmt.Sprintf("the name of the book is %s and the author of the book is %s", b.Name, b.Author)
	return x
}

func (c Car) Details() string {
	x := fmt.Sprintf("the name of the car is %s and the manufacturer of the Car is %s", c.Name, c.Manufacturer)
	return x
}

func reusableDetails(r []Printer) {
	for _, info := range r {
		fmt.Println(info.Details())
	}
}
func ReusablePrinter() {
	fmt.Println("")
	fmt.Println("Reusable Printer")
	b := Book{
		Name:   "ABC of Golang",
		Author: "Google",
	}
	c := Car{
		Name:         "Heliux",
		Manufacturer: "Toyota",
	}
	detail := []Printer{}
	detail = append(detail, b)
	detail = append(detail, c)
	reusableDetails(detail)
}
