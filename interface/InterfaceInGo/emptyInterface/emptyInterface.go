package emptyinterface

import (
	"fmt"
)

func Describe(i interface{}) {
	fmt.Printf("The type of input is %T, and the value of input is '%v'\n", i, i)
}

func EmptyInterface() {
	fmt.Println("\nEmpty Interface")
	s := "Hey this is sagar"
	Describe(s)

	stru := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "sagar",
		lastName:  "parmar",
		age:       99,
	}
	Describe(stru)
}
