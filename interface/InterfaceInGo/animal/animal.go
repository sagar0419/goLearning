// Animal Sounds

package animal

import "fmt"

type AnimalSpeak interface {
	speak()
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}
type Cow struct {
	Name string
}

func (c Cat) speak() {
	fmt.Println(c.Name, "say meoww!!!")
}

func (c Cow) speak() {
	fmt.Println(c.Name, "say meoww!!!")
}

func (d Dog) speak() {
	fmt.Println(d.Name, "say meoww!!!")
}

func AnimalSpealGreet() {
	// var as string
	fmt.Println("")
	c := Cat{
		"Pinky",
	}
	d := Dog{
		"Golu",
	}
	s := Cow{
		"Hema",
	}
	c.speak()
	d.speak()
	s.speak()
}
