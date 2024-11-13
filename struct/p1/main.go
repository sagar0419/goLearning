package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := Person{
		firstName: "sagar",
		lastName:  "parmar",
		age:       32,
	}
	fmt.Println(p1)
}
