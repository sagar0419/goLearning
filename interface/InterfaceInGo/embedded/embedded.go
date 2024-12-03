package embedded

import (
	"fmt"
)

type Flyer interface {
	Fly() string
}
type Swimmer interface {
	Swim() string
}

type Amphibian interface {
	Flyer
	Swimmer
}

type Duck struct {
	Name string
}

func (d Duck) Fly() string {
	return fmt.Sprintf("%s can fly.", d.Name)
}

func (d Duck) Swim() string {
	return fmt.Sprintf("%s can swim.", d.Name)
}

func AmphibianInterface() {
	fmt.Println("\nAmphibian Interface")
	duck := Duck{Name: "Donald the Duck"}

	var amphibian Amphibian = duck
	fmt.Println(amphibian.Fly())
	fmt.Println(amphibian.Swim())
}
