package main

import "fmt"

type Product struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

func Equal(Product1, Product2 Product) bool {
	return Product1 == Product2
}

func main() {
	Product1 := Product{
		Name:     "Dell",
		Category: "Laptop",
		Price:    88.5,
	}

	Product2 := Product{
		Name:     "Dell",
		Category: "Monitor",
		Price:    88,
	}
	areEqual := Equal(Product1, Product2)
	fmt.Println("Are the values equal???\n", areEqual)

	Product3 := Product{
		Name:     "lenovo",
		Category: "Laptop",
		Price:    70.5,
	}

	Product4 := Product{
		Name:     "lenovo",
		Category: "Laptop",
		Price:    70.5,
	}
	areEqual = Equal(Product3, Product4)
	fmt.Println("Are the values equal???\n", areEqual)
}
