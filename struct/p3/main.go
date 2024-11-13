package main

import "fmt"

type BankAccount struct {
	accountNumber string
	balance       float64
}

func (b *BankAccount) deposit(amount float64) {
	b.balance += amount
	fmt.Printf("The existing bank balance is %.2f\n The new bank balance is %.2f\n", amount, b.balance)
}

func main() {
	ba := BankAccount{
		accountNumber: "2017562",
		balance:       20.54,
	}
	ba.deposit(60)
}
