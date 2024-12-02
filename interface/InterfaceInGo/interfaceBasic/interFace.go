// Interface Baiscs

package interfaceingo

import "fmt"

type Speak interface {
	greet()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) greet() {
	fmt.Println(d.Name, "Says Woof")
}

func (c Cat) greet() {
	fmt.Println(c.Name, "Says Meow")
}

func GreetInterface() {
	fmt.Println("")
	fmt.Println("Greet Interface")
	var s Speak
	d1 := Dog{"Moti"}
	s = d1
	s.greet()

	c1 := Cat{"Moti"}
	s = c1
	s.greet()
}

type Person struct {
	Name string
}

type Name interface {
	Greeter()
}

func (p Person) Greeter() {
	fmt.Println("Hello ", p.Name)
}

func GreeterInterface() {
	fmt.Println("")
	p := Person{
		Name: "sagar",
	}
	p.Greeter()
}
