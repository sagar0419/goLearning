package main

import "fmt"

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func (b *Book) applyDiscount(discount float64) {
	disc := (discount / 100) * b.Price
	b.Price -= disc
}

func (b Book) getBookInfo() {
	fmt.Printf("Book is '%s' and the author is '%s' is of %f\n", b.Title, b.Author, b.Price)
}

func main() {
	purchasedBook := Book{
		Title:  "The Superman",
		Author: "Warner Bros",
		Price:  999.99,
	}
	purchasedBook.applyDiscount(9)

	purchasedBook.getBookInfo()
}
